package plasmainteraction

import (
	"fmt"
	"strconv"
	"time"
)

const loopTimeMS = 1000

type BlockInformation struct {
	blockNumber     uint32
	blockHash       [32]byte
	blockMerkleRoot [32]byte
}

type BlockProcessingLoop struct {
	controlChannel            chan int
	processingRequestsChannel chan BlockInformation
	depositChecksChannel      chan DepositIndexCheckoutRequest
	withdrawChallengesChannel chan WithdrawChallengeRequest
	processor                 *BlockProcessor
}

func NewBlockProcessingLoop(processor *BlockProcessor) *BlockProcessingLoop {
	cont := make(chan int)
	bl := make(chan BlockInformation)
	deps := make(chan DepositIndexCheckoutRequest)
	withdraws := make(chan WithdrawChallengeRequest)
	newInstance := &BlockProcessingLoop{cont, bl, deps, withdraws, processor}
	return newInstance
}

func (p *BlockProcessingLoop) Run() {
	loopFunction := func() {
		for {
			select {
			case controlMessage := <-p.controlChannel:
				fmt.Println("received control message")
				fmt.Println(strconv.Itoa(controlMessage))
			default:
				fmt.Println("continue to process")
			}
			select {
			case newBlockInfo := <-p.processingRequestsChannel:
				continue
				fmt.Println(newBlockInfo)
				// p.processor.ProcessBlock(newBlockInfo.)
			default:
				time.Sleep(loopTimeMS)
			}
		}

	}
	go func() {
		loopFunction()
	}()
}
