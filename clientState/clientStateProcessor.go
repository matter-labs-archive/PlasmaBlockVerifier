package clientState

import (
	"encoding/binary"
	"errors"
	"sync"

	"github.com/dgraph-io/badger"
)

type ClientStateProcessor struct {
	db    *badger.DB
	mutex *sync.Mutex
}

var lastEthereumBlockKey = []byte("ethblock")
var lastPlasmaBlockKey = []byte("plasmablock")
var isByzantineKey = []byte("byzantine")

var state *ClientStateProcessor = nil

func InitStateProcessor(db *badger.DB) {
	mutex := new(sync.Mutex)
	state = &ClientStateProcessor{db, mutex}
}

func GetLastEthereumBlock() uint64 {
	if state == nil {
		return 0
	}
	var value []byte
	err := state.db.View(func(txn *badger.Txn) error {
		item, err := txn.Get(lastEthereumBlockKey)
		if err != nil {
			return err
		}
		val, err := item.Value()
		if err != nil {
			return err
		}
		value = val
		return nil
	})
	if err != nil {
		return 0
	}
	blockNumber := binary.BigEndian.Uint64(value)
	return blockNumber
}

func SetLastEthereumBlock(blockNumber uint64) error {
	if state == nil {
		return errors.New("State processor is not initialized")
	}
	state.mutex.Lock()
	defer state.mutex.Unlock()
	value := make([]byte, 8)
	binary.BigEndian.PutUint64(value, blockNumber)
	err := state.db.Update(func(txn *badger.Txn) error {
		err := txn.Set(lastEthereumBlockKey, value)
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return err
	}
	return nil
}
