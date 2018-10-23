package plasmainteraction

import (
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
)

const MaxBlockDelay = 300 // 5 minutes
const blockStorageURL = "https://plasma.ams3.digitaloceanspaces.com/plasma/"

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
			resp, err := http.Get(blockStorageURL + blockNumberString)
			defer resp.Body.Close()
			elapsed := time.Since(start)
			if err != nil {
				if elapsed.Seconds() > MaxBlockDelay {
					resultChannel <- []byte{}
					return
				} else {
					time.Sleep(5000000000) // 5 seconds
					continue
				}
			}
			body, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				if elapsed.Seconds() > MaxBlockDelay {
					resultChannel <- []byte{}
					return
				} else {
					time.Sleep(5000000000) // 5 seconds
					continue
				}
			}
			resultChannel <- body
			return
		}
	}(blockNumber)
	return resultChannel
}

// // mocked for now, to get blocks from local storage for testing
// func (p *BlockDownloader) GetBlock(blockNumber uint32) chan []byte {
// 	resultChannel := make(chan []byte)
// 	go func(bn uint32) {
// 		start := time.Now()
// 		for {
// 			blockNumberString := strconv.FormatUint(uint64(blockNumber), 10)
// 			data, err := ioutil.ReadFile("./plasmainteraction/testBlocks/" + blockNumberString)
// 			elapsed := time.Since(start)
// 			if err != nil && elapsed.Seconds() > MaxBlockDelay {
// 				resultChannel <- []byte{}
// 				return
// 			} else {
// 				resultChannel <- data
// 				return
// 			}
// 		}
// 	}(blockNumber)
// 	return resultChannel
// }
