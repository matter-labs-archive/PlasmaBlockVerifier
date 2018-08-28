package plasmainteraction

import (
	"fmt"
	"strconv"
	"time"
)

const loopTimeMS = 1000

type BlockInformation struct {
	BlockNumber     uint32
	BlockHash       [32]byte
	BlockMerkleRoot [32]byte
}

// TODO add checker for exit events!
type BlockProcessingLoop struct {
	ControlChannel            chan int
	ProcessingRequestsChannel chan *BlockInformation
	DepositChecksChannel      chan *DepositIndexCheckoutRequest
	WithdrawChallengesChannel chan *WithdrawChallengeRequest
	Processor                 *BlockProcessor
}

func NewBlockProcessingLoop(processor *BlockProcessor) *BlockProcessingLoop {
	cont := make(chan int)
	bl := make(chan *BlockInformation)
	deps := make(chan *DepositIndexCheckoutRequest)
	withdraws := make(chan *WithdrawChallengeRequest)
	newInstance := &BlockProcessingLoop{cont, bl, deps, withdraws, processor}
	return newInstance
}

func (p *BlockProcessingLoop) Run() {
	blockDownloader := NewBlockDownloader()
	loopFunction := func() {
		for {
			select {
			case controlMessage := <-p.ControlChannel:
				fmt.Println("Received control message")
				fmt.Println(strconv.Itoa(controlMessage))
			default:
				// fmt.Println("continue to process")
			}
			select {
			case newBlockInfo := <-p.ProcessingRequestsChannel:
				newBlockBytes := <-blockDownloader.GetBlock(newBlockInfo.BlockNumber)
				if len(newBlockBytes) == 0 {
					panic("Block is probably withheld")
				}
				deps, withs, err := p.Processor.ProcessBlock(newBlockBytes, newBlockInfo.BlockHash[:], newBlockInfo.BlockMerkleRoot[:])
				if err != nil {
					panic("Unrecoverable error, chain is byzantine")
				}
				for _, d := range deps {
					p.DepositChecksChannel <- d
				}
				for _, w := range withs {
					p.WithdrawChallengesChannel <- w
				}
			default:
				// fmt.Println("No block to process")
			}
			time.Sleep(loopTimeMS)
		}
	}
	go func() {
		loopFunction()
	}()
}
