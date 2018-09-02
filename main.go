package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"time"

	database "github.com/matterinc/PlasmaBlockVerifier/database"
	eth "github.com/matterinc/PlasmaBlockVerifier/ethereuminteraction"
	dispatch "github.com/matterinc/PlasmaBlockVerifier/grandCentralDispatch"
	plasma "github.com/matterinc/PlasmaBlockVerifier/plasmainteraction"
)

func main() {
	database.PurgeTestDatabase()
	db, err := database.OpenTestDatabase()
	ethereumLoop := eth.NewEthereumNetworkEventDispatcher("http://127.0.0.1:8545", "0x345ca3e014aaf5dca488057592ee47305d9b3e10")
	depositCheckoutsProcessor := eth.NewDepositCheckoutProcessor(db, ethereumLoop.PlasmaParentContract, ethereumLoop.BlockStorageContract)
	withdrawChallengeProcessor := eth.NewWithdrawChallengeProcessor(ethereumLoop.PlasmaParentContract, db)
	withdrawStartedProcessor := eth.NewWithdrawStartedProcessor(db)
	blockProcessor := plasma.NewBlockProcessor(db, 100, 100)
	processingLoop := plasma.NewBlockProcessingLoop(blockProcessor)

	ethereumLoop.Run(0, dispatch.BlockInformationChannel, withdrawStartedProcessor)
	processingLoop.Run(dispatch.BlockInformationChannel, depositCheckoutsProcessor, withdrawChallengeProcessor)

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
