package ethereuminteraction

import (
	"bytes"
	"io"
	"log"

	"github.com/dgraph-io/badger"

	blockdownloader "github.com/matterinc/PlasmaBlockVerifier/blockdownloader"
	ABI "github.com/matterinc/PlasmaBlockVerifier/contractABI"
	"github.com/matterinc/PlasmaBlockVerifier/messageStructures"
	"github.com/matterinc/PlasmaCommons/block"
	"github.com/matterinc/PlasmaCommons/transaction"
)

type WithdrawChallengeProcessor struct {
	plasmaParent *ABI.PlasmaParent
	db           *badger.DB
}

type WithdrawChallengeResult struct {
	AlreadyChallenged bool
	Error             error
	ExitRecordHash    [22]byte
	ProofData         []byte
	TxData            []byte
	BlockNumber       uint32
	InputNumber       uint8
}

func NewWithdrawChallengeProcessor(plasmaParent *ABI.PlasmaParent, db *badger.DB) *WithdrawChallengeProcessor {
	newInstance := &WithdrawChallengeProcessor{plasmaParent, db}
	return newInstance
}

func (p *WithdrawChallengeProcessor) Process(challengeRequest *messageStructures.WithdrawChallengeRequest) chan WithdrawChallengeResult {
	resultChannel := make(chan WithdrawChallengeResult, 1)
	go func() {
		result := WithdrawChallengeResult{}
		var spendingIndex []byte
		if challengeRequest.SpendingTransactionIndex != nil {
			// challenge came from block processing, so there have existed an exit event before
			spendingIndex = challengeRequest.SpendingTransactionIndex
		} else {
			var utxoIndex [transaction.UTXOIndexLength]byte
			copy(utxoIndex[:], challengeRequest.UtxoIndex[4:])
			details := transaction.ParseIndexIntoUTXOdetails(utxoIndex)
			log.Printf("Processign withdraw challenge for %v", details)
			shortIndex := transaction.PackUTXOnumber(details.BlockNumber, details.TransactionNumber, details.OutputNumber)
			index := []byte{}
			index = append(index, []byte("spend")...)
			index = append(index, shortIndex...)
			// challenge came from the exit event processor, so there was a spend and then an exit event
			err := p.db.View(func(txn *badger.Txn) error {
				item, err := txn.Get(index)
				if err != nil {
					return err
				}
				value, err := item.Value()
				if err != nil {
					return err
				}
				spendingIndex = value
				return nil
			})
			if err != nil {
				result.Error = err
				resultChannel <- result
				close(resultChannel)
				return
			}
		}
		var exitRecordHash []byte
		// either we already know the exit record that has to be challenged,
		// or need to get one from Ethereum
		if challengeRequest.PartialHash != nil {
			exitRecordHash = challengeRequest.PartialHash
		} else {

		}
		spendingBlock, spendingTxNumber, spendingInput, err := transaction.ParseUTXOnumber(spendingIndex)
		if err != nil {
			result.Error = err
			resultChannel <- result
			close(resultChannel)
			return
		}
		blockDownloader := blockdownloader.NewBlockDownloader()
		blockChannel := blockDownloader.GetBlock(spendingBlock)
		blockBytes := <-blockChannel
		var exitRecordHashFixedLength [22]byte
		copy(exitRecordHashFixedLength[:], exitRecordHash)
		exitRecord, err := p.plasmaParent.ExitRecords(nil, exitRecordHashFixedLength)
		if !exitRecord.IsValid {
			result.AlreadyChallenged = true
			resultChannel <- result
			close(resultChannel)
			return
		}
		parsedBlock, err := block.NewBlockFromBytes(blockBytes)
		if err != nil {
			result.Error = err
			resultChannel <- result
			close(resultChannel)
			return
		}
		proof, err := parsedBlock.MerkleTree.ProvideBinaryProof(int(spendingTxNumber))
		tx := parsedBlock.Transactions[int(spendingTxNumber)]
		var b bytes.Buffer
		i := io.Writer(&b)
		err = tx.EncodeRLP(i)
		if err != nil {
			result.Error = err
			resultChannel <- result
			close(resultChannel)
			return
		}
		rawTransaction := b.Bytes()
		log.Printf("Prepared challenge by using a block %v, tx %v, output %v \n", spendingBlock, spendingTxNumber, spendingInput)
		result.ExitRecordHash = exitRecordHashFixedLength
		result.ProofData = proof
		result.TxData = rawTransaction
		result.BlockNumber = spendingBlock
		result.InputNumber = spendingInput
		resultChannel <- result
		close(resultChannel)
		return
	}()
	return resultChannel

}
