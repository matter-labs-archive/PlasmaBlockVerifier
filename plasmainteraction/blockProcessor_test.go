package plasmainteraction

import (
	"io/ioutil"
	"testing"

	"github.com/dgraph-io/badger"
	database "github.com/matterinc/PlasmaBlockVerifier/database"
)

func TestBlockProcessor(t *testing.T) {
	err := database.PurgeTestDatabase()
	if err != nil {
		t.Fatal(err)
	}
	db, err := database.OpenTestDatabase()
	if err != nil {
		t.Fatal(err)
	}

	processor := NewBlockProcessor(db, 1, 100)
	data, err := ioutil.ReadFile("testBlocks/1")
	deposits, withdraws, err := processor.ProcessBlock(data)
	if err != nil {
		t.Fatal(err)
	}
	if len(deposits) != 1 {
		t.Fail()
	}
	if len(withdraws) != 0 {
		t.Fail()
	}

	data, err = ioutil.ReadFile("testBlocks/2")
	deposits, withdraws, err = processor.ProcessBlock(data)
	if err != nil {
		t.Fatal(err)
	}
	if len(deposits) != 0 {
		t.Fail()
	}
	if len(withdraws) != 0 {
		t.Fail()
	}

	data, err = ioutil.ReadFile("testBlocks/3")
	deposits, withdraws, err = processor.ProcessBlock(data)
	if err != nil {
		t.Fatal(err)
	}
	if len(deposits) != 0 {
		t.Fail()
	}
	if len(withdraws) != 0 {
		t.Fail()
	}
	counter := 0
	err = db.View(func(txn *badger.Txn) error {
		it := txn.NewIterator(badger.DefaultIteratorOptions)
		defer it.Close()
		prefix := []byte("utxo")
		for it.Seek(prefix); it.ValidForPrefix(prefix); it.Next() {
			// item := it.Item()
			// k := item.Key()
			// v, err := item.Value()
			// if err != nil {
			// 	return err
			// }
			// fmt.Printf("key=%s, value=%s\n", k, v)
			counter++
		}
		return nil
	})
	if counter != 1 {
		t.Fail()
	}
}
