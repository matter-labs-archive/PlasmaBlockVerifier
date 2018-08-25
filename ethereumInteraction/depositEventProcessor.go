package ethereuminteraction

import (
	"errors"
	"math/big"

	"github.com/dgraph-io/badger"
	"github.com/ethereum/go-ethereum/common"
	ABI "github.com/matterinc/PlasmaBlockVerifier/contractABI"
)

var DepositIndexPrefix = []byte("deposit")

type DepositEventProcessor struct {
	db *badger.DB
}

func NewDepositEventProcessor(db *badger.DB) *DepositEventProcessor {
	newInstance := &DepositEventProcessor{db}
	return newInstance
}

func (p *DepositEventProcessor) Process(event *ABI.PlasmaParentDepositEvent) (bool, error) {
	// depositFrom := event.From
	depositIndex := event.DepositIndex
	// depositAmount := event.Amount
	depositIndexBytes, err := makeDepositIndex(depositIndex)
	if err != nil {
		return false, err
	}
	depositIndexKey := []byte{}
	depositIndexKey = append(depositIndexKey, DepositIndexPrefix...)
	depositIndexKey = append(depositIndexKey, depositIndexBytes...)
	err = p.db.Update(func(txn *badger.Txn) error {
		_, err := txn.Get(depositIndexKey)
		if err == nil {
			return errors.New("Duplicate deposit")
		}
		err = txn.Set(depositIndexKey, []byte{0x01})
		return err
	})
	if err != nil {
		return false, err
	}
	return true, nil
}

func makeDepositIndex(depositIndex *big.Int) ([]byte, error) {
	beBytes := depositIndex.Bytes()
	if len(beBytes) > 32 {
		return nil, errors.New("Invalid index length")
	}
	return common.LeftPadBytes(beBytes, 32), nil
}