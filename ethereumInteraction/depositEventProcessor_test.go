package ethereuminteraction

import (
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	ABI "github.com/matterinc/PlasmaBlockVerifier/contractABI"
	database "github.com/matterinc/PlasmaBlockVerifier/database"
)

func TestSingleDepositEvent(t *testing.T) {
	err := database.PurgeTestDatabase()
	if err != nil {
		t.Fatal(err)
	}
	db, err := database.OpenTestDatabase()
	if err != nil {
		t.Fatal(err)
	}
	addr := common.Address{}
	valueBI := big.NewInt(1000)
	indexBI := big.NewInt(1)
	depositEvent := &ABI.PlasmaParentDepositEvent{}
	depositEvent.From = addr
	depositEvent.Amount = valueBI
	depositEvent.DepositIndex = indexBI

	processor := NewDepositEventProcessor(db)
	success, err := processor.Process(depositEvent)
	if err != nil {
		t.Fatal(err)
	}
	if success != true {
		t.Fail()
	}
}

func TestDuplicateDepositEvent(t *testing.T) {
	err := database.PurgeTestDatabase()
	if err != nil {
		t.Fatal(err)
	}
	db, err := database.OpenTestDatabase()
	if err != nil {
		t.Fatal(err)
	}
	addr := common.Address{}
	valueBI := big.NewInt(1000)
	indexBI := big.NewInt(1)
	depositEvent := &ABI.PlasmaParentDepositEvent{}
	depositEvent.From = addr
	depositEvent.Amount = valueBI
	depositEvent.DepositIndex = indexBI

	processor := NewDepositEventProcessor(db)
	success, err := processor.Process(depositEvent)
	if err != nil {
		t.Fatal(err)
	}
	if success != true {
		t.Fail()
	}

	success, err = processor.Process(depositEvent)
	if err == nil {
		t.Fail()
	}
}
