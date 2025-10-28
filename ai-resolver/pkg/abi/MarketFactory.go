// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package abi

import (
	"errors"
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
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// MarketFactoryMarket is an auto generated low-level Go binding around an user-defined struct.
type MarketFactoryMarket struct {
	Id              *big.Int
	Creator         common.Address
	Amm             common.Address
	CollateralToken common.Address
	CloseTime       *big.Int
	Category        string
	MetadataURI     string
	CreatorStake    *big.Int
	StakeRefunded   bool
	Status          uint8
}

// MarketFactoryMarketParams is an auto generated low-level Go binding around an user-defined struct.
type MarketFactoryMarketParams struct {
	CollateralToken common.Address
	CloseTime       *big.Int
	Category        string
	MetadataURI     string
	CreatorStake    *big.Int
}

// MarketFactoryMetaData contains all meta data concerning the MarketFactory contract.
var MarketFactoryMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_outcomeToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_feeSplitter\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_horizonPerks\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_horizonToken\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"allMarketIds\",\"inputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"createMarket\",\"inputs\":[{\"name\":\"params\",\"type\":\"tuple\",\"internalType\":\"structMarketFactory.MarketParams\",\"components\":[{\"name\":\"collateralToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"closeTime\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"category\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"metadataURI\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"creatorStake\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[{\"name\":\"marketId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"feeSplitter\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractFeeSplitter\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getActiveMarkets\",\"inputs\":[{\"name\":\"offset\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"limit\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple[]\",\"internalType\":\"structMarketFactory.Market[]\",\"components\":[{\"name\":\"id\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"creator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amm\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"collateralToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"closeTime\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"category\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"metadataURI\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"creatorStake\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"stakeRefunded\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"status\",\"type\":\"uint8\",\"internalType\":\"enumMarketFactory.MarketStatus\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getAllMarketIds\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getMarket\",\"inputs\":[{\"name\":\"marketId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structMarketFactory.Market\",\"components\":[{\"name\":\"id\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"creator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amm\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"collateralToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"closeTime\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"category\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"metadataURI\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"creatorStake\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"stakeRefunded\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"status\",\"type\":\"uint8\",\"internalType\":\"enumMarketFactory.MarketStatus\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getMarketCount\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getMarketIdsByCategory\",\"inputs\":[{\"name\":\"category\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getMarketIdsByCreator\",\"inputs\":[{\"name\":\"creator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getMarkets\",\"inputs\":[{\"name\":\"offset\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"limit\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple[]\",\"internalType\":\"structMarketFactory.Market[]\",\"components\":[{\"name\":\"id\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"creator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amm\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"collateralToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"closeTime\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"category\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"metadataURI\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"creatorStake\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"stakeRefunded\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"status\",\"type\":\"uint8\",\"internalType\":\"enumMarketFactory.MarketStatus\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"horizonPerks\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractHorizonPerks\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"horizonToken\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractHorizonToken\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"marketExists\",\"inputs\":[{\"name\":\"marketId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"markets\",\"inputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"id\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"creator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amm\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"collateralToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"closeTime\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"category\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"metadataURI\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"creatorStake\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"stakeRefunded\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"status\",\"type\":\"uint8\",\"internalType\":\"enumMarketFactory.MarketStatus\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"marketsByCategory\",\"inputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"marketsByCreator\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"minCreatorStake\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"nextMarketId\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"outcomeToken\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractOutcomeToken\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"refundCreatorStake\",\"inputs\":[{\"name\":\"marketId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setMinCreatorStake\",\"inputs\":[{\"name\":\"newMinStake\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"updateMarketStatus\",\"inputs\":[{\"name\":\"marketId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"CreatorStakeRefunded\",\"inputs\":[{\"name\":\"marketId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"creator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"MarketCreated\",\"inputs\":[{\"name\":\"marketId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"creator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"ammAddress\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"collateralToken\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"closeTime\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"category\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"metadataURI\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"creatorStake\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"MarketStatusUpdated\",\"inputs\":[{\"name\":\"marketId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"oldStatus\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"enumMarketFactory.MarketStatus\"},{\"name\":\"newStatus\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"enumMarketFactory.MarketStatus\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"MinCreatorStakeUpdated\",\"inputs\":[{\"name\":\"oldStake\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"newStake\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"InvalidAddress\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidCategory\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidCloseTime\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidCollateral\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidCreatorStake\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"MarketDoesNotExist\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"MarketNotResolved\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotMarketCreator\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OwnableInvalidOwner\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"OwnableUnauthorizedAccount\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ReentrancyGuardReentrantCall\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"SafeERC20FailedOperation\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"StakeAlreadyClaimed\",\"inputs\":[]}]",
}

// MarketFactoryABI is the input ABI used to generate the binding from.
// Deprecated: Use MarketFactoryMetaData.ABI instead.
var MarketFactoryABI = MarketFactoryMetaData.ABI

// MarketFactory is an auto generated Go binding around an Ethereum contract.
type MarketFactory struct {
	MarketFactoryCaller     // Read-only binding to the contract
	MarketFactoryTransactor // Write-only binding to the contract
	MarketFactoryFilterer   // Log filterer for contract events
}

// MarketFactoryCaller is an auto generated read-only Go binding around an Ethereum contract.
type MarketFactoryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MarketFactoryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MarketFactoryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MarketFactoryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MarketFactoryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MarketFactorySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MarketFactorySession struct {
	Contract     *MarketFactory    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MarketFactoryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MarketFactoryCallerSession struct {
	Contract *MarketFactoryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// MarketFactoryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MarketFactoryTransactorSession struct {
	Contract     *MarketFactoryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// MarketFactoryRaw is an auto generated low-level Go binding around an Ethereum contract.
type MarketFactoryRaw struct {
	Contract *MarketFactory // Generic contract binding to access the raw methods on
}

// MarketFactoryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MarketFactoryCallerRaw struct {
	Contract *MarketFactoryCaller // Generic read-only contract binding to access the raw methods on
}

// MarketFactoryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MarketFactoryTransactorRaw struct {
	Contract *MarketFactoryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMarketFactory creates a new instance of MarketFactory, bound to a specific deployed contract.
func NewMarketFactory(address common.Address, backend bind.ContractBackend) (*MarketFactory, error) {
	contract, err := bindMarketFactory(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MarketFactory{MarketFactoryCaller: MarketFactoryCaller{contract: contract}, MarketFactoryTransactor: MarketFactoryTransactor{contract: contract}, MarketFactoryFilterer: MarketFactoryFilterer{contract: contract}}, nil
}

// NewMarketFactoryCaller creates a new read-only instance of MarketFactory, bound to a specific deployed contract.
func NewMarketFactoryCaller(address common.Address, caller bind.ContractCaller) (*MarketFactoryCaller, error) {
	contract, err := bindMarketFactory(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MarketFactoryCaller{contract: contract}, nil
}

// NewMarketFactoryTransactor creates a new write-only instance of MarketFactory, bound to a specific deployed contract.
func NewMarketFactoryTransactor(address common.Address, transactor bind.ContractTransactor) (*MarketFactoryTransactor, error) {
	contract, err := bindMarketFactory(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MarketFactoryTransactor{contract: contract}, nil
}

// NewMarketFactoryFilterer creates a new log filterer instance of MarketFactory, bound to a specific deployed contract.
func NewMarketFactoryFilterer(address common.Address, filterer bind.ContractFilterer) (*MarketFactoryFilterer, error) {
	contract, err := bindMarketFactory(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MarketFactoryFilterer{contract: contract}, nil
}

// bindMarketFactory binds a generic wrapper to an already deployed contract.
func bindMarketFactory(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := MarketFactoryMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MarketFactory *MarketFactoryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MarketFactory.Contract.MarketFactoryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MarketFactory *MarketFactoryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MarketFactory.Contract.MarketFactoryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MarketFactory *MarketFactoryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MarketFactory.Contract.MarketFactoryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MarketFactory *MarketFactoryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MarketFactory.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MarketFactory *MarketFactoryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MarketFactory.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MarketFactory *MarketFactoryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MarketFactory.Contract.contract.Transact(opts, method, params...)
}

// AllMarketIds is a free data retrieval call binding the contract method 0x23188024.
//
// Solidity: function allMarketIds(uint256 ) view returns(uint256)
func (_MarketFactory *MarketFactoryCaller) AllMarketIds(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _MarketFactory.contract.Call(opts, &out, "allMarketIds", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AllMarketIds is a free data retrieval call binding the contract method 0x23188024.
//
// Solidity: function allMarketIds(uint256 ) view returns(uint256)
func (_MarketFactory *MarketFactorySession) AllMarketIds(arg0 *big.Int) (*big.Int, error) {
	return _MarketFactory.Contract.AllMarketIds(&_MarketFactory.CallOpts, arg0)
}

// AllMarketIds is a free data retrieval call binding the contract method 0x23188024.
//
// Solidity: function allMarketIds(uint256 ) view returns(uint256)
func (_MarketFactory *MarketFactoryCallerSession) AllMarketIds(arg0 *big.Int) (*big.Int, error) {
	return _MarketFactory.Contract.AllMarketIds(&_MarketFactory.CallOpts, arg0)
}

// FeeSplitter is a free data retrieval call binding the contract method 0x6052970c.
//
// Solidity: function feeSplitter() view returns(address)
func (_MarketFactory *MarketFactoryCaller) FeeSplitter(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MarketFactory.contract.Call(opts, &out, "feeSplitter")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FeeSplitter is a free data retrieval call binding the contract method 0x6052970c.
//
// Solidity: function feeSplitter() view returns(address)
func (_MarketFactory *MarketFactorySession) FeeSplitter() (common.Address, error) {
	return _MarketFactory.Contract.FeeSplitter(&_MarketFactory.CallOpts)
}

// FeeSplitter is a free data retrieval call binding the contract method 0x6052970c.
//
// Solidity: function feeSplitter() view returns(address)
func (_MarketFactory *MarketFactoryCallerSession) FeeSplitter() (common.Address, error) {
	return _MarketFactory.Contract.FeeSplitter(&_MarketFactory.CallOpts)
}

// GetActiveMarkets is a free data retrieval call binding the contract method 0xa04ddcad.
//
// Solidity: function getActiveMarkets(uint256 offset, uint256 limit) view returns((uint256,address,address,address,uint256,string,string,uint256,bool,uint8)[])
func (_MarketFactory *MarketFactoryCaller) GetActiveMarkets(opts *bind.CallOpts, offset *big.Int, limit *big.Int) ([]MarketFactoryMarket, error) {
	var out []interface{}
	err := _MarketFactory.contract.Call(opts, &out, "getActiveMarkets", offset, limit)

	if err != nil {
		return *new([]MarketFactoryMarket), err
	}

	out0 := *abi.ConvertType(out[0], new([]MarketFactoryMarket)).(*[]MarketFactoryMarket)

	return out0, err

}

// GetActiveMarkets is a free data retrieval call binding the contract method 0xa04ddcad.
//
// Solidity: function getActiveMarkets(uint256 offset, uint256 limit) view returns((uint256,address,address,address,uint256,string,string,uint256,bool,uint8)[])
func (_MarketFactory *MarketFactorySession) GetActiveMarkets(offset *big.Int, limit *big.Int) ([]MarketFactoryMarket, error) {
	return _MarketFactory.Contract.GetActiveMarkets(&_MarketFactory.CallOpts, offset, limit)
}

// GetActiveMarkets is a free data retrieval call binding the contract method 0xa04ddcad.
//
// Solidity: function getActiveMarkets(uint256 offset, uint256 limit) view returns((uint256,address,address,address,uint256,string,string,uint256,bool,uint8)[])
func (_MarketFactory *MarketFactoryCallerSession) GetActiveMarkets(offset *big.Int, limit *big.Int) ([]MarketFactoryMarket, error) {
	return _MarketFactory.Contract.GetActiveMarkets(&_MarketFactory.CallOpts, offset, limit)
}

// GetAllMarketIds is a free data retrieval call binding the contract method 0xb85ed636.
//
// Solidity: function getAllMarketIds() view returns(uint256[])
func (_MarketFactory *MarketFactoryCaller) GetAllMarketIds(opts *bind.CallOpts) ([]*big.Int, error) {
	var out []interface{}
	err := _MarketFactory.contract.Call(opts, &out, "getAllMarketIds")

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetAllMarketIds is a free data retrieval call binding the contract method 0xb85ed636.
//
// Solidity: function getAllMarketIds() view returns(uint256[])
func (_MarketFactory *MarketFactorySession) GetAllMarketIds() ([]*big.Int, error) {
	return _MarketFactory.Contract.GetAllMarketIds(&_MarketFactory.CallOpts)
}

// GetAllMarketIds is a free data retrieval call binding the contract method 0xb85ed636.
//
// Solidity: function getAllMarketIds() view returns(uint256[])
func (_MarketFactory *MarketFactoryCallerSession) GetAllMarketIds() ([]*big.Int, error) {
	return _MarketFactory.Contract.GetAllMarketIds(&_MarketFactory.CallOpts)
}

// GetMarket is a free data retrieval call binding the contract method 0xeb44fdd3.
//
// Solidity: function getMarket(uint256 marketId) view returns((uint256,address,address,address,uint256,string,string,uint256,bool,uint8))
func (_MarketFactory *MarketFactoryCaller) GetMarket(opts *bind.CallOpts, marketId *big.Int) (MarketFactoryMarket, error) {
	var out []interface{}
	err := _MarketFactory.contract.Call(opts, &out, "getMarket", marketId)

	if err != nil {
		return *new(MarketFactoryMarket), err
	}

	out0 := *abi.ConvertType(out[0], new(MarketFactoryMarket)).(*MarketFactoryMarket)

	return out0, err

}

// GetMarket is a free data retrieval call binding the contract method 0xeb44fdd3.
//
// Solidity: function getMarket(uint256 marketId) view returns((uint256,address,address,address,uint256,string,string,uint256,bool,uint8))
func (_MarketFactory *MarketFactorySession) GetMarket(marketId *big.Int) (MarketFactoryMarket, error) {
	return _MarketFactory.Contract.GetMarket(&_MarketFactory.CallOpts, marketId)
}

// GetMarket is a free data retrieval call binding the contract method 0xeb44fdd3.
//
// Solidity: function getMarket(uint256 marketId) view returns((uint256,address,address,address,uint256,string,string,uint256,bool,uint8))
func (_MarketFactory *MarketFactoryCallerSession) GetMarket(marketId *big.Int) (MarketFactoryMarket, error) {
	return _MarketFactory.Contract.GetMarket(&_MarketFactory.CallOpts, marketId)
}

// GetMarketCount is a free data retrieval call binding the contract method 0xfd69f3c2.
//
// Solidity: function getMarketCount() view returns(uint256)
func (_MarketFactory *MarketFactoryCaller) GetMarketCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MarketFactory.contract.Call(opts, &out, "getMarketCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetMarketCount is a free data retrieval call binding the contract method 0xfd69f3c2.
//
// Solidity: function getMarketCount() view returns(uint256)
func (_MarketFactory *MarketFactorySession) GetMarketCount() (*big.Int, error) {
	return _MarketFactory.Contract.GetMarketCount(&_MarketFactory.CallOpts)
}

// GetMarketCount is a free data retrieval call binding the contract method 0xfd69f3c2.
//
// Solidity: function getMarketCount() view returns(uint256)
func (_MarketFactory *MarketFactoryCallerSession) GetMarketCount() (*big.Int, error) {
	return _MarketFactory.Contract.GetMarketCount(&_MarketFactory.CallOpts)
}

// GetMarketIdsByCategory is a free data retrieval call binding the contract method 0xd893f617.
//
// Solidity: function getMarketIdsByCategory(string category) view returns(uint256[])
func (_MarketFactory *MarketFactoryCaller) GetMarketIdsByCategory(opts *bind.CallOpts, category string) ([]*big.Int, error) {
	var out []interface{}
	err := _MarketFactory.contract.Call(opts, &out, "getMarketIdsByCategory", category)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetMarketIdsByCategory is a free data retrieval call binding the contract method 0xd893f617.
//
// Solidity: function getMarketIdsByCategory(string category) view returns(uint256[])
func (_MarketFactory *MarketFactorySession) GetMarketIdsByCategory(category string) ([]*big.Int, error) {
	return _MarketFactory.Contract.GetMarketIdsByCategory(&_MarketFactory.CallOpts, category)
}

// GetMarketIdsByCategory is a free data retrieval call binding the contract method 0xd893f617.
//
// Solidity: function getMarketIdsByCategory(string category) view returns(uint256[])
func (_MarketFactory *MarketFactoryCallerSession) GetMarketIdsByCategory(category string) ([]*big.Int, error) {
	return _MarketFactory.Contract.GetMarketIdsByCategory(&_MarketFactory.CallOpts, category)
}

// GetMarketIdsByCreator is a free data retrieval call binding the contract method 0xb51c6ca9.
//
// Solidity: function getMarketIdsByCreator(address creator) view returns(uint256[])
func (_MarketFactory *MarketFactoryCaller) GetMarketIdsByCreator(opts *bind.CallOpts, creator common.Address) ([]*big.Int, error) {
	var out []interface{}
	err := _MarketFactory.contract.Call(opts, &out, "getMarketIdsByCreator", creator)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetMarketIdsByCreator is a free data retrieval call binding the contract method 0xb51c6ca9.
//
// Solidity: function getMarketIdsByCreator(address creator) view returns(uint256[])
func (_MarketFactory *MarketFactorySession) GetMarketIdsByCreator(creator common.Address) ([]*big.Int, error) {
	return _MarketFactory.Contract.GetMarketIdsByCreator(&_MarketFactory.CallOpts, creator)
}

// GetMarketIdsByCreator is a free data retrieval call binding the contract method 0xb51c6ca9.
//
// Solidity: function getMarketIdsByCreator(address creator) view returns(uint256[])
func (_MarketFactory *MarketFactoryCallerSession) GetMarketIdsByCreator(creator common.Address) ([]*big.Int, error) {
	return _MarketFactory.Contract.GetMarketIdsByCreator(&_MarketFactory.CallOpts, creator)
}

// GetMarkets is a free data retrieval call binding the contract method 0x80968d48.
//
// Solidity: function getMarkets(uint256 offset, uint256 limit) view returns((uint256,address,address,address,uint256,string,string,uint256,bool,uint8)[])
func (_MarketFactory *MarketFactoryCaller) GetMarkets(opts *bind.CallOpts, offset *big.Int, limit *big.Int) ([]MarketFactoryMarket, error) {
	var out []interface{}
	err := _MarketFactory.contract.Call(opts, &out, "getMarkets", offset, limit)

	if err != nil {
		return *new([]MarketFactoryMarket), err
	}

	out0 := *abi.ConvertType(out[0], new([]MarketFactoryMarket)).(*[]MarketFactoryMarket)

	return out0, err

}

// GetMarkets is a free data retrieval call binding the contract method 0x80968d48.
//
// Solidity: function getMarkets(uint256 offset, uint256 limit) view returns((uint256,address,address,address,uint256,string,string,uint256,bool,uint8)[])
func (_MarketFactory *MarketFactorySession) GetMarkets(offset *big.Int, limit *big.Int) ([]MarketFactoryMarket, error) {
	return _MarketFactory.Contract.GetMarkets(&_MarketFactory.CallOpts, offset, limit)
}

// GetMarkets is a free data retrieval call binding the contract method 0x80968d48.
//
// Solidity: function getMarkets(uint256 offset, uint256 limit) view returns((uint256,address,address,address,uint256,string,string,uint256,bool,uint8)[])
func (_MarketFactory *MarketFactoryCallerSession) GetMarkets(offset *big.Int, limit *big.Int) ([]MarketFactoryMarket, error) {
	return _MarketFactory.Contract.GetMarkets(&_MarketFactory.CallOpts, offset, limit)
}

// HorizonPerks is a free data retrieval call binding the contract method 0xffe02e34.
//
// Solidity: function horizonPerks() view returns(address)
func (_MarketFactory *MarketFactoryCaller) HorizonPerks(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MarketFactory.contract.Call(opts, &out, "horizonPerks")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// HorizonPerks is a free data retrieval call binding the contract method 0xffe02e34.
//
// Solidity: function horizonPerks() view returns(address)
func (_MarketFactory *MarketFactorySession) HorizonPerks() (common.Address, error) {
	return _MarketFactory.Contract.HorizonPerks(&_MarketFactory.CallOpts)
}

// HorizonPerks is a free data retrieval call binding the contract method 0xffe02e34.
//
// Solidity: function horizonPerks() view returns(address)
func (_MarketFactory *MarketFactoryCallerSession) HorizonPerks() (common.Address, error) {
	return _MarketFactory.Contract.HorizonPerks(&_MarketFactory.CallOpts)
}

// HorizonToken is a free data retrieval call binding the contract method 0x326046b1.
//
// Solidity: function horizonToken() view returns(address)
func (_MarketFactory *MarketFactoryCaller) HorizonToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MarketFactory.contract.Call(opts, &out, "horizonToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// HorizonToken is a free data retrieval call binding the contract method 0x326046b1.
//
// Solidity: function horizonToken() view returns(address)
func (_MarketFactory *MarketFactorySession) HorizonToken() (common.Address, error) {
	return _MarketFactory.Contract.HorizonToken(&_MarketFactory.CallOpts)
}

// HorizonToken is a free data retrieval call binding the contract method 0x326046b1.
//
// Solidity: function horizonToken() view returns(address)
func (_MarketFactory *MarketFactoryCallerSession) HorizonToken() (common.Address, error) {
	return _MarketFactory.Contract.HorizonToken(&_MarketFactory.CallOpts)
}

// MarketExists is a free data retrieval call binding the contract method 0xec69a654.
//
// Solidity: function marketExists(uint256 marketId) view returns(bool)
func (_MarketFactory *MarketFactoryCaller) MarketExists(opts *bind.CallOpts, marketId *big.Int) (bool, error) {
	var out []interface{}
	err := _MarketFactory.contract.Call(opts, &out, "marketExists", marketId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// MarketExists is a free data retrieval call binding the contract method 0xec69a654.
//
// Solidity: function marketExists(uint256 marketId) view returns(bool)
func (_MarketFactory *MarketFactorySession) MarketExists(marketId *big.Int) (bool, error) {
	return _MarketFactory.Contract.MarketExists(&_MarketFactory.CallOpts, marketId)
}

// MarketExists is a free data retrieval call binding the contract method 0xec69a654.
//
// Solidity: function marketExists(uint256 marketId) view returns(bool)
func (_MarketFactory *MarketFactoryCallerSession) MarketExists(marketId *big.Int) (bool, error) {
	return _MarketFactory.Contract.MarketExists(&_MarketFactory.CallOpts, marketId)
}

// Markets is a free data retrieval call binding the contract method 0xb1283e77.
//
// Solidity: function markets(uint256 ) view returns(uint256 id, address creator, address amm, address collateralToken, uint256 closeTime, string category, string metadataURI, uint256 creatorStake, bool stakeRefunded, uint8 status)
func (_MarketFactory *MarketFactoryCaller) Markets(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Id              *big.Int
	Creator         common.Address
	Amm             common.Address
	CollateralToken common.Address
	CloseTime       *big.Int
	Category        string
	MetadataURI     string
	CreatorStake    *big.Int
	StakeRefunded   bool
	Status          uint8
}, error) {
	var out []interface{}
	err := _MarketFactory.contract.Call(opts, &out, "markets", arg0)

	outstruct := new(struct {
		Id              *big.Int
		Creator         common.Address
		Amm             common.Address
		CollateralToken common.Address
		CloseTime       *big.Int
		Category        string
		MetadataURI     string
		CreatorStake    *big.Int
		StakeRefunded   bool
		Status          uint8
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Id = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Creator = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.Amm = *abi.ConvertType(out[2], new(common.Address)).(*common.Address)
	outstruct.CollateralToken = *abi.ConvertType(out[3], new(common.Address)).(*common.Address)
	outstruct.CloseTime = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.Category = *abi.ConvertType(out[5], new(string)).(*string)
	outstruct.MetadataURI = *abi.ConvertType(out[6], new(string)).(*string)
	outstruct.CreatorStake = *abi.ConvertType(out[7], new(*big.Int)).(**big.Int)
	outstruct.StakeRefunded = *abi.ConvertType(out[8], new(bool)).(*bool)
	outstruct.Status = *abi.ConvertType(out[9], new(uint8)).(*uint8)

	return *outstruct, err

}

// Markets is a free data retrieval call binding the contract method 0xb1283e77.
//
// Solidity: function markets(uint256 ) view returns(uint256 id, address creator, address amm, address collateralToken, uint256 closeTime, string category, string metadataURI, uint256 creatorStake, bool stakeRefunded, uint8 status)
func (_MarketFactory *MarketFactorySession) Markets(arg0 *big.Int) (struct {
	Id              *big.Int
	Creator         common.Address
	Amm             common.Address
	CollateralToken common.Address
	CloseTime       *big.Int
	Category        string
	MetadataURI     string
	CreatorStake    *big.Int
	StakeRefunded   bool
	Status          uint8
}, error) {
	return _MarketFactory.Contract.Markets(&_MarketFactory.CallOpts, arg0)
}

// Markets is a free data retrieval call binding the contract method 0xb1283e77.
//
// Solidity: function markets(uint256 ) view returns(uint256 id, address creator, address amm, address collateralToken, uint256 closeTime, string category, string metadataURI, uint256 creatorStake, bool stakeRefunded, uint8 status)
func (_MarketFactory *MarketFactoryCallerSession) Markets(arg0 *big.Int) (struct {
	Id              *big.Int
	Creator         common.Address
	Amm             common.Address
	CollateralToken common.Address
	CloseTime       *big.Int
	Category        string
	MetadataURI     string
	CreatorStake    *big.Int
	StakeRefunded   bool
	Status          uint8
}, error) {
	return _MarketFactory.Contract.Markets(&_MarketFactory.CallOpts, arg0)
}

// MarketsByCategory is a free data retrieval call binding the contract method 0xaff14fe8.
//
// Solidity: function marketsByCategory(string , uint256 ) view returns(uint256)
func (_MarketFactory *MarketFactoryCaller) MarketsByCategory(opts *bind.CallOpts, arg0 string, arg1 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _MarketFactory.contract.Call(opts, &out, "marketsByCategory", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MarketsByCategory is a free data retrieval call binding the contract method 0xaff14fe8.
//
// Solidity: function marketsByCategory(string , uint256 ) view returns(uint256)
func (_MarketFactory *MarketFactorySession) MarketsByCategory(arg0 string, arg1 *big.Int) (*big.Int, error) {
	return _MarketFactory.Contract.MarketsByCategory(&_MarketFactory.CallOpts, arg0, arg1)
}

// MarketsByCategory is a free data retrieval call binding the contract method 0xaff14fe8.
//
// Solidity: function marketsByCategory(string , uint256 ) view returns(uint256)
func (_MarketFactory *MarketFactoryCallerSession) MarketsByCategory(arg0 string, arg1 *big.Int) (*big.Int, error) {
	return _MarketFactory.Contract.MarketsByCategory(&_MarketFactory.CallOpts, arg0, arg1)
}

// MarketsByCreator is a free data retrieval call binding the contract method 0x13ba3d14.
//
// Solidity: function marketsByCreator(address , uint256 ) view returns(uint256)
func (_MarketFactory *MarketFactoryCaller) MarketsByCreator(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _MarketFactory.contract.Call(opts, &out, "marketsByCreator", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MarketsByCreator is a free data retrieval call binding the contract method 0x13ba3d14.
//
// Solidity: function marketsByCreator(address , uint256 ) view returns(uint256)
func (_MarketFactory *MarketFactorySession) MarketsByCreator(arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	return _MarketFactory.Contract.MarketsByCreator(&_MarketFactory.CallOpts, arg0, arg1)
}

// MarketsByCreator is a free data retrieval call binding the contract method 0x13ba3d14.
//
// Solidity: function marketsByCreator(address , uint256 ) view returns(uint256)
func (_MarketFactory *MarketFactoryCallerSession) MarketsByCreator(arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	return _MarketFactory.Contract.MarketsByCreator(&_MarketFactory.CallOpts, arg0, arg1)
}

// MinCreatorStake is a free data retrieval call binding the contract method 0xa395f823.
//
// Solidity: function minCreatorStake() view returns(uint256)
func (_MarketFactory *MarketFactoryCaller) MinCreatorStake(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MarketFactory.contract.Call(opts, &out, "minCreatorStake")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinCreatorStake is a free data retrieval call binding the contract method 0xa395f823.
//
// Solidity: function minCreatorStake() view returns(uint256)
func (_MarketFactory *MarketFactorySession) MinCreatorStake() (*big.Int, error) {
	return _MarketFactory.Contract.MinCreatorStake(&_MarketFactory.CallOpts)
}

// MinCreatorStake is a free data retrieval call binding the contract method 0xa395f823.
//
// Solidity: function minCreatorStake() view returns(uint256)
func (_MarketFactory *MarketFactoryCallerSession) MinCreatorStake() (*big.Int, error) {
	return _MarketFactory.Contract.MinCreatorStake(&_MarketFactory.CallOpts)
}

// NextMarketId is a free data retrieval call binding the contract method 0x406ef2ef.
//
// Solidity: function nextMarketId() view returns(uint256)
func (_MarketFactory *MarketFactoryCaller) NextMarketId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MarketFactory.contract.Call(opts, &out, "nextMarketId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NextMarketId is a free data retrieval call binding the contract method 0x406ef2ef.
//
// Solidity: function nextMarketId() view returns(uint256)
func (_MarketFactory *MarketFactorySession) NextMarketId() (*big.Int, error) {
	return _MarketFactory.Contract.NextMarketId(&_MarketFactory.CallOpts)
}

// NextMarketId is a free data retrieval call binding the contract method 0x406ef2ef.
//
// Solidity: function nextMarketId() view returns(uint256)
func (_MarketFactory *MarketFactoryCallerSession) NextMarketId() (*big.Int, error) {
	return _MarketFactory.Contract.NextMarketId(&_MarketFactory.CallOpts)
}

// OutcomeToken is a free data retrieval call binding the contract method 0xa998d6d8.
//
// Solidity: function outcomeToken() view returns(address)
func (_MarketFactory *MarketFactoryCaller) OutcomeToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MarketFactory.contract.Call(opts, &out, "outcomeToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OutcomeToken is a free data retrieval call binding the contract method 0xa998d6d8.
//
// Solidity: function outcomeToken() view returns(address)
func (_MarketFactory *MarketFactorySession) OutcomeToken() (common.Address, error) {
	return _MarketFactory.Contract.OutcomeToken(&_MarketFactory.CallOpts)
}

// OutcomeToken is a free data retrieval call binding the contract method 0xa998d6d8.
//
// Solidity: function outcomeToken() view returns(address)
func (_MarketFactory *MarketFactoryCallerSession) OutcomeToken() (common.Address, error) {
	return _MarketFactory.Contract.OutcomeToken(&_MarketFactory.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_MarketFactory *MarketFactoryCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MarketFactory.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_MarketFactory *MarketFactorySession) Owner() (common.Address, error) {
	return _MarketFactory.Contract.Owner(&_MarketFactory.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_MarketFactory *MarketFactoryCallerSession) Owner() (common.Address, error) {
	return _MarketFactory.Contract.Owner(&_MarketFactory.CallOpts)
}

// CreateMarket is a paid mutator transaction binding the contract method 0xf7cbc21c.
//
// Solidity: function createMarket((address,uint256,string,string,uint256) params) returns(uint256 marketId)
func (_MarketFactory *MarketFactoryTransactor) CreateMarket(opts *bind.TransactOpts, params MarketFactoryMarketParams) (*types.Transaction, error) {
	return _MarketFactory.contract.Transact(opts, "createMarket", params)
}

// CreateMarket is a paid mutator transaction binding the contract method 0xf7cbc21c.
//
// Solidity: function createMarket((address,uint256,string,string,uint256) params) returns(uint256 marketId)
func (_MarketFactory *MarketFactorySession) CreateMarket(params MarketFactoryMarketParams) (*types.Transaction, error) {
	return _MarketFactory.Contract.CreateMarket(&_MarketFactory.TransactOpts, params)
}

// CreateMarket is a paid mutator transaction binding the contract method 0xf7cbc21c.
//
// Solidity: function createMarket((address,uint256,string,string,uint256) params) returns(uint256 marketId)
func (_MarketFactory *MarketFactoryTransactorSession) CreateMarket(params MarketFactoryMarketParams) (*types.Transaction, error) {
	return _MarketFactory.Contract.CreateMarket(&_MarketFactory.TransactOpts, params)
}

// RefundCreatorStake is a paid mutator transaction binding the contract method 0x0c080604.
//
// Solidity: function refundCreatorStake(uint256 marketId) returns()
func (_MarketFactory *MarketFactoryTransactor) RefundCreatorStake(opts *bind.TransactOpts, marketId *big.Int) (*types.Transaction, error) {
	return _MarketFactory.contract.Transact(opts, "refundCreatorStake", marketId)
}

// RefundCreatorStake is a paid mutator transaction binding the contract method 0x0c080604.
//
// Solidity: function refundCreatorStake(uint256 marketId) returns()
func (_MarketFactory *MarketFactorySession) RefundCreatorStake(marketId *big.Int) (*types.Transaction, error) {
	return _MarketFactory.Contract.RefundCreatorStake(&_MarketFactory.TransactOpts, marketId)
}

// RefundCreatorStake is a paid mutator transaction binding the contract method 0x0c080604.
//
// Solidity: function refundCreatorStake(uint256 marketId) returns()
func (_MarketFactory *MarketFactoryTransactorSession) RefundCreatorStake(marketId *big.Int) (*types.Transaction, error) {
	return _MarketFactory.Contract.RefundCreatorStake(&_MarketFactory.TransactOpts, marketId)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_MarketFactory *MarketFactoryTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MarketFactory.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_MarketFactory *MarketFactorySession) RenounceOwnership() (*types.Transaction, error) {
	return _MarketFactory.Contract.RenounceOwnership(&_MarketFactory.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_MarketFactory *MarketFactoryTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _MarketFactory.Contract.RenounceOwnership(&_MarketFactory.TransactOpts)
}

// SetMinCreatorStake is a paid mutator transaction binding the contract method 0x31d1d650.
//
// Solidity: function setMinCreatorStake(uint256 newMinStake) returns()
func (_MarketFactory *MarketFactoryTransactor) SetMinCreatorStake(opts *bind.TransactOpts, newMinStake *big.Int) (*types.Transaction, error) {
	return _MarketFactory.contract.Transact(opts, "setMinCreatorStake", newMinStake)
}

// SetMinCreatorStake is a paid mutator transaction binding the contract method 0x31d1d650.
//
// Solidity: function setMinCreatorStake(uint256 newMinStake) returns()
func (_MarketFactory *MarketFactorySession) SetMinCreatorStake(newMinStake *big.Int) (*types.Transaction, error) {
	return _MarketFactory.Contract.SetMinCreatorStake(&_MarketFactory.TransactOpts, newMinStake)
}

// SetMinCreatorStake is a paid mutator transaction binding the contract method 0x31d1d650.
//
// Solidity: function setMinCreatorStake(uint256 newMinStake) returns()
func (_MarketFactory *MarketFactoryTransactorSession) SetMinCreatorStake(newMinStake *big.Int) (*types.Transaction, error) {
	return _MarketFactory.Contract.SetMinCreatorStake(&_MarketFactory.TransactOpts, newMinStake)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_MarketFactory *MarketFactoryTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _MarketFactory.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_MarketFactory *MarketFactorySession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _MarketFactory.Contract.TransferOwnership(&_MarketFactory.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_MarketFactory *MarketFactoryTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _MarketFactory.Contract.TransferOwnership(&_MarketFactory.TransactOpts, newOwner)
}

// UpdateMarketStatus is a paid mutator transaction binding the contract method 0x5710b285.
//
// Solidity: function updateMarketStatus(uint256 marketId) returns()
func (_MarketFactory *MarketFactoryTransactor) UpdateMarketStatus(opts *bind.TransactOpts, marketId *big.Int) (*types.Transaction, error) {
	return _MarketFactory.contract.Transact(opts, "updateMarketStatus", marketId)
}

// UpdateMarketStatus is a paid mutator transaction binding the contract method 0x5710b285.
//
// Solidity: function updateMarketStatus(uint256 marketId) returns()
func (_MarketFactory *MarketFactorySession) UpdateMarketStatus(marketId *big.Int) (*types.Transaction, error) {
	return _MarketFactory.Contract.UpdateMarketStatus(&_MarketFactory.TransactOpts, marketId)
}

// UpdateMarketStatus is a paid mutator transaction binding the contract method 0x5710b285.
//
// Solidity: function updateMarketStatus(uint256 marketId) returns()
func (_MarketFactory *MarketFactoryTransactorSession) UpdateMarketStatus(marketId *big.Int) (*types.Transaction, error) {
	return _MarketFactory.Contract.UpdateMarketStatus(&_MarketFactory.TransactOpts, marketId)
}

// MarketFactoryCreatorStakeRefundedIterator is returned from FilterCreatorStakeRefunded and is used to iterate over the raw logs and unpacked data for CreatorStakeRefunded events raised by the MarketFactory contract.
type MarketFactoryCreatorStakeRefundedIterator struct {
	Event *MarketFactoryCreatorStakeRefunded // Event containing the contract specifics and raw log

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
func (it *MarketFactoryCreatorStakeRefundedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MarketFactoryCreatorStakeRefunded)
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
		it.Event = new(MarketFactoryCreatorStakeRefunded)
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
func (it *MarketFactoryCreatorStakeRefundedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MarketFactoryCreatorStakeRefundedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MarketFactoryCreatorStakeRefunded represents a CreatorStakeRefunded event raised by the MarketFactory contract.
type MarketFactoryCreatorStakeRefunded struct {
	MarketId *big.Int
	Creator  common.Address
	Amount   *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterCreatorStakeRefunded is a free log retrieval operation binding the contract event 0xc070177e58281314067d38f9a79cd7989153b56d3388f9046a49c160e26c68f9.
//
// Solidity: event CreatorStakeRefunded(uint256 indexed marketId, address indexed creator, uint256 amount)
func (_MarketFactory *MarketFactoryFilterer) FilterCreatorStakeRefunded(opts *bind.FilterOpts, marketId []*big.Int, creator []common.Address) (*MarketFactoryCreatorStakeRefundedIterator, error) {

	var marketIdRule []interface{}
	for _, marketIdItem := range marketId {
		marketIdRule = append(marketIdRule, marketIdItem)
	}
	var creatorRule []interface{}
	for _, creatorItem := range creator {
		creatorRule = append(creatorRule, creatorItem)
	}

	logs, sub, err := _MarketFactory.contract.FilterLogs(opts, "CreatorStakeRefunded", marketIdRule, creatorRule)
	if err != nil {
		return nil, err
	}
	return &MarketFactoryCreatorStakeRefundedIterator{contract: _MarketFactory.contract, event: "CreatorStakeRefunded", logs: logs, sub: sub}, nil
}

// WatchCreatorStakeRefunded is a free log subscription operation binding the contract event 0xc070177e58281314067d38f9a79cd7989153b56d3388f9046a49c160e26c68f9.
//
// Solidity: event CreatorStakeRefunded(uint256 indexed marketId, address indexed creator, uint256 amount)
func (_MarketFactory *MarketFactoryFilterer) WatchCreatorStakeRefunded(opts *bind.WatchOpts, sink chan<- *MarketFactoryCreatorStakeRefunded, marketId []*big.Int, creator []common.Address) (event.Subscription, error) {

	var marketIdRule []interface{}
	for _, marketIdItem := range marketId {
		marketIdRule = append(marketIdRule, marketIdItem)
	}
	var creatorRule []interface{}
	for _, creatorItem := range creator {
		creatorRule = append(creatorRule, creatorItem)
	}

	logs, sub, err := _MarketFactory.contract.WatchLogs(opts, "CreatorStakeRefunded", marketIdRule, creatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MarketFactoryCreatorStakeRefunded)
				if err := _MarketFactory.contract.UnpackLog(event, "CreatorStakeRefunded", log); err != nil {
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

// ParseCreatorStakeRefunded is a log parse operation binding the contract event 0xc070177e58281314067d38f9a79cd7989153b56d3388f9046a49c160e26c68f9.
//
// Solidity: event CreatorStakeRefunded(uint256 indexed marketId, address indexed creator, uint256 amount)
func (_MarketFactory *MarketFactoryFilterer) ParseCreatorStakeRefunded(log types.Log) (*MarketFactoryCreatorStakeRefunded, error) {
	event := new(MarketFactoryCreatorStakeRefunded)
	if err := _MarketFactory.contract.UnpackLog(event, "CreatorStakeRefunded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MarketFactoryMarketCreatedIterator is returned from FilterMarketCreated and is used to iterate over the raw logs and unpacked data for MarketCreated events raised by the MarketFactory contract.
type MarketFactoryMarketCreatedIterator struct {
	Event *MarketFactoryMarketCreated // Event containing the contract specifics and raw log

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
func (it *MarketFactoryMarketCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MarketFactoryMarketCreated)
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
		it.Event = new(MarketFactoryMarketCreated)
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
func (it *MarketFactoryMarketCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MarketFactoryMarketCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MarketFactoryMarketCreated represents a MarketCreated event raised by the MarketFactory contract.
type MarketFactoryMarketCreated struct {
	MarketId        *big.Int
	Creator         common.Address
	AmmAddress      common.Address
	CollateralToken common.Address
	CloseTime       *big.Int
	Category        string
	MetadataURI     string
	CreatorStake    *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterMarketCreated is a free log retrieval operation binding the contract event 0xeffb0542c4d33ec0e0dd841f6551917e437c10514420b319620b21638298ca0a.
//
// Solidity: event MarketCreated(uint256 indexed marketId, address indexed creator, address indexed ammAddress, address collateralToken, uint256 closeTime, string category, string metadataURI, uint256 creatorStake)
func (_MarketFactory *MarketFactoryFilterer) FilterMarketCreated(opts *bind.FilterOpts, marketId []*big.Int, creator []common.Address, ammAddress []common.Address) (*MarketFactoryMarketCreatedIterator, error) {

	var marketIdRule []interface{}
	for _, marketIdItem := range marketId {
		marketIdRule = append(marketIdRule, marketIdItem)
	}
	var creatorRule []interface{}
	for _, creatorItem := range creator {
		creatorRule = append(creatorRule, creatorItem)
	}
	var ammAddressRule []interface{}
	for _, ammAddressItem := range ammAddress {
		ammAddressRule = append(ammAddressRule, ammAddressItem)
	}

	logs, sub, err := _MarketFactory.contract.FilterLogs(opts, "MarketCreated", marketIdRule, creatorRule, ammAddressRule)
	if err != nil {
		return nil, err
	}
	return &MarketFactoryMarketCreatedIterator{contract: _MarketFactory.contract, event: "MarketCreated", logs: logs, sub: sub}, nil
}

// WatchMarketCreated is a free log subscription operation binding the contract event 0xeffb0542c4d33ec0e0dd841f6551917e437c10514420b319620b21638298ca0a.
//
// Solidity: event MarketCreated(uint256 indexed marketId, address indexed creator, address indexed ammAddress, address collateralToken, uint256 closeTime, string category, string metadataURI, uint256 creatorStake)
func (_MarketFactory *MarketFactoryFilterer) WatchMarketCreated(opts *bind.WatchOpts, sink chan<- *MarketFactoryMarketCreated, marketId []*big.Int, creator []common.Address, ammAddress []common.Address) (event.Subscription, error) {

	var marketIdRule []interface{}
	for _, marketIdItem := range marketId {
		marketIdRule = append(marketIdRule, marketIdItem)
	}
	var creatorRule []interface{}
	for _, creatorItem := range creator {
		creatorRule = append(creatorRule, creatorItem)
	}
	var ammAddressRule []interface{}
	for _, ammAddressItem := range ammAddress {
		ammAddressRule = append(ammAddressRule, ammAddressItem)
	}

	logs, sub, err := _MarketFactory.contract.WatchLogs(opts, "MarketCreated", marketIdRule, creatorRule, ammAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MarketFactoryMarketCreated)
				if err := _MarketFactory.contract.UnpackLog(event, "MarketCreated", log); err != nil {
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

// ParseMarketCreated is a log parse operation binding the contract event 0xeffb0542c4d33ec0e0dd841f6551917e437c10514420b319620b21638298ca0a.
//
// Solidity: event MarketCreated(uint256 indexed marketId, address indexed creator, address indexed ammAddress, address collateralToken, uint256 closeTime, string category, string metadataURI, uint256 creatorStake)
func (_MarketFactory *MarketFactoryFilterer) ParseMarketCreated(log types.Log) (*MarketFactoryMarketCreated, error) {
	event := new(MarketFactoryMarketCreated)
	if err := _MarketFactory.contract.UnpackLog(event, "MarketCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MarketFactoryMarketStatusUpdatedIterator is returned from FilterMarketStatusUpdated and is used to iterate over the raw logs and unpacked data for MarketStatusUpdated events raised by the MarketFactory contract.
type MarketFactoryMarketStatusUpdatedIterator struct {
	Event *MarketFactoryMarketStatusUpdated // Event containing the contract specifics and raw log

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
func (it *MarketFactoryMarketStatusUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MarketFactoryMarketStatusUpdated)
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
		it.Event = new(MarketFactoryMarketStatusUpdated)
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
func (it *MarketFactoryMarketStatusUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MarketFactoryMarketStatusUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MarketFactoryMarketStatusUpdated represents a MarketStatusUpdated event raised by the MarketFactory contract.
type MarketFactoryMarketStatusUpdated struct {
	MarketId  *big.Int
	OldStatus uint8
	NewStatus uint8
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterMarketStatusUpdated is a free log retrieval operation binding the contract event 0xc30cef39d3c71bedadb4e96f142426c7ac02a36a432dd670af23a049a899005d.
//
// Solidity: event MarketStatusUpdated(uint256 indexed marketId, uint8 oldStatus, uint8 newStatus)
func (_MarketFactory *MarketFactoryFilterer) FilterMarketStatusUpdated(opts *bind.FilterOpts, marketId []*big.Int) (*MarketFactoryMarketStatusUpdatedIterator, error) {

	var marketIdRule []interface{}
	for _, marketIdItem := range marketId {
		marketIdRule = append(marketIdRule, marketIdItem)
	}

	logs, sub, err := _MarketFactory.contract.FilterLogs(opts, "MarketStatusUpdated", marketIdRule)
	if err != nil {
		return nil, err
	}
	return &MarketFactoryMarketStatusUpdatedIterator{contract: _MarketFactory.contract, event: "MarketStatusUpdated", logs: logs, sub: sub}, nil
}

// WatchMarketStatusUpdated is a free log subscription operation binding the contract event 0xc30cef39d3c71bedadb4e96f142426c7ac02a36a432dd670af23a049a899005d.
//
// Solidity: event MarketStatusUpdated(uint256 indexed marketId, uint8 oldStatus, uint8 newStatus)
func (_MarketFactory *MarketFactoryFilterer) WatchMarketStatusUpdated(opts *bind.WatchOpts, sink chan<- *MarketFactoryMarketStatusUpdated, marketId []*big.Int) (event.Subscription, error) {

	var marketIdRule []interface{}
	for _, marketIdItem := range marketId {
		marketIdRule = append(marketIdRule, marketIdItem)
	}

	logs, sub, err := _MarketFactory.contract.WatchLogs(opts, "MarketStatusUpdated", marketIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MarketFactoryMarketStatusUpdated)
				if err := _MarketFactory.contract.UnpackLog(event, "MarketStatusUpdated", log); err != nil {
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

// ParseMarketStatusUpdated is a log parse operation binding the contract event 0xc30cef39d3c71bedadb4e96f142426c7ac02a36a432dd670af23a049a899005d.
//
// Solidity: event MarketStatusUpdated(uint256 indexed marketId, uint8 oldStatus, uint8 newStatus)
func (_MarketFactory *MarketFactoryFilterer) ParseMarketStatusUpdated(log types.Log) (*MarketFactoryMarketStatusUpdated, error) {
	event := new(MarketFactoryMarketStatusUpdated)
	if err := _MarketFactory.contract.UnpackLog(event, "MarketStatusUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MarketFactoryMinCreatorStakeUpdatedIterator is returned from FilterMinCreatorStakeUpdated and is used to iterate over the raw logs and unpacked data for MinCreatorStakeUpdated events raised by the MarketFactory contract.
type MarketFactoryMinCreatorStakeUpdatedIterator struct {
	Event *MarketFactoryMinCreatorStakeUpdated // Event containing the contract specifics and raw log

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
func (it *MarketFactoryMinCreatorStakeUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MarketFactoryMinCreatorStakeUpdated)
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
		it.Event = new(MarketFactoryMinCreatorStakeUpdated)
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
func (it *MarketFactoryMinCreatorStakeUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MarketFactoryMinCreatorStakeUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MarketFactoryMinCreatorStakeUpdated represents a MinCreatorStakeUpdated event raised by the MarketFactory contract.
type MarketFactoryMinCreatorStakeUpdated struct {
	OldStake *big.Int
	NewStake *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterMinCreatorStakeUpdated is a free log retrieval operation binding the contract event 0xb2193ba7a75ea9a3671a8dda768d04598e7c2fed659cf26f16d2c7362e510d99.
//
// Solidity: event MinCreatorStakeUpdated(uint256 oldStake, uint256 newStake)
func (_MarketFactory *MarketFactoryFilterer) FilterMinCreatorStakeUpdated(opts *bind.FilterOpts) (*MarketFactoryMinCreatorStakeUpdatedIterator, error) {

	logs, sub, err := _MarketFactory.contract.FilterLogs(opts, "MinCreatorStakeUpdated")
	if err != nil {
		return nil, err
	}
	return &MarketFactoryMinCreatorStakeUpdatedIterator{contract: _MarketFactory.contract, event: "MinCreatorStakeUpdated", logs: logs, sub: sub}, nil
}

// WatchMinCreatorStakeUpdated is a free log subscription operation binding the contract event 0xb2193ba7a75ea9a3671a8dda768d04598e7c2fed659cf26f16d2c7362e510d99.
//
// Solidity: event MinCreatorStakeUpdated(uint256 oldStake, uint256 newStake)
func (_MarketFactory *MarketFactoryFilterer) WatchMinCreatorStakeUpdated(opts *bind.WatchOpts, sink chan<- *MarketFactoryMinCreatorStakeUpdated) (event.Subscription, error) {

	logs, sub, err := _MarketFactory.contract.WatchLogs(opts, "MinCreatorStakeUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MarketFactoryMinCreatorStakeUpdated)
				if err := _MarketFactory.contract.UnpackLog(event, "MinCreatorStakeUpdated", log); err != nil {
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

// ParseMinCreatorStakeUpdated is a log parse operation binding the contract event 0xb2193ba7a75ea9a3671a8dda768d04598e7c2fed659cf26f16d2c7362e510d99.
//
// Solidity: event MinCreatorStakeUpdated(uint256 oldStake, uint256 newStake)
func (_MarketFactory *MarketFactoryFilterer) ParseMinCreatorStakeUpdated(log types.Log) (*MarketFactoryMinCreatorStakeUpdated, error) {
	event := new(MarketFactoryMinCreatorStakeUpdated)
	if err := _MarketFactory.contract.UnpackLog(event, "MinCreatorStakeUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MarketFactoryOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the MarketFactory contract.
type MarketFactoryOwnershipTransferredIterator struct {
	Event *MarketFactoryOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *MarketFactoryOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MarketFactoryOwnershipTransferred)
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
		it.Event = new(MarketFactoryOwnershipTransferred)
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
func (it *MarketFactoryOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MarketFactoryOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MarketFactoryOwnershipTransferred represents a OwnershipTransferred event raised by the MarketFactory contract.
type MarketFactoryOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_MarketFactory *MarketFactoryFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*MarketFactoryOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _MarketFactory.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &MarketFactoryOwnershipTransferredIterator{contract: _MarketFactory.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_MarketFactory *MarketFactoryFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *MarketFactoryOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _MarketFactory.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MarketFactoryOwnershipTransferred)
				if err := _MarketFactory.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_MarketFactory *MarketFactoryFilterer) ParseOwnershipTransferred(log types.Log) (*MarketFactoryOwnershipTransferred, error) {
	event := new(MarketFactoryOwnershipTransferred)
	if err := _MarketFactory.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
