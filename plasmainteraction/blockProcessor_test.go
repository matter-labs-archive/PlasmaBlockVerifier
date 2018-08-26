package plasmainteraction

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"strconv"
	"testing"
	"time"

	"github.com/shamatar/go-plasma/block"

	"github.com/dgraph-io/badger"
	"github.com/ethereum/go-ethereum/common"
	database "github.com/matterinc/PlasmaBlockVerifier/database"
	"github.com/shamatar/go-plasma/transaction"
	"github.com/shamatar/go-plasma/types"
)

var testRecipientAccount = common.FromHex("0xf62803ffaddda373d44b10bf6bb404909be0e66b")
var testSenderAccount = common.FromHex("0x627306090abab3a6e1400e9345bc60c78a8bef57")
var senderPrivateKey = common.FromHex("0xc87509a1c067bbde78beb793e6fa76530b6382a4c0241e5e4a9ec0a0f44dc0d3")

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
	deposits, withdraws, err := processor.ProcessBlockWithoutCommitment(data)
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
	deposits, withdraws, err = processor.ProcessBlockWithoutCommitment(data)
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
	deposits, withdraws, err = processor.ProcessBlockWithoutCommitment(data)
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
			counter++
		}
		return nil
	})
	if counter != 1 {
		t.Fail()
	}

	historyCounter := 0
	err = db.View(func(txn *badger.Txn) error {
		it := txn.NewIterator(badger.DefaultIteratorOptions)
		defer it.Close()
		prefix := []byte("spend")
		for it.Seek(prefix); it.ValidForPrefix(prefix); it.Next() {
			historyCounter++
		}
		return nil
	})
	if historyCounter != 2 {
		t.Fail()
	}
}

func createTransferTransaction(blockNumber int, txNumberInBlock int, outputNumberInTransaction int, value string, testPrivateKey []byte) (*transaction.SignedTransaction, error) {
	bn := types.NewBigInt(int64(blockNumber))
	tn := types.NewBigInt(int64(txNumberInBlock))
	in := types.NewBigInt(int64(outputNumberInTransaction))
	v := types.NewBigInt(0)
	v.SetString(value, 10)
	input := &transaction.TransactionInput{}
	err := input.SetFields(bn, tn, in, v)
	if err != nil {
		return nil, err
	}
	bn = types.NewBigInt(0)
	to := common.Address{}
	copy(to[:], testRecipientAccount)
	output := &transaction.TransactionOutput{}
	err = output.SetFields(bn, to, v)
	if err != nil {
		return nil, err
	}

	inputs := []*transaction.TransactionInput{input}
	outputs := []*transaction.TransactionOutput{output}
	txType := transaction.TransactionTypeSplit
	tx, err := transaction.NewUnsignedTransaction(txType, inputs, outputs)
	emptyBytes := [32]byte{}
	signed, err := transaction.NewSignedTransaction(tx, []byte{0x00}, emptyBytes[:], emptyBytes[:])
	signed.Sign(testPrivateKey)
	return signed, nil
	// var b bytes.Buffer
	// i := io.Writer(&b)
	// err = signed.EncodeRLP(i)
	// if err != nil {
	// 	return nil, err
	// }
	// a := b.Bytes()
	// return a, nil
}

func TestLargeBlockApplication(t *testing.T) {
	blockSize := 100000
	blockNumber := rand.Uint32()
	amount := "100"
	_ = database.PurgeTestDatabase()
	db, _ := database.OpenTestDatabase()
	for i := 0; i < blockSize; i++ {
		index := transaction.PackUTXOnumber(blockNumber, uint32(i), 0)
		utxoIndex := append(UtxoIndexPrefix, testSenderAccount...)
		utxoIndex = append(utxoIndex, index...)
		v := types.NewBigInt(0)
		v.SetString(amount, 10)
		vBytes, _ := v.GetLeftPaddedBytes(32)
		utxoIndex = append(utxoIndex, vBytes...)
		_ = db.Update(func(txn *badger.Txn) error {
			err := txn.Set(utxoIndex, []byte{0x01})
			return err
		})
		if i%1000 == 0 {
			fmt.Println("Created " + strconv.Itoa(i) + " for test")
		}
	}
	counter := 0
	_ = db.View(func(txn *badger.Txn) error {
		it := txn.NewIterator(badger.DefaultIteratorOptions)
		defer it.Close()
		prefix := []byte("utxo")
		for it.Seek(prefix); it.ValidForPrefix(prefix); it.Next() {
			counter++
		}
		return nil
	})
	if counter != blockSize {
		t.Fail()
	}
	allTXes := make([]*transaction.SignedTransaction, blockSize)
	fmt.Println("Creating transactions block")
	for i := 0; i < blockSize; i++ {
		tx, err := createTransferTransaction(int(blockNumber), i, 0, amount, senderPrivateKey)
		if err != nil {
			t.Fatal(err)
		}
		allTXes[i] = tx
		if i%1000 == 0 {
			fmt.Println("Prepared transactions = " + strconv.Itoa(i))
		}
	}
	newBlockNumber := rand.Uint32()
	fmt.Println("Start to prepare a block")
	start := time.Now()
	block, err := block.NewBlock(newBlockNumber, allTXes, make([]byte, 32))
	elapsed := time.Since(start)
	blockPreparationTime := elapsed.Seconds()
	fmt.Printf("Block preparation taken %f seconds \n", blockPreparationTime)
	if err != nil {
		t.Fatal(err)
	}
	blockProcessor := NewBlockProcessor(db, 100, 100)
	fmt.Println("Applying parsed(!) block")
	start = time.Now()
	dRequests, wRequests, err := blockProcessor.ProcessParsedBlock(block)
	elapsed = time.Since(start)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(len(dRequests))
	fmt.Println(len(wRequests))
	fmt.Println("Verified " + strconv.Itoa(blockSize) + " succesfully")
	totalSpeed := float64(blockSize) / (elapsed.Seconds() + blockPreparationTime)
	fmt.Println("Total vefirication speed (including block validation, but not parsing) = " + fmt.Sprintf("%f", totalSpeed))

	txApplicationSpeed := float64(blockSize) / elapsed.Seconds()
	fmt.Println("TX application speed (no block preparation) = " + fmt.Sprintf("%f", txApplicationSpeed))

	counter = 0
	_ = db.View(func(txn *badger.Txn) error {
		it := txn.NewIterator(badger.DefaultIteratorOptions)
		defer it.Close()
		prefix := []byte("utxo")
		prefix = append(prefix, testRecipientAccount...)
		for it.Seek(prefix); it.ValidForPrefix(prefix); it.Next() {
			counter++
			// item := it.Item()
			// k := item.Key()
			// v, err := item.Value()
			// if err != nil {
			// 	return err
			// }
			// fmt.Printf("key=%s, value=%s\n", common.ToHex(k), common.ToHex(v))
		}
		return nil
	})
	if counter != blockSize {
		t.Fatal("Invalid number of UTXOs for new owner")
	}

	counter = 0
	_ = db.View(func(txn *badger.Txn) error {
		it := txn.NewIterator(badger.DefaultIteratorOptions)
		defer it.Close()
		prefix := []byte("utxo")
		prefix = append(prefix, testSenderAccount...)
		for it.Seek(prefix); it.ValidForPrefix(prefix); it.Next() {
			counter++
			// item := it.Item()
			// k := item.Key()
			// v, err := item.Value()
			// if err != nil {
			// 	return err
			// }
			// fmt.Printf("key=%s, value=%s\n", common.ToHex(k), common.ToHex(v))
		}
		return nil
	})
	if counter != 0 {
		t.Fatal("Invalid number of UTXOs left for old owner")
	}
}

func TestLargeBlockApplicationWithWithdraws(t *testing.T) {
	blockSize := 100000
	blockNumber := rand.Uint32()
	amount := "100"
	withdrawChance := 10
	_ = database.PurgeTestDatabase()
	db, _ := database.OpenTestDatabase()
	withdrawsToExpect := 0
	for i := 0; i < blockSize; i++ {
		if int(rand.Uint32())%100 <= withdrawChance {
			withdrawsToExpect++
			continue
		}
		index := transaction.PackUTXOnumber(blockNumber, uint32(i), 0)
		utxoIndex := append(UtxoIndexPrefix, testSenderAccount...)
		utxoIndex = append(utxoIndex, index...)
		v := types.NewBigInt(0)
		v.SetString(amount, 10)
		vBytes, _ := v.GetLeftPaddedBytes(32)
		utxoIndex = append(utxoIndex, vBytes...)
		_ = db.Update(func(txn *badger.Txn) error {
			err := txn.Set(utxoIndex, []byte{0x01})
			return err
		})
		if i%1000 == 0 {
			fmt.Println("Created " + strconv.Itoa(i) + " for test")
		}
	}
	counter := 0
	_ = db.View(func(txn *badger.Txn) error {
		it := txn.NewIterator(badger.DefaultIteratorOptions)
		defer it.Close()
		prefix := []byte("utxo")
		for it.Seek(prefix); it.ValidForPrefix(prefix); it.Next() {
			counter++
		}
		return nil
	})
	if counter != blockSize-withdrawsToExpect {
		t.Fatal("Invalid number of created UTXOs")
	}
	allTXes := make([]*transaction.SignedTransaction, blockSize)
	fmt.Println("Creating transactions block")
	for i := 0; i < blockSize; i++ {
		tx, err := createTransferTransaction(int(blockNumber), i, 0, amount, senderPrivateKey)
		if err != nil {
			t.Fatal(err)
		}
		allTXes[i] = tx
		if i%1000 == 0 {
			fmt.Println("Prepared transactions = " + strconv.Itoa(i))
		}
	}
	newBlockNumber := rand.Uint32()
	fmt.Println("Start to prepare a block")
	start := time.Now()
	block, err := block.NewBlock(newBlockNumber, allTXes, make([]byte, 32))
	elapsed := time.Since(start)
	blockPreparationTime := elapsed.Seconds()
	fmt.Printf("Block preparation taken %f seconds \n", blockPreparationTime)
	if err != nil {
		t.Fatal(err)
	}
	blockProcessor := NewBlockProcessor(db, 100, 100)
	fmt.Println("Applying parsed(!) block")
	start = time.Now()
	dRequests, wRequests, err := blockProcessor.ProcessParsedBlock(block)
	elapsed = time.Since(start)
	if err != nil {
		t.Fatal(err)
	}
	if withdrawsToExpect != len(wRequests) {
		t.Fatal("Invalid number of withdraw requests")
	}
	fmt.Println(len(dRequests))
	fmt.Println(len(wRequests))
	fmt.Println("Verified " + strconv.Itoa(blockSize) + " succesfully")
	totalSpeed := float64(blockSize) / (elapsed.Seconds() + blockPreparationTime)
	fmt.Println("Total vefirication speed (including block validation, but not parsing) = " + fmt.Sprintf("%f", totalSpeed))

	txApplicationSpeed := float64(blockSize) / elapsed.Seconds()
	fmt.Println("TX application speed (no block preparation) = " + fmt.Sprintf("%f", txApplicationSpeed))

	counter = 0
	_ = db.View(func(txn *badger.Txn) error {
		it := txn.NewIterator(badger.DefaultIteratorOptions)
		defer it.Close()
		prefix := []byte("utxo")
		prefix = append(prefix, testRecipientAccount...)
		for it.Seek(prefix); it.ValidForPrefix(prefix); it.Next() {
			counter++
			// item := it.Item()
			// k := item.Key()
			// v, err := item.Value()
			// if err != nil {
			// 	return err
			// }
			// fmt.Printf("key=%s, value=%s\n", common.ToHex(k), common.ToHex(v))
		}
		return nil
	})
	if counter != blockSize-withdrawsToExpect {
		t.Fatal("Invalid number of UTXOs for new owner")
	}

	counter = 0
	_ = db.View(func(txn *badger.Txn) error {
		it := txn.NewIterator(badger.DefaultIteratorOptions)
		defer it.Close()
		prefix := []byte("utxo")
		prefix = append(prefix, testSenderAccount...)
		for it.Seek(prefix); it.ValidForPrefix(prefix); it.Next() {
			counter++
			// item := it.Item()
			// k := item.Key()
			// v, err := item.Value()
			// if err != nil {
			// 	return err
			// }
			// fmt.Printf("key=%s, value=%s\n", common.ToHex(k), common.ToHex(v))
		}
		return nil
	})
	if counter != 0 {
		t.Fatal("Invalid number of UTXOs left for old owner")
	}
}
