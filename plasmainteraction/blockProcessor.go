package plasmainteraction

import (
	"encoding/binary"
	"fmt"
	"math/big"
	"sync"

	"github.com/dgraph-io/badger"
	common "github.com/ethereum/go-ethereum/common"
	"github.com/shamatar/go-plasma/block"
	"github.com/shamatar/go-plasma/transaction"
)

var (
	UtxoIndexPrefix      = []byte("utxo")
	BlockNumberKey       = []byte("blockNumber")
	TransactionNumberKey = []byte("txNumber")
	SpendingIndexKey     = []byte("spend")
)

type TransactionPayload struct {
	txNumber uint32
	tx       *transaction.SignedTransaction
}

type DepositIndexCheckoutRequest struct {
	index    *big.Int
	value    *big.Int
	address  common.Address
	operator common.Address
}

type WithdrawChallengeRequest struct {
	utxoIndex                []byte
	spendingTransactionIndex []byte
}

type PreprocessedTransactionPayload struct {
	txNumber               uint32
	keysToDelete           [][]byte
	keysToWrite            [][]byte
	spendingIndexesToWrite [][2][]byte
	depositCheckoutRequest *DepositIndexCheckoutRequest
}

type ResultPayload struct {
	txNumber                 uint32
	result                   bool
	nonrecoverableError      error
	depositIndexRequest      *DepositIndexCheckoutRequest
	withdrawChallengeRequest *WithdrawChallengeRequest
}

type BlockProcessor struct {
	db                 *badger.DB
	concurrencyLimit   int
	sliceSize          int
	concurrencyChannel chan bool
	resultsChannel     chan bool
}

func NewBlockProcessor(db *badger.DB, concurrencyLimit int, sliceSize int) *BlockProcessor {
	concurrencyChannel := make(chan bool, concurrencyLimit)
	resultsChannel := make(chan bool, concurrencyLimit)
	newInstance := &BlockProcessor{db, concurrencyLimit, sliceSize, concurrencyChannel, resultsChannel}
	return newInstance
}

func (p *BlockProcessor) ProcessBlock(blockBytes []byte) ([]*DepositIndexCheckoutRequest, []*WithdrawChallengeRequest, error) {
	parsedBlock, err := block.NewBlockFromBytes(blockBytes)
	if err != nil {
		return nil, nil, err
	}
	err = parsedBlock.Validate()
	if err != nil {
		return nil, nil, err
	}
	blockNumberBytes := parsedBlock.BlockHeader.BlockNumber[:]
	blockNumber := binary.BigEndian.Uint32(blockNumberBytes)
	numTransactions := len(parsedBlock.Transactions)
	numSlices := numTransactions / p.sliceSize
	if numTransactions%p.sliceSize != 0 {
		numSlices++
	}
	sliceSize := p.sliceSize
	resChannels := make([]<-chan []ResultPayload, numSlices)
	for i := 0; i < numSlices; i++ {
		minTxNumber := uint32(0)
		maxTxNumber := uint32(0)
		if (i+1)*sliceSize < numTransactions {
			minTxNumber = uint32(i * sliceSize)
			maxTxNumber = uint32((i+1)*sliceSize) - 1
		} else {
			minTxNumber = uint32(i * sliceSize)
			maxTxNumber = uint32(numTransactions) - 1
		}
		currentSlice := parsedBlock.Transactions[minTxNumber : maxTxNumber+1]
		resChannels[i] = p.startWorker(minTxNumber, maxTxNumber, currentSlice, blockNumber)
	}

	flattenedResults := merge(resChannels)
	depositChecks := []*DepositIndexCheckoutRequest{}
	withdrawChecks := []*WithdrawChallengeRequest{}
	for result := range flattenedResults {
		if result.result == false {
			fmt.Println("Unsuccesfull result")
		}
		if result.nonrecoverableError != nil {
			panic(result.nonrecoverableError)
		}
		if result.depositIndexRequest != nil {
			depositChecks = append(depositChecks, result.depositIndexRequest)
		}
		if result.withdrawChallengeRequest != nil {
			withdrawChecks = append(withdrawChecks, result.withdrawChallengeRequest)
		}
	}
	return depositChecks, withdrawChecks, nil
}

func merge(inputs []<-chan []ResultPayload) <-chan ResultPayload {
	var wg sync.WaitGroup
	out := make(chan ResultPayload)

	// Start an output goroutine for each input channel in cs.  output
	// copies values from c to out until c is closed, then calls wg.Done.
	output := func(c <-chan []ResultPayload) {
		for n := range c {
			for _, r := range n {
				out <- r
			}
		}
		wg.Done()
	}
	wg.Add(len(inputs))
	for _, c := range inputs {
		go output(c)
	}

	// Start a goroutine to close out once all the output goroutines are
	// done.  This must start after the wg.Add call.
	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}

func (p *BlockProcessor) startWorker(startIndex, endIndex uint32, txes []*transaction.SignedTransaction, blockNumber uint32) <-chan []ResultPayload {
	finished := make(chan []ResultPayload)
	go func() {
		defer close(finished)
		preprocessed, err := p.PreprocessTransactions(startIndex, endIndex, txes, blockNumber)
		if err != nil {
			panic(err)
		}
		result, err := p.ProcessTransactionsSlice(preprocessed)
		if err != nil {
			result, err := p.ProcessSliceInSequence(preprocessed)
			if err != nil {
				panic(err)
			}
			finished <- result
		} else {
			finished <- result
		}
	}()
	return finished
}

// preprocessing step does the following:
// 1. check for deposit transactions that there is a deposit record in database
// 2. parse transaction signature
// 3. make an array of UTXO indexes that should be marked as spent and write transaction index
// 4. make an array of fresh UTXOs
func (p *BlockProcessor) PreprocessTransactions(startIndex, endIndex uint32, txes []*transaction.SignedTransaction, blockNumber uint32) ([]*PreprocessedTransactionPayload, error) {
	numTxes := endIndex - startIndex + 1
	preprocessedPayloads := make([]*PreprocessedTransactionPayload, numTxes)
	// keysToDelete := make([][][]byte, numTxes)
	// keysToAdd := make([][][]byte, numTxes)
	// spendingIndexesToWrite := make([][][2][]byte, numTxes)
	for i := 0; i < int(numTxes); i++ {
		tx := txes[i]
		transactionNumber := startIndex + uint32(i)
		transactionType := tx.UnsignedTransaction.TransactionType
		payload := &PreprocessedTransactionPayload{transactionNumber, nil, nil, nil, nil}
		// iterate over inputs, add records that should be removed from UTXO set and added to spending index
		if transactionType[0] == transaction.TransactionTypeFund {
			request, err := p.MakeCheckDepositRecord(tx)
			if err != nil {
				return nil, err
			}
			payload.depositCheckoutRequest = request
		} else {
			numInputs := len(tx.UnsignedTransaction.Inputs)
			forDelete := make([][]byte, numInputs)
			forIndex := make([][2][]byte, numInputs)
			for j := range tx.UnsignedTransaction.Inputs {
				inputIndex, err := transaction.CreateCorrespondingUTXOIndexForInput(tx, j)
				if err != nil {
					return nil, err
				}
				prefixedInputIndex := append(UtxoIndexPrefix, inputIndex[:]...)
				spendingHistoryIndex, err := transaction.CreateShortUTXOIndexForInput(tx, j)
				if err != nil {
					return nil, err
				}
				spendingHistoryKey := append(SpendingIndexKey, spendingHistoryIndex...)
				spendingHistoryValue := transaction.PackUTXOnumber(blockNumber, transactionNumber, uint8(j))
				forDelete[j] = prefixedInputIndex
				spendingIndex := [2][]byte{}
				spendingIndex[0] = spendingHistoryKey
				spendingIndex[1] = spendingHistoryValue
				forIndex[j] = spendingIndex
			}
			payload.keysToDelete = forDelete
			payload.spendingIndexesToWrite = forIndex
		}

		numOutputs := len(tx.UnsignedTransaction.Outputs)
		forAdd := make([][]byte, numOutputs)
		// now we should iterate over the outputs to create new UTXOs in an index
		for j := range tx.UnsignedTransaction.Outputs {
			newUTXOindex, err := transaction.CreateUTXOIndexForOutput(tx, blockNumber, transactionNumber, j)
			if err != nil {
				return nil, err
			}
			fullUTXOindex := append(UtxoIndexPrefix, newUTXOindex[:]...)
			forAdd[j] = fullUTXOindex
		}
		payload.keysToWrite = forAdd
		preprocessedPayloads[i] = payload
	}
	return preprocessedPayloads, nil
}

// during the block processing parse database for deposits during parsing
func (p *BlockProcessor) MakeCheckDepositRecord(tx *transaction.SignedTransaction) (*DepositIndexCheckoutRequest, error) {
	depositIndexBytes := tx.UnsignedTransaction.Inputs[0].Value[:]
	depositIndex := big.NewInt(0).SetBytes(depositIndexBytes)
	depositFor := tx.UnsignedTransaction.Outputs[0].GetToAddress()
	depositAmount := big.NewInt(0).SetBytes(tx.UnsignedTransaction.Outputs[0].Value[:])
	sender, err := tx.GetFrom()
	if err != nil {
		return nil, err
	}
	depositForCast := common.Address{}
	copy(depositForCast[:], depositFor[:])
	operatorCast := common.Address{}
	copy(operatorCast[:], sender[:])
	request := &DepositIndexCheckoutRequest{depositIndex, depositAmount, depositForCast, operatorCast}
	return request, nil
}

func (p *BlockProcessor) ProcessTransactionsSlice(preprocessed []*PreprocessedTransactionPayload) ([]ResultPayload, error) {
	results := make([]ResultPayload, len(preprocessed))
	txn := p.db.NewTransaction(true)
	defer txn.Discard()
	for i, payload := range preprocessed {
		// process either a deposit transaction or work with UTXO indexes
		if payload.depositCheckoutRequest != nil {
			results[i] = ResultPayload{payload.txNumber, true, nil, payload.depositCheckoutRequest, nil}
		} else {
			for _, toDelete := range payload.keysToDelete {
				err := txn.Delete(toDelete)
				if err == badger.ErrTxnTooBig {
					err := txn.Commit(nil)
					if err != nil {
						return nil, err
					}
					txn = p.db.NewTransaction(true)
					err = txn.Delete(toDelete)
					if err != nil {
						return nil, err
					}
				} else if err != nil {
					return nil, err
				}
			}

			for _, toIndex := range payload.spendingIndexesToWrite {
				err := txn.Set(toIndex[0], toIndex[1])
				if err == badger.ErrTxnTooBig {
					err := txn.Commit(nil)
					if err != nil {
						return nil, err
					}
					txn = p.db.NewTransaction(true)
					err = txn.Set(toIndex[0], toIndex[1])
					if err != nil {
						return nil, err
					}
				} else if err != nil {
					return nil, err
				}
			}
			res := ResultPayload{payload.txNumber, true, nil, nil, nil}
			results[i] = res
		}

		// process new UTXOs
		for _, toAdd := range payload.keysToWrite {
			err := txn.Set(toAdd, []byte{0x01})
			if err == badger.ErrTxnTooBig {
				err := txn.Commit(nil)
				if err != nil {
					return nil, err
				}
				txn = p.db.NewTransaction(true)
				err = txn.Set(toAdd, []byte{0x01})
				if err != nil {
					return nil, err
				}
			} else if err != nil {
				return nil, err
			}
		}
	}
	err := txn.Commit(nil)
	if err != nil {
		return nil, err
	}
	return results, nil
}

func (p *BlockProcessor) ProcessSliceInSequence(preprocessed []*PreprocessedTransactionPayload) ([]ResultPayload, error) {
	results := make([]ResultPayload, len(preprocessed))
	return results, nil
}
