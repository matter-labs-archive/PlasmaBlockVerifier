package restAPI

import (
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/dgraph-io/badger"
	"github.com/gorilla/mux"
)

type RestAPI struct {
	ControlChannel chan int
	db             *badger.DB
}

func NewRestAPI(db *badger.DB) *RestAPI {
	ch := make(chan int)
	return &RestAPI{ch, db}
}

func (p *RestAPI) Start(port int) error {
	r := mux.NewRouter()
	listUTXOhandler := NewListUTXOsHandler(p.db)
	// Add your routes as needed
	r.HandleFunc("/listUTXOs", listUTXOhandler.handler).Methods("POST")
	// r.HandleFunc("/listDeposits", handler).Methods("POST")
	// r.HandleFunc("/listWithdraws", handler).Methods("GET")
	srv := &http.Server{
		Addr: "0.0.0.0:" + strconv.Itoa(port),
		// Good practice to set timeouts to avoid Slowloris attacks.
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      r, // Pass our instance of gorilla/mux in.
	}

	// Run our server in a goroutine so that it doesn't block.
	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Println(err)
		}
	}()
	go func() {
		_ = <-p.ControlChannel
		srv.Close()
	}()
	return nil
}
