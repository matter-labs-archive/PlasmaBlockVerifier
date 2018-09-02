package plasmainteraction

import (
	"fmt"
	"strconv"
	"time"

	ethereuminteraction "github.com/matterinc/PlasmaBlockVerifier/ethereuminteraction"
	"github.com/matterinc/PlasmaBlockVerifier/messageStructures"
)

const loopTimeMS = 1000

// TODO add checker for exit events!
type BlockProcessingLoop struct {
	ControlChannel chan int
	// ProcessingRequestsChannel chan *BlockInformation
	// DepositChecksChannel      chan *DepositIndexCheckoutRequest
	// WithdrawChallengesChannel chan *WithdrawChallengeRequest
	Processor *BlockProcessor
}

func NewBlockProcessingLoop(processor *BlockProcessor) *BlockProcessingLoop {
	cont := make(chan int)
	// bl := make(chan *BlockInformation)
	// deps := make(chan *DepositIndexCheckoutRequest)
	// withdraws := make(chan *WithdrawChallengeRequest)
	// newInstance := &BlockProcessingLoop{cont, bl, deps, withdraws, processor}
	newInstance := &BlockProcessingLoop{cont, processor}
	return newInstance
}

func (p *BlockProcessingLoop) Run(blockInfoChannel <-chan *messageStructures.BlockInformation,
	depositCheckoutProcessor *ethereuminteraction.DepositCheckoutProcessor,
	withdrawChallengeProcessor *ethereuminteraction.WithdrawChallengeProcessor) {
	blockDownloader := NewBlockDownloader()
	depositCheckoutsChannel := make(chan *messageStructures.DepositIndexCheckoutRequest)
	withdrawChallengesChannel := make(chan *messageStructures.WithdrawChallengeRequest)
	blockFunction := func() {
		for {
			select {
			case controlMessage := <-p.ControlChannel:
				fmt.Println("Received control message")
				fmt.Println(strconv.Itoa(controlMessage))
			default:
				// fmt.Println("continue to process")
			}
			select {
			case newBlockInfo := <-blockInfoChannel:
				newBlockBytes := <-blockDownloader.GetBlock(newBlockInfo.BlockNumber)
				if len(newBlockBytes) == 0 {
					panic("Block is probably withheld")
				}
				deps, withs, err := p.Processor.ProcessBlock(newBlockBytes, newBlockInfo.BlockHash[:], newBlockInfo.BlockMerkleRoot[:])
				if err != nil {
					panic("Unrecoverable error, chain is byzantine")
				}
				for _, d := range deps {
					fmt.Println("Processing deposit checks")
					depositCheckoutsChannel <- d
				}
				for _, w := range withs {
					fmt.Println("Processing exit challenges")
					withdrawChallengesChannel <- w
				}
			default:
				// fmt.Println("No block to process")
			}
		}
	}
	ethereumFunction := func() {
		for {
			select {
			case depositCheckoutRequest := <-depositCheckoutsChannel:
				success, err := depositCheckoutProcessor.Process(depositCheckoutRequest)
				if !success || err != nil {
					panic("Unrecoverable error, chain is byzantine")
				}
			default:
				// fmt.Println("No deposits to process")
			}
			select {
			case challengeRequest := <-withdrawChallengesChannel:
				success, err := withdrawChallengeProcessor.Process(challengeRequest)
				if !success || err != nil {
					panic("Unrecoverable error, chain is byzantine")
				}
			default:
				// fmt.Println("No withdraw challenges to process")
			}
			time.Sleep(loopTimeMS)
		}
	}
	go blockFunction()
	go ethereumFunction()
}
