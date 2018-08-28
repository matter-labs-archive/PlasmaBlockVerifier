package plasmainteraction

import (
	"io/ioutil"
	"strconv"
	"time"
)

const MaxBlockDelay = 300 // 5 minutes

type BlockDownloader struct {
}

func NewBlockDownloader() *BlockDownloader {
	newInstance := &BlockDownloader{}
	return newInstance
}

// mocked for now, to get blocks from local storage for testing
func (p *BlockDownloader) GetBlock(blockNumber uint32) chan []byte {
	resultChannel := make(chan []byte)
	go func(bn uint32) {
		start := time.Now()
		for {
			blockNumberString := strconv.FormatUint(uint64(blockNumber), 10)
			data, err := ioutil.ReadFile("./plasmainteraction/testBlocks/" + blockNumberString)
			elapsed := time.Since(start)
			if err != nil && elapsed.Seconds() > MaxBlockDelay {
				resultChannel <- []byte{}
				return
			} else {
				resultChannel <- data
				return
			}
		}
	}(blockNumber)
	return resultChannel
}
