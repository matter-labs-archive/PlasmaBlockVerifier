package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"time"

	database "github.com/matterinc/PlasmaBlockVerifier/database"
	eth "github.com/matterinc/PlasmaBlockVerifier/ethereuminteraction"
	plasma "github.com/matterinc/PlasmaBlockVerifier/plasmainteraction"
)

func main() {
	database.PurgeTestDatabase()
	db, err := database.OpenTestDatabase()

	blockProcessor := plasma.NewBlockProcessor(db, 100, 100)
	processingLoop := plasma.NewBlockProcessingLoop(blockProcessor)
	ethereumLoop := eth.NewEthereumNetworkEventDispatcher("http://127.0.0.1:8545", "0x82d50ad3c1091866e258fd0f1a7cc9674609d254")
	ethereumLoop.Run(0, processingLoop.ProcessingRequestsChannel)
	processingLoop.Run()

	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	wait := time.Second * 15
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c
	_, cancel := context.WithTimeout(context.Background(), wait)
	defer cancel()
	log.Println("Shutting down")
	os.Exit(0)
}
