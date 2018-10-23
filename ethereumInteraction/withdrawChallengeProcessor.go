package ethereuminteraction

import (
	"log"

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
			return false, err
		}
	}
	spendingBlock, spendingTxNumber, spendingInput, err := transaction.ParseUTXOnumber(spendingIndex)
	if err != nil {
		return false, err
	}
	log.Printf("Preparing challenge by using a block %v, tx %v, output %v \n", spendingBlock, spendingTxNumber, spendingInput)
	return true, nil
}
