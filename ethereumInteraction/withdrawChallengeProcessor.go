package ethereuminteraction

import (
	"fmt"

	"github.com/dgraph-io/badger"

	ABI "github.com/matterinc/PlasmaBlockVerifier/contractABI"
	"github.com/matterinc/PlasmaBlockVerifier/messageStructures"
	"github.com/matterinc/PlasmaCommons/transaction"
)

type WithdrawChallengeProcessor struct {
	plasmaParent *ABI.PlasmaParent
	db           *badger.DB
}

func NewWithdrawChallengeProcessor(plasmaParent *ABI.PlasmaParent, db *badger.DB) *WithdrawChallengeProcessor {
	newInstance := &WithdrawChallengeProcessor{plasmaParent, db}
	return newInstance
}

func (p *WithdrawChallengeProcessor) Process(challengeRequest *messageStructures.WithdrawChallengeRequest) (bool, error) {
	var utxoIndex [transaction.UTXOIndexLength]byte
	copy(challengeRequest.UtxoIndex[4:], utxoIndex[:])
	details := transaction.ParseIndexIntoUTXOdetails(utxoIndex)
	fmt.Println("Processign withdraw for " + fmt.Sprintln(details))
	shortIndex := transaction.PackUTXOnumber(details.BlockNumber, details.TransactionNumber, details.OutputNumber)
	index := []byte{}
	index = append(index, []byte("spend")...)
	index = append(index, shortIndex...)
	var spendingIndex []byte
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
		return false, err
	}
	spendingBlock, spendingTxNumber, spendingInput, err := transaction.ParseUTXOnumber(spendingIndex)
	if err != nil {
		return false, err
	}
	fmt.Printf("Preparing challenge by using a block %i, tx %i, output %i", spendingBlock, spendingTxNumber, spendingInput)
	return true, nil
}
