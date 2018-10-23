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
	restapi "github.com/matterinc/PlasmaBlockVerifier/restAPI"
)

func main() {
	database.PurgeTestDatabase()
	db, err := database.OpenTestDatabase()
	if err != nil {
		log.Fatalln(err)
	}
	defer db.Close()

	ethereumLoop := eth.NewEthereumNetworkEventDispatcher("https://mainnet.infura.io", "0xc71d3638e2f14c115eb53163ae97ae76eae9d7f6")
	depositCheckoutsProcessor := eth.NewDepositCheckoutProcessor(db, ethereumLoop.PlasmaParentContract, ethereumLoop.BlockStorageContract)
	withdrawChallengeProcessor := eth.NewWithdrawChallengeProcessor(ethereumLoop.PlasmaParentContract, db)
	withdrawStartedProcessor := eth.NewWithdrawStartedProcessor(db)
	blockProcessor := plasma.NewBlockProcessor(db, 100, 100)
	processingLoop := plasma.NewBlockProcessingLoop(blockProcessor)

	// ethereumLoop.Run(6525445, dispatch.BlockInformationChannel, withdrawStartedProcessor, dispatch.ControlChannel)
	// ethereumLoop.Run(6527670, dispatch.BlockInformationChannel, withdrawStartedProcessor)
	ethereumLoop.Run(6530700, dispatch.BlockInformationChannel, withdrawStartedProcessor)
	processingLoop.Run(dispatch.BlockInformationChannel, depositCheckoutsProcessor, withdrawChallengeProcessor)

	restAPI := restapi.NewRestAPI(db)
	restAPI.Start(8000)
	wait := time.Second * 15
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c
	_, cancel := context.WithTimeout(context.Background(), wait)
	ethereumLoop.SignalChannel <- 1
	processingLoop.ControlChannel <- 1
	restAPI.ControlChannel <- 1
	defer cancel()
	log.Println("Shutting down")
	os.Exit(0)
}
