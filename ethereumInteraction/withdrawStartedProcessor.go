package ethereuminteraction

import (
	"log"

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
	blockNumber := event.BlockNumber
	transactionNumber := event.TransactionNumber
	outputNumber := event.OutputNumber
	utxoIndex, err := transaction.CreateFullUTXOindexFromDetails(blockNumber, transactionNumber, outputNumber, event.From, event.Amount)
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
	shortIndex := transaction.PackUTXOnumber(blockNumber, transactionNumber, outputNumber)
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
	log.Printf("Preparing challenge by using a block %i, tx %i, output %i", spendingBlock, spendingTxNumber, spendingInput)
	request := &messageStructures.WithdrawChallengeRequest{event.PartialHash, shortIndex, spendingIndex}
	// request := &messageStructures.WithdrawChallengeRequest{event.PartialHash, blockNumber, transactionNumber, outputNumber, spendingBlock, spendingTxNumber, spendingInput, true, shortIndex, spendingIndex}
	return request, nil
}
