package plasmainteraction

import (
	"io/ioutil"
	"strconv"
)

type BlockDownloader struct {
}

func NewBlockDownloader() *BlockDownloader {
	newInstance := &BlockDownloader{}
	return newInstance
}

// mocked for now, to get blocks from local storage for testing
func (p *BlockDownloader) GetBlock(blockNumber uint32) ([]byte, error) {
	blockNumberString := strconv.FormatUint(uint64(blockNumber), 10)
	data, err := ioutil.ReadFile("testBlocks/" + blockNumberString)
	return data, err
}
