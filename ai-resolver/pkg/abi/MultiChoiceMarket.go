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

// IMarketMarketInfo is an auto generated low-level Go binding around an user-defined struct.
type IMarketMarketInfo struct {
	MarketId        *big.Int
	MarketType      uint8
	CollateralToken common.Address
	CloseTime       *big.Int
	OutcomeCount    *big.Int
	IsResolved      bool
	IsPaused        bool
}

// MultiChoiceMarketMetaData contains all meta data concerning the MultiChoiceMarket contract.
var MultiChoiceMarketMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_marketId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_collateralToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_outcomeToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_feeSplitter\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_horizonPerks\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_closeTime\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_outcomeCount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_liquidityParameter\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"MINIMUM_LIQUIDITY\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"PRICE_PRECISION\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"addLiquidity\",\"inputs\":[{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"lpTokens\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"admin\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"allowance\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"spender\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"approve\",\"inputs\":[{\"name\":\"spender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"balanceOf\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"buy\",\"inputs\":[{\"name\":\"outcomeId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"collateralIn\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"minTokensOut\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"tokensOut\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"closeTime\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"collateralToken\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIERC20\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"decimals\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"feeSplitter\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractFeeSplitter\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"fundRedemptions\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getAllPrices\",\"inputs\":[],\"outputs\":[{\"name\":\"prices\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getMarketInfo\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structIMarket.MarketInfo\",\"components\":[{\"name\":\"marketId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"marketType\",\"type\":\"uint8\",\"internalType\":\"enumIMarket.MarketType\"},{\"name\":\"collateralToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"closeTime\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"outcomeCount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"isResolved\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"isPaused\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getMarketType\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"enumIMarket.MarketType\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getOutcomeCount\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getOutcomeReserves\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getPrice\",\"inputs\":[{\"name\":\"outcomeId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"price\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getQuoteBuy\",\"inputs\":[{\"name\":\"outcomeId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"collateralIn\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"user\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"tokensOut\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"fee\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getQuoteSell\",\"inputs\":[{\"name\":\"outcomeId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"tokensIn\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"user\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"collateralOut\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"fee\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"horizonPerks\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractHorizonPerks\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"liquidityParameter\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"marketId\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"marketType\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"enumIMarket.MarketType\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"name\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"onERC1155BatchReceived\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"},{\"name\":\"\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"},{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"onERC1155Received\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"outcomeCount\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"outcomeReserves\",\"inputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"outcomeToken\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractOutcomeToken\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pause\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"paused\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"removeLiquidity\",\"inputs\":[{\"name\":\"lpTokens\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"collateralOut\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"sell\",\"inputs\":[{\"name\":\"outcomeId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"tokensIn\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"minCollateralOut\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"collateralOut\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setAdmin\",\"inputs\":[{\"name\":\"newAdmin\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"supportsInterface\",\"inputs\":[{\"name\":\"interfaceId\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"symbol\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"totalCollateral\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"totalSupply\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"transfer\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferFrom\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"unpause\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"Approval\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"spender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"LiquidityChanged\",\"inputs\":[{\"name\":\"provider\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"isAddition\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"LiquidityParameterUpdated\",\"inputs\":[{\"name\":\"oldB\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"newB\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OutcomeReservesUpdated\",\"inputs\":[{\"name\":\"reserves\",\"type\":\"uint256[]\",\"indexed\":false,\"internalType\":\"uint256[]\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Paused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Trade\",\"inputs\":[{\"name\":\"trader\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"outcomeId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"amountIn\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"amountOut\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"fee\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"isBuy\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Transfer\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Unpaused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"ERC20InsufficientAllowance\",\"inputs\":[{\"name\":\"spender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"allowance\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"needed\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"ERC20InsufficientBalance\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"balance\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"needed\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"ERC20InvalidApprover\",\"inputs\":[{\"name\":\"approver\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ERC20InvalidReceiver\",\"inputs\":[{\"name\":\"receiver\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ERC20InvalidSender\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ERC20InvalidSpender\",\"inputs\":[{\"name\":\"spender\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"EnforcedPause\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ExpectedPause\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InsufficientLPTokens\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InsufficientLiquidity\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InsufficientReserves\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidAmount\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidLiquidityParameter\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidOutcomeCount\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidOutcomeId\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidState\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"MarketClosed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"MarketResolved\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"MinimumLiquidityRequired\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"PRBMath_MulDiv18_Overflow\",\"inputs\":[{\"name\":\"x\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"PRBMath_MulDiv_Overflow\",\"inputs\":[{\"name\":\"x\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"y\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"denominator\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"PRBMath_UD60x18_Exp2_InputTooBig\",\"inputs\":[{\"name\":\"x\",\"type\":\"uint256\",\"internalType\":\"UD60x18\"}]},{\"type\":\"error\",\"name\":\"PRBMath_UD60x18_Exp_InputTooBig\",\"inputs\":[{\"name\":\"x\",\"type\":\"uint256\",\"internalType\":\"UD60x18\"}]},{\"type\":\"error\",\"name\":\"PRBMath_UD60x18_Log_InputTooSmall\",\"inputs\":[{\"name\":\"x\",\"type\":\"uint256\",\"internalType\":\"UD60x18\"}]},{\"type\":\"error\",\"name\":\"PriceCalculationOverflow\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ReentrancyGuardReentrantCall\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"SafeERC20FailedOperation\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"SlippageExceeded\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"Unauthorized\",\"inputs\":[]}]",
}

// MultiChoiceMarketABI is the input ABI used to generate the binding from.
// Deprecated: Use MultiChoiceMarketMetaData.ABI instead.
var MultiChoiceMarketABI = MultiChoiceMarketMetaData.ABI

// MultiChoiceMarket is an auto generated Go binding around an Ethereum contract.
type MultiChoiceMarket struct {
	MultiChoiceMarketCaller     // Read-only binding to the contract
	MultiChoiceMarketTransactor // Write-only binding to the contract
	MultiChoiceMarketFilterer   // Log filterer for contract events
}

// MultiChoiceMarketCaller is an auto generated read-only Go binding around an Ethereum contract.
type MultiChoiceMarketCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MultiChoiceMarketTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MultiChoiceMarketTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MultiChoiceMarketFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MultiChoiceMarketFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MultiChoiceMarketSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MultiChoiceMarketSession struct {
	Contract     *MultiChoiceMarket // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// MultiChoiceMarketCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MultiChoiceMarketCallerSession struct {
	Contract *MultiChoiceMarketCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// MultiChoiceMarketTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MultiChoiceMarketTransactorSession struct {
	Contract     *MultiChoiceMarketTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// MultiChoiceMarketRaw is an auto generated low-level Go binding around an Ethereum contract.
type MultiChoiceMarketRaw struct {
	Contract *MultiChoiceMarket // Generic contract binding to access the raw methods on
}

// MultiChoiceMarketCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MultiChoiceMarketCallerRaw struct {
	Contract *MultiChoiceMarketCaller // Generic read-only contract binding to access the raw methods on
}

// MultiChoiceMarketTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MultiChoiceMarketTransactorRaw struct {
	Contract *MultiChoiceMarketTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMultiChoiceMarket creates a new instance of MultiChoiceMarket, bound to a specific deployed contract.
func NewMultiChoiceMarket(address common.Address, backend bind.ContractBackend) (*MultiChoiceMarket, error) {
	contract, err := bindMultiChoiceMarket(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MultiChoiceMarket{MultiChoiceMarketCaller: MultiChoiceMarketCaller{contract: contract}, MultiChoiceMarketTransactor: MultiChoiceMarketTransactor{contract: contract}, MultiChoiceMarketFilterer: MultiChoiceMarketFilterer{contract: contract}}, nil
}

// NewMultiChoiceMarketCaller creates a new read-only instance of MultiChoiceMarket, bound to a specific deployed contract.
func NewMultiChoiceMarketCaller(address common.Address, caller bind.ContractCaller) (*MultiChoiceMarketCaller, error) {
	contract, err := bindMultiChoiceMarket(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MultiChoiceMarketCaller{contract: contract}, nil
}

// NewMultiChoiceMarketTransactor creates a new write-only instance of MultiChoiceMarket, bound to a specific deployed contract.
func NewMultiChoiceMarketTransactor(address common.Address, transactor bind.ContractTransactor) (*MultiChoiceMarketTransactor, error) {
	contract, err := bindMultiChoiceMarket(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MultiChoiceMarketTransactor{contract: contract}, nil
}

// NewMultiChoiceMarketFilterer creates a new log filterer instance of MultiChoiceMarket, bound to a specific deployed contract.
func NewMultiChoiceMarketFilterer(address common.Address, filterer bind.ContractFilterer) (*MultiChoiceMarketFilterer, error) {
	contract, err := bindMultiChoiceMarket(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MultiChoiceMarketFilterer{contract: contract}, nil
}

// bindMultiChoiceMarket binds a generic wrapper to an already deployed contract.
func bindMultiChoiceMarket(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := MultiChoiceMarketMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MultiChoiceMarket *MultiChoiceMarketRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MultiChoiceMarket.Contract.MultiChoiceMarketCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MultiChoiceMarket *MultiChoiceMarketRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MultiChoiceMarket.Contract.MultiChoiceMarketTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MultiChoiceMarket *MultiChoiceMarketRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MultiChoiceMarket.Contract.MultiChoiceMarketTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MultiChoiceMarket *MultiChoiceMarketCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MultiChoiceMarket.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MultiChoiceMarket *MultiChoiceMarketTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MultiChoiceMarket.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MultiChoiceMarket *MultiChoiceMarketTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MultiChoiceMarket.Contract.contract.Transact(opts, method, params...)
}

// MINIMUMLIQUIDITY is a free data retrieval call binding the contract method 0xba9a7a56.
//
// Solidity: function MINIMUM_LIQUIDITY() view returns(uint256)
func (_MultiChoiceMarket *MultiChoiceMarketCaller) MINIMUMLIQUIDITY(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MultiChoiceMarket.contract.Call(opts, &out, "MINIMUM_LIQUIDITY")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MINIMUMLIQUIDITY is a free data retrieval call binding the contract method 0xba9a7a56.
//
// Solidity: function MINIMUM_LIQUIDITY() view returns(uint256)
func (_MultiChoiceMarket *MultiChoiceMarketSession) MINIMUMLIQUIDITY() (*big.Int, error) {
	return _MultiChoiceMarket.Contract.MINIMUMLIQUIDITY(&_MultiChoiceMarket.CallOpts)
}

// MINIMUMLIQUIDITY is a free data retrieval call binding the contract method 0xba9a7a56.
//
// Solidity: function MINIMUM_LIQUIDITY() view returns(uint256)
func (_MultiChoiceMarket *MultiChoiceMarketCallerSession) MINIMUMLIQUIDITY() (*big.Int, error) {
	return _MultiChoiceMarket.Contract.MINIMUMLIQUIDITY(&_MultiChoiceMarket.CallOpts)
}

// PRICEPRECISION is a free data retrieval call binding the contract method 0x95082d25.
//
// Solidity: function PRICE_PRECISION() view returns(uint256)
func (_MultiChoiceMarket *MultiChoiceMarketCaller) PRICEPRECISION(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MultiChoiceMarket.contract.Call(opts, &out, "PRICE_PRECISION")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PRICEPRECISION is a free data retrieval call binding the contract method 0x95082d25.
//
// Solidity: function PRICE_PRECISION() view returns(uint256)
func (_MultiChoiceMarket *MultiChoiceMarketSession) PRICEPRECISION() (*big.Int, error) {
	return _MultiChoiceMarket.Contract.PRICEPRECISION(&_MultiChoiceMarket.CallOpts)
}

// PRICEPRECISION is a free data retrieval call binding the contract method 0x95082d25.
//
// Solidity: function PRICE_PRECISION() view returns(uint256)
func (_MultiChoiceMarket *MultiChoiceMarketCallerSession) PRICEPRECISION() (*big.Int, error) {
	return _MultiChoiceMarket.Contract.PRICEPRECISION(&_MultiChoiceMarket.CallOpts)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_MultiChoiceMarket *MultiChoiceMarketCaller) Admin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MultiChoiceMarket.contract.Call(opts, &out, "admin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_MultiChoiceMarket *MultiChoiceMarketSession) Admin() (common.Address, error) {
	return _MultiChoiceMarket.Contract.Admin(&_MultiChoiceMarket.CallOpts)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_MultiChoiceMarket *MultiChoiceMarketCallerSession) Admin() (common.Address, error) {
	return _MultiChoiceMarket.Contract.Admin(&_MultiChoiceMarket.CallOpts)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_MultiChoiceMarket *MultiChoiceMarketCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _MultiChoiceMarket.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_MultiChoiceMarket *MultiChoiceMarketSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _MultiChoiceMarket.Contract.Allowance(&_MultiChoiceMarket.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_MultiChoiceMarket *MultiChoiceMarketCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _MultiChoiceMarket.Contract.Allowance(&_MultiChoiceMarket.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_MultiChoiceMarket *MultiChoiceMarketCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _MultiChoiceMarket.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_MultiChoiceMarket *MultiChoiceMarketSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _MultiChoiceMarket.Contract.BalanceOf(&_MultiChoiceMarket.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_MultiChoiceMarket *MultiChoiceMarketCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _MultiChoiceMarket.Contract.BalanceOf(&_MultiChoiceMarket.CallOpts, account)
}

// CloseTime is a free data retrieval call binding the contract method 0x627749e6.
//
// Solidity: function closeTime() view returns(uint256)
func (_MultiChoiceMarket *MultiChoiceMarketCaller) CloseTime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MultiChoiceMarket.contract.Call(opts, &out, "closeTime")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CloseTime is a free data retrieval call binding the contract method 0x627749e6.
//
// Solidity: function closeTime() view returns(uint256)
func (_MultiChoiceMarket *MultiChoiceMarketSession) CloseTime() (*big.Int, error) {
	return _MultiChoiceMarket.Contract.CloseTime(&_MultiChoiceMarket.CallOpts)
}

// CloseTime is a free data retrieval call binding the contract method 0x627749e6.
//
// Solidity: function closeTime() view returns(uint256)
func (_MultiChoiceMarket *MultiChoiceMarketCallerSession) CloseTime() (*big.Int, error) {
	return _MultiChoiceMarket.Contract.CloseTime(&_MultiChoiceMarket.CallOpts)
}

// CollateralToken is a free data retrieval call binding the contract method 0xb2016bd4.
//
// Solidity: function collateralToken() view returns(address)
func (_MultiChoiceMarket *MultiChoiceMarketCaller) CollateralToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MultiChoiceMarket.contract.Call(opts, &out, "collateralToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CollateralToken is a free data retrieval call binding the contract method 0xb2016bd4.
//
// Solidity: function collateralToken() view returns(address)
func (_MultiChoiceMarket *MultiChoiceMarketSession) CollateralToken() (common.Address, error) {
	return _MultiChoiceMarket.Contract.CollateralToken(&_MultiChoiceMarket.CallOpts)
}

// CollateralToken is a free data retrieval call binding the contract method 0xb2016bd4.
//
// Solidity: function collateralToken() view returns(address)
func (_MultiChoiceMarket *MultiChoiceMarketCallerSession) CollateralToken() (common.Address, error) {
	return _MultiChoiceMarket.Contract.CollateralToken(&_MultiChoiceMarket.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_MultiChoiceMarket *MultiChoiceMarketCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _MultiChoiceMarket.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_MultiChoiceMarket *MultiChoiceMarketSession) Decimals() (uint8, error) {
	return _MultiChoiceMarket.Contract.Decimals(&_MultiChoiceMarket.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_MultiChoiceMarket *MultiChoiceMarketCallerSession) Decimals() (uint8, error) {
	return _MultiChoiceMarket.Contract.Decimals(&_MultiChoiceMarket.CallOpts)
}

// FeeSplitter is a free data retrieval call binding the contract method 0x6052970c.
//
// Solidity: function feeSplitter() view returns(address)
func (_MultiChoiceMarket *MultiChoiceMarketCaller) FeeSplitter(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MultiChoiceMarket.contract.Call(opts, &out, "feeSplitter")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FeeSplitter is a free data retrieval call binding the contract method 0x6052970c.
//
// Solidity: function feeSplitter() view returns(address)
func (_MultiChoiceMarket *MultiChoiceMarketSession) FeeSplitter() (common.Address, error) {
	return _MultiChoiceMarket.Contract.FeeSplitter(&_MultiChoiceMarket.CallOpts)
}

// FeeSplitter is a free data retrieval call binding the contract method 0x6052970c.
//
// Solidity: function feeSplitter() view returns(address)
func (_MultiChoiceMarket *MultiChoiceMarketCallerSession) FeeSplitter() (common.Address, error) {
	return _MultiChoiceMarket.Contract.FeeSplitter(&_MultiChoiceMarket.CallOpts)
}

// GetAllPrices is a free data retrieval call binding the contract method 0x445df9d6.
//
// Solidity: function getAllPrices() view returns(uint256[] prices)
func (_MultiChoiceMarket *MultiChoiceMarketCaller) GetAllPrices(opts *bind.CallOpts) ([]*big.Int, error) {
	var out []interface{}
	err := _MultiChoiceMarket.contract.Call(opts, &out, "getAllPrices")

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetAllPrices is a free data retrieval call binding the contract method 0x445df9d6.
//
// Solidity: function getAllPrices() view returns(uint256[] prices)
func (_MultiChoiceMarket *MultiChoiceMarketSession) GetAllPrices() ([]*big.Int, error) {
	return _MultiChoiceMarket.Contract.GetAllPrices(&_MultiChoiceMarket.CallOpts)
}

// GetAllPrices is a free data retrieval call binding the contract method 0x445df9d6.
//
// Solidity: function getAllPrices() view returns(uint256[] prices)
func (_MultiChoiceMarket *MultiChoiceMarketCallerSession) GetAllPrices() ([]*big.Int, error) {
	return _MultiChoiceMarket.Contract.GetAllPrices(&_MultiChoiceMarket.CallOpts)
}

// GetMarketInfo is a free data retrieval call binding the contract method 0x23341a05.
//
// Solidity: function getMarketInfo() view returns((uint256,uint8,address,uint256,uint256,bool,bool))
func (_MultiChoiceMarket *MultiChoiceMarketCaller) GetMarketInfo(opts *bind.CallOpts) (IMarketMarketInfo, error) {
	var out []interface{}
	err := _MultiChoiceMarket.contract.Call(opts, &out, "getMarketInfo")

	if err != nil {
		return *new(IMarketMarketInfo), err
	}

	out0 := *abi.ConvertType(out[0], new(IMarketMarketInfo)).(*IMarketMarketInfo)

	return out0, err

}

// GetMarketInfo is a free data retrieval call binding the contract method 0x23341a05.
//
// Solidity: function getMarketInfo() view returns((uint256,uint8,address,uint256,uint256,bool,bool))
func (_MultiChoiceMarket *MultiChoiceMarketSession) GetMarketInfo() (IMarketMarketInfo, error) {
	return _MultiChoiceMarket.Contract.GetMarketInfo(&_MultiChoiceMarket.CallOpts)
}

// GetMarketInfo is a free data retrieval call binding the contract method 0x23341a05.
//
// Solidity: function getMarketInfo() view returns((uint256,uint8,address,uint256,uint256,bool,bool))
func (_MultiChoiceMarket *MultiChoiceMarketCallerSession) GetMarketInfo() (IMarketMarketInfo, error) {
	return _MultiChoiceMarket.Contract.GetMarketInfo(&_MultiChoiceMarket.CallOpts)
}

// GetMarketType is a free data retrieval call binding the contract method 0x33e7a1d0.
//
// Solidity: function getMarketType() view returns(uint8)
func (_MultiChoiceMarket *MultiChoiceMarketCaller) GetMarketType(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _MultiChoiceMarket.contract.Call(opts, &out, "getMarketType")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// GetMarketType is a free data retrieval call binding the contract method 0x33e7a1d0.
//
// Solidity: function getMarketType() view returns(uint8)
func (_MultiChoiceMarket *MultiChoiceMarketSession) GetMarketType() (uint8, error) {
	return _MultiChoiceMarket.Contract.GetMarketType(&_MultiChoiceMarket.CallOpts)
}

// GetMarketType is a free data retrieval call binding the contract method 0x33e7a1d0.
//
// Solidity: function getMarketType() view returns(uint8)
func (_MultiChoiceMarket *MultiChoiceMarketCallerSession) GetMarketType() (uint8, error) {
	return _MultiChoiceMarket.Contract.GetMarketType(&_MultiChoiceMarket.CallOpts)
}

// GetOutcomeCount is a free data retrieval call binding the contract method 0x7dc8f086.
//
// Solidity: function getOutcomeCount() view returns(uint256)
func (_MultiChoiceMarket *MultiChoiceMarketCaller) GetOutcomeCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MultiChoiceMarket.contract.Call(opts, &out, "getOutcomeCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetOutcomeCount is a free data retrieval call binding the contract method 0x7dc8f086.
//
// Solidity: function getOutcomeCount() view returns(uint256)
func (_MultiChoiceMarket *MultiChoiceMarketSession) GetOutcomeCount() (*big.Int, error) {
	return _MultiChoiceMarket.Contract.GetOutcomeCount(&_MultiChoiceMarket.CallOpts)
}

// GetOutcomeCount is a free data retrieval call binding the contract method 0x7dc8f086.
//
// Solidity: function getOutcomeCount() view returns(uint256)
func (_MultiChoiceMarket *MultiChoiceMarketCallerSession) GetOutcomeCount() (*big.Int, error) {
	return _MultiChoiceMarket.Contract.GetOutcomeCount(&_MultiChoiceMarket.CallOpts)
}

// GetOutcomeReserves is a free data retrieval call binding the contract method 0x5ed76ef8.
//
// Solidity: function getOutcomeReserves() view returns(uint256[])
func (_MultiChoiceMarket *MultiChoiceMarketCaller) GetOutcomeReserves(opts *bind.CallOpts) ([]*big.Int, error) {
	var out []interface{}
	err := _MultiChoiceMarket.contract.Call(opts, &out, "getOutcomeReserves")

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetOutcomeReserves is a free data retrieval call binding the contract method 0x5ed76ef8.
//
// Solidity: function getOutcomeReserves() view returns(uint256[])
func (_MultiChoiceMarket *MultiChoiceMarketSession) GetOutcomeReserves() ([]*big.Int, error) {
	return _MultiChoiceMarket.Contract.GetOutcomeReserves(&_MultiChoiceMarket.CallOpts)
}

// GetOutcomeReserves is a free data retrieval call binding the contract method 0x5ed76ef8.
//
// Solidity: function getOutcomeReserves() view returns(uint256[])
func (_MultiChoiceMarket *MultiChoiceMarketCallerSession) GetOutcomeReserves() ([]*big.Int, error) {
	return _MultiChoiceMarket.Contract.GetOutcomeReserves(&_MultiChoiceMarket.CallOpts)
}

// GetPrice is a free data retrieval call binding the contract method 0xe7572230.
//
// Solidity: function getPrice(uint256 outcomeId) view returns(uint256 price)
func (_MultiChoiceMarket *MultiChoiceMarketCaller) GetPrice(opts *bind.CallOpts, outcomeId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _MultiChoiceMarket.contract.Call(opts, &out, "getPrice", outcomeId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetPrice is a free data retrieval call binding the contract method 0xe7572230.
//
// Solidity: function getPrice(uint256 outcomeId) view returns(uint256 price)
func (_MultiChoiceMarket *MultiChoiceMarketSession) GetPrice(outcomeId *big.Int) (*big.Int, error) {
	return _MultiChoiceMarket.Contract.GetPrice(&_MultiChoiceMarket.CallOpts, outcomeId)
}

// GetPrice is a free data retrieval call binding the contract method 0xe7572230.
//
// Solidity: function getPrice(uint256 outcomeId) view returns(uint256 price)
func (_MultiChoiceMarket *MultiChoiceMarketCallerSession) GetPrice(outcomeId *big.Int) (*big.Int, error) {
	return _MultiChoiceMarket.Contract.GetPrice(&_MultiChoiceMarket.CallOpts, outcomeId)
}

// GetQuoteBuy is a free data retrieval call binding the contract method 0xca6d5811.
//
// Solidity: function getQuoteBuy(uint256 outcomeId, uint256 collateralIn, address user) view returns(uint256 tokensOut, uint256 fee)
func (_MultiChoiceMarket *MultiChoiceMarketCaller) GetQuoteBuy(opts *bind.CallOpts, outcomeId *big.Int, collateralIn *big.Int, user common.Address) (struct {
	TokensOut *big.Int
	Fee       *big.Int
}, error) {
	var out []interface{}
	err := _MultiChoiceMarket.contract.Call(opts, &out, "getQuoteBuy", outcomeId, collateralIn, user)

	outstruct := new(struct {
		TokensOut *big.Int
		Fee       *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.TokensOut = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Fee = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetQuoteBuy is a free data retrieval call binding the contract method 0xca6d5811.
//
// Solidity: function getQuoteBuy(uint256 outcomeId, uint256 collateralIn, address user) view returns(uint256 tokensOut, uint256 fee)
func (_MultiChoiceMarket *MultiChoiceMarketSession) GetQuoteBuy(outcomeId *big.Int, collateralIn *big.Int, user common.Address) (struct {
	TokensOut *big.Int
	Fee       *big.Int
}, error) {
	return _MultiChoiceMarket.Contract.GetQuoteBuy(&_MultiChoiceMarket.CallOpts, outcomeId, collateralIn, user)
}

// GetQuoteBuy is a free data retrieval call binding the contract method 0xca6d5811.
//
// Solidity: function getQuoteBuy(uint256 outcomeId, uint256 collateralIn, address user) view returns(uint256 tokensOut, uint256 fee)
func (_MultiChoiceMarket *MultiChoiceMarketCallerSession) GetQuoteBuy(outcomeId *big.Int, collateralIn *big.Int, user common.Address) (struct {
	TokensOut *big.Int
	Fee       *big.Int
}, error) {
	return _MultiChoiceMarket.Contract.GetQuoteBuy(&_MultiChoiceMarket.CallOpts, outcomeId, collateralIn, user)
}

// GetQuoteSell is a free data retrieval call binding the contract method 0x8b5e8a24.
//
// Solidity: function getQuoteSell(uint256 outcomeId, uint256 tokensIn, address user) view returns(uint256 collateralOut, uint256 fee)
func (_MultiChoiceMarket *MultiChoiceMarketCaller) GetQuoteSell(opts *bind.CallOpts, outcomeId *big.Int, tokensIn *big.Int, user common.Address) (struct {
	CollateralOut *big.Int
	Fee           *big.Int
}, error) {
	var out []interface{}
	err := _MultiChoiceMarket.contract.Call(opts, &out, "getQuoteSell", outcomeId, tokensIn, user)

	outstruct := new(struct {
		CollateralOut *big.Int
		Fee           *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.CollateralOut = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Fee = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetQuoteSell is a free data retrieval call binding the contract method 0x8b5e8a24.
//
// Solidity: function getQuoteSell(uint256 outcomeId, uint256 tokensIn, address user) view returns(uint256 collateralOut, uint256 fee)
func (_MultiChoiceMarket *MultiChoiceMarketSession) GetQuoteSell(outcomeId *big.Int, tokensIn *big.Int, user common.Address) (struct {
	CollateralOut *big.Int
	Fee           *big.Int
}, error) {
	return _MultiChoiceMarket.Contract.GetQuoteSell(&_MultiChoiceMarket.CallOpts, outcomeId, tokensIn, user)
}

// GetQuoteSell is a free data retrieval call binding the contract method 0x8b5e8a24.
//
// Solidity: function getQuoteSell(uint256 outcomeId, uint256 tokensIn, address user) view returns(uint256 collateralOut, uint256 fee)
func (_MultiChoiceMarket *MultiChoiceMarketCallerSession) GetQuoteSell(outcomeId *big.Int, tokensIn *big.Int, user common.Address) (struct {
	CollateralOut *big.Int
	Fee           *big.Int
}, error) {
	return _MultiChoiceMarket.Contract.GetQuoteSell(&_MultiChoiceMarket.CallOpts, outcomeId, tokensIn, user)
}

// HorizonPerks is a free data retrieval call binding the contract method 0xffe02e34.
//
// Solidity: function horizonPerks() view returns(address)
func (_MultiChoiceMarket *MultiChoiceMarketCaller) HorizonPerks(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MultiChoiceMarket.contract.Call(opts, &out, "horizonPerks")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// HorizonPerks is a free data retrieval call binding the contract method 0xffe02e34.
//
// Solidity: function horizonPerks() view returns(address)
func (_MultiChoiceMarket *MultiChoiceMarketSession) HorizonPerks() (common.Address, error) {
	return _MultiChoiceMarket.Contract.HorizonPerks(&_MultiChoiceMarket.CallOpts)
}

// HorizonPerks is a free data retrieval call binding the contract method 0xffe02e34.
//
// Solidity: function horizonPerks() view returns(address)
func (_MultiChoiceMarket *MultiChoiceMarketCallerSession) HorizonPerks() (common.Address, error) {
	return _MultiChoiceMarket.Contract.HorizonPerks(&_MultiChoiceMarket.CallOpts)
}

// LiquidityParameter is a free data retrieval call binding the contract method 0x3a69a1be.
//
// Solidity: function liquidityParameter() view returns(uint256)
func (_MultiChoiceMarket *MultiChoiceMarketCaller) LiquidityParameter(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MultiChoiceMarket.contract.Call(opts, &out, "liquidityParameter")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LiquidityParameter is a free data retrieval call binding the contract method 0x3a69a1be.
//
// Solidity: function liquidityParameter() view returns(uint256)
func (_MultiChoiceMarket *MultiChoiceMarketSession) LiquidityParameter() (*big.Int, error) {
	return _MultiChoiceMarket.Contract.LiquidityParameter(&_MultiChoiceMarket.CallOpts)
}

// LiquidityParameter is a free data retrieval call binding the contract method 0x3a69a1be.
//
// Solidity: function liquidityParameter() view returns(uint256)
func (_MultiChoiceMarket *MultiChoiceMarketCallerSession) LiquidityParameter() (*big.Int, error) {
	return _MultiChoiceMarket.Contract.LiquidityParameter(&_MultiChoiceMarket.CallOpts)
}

// MarketId is a free data retrieval call binding the contract method 0x6ed71ede.
//
// Solidity: function marketId() view returns(uint256)
func (_MultiChoiceMarket *MultiChoiceMarketCaller) MarketId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MultiChoiceMarket.contract.Call(opts, &out, "marketId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MarketId is a free data retrieval call binding the contract method 0x6ed71ede.
//
// Solidity: function marketId() view returns(uint256)
func (_MultiChoiceMarket *MultiChoiceMarketSession) MarketId() (*big.Int, error) {
	return _MultiChoiceMarket.Contract.MarketId(&_MultiChoiceMarket.CallOpts)
}

// MarketId is a free data retrieval call binding the contract method 0x6ed71ede.
//
// Solidity: function marketId() view returns(uint256)
func (_MultiChoiceMarket *MultiChoiceMarketCallerSession) MarketId() (*big.Int, error) {
	return _MultiChoiceMarket.Contract.MarketId(&_MultiChoiceMarket.CallOpts)
}

// MarketType is a free data retrieval call binding the contract method 0x2dd48909.
//
// Solidity: function marketType() view returns(uint8)
func (_MultiChoiceMarket *MultiChoiceMarketCaller) MarketType(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _MultiChoiceMarket.contract.Call(opts, &out, "marketType")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// MarketType is a free data retrieval call binding the contract method 0x2dd48909.
//
// Solidity: function marketType() view returns(uint8)
func (_MultiChoiceMarket *MultiChoiceMarketSession) MarketType() (uint8, error) {
	return _MultiChoiceMarket.Contract.MarketType(&_MultiChoiceMarket.CallOpts)
}

// MarketType is a free data retrieval call binding the contract method 0x2dd48909.
//
// Solidity: function marketType() view returns(uint8)
func (_MultiChoiceMarket *MultiChoiceMarketCallerSession) MarketType() (uint8, error) {
	return _MultiChoiceMarket.Contract.MarketType(&_MultiChoiceMarket.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_MultiChoiceMarket *MultiChoiceMarketCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _MultiChoiceMarket.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_MultiChoiceMarket *MultiChoiceMarketSession) Name() (string, error) {
	return _MultiChoiceMarket.Contract.Name(&_MultiChoiceMarket.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_MultiChoiceMarket *MultiChoiceMarketCallerSession) Name() (string, error) {
	return _MultiChoiceMarket.Contract.Name(&_MultiChoiceMarket.CallOpts)
}

// OnERC1155BatchReceived is a free data retrieval call binding the contract method 0xbc197c81.
//
// Solidity: function onERC1155BatchReceived(address , address , uint256[] , uint256[] , bytes ) pure returns(bytes4)
func (_MultiChoiceMarket *MultiChoiceMarketCaller) OnERC1155BatchReceived(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address, arg2 []*big.Int, arg3 []*big.Int, arg4 []byte) ([4]byte, error) {
	var out []interface{}
	err := _MultiChoiceMarket.contract.Call(opts, &out, "onERC1155BatchReceived", arg0, arg1, arg2, arg3, arg4)

	if err != nil {
		return *new([4]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([4]byte)).(*[4]byte)

	return out0, err

}

// OnERC1155BatchReceived is a free data retrieval call binding the contract method 0xbc197c81.
//
// Solidity: function onERC1155BatchReceived(address , address , uint256[] , uint256[] , bytes ) pure returns(bytes4)
func (_MultiChoiceMarket *MultiChoiceMarketSession) OnERC1155BatchReceived(arg0 common.Address, arg1 common.Address, arg2 []*big.Int, arg3 []*big.Int, arg4 []byte) ([4]byte, error) {
	return _MultiChoiceMarket.Contract.OnERC1155BatchReceived(&_MultiChoiceMarket.CallOpts, arg0, arg1, arg2, arg3, arg4)
}

// OnERC1155BatchReceived is a free data retrieval call binding the contract method 0xbc197c81.
//
// Solidity: function onERC1155BatchReceived(address , address , uint256[] , uint256[] , bytes ) pure returns(bytes4)
func (_MultiChoiceMarket *MultiChoiceMarketCallerSession) OnERC1155BatchReceived(arg0 common.Address, arg1 common.Address, arg2 []*big.Int, arg3 []*big.Int, arg4 []byte) ([4]byte, error) {
	return _MultiChoiceMarket.Contract.OnERC1155BatchReceived(&_MultiChoiceMarket.CallOpts, arg0, arg1, arg2, arg3, arg4)
}

// OnERC1155Received is a free data retrieval call binding the contract method 0xf23a6e61.
//
// Solidity: function onERC1155Received(address , address , uint256 , uint256 , bytes ) pure returns(bytes4)
func (_MultiChoiceMarket *MultiChoiceMarketCaller) OnERC1155Received(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 *big.Int, arg4 []byte) ([4]byte, error) {
	var out []interface{}
	err := _MultiChoiceMarket.contract.Call(opts, &out, "onERC1155Received", arg0, arg1, arg2, arg3, arg4)

	if err != nil {
		return *new([4]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([4]byte)).(*[4]byte)

	return out0, err

}

// OnERC1155Received is a free data retrieval call binding the contract method 0xf23a6e61.
//
// Solidity: function onERC1155Received(address , address , uint256 , uint256 , bytes ) pure returns(bytes4)
func (_MultiChoiceMarket *MultiChoiceMarketSession) OnERC1155Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 *big.Int, arg4 []byte) ([4]byte, error) {
	return _MultiChoiceMarket.Contract.OnERC1155Received(&_MultiChoiceMarket.CallOpts, arg0, arg1, arg2, arg3, arg4)
}

// OnERC1155Received is a free data retrieval call binding the contract method 0xf23a6e61.
//
// Solidity: function onERC1155Received(address , address , uint256 , uint256 , bytes ) pure returns(bytes4)
func (_MultiChoiceMarket *MultiChoiceMarketCallerSession) OnERC1155Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 *big.Int, arg4 []byte) ([4]byte, error) {
	return _MultiChoiceMarket.Contract.OnERC1155Received(&_MultiChoiceMarket.CallOpts, arg0, arg1, arg2, arg3, arg4)
}

// OutcomeCount is a free data retrieval call binding the contract method 0xd300cb31.
//
// Solidity: function outcomeCount() view returns(uint256)
func (_MultiChoiceMarket *MultiChoiceMarketCaller) OutcomeCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MultiChoiceMarket.contract.Call(opts, &out, "outcomeCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// OutcomeCount is a free data retrieval call binding the contract method 0xd300cb31.
//
// Solidity: function outcomeCount() view returns(uint256)
func (_MultiChoiceMarket *MultiChoiceMarketSession) OutcomeCount() (*big.Int, error) {
	return _MultiChoiceMarket.Contract.OutcomeCount(&_MultiChoiceMarket.CallOpts)
}

// OutcomeCount is a free data retrieval call binding the contract method 0xd300cb31.
//
// Solidity: function outcomeCount() view returns(uint256)
func (_MultiChoiceMarket *MultiChoiceMarketCallerSession) OutcomeCount() (*big.Int, error) {
	return _MultiChoiceMarket.Contract.OutcomeCount(&_MultiChoiceMarket.CallOpts)
}

// OutcomeReserves is a free data retrieval call binding the contract method 0x48b61ae6.
//
// Solidity: function outcomeReserves(uint256 ) view returns(uint256)
func (_MultiChoiceMarket *MultiChoiceMarketCaller) OutcomeReserves(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _MultiChoiceMarket.contract.Call(opts, &out, "outcomeReserves", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// OutcomeReserves is a free data retrieval call binding the contract method 0x48b61ae6.
//
// Solidity: function outcomeReserves(uint256 ) view returns(uint256)
func (_MultiChoiceMarket *MultiChoiceMarketSession) OutcomeReserves(arg0 *big.Int) (*big.Int, error) {
	return _MultiChoiceMarket.Contract.OutcomeReserves(&_MultiChoiceMarket.CallOpts, arg0)
}

// OutcomeReserves is a free data retrieval call binding the contract method 0x48b61ae6.
//
// Solidity: function outcomeReserves(uint256 ) view returns(uint256)
func (_MultiChoiceMarket *MultiChoiceMarketCallerSession) OutcomeReserves(arg0 *big.Int) (*big.Int, error) {
	return _MultiChoiceMarket.Contract.OutcomeReserves(&_MultiChoiceMarket.CallOpts, arg0)
}

// OutcomeToken is a free data retrieval call binding the contract method 0xa998d6d8.
//
// Solidity: function outcomeToken() view returns(address)
func (_MultiChoiceMarket *MultiChoiceMarketCaller) OutcomeToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MultiChoiceMarket.contract.Call(opts, &out, "outcomeToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OutcomeToken is a free data retrieval call binding the contract method 0xa998d6d8.
//
// Solidity: function outcomeToken() view returns(address)
func (_MultiChoiceMarket *MultiChoiceMarketSession) OutcomeToken() (common.Address, error) {
	return _MultiChoiceMarket.Contract.OutcomeToken(&_MultiChoiceMarket.CallOpts)
}

// OutcomeToken is a free data retrieval call binding the contract method 0xa998d6d8.
//
// Solidity: function outcomeToken() view returns(address)
func (_MultiChoiceMarket *MultiChoiceMarketCallerSession) OutcomeToken() (common.Address, error) {
	return _MultiChoiceMarket.Contract.OutcomeToken(&_MultiChoiceMarket.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_MultiChoiceMarket *MultiChoiceMarketCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _MultiChoiceMarket.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_MultiChoiceMarket *MultiChoiceMarketSession) Paused() (bool, error) {
	return _MultiChoiceMarket.Contract.Paused(&_MultiChoiceMarket.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_MultiChoiceMarket *MultiChoiceMarketCallerSession) Paused() (bool, error) {
	return _MultiChoiceMarket.Contract.Paused(&_MultiChoiceMarket.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) pure returns(bool)
func (_MultiChoiceMarket *MultiChoiceMarketCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _MultiChoiceMarket.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) pure returns(bool)
func (_MultiChoiceMarket *MultiChoiceMarketSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _MultiChoiceMarket.Contract.SupportsInterface(&_MultiChoiceMarket.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) pure returns(bool)
func (_MultiChoiceMarket *MultiChoiceMarketCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _MultiChoiceMarket.Contract.SupportsInterface(&_MultiChoiceMarket.CallOpts, interfaceId)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_MultiChoiceMarket *MultiChoiceMarketCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _MultiChoiceMarket.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_MultiChoiceMarket *MultiChoiceMarketSession) Symbol() (string, error) {
	return _MultiChoiceMarket.Contract.Symbol(&_MultiChoiceMarket.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_MultiChoiceMarket *MultiChoiceMarketCallerSession) Symbol() (string, error) {
	return _MultiChoiceMarket.Contract.Symbol(&_MultiChoiceMarket.CallOpts)
}

// TotalCollateral is a free data retrieval call binding the contract method 0x4ac8eb5f.
//
// Solidity: function totalCollateral() view returns(uint256)
func (_MultiChoiceMarket *MultiChoiceMarketCaller) TotalCollateral(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MultiChoiceMarket.contract.Call(opts, &out, "totalCollateral")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalCollateral is a free data retrieval call binding the contract method 0x4ac8eb5f.
//
// Solidity: function totalCollateral() view returns(uint256)
func (_MultiChoiceMarket *MultiChoiceMarketSession) TotalCollateral() (*big.Int, error) {
	return _MultiChoiceMarket.Contract.TotalCollateral(&_MultiChoiceMarket.CallOpts)
}

// TotalCollateral is a free data retrieval call binding the contract method 0x4ac8eb5f.
//
// Solidity: function totalCollateral() view returns(uint256)
func (_MultiChoiceMarket *MultiChoiceMarketCallerSession) TotalCollateral() (*big.Int, error) {
	return _MultiChoiceMarket.Contract.TotalCollateral(&_MultiChoiceMarket.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_MultiChoiceMarket *MultiChoiceMarketCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MultiChoiceMarket.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_MultiChoiceMarket *MultiChoiceMarketSession) TotalSupply() (*big.Int, error) {
	return _MultiChoiceMarket.Contract.TotalSupply(&_MultiChoiceMarket.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_MultiChoiceMarket *MultiChoiceMarketCallerSession) TotalSupply() (*big.Int, error) {
	return _MultiChoiceMarket.Contract.TotalSupply(&_MultiChoiceMarket.CallOpts)
}

// AddLiquidity is a paid mutator transaction binding the contract method 0x51c6590a.
//
// Solidity: function addLiquidity(uint256 amount) returns(uint256 lpTokens)
func (_MultiChoiceMarket *MultiChoiceMarketTransactor) AddLiquidity(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _MultiChoiceMarket.contract.Transact(opts, "addLiquidity", amount)
}

// AddLiquidity is a paid mutator transaction binding the contract method 0x51c6590a.
//
// Solidity: function addLiquidity(uint256 amount) returns(uint256 lpTokens)
func (_MultiChoiceMarket *MultiChoiceMarketSession) AddLiquidity(amount *big.Int) (*types.Transaction, error) {
	return _MultiChoiceMarket.Contract.AddLiquidity(&_MultiChoiceMarket.TransactOpts, amount)
}

// AddLiquidity is a paid mutator transaction binding the contract method 0x51c6590a.
//
// Solidity: function addLiquidity(uint256 amount) returns(uint256 lpTokens)
func (_MultiChoiceMarket *MultiChoiceMarketTransactorSession) AddLiquidity(amount *big.Int) (*types.Transaction, error) {
	return _MultiChoiceMarket.Contract.AddLiquidity(&_MultiChoiceMarket.TransactOpts, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_MultiChoiceMarket *MultiChoiceMarketTransactor) Approve(opts *bind.TransactOpts, spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _MultiChoiceMarket.contract.Transact(opts, "approve", spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_MultiChoiceMarket *MultiChoiceMarketSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _MultiChoiceMarket.Contract.Approve(&_MultiChoiceMarket.TransactOpts, spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_MultiChoiceMarket *MultiChoiceMarketTransactorSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _MultiChoiceMarket.Contract.Approve(&_MultiChoiceMarket.TransactOpts, spender, value)
}

// Buy is a paid mutator transaction binding the contract method 0x40993b26.
//
// Solidity: function buy(uint256 outcomeId, uint256 collateralIn, uint256 minTokensOut) returns(uint256 tokensOut)
func (_MultiChoiceMarket *MultiChoiceMarketTransactor) Buy(opts *bind.TransactOpts, outcomeId *big.Int, collateralIn *big.Int, minTokensOut *big.Int) (*types.Transaction, error) {
	return _MultiChoiceMarket.contract.Transact(opts, "buy", outcomeId, collateralIn, minTokensOut)
}

// Buy is a paid mutator transaction binding the contract method 0x40993b26.
//
// Solidity: function buy(uint256 outcomeId, uint256 collateralIn, uint256 minTokensOut) returns(uint256 tokensOut)
func (_MultiChoiceMarket *MultiChoiceMarketSession) Buy(outcomeId *big.Int, collateralIn *big.Int, minTokensOut *big.Int) (*types.Transaction, error) {
	return _MultiChoiceMarket.Contract.Buy(&_MultiChoiceMarket.TransactOpts, outcomeId, collateralIn, minTokensOut)
}

// Buy is a paid mutator transaction binding the contract method 0x40993b26.
//
// Solidity: function buy(uint256 outcomeId, uint256 collateralIn, uint256 minTokensOut) returns(uint256 tokensOut)
func (_MultiChoiceMarket *MultiChoiceMarketTransactorSession) Buy(outcomeId *big.Int, collateralIn *big.Int, minTokensOut *big.Int) (*types.Transaction, error) {
	return _MultiChoiceMarket.Contract.Buy(&_MultiChoiceMarket.TransactOpts, outcomeId, collateralIn, minTokensOut)
}

// FundRedemptions is a paid mutator transaction binding the contract method 0x281155ba.
//
// Solidity: function fundRedemptions() returns()
func (_MultiChoiceMarket *MultiChoiceMarketTransactor) FundRedemptions(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MultiChoiceMarket.contract.Transact(opts, "fundRedemptions")
}

// FundRedemptions is a paid mutator transaction binding the contract method 0x281155ba.
//
// Solidity: function fundRedemptions() returns()
func (_MultiChoiceMarket *MultiChoiceMarketSession) FundRedemptions() (*types.Transaction, error) {
	return _MultiChoiceMarket.Contract.FundRedemptions(&_MultiChoiceMarket.TransactOpts)
}

// FundRedemptions is a paid mutator transaction binding the contract method 0x281155ba.
//
// Solidity: function fundRedemptions() returns()
func (_MultiChoiceMarket *MultiChoiceMarketTransactorSession) FundRedemptions() (*types.Transaction, error) {
	return _MultiChoiceMarket.Contract.FundRedemptions(&_MultiChoiceMarket.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_MultiChoiceMarket *MultiChoiceMarketTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MultiChoiceMarket.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_MultiChoiceMarket *MultiChoiceMarketSession) Pause() (*types.Transaction, error) {
	return _MultiChoiceMarket.Contract.Pause(&_MultiChoiceMarket.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_MultiChoiceMarket *MultiChoiceMarketTransactorSession) Pause() (*types.Transaction, error) {
	return _MultiChoiceMarket.Contract.Pause(&_MultiChoiceMarket.TransactOpts)
}

// RemoveLiquidity is a paid mutator transaction binding the contract method 0x9c8f9f23.
//
// Solidity: function removeLiquidity(uint256 lpTokens) returns(uint256 collateralOut)
func (_MultiChoiceMarket *MultiChoiceMarketTransactor) RemoveLiquidity(opts *bind.TransactOpts, lpTokens *big.Int) (*types.Transaction, error) {
	return _MultiChoiceMarket.contract.Transact(opts, "removeLiquidity", lpTokens)
}

// RemoveLiquidity is a paid mutator transaction binding the contract method 0x9c8f9f23.
//
// Solidity: function removeLiquidity(uint256 lpTokens) returns(uint256 collateralOut)
func (_MultiChoiceMarket *MultiChoiceMarketSession) RemoveLiquidity(lpTokens *big.Int) (*types.Transaction, error) {
	return _MultiChoiceMarket.Contract.RemoveLiquidity(&_MultiChoiceMarket.TransactOpts, lpTokens)
}

// RemoveLiquidity is a paid mutator transaction binding the contract method 0x9c8f9f23.
//
// Solidity: function removeLiquidity(uint256 lpTokens) returns(uint256 collateralOut)
func (_MultiChoiceMarket *MultiChoiceMarketTransactorSession) RemoveLiquidity(lpTokens *big.Int) (*types.Transaction, error) {
	return _MultiChoiceMarket.Contract.RemoveLiquidity(&_MultiChoiceMarket.TransactOpts, lpTokens)
}

// Sell is a paid mutator transaction binding the contract method 0xd3c9727c.
//
// Solidity: function sell(uint256 outcomeId, uint256 tokensIn, uint256 minCollateralOut) returns(uint256 collateralOut)
func (_MultiChoiceMarket *MultiChoiceMarketTransactor) Sell(opts *bind.TransactOpts, outcomeId *big.Int, tokensIn *big.Int, minCollateralOut *big.Int) (*types.Transaction, error) {
	return _MultiChoiceMarket.contract.Transact(opts, "sell", outcomeId, tokensIn, minCollateralOut)
}

// Sell is a paid mutator transaction binding the contract method 0xd3c9727c.
//
// Solidity: function sell(uint256 outcomeId, uint256 tokensIn, uint256 minCollateralOut) returns(uint256 collateralOut)
func (_MultiChoiceMarket *MultiChoiceMarketSession) Sell(outcomeId *big.Int, tokensIn *big.Int, minCollateralOut *big.Int) (*types.Transaction, error) {
	return _MultiChoiceMarket.Contract.Sell(&_MultiChoiceMarket.TransactOpts, outcomeId, tokensIn, minCollateralOut)
}

// Sell is a paid mutator transaction binding the contract method 0xd3c9727c.
//
// Solidity: function sell(uint256 outcomeId, uint256 tokensIn, uint256 minCollateralOut) returns(uint256 collateralOut)
func (_MultiChoiceMarket *MultiChoiceMarketTransactorSession) Sell(outcomeId *big.Int, tokensIn *big.Int, minCollateralOut *big.Int) (*types.Transaction, error) {
	return _MultiChoiceMarket.Contract.Sell(&_MultiChoiceMarket.TransactOpts, outcomeId, tokensIn, minCollateralOut)
}

// SetAdmin is a paid mutator transaction binding the contract method 0x704b6c02.
//
// Solidity: function setAdmin(address newAdmin) returns()
func (_MultiChoiceMarket *MultiChoiceMarketTransactor) SetAdmin(opts *bind.TransactOpts, newAdmin common.Address) (*types.Transaction, error) {
	return _MultiChoiceMarket.contract.Transact(opts, "setAdmin", newAdmin)
}

// SetAdmin is a paid mutator transaction binding the contract method 0x704b6c02.
//
// Solidity: function setAdmin(address newAdmin) returns()
func (_MultiChoiceMarket *MultiChoiceMarketSession) SetAdmin(newAdmin common.Address) (*types.Transaction, error) {
	return _MultiChoiceMarket.Contract.SetAdmin(&_MultiChoiceMarket.TransactOpts, newAdmin)
}

// SetAdmin is a paid mutator transaction binding the contract method 0x704b6c02.
//
// Solidity: function setAdmin(address newAdmin) returns()
func (_MultiChoiceMarket *MultiChoiceMarketTransactorSession) SetAdmin(newAdmin common.Address) (*types.Transaction, error) {
	return _MultiChoiceMarket.Contract.SetAdmin(&_MultiChoiceMarket.TransactOpts, newAdmin)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_MultiChoiceMarket *MultiChoiceMarketTransactor) Transfer(opts *bind.TransactOpts, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _MultiChoiceMarket.contract.Transact(opts, "transfer", to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_MultiChoiceMarket *MultiChoiceMarketSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _MultiChoiceMarket.Contract.Transfer(&_MultiChoiceMarket.TransactOpts, to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_MultiChoiceMarket *MultiChoiceMarketTransactorSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _MultiChoiceMarket.Contract.Transfer(&_MultiChoiceMarket.TransactOpts, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_MultiChoiceMarket *MultiChoiceMarketTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _MultiChoiceMarket.contract.Transact(opts, "transferFrom", from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_MultiChoiceMarket *MultiChoiceMarketSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _MultiChoiceMarket.Contract.TransferFrom(&_MultiChoiceMarket.TransactOpts, from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_MultiChoiceMarket *MultiChoiceMarketTransactorSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _MultiChoiceMarket.Contract.TransferFrom(&_MultiChoiceMarket.TransactOpts, from, to, value)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_MultiChoiceMarket *MultiChoiceMarketTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MultiChoiceMarket.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_MultiChoiceMarket *MultiChoiceMarketSession) Unpause() (*types.Transaction, error) {
	return _MultiChoiceMarket.Contract.Unpause(&_MultiChoiceMarket.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_MultiChoiceMarket *MultiChoiceMarketTransactorSession) Unpause() (*types.Transaction, error) {
	return _MultiChoiceMarket.Contract.Unpause(&_MultiChoiceMarket.TransactOpts)
}

// MultiChoiceMarketApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the MultiChoiceMarket contract.
type MultiChoiceMarketApprovalIterator struct {
	Event *MultiChoiceMarketApproval // Event containing the contract specifics and raw log

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
func (it *MultiChoiceMarketApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MultiChoiceMarketApproval)
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
		it.Event = new(MultiChoiceMarketApproval)
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
func (it *MultiChoiceMarketApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MultiChoiceMarketApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MultiChoiceMarketApproval represents a Approval event raised by the MultiChoiceMarket contract.
type MultiChoiceMarketApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_MultiChoiceMarket *MultiChoiceMarketFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*MultiChoiceMarketApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _MultiChoiceMarket.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &MultiChoiceMarketApprovalIterator{contract: _MultiChoiceMarket.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_MultiChoiceMarket *MultiChoiceMarketFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *MultiChoiceMarketApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _MultiChoiceMarket.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MultiChoiceMarketApproval)
				if err := _MultiChoiceMarket.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_MultiChoiceMarket *MultiChoiceMarketFilterer) ParseApproval(log types.Log) (*MultiChoiceMarketApproval, error) {
	event := new(MultiChoiceMarketApproval)
	if err := _MultiChoiceMarket.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MultiChoiceMarketLiquidityChangedIterator is returned from FilterLiquidityChanged and is used to iterate over the raw logs and unpacked data for LiquidityChanged events raised by the MultiChoiceMarket contract.
type MultiChoiceMarketLiquidityChangedIterator struct {
	Event *MultiChoiceMarketLiquidityChanged // Event containing the contract specifics and raw log

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
func (it *MultiChoiceMarketLiquidityChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MultiChoiceMarketLiquidityChanged)
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
		it.Event = new(MultiChoiceMarketLiquidityChanged)
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
func (it *MultiChoiceMarketLiquidityChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MultiChoiceMarketLiquidityChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MultiChoiceMarketLiquidityChanged represents a LiquidityChanged event raised by the MultiChoiceMarket contract.
type MultiChoiceMarketLiquidityChanged struct {
	Provider   common.Address
	Amount     *big.Int
	IsAddition bool
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterLiquidityChanged is a free log retrieval operation binding the contract event 0xb029a6414a0c6d2e4fa2e5287326aa8a8c7191f9f5ced9799754a380471458d4.
//
// Solidity: event LiquidityChanged(address indexed provider, uint256 amount, bool isAddition)
func (_MultiChoiceMarket *MultiChoiceMarketFilterer) FilterLiquidityChanged(opts *bind.FilterOpts, provider []common.Address) (*MultiChoiceMarketLiquidityChangedIterator, error) {

	var providerRule []interface{}
	for _, providerItem := range provider {
		providerRule = append(providerRule, providerItem)
	}

	logs, sub, err := _MultiChoiceMarket.contract.FilterLogs(opts, "LiquidityChanged", providerRule)
	if err != nil {
		return nil, err
	}
	return &MultiChoiceMarketLiquidityChangedIterator{contract: _MultiChoiceMarket.contract, event: "LiquidityChanged", logs: logs, sub: sub}, nil
}

// WatchLiquidityChanged is a free log subscription operation binding the contract event 0xb029a6414a0c6d2e4fa2e5287326aa8a8c7191f9f5ced9799754a380471458d4.
//
// Solidity: event LiquidityChanged(address indexed provider, uint256 amount, bool isAddition)
func (_MultiChoiceMarket *MultiChoiceMarketFilterer) WatchLiquidityChanged(opts *bind.WatchOpts, sink chan<- *MultiChoiceMarketLiquidityChanged, provider []common.Address) (event.Subscription, error) {

	var providerRule []interface{}
	for _, providerItem := range provider {
		providerRule = append(providerRule, providerItem)
	}

	logs, sub, err := _MultiChoiceMarket.contract.WatchLogs(opts, "LiquidityChanged", providerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MultiChoiceMarketLiquidityChanged)
				if err := _MultiChoiceMarket.contract.UnpackLog(event, "LiquidityChanged", log); err != nil {
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

// ParseLiquidityChanged is a log parse operation binding the contract event 0xb029a6414a0c6d2e4fa2e5287326aa8a8c7191f9f5ced9799754a380471458d4.
//
// Solidity: event LiquidityChanged(address indexed provider, uint256 amount, bool isAddition)
func (_MultiChoiceMarket *MultiChoiceMarketFilterer) ParseLiquidityChanged(log types.Log) (*MultiChoiceMarketLiquidityChanged, error) {
	event := new(MultiChoiceMarketLiquidityChanged)
	if err := _MultiChoiceMarket.contract.UnpackLog(event, "LiquidityChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MultiChoiceMarketLiquidityParameterUpdatedIterator is returned from FilterLiquidityParameterUpdated and is used to iterate over the raw logs and unpacked data for LiquidityParameterUpdated events raised by the MultiChoiceMarket contract.
type MultiChoiceMarketLiquidityParameterUpdatedIterator struct {
	Event *MultiChoiceMarketLiquidityParameterUpdated // Event containing the contract specifics and raw log

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
func (it *MultiChoiceMarketLiquidityParameterUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MultiChoiceMarketLiquidityParameterUpdated)
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
		it.Event = new(MultiChoiceMarketLiquidityParameterUpdated)
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
func (it *MultiChoiceMarketLiquidityParameterUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MultiChoiceMarketLiquidityParameterUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MultiChoiceMarketLiquidityParameterUpdated represents a LiquidityParameterUpdated event raised by the MultiChoiceMarket contract.
type MultiChoiceMarketLiquidityParameterUpdated struct {
	OldB *big.Int
	NewB *big.Int
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterLiquidityParameterUpdated is a free log retrieval operation binding the contract event 0x6e515712d7c1b079931bc2e5abaaa2232d0bbaec22a28a19a2500a3790dbba94.
//
// Solidity: event LiquidityParameterUpdated(uint256 oldB, uint256 newB)
func (_MultiChoiceMarket *MultiChoiceMarketFilterer) FilterLiquidityParameterUpdated(opts *bind.FilterOpts) (*MultiChoiceMarketLiquidityParameterUpdatedIterator, error) {

	logs, sub, err := _MultiChoiceMarket.contract.FilterLogs(opts, "LiquidityParameterUpdated")
	if err != nil {
		return nil, err
	}
	return &MultiChoiceMarketLiquidityParameterUpdatedIterator{contract: _MultiChoiceMarket.contract, event: "LiquidityParameterUpdated", logs: logs, sub: sub}, nil
}

// WatchLiquidityParameterUpdated is a free log subscription operation binding the contract event 0x6e515712d7c1b079931bc2e5abaaa2232d0bbaec22a28a19a2500a3790dbba94.
//
// Solidity: event LiquidityParameterUpdated(uint256 oldB, uint256 newB)
func (_MultiChoiceMarket *MultiChoiceMarketFilterer) WatchLiquidityParameterUpdated(opts *bind.WatchOpts, sink chan<- *MultiChoiceMarketLiquidityParameterUpdated) (event.Subscription, error) {

	logs, sub, err := _MultiChoiceMarket.contract.WatchLogs(opts, "LiquidityParameterUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MultiChoiceMarketLiquidityParameterUpdated)
				if err := _MultiChoiceMarket.contract.UnpackLog(event, "LiquidityParameterUpdated", log); err != nil {
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

// ParseLiquidityParameterUpdated is a log parse operation binding the contract event 0x6e515712d7c1b079931bc2e5abaaa2232d0bbaec22a28a19a2500a3790dbba94.
//
// Solidity: event LiquidityParameterUpdated(uint256 oldB, uint256 newB)
func (_MultiChoiceMarket *MultiChoiceMarketFilterer) ParseLiquidityParameterUpdated(log types.Log) (*MultiChoiceMarketLiquidityParameterUpdated, error) {
	event := new(MultiChoiceMarketLiquidityParameterUpdated)
	if err := _MultiChoiceMarket.contract.UnpackLog(event, "LiquidityParameterUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MultiChoiceMarketOutcomeReservesUpdatedIterator is returned from FilterOutcomeReservesUpdated and is used to iterate over the raw logs and unpacked data for OutcomeReservesUpdated events raised by the MultiChoiceMarket contract.
type MultiChoiceMarketOutcomeReservesUpdatedIterator struct {
	Event *MultiChoiceMarketOutcomeReservesUpdated // Event containing the contract specifics and raw log

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
func (it *MultiChoiceMarketOutcomeReservesUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MultiChoiceMarketOutcomeReservesUpdated)
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
		it.Event = new(MultiChoiceMarketOutcomeReservesUpdated)
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
func (it *MultiChoiceMarketOutcomeReservesUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MultiChoiceMarketOutcomeReservesUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MultiChoiceMarketOutcomeReservesUpdated represents a OutcomeReservesUpdated event raised by the MultiChoiceMarket contract.
type MultiChoiceMarketOutcomeReservesUpdated struct {
	Reserves []*big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterOutcomeReservesUpdated is a free log retrieval operation binding the contract event 0x76777e72733ee878fe26202d0e294634bc9fb2c08f336f7f0188875327fab256.
//
// Solidity: event OutcomeReservesUpdated(uint256[] reserves)
func (_MultiChoiceMarket *MultiChoiceMarketFilterer) FilterOutcomeReservesUpdated(opts *bind.FilterOpts) (*MultiChoiceMarketOutcomeReservesUpdatedIterator, error) {

	logs, sub, err := _MultiChoiceMarket.contract.FilterLogs(opts, "OutcomeReservesUpdated")
	if err != nil {
		return nil, err
	}
	return &MultiChoiceMarketOutcomeReservesUpdatedIterator{contract: _MultiChoiceMarket.contract, event: "OutcomeReservesUpdated", logs: logs, sub: sub}, nil
}

// WatchOutcomeReservesUpdated is a free log subscription operation binding the contract event 0x76777e72733ee878fe26202d0e294634bc9fb2c08f336f7f0188875327fab256.
//
// Solidity: event OutcomeReservesUpdated(uint256[] reserves)
func (_MultiChoiceMarket *MultiChoiceMarketFilterer) WatchOutcomeReservesUpdated(opts *bind.WatchOpts, sink chan<- *MultiChoiceMarketOutcomeReservesUpdated) (event.Subscription, error) {

	logs, sub, err := _MultiChoiceMarket.contract.WatchLogs(opts, "OutcomeReservesUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MultiChoiceMarketOutcomeReservesUpdated)
				if err := _MultiChoiceMarket.contract.UnpackLog(event, "OutcomeReservesUpdated", log); err != nil {
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

// ParseOutcomeReservesUpdated is a log parse operation binding the contract event 0x76777e72733ee878fe26202d0e294634bc9fb2c08f336f7f0188875327fab256.
//
// Solidity: event OutcomeReservesUpdated(uint256[] reserves)
func (_MultiChoiceMarket *MultiChoiceMarketFilterer) ParseOutcomeReservesUpdated(log types.Log) (*MultiChoiceMarketOutcomeReservesUpdated, error) {
	event := new(MultiChoiceMarketOutcomeReservesUpdated)
	if err := _MultiChoiceMarket.contract.UnpackLog(event, "OutcomeReservesUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MultiChoiceMarketPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the MultiChoiceMarket contract.
type MultiChoiceMarketPausedIterator struct {
	Event *MultiChoiceMarketPaused // Event containing the contract specifics and raw log

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
func (it *MultiChoiceMarketPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MultiChoiceMarketPaused)
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
		it.Event = new(MultiChoiceMarketPaused)
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
func (it *MultiChoiceMarketPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MultiChoiceMarketPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MultiChoiceMarketPaused represents a Paused event raised by the MultiChoiceMarket contract.
type MultiChoiceMarketPaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_MultiChoiceMarket *MultiChoiceMarketFilterer) FilterPaused(opts *bind.FilterOpts) (*MultiChoiceMarketPausedIterator, error) {

	logs, sub, err := _MultiChoiceMarket.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &MultiChoiceMarketPausedIterator{contract: _MultiChoiceMarket.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_MultiChoiceMarket *MultiChoiceMarketFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *MultiChoiceMarketPaused) (event.Subscription, error) {

	logs, sub, err := _MultiChoiceMarket.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MultiChoiceMarketPaused)
				if err := _MultiChoiceMarket.contract.UnpackLog(event, "Paused", log); err != nil {
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

// ParsePaused is a log parse operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_MultiChoiceMarket *MultiChoiceMarketFilterer) ParsePaused(log types.Log) (*MultiChoiceMarketPaused, error) {
	event := new(MultiChoiceMarketPaused)
	if err := _MultiChoiceMarket.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MultiChoiceMarketTradeIterator is returned from FilterTrade and is used to iterate over the raw logs and unpacked data for Trade events raised by the MultiChoiceMarket contract.
type MultiChoiceMarketTradeIterator struct {
	Event *MultiChoiceMarketTrade // Event containing the contract specifics and raw log

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
func (it *MultiChoiceMarketTradeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MultiChoiceMarketTrade)
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
		it.Event = new(MultiChoiceMarketTrade)
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
func (it *MultiChoiceMarketTradeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MultiChoiceMarketTradeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MultiChoiceMarketTrade represents a Trade event raised by the MultiChoiceMarket contract.
type MultiChoiceMarketTrade struct {
	Trader    common.Address
	OutcomeId *big.Int
	AmountIn  *big.Int
	AmountOut *big.Int
	Fee       *big.Int
	IsBuy     bool
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterTrade is a free log retrieval operation binding the contract event 0xe34b2a81bbc1e1a545c34243f3cc283b6b6d1f4c2153be2a47b2612247e45865.
//
// Solidity: event Trade(address indexed trader, uint256 indexed outcomeId, uint256 amountIn, uint256 amountOut, uint256 fee, bool isBuy)
func (_MultiChoiceMarket *MultiChoiceMarketFilterer) FilterTrade(opts *bind.FilterOpts, trader []common.Address, outcomeId []*big.Int) (*MultiChoiceMarketTradeIterator, error) {

	var traderRule []interface{}
	for _, traderItem := range trader {
		traderRule = append(traderRule, traderItem)
	}
	var outcomeIdRule []interface{}
	for _, outcomeIdItem := range outcomeId {
		outcomeIdRule = append(outcomeIdRule, outcomeIdItem)
	}

	logs, sub, err := _MultiChoiceMarket.contract.FilterLogs(opts, "Trade", traderRule, outcomeIdRule)
	if err != nil {
		return nil, err
	}
	return &MultiChoiceMarketTradeIterator{contract: _MultiChoiceMarket.contract, event: "Trade", logs: logs, sub: sub}, nil
}

// WatchTrade is a free log subscription operation binding the contract event 0xe34b2a81bbc1e1a545c34243f3cc283b6b6d1f4c2153be2a47b2612247e45865.
//
// Solidity: event Trade(address indexed trader, uint256 indexed outcomeId, uint256 amountIn, uint256 amountOut, uint256 fee, bool isBuy)
func (_MultiChoiceMarket *MultiChoiceMarketFilterer) WatchTrade(opts *bind.WatchOpts, sink chan<- *MultiChoiceMarketTrade, trader []common.Address, outcomeId []*big.Int) (event.Subscription, error) {

	var traderRule []interface{}
	for _, traderItem := range trader {
		traderRule = append(traderRule, traderItem)
	}
	var outcomeIdRule []interface{}
	for _, outcomeIdItem := range outcomeId {
		outcomeIdRule = append(outcomeIdRule, outcomeIdItem)
	}

	logs, sub, err := _MultiChoiceMarket.contract.WatchLogs(opts, "Trade", traderRule, outcomeIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MultiChoiceMarketTrade)
				if err := _MultiChoiceMarket.contract.UnpackLog(event, "Trade", log); err != nil {
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

// ParseTrade is a log parse operation binding the contract event 0xe34b2a81bbc1e1a545c34243f3cc283b6b6d1f4c2153be2a47b2612247e45865.
//
// Solidity: event Trade(address indexed trader, uint256 indexed outcomeId, uint256 amountIn, uint256 amountOut, uint256 fee, bool isBuy)
func (_MultiChoiceMarket *MultiChoiceMarketFilterer) ParseTrade(log types.Log) (*MultiChoiceMarketTrade, error) {
	event := new(MultiChoiceMarketTrade)
	if err := _MultiChoiceMarket.contract.UnpackLog(event, "Trade", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MultiChoiceMarketTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the MultiChoiceMarket contract.
type MultiChoiceMarketTransferIterator struct {
	Event *MultiChoiceMarketTransfer // Event containing the contract specifics and raw log

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
func (it *MultiChoiceMarketTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MultiChoiceMarketTransfer)
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
		it.Event = new(MultiChoiceMarketTransfer)
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
func (it *MultiChoiceMarketTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MultiChoiceMarketTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MultiChoiceMarketTransfer represents a Transfer event raised by the MultiChoiceMarket contract.
type MultiChoiceMarketTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_MultiChoiceMarket *MultiChoiceMarketFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*MultiChoiceMarketTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _MultiChoiceMarket.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &MultiChoiceMarketTransferIterator{contract: _MultiChoiceMarket.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_MultiChoiceMarket *MultiChoiceMarketFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *MultiChoiceMarketTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _MultiChoiceMarket.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MultiChoiceMarketTransfer)
				if err := _MultiChoiceMarket.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_MultiChoiceMarket *MultiChoiceMarketFilterer) ParseTransfer(log types.Log) (*MultiChoiceMarketTransfer, error) {
	event := new(MultiChoiceMarketTransfer)
	if err := _MultiChoiceMarket.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MultiChoiceMarketUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the MultiChoiceMarket contract.
type MultiChoiceMarketUnpausedIterator struct {
	Event *MultiChoiceMarketUnpaused // Event containing the contract specifics and raw log

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
func (it *MultiChoiceMarketUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MultiChoiceMarketUnpaused)
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
		it.Event = new(MultiChoiceMarketUnpaused)
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
func (it *MultiChoiceMarketUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MultiChoiceMarketUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MultiChoiceMarketUnpaused represents a Unpaused event raised by the MultiChoiceMarket contract.
type MultiChoiceMarketUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_MultiChoiceMarket *MultiChoiceMarketFilterer) FilterUnpaused(opts *bind.FilterOpts) (*MultiChoiceMarketUnpausedIterator, error) {

	logs, sub, err := _MultiChoiceMarket.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &MultiChoiceMarketUnpausedIterator{contract: _MultiChoiceMarket.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_MultiChoiceMarket *MultiChoiceMarketFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *MultiChoiceMarketUnpaused) (event.Subscription, error) {

	logs, sub, err := _MultiChoiceMarket.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MultiChoiceMarketUnpaused)
				if err := _MultiChoiceMarket.contract.UnpackLog(event, "Unpaused", log); err != nil {
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

// ParseUnpaused is a log parse operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_MultiChoiceMarket *MultiChoiceMarketFilterer) ParseUnpaused(log types.Log) (*MultiChoiceMarketUnpaused, error) {
	event := new(MultiChoiceMarketUnpaused)
	if err := _MultiChoiceMarket.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
