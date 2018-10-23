package restAPI

import (
	"encoding/binary"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/dgraph-io/badger"
	"github.com/ethereum/go-ethereum/common"
	"github.com/matterinc/PlasmaCommons/transaction"
)

type ListUTXOsHandler struct {
	db *badger.DB
}

type listUTXOsRequest struct {
	For               string `json:"for"`
	BlockNumber       int    `json:"blockNumber"`
	TransactionNumber int    `json:"transactionNumber"`
	OutputNumber      int    `json:"outputNumber"`
	Limit             int    `json:"limit,omitempty"`
}

type singleUTXOdetails struct {
	BlockNumber       int    `json:"blockNumber"`
	TransactionNumber int    `json:"transactionNumber"`
	OutputNumber      int    `json:"outputNumber"`
	Value             string `json:"value"`
}

type listUTXOsResponse struct {
	Error bool                `json:"error"`
	UTXOs []singleUTXOdetails `json:"utxos"`
}

func NewListUTXOsHandler(db *badger.DB) *ListUTXOsHandler {
	return &ListUTXOsHandler{db}
}

func (p *ListUTXOsHandler) handler(w http.ResponseWriter, r *http.Request) {
	var requestJSON listUTXOsRequest
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		writeEmptyUTXOResponse(w)
		return
	}
	err = json.Unmarshal(body, &requestJSON)
	if err != nil {
		writeEmptyUTXOResponse(w)
		return
	}

	forBytes := common.FromHex(requestJSON.For)
	address := common.Address{}
	copy(address[:], forBytes)
	blockNumber := uint32(requestJSON.BlockNumber)
	transactionNumber := uint32(requestJSON.TransactionNumber)
	outputNumber := uint8(requestJSON.OutputNumber)
	limit := 50
	if requestJSON.Limit != 0 {
		limit = requestJSON.Limit
	}
	if limit > 100 {
		limit = 100
	}
	// limit := 0
	utxos, err := p.getUTXOsForAddress(address, blockNumber, transactionNumber, outputNumber, limit)
	if err != nil {
		writeEmptyUTXOResponse(w)
		return
	}
	details := make([]singleUTXOdetails, len(utxos))
	for i, utxo := range utxos {
		detail := transaction.ParseIndexIntoUTXOdetails(utxo)
		responseDetails := singleUTXOdetails{int(detail.BlockNumber), int(detail.TransactionNumber),
			int(detail.OutputNumber), detail.Value}
		details[i] = responseDetails
	}
	writeUTXOResponse(w, details)
	return
}

func (p *ListUTXOsHandler) getUTXOsForAddress(address common.Address, afterBlock uint32, afterTransaction uint32, afterOutput uint8, limit int) ([][transaction.UTXOIndexLength]byte, error) {
	prefix := []byte("utxo")
	prefix = append(prefix, address[:]...)
	blockNumberBuffer := make([]byte, transaction.BlockNumberLength)
	binary.BigEndian.PutUint32(blockNumberBuffer, afterBlock)
	transactionNumberBuffer := make([]byte, transaction.TransactionNumberLength)
	binary.BigEndian.PutUint32(transactionNumberBuffer, afterTransaction)
	outputNumberBuffer := make([]byte, transaction.OutputNumberLength)
	outputNumberBuffer[0] = afterOutput
	valueBuffer := make([]byte, transaction.ValueLength)
	startingIndex := []byte("utxo")
	startingIndex = append(startingIndex, address[:]...)
	startingIndex = append(startingIndex, blockNumberBuffer...)
	startingIndex = append(startingIndex, transactionNumberBuffer...)
	startingIndex = append(startingIndex, outputNumberBuffer...)
	startingIndex = append(startingIndex, valueBuffer...)

	values := make([][]byte, limit)
	counter := 0

	err := p.db.View(func(txn *badger.Txn) error {
		opts := badger.DefaultIteratorOptions
		opts.PrefetchValues = false
		it := txn.NewIterator(opts)
		defer it.Close()
		for it.Seek(startingIndex); it.ValidForPrefix(prefix); it.Next() {
			item := it.Item()
			k := item.Key()
			values[counter] = k
			counter++
			if counter > limit {
				break
			}
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	toCutFromKey := 4
	toReturn := make([][transaction.UTXOIndexLength]byte, counter)
	for i := 0; i < counter; i++ {
		key := values[i]
		index := [transaction.UTXOIndexLength]byte{}
		keyMeaningfulPart := key[toCutFromKey:]
		copy(index[:], keyMeaningfulPart)
		toReturn[i] = index
	}
	return toReturn, nil
}

func writeEmptyUTXOResponse(w http.ResponseWriter) {
	response := listUTXOsResponse{true, []singleUTXOdetails{}}
	w.WriteHeader(200)
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	w.Header().Set("Content-Type", "application/json")
	body, _ := json.Marshal(response)
	w.Write(body)
}

func writeUTXOResponse(w http.ResponseWriter, details []singleUTXOdetails) {
	response := listUTXOsResponse{false, details}
	w.WriteHeader(200)
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	w.Header().Set("Content-Type", "application/json")
	body, _ := json.Marshal(response)
	w.Write(body)
}
