package ethereuminteraction

import (
	"context"
	"fmt"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/common"
	ethclient "github.com/ethereum/go-ethereum/ethclient"
	ABI "github.com/matterinc/PlasmaBlockVerifier/contractABI"
	messageStructures "github.com/matterinc/PlasmaBlockVerifier/messageStructures"
)

const loopDelay = 10000000 // 10 seconds

type EthereumNetworkEventDispatcher struct {
	ConnectionString     string
	Client               *ethclient.Client
	SignalChannel        chan int
	PlasmaParentContract *ABI.PlasmaParent
	BlockStorageContract *ABI.PlasmaBlockStorage
}

func NewEthereumNetworkEventDispatcher(connection, contractAddress string) *EthereumNetworkEventDispatcher {
	conn, err := ethclient.Dial(connection)
	if err != nil {
		return nil
	}
	signalChan := make(chan int)
	address := common.HexToAddress(contractAddress)
	plasma, err := ABI.NewPlasmaParent(address, conn)
	blockStorageAddress, err := plasma.PlasmaParentCaller.BlockStorage(nil)
	if err != nil {
		panic("Can not get block storage contract address")
	}
	blockStorage, err := ABI.NewPlasmaBlockStorage(blockStorageAddress, conn)
	newInstance := &EthereumNetworkEventDispatcher{connection, conn, signalChan, plasma, blockStorage}
	return newInstance
}

func (p *EthereumNetworkEventDispatcher) Run(fromBlockNumber int64, blockProcessorLoopChannel chan<- *messageStructures.BlockInformation) {
	lastBlockNumber := big.NewInt(fromBlockNumber - 1)
	if fromBlockNumber <= 1 {
		lastBlockNumber.SetUint64(0)
	}
	one := big.NewInt(1)
	eventLoop := func() {
		for {
			time.Sleep(loopDelay)
			newBlockNumber := new(big.Int).Add(lastBlockNumber, one)
			_, err := p.Client.HeaderByNumber(context.TODO(), newBlockNumber)
			if err != nil {
				continue
			}
			fmt.Println("Processing Ethereum block " + newBlockNumber.String())
			eventIterator, err := p.BlockStorageContract.PlasmaBlockStorageFilterer.FilterBlockHeaderSubmitted(nil, []*big.Int{newBlockNumber}, nil)
			if err != nil {
				fmt.Println(err)
				continue
			}
			for eventIterator.Next() {
				ev := eventIterator.Event
				fmt.Println("Processing Plasma block " + ev.BlockNumber.String())
				plasmaBlockNumber := uint32(ev.BlockNumber.Uint64())
				merkleRoot := ev.MerkleRoot
				blockProcessingInformation := &messageStructures.BlockInformation{BlockNumber: plasmaBlockNumber, BlockHash: [32]byte{}, BlockMerkleRoot: merkleRoot}
				blockProcessorLoopChannel <- blockProcessingInformation
			}
			lastBlockNumber = newBlockNumber
		}
	}
	go eventLoop()
}
