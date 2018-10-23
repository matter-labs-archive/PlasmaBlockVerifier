// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contractABI

import (
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = abi.U256
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// PlasmaBlockStorageABI is the input ABI used to generate the binding from.
const PlasmaBlockStorageABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"operators\",\"outputs\":[{\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"lastBlockNumber\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"weekOldBlockNumber\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"blockHeaderLength\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"hashOfLastSubmittedBlock\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"blocks\",\"outputs\":[{\"name\":\"numberOfTransactions\",\"type\":\"uint32\"},{\"name\":\"submittedAt\",\"type\":\"uint64\"},{\"name\":\"merkleRootHash\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_blockNumber\",\"type\":\"uint256\"},{\"indexed\":true,\"name\":\"_merkleRoot\",\"type\":\"bytes32\"}],\"name\":\"BlockHeaderSubmitted\",\"type\":\"event\"},{\"constant\":false,\"inputs\":[{\"name\":\"_newOwner\",\"type\":\"address\"}],\"name\":\"setOwner\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_op\",\"type\":\"address\"},{\"name\":\"_status\",\"type\":\"uint256\"}],\"name\":\"setOperator\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_operator\",\"type\":\"address\"}],\"name\":\"isOperator\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_operator\",\"type\":\"address\"}],\"name\":\"canSignBlocks\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"incrementWeekOldCounter\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_headers\",\"type\":\"bytes\"}],\"name\":\"submitBlockHeaders\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_blockNumber\",\"type\":\"uint32\"}],\"name\":\"getBlockInformation\",\"outputs\":[{\"name\":\"submittedAt\",\"type\":\"uint256\"},{\"name\":\"numberOfTransactions\",\"type\":\"uint32\"},{\"name\":\"merkleRoot\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_blockNumber\",\"type\":\"uint32\"}],\"name\":\"getMerkleRoot\",\"outputs\":[{\"name\":\"merkleRoot\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_blockNumber\",\"type\":\"uint32\"}],\"name\":\"getSubmissionTime\",\"outputs\":[{\"name\":\"submittedAt\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_blockNumber\",\"type\":\"uint32\"}],\"name\":\"getNumberOfTransactions\",\"outputs\":[{\"name\":\"numberOfTransaction\",\"type\":\"uint32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// PlasmaBlockStorage is an auto generated Go binding around an Ethereum contract.
type PlasmaBlockStorage struct {
	PlasmaBlockStorageCaller     // Read-only binding to the contract
	PlasmaBlockStorageTransactor // Write-only binding to the contract
	PlasmaBlockStorageFilterer   // Log filterer for contract events
}

// PlasmaBlockStorageCaller is an auto generated read-only Go binding around an Ethereum contract.
type PlasmaBlockStorageCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PlasmaBlockStorageTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PlasmaBlockStorageTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PlasmaBlockStorageFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PlasmaBlockStorageFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PlasmaBlockStorageSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PlasmaBlockStorageSession struct {
	Contract     *PlasmaBlockStorage // Generic contract binding to set the session for
	CallOpts     bind.CallOpts       // Call options to use throughout this session
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// PlasmaBlockStorageCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PlasmaBlockStorageCallerSession struct {
	Contract *PlasmaBlockStorageCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts             // Call options to use throughout this session
}

// PlasmaBlockStorageTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PlasmaBlockStorageTransactorSession struct {
	Contract     *PlasmaBlockStorageTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// PlasmaBlockStorageRaw is an auto generated low-level Go binding around an Ethereum contract.
type PlasmaBlockStorageRaw struct {
	Contract *PlasmaBlockStorage // Generic contract binding to access the raw methods on
}

// PlasmaBlockStorageCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PlasmaBlockStorageCallerRaw struct {
	Contract *PlasmaBlockStorageCaller // Generic read-only contract binding to access the raw methods on
}

// PlasmaBlockStorageTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PlasmaBlockStorageTransactorRaw struct {
	Contract *PlasmaBlockStorageTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPlasmaBlockStorage creates a new instance of PlasmaBlockStorage, bound to a specific deployed contract.
func NewPlasmaBlockStorage(address common.Address, backend bind.ContractBackend) (*PlasmaBlockStorage, error) {
	contract, err := bindPlasmaBlockStorage(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &PlasmaBlockStorage{PlasmaBlockStorageCaller: PlasmaBlockStorageCaller{contract: contract}, PlasmaBlockStorageTransactor: PlasmaBlockStorageTransactor{contract: contract}, PlasmaBlockStorageFilterer: PlasmaBlockStorageFilterer{contract: contract}}, nil
}

// NewPlasmaBlockStorageCaller creates a new read-only instance of PlasmaBlockStorage, bound to a specific deployed contract.
func NewPlasmaBlockStorageCaller(address common.Address, caller bind.ContractCaller) (*PlasmaBlockStorageCaller, error) {
	contract, err := bindPlasmaBlockStorage(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PlasmaBlockStorageCaller{contract: contract}, nil
}

// NewPlasmaBlockStorageTransactor creates a new write-only instance of PlasmaBlockStorage, bound to a specific deployed contract.
func NewPlasmaBlockStorageTransactor(address common.Address, transactor bind.ContractTransactor) (*PlasmaBlockStorageTransactor, error) {
	contract, err := bindPlasmaBlockStorage(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PlasmaBlockStorageTransactor{contract: contract}, nil
}

// NewPlasmaBlockStorageFilterer creates a new log filterer instance of PlasmaBlockStorage, bound to a specific deployed contract.
func NewPlasmaBlockStorageFilterer(address common.Address, filterer bind.ContractFilterer) (*PlasmaBlockStorageFilterer, error) {
	contract, err := bindPlasmaBlockStorage(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PlasmaBlockStorageFilterer{contract: contract}, nil
}

// bindPlasmaBlockStorage binds a generic wrapper to an already deployed contract.
func bindPlasmaBlockStorage(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(PlasmaBlockStorageABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PlasmaBlockStorage *PlasmaBlockStorageRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _PlasmaBlockStorage.Contract.PlasmaBlockStorageCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PlasmaBlockStorage *PlasmaBlockStorageRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PlasmaBlockStorage.Contract.PlasmaBlockStorageTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PlasmaBlockStorage *PlasmaBlockStorageRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PlasmaBlockStorage.Contract.PlasmaBlockStorageTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PlasmaBlockStorage *PlasmaBlockStorageCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _PlasmaBlockStorage.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PlasmaBlockStorage *PlasmaBlockStorageTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PlasmaBlockStorage.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PlasmaBlockStorage *PlasmaBlockStorageTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PlasmaBlockStorage.Contract.contract.Transact(opts, method, params...)
}

// BlockHeaderLength is a free data retrieval call binding the contract method 0x68349324.
//
// Solidity: function blockHeaderLength() constant returns(uint256)
func (_PlasmaBlockStorage *PlasmaBlockStorageCaller) BlockHeaderLength(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _PlasmaBlockStorage.contract.Call(opts, out, "blockHeaderLength")
	return *ret0, err
}

// BlockHeaderLength is a free data retrieval call binding the contract method 0x68349324.
//
// Solidity: function blockHeaderLength() constant returns(uint256)
func (_PlasmaBlockStorage *PlasmaBlockStorageSession) BlockHeaderLength() (*big.Int, error) {
	return _PlasmaBlockStorage.Contract.BlockHeaderLength(&_PlasmaBlockStorage.CallOpts)
}

// BlockHeaderLength is a free data retrieval call binding the contract method 0x68349324.
//
// Solidity: function blockHeaderLength() constant returns(uint256)
func (_PlasmaBlockStorage *PlasmaBlockStorageCallerSession) BlockHeaderLength() (*big.Int, error) {
	return _PlasmaBlockStorage.Contract.BlockHeaderLength(&_PlasmaBlockStorage.CallOpts)
}

// Blocks is a free data retrieval call binding the contract method 0xf25b3f99.
//
// Solidity: function blocks( uint256) constant returns(numberOfTransactions uint32, submittedAt uint64, merkleRootHash bytes32)
func (_PlasmaBlockStorage *PlasmaBlockStorageCaller) Blocks(opts *bind.CallOpts, arg0 *big.Int) (struct {
	NumberOfTransactions uint32
	SubmittedAt          uint64
	MerkleRootHash       [32]byte
}, error) {
	ret := new(struct {
		NumberOfTransactions uint32
		SubmittedAt          uint64
		MerkleRootHash       [32]byte
	})
	out := ret
	err := _PlasmaBlockStorage.contract.Call(opts, out, "blocks", arg0)
	return *ret, err
}

// Blocks is a free data retrieval call binding the contract method 0xf25b3f99.
//
// Solidity: function blocks( uint256) constant returns(numberOfTransactions uint32, submittedAt uint64, merkleRootHash bytes32)
func (_PlasmaBlockStorage *PlasmaBlockStorageSession) Blocks(arg0 *big.Int) (struct {
	NumberOfTransactions uint32
	SubmittedAt          uint64
	MerkleRootHash       [32]byte
}, error) {
	return _PlasmaBlockStorage.Contract.Blocks(&_PlasmaBlockStorage.CallOpts, arg0)
}

// Blocks is a free data retrieval call binding the contract method 0xf25b3f99.
//
// Solidity: function blocks( uint256) constant returns(numberOfTransactions uint32, submittedAt uint64, merkleRootHash bytes32)
func (_PlasmaBlockStorage *PlasmaBlockStorageCallerSession) Blocks(arg0 *big.Int) (struct {
	NumberOfTransactions uint32
	SubmittedAt          uint64
	MerkleRootHash       [32]byte
}, error) {
	return _PlasmaBlockStorage.Contract.Blocks(&_PlasmaBlockStorage.CallOpts, arg0)
}

// CanSignBlocks is a free data retrieval call binding the contract method 0xe2a56a44.
//
// Solidity: function canSignBlocks(_operator address) constant returns(bool)
func (_PlasmaBlockStorage *PlasmaBlockStorageCaller) CanSignBlocks(opts *bind.CallOpts, _operator common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _PlasmaBlockStorage.contract.Call(opts, out, "canSignBlocks", _operator)
	return *ret0, err
}

// CanSignBlocks is a free data retrieval call binding the contract method 0xe2a56a44.
//
// Solidity: function canSignBlocks(_operator address) constant returns(bool)
func (_PlasmaBlockStorage *PlasmaBlockStorageSession) CanSignBlocks(_operator common.Address) (bool, error) {
	return _PlasmaBlockStorage.Contract.CanSignBlocks(&_PlasmaBlockStorage.CallOpts, _operator)
}

// CanSignBlocks is a free data retrieval call binding the contract method 0xe2a56a44.
//
// Solidity: function canSignBlocks(_operator address) constant returns(bool)
func (_PlasmaBlockStorage *PlasmaBlockStorageCallerSession) CanSignBlocks(_operator common.Address) (bool, error) {
	return _PlasmaBlockStorage.Contract.CanSignBlocks(&_PlasmaBlockStorage.CallOpts, _operator)
}

// GetBlockInformation is a free data retrieval call binding the contract method 0xd1c55ee8.
//
// Solidity: function getBlockInformation(_blockNumber uint32) constant returns(submittedAt uint256, numberOfTransactions uint32, merkleRoot bytes32)
func (_PlasmaBlockStorage *PlasmaBlockStorageCaller) GetBlockInformation(opts *bind.CallOpts, _blockNumber uint32) (struct {
	SubmittedAt          *big.Int
	NumberOfTransactions uint32
	MerkleRoot           [32]byte
}, error) {
	ret := new(struct {
		SubmittedAt          *big.Int
		NumberOfTransactions uint32
		MerkleRoot           [32]byte
	})
	out := ret
	err := _PlasmaBlockStorage.contract.Call(opts, out, "getBlockInformation", _blockNumber)
	return *ret, err
}

// GetBlockInformation is a free data retrieval call binding the contract method 0xd1c55ee8.
//
// Solidity: function getBlockInformation(_blockNumber uint32) constant returns(submittedAt uint256, numberOfTransactions uint32, merkleRoot bytes32)
func (_PlasmaBlockStorage *PlasmaBlockStorageSession) GetBlockInformation(_blockNumber uint32) (struct {
	SubmittedAt          *big.Int
	NumberOfTransactions uint32
	MerkleRoot           [32]byte
}, error) {
	return _PlasmaBlockStorage.Contract.GetBlockInformation(&_PlasmaBlockStorage.CallOpts, _blockNumber)
}

// GetBlockInformation is a free data retrieval call binding the contract method 0xd1c55ee8.
//
// Solidity: function getBlockInformation(_blockNumber uint32) constant returns(submittedAt uint256, numberOfTransactions uint32, merkleRoot bytes32)
func (_PlasmaBlockStorage *PlasmaBlockStorageCallerSession) GetBlockInformation(_blockNumber uint32) (struct {
	SubmittedAt          *big.Int
	NumberOfTransactions uint32
	MerkleRoot           [32]byte
}, error) {
	return _PlasmaBlockStorage.Contract.GetBlockInformation(&_PlasmaBlockStorage.CallOpts, _blockNumber)
}

// GetMerkleRoot is a free data retrieval call binding the contract method 0x066e89a6.
//
// Solidity: function getMerkleRoot(_blockNumber uint32) constant returns(merkleRoot bytes32)
func (_PlasmaBlockStorage *PlasmaBlockStorageCaller) GetMerkleRoot(opts *bind.CallOpts, _blockNumber uint32) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _PlasmaBlockStorage.contract.Call(opts, out, "getMerkleRoot", _blockNumber)
	return *ret0, err
}

// GetMerkleRoot is a free data retrieval call binding the contract method 0x066e89a6.
//
// Solidity: function getMerkleRoot(_blockNumber uint32) constant returns(merkleRoot bytes32)
func (_PlasmaBlockStorage *PlasmaBlockStorageSession) GetMerkleRoot(_blockNumber uint32) ([32]byte, error) {
	return _PlasmaBlockStorage.Contract.GetMerkleRoot(&_PlasmaBlockStorage.CallOpts, _blockNumber)
}

// GetMerkleRoot is a free data retrieval call binding the contract method 0x066e89a6.
//
// Solidity: function getMerkleRoot(_blockNumber uint32) constant returns(merkleRoot bytes32)
func (_PlasmaBlockStorage *PlasmaBlockStorageCallerSession) GetMerkleRoot(_blockNumber uint32) ([32]byte, error) {
	return _PlasmaBlockStorage.Contract.GetMerkleRoot(&_PlasmaBlockStorage.CallOpts, _blockNumber)
}

// GetNumberOfTransactions is a free data retrieval call binding the contract method 0x4cbb0ed8.
//
// Solidity: function getNumberOfTransactions(_blockNumber uint32) constant returns(numberOfTransaction uint32)
func (_PlasmaBlockStorage *PlasmaBlockStorageCaller) GetNumberOfTransactions(opts *bind.CallOpts, _blockNumber uint32) (uint32, error) {
	var (
		ret0 = new(uint32)
	)
	out := ret0
	err := _PlasmaBlockStorage.contract.Call(opts, out, "getNumberOfTransactions", _blockNumber)
	return *ret0, err
}

// GetNumberOfTransactions is a free data retrieval call binding the contract method 0x4cbb0ed8.
//
// Solidity: function getNumberOfTransactions(_blockNumber uint32) constant returns(numberOfTransaction uint32)
func (_PlasmaBlockStorage *PlasmaBlockStorageSession) GetNumberOfTransactions(_blockNumber uint32) (uint32, error) {
	return _PlasmaBlockStorage.Contract.GetNumberOfTransactions(&_PlasmaBlockStorage.CallOpts, _blockNumber)
}

// GetNumberOfTransactions is a free data retrieval call binding the contract method 0x4cbb0ed8.
//
// Solidity: function getNumberOfTransactions(_blockNumber uint32) constant returns(numberOfTransaction uint32)
func (_PlasmaBlockStorage *PlasmaBlockStorageCallerSession) GetNumberOfTransactions(_blockNumber uint32) (uint32, error) {
	return _PlasmaBlockStorage.Contract.GetNumberOfTransactions(&_PlasmaBlockStorage.CallOpts, _blockNumber)
}

// GetSubmissionTime is a free data retrieval call binding the contract method 0xc3943844.
//
// Solidity: function getSubmissionTime(_blockNumber uint32) constant returns(submittedAt uint256)
func (_PlasmaBlockStorage *PlasmaBlockStorageCaller) GetSubmissionTime(opts *bind.CallOpts, _blockNumber uint32) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _PlasmaBlockStorage.contract.Call(opts, out, "getSubmissionTime", _blockNumber)
	return *ret0, err
}

// GetSubmissionTime is a free data retrieval call binding the contract method 0xc3943844.
//
// Solidity: function getSubmissionTime(_blockNumber uint32) constant returns(submittedAt uint256)
func (_PlasmaBlockStorage *PlasmaBlockStorageSession) GetSubmissionTime(_blockNumber uint32) (*big.Int, error) {
	return _PlasmaBlockStorage.Contract.GetSubmissionTime(&_PlasmaBlockStorage.CallOpts, _blockNumber)
}

// GetSubmissionTime is a free data retrieval call binding the contract method 0xc3943844.
//
// Solidity: function getSubmissionTime(_blockNumber uint32) constant returns(submittedAt uint256)
func (_PlasmaBlockStorage *PlasmaBlockStorageCallerSession) GetSubmissionTime(_blockNumber uint32) (*big.Int, error) {
	return _PlasmaBlockStorage.Contract.GetSubmissionTime(&_PlasmaBlockStorage.CallOpts, _blockNumber)
}

// HashOfLastSubmittedBlock is a free data retrieval call binding the contract method 0x8649d43b.
//
// Solidity: function hashOfLastSubmittedBlock() constant returns(bytes32)
func (_PlasmaBlockStorage *PlasmaBlockStorageCaller) HashOfLastSubmittedBlock(opts *bind.CallOpts) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _PlasmaBlockStorage.contract.Call(opts, out, "hashOfLastSubmittedBlock")
	return *ret0, err
}

// HashOfLastSubmittedBlock is a free data retrieval call binding the contract method 0x8649d43b.
//
// Solidity: function hashOfLastSubmittedBlock() constant returns(bytes32)
func (_PlasmaBlockStorage *PlasmaBlockStorageSession) HashOfLastSubmittedBlock() ([32]byte, error) {
	return _PlasmaBlockStorage.Contract.HashOfLastSubmittedBlock(&_PlasmaBlockStorage.CallOpts)
}

// HashOfLastSubmittedBlock is a free data retrieval call binding the contract method 0x8649d43b.
//
// Solidity: function hashOfLastSubmittedBlock() constant returns(bytes32)
func (_PlasmaBlockStorage *PlasmaBlockStorageCallerSession) HashOfLastSubmittedBlock() ([32]byte, error) {
	return _PlasmaBlockStorage.Contract.HashOfLastSubmittedBlock(&_PlasmaBlockStorage.CallOpts)
}

// IsOperator is a free data retrieval call binding the contract method 0x6d70f7ae.
//
// Solidity: function isOperator(_operator address) constant returns(bool)
func (_PlasmaBlockStorage *PlasmaBlockStorageCaller) IsOperator(opts *bind.CallOpts, _operator common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _PlasmaBlockStorage.contract.Call(opts, out, "isOperator", _operator)
	return *ret0, err
}

// IsOperator is a free data retrieval call binding the contract method 0x6d70f7ae.
//
// Solidity: function isOperator(_operator address) constant returns(bool)
func (_PlasmaBlockStorage *PlasmaBlockStorageSession) IsOperator(_operator common.Address) (bool, error) {
	return _PlasmaBlockStorage.Contract.IsOperator(&_PlasmaBlockStorage.CallOpts, _operator)
}

// IsOperator is a free data retrieval call binding the contract method 0x6d70f7ae.
//
// Solidity: function isOperator(_operator address) constant returns(bool)
func (_PlasmaBlockStorage *PlasmaBlockStorageCallerSession) IsOperator(_operator common.Address) (bool, error) {
	return _PlasmaBlockStorage.Contract.IsOperator(&_PlasmaBlockStorage.CallOpts, _operator)
}

// LastBlockNumber is a free data retrieval call binding the contract method 0x2552317c.
//
// Solidity: function lastBlockNumber() constant returns(uint256)
func (_PlasmaBlockStorage *PlasmaBlockStorageCaller) LastBlockNumber(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _PlasmaBlockStorage.contract.Call(opts, out, "lastBlockNumber")
	return *ret0, err
}

// LastBlockNumber is a free data retrieval call binding the contract method 0x2552317c.
//
// Solidity: function lastBlockNumber() constant returns(uint256)
func (_PlasmaBlockStorage *PlasmaBlockStorageSession) LastBlockNumber() (*big.Int, error) {
	return _PlasmaBlockStorage.Contract.LastBlockNumber(&_PlasmaBlockStorage.CallOpts)
}

// LastBlockNumber is a free data retrieval call binding the contract method 0x2552317c.
//
// Solidity: function lastBlockNumber() constant returns(uint256)
func (_PlasmaBlockStorage *PlasmaBlockStorageCallerSession) LastBlockNumber() (*big.Int, error) {
	return _PlasmaBlockStorage.Contract.LastBlockNumber(&_PlasmaBlockStorage.CallOpts)
}

// Operators is a free data retrieval call binding the contract method 0x13e7c9d8.
//
// Solidity: function operators( address) constant returns(uint8)
func (_PlasmaBlockStorage *PlasmaBlockStorageCaller) Operators(opts *bind.CallOpts, arg0 common.Address) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _PlasmaBlockStorage.contract.Call(opts, out, "operators", arg0)
	return *ret0, err
}

// Operators is a free data retrieval call binding the contract method 0x13e7c9d8.
//
// Solidity: function operators( address) constant returns(uint8)
func (_PlasmaBlockStorage *PlasmaBlockStorageSession) Operators(arg0 common.Address) (uint8, error) {
	return _PlasmaBlockStorage.Contract.Operators(&_PlasmaBlockStorage.CallOpts, arg0)
}

// Operators is a free data retrieval call binding the contract method 0x13e7c9d8.
//
// Solidity: function operators( address) constant returns(uint8)
func (_PlasmaBlockStorage *PlasmaBlockStorageCallerSession) Operators(arg0 common.Address) (uint8, error) {
	return _PlasmaBlockStorage.Contract.Operators(&_PlasmaBlockStorage.CallOpts, arg0)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_PlasmaBlockStorage *PlasmaBlockStorageCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _PlasmaBlockStorage.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_PlasmaBlockStorage *PlasmaBlockStorageSession) Owner() (common.Address, error) {
	return _PlasmaBlockStorage.Contract.Owner(&_PlasmaBlockStorage.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_PlasmaBlockStorage *PlasmaBlockStorageCallerSession) Owner() (common.Address, error) {
	return _PlasmaBlockStorage.Contract.Owner(&_PlasmaBlockStorage.CallOpts)
}

// WeekOldBlockNumber is a free data retrieval call binding the contract method 0x57edff07.
//
// Solidity: function weekOldBlockNumber() constant returns(uint256)
func (_PlasmaBlockStorage *PlasmaBlockStorageCaller) WeekOldBlockNumber(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _PlasmaBlockStorage.contract.Call(opts, out, "weekOldBlockNumber")
	return *ret0, err
}

// WeekOldBlockNumber is a free data retrieval call binding the contract method 0x57edff07.
//
// Solidity: function weekOldBlockNumber() constant returns(uint256)
func (_PlasmaBlockStorage *PlasmaBlockStorageSession) WeekOldBlockNumber() (*big.Int, error) {
	return _PlasmaBlockStorage.Contract.WeekOldBlockNumber(&_PlasmaBlockStorage.CallOpts)
}

// WeekOldBlockNumber is a free data retrieval call binding the contract method 0x57edff07.
//
// Solidity: function weekOldBlockNumber() constant returns(uint256)
func (_PlasmaBlockStorage *PlasmaBlockStorageCallerSession) WeekOldBlockNumber() (*big.Int, error) {
	return _PlasmaBlockStorage.Contract.WeekOldBlockNumber(&_PlasmaBlockStorage.CallOpts)
}

// IncrementWeekOldCounter is a paid mutator transaction binding the contract method 0x32bc8990.
//
// Solidity: function incrementWeekOldCounter() returns()
func (_PlasmaBlockStorage *PlasmaBlockStorageTransactor) IncrementWeekOldCounter(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PlasmaBlockStorage.contract.Transact(opts, "incrementWeekOldCounter")
}

// IncrementWeekOldCounter is a paid mutator transaction binding the contract method 0x32bc8990.
//
// Solidity: function incrementWeekOldCounter() returns()
func (_PlasmaBlockStorage *PlasmaBlockStorageSession) IncrementWeekOldCounter() (*types.Transaction, error) {
	return _PlasmaBlockStorage.Contract.IncrementWeekOldCounter(&_PlasmaBlockStorage.TransactOpts)
}

// IncrementWeekOldCounter is a paid mutator transaction binding the contract method 0x32bc8990.
//
// Solidity: function incrementWeekOldCounter() returns()
func (_PlasmaBlockStorage *PlasmaBlockStorageTransactorSession) IncrementWeekOldCounter() (*types.Transaction, error) {
	return _PlasmaBlockStorage.Contract.IncrementWeekOldCounter(&_PlasmaBlockStorage.TransactOpts)
}

// SetOperator is a paid mutator transaction binding the contract method 0x77754136.
//
// Solidity: function setOperator(_op address, _status uint256) returns(success bool)
func (_PlasmaBlockStorage *PlasmaBlockStorageTransactor) SetOperator(opts *bind.TransactOpts, _op common.Address, _status *big.Int) (*types.Transaction, error) {
	return _PlasmaBlockStorage.contract.Transact(opts, "setOperator", _op, _status)
}

// SetOperator is a paid mutator transaction binding the contract method 0x77754136.
//
// Solidity: function setOperator(_op address, _status uint256) returns(success bool)
func (_PlasmaBlockStorage *PlasmaBlockStorageSession) SetOperator(_op common.Address, _status *big.Int) (*types.Transaction, error) {
	return _PlasmaBlockStorage.Contract.SetOperator(&_PlasmaBlockStorage.TransactOpts, _op, _status)
}

// SetOperator is a paid mutator transaction binding the contract method 0x77754136.
//
// Solidity: function setOperator(_op address, _status uint256) returns(success bool)
func (_PlasmaBlockStorage *PlasmaBlockStorageTransactorSession) SetOperator(_op common.Address, _status *big.Int) (*types.Transaction, error) {
	return _PlasmaBlockStorage.Contract.SetOperator(&_PlasmaBlockStorage.TransactOpts, _op, _status)
}

// SetOwner is a paid mutator transaction binding the contract method 0x13af4035.
//
// Solidity: function setOwner(_newOwner address) returns()
func (_PlasmaBlockStorage *PlasmaBlockStorageTransactor) SetOwner(opts *bind.TransactOpts, _newOwner common.Address) (*types.Transaction, error) {
	return _PlasmaBlockStorage.contract.Transact(opts, "setOwner", _newOwner)
}

// SetOwner is a paid mutator transaction binding the contract method 0x13af4035.
//
// Solidity: function setOwner(_newOwner address) returns()
func (_PlasmaBlockStorage *PlasmaBlockStorageSession) SetOwner(_newOwner common.Address) (*types.Transaction, error) {
	return _PlasmaBlockStorage.Contract.SetOwner(&_PlasmaBlockStorage.TransactOpts, _newOwner)
}

// SetOwner is a paid mutator transaction binding the contract method 0x13af4035.
//
// Solidity: function setOwner(_newOwner address) returns()
func (_PlasmaBlockStorage *PlasmaBlockStorageTransactorSession) SetOwner(_newOwner common.Address) (*types.Transaction, error) {
	return _PlasmaBlockStorage.Contract.SetOwner(&_PlasmaBlockStorage.TransactOpts, _newOwner)
}

// SubmitBlockHeaders is a paid mutator transaction binding the contract method 0x8d92ce46.
//
// Solidity: function submitBlockHeaders(_headers bytes) returns(success bool)
func (_PlasmaBlockStorage *PlasmaBlockStorageTransactor) SubmitBlockHeaders(opts *bind.TransactOpts, _headers []byte) (*types.Transaction, error) {
	return _PlasmaBlockStorage.contract.Transact(opts, "submitBlockHeaders", _headers)
}

// SubmitBlockHeaders is a paid mutator transaction binding the contract method 0x8d92ce46.
//
// Solidity: function submitBlockHeaders(_headers bytes) returns(success bool)
func (_PlasmaBlockStorage *PlasmaBlockStorageSession) SubmitBlockHeaders(_headers []byte) (*types.Transaction, error) {
	return _PlasmaBlockStorage.Contract.SubmitBlockHeaders(&_PlasmaBlockStorage.TransactOpts, _headers)
}

// SubmitBlockHeaders is a paid mutator transaction binding the contract method 0x8d92ce46.
//
// Solidity: function submitBlockHeaders(_headers bytes) returns(success bool)
func (_PlasmaBlockStorage *PlasmaBlockStorageTransactorSession) SubmitBlockHeaders(_headers []byte) (*types.Transaction, error) {
	return _PlasmaBlockStorage.Contract.SubmitBlockHeaders(&_PlasmaBlockStorage.TransactOpts, _headers)
}

// PlasmaBlockStorageBlockHeaderSubmittedIterator is returned from FilterBlockHeaderSubmitted and is used to iterate over the raw logs and unpacked data for BlockHeaderSubmitted events raised by the PlasmaBlockStorage contract.
type PlasmaBlockStorageBlockHeaderSubmittedIterator struct {
	Event *PlasmaBlockStorageBlockHeaderSubmitted // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *PlasmaBlockStorageBlockHeaderSubmittedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PlasmaBlockStorageBlockHeaderSubmitted)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(PlasmaBlockStorageBlockHeaderSubmitted)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *PlasmaBlockStorageBlockHeaderSubmittedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PlasmaBlockStorageBlockHeaderSubmittedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PlasmaBlockStorageBlockHeaderSubmitted represents a BlockHeaderSubmitted event raised by the PlasmaBlockStorage contract.
type PlasmaBlockStorageBlockHeaderSubmitted struct {
	BlockNumber *big.Int
	MerkleRoot  [32]byte
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterBlockHeaderSubmitted is a free log retrieval operation binding the contract event 0x2b21e5460f086f577c26c0ac6d9707fd893ade470ca9d6cc9bdc8e07dd2312d8.
//
// Solidity: e BlockHeaderSubmitted(_blockNumber indexed uint256, _merkleRoot indexed bytes32)
func (_PlasmaBlockStorage *PlasmaBlockStorageFilterer) FilterBlockHeaderSubmitted(opts *bind.FilterOpts, _blockNumber []*big.Int, _merkleRoot [][32]byte) (*PlasmaBlockStorageBlockHeaderSubmittedIterator, error) {

	var _blockNumberRule []interface{}
	for _, _blockNumberItem := range _blockNumber {
		_blockNumberRule = append(_blockNumberRule, _blockNumberItem)
	}
	var _merkleRootRule []interface{}
	for _, _merkleRootItem := range _merkleRoot {
		_merkleRootRule = append(_merkleRootRule, _merkleRootItem)
	}

	logs, sub, err := _PlasmaBlockStorage.contract.FilterLogs(opts, "BlockHeaderSubmitted", _blockNumberRule, _merkleRootRule)
	if err != nil {
		return nil, err
	}
	return &PlasmaBlockStorageBlockHeaderSubmittedIterator{contract: _PlasmaBlockStorage.contract, event: "BlockHeaderSubmitted", logs: logs, sub: sub}, nil
}

// WatchBlockHeaderSubmitted is a free log subscription operation binding the contract event 0x2b21e5460f086f577c26c0ac6d9707fd893ade470ca9d6cc9bdc8e07dd2312d8.
//
// Solidity: e BlockHeaderSubmitted(_blockNumber indexed uint256, _merkleRoot indexed bytes32)
func (_PlasmaBlockStorage *PlasmaBlockStorageFilterer) WatchBlockHeaderSubmitted(opts *bind.WatchOpts, sink chan<- *PlasmaBlockStorageBlockHeaderSubmitted, _blockNumber []*big.Int, _merkleRoot [][32]byte) (event.Subscription, error) {

	var _blockNumberRule []interface{}
	for _, _blockNumberItem := range _blockNumber {
		_blockNumberRule = append(_blockNumberRule, _blockNumberItem)
	}
	var _merkleRootRule []interface{}
	for _, _merkleRootItem := range _merkleRoot {
		_merkleRootRule = append(_merkleRootRule, _merkleRootItem)
	}

	logs, sub, err := _PlasmaBlockStorage.contract.WatchLogs(opts, "BlockHeaderSubmitted", _blockNumberRule, _merkleRootRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PlasmaBlockStorageBlockHeaderSubmitted)
				if err := _PlasmaBlockStorage.contract.UnpackLog(event, "BlockHeaderSubmitted", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}
