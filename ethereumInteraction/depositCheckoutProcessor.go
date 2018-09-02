package ethereuminteraction

import (
	"bytes"
	"errors"
	"fmt"
	"math/big"

	"github.com/matterinc/PlasmaBlockVerifier/messageStructures"

	"github.com/dgraph-io/badger"
	"github.com/ethereum/go-ethereum/common"
	ABI "github.com/matterinc/PlasmaBlockVerifier/contractABI"
)

var DepositIndexPrefix = []byte("deposit")

type DepositCheckoutProcessor struct {
	db           *badger.DB
	plasmaParent *ABI.PlasmaParent
	blockStorage *ABI.PlasmaBlockStorage
}

func NewDepositCheckoutProcessor(db *badger.DB, plasmaParent *ABI.PlasmaParent, blockStorage *ABI.PlasmaBlockStorage) *DepositCheckoutProcessor {
	newInstance := &DepositCheckoutProcessor{db, plasmaParent, blockStorage}
	return newInstance
}

func (p *DepositCheckoutProcessor) Process(checkoutRequest *messageStructures.DepositIndexCheckoutRequest) (bool, error) {
	// depositFrom := event.From
	depositIndex := checkoutRequest.Index
	fmt.Println("Processing deposit for index " + depositIndex.String())
	// depositAmount := event.Amount
	depositIndexBytes, err := makeDepositIndex(depositIndex)
	if err != nil {
		return false, err
	}
	depositIndexKey := []byte{}
	depositIndexKey = append(depositIndexKey, DepositIndexPrefix...)
	depositIndexKey = append(depositIndexKey, depositIndexBytes...)
	err = p.db.View(func(txn *badger.Txn) error {
		_, err := txn.Get(depositIndexKey)
		if err == nil {
			return errors.New("Duplicate deposit")
		}
		return nil
	})
	if err != nil {
		return false, err
	}

	// check the ethereum for an existence of the deposit and correspondance

	isValidSigner, err := p.blockStorage.CanSignBlocks(nil, checkoutRequest.Operator)
	if !isValidSigner || err != nil {
		return false, err
	}

	record, err := p.plasmaParent.DepositRecords(nil, depositIndex)
	if err != nil {
		return false, err
	}

	if record.Amount.Cmp(checkoutRequest.Value) != 0 {
		return false, errors.New("Deposit amount mismatch")
	}

	if bytes.Compare(record.From[:], checkoutRequest.Address[:]) != 0 {
		return false, errors.New("Deposit owner mismatch")
	}

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
