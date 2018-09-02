package ethereuminteraction

import (
	"fmt"

	"github.com/dgraph-io/badger"
	"github.com/matterinc/PlasmaBlockVerifier/messageStructures"
	"github.com/matterinc/PlasmaCommons/transaction"
)

type WithdrawStartedProcessor struct {
	db *badger.DB
}

func NewWithdrawStartedProcessor(db *badger.DB) *WithdrawStartedProcessor {
	newInstance := &WithdrawStartedProcessor{db}
	return newInstance
}

func (p *WithdrawStartedProcessor) Process(event *messageStructures.WithdrawStartedInformation) (*messageStructures.WithdrawChallengeRequest, error) {
	// process only non-limbo for now
	if event.IsLimbo {
		return nil, nil
	}
	// utxoIndex := event.Index.Bytes()
	// // transaction.ParseUTXOindexNumberIntoDetails()
	// if len(utxoIndex) < 9 {
	// 	padding := make([]byte, 9-len(utxoIndex))
	// 	utxoIndex = append(padding, utxoIndex...)
	// }
	// blockNumber := binary.BigEndian.Uint32(utxoIndex[0:4])
	// txNumber := binary.BigEndian.Uint32(utxoIndex[4:8])
	// outputNumber := utxoIndex[8]
	details, err := transaction.ParseUTXOindexBigIntIntoDetails(event.Index)
	if err != nil {
		return nil, err
	}
	utxoIndex, err := transaction.CreateFullUTXOindexFromDetails(details.BlockNumber, details.TransactionNumber, details.OutputNumber, event.From, event.Amount)
	if err != nil {
		return nil, err
	}
	index := []byte{}
	index = append(index, []byte("utxo")...)
	index = append(index, utxoIndex[:]...)
	err = p.db.Update(func(txn *badger.Txn) error {
		_, err := txn.Get(index)
		if err != nil {
			return err
		}
		err = txn.Delete(index)
		if err != nil {
			return err
		}
		return nil
	})
	if err == nil {
		return nil, nil
	}
	var spendingIndex []byte
	shortIndex := transaction.PackUTXOnumber(details.BlockNumber, details.TransactionNumber, details.OutputNumber)
	lookupIndex := []byte("spend")
	lookupIndex = append(lookupIndex, shortIndex...)
	err = p.db.View(func(txn *badger.Txn) error {
		item, err := txn.Get(lookupIndex)
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
		return nil, err
	}
	spendingBlock, spendingTxNumber, spendingInput, err := transaction.ParseUTXOnumber(spendingIndex)
	if err != nil {
		return nil, err
	}
	fmt.Printf("Preparing challenge by using a block %i, tx %i, output %i", spendingBlock, spendingTxNumber, spendingInput)
	return nil, nil
}
