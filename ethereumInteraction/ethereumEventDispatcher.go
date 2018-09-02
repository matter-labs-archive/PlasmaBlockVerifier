package ethereuminteraction

import (
	"context"
	"fmt"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"

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

func (p *EthereumNetworkEventDispatcher) Run(fromBlockNumber int64, blockProcessorLoopChannel chan<- *messageStructures.BlockInformation, withdrawProcessor *WithdrawStartedProcessor) {
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
			blockNumberUInt64 := newBlockNumber.Uint64()

			filterOptions := &bind.FilterOpts{blockNumberUInt64, &blockNumberUInt64, context.TODO()}
			exitEventsIterator, err := p.PlasmaParentContract.PlasmaParentFilterer.FilterExitStartedEvent(filterOptions, nil, nil, nil)
			if err != nil {
				fmt.Println(err)
				continue
			}
			for exitEventsIterator.Next() {
				ev := exitEventsIterator.Event
				fmt.Println("Processing exit event")
				fullInfo, err := p.PlasmaParentContract.PlasmaParentCaller.PublishedUTXOs(nil, ev.Index)
				if err != nil {
					fmt.Println(err)
					panic(err)
				}

				info := &messageStructures.WithdrawStartedInformation{false, ev.Index, nil, ev.From, fullInfo.Value}
				_, err = withdrawProcessor.Process(info)
				if err != nil {
					fmt.Println(err)
					panic(err)
				}
			}
			blocksEventIterator, err := p.BlockStorageContract.PlasmaBlockStorageFilterer.FilterBlockHeaderSubmitted(filterOptions, nil, nil)
			if err != nil {
				fmt.Println(err)
				continue
			}
			for blocksEventIterator.Next() {
				ev := blocksEventIterator.Event
				fmt.Println("Processing Plasma block " + ev.BlockNumber.String() + " in Ethereum block " + newBlockNumber.String())
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
