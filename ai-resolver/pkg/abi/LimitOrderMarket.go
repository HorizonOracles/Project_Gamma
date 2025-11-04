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


// LimitOrderMarketOrder is an auto generated low-level Go binding around an user-defined struct.
type LimitOrderMarketOrder struct {
	Trader    common.Address
	OutcomeId uint8
	Price     *big.Int
	Amount    *big.Int
	Filled    *big.Int
	TotalCost *big.Int
	Timestamp *big.Int
	IsBuy     bool
	OrderType uint8
	IsActive  bool
}

// LimitOrderMarketMetaData contains all meta data concerning the LimitOrderMarket contract.
var LimitOrderMarketMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_marketId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_collateralToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_outcomeToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_feeSplitter\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_horizonPerks\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_closeTime\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_outcomeCount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"MAX_ORDERS_PER_SIDE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"MINIMUM_LIQUIDITY\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"PRICE_PRECISION\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"addLiquidity\",\"inputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"admin\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"allowance\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"spender\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"approve\",\"inputs\":[{\"name\":\"spender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"balanceOf\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"buy\",\"inputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"buyOrdersByOutcome\",\"inputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"cancelOrder\",\"inputs\":[{\"name\":\"orderId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"closeTime\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"collateralToken\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIERC20\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"decimals\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"feeSplitter\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractFeeSplitter\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"fundRedemptions\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getBestAsk\",\"inputs\":[{\"name\":\"outcomeId\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getBestBid\",\"inputs\":[{\"name\":\"outcomeId\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getBuyOrders\",\"inputs\":[{\"name\":\"outcomeId\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getMarketInfo\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structIMarket.MarketInfo\",\"components\":[{\"name\":\"marketId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"marketType\",\"type\":\"uint8\",\"internalType\":\"enumIMarket.MarketType\"},{\"name\":\"collateralToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"closeTime\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"outcomeCount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"isResolved\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"isPaused\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getMarketType\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"enumIMarket.MarketType\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getOrder\",\"inputs\":[{\"name\":\"orderId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structLimitOrderMarket.Order\",\"components\":[{\"name\":\"trader\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"outcomeId\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"price\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"filled\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"totalCost\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"timestamp\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"isBuy\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"orderType\",\"type\":\"uint8\",\"internalType\":\"enumLimitOrderMarket.OrderType\"},{\"name\":\"isActive\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getOutcomeCount\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getPrice\",\"inputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"getQuoteBuy\",\"inputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"getQuoteSell\",\"inputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"getSellOrders\",\"inputs\":[{\"name\":\"outcomeId\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getSpread\",\"inputs\":[{\"name\":\"outcomeId\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"outputs\":[{\"name\":\"spread\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getUserOrders\",\"inputs\":[{\"name\":\"user\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"horizonPerks\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractHorizonPerks\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"marketId\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"marketType\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"enumIMarket.MarketType\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"name\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"onERC1155BatchReceived\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"},{\"name\":\"\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"},{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"onERC1155Received\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"orders\",\"inputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"trader\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"outcomeId\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"price\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"filled\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"totalCost\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"timestamp\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"isBuy\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"orderType\",\"type\":\"uint8\",\"internalType\":\"enumLimitOrderMarket.OrderType\"},{\"name\":\"isActive\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"outcomeCount\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"outcomeToken\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractOutcomeToken\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pause\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"paused\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"placeMarketOrder\",\"inputs\":[{\"name\":\"outcomeId\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"isBuy\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"maxSlippage\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"filledAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"placeOrder\",\"inputs\":[{\"name\":\"outcomeId\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"isBuy\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"price\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"orderType\",\"type\":\"uint8\",\"internalType\":\"enumLimitOrderMarket.OrderType\"}],\"outputs\":[{\"name\":\"orderId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"removeLiquidity\",\"inputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"sell\",\"inputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"sellOrdersByOutcome\",\"inputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"setAdmin\",\"inputs\":[{\"name\":\"newAdmin\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"supportsInterface\",\"inputs\":[{\"name\":\"interfaceId\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"symbol\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"totalCollateral\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"totalSupply\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"transfer\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferFrom\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"unpause\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"userOrders\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"Approval\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"spender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"LiquidityChanged\",\"inputs\":[{\"name\":\"provider\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"isAddition\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OrderCancelled\",\"inputs\":[{\"name\":\"orderId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"trader\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amountRemaining\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OrderMatched\",\"inputs\":[{\"name\":\"buyOrderId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"sellOrderId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"outcomeId\",\"type\":\"uint8\",\"indexed\":true,\"internalType\":\"uint8\"},{\"name\":\"price\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"buyer\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"seller\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OrderPartiallyFilled\",\"inputs\":[{\"name\":\"orderId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"filledAmount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"remainingAmount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OrderPlaced\",\"inputs\":[{\"name\":\"orderId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"trader\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"outcomeId\",\"type\":\"uint8\",\"indexed\":true,\"internalType\":\"uint8\"},{\"name\":\"isBuy\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"},{\"name\":\"price\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"orderType\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"enumLimitOrderMarket.OrderType\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Paused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Trade\",\"inputs\":[{\"name\":\"trader\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"outcomeId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"amountIn\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"amountOut\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"fee\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"isBuy\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Transfer\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Unpaused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"ERC20InsufficientAllowance\",\"inputs\":[{\"name\":\"spender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"allowance\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"needed\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"ERC20InsufficientBalance\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"balance\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"needed\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"ERC20InvalidApprover\",\"inputs\":[{\"name\":\"approver\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ERC20InvalidReceiver\",\"inputs\":[{\"name\":\"receiver\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ERC20InvalidSender\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ERC20InvalidSpender\",\"inputs\":[{\"name\":\"spender\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"EnforcedPause\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ExpectedPause\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InsufficientLPTokens\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InsufficientLiquidity\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidAmount\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidOutcomeId\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidPrice\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidState\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"MarketClosed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"MarketResolved\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"MaxOrdersReached\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"MinimumLiquidityRequired\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NoMatchingOrders\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OrderAlreadyFilled\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OrderNotFound\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OrderNotOwned\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"PostOnlyWouldMatch\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ReentrancyGuardReentrantCall\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"SafeERC20FailedOperation\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"SlippageExceeded\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"Unauthorized\",\"inputs\":[]}]",
}

// LimitOrderMarketABI is the input ABI used to generate the binding from.
// Deprecated: Use LimitOrderMarketMetaData.ABI instead.
var LimitOrderMarketABI = LimitOrderMarketMetaData.ABI

// LimitOrderMarket is an auto generated Go binding around an Ethereum contract.
type LimitOrderMarket struct {
	LimitOrderMarketCaller     // Read-only binding to the contract
	LimitOrderMarketTransactor // Write-only binding to the contract
	LimitOrderMarketFilterer   // Log filterer for contract events
}

// LimitOrderMarketCaller is an auto generated read-only Go binding around an Ethereum contract.
type LimitOrderMarketCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LimitOrderMarketTransactor is an auto generated write-only Go binding around an Ethereum contract.
type LimitOrderMarketTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LimitOrderMarketFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type LimitOrderMarketFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LimitOrderMarketSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type LimitOrderMarketSession struct {
	Contract     *LimitOrderMarket // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// LimitOrderMarketCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type LimitOrderMarketCallerSession struct {
	Contract *LimitOrderMarketCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// LimitOrderMarketTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type LimitOrderMarketTransactorSession struct {
	Contract     *LimitOrderMarketTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// LimitOrderMarketRaw is an auto generated low-level Go binding around an Ethereum contract.
type LimitOrderMarketRaw struct {
	Contract *LimitOrderMarket // Generic contract binding to access the raw methods on
}

// LimitOrderMarketCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type LimitOrderMarketCallerRaw struct {
	Contract *LimitOrderMarketCaller // Generic read-only contract binding to access the raw methods on
}

// LimitOrderMarketTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type LimitOrderMarketTransactorRaw struct {
	Contract *LimitOrderMarketTransactor // Generic write-only contract binding to access the raw methods on
}

// NewLimitOrderMarket creates a new instance of LimitOrderMarket, bound to a specific deployed contract.
func NewLimitOrderMarket(address common.Address, backend bind.ContractBackend) (*LimitOrderMarket, error) {
	contract, err := bindLimitOrderMarket(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &LimitOrderMarket{LimitOrderMarketCaller: LimitOrderMarketCaller{contract: contract}, LimitOrderMarketTransactor: LimitOrderMarketTransactor{contract: contract}, LimitOrderMarketFilterer: LimitOrderMarketFilterer{contract: contract}}, nil
}

// NewLimitOrderMarketCaller creates a new read-only instance of LimitOrderMarket, bound to a specific deployed contract.
func NewLimitOrderMarketCaller(address common.Address, caller bind.ContractCaller) (*LimitOrderMarketCaller, error) {
	contract, err := bindLimitOrderMarket(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &LimitOrderMarketCaller{contract: contract}, nil
}

// NewLimitOrderMarketTransactor creates a new write-only instance of LimitOrderMarket, bound to a specific deployed contract.
func NewLimitOrderMarketTransactor(address common.Address, transactor bind.ContractTransactor) (*LimitOrderMarketTransactor, error) {
	contract, err := bindLimitOrderMarket(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &LimitOrderMarketTransactor{contract: contract}, nil
}

// NewLimitOrderMarketFilterer creates a new log filterer instance of LimitOrderMarket, bound to a specific deployed contract.
func NewLimitOrderMarketFilterer(address common.Address, filterer bind.ContractFilterer) (*LimitOrderMarketFilterer, error) {
	contract, err := bindLimitOrderMarket(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &LimitOrderMarketFilterer{contract: contract}, nil
}

// bindLimitOrderMarket binds a generic wrapper to an already deployed contract.
func bindLimitOrderMarket(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := LimitOrderMarketMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_LimitOrderMarket *LimitOrderMarketRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _LimitOrderMarket.Contract.LimitOrderMarketCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_LimitOrderMarket *LimitOrderMarketRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LimitOrderMarket.Contract.LimitOrderMarketTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_LimitOrderMarket *LimitOrderMarketRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _LimitOrderMarket.Contract.LimitOrderMarketTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_LimitOrderMarket *LimitOrderMarketCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _LimitOrderMarket.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_LimitOrderMarket *LimitOrderMarketTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LimitOrderMarket.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_LimitOrderMarket *LimitOrderMarketTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _LimitOrderMarket.Contract.contract.Transact(opts, method, params...)
}

// MAXORDERSPERSIDE is a free data retrieval call binding the contract method 0x771e3819.
//
// Solidity: function MAX_ORDERS_PER_SIDE() view returns(uint256)
func (_LimitOrderMarket *LimitOrderMarketCaller) MAXORDERSPERSIDE(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _LimitOrderMarket.contract.Call(opts, &out, "MAX_ORDERS_PER_SIDE")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXORDERSPERSIDE is a free data retrieval call binding the contract method 0x771e3819.
//
// Solidity: function MAX_ORDERS_PER_SIDE() view returns(uint256)
func (_LimitOrderMarket *LimitOrderMarketSession) MAXORDERSPERSIDE() (*big.Int, error) {
	return _LimitOrderMarket.Contract.MAXORDERSPERSIDE(&_LimitOrderMarket.CallOpts)
}

// MAXORDERSPERSIDE is a free data retrieval call binding the contract method 0x771e3819.
//
// Solidity: function MAX_ORDERS_PER_SIDE() view returns(uint256)
func (_LimitOrderMarket *LimitOrderMarketCallerSession) MAXORDERSPERSIDE() (*big.Int, error) {
	return _LimitOrderMarket.Contract.MAXORDERSPERSIDE(&_LimitOrderMarket.CallOpts)
}

// MINIMUMLIQUIDITY is a free data retrieval call binding the contract method 0xba9a7a56.
//
// Solidity: function MINIMUM_LIQUIDITY() view returns(uint256)
func (_LimitOrderMarket *LimitOrderMarketCaller) MINIMUMLIQUIDITY(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _LimitOrderMarket.contract.Call(opts, &out, "MINIMUM_LIQUIDITY")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MINIMUMLIQUIDITY is a free data retrieval call binding the contract method 0xba9a7a56.
//
// Solidity: function MINIMUM_LIQUIDITY() view returns(uint256)
func (_LimitOrderMarket *LimitOrderMarketSession) MINIMUMLIQUIDITY() (*big.Int, error) {
	return _LimitOrderMarket.Contract.MINIMUMLIQUIDITY(&_LimitOrderMarket.CallOpts)
}

// MINIMUMLIQUIDITY is a free data retrieval call binding the contract method 0xba9a7a56.
//
// Solidity: function MINIMUM_LIQUIDITY() view returns(uint256)
func (_LimitOrderMarket *LimitOrderMarketCallerSession) MINIMUMLIQUIDITY() (*big.Int, error) {
	return _LimitOrderMarket.Contract.MINIMUMLIQUIDITY(&_LimitOrderMarket.CallOpts)
}

// PRICEPRECISION is a free data retrieval call binding the contract method 0x95082d25.
//
// Solidity: function PRICE_PRECISION() view returns(uint256)
func (_LimitOrderMarket *LimitOrderMarketCaller) PRICEPRECISION(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _LimitOrderMarket.contract.Call(opts, &out, "PRICE_PRECISION")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PRICEPRECISION is a free data retrieval call binding the contract method 0x95082d25.
//
// Solidity: function PRICE_PRECISION() view returns(uint256)
func (_LimitOrderMarket *LimitOrderMarketSession) PRICEPRECISION() (*big.Int, error) {
	return _LimitOrderMarket.Contract.PRICEPRECISION(&_LimitOrderMarket.CallOpts)
}

// PRICEPRECISION is a free data retrieval call binding the contract method 0x95082d25.
//
// Solidity: function PRICE_PRECISION() view returns(uint256)
func (_LimitOrderMarket *LimitOrderMarketCallerSession) PRICEPRECISION() (*big.Int, error) {
	return _LimitOrderMarket.Contract.PRICEPRECISION(&_LimitOrderMarket.CallOpts)
}

// AddLiquidity is a free data retrieval call binding the contract method 0x51c6590a.
//
// Solidity: function addLiquidity(uint256 ) pure returns(uint256)
func (_LimitOrderMarket *LimitOrderMarketCaller) AddLiquidity(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _LimitOrderMarket.contract.Call(opts, &out, "addLiquidity", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AddLiquidity is a free data retrieval call binding the contract method 0x51c6590a.
//
// Solidity: function addLiquidity(uint256 ) pure returns(uint256)
func (_LimitOrderMarket *LimitOrderMarketSession) AddLiquidity(arg0 *big.Int) (*big.Int, error) {
	return _LimitOrderMarket.Contract.AddLiquidity(&_LimitOrderMarket.CallOpts, arg0)
}

// AddLiquidity is a free data retrieval call binding the contract method 0x51c6590a.
//
// Solidity: function addLiquidity(uint256 ) pure returns(uint256)
func (_LimitOrderMarket *LimitOrderMarketCallerSession) AddLiquidity(arg0 *big.Int) (*big.Int, error) {
	return _LimitOrderMarket.Contract.AddLiquidity(&_LimitOrderMarket.CallOpts, arg0)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_LimitOrderMarket *LimitOrderMarketCaller) Admin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _LimitOrderMarket.contract.Call(opts, &out, "admin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_LimitOrderMarket *LimitOrderMarketSession) Admin() (common.Address, error) {
	return _LimitOrderMarket.Contract.Admin(&_LimitOrderMarket.CallOpts)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_LimitOrderMarket *LimitOrderMarketCallerSession) Admin() (common.Address, error) {
	return _LimitOrderMarket.Contract.Admin(&_LimitOrderMarket.CallOpts)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_LimitOrderMarket *LimitOrderMarketCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _LimitOrderMarket.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_LimitOrderMarket *LimitOrderMarketSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _LimitOrderMarket.Contract.Allowance(&_LimitOrderMarket.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_LimitOrderMarket *LimitOrderMarketCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _LimitOrderMarket.Contract.Allowance(&_LimitOrderMarket.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_LimitOrderMarket *LimitOrderMarketCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _LimitOrderMarket.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_LimitOrderMarket *LimitOrderMarketSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _LimitOrderMarket.Contract.BalanceOf(&_LimitOrderMarket.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_LimitOrderMarket *LimitOrderMarketCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _LimitOrderMarket.Contract.BalanceOf(&_LimitOrderMarket.CallOpts, account)
}

// Buy is a free data retrieval call binding the contract method 0x40993b26.
//
// Solidity: function buy(uint256 , uint256 , uint256 ) pure returns(uint256)
func (_LimitOrderMarket *LimitOrderMarketCaller) Buy(opts *bind.CallOpts, arg0 *big.Int, arg1 *big.Int, arg2 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _LimitOrderMarket.contract.Call(opts, &out, "buy", arg0, arg1, arg2)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Buy is a free data retrieval call binding the contract method 0x40993b26.
//
// Solidity: function buy(uint256 , uint256 , uint256 ) pure returns(uint256)
func (_LimitOrderMarket *LimitOrderMarketSession) Buy(arg0 *big.Int, arg1 *big.Int, arg2 *big.Int) (*big.Int, error) {
	return _LimitOrderMarket.Contract.Buy(&_LimitOrderMarket.CallOpts, arg0, arg1, arg2)
}

// Buy is a free data retrieval call binding the contract method 0x40993b26.
//
// Solidity: function buy(uint256 , uint256 , uint256 ) pure returns(uint256)
func (_LimitOrderMarket *LimitOrderMarketCallerSession) Buy(arg0 *big.Int, arg1 *big.Int, arg2 *big.Int) (*big.Int, error) {
	return _LimitOrderMarket.Contract.Buy(&_LimitOrderMarket.CallOpts, arg0, arg1, arg2)
}

// BuyOrdersByOutcome is a free data retrieval call binding the contract method 0xf78898dc.
//
// Solidity: function buyOrdersByOutcome(uint8 , uint256 ) view returns(bytes32)
func (_LimitOrderMarket *LimitOrderMarketCaller) BuyOrdersByOutcome(opts *bind.CallOpts, arg0 uint8, arg1 *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _LimitOrderMarket.contract.Call(opts, &out, "buyOrdersByOutcome", arg0, arg1)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// BuyOrdersByOutcome is a free data retrieval call binding the contract method 0xf78898dc.
//
// Solidity: function buyOrdersByOutcome(uint8 , uint256 ) view returns(bytes32)
func (_LimitOrderMarket *LimitOrderMarketSession) BuyOrdersByOutcome(arg0 uint8, arg1 *big.Int) ([32]byte, error) {
	return _LimitOrderMarket.Contract.BuyOrdersByOutcome(&_LimitOrderMarket.CallOpts, arg0, arg1)
}

// BuyOrdersByOutcome is a free data retrieval call binding the contract method 0xf78898dc.
//
// Solidity: function buyOrdersByOutcome(uint8 , uint256 ) view returns(bytes32)
func (_LimitOrderMarket *LimitOrderMarketCallerSession) BuyOrdersByOutcome(arg0 uint8, arg1 *big.Int) ([32]byte, error) {
	return _LimitOrderMarket.Contract.BuyOrdersByOutcome(&_LimitOrderMarket.CallOpts, arg0, arg1)
}

// CloseTime is a free data retrieval call binding the contract method 0x627749e6.
//
// Solidity: function closeTime() view returns(uint256)
func (_LimitOrderMarket *LimitOrderMarketCaller) CloseTime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _LimitOrderMarket.contract.Call(opts, &out, "closeTime")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CloseTime is a free data retrieval call binding the contract method 0x627749e6.
//
// Solidity: function closeTime() view returns(uint256)
func (_LimitOrderMarket *LimitOrderMarketSession) CloseTime() (*big.Int, error) {
	return _LimitOrderMarket.Contract.CloseTime(&_LimitOrderMarket.CallOpts)
}

// CloseTime is a free data retrieval call binding the contract method 0x627749e6.
//
// Solidity: function closeTime() view returns(uint256)
func (_LimitOrderMarket *LimitOrderMarketCallerSession) CloseTime() (*big.Int, error) {
	return _LimitOrderMarket.Contract.CloseTime(&_LimitOrderMarket.CallOpts)
}

// CollateralToken is a free data retrieval call binding the contract method 0xb2016bd4.
//
// Solidity: function collateralToken() view returns(address)
func (_LimitOrderMarket *LimitOrderMarketCaller) CollateralToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _LimitOrderMarket.contract.Call(opts, &out, "collateralToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CollateralToken is a free data retrieval call binding the contract method 0xb2016bd4.
//
// Solidity: function collateralToken() view returns(address)
func (_LimitOrderMarket *LimitOrderMarketSession) CollateralToken() (common.Address, error) {
	return _LimitOrderMarket.Contract.CollateralToken(&_LimitOrderMarket.CallOpts)
}

// CollateralToken is a free data retrieval call binding the contract method 0xb2016bd4.
//
// Solidity: function collateralToken() view returns(address)
func (_LimitOrderMarket *LimitOrderMarketCallerSession) CollateralToken() (common.Address, error) {
	return _LimitOrderMarket.Contract.CollateralToken(&_LimitOrderMarket.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_LimitOrderMarket *LimitOrderMarketCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _LimitOrderMarket.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_LimitOrderMarket *LimitOrderMarketSession) Decimals() (uint8, error) {
	return _LimitOrderMarket.Contract.Decimals(&_LimitOrderMarket.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_LimitOrderMarket *LimitOrderMarketCallerSession) Decimals() (uint8, error) {
	return _LimitOrderMarket.Contract.Decimals(&_LimitOrderMarket.CallOpts)
}

// FeeSplitter is a free data retrieval call binding the contract method 0x6052970c.
//
// Solidity: function feeSplitter() view returns(address)
func (_LimitOrderMarket *LimitOrderMarketCaller) FeeSplitter(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _LimitOrderMarket.contract.Call(opts, &out, "feeSplitter")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FeeSplitter is a free data retrieval call binding the contract method 0x6052970c.
//
// Solidity: function feeSplitter() view returns(address)
func (_LimitOrderMarket *LimitOrderMarketSession) FeeSplitter() (common.Address, error) {
	return _LimitOrderMarket.Contract.FeeSplitter(&_LimitOrderMarket.CallOpts)
}

// FeeSplitter is a free data retrieval call binding the contract method 0x6052970c.
//
// Solidity: function feeSplitter() view returns(address)
func (_LimitOrderMarket *LimitOrderMarketCallerSession) FeeSplitter() (common.Address, error) {
	return _LimitOrderMarket.Contract.FeeSplitter(&_LimitOrderMarket.CallOpts)
}

// GetBestAsk is a free data retrieval call binding the contract method 0xd965ce64.
//
// Solidity: function getBestAsk(uint8 outcomeId) view returns(uint256)
func (_LimitOrderMarket *LimitOrderMarketCaller) GetBestAsk(opts *bind.CallOpts, outcomeId uint8) (*big.Int, error) {
	var out []interface{}
	err := _LimitOrderMarket.contract.Call(opts, &out, "getBestAsk", outcomeId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetBestAsk is a free data retrieval call binding the contract method 0xd965ce64.
//
// Solidity: function getBestAsk(uint8 outcomeId) view returns(uint256)
func (_LimitOrderMarket *LimitOrderMarketSession) GetBestAsk(outcomeId uint8) (*big.Int, error) {
	return _LimitOrderMarket.Contract.GetBestAsk(&_LimitOrderMarket.CallOpts, outcomeId)
}

// GetBestAsk is a free data retrieval call binding the contract method 0xd965ce64.
//
// Solidity: function getBestAsk(uint8 outcomeId) view returns(uint256)
func (_LimitOrderMarket *LimitOrderMarketCallerSession) GetBestAsk(outcomeId uint8) (*big.Int, error) {
	return _LimitOrderMarket.Contract.GetBestAsk(&_LimitOrderMarket.CallOpts, outcomeId)
}

// GetBestBid is a free data retrieval call binding the contract method 0x7aa582c1.
//
// Solidity: function getBestBid(uint8 outcomeId) view returns(uint256)
func (_LimitOrderMarket *LimitOrderMarketCaller) GetBestBid(opts *bind.CallOpts, outcomeId uint8) (*big.Int, error) {
	var out []interface{}
	err := _LimitOrderMarket.contract.Call(opts, &out, "getBestBid", outcomeId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetBestBid is a free data retrieval call binding the contract method 0x7aa582c1.
//
// Solidity: function getBestBid(uint8 outcomeId) view returns(uint256)
func (_LimitOrderMarket *LimitOrderMarketSession) GetBestBid(outcomeId uint8) (*big.Int, error) {
	return _LimitOrderMarket.Contract.GetBestBid(&_LimitOrderMarket.CallOpts, outcomeId)
}

// GetBestBid is a free data retrieval call binding the contract method 0x7aa582c1.
//
// Solidity: function getBestBid(uint8 outcomeId) view returns(uint256)
func (_LimitOrderMarket *LimitOrderMarketCallerSession) GetBestBid(outcomeId uint8) (*big.Int, error) {
	return _LimitOrderMarket.Contract.GetBestBid(&_LimitOrderMarket.CallOpts, outcomeId)
}

// GetBuyOrders is a free data retrieval call binding the contract method 0x3bb15184.
//
// Solidity: function getBuyOrders(uint8 outcomeId) view returns(bytes32[])
func (_LimitOrderMarket *LimitOrderMarketCaller) GetBuyOrders(opts *bind.CallOpts, outcomeId uint8) ([][32]byte, error) {
	var out []interface{}
	err := _LimitOrderMarket.contract.Call(opts, &out, "getBuyOrders", outcomeId)

	if err != nil {
		return *new([][32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([][32]byte)).(*[][32]byte)

	return out0, err

}

// GetBuyOrders is a free data retrieval call binding the contract method 0x3bb15184.
//
// Solidity: function getBuyOrders(uint8 outcomeId) view returns(bytes32[])
func (_LimitOrderMarket *LimitOrderMarketSession) GetBuyOrders(outcomeId uint8) ([][32]byte, error) {
	return _LimitOrderMarket.Contract.GetBuyOrders(&_LimitOrderMarket.CallOpts, outcomeId)
}

// GetBuyOrders is a free data retrieval call binding the contract method 0x3bb15184.
//
// Solidity: function getBuyOrders(uint8 outcomeId) view returns(bytes32[])
func (_LimitOrderMarket *LimitOrderMarketCallerSession) GetBuyOrders(outcomeId uint8) ([][32]byte, error) {
	return _LimitOrderMarket.Contract.GetBuyOrders(&_LimitOrderMarket.CallOpts, outcomeId)
}

// GetMarketInfo is a free data retrieval call binding the contract method 0x23341a05.
//
// Solidity: function getMarketInfo() view returns((uint256,uint8,address,uint256,uint256,bool,bool))
func (_LimitOrderMarket *LimitOrderMarketCaller) GetMarketInfo(opts *bind.CallOpts) (IMarketMarketInfo, error) {
	var out []interface{}
	err := _LimitOrderMarket.contract.Call(opts, &out, "getMarketInfo")

	if err != nil {
		return *new(IMarketMarketInfo), err
	}

	out0 := *abi.ConvertType(out[0], new(IMarketMarketInfo)).(*IMarketMarketInfo)

	return out0, err

}

// GetMarketInfo is a free data retrieval call binding the contract method 0x23341a05.
//
// Solidity: function getMarketInfo() view returns((uint256,uint8,address,uint256,uint256,bool,bool))
func (_LimitOrderMarket *LimitOrderMarketSession) GetMarketInfo() (IMarketMarketInfo, error) {
	return _LimitOrderMarket.Contract.GetMarketInfo(&_LimitOrderMarket.CallOpts)
}

// GetMarketInfo is a free data retrieval call binding the contract method 0x23341a05.
//
// Solidity: function getMarketInfo() view returns((uint256,uint8,address,uint256,uint256,bool,bool))
func (_LimitOrderMarket *LimitOrderMarketCallerSession) GetMarketInfo() (IMarketMarketInfo, error) {
	return _LimitOrderMarket.Contract.GetMarketInfo(&_LimitOrderMarket.CallOpts)
}

// GetMarketType is a free data retrieval call binding the contract method 0x33e7a1d0.
//
// Solidity: function getMarketType() view returns(uint8)
func (_LimitOrderMarket *LimitOrderMarketCaller) GetMarketType(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _LimitOrderMarket.contract.Call(opts, &out, "getMarketType")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// GetMarketType is a free data retrieval call binding the contract method 0x33e7a1d0.
//
// Solidity: function getMarketType() view returns(uint8)
func (_LimitOrderMarket *LimitOrderMarketSession) GetMarketType() (uint8, error) {
	return _LimitOrderMarket.Contract.GetMarketType(&_LimitOrderMarket.CallOpts)
}

// GetMarketType is a free data retrieval call binding the contract method 0x33e7a1d0.
//
// Solidity: function getMarketType() view returns(uint8)
func (_LimitOrderMarket *LimitOrderMarketCallerSession) GetMarketType() (uint8, error) {
	return _LimitOrderMarket.Contract.GetMarketType(&_LimitOrderMarket.CallOpts)
}

// GetOrder is a free data retrieval call binding the contract method 0x5778472a.
//
// Solidity: function getOrder(bytes32 orderId) view returns((address,uint8,uint256,uint256,uint256,uint256,uint256,bool,uint8,bool))
func (_LimitOrderMarket *LimitOrderMarketCaller) GetOrder(opts *bind.CallOpts, orderId [32]byte) (LimitOrderMarketOrder, error) {
	var out []interface{}
	err := _LimitOrderMarket.contract.Call(opts, &out, "getOrder", orderId)

	if err != nil {
		return *new(LimitOrderMarketOrder), err
	}

	out0 := *abi.ConvertType(out[0], new(LimitOrderMarketOrder)).(*LimitOrderMarketOrder)

	return out0, err

}

// GetOrder is a free data retrieval call binding the contract method 0x5778472a.
//
// Solidity: function getOrder(bytes32 orderId) view returns((address,uint8,uint256,uint256,uint256,uint256,uint256,bool,uint8,bool))
func (_LimitOrderMarket *LimitOrderMarketSession) GetOrder(orderId [32]byte) (LimitOrderMarketOrder, error) {
	return _LimitOrderMarket.Contract.GetOrder(&_LimitOrderMarket.CallOpts, orderId)
}

// GetOrder is a free data retrieval call binding the contract method 0x5778472a.
//
// Solidity: function getOrder(bytes32 orderId) view returns((address,uint8,uint256,uint256,uint256,uint256,uint256,bool,uint8,bool))
func (_LimitOrderMarket *LimitOrderMarketCallerSession) GetOrder(orderId [32]byte) (LimitOrderMarketOrder, error) {
	return _LimitOrderMarket.Contract.GetOrder(&_LimitOrderMarket.CallOpts, orderId)
}

// GetOutcomeCount is a free data retrieval call binding the contract method 0x7dc8f086.
//
// Solidity: function getOutcomeCount() view returns(uint256)
func (_LimitOrderMarket *LimitOrderMarketCaller) GetOutcomeCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _LimitOrderMarket.contract.Call(opts, &out, "getOutcomeCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetOutcomeCount is a free data retrieval call binding the contract method 0x7dc8f086.
//
// Solidity: function getOutcomeCount() view returns(uint256)
func (_LimitOrderMarket *LimitOrderMarketSession) GetOutcomeCount() (*big.Int, error) {
	return _LimitOrderMarket.Contract.GetOutcomeCount(&_LimitOrderMarket.CallOpts)
}

// GetOutcomeCount is a free data retrieval call binding the contract method 0x7dc8f086.
//
// Solidity: function getOutcomeCount() view returns(uint256)
func (_LimitOrderMarket *LimitOrderMarketCallerSession) GetOutcomeCount() (*big.Int, error) {
	return _LimitOrderMarket.Contract.GetOutcomeCount(&_LimitOrderMarket.CallOpts)
}

// GetPrice is a free data retrieval call binding the contract method 0xe7572230.
//
// Solidity: function getPrice(uint256 ) pure returns(uint256)
func (_LimitOrderMarket *LimitOrderMarketCaller) GetPrice(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _LimitOrderMarket.contract.Call(opts, &out, "getPrice", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetPrice is a free data retrieval call binding the contract method 0xe7572230.
//
// Solidity: function getPrice(uint256 ) pure returns(uint256)
func (_LimitOrderMarket *LimitOrderMarketSession) GetPrice(arg0 *big.Int) (*big.Int, error) {
	return _LimitOrderMarket.Contract.GetPrice(&_LimitOrderMarket.CallOpts, arg0)
}

// GetPrice is a free data retrieval call binding the contract method 0xe7572230.
//
// Solidity: function getPrice(uint256 ) pure returns(uint256)
func (_LimitOrderMarket *LimitOrderMarketCallerSession) GetPrice(arg0 *big.Int) (*big.Int, error) {
	return _LimitOrderMarket.Contract.GetPrice(&_LimitOrderMarket.CallOpts, arg0)
}

// GetQuoteBuy is a free data retrieval call binding the contract method 0xca6d5811.
//
// Solidity: function getQuoteBuy(uint256 , uint256 , address ) pure returns(uint256, uint256)
func (_LimitOrderMarket *LimitOrderMarketCaller) GetQuoteBuy(opts *bind.CallOpts, arg0 *big.Int, arg1 *big.Int, arg2 common.Address) (*big.Int, *big.Int, error) {
	var out []interface{}
	err := _LimitOrderMarket.contract.Call(opts, &out, "getQuoteBuy", arg0, arg1, arg2)

	if err != nil {
		return *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return out0, out1, err

}

// GetQuoteBuy is a free data retrieval call binding the contract method 0xca6d5811.
//
// Solidity: function getQuoteBuy(uint256 , uint256 , address ) pure returns(uint256, uint256)
func (_LimitOrderMarket *LimitOrderMarketSession) GetQuoteBuy(arg0 *big.Int, arg1 *big.Int, arg2 common.Address) (*big.Int, *big.Int, error) {
	return _LimitOrderMarket.Contract.GetQuoteBuy(&_LimitOrderMarket.CallOpts, arg0, arg1, arg2)
}

// GetQuoteBuy is a free data retrieval call binding the contract method 0xca6d5811.
//
// Solidity: function getQuoteBuy(uint256 , uint256 , address ) pure returns(uint256, uint256)
func (_LimitOrderMarket *LimitOrderMarketCallerSession) GetQuoteBuy(arg0 *big.Int, arg1 *big.Int, arg2 common.Address) (*big.Int, *big.Int, error) {
	return _LimitOrderMarket.Contract.GetQuoteBuy(&_LimitOrderMarket.CallOpts, arg0, arg1, arg2)
}

// GetQuoteSell is a free data retrieval call binding the contract method 0x8b5e8a24.
//
// Solidity: function getQuoteSell(uint256 , uint256 , address ) pure returns(uint256, uint256)
func (_LimitOrderMarket *LimitOrderMarketCaller) GetQuoteSell(opts *bind.CallOpts, arg0 *big.Int, arg1 *big.Int, arg2 common.Address) (*big.Int, *big.Int, error) {
	var out []interface{}
	err := _LimitOrderMarket.contract.Call(opts, &out, "getQuoteSell", arg0, arg1, arg2)

	if err != nil {
		return *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return out0, out1, err

}

// GetQuoteSell is a free data retrieval call binding the contract method 0x8b5e8a24.
//
// Solidity: function getQuoteSell(uint256 , uint256 , address ) pure returns(uint256, uint256)
func (_LimitOrderMarket *LimitOrderMarketSession) GetQuoteSell(arg0 *big.Int, arg1 *big.Int, arg2 common.Address) (*big.Int, *big.Int, error) {
	return _LimitOrderMarket.Contract.GetQuoteSell(&_LimitOrderMarket.CallOpts, arg0, arg1, arg2)
}

// GetQuoteSell is a free data retrieval call binding the contract method 0x8b5e8a24.
//
// Solidity: function getQuoteSell(uint256 , uint256 , address ) pure returns(uint256, uint256)
func (_LimitOrderMarket *LimitOrderMarketCallerSession) GetQuoteSell(arg0 *big.Int, arg1 *big.Int, arg2 common.Address) (*big.Int, *big.Int, error) {
	return _LimitOrderMarket.Contract.GetQuoteSell(&_LimitOrderMarket.CallOpts, arg0, arg1, arg2)
}

// GetSellOrders is a free data retrieval call binding the contract method 0x26d0c772.
//
// Solidity: function getSellOrders(uint8 outcomeId) view returns(bytes32[])
func (_LimitOrderMarket *LimitOrderMarketCaller) GetSellOrders(opts *bind.CallOpts, outcomeId uint8) ([][32]byte, error) {
	var out []interface{}
	err := _LimitOrderMarket.contract.Call(opts, &out, "getSellOrders", outcomeId)

	if err != nil {
		return *new([][32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([][32]byte)).(*[][32]byte)

	return out0, err

}

// GetSellOrders is a free data retrieval call binding the contract method 0x26d0c772.
//
// Solidity: function getSellOrders(uint8 outcomeId) view returns(bytes32[])
func (_LimitOrderMarket *LimitOrderMarketSession) GetSellOrders(outcomeId uint8) ([][32]byte, error) {
	return _LimitOrderMarket.Contract.GetSellOrders(&_LimitOrderMarket.CallOpts, outcomeId)
}

// GetSellOrders is a free data retrieval call binding the contract method 0x26d0c772.
//
// Solidity: function getSellOrders(uint8 outcomeId) view returns(bytes32[])
func (_LimitOrderMarket *LimitOrderMarketCallerSession) GetSellOrders(outcomeId uint8) ([][32]byte, error) {
	return _LimitOrderMarket.Contract.GetSellOrders(&_LimitOrderMarket.CallOpts, outcomeId)
}

// GetSpread is a free data retrieval call binding the contract method 0xf8ff82ea.
//
// Solidity: function getSpread(uint8 outcomeId) view returns(uint256 spread)
func (_LimitOrderMarket *LimitOrderMarketCaller) GetSpread(opts *bind.CallOpts, outcomeId uint8) (*big.Int, error) {
	var out []interface{}
	err := _LimitOrderMarket.contract.Call(opts, &out, "getSpread", outcomeId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetSpread is a free data retrieval call binding the contract method 0xf8ff82ea.
//
// Solidity: function getSpread(uint8 outcomeId) view returns(uint256 spread)
func (_LimitOrderMarket *LimitOrderMarketSession) GetSpread(outcomeId uint8) (*big.Int, error) {
	return _LimitOrderMarket.Contract.GetSpread(&_LimitOrderMarket.CallOpts, outcomeId)
}

// GetSpread is a free data retrieval call binding the contract method 0xf8ff82ea.
//
// Solidity: function getSpread(uint8 outcomeId) view returns(uint256 spread)
func (_LimitOrderMarket *LimitOrderMarketCallerSession) GetSpread(outcomeId uint8) (*big.Int, error) {
	return _LimitOrderMarket.Contract.GetSpread(&_LimitOrderMarket.CallOpts, outcomeId)
}

// GetUserOrders is a free data retrieval call binding the contract method 0x63c69f08.
//
// Solidity: function getUserOrders(address user) view returns(bytes32[])
func (_LimitOrderMarket *LimitOrderMarketCaller) GetUserOrders(opts *bind.CallOpts, user common.Address) ([][32]byte, error) {
	var out []interface{}
	err := _LimitOrderMarket.contract.Call(opts, &out, "getUserOrders", user)

	if err != nil {
		return *new([][32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([][32]byte)).(*[][32]byte)

	return out0, err

}

// GetUserOrders is a free data retrieval call binding the contract method 0x63c69f08.
//
// Solidity: function getUserOrders(address user) view returns(bytes32[])
func (_LimitOrderMarket *LimitOrderMarketSession) GetUserOrders(user common.Address) ([][32]byte, error) {
	return _LimitOrderMarket.Contract.GetUserOrders(&_LimitOrderMarket.CallOpts, user)
}

// GetUserOrders is a free data retrieval call binding the contract method 0x63c69f08.
//
// Solidity: function getUserOrders(address user) view returns(bytes32[])
func (_LimitOrderMarket *LimitOrderMarketCallerSession) GetUserOrders(user common.Address) ([][32]byte, error) {
	return _LimitOrderMarket.Contract.GetUserOrders(&_LimitOrderMarket.CallOpts, user)
}

// HorizonPerks is a free data retrieval call binding the contract method 0xffe02e34.
//
// Solidity: function horizonPerks() view returns(address)
func (_LimitOrderMarket *LimitOrderMarketCaller) HorizonPerks(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _LimitOrderMarket.contract.Call(opts, &out, "horizonPerks")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// HorizonPerks is a free data retrieval call binding the contract method 0xffe02e34.
//
// Solidity: function horizonPerks() view returns(address)
func (_LimitOrderMarket *LimitOrderMarketSession) HorizonPerks() (common.Address, error) {
	return _LimitOrderMarket.Contract.HorizonPerks(&_LimitOrderMarket.CallOpts)
}

// HorizonPerks is a free data retrieval call binding the contract method 0xffe02e34.
//
// Solidity: function horizonPerks() view returns(address)
func (_LimitOrderMarket *LimitOrderMarketCallerSession) HorizonPerks() (common.Address, error) {
	return _LimitOrderMarket.Contract.HorizonPerks(&_LimitOrderMarket.CallOpts)
}

// MarketId is a free data retrieval call binding the contract method 0x6ed71ede.
//
// Solidity: function marketId() view returns(uint256)
func (_LimitOrderMarket *LimitOrderMarketCaller) MarketId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _LimitOrderMarket.contract.Call(opts, &out, "marketId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MarketId is a free data retrieval call binding the contract method 0x6ed71ede.
//
// Solidity: function marketId() view returns(uint256)
func (_LimitOrderMarket *LimitOrderMarketSession) MarketId() (*big.Int, error) {
	return _LimitOrderMarket.Contract.MarketId(&_LimitOrderMarket.CallOpts)
}

// MarketId is a free data retrieval call binding the contract method 0x6ed71ede.
//
// Solidity: function marketId() view returns(uint256)
func (_LimitOrderMarket *LimitOrderMarketCallerSession) MarketId() (*big.Int, error) {
	return _LimitOrderMarket.Contract.MarketId(&_LimitOrderMarket.CallOpts)
}

// MarketType is a free data retrieval call binding the contract method 0x2dd48909.
//
// Solidity: function marketType() view returns(uint8)
func (_LimitOrderMarket *LimitOrderMarketCaller) MarketType(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _LimitOrderMarket.contract.Call(opts, &out, "marketType")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// MarketType is a free data retrieval call binding the contract method 0x2dd48909.
//
// Solidity: function marketType() view returns(uint8)
func (_LimitOrderMarket *LimitOrderMarketSession) MarketType() (uint8, error) {
	return _LimitOrderMarket.Contract.MarketType(&_LimitOrderMarket.CallOpts)
}

// MarketType is a free data retrieval call binding the contract method 0x2dd48909.
//
// Solidity: function marketType() view returns(uint8)
func (_LimitOrderMarket *LimitOrderMarketCallerSession) MarketType() (uint8, error) {
	return _LimitOrderMarket.Contract.MarketType(&_LimitOrderMarket.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_LimitOrderMarket *LimitOrderMarketCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _LimitOrderMarket.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_LimitOrderMarket *LimitOrderMarketSession) Name() (string, error) {
	return _LimitOrderMarket.Contract.Name(&_LimitOrderMarket.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_LimitOrderMarket *LimitOrderMarketCallerSession) Name() (string, error) {
	return _LimitOrderMarket.Contract.Name(&_LimitOrderMarket.CallOpts)
}

// OnERC1155BatchReceived is a free data retrieval call binding the contract method 0xbc197c81.
//
// Solidity: function onERC1155BatchReceived(address , address , uint256[] , uint256[] , bytes ) pure returns(bytes4)
func (_LimitOrderMarket *LimitOrderMarketCaller) OnERC1155BatchReceived(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address, arg2 []*big.Int, arg3 []*big.Int, arg4 []byte) ([4]byte, error) {
	var out []interface{}
	err := _LimitOrderMarket.contract.Call(opts, &out, "onERC1155BatchReceived", arg0, arg1, arg2, arg3, arg4)

	if err != nil {
		return *new([4]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([4]byte)).(*[4]byte)

	return out0, err

}

// OnERC1155BatchReceived is a free data retrieval call binding the contract method 0xbc197c81.
//
// Solidity: function onERC1155BatchReceived(address , address , uint256[] , uint256[] , bytes ) pure returns(bytes4)
func (_LimitOrderMarket *LimitOrderMarketSession) OnERC1155BatchReceived(arg0 common.Address, arg1 common.Address, arg2 []*big.Int, arg3 []*big.Int, arg4 []byte) ([4]byte, error) {
	return _LimitOrderMarket.Contract.OnERC1155BatchReceived(&_LimitOrderMarket.CallOpts, arg0, arg1, arg2, arg3, arg4)
}

// OnERC1155BatchReceived is a free data retrieval call binding the contract method 0xbc197c81.
//
// Solidity: function onERC1155BatchReceived(address , address , uint256[] , uint256[] , bytes ) pure returns(bytes4)
func (_LimitOrderMarket *LimitOrderMarketCallerSession) OnERC1155BatchReceived(arg0 common.Address, arg1 common.Address, arg2 []*big.Int, arg3 []*big.Int, arg4 []byte) ([4]byte, error) {
	return _LimitOrderMarket.Contract.OnERC1155BatchReceived(&_LimitOrderMarket.CallOpts, arg0, arg1, arg2, arg3, arg4)
}

// OnERC1155Received is a free data retrieval call binding the contract method 0xf23a6e61.
//
// Solidity: function onERC1155Received(address , address , uint256 , uint256 , bytes ) pure returns(bytes4)
func (_LimitOrderMarket *LimitOrderMarketCaller) OnERC1155Received(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 *big.Int, arg4 []byte) ([4]byte, error) {
	var out []interface{}
	err := _LimitOrderMarket.contract.Call(opts, &out, "onERC1155Received", arg0, arg1, arg2, arg3, arg4)

	if err != nil {
		return *new([4]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([4]byte)).(*[4]byte)

	return out0, err

}

// OnERC1155Received is a free data retrieval call binding the contract method 0xf23a6e61.
//
// Solidity: function onERC1155Received(address , address , uint256 , uint256 , bytes ) pure returns(bytes4)
func (_LimitOrderMarket *LimitOrderMarketSession) OnERC1155Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 *big.Int, arg4 []byte) ([4]byte, error) {
	return _LimitOrderMarket.Contract.OnERC1155Received(&_LimitOrderMarket.CallOpts, arg0, arg1, arg2, arg3, arg4)
}

// OnERC1155Received is a free data retrieval call binding the contract method 0xf23a6e61.
//
// Solidity: function onERC1155Received(address , address , uint256 , uint256 , bytes ) pure returns(bytes4)
func (_LimitOrderMarket *LimitOrderMarketCallerSession) OnERC1155Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 *big.Int, arg4 []byte) ([4]byte, error) {
	return _LimitOrderMarket.Contract.OnERC1155Received(&_LimitOrderMarket.CallOpts, arg0, arg1, arg2, arg3, arg4)
}

// Orders is a free data retrieval call binding the contract method 0x9c3f1e90.
//
// Solidity: function orders(bytes32 ) view returns(address trader, uint8 outcomeId, uint256 price, uint256 amount, uint256 filled, uint256 totalCost, uint256 timestamp, bool isBuy, uint8 orderType, bool isActive)
func (_LimitOrderMarket *LimitOrderMarketCaller) Orders(opts *bind.CallOpts, arg0 [32]byte) (struct {
	Trader    common.Address
	OutcomeId uint8
	Price     *big.Int
	Amount    *big.Int
	Filled    *big.Int
	TotalCost *big.Int
	Timestamp *big.Int
	IsBuy     bool
	OrderType uint8
	IsActive  bool
}, error) {
	var out []interface{}
	err := _LimitOrderMarket.contract.Call(opts, &out, "orders", arg0)

	outstruct := new(struct {
		Trader    common.Address
		OutcomeId uint8
		Price     *big.Int
		Amount    *big.Int
		Filled    *big.Int
		TotalCost *big.Int
		Timestamp *big.Int
		IsBuy     bool
		OrderType uint8
		IsActive  bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Trader = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.OutcomeId = *abi.ConvertType(out[1], new(uint8)).(*uint8)
	outstruct.Price = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.Amount = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.Filled = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.TotalCost = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)
	outstruct.Timestamp = *abi.ConvertType(out[6], new(*big.Int)).(**big.Int)
	outstruct.IsBuy = *abi.ConvertType(out[7], new(bool)).(*bool)
	outstruct.OrderType = *abi.ConvertType(out[8], new(uint8)).(*uint8)
	outstruct.IsActive = *abi.ConvertType(out[9], new(bool)).(*bool)

	return *outstruct, err

}

// Orders is a free data retrieval call binding the contract method 0x9c3f1e90.
//
// Solidity: function orders(bytes32 ) view returns(address trader, uint8 outcomeId, uint256 price, uint256 amount, uint256 filled, uint256 totalCost, uint256 timestamp, bool isBuy, uint8 orderType, bool isActive)
func (_LimitOrderMarket *LimitOrderMarketSession) Orders(arg0 [32]byte) (struct {
	Trader    common.Address
	OutcomeId uint8
	Price     *big.Int
	Amount    *big.Int
	Filled    *big.Int
	TotalCost *big.Int
	Timestamp *big.Int
	IsBuy     bool
	OrderType uint8
	IsActive  bool
}, error) {
	return _LimitOrderMarket.Contract.Orders(&_LimitOrderMarket.CallOpts, arg0)
}

// Orders is a free data retrieval call binding the contract method 0x9c3f1e90.
//
// Solidity: function orders(bytes32 ) view returns(address trader, uint8 outcomeId, uint256 price, uint256 amount, uint256 filled, uint256 totalCost, uint256 timestamp, bool isBuy, uint8 orderType, bool isActive)
func (_LimitOrderMarket *LimitOrderMarketCallerSession) Orders(arg0 [32]byte) (struct {
	Trader    common.Address
	OutcomeId uint8
	Price     *big.Int
	Amount    *big.Int
	Filled    *big.Int
	TotalCost *big.Int
	Timestamp *big.Int
	IsBuy     bool
	OrderType uint8
	IsActive  bool
}, error) {
	return _LimitOrderMarket.Contract.Orders(&_LimitOrderMarket.CallOpts, arg0)
}

// OutcomeCount is a free data retrieval call binding the contract method 0xd300cb31.
//
// Solidity: function outcomeCount() view returns(uint256)
func (_LimitOrderMarket *LimitOrderMarketCaller) OutcomeCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _LimitOrderMarket.contract.Call(opts, &out, "outcomeCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// OutcomeCount is a free data retrieval call binding the contract method 0xd300cb31.
//
// Solidity: function outcomeCount() view returns(uint256)
func (_LimitOrderMarket *LimitOrderMarketSession) OutcomeCount() (*big.Int, error) {
	return _LimitOrderMarket.Contract.OutcomeCount(&_LimitOrderMarket.CallOpts)
}

// OutcomeCount is a free data retrieval call binding the contract method 0xd300cb31.
//
// Solidity: function outcomeCount() view returns(uint256)
func (_LimitOrderMarket *LimitOrderMarketCallerSession) OutcomeCount() (*big.Int, error) {
	return _LimitOrderMarket.Contract.OutcomeCount(&_LimitOrderMarket.CallOpts)
}

// OutcomeToken is a free data retrieval call binding the contract method 0xa998d6d8.
//
// Solidity: function outcomeToken() view returns(address)
func (_LimitOrderMarket *LimitOrderMarketCaller) OutcomeToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _LimitOrderMarket.contract.Call(opts, &out, "outcomeToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OutcomeToken is a free data retrieval call binding the contract method 0xa998d6d8.
//
// Solidity: function outcomeToken() view returns(address)
func (_LimitOrderMarket *LimitOrderMarketSession) OutcomeToken() (common.Address, error) {
	return _LimitOrderMarket.Contract.OutcomeToken(&_LimitOrderMarket.CallOpts)
}

// OutcomeToken is a free data retrieval call binding the contract method 0xa998d6d8.
//
// Solidity: function outcomeToken() view returns(address)
func (_LimitOrderMarket *LimitOrderMarketCallerSession) OutcomeToken() (common.Address, error) {
	return _LimitOrderMarket.Contract.OutcomeToken(&_LimitOrderMarket.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_LimitOrderMarket *LimitOrderMarketCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _LimitOrderMarket.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_LimitOrderMarket *LimitOrderMarketSession) Paused() (bool, error) {
	return _LimitOrderMarket.Contract.Paused(&_LimitOrderMarket.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_LimitOrderMarket *LimitOrderMarketCallerSession) Paused() (bool, error) {
	return _LimitOrderMarket.Contract.Paused(&_LimitOrderMarket.CallOpts)
}

// RemoveLiquidity is a free data retrieval call binding the contract method 0x9c8f9f23.
//
// Solidity: function removeLiquidity(uint256 ) pure returns(uint256)
func (_LimitOrderMarket *LimitOrderMarketCaller) RemoveLiquidity(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _LimitOrderMarket.contract.Call(opts, &out, "removeLiquidity", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RemoveLiquidity is a free data retrieval call binding the contract method 0x9c8f9f23.
//
// Solidity: function removeLiquidity(uint256 ) pure returns(uint256)
func (_LimitOrderMarket *LimitOrderMarketSession) RemoveLiquidity(arg0 *big.Int) (*big.Int, error) {
	return _LimitOrderMarket.Contract.RemoveLiquidity(&_LimitOrderMarket.CallOpts, arg0)
}

// RemoveLiquidity is a free data retrieval call binding the contract method 0x9c8f9f23.
//
// Solidity: function removeLiquidity(uint256 ) pure returns(uint256)
func (_LimitOrderMarket *LimitOrderMarketCallerSession) RemoveLiquidity(arg0 *big.Int) (*big.Int, error) {
	return _LimitOrderMarket.Contract.RemoveLiquidity(&_LimitOrderMarket.CallOpts, arg0)
}

// Sell is a free data retrieval call binding the contract method 0xd3c9727c.
//
// Solidity: function sell(uint256 , uint256 , uint256 ) pure returns(uint256)
func (_LimitOrderMarket *LimitOrderMarketCaller) Sell(opts *bind.CallOpts, arg0 *big.Int, arg1 *big.Int, arg2 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _LimitOrderMarket.contract.Call(opts, &out, "sell", arg0, arg1, arg2)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Sell is a free data retrieval call binding the contract method 0xd3c9727c.
//
// Solidity: function sell(uint256 , uint256 , uint256 ) pure returns(uint256)
func (_LimitOrderMarket *LimitOrderMarketSession) Sell(arg0 *big.Int, arg1 *big.Int, arg2 *big.Int) (*big.Int, error) {
	return _LimitOrderMarket.Contract.Sell(&_LimitOrderMarket.CallOpts, arg0, arg1, arg2)
}

// Sell is a free data retrieval call binding the contract method 0xd3c9727c.
//
// Solidity: function sell(uint256 , uint256 , uint256 ) pure returns(uint256)
func (_LimitOrderMarket *LimitOrderMarketCallerSession) Sell(arg0 *big.Int, arg1 *big.Int, arg2 *big.Int) (*big.Int, error) {
	return _LimitOrderMarket.Contract.Sell(&_LimitOrderMarket.CallOpts, arg0, arg1, arg2)
}

// SellOrdersByOutcome is a free data retrieval call binding the contract method 0x4ef91484.
//
// Solidity: function sellOrdersByOutcome(uint8 , uint256 ) view returns(bytes32)
func (_LimitOrderMarket *LimitOrderMarketCaller) SellOrdersByOutcome(opts *bind.CallOpts, arg0 uint8, arg1 *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _LimitOrderMarket.contract.Call(opts, &out, "sellOrdersByOutcome", arg0, arg1)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// SellOrdersByOutcome is a free data retrieval call binding the contract method 0x4ef91484.
//
// Solidity: function sellOrdersByOutcome(uint8 , uint256 ) view returns(bytes32)
func (_LimitOrderMarket *LimitOrderMarketSession) SellOrdersByOutcome(arg0 uint8, arg1 *big.Int) ([32]byte, error) {
	return _LimitOrderMarket.Contract.SellOrdersByOutcome(&_LimitOrderMarket.CallOpts, arg0, arg1)
}

// SellOrdersByOutcome is a free data retrieval call binding the contract method 0x4ef91484.
//
// Solidity: function sellOrdersByOutcome(uint8 , uint256 ) view returns(bytes32)
func (_LimitOrderMarket *LimitOrderMarketCallerSession) SellOrdersByOutcome(arg0 uint8, arg1 *big.Int) ([32]byte, error) {
	return _LimitOrderMarket.Contract.SellOrdersByOutcome(&_LimitOrderMarket.CallOpts, arg0, arg1)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) pure returns(bool)
func (_LimitOrderMarket *LimitOrderMarketCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _LimitOrderMarket.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) pure returns(bool)
func (_LimitOrderMarket *LimitOrderMarketSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _LimitOrderMarket.Contract.SupportsInterface(&_LimitOrderMarket.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) pure returns(bool)
func (_LimitOrderMarket *LimitOrderMarketCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _LimitOrderMarket.Contract.SupportsInterface(&_LimitOrderMarket.CallOpts, interfaceId)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_LimitOrderMarket *LimitOrderMarketCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _LimitOrderMarket.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_LimitOrderMarket *LimitOrderMarketSession) Symbol() (string, error) {
	return _LimitOrderMarket.Contract.Symbol(&_LimitOrderMarket.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_LimitOrderMarket *LimitOrderMarketCallerSession) Symbol() (string, error) {
	return _LimitOrderMarket.Contract.Symbol(&_LimitOrderMarket.CallOpts)
}

// TotalCollateral is a free data retrieval call binding the contract method 0x4ac8eb5f.
//
// Solidity: function totalCollateral() view returns(uint256)
func (_LimitOrderMarket *LimitOrderMarketCaller) TotalCollateral(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _LimitOrderMarket.contract.Call(opts, &out, "totalCollateral")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalCollateral is a free data retrieval call binding the contract method 0x4ac8eb5f.
//
// Solidity: function totalCollateral() view returns(uint256)
func (_LimitOrderMarket *LimitOrderMarketSession) TotalCollateral() (*big.Int, error) {
	return _LimitOrderMarket.Contract.TotalCollateral(&_LimitOrderMarket.CallOpts)
}

// TotalCollateral is a free data retrieval call binding the contract method 0x4ac8eb5f.
//
// Solidity: function totalCollateral() view returns(uint256)
func (_LimitOrderMarket *LimitOrderMarketCallerSession) TotalCollateral() (*big.Int, error) {
	return _LimitOrderMarket.Contract.TotalCollateral(&_LimitOrderMarket.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_LimitOrderMarket *LimitOrderMarketCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _LimitOrderMarket.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_LimitOrderMarket *LimitOrderMarketSession) TotalSupply() (*big.Int, error) {
	return _LimitOrderMarket.Contract.TotalSupply(&_LimitOrderMarket.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_LimitOrderMarket *LimitOrderMarketCallerSession) TotalSupply() (*big.Int, error) {
	return _LimitOrderMarket.Contract.TotalSupply(&_LimitOrderMarket.CallOpts)
}

// UserOrders is a free data retrieval call binding the contract method 0x856652e9.
//
// Solidity: function userOrders(address , uint256 ) view returns(bytes32)
func (_LimitOrderMarket *LimitOrderMarketCaller) UserOrders(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _LimitOrderMarket.contract.Call(opts, &out, "userOrders", arg0, arg1)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// UserOrders is a free data retrieval call binding the contract method 0x856652e9.
//
// Solidity: function userOrders(address , uint256 ) view returns(bytes32)
func (_LimitOrderMarket *LimitOrderMarketSession) UserOrders(arg0 common.Address, arg1 *big.Int) ([32]byte, error) {
	return _LimitOrderMarket.Contract.UserOrders(&_LimitOrderMarket.CallOpts, arg0, arg1)
}

// UserOrders is a free data retrieval call binding the contract method 0x856652e9.
//
// Solidity: function userOrders(address , uint256 ) view returns(bytes32)
func (_LimitOrderMarket *LimitOrderMarketCallerSession) UserOrders(arg0 common.Address, arg1 *big.Int) ([32]byte, error) {
	return _LimitOrderMarket.Contract.UserOrders(&_LimitOrderMarket.CallOpts, arg0, arg1)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_LimitOrderMarket *LimitOrderMarketTransactor) Approve(opts *bind.TransactOpts, spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _LimitOrderMarket.contract.Transact(opts, "approve", spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_LimitOrderMarket *LimitOrderMarketSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _LimitOrderMarket.Contract.Approve(&_LimitOrderMarket.TransactOpts, spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_LimitOrderMarket *LimitOrderMarketTransactorSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _LimitOrderMarket.Contract.Approve(&_LimitOrderMarket.TransactOpts, spender, value)
}

// CancelOrder is a paid mutator transaction binding the contract method 0x7489ec23.
//
// Solidity: function cancelOrder(bytes32 orderId) returns()
func (_LimitOrderMarket *LimitOrderMarketTransactor) CancelOrder(opts *bind.TransactOpts, orderId [32]byte) (*types.Transaction, error) {
	return _LimitOrderMarket.contract.Transact(opts, "cancelOrder", orderId)
}

// CancelOrder is a paid mutator transaction binding the contract method 0x7489ec23.
//
// Solidity: function cancelOrder(bytes32 orderId) returns()
func (_LimitOrderMarket *LimitOrderMarketSession) CancelOrder(orderId [32]byte) (*types.Transaction, error) {
	return _LimitOrderMarket.Contract.CancelOrder(&_LimitOrderMarket.TransactOpts, orderId)
}

// CancelOrder is a paid mutator transaction binding the contract method 0x7489ec23.
//
// Solidity: function cancelOrder(bytes32 orderId) returns()
func (_LimitOrderMarket *LimitOrderMarketTransactorSession) CancelOrder(orderId [32]byte) (*types.Transaction, error) {
	return _LimitOrderMarket.Contract.CancelOrder(&_LimitOrderMarket.TransactOpts, orderId)
}

// FundRedemptions is a paid mutator transaction binding the contract method 0x281155ba.
//
// Solidity: function fundRedemptions() returns()
func (_LimitOrderMarket *LimitOrderMarketTransactor) FundRedemptions(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LimitOrderMarket.contract.Transact(opts, "fundRedemptions")
}

// FundRedemptions is a paid mutator transaction binding the contract method 0x281155ba.
//
// Solidity: function fundRedemptions() returns()
func (_LimitOrderMarket *LimitOrderMarketSession) FundRedemptions() (*types.Transaction, error) {
	return _LimitOrderMarket.Contract.FundRedemptions(&_LimitOrderMarket.TransactOpts)
}

// FundRedemptions is a paid mutator transaction binding the contract method 0x281155ba.
//
// Solidity: function fundRedemptions() returns()
func (_LimitOrderMarket *LimitOrderMarketTransactorSession) FundRedemptions() (*types.Transaction, error) {
	return _LimitOrderMarket.Contract.FundRedemptions(&_LimitOrderMarket.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_LimitOrderMarket *LimitOrderMarketTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LimitOrderMarket.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_LimitOrderMarket *LimitOrderMarketSession) Pause() (*types.Transaction, error) {
	return _LimitOrderMarket.Contract.Pause(&_LimitOrderMarket.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_LimitOrderMarket *LimitOrderMarketTransactorSession) Pause() (*types.Transaction, error) {
	return _LimitOrderMarket.Contract.Pause(&_LimitOrderMarket.TransactOpts)
}

// PlaceMarketOrder is a paid mutator transaction binding the contract method 0xd775205a.
//
// Solidity: function placeMarketOrder(uint8 outcomeId, bool isBuy, uint256 amount, uint256 maxSlippage) returns(uint256 filledAmount)
func (_LimitOrderMarket *LimitOrderMarketTransactor) PlaceMarketOrder(opts *bind.TransactOpts, outcomeId uint8, isBuy bool, amount *big.Int, maxSlippage *big.Int) (*types.Transaction, error) {
	return _LimitOrderMarket.contract.Transact(opts, "placeMarketOrder", outcomeId, isBuy, amount, maxSlippage)
}

// PlaceMarketOrder is a paid mutator transaction binding the contract method 0xd775205a.
//
// Solidity: function placeMarketOrder(uint8 outcomeId, bool isBuy, uint256 amount, uint256 maxSlippage) returns(uint256 filledAmount)
func (_LimitOrderMarket *LimitOrderMarketSession) PlaceMarketOrder(outcomeId uint8, isBuy bool, amount *big.Int, maxSlippage *big.Int) (*types.Transaction, error) {
	return _LimitOrderMarket.Contract.PlaceMarketOrder(&_LimitOrderMarket.TransactOpts, outcomeId, isBuy, amount, maxSlippage)
}

// PlaceMarketOrder is a paid mutator transaction binding the contract method 0xd775205a.
//
// Solidity: function placeMarketOrder(uint8 outcomeId, bool isBuy, uint256 amount, uint256 maxSlippage) returns(uint256 filledAmount)
func (_LimitOrderMarket *LimitOrderMarketTransactorSession) PlaceMarketOrder(outcomeId uint8, isBuy bool, amount *big.Int, maxSlippage *big.Int) (*types.Transaction, error) {
	return _LimitOrderMarket.Contract.PlaceMarketOrder(&_LimitOrderMarket.TransactOpts, outcomeId, isBuy, amount, maxSlippage)
}

// PlaceOrder is a paid mutator transaction binding the contract method 0x61a87b87.
//
// Solidity: function placeOrder(uint8 outcomeId, bool isBuy, uint256 price, uint256 amount, uint8 orderType) returns(bytes32 orderId)
func (_LimitOrderMarket *LimitOrderMarketTransactor) PlaceOrder(opts *bind.TransactOpts, outcomeId uint8, isBuy bool, price *big.Int, amount *big.Int, orderType uint8) (*types.Transaction, error) {
	return _LimitOrderMarket.contract.Transact(opts, "placeOrder", outcomeId, isBuy, price, amount, orderType)
}

// PlaceOrder is a paid mutator transaction binding the contract method 0x61a87b87.
//
// Solidity: function placeOrder(uint8 outcomeId, bool isBuy, uint256 price, uint256 amount, uint8 orderType) returns(bytes32 orderId)
func (_LimitOrderMarket *LimitOrderMarketSession) PlaceOrder(outcomeId uint8, isBuy bool, price *big.Int, amount *big.Int, orderType uint8) (*types.Transaction, error) {
	return _LimitOrderMarket.Contract.PlaceOrder(&_LimitOrderMarket.TransactOpts, outcomeId, isBuy, price, amount, orderType)
}

// PlaceOrder is a paid mutator transaction binding the contract method 0x61a87b87.
//
// Solidity: function placeOrder(uint8 outcomeId, bool isBuy, uint256 price, uint256 amount, uint8 orderType) returns(bytes32 orderId)
func (_LimitOrderMarket *LimitOrderMarketTransactorSession) PlaceOrder(outcomeId uint8, isBuy bool, price *big.Int, amount *big.Int, orderType uint8) (*types.Transaction, error) {
	return _LimitOrderMarket.Contract.PlaceOrder(&_LimitOrderMarket.TransactOpts, outcomeId, isBuy, price, amount, orderType)
}

// SetAdmin is a paid mutator transaction binding the contract method 0x704b6c02.
//
// Solidity: function setAdmin(address newAdmin) returns()
func (_LimitOrderMarket *LimitOrderMarketTransactor) SetAdmin(opts *bind.TransactOpts, newAdmin common.Address) (*types.Transaction, error) {
	return _LimitOrderMarket.contract.Transact(opts, "setAdmin", newAdmin)
}

// SetAdmin is a paid mutator transaction binding the contract method 0x704b6c02.
//
// Solidity: function setAdmin(address newAdmin) returns()
func (_LimitOrderMarket *LimitOrderMarketSession) SetAdmin(newAdmin common.Address) (*types.Transaction, error) {
	return _LimitOrderMarket.Contract.SetAdmin(&_LimitOrderMarket.TransactOpts, newAdmin)
}

// SetAdmin is a paid mutator transaction binding the contract method 0x704b6c02.
//
// Solidity: function setAdmin(address newAdmin) returns()
func (_LimitOrderMarket *LimitOrderMarketTransactorSession) SetAdmin(newAdmin common.Address) (*types.Transaction, error) {
	return _LimitOrderMarket.Contract.SetAdmin(&_LimitOrderMarket.TransactOpts, newAdmin)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_LimitOrderMarket *LimitOrderMarketTransactor) Transfer(opts *bind.TransactOpts, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _LimitOrderMarket.contract.Transact(opts, "transfer", to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_LimitOrderMarket *LimitOrderMarketSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _LimitOrderMarket.Contract.Transfer(&_LimitOrderMarket.TransactOpts, to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_LimitOrderMarket *LimitOrderMarketTransactorSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _LimitOrderMarket.Contract.Transfer(&_LimitOrderMarket.TransactOpts, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_LimitOrderMarket *LimitOrderMarketTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _LimitOrderMarket.contract.Transact(opts, "transferFrom", from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_LimitOrderMarket *LimitOrderMarketSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _LimitOrderMarket.Contract.TransferFrom(&_LimitOrderMarket.TransactOpts, from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_LimitOrderMarket *LimitOrderMarketTransactorSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _LimitOrderMarket.Contract.TransferFrom(&_LimitOrderMarket.TransactOpts, from, to, value)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_LimitOrderMarket *LimitOrderMarketTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LimitOrderMarket.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_LimitOrderMarket *LimitOrderMarketSession) Unpause() (*types.Transaction, error) {
	return _LimitOrderMarket.Contract.Unpause(&_LimitOrderMarket.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_LimitOrderMarket *LimitOrderMarketTransactorSession) Unpause() (*types.Transaction, error) {
	return _LimitOrderMarket.Contract.Unpause(&_LimitOrderMarket.TransactOpts)
}

// LimitOrderMarketApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the LimitOrderMarket contract.
type LimitOrderMarketApprovalIterator struct {
	Event *LimitOrderMarketApproval // Event containing the contract specifics and raw log

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
func (it *LimitOrderMarketApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LimitOrderMarketApproval)
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
		it.Event = new(LimitOrderMarketApproval)
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
func (it *LimitOrderMarketApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LimitOrderMarketApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LimitOrderMarketApproval represents a Approval event raised by the LimitOrderMarket contract.
type LimitOrderMarketApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_LimitOrderMarket *LimitOrderMarketFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*LimitOrderMarketApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _LimitOrderMarket.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &LimitOrderMarketApprovalIterator{contract: _LimitOrderMarket.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_LimitOrderMarket *LimitOrderMarketFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *LimitOrderMarketApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _LimitOrderMarket.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LimitOrderMarketApproval)
				if err := _LimitOrderMarket.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_LimitOrderMarket *LimitOrderMarketFilterer) ParseApproval(log types.Log) (*LimitOrderMarketApproval, error) {
	event := new(LimitOrderMarketApproval)
	if err := _LimitOrderMarket.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LimitOrderMarketLiquidityChangedIterator is returned from FilterLiquidityChanged and is used to iterate over the raw logs and unpacked data for LiquidityChanged events raised by the LimitOrderMarket contract.
type LimitOrderMarketLiquidityChangedIterator struct {
	Event *LimitOrderMarketLiquidityChanged // Event containing the contract specifics and raw log

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
func (it *LimitOrderMarketLiquidityChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LimitOrderMarketLiquidityChanged)
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
		it.Event = new(LimitOrderMarketLiquidityChanged)
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
func (it *LimitOrderMarketLiquidityChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LimitOrderMarketLiquidityChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LimitOrderMarketLiquidityChanged represents a LiquidityChanged event raised by the LimitOrderMarket contract.
type LimitOrderMarketLiquidityChanged struct {
	Provider   common.Address
	Amount     *big.Int
	IsAddition bool
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterLiquidityChanged is a free log retrieval operation binding the contract event 0xb029a6414a0c6d2e4fa2e5287326aa8a8c7191f9f5ced9799754a380471458d4.
//
// Solidity: event LiquidityChanged(address indexed provider, uint256 amount, bool isAddition)
func (_LimitOrderMarket *LimitOrderMarketFilterer) FilterLiquidityChanged(opts *bind.FilterOpts, provider []common.Address) (*LimitOrderMarketLiquidityChangedIterator, error) {

	var providerRule []interface{}
	for _, providerItem := range provider {
		providerRule = append(providerRule, providerItem)
	}

	logs, sub, err := _LimitOrderMarket.contract.FilterLogs(opts, "LiquidityChanged", providerRule)
	if err != nil {
		return nil, err
	}
	return &LimitOrderMarketLiquidityChangedIterator{contract: _LimitOrderMarket.contract, event: "LiquidityChanged", logs: logs, sub: sub}, nil
}

// WatchLiquidityChanged is a free log subscription operation binding the contract event 0xb029a6414a0c6d2e4fa2e5287326aa8a8c7191f9f5ced9799754a380471458d4.
//
// Solidity: event LiquidityChanged(address indexed provider, uint256 amount, bool isAddition)
func (_LimitOrderMarket *LimitOrderMarketFilterer) WatchLiquidityChanged(opts *bind.WatchOpts, sink chan<- *LimitOrderMarketLiquidityChanged, provider []common.Address) (event.Subscription, error) {

	var providerRule []interface{}
	for _, providerItem := range provider {
		providerRule = append(providerRule, providerItem)
	}

	logs, sub, err := _LimitOrderMarket.contract.WatchLogs(opts, "LiquidityChanged", providerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LimitOrderMarketLiquidityChanged)
				if err := _LimitOrderMarket.contract.UnpackLog(event, "LiquidityChanged", log); err != nil {
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
func (_LimitOrderMarket *LimitOrderMarketFilterer) ParseLiquidityChanged(log types.Log) (*LimitOrderMarketLiquidityChanged, error) {
	event := new(LimitOrderMarketLiquidityChanged)
	if err := _LimitOrderMarket.contract.UnpackLog(event, "LiquidityChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LimitOrderMarketOrderCancelledIterator is returned from FilterOrderCancelled and is used to iterate over the raw logs and unpacked data for OrderCancelled events raised by the LimitOrderMarket contract.
type LimitOrderMarketOrderCancelledIterator struct {
	Event *LimitOrderMarketOrderCancelled // Event containing the contract specifics and raw log

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
func (it *LimitOrderMarketOrderCancelledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LimitOrderMarketOrderCancelled)
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
		it.Event = new(LimitOrderMarketOrderCancelled)
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
func (it *LimitOrderMarketOrderCancelledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LimitOrderMarketOrderCancelledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LimitOrderMarketOrderCancelled represents a OrderCancelled event raised by the LimitOrderMarket contract.
type LimitOrderMarketOrderCancelled struct {
	OrderId         [32]byte
	Trader          common.Address
	AmountRemaining *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterOrderCancelled is a free log retrieval operation binding the contract event 0xb2705df32ac67fc3101f496cd7036bf59074a603544d97d73650b6f09744986a.
//
// Solidity: event OrderCancelled(bytes32 indexed orderId, address indexed trader, uint256 amountRemaining)
func (_LimitOrderMarket *LimitOrderMarketFilterer) FilterOrderCancelled(opts *bind.FilterOpts, orderId [][32]byte, trader []common.Address) (*LimitOrderMarketOrderCancelledIterator, error) {

	var orderIdRule []interface{}
	for _, orderIdItem := range orderId {
		orderIdRule = append(orderIdRule, orderIdItem)
	}
	var traderRule []interface{}
	for _, traderItem := range trader {
		traderRule = append(traderRule, traderItem)
	}

	logs, sub, err := _LimitOrderMarket.contract.FilterLogs(opts, "OrderCancelled", orderIdRule, traderRule)
	if err != nil {
		return nil, err
	}
	return &LimitOrderMarketOrderCancelledIterator{contract: _LimitOrderMarket.contract, event: "OrderCancelled", logs: logs, sub: sub}, nil
}

// WatchOrderCancelled is a free log subscription operation binding the contract event 0xb2705df32ac67fc3101f496cd7036bf59074a603544d97d73650b6f09744986a.
//
// Solidity: event OrderCancelled(bytes32 indexed orderId, address indexed trader, uint256 amountRemaining)
func (_LimitOrderMarket *LimitOrderMarketFilterer) WatchOrderCancelled(opts *bind.WatchOpts, sink chan<- *LimitOrderMarketOrderCancelled, orderId [][32]byte, trader []common.Address) (event.Subscription, error) {

	var orderIdRule []interface{}
	for _, orderIdItem := range orderId {
		orderIdRule = append(orderIdRule, orderIdItem)
	}
	var traderRule []interface{}
	for _, traderItem := range trader {
		traderRule = append(traderRule, traderItem)
	}

	logs, sub, err := _LimitOrderMarket.contract.WatchLogs(opts, "OrderCancelled", orderIdRule, traderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LimitOrderMarketOrderCancelled)
				if err := _LimitOrderMarket.contract.UnpackLog(event, "OrderCancelled", log); err != nil {
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

// ParseOrderCancelled is a log parse operation binding the contract event 0xb2705df32ac67fc3101f496cd7036bf59074a603544d97d73650b6f09744986a.
//
// Solidity: event OrderCancelled(bytes32 indexed orderId, address indexed trader, uint256 amountRemaining)
func (_LimitOrderMarket *LimitOrderMarketFilterer) ParseOrderCancelled(log types.Log) (*LimitOrderMarketOrderCancelled, error) {
	event := new(LimitOrderMarketOrderCancelled)
	if err := _LimitOrderMarket.contract.UnpackLog(event, "OrderCancelled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LimitOrderMarketOrderMatchedIterator is returned from FilterOrderMatched and is used to iterate over the raw logs and unpacked data for OrderMatched events raised by the LimitOrderMarket contract.
type LimitOrderMarketOrderMatchedIterator struct {
	Event *LimitOrderMarketOrderMatched // Event containing the contract specifics and raw log

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
func (it *LimitOrderMarketOrderMatchedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LimitOrderMarketOrderMatched)
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
		it.Event = new(LimitOrderMarketOrderMatched)
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
func (it *LimitOrderMarketOrderMatchedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LimitOrderMarketOrderMatchedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LimitOrderMarketOrderMatched represents a OrderMatched event raised by the LimitOrderMarket contract.
type LimitOrderMarketOrderMatched struct {
	BuyOrderId  [32]byte
	SellOrderId [32]byte
	OutcomeId   uint8
	Price       *big.Int
	Amount      *big.Int
	Buyer       common.Address
	Seller      common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterOrderMatched is a free log retrieval operation binding the contract event 0xa674596fa044bf929b1aa6b81415a6a1534532f20319761ad36f7547e243b34a.
//
// Solidity: event OrderMatched(bytes32 indexed buyOrderId, bytes32 indexed sellOrderId, uint8 indexed outcomeId, uint256 price, uint256 amount, address buyer, address seller)
func (_LimitOrderMarket *LimitOrderMarketFilterer) FilterOrderMatched(opts *bind.FilterOpts, buyOrderId [][32]byte, sellOrderId [][32]byte, outcomeId []uint8) (*LimitOrderMarketOrderMatchedIterator, error) {

	var buyOrderIdRule []interface{}
	for _, buyOrderIdItem := range buyOrderId {
		buyOrderIdRule = append(buyOrderIdRule, buyOrderIdItem)
	}
	var sellOrderIdRule []interface{}
	for _, sellOrderIdItem := range sellOrderId {
		sellOrderIdRule = append(sellOrderIdRule, sellOrderIdItem)
	}
	var outcomeIdRule []interface{}
	for _, outcomeIdItem := range outcomeId {
		outcomeIdRule = append(outcomeIdRule, outcomeIdItem)
	}

	logs, sub, err := _LimitOrderMarket.contract.FilterLogs(opts, "OrderMatched", buyOrderIdRule, sellOrderIdRule, outcomeIdRule)
	if err != nil {
		return nil, err
	}
	return &LimitOrderMarketOrderMatchedIterator{contract: _LimitOrderMarket.contract, event: "OrderMatched", logs: logs, sub: sub}, nil
}

// WatchOrderMatched is a free log subscription operation binding the contract event 0xa674596fa044bf929b1aa6b81415a6a1534532f20319761ad36f7547e243b34a.
//
// Solidity: event OrderMatched(bytes32 indexed buyOrderId, bytes32 indexed sellOrderId, uint8 indexed outcomeId, uint256 price, uint256 amount, address buyer, address seller)
func (_LimitOrderMarket *LimitOrderMarketFilterer) WatchOrderMatched(opts *bind.WatchOpts, sink chan<- *LimitOrderMarketOrderMatched, buyOrderId [][32]byte, sellOrderId [][32]byte, outcomeId []uint8) (event.Subscription, error) {

	var buyOrderIdRule []interface{}
	for _, buyOrderIdItem := range buyOrderId {
		buyOrderIdRule = append(buyOrderIdRule, buyOrderIdItem)
	}
	var sellOrderIdRule []interface{}
	for _, sellOrderIdItem := range sellOrderId {
		sellOrderIdRule = append(sellOrderIdRule, sellOrderIdItem)
	}
	var outcomeIdRule []interface{}
	for _, outcomeIdItem := range outcomeId {
		outcomeIdRule = append(outcomeIdRule, outcomeIdItem)
	}

	logs, sub, err := _LimitOrderMarket.contract.WatchLogs(opts, "OrderMatched", buyOrderIdRule, sellOrderIdRule, outcomeIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LimitOrderMarketOrderMatched)
				if err := _LimitOrderMarket.contract.UnpackLog(event, "OrderMatched", log); err != nil {
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

// ParseOrderMatched is a log parse operation binding the contract event 0xa674596fa044bf929b1aa6b81415a6a1534532f20319761ad36f7547e243b34a.
//
// Solidity: event OrderMatched(bytes32 indexed buyOrderId, bytes32 indexed sellOrderId, uint8 indexed outcomeId, uint256 price, uint256 amount, address buyer, address seller)
func (_LimitOrderMarket *LimitOrderMarketFilterer) ParseOrderMatched(log types.Log) (*LimitOrderMarketOrderMatched, error) {
	event := new(LimitOrderMarketOrderMatched)
	if err := _LimitOrderMarket.contract.UnpackLog(event, "OrderMatched", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LimitOrderMarketOrderPartiallyFilledIterator is returned from FilterOrderPartiallyFilled and is used to iterate over the raw logs and unpacked data for OrderPartiallyFilled events raised by the LimitOrderMarket contract.
type LimitOrderMarketOrderPartiallyFilledIterator struct {
	Event *LimitOrderMarketOrderPartiallyFilled // Event containing the contract specifics and raw log

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
func (it *LimitOrderMarketOrderPartiallyFilledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LimitOrderMarketOrderPartiallyFilled)
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
		it.Event = new(LimitOrderMarketOrderPartiallyFilled)
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
func (it *LimitOrderMarketOrderPartiallyFilledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LimitOrderMarketOrderPartiallyFilledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LimitOrderMarketOrderPartiallyFilled represents a OrderPartiallyFilled event raised by the LimitOrderMarket contract.
type LimitOrderMarketOrderPartiallyFilled struct {
	OrderId         [32]byte
	FilledAmount    *big.Int
	RemainingAmount *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterOrderPartiallyFilled is a free log retrieval operation binding the contract event 0xc6be4a2e8fa0a1c95534e65d292f6be37928ffe99fd9062096dfbd66866ec26a.
//
// Solidity: event OrderPartiallyFilled(bytes32 indexed orderId, uint256 filledAmount, uint256 remainingAmount)
func (_LimitOrderMarket *LimitOrderMarketFilterer) FilterOrderPartiallyFilled(opts *bind.FilterOpts, orderId [][32]byte) (*LimitOrderMarketOrderPartiallyFilledIterator, error) {

	var orderIdRule []interface{}
	for _, orderIdItem := range orderId {
		orderIdRule = append(orderIdRule, orderIdItem)
	}

	logs, sub, err := _LimitOrderMarket.contract.FilterLogs(opts, "OrderPartiallyFilled", orderIdRule)
	if err != nil {
		return nil, err
	}
	return &LimitOrderMarketOrderPartiallyFilledIterator{contract: _LimitOrderMarket.contract, event: "OrderPartiallyFilled", logs: logs, sub: sub}, nil
}

// WatchOrderPartiallyFilled is a free log subscription operation binding the contract event 0xc6be4a2e8fa0a1c95534e65d292f6be37928ffe99fd9062096dfbd66866ec26a.
//
// Solidity: event OrderPartiallyFilled(bytes32 indexed orderId, uint256 filledAmount, uint256 remainingAmount)
func (_LimitOrderMarket *LimitOrderMarketFilterer) WatchOrderPartiallyFilled(opts *bind.WatchOpts, sink chan<- *LimitOrderMarketOrderPartiallyFilled, orderId [][32]byte) (event.Subscription, error) {

	var orderIdRule []interface{}
	for _, orderIdItem := range orderId {
		orderIdRule = append(orderIdRule, orderIdItem)
	}

	logs, sub, err := _LimitOrderMarket.contract.WatchLogs(opts, "OrderPartiallyFilled", orderIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LimitOrderMarketOrderPartiallyFilled)
				if err := _LimitOrderMarket.contract.UnpackLog(event, "OrderPartiallyFilled", log); err != nil {
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

// ParseOrderPartiallyFilled is a log parse operation binding the contract event 0xc6be4a2e8fa0a1c95534e65d292f6be37928ffe99fd9062096dfbd66866ec26a.
//
// Solidity: event OrderPartiallyFilled(bytes32 indexed orderId, uint256 filledAmount, uint256 remainingAmount)
func (_LimitOrderMarket *LimitOrderMarketFilterer) ParseOrderPartiallyFilled(log types.Log) (*LimitOrderMarketOrderPartiallyFilled, error) {
	event := new(LimitOrderMarketOrderPartiallyFilled)
	if err := _LimitOrderMarket.contract.UnpackLog(event, "OrderPartiallyFilled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LimitOrderMarketOrderPlacedIterator is returned from FilterOrderPlaced and is used to iterate over the raw logs and unpacked data for OrderPlaced events raised by the LimitOrderMarket contract.
type LimitOrderMarketOrderPlacedIterator struct {
	Event *LimitOrderMarketOrderPlaced // Event containing the contract specifics and raw log

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
func (it *LimitOrderMarketOrderPlacedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LimitOrderMarketOrderPlaced)
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
		it.Event = new(LimitOrderMarketOrderPlaced)
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
func (it *LimitOrderMarketOrderPlacedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LimitOrderMarketOrderPlacedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LimitOrderMarketOrderPlaced represents a OrderPlaced event raised by the LimitOrderMarket contract.
type LimitOrderMarketOrderPlaced struct {
	OrderId   [32]byte
	Trader    common.Address
	OutcomeId uint8
	IsBuy     bool
	Price     *big.Int
	Amount    *big.Int
	OrderType uint8
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterOrderPlaced is a free log retrieval operation binding the contract event 0xa8365d40fe5851a7e98a443de8467fe15bf002b7a850d4bcaf9b5fef1f0e0033.
//
// Solidity: event OrderPlaced(bytes32 indexed orderId, address indexed trader, uint8 indexed outcomeId, bool isBuy, uint256 price, uint256 amount, uint8 orderType)
func (_LimitOrderMarket *LimitOrderMarketFilterer) FilterOrderPlaced(opts *bind.FilterOpts, orderId [][32]byte, trader []common.Address, outcomeId []uint8) (*LimitOrderMarketOrderPlacedIterator, error) {

	var orderIdRule []interface{}
	for _, orderIdItem := range orderId {
		orderIdRule = append(orderIdRule, orderIdItem)
	}
	var traderRule []interface{}
	for _, traderItem := range trader {
		traderRule = append(traderRule, traderItem)
	}
	var outcomeIdRule []interface{}
	for _, outcomeIdItem := range outcomeId {
		outcomeIdRule = append(outcomeIdRule, outcomeIdItem)
	}

	logs, sub, err := _LimitOrderMarket.contract.FilterLogs(opts, "OrderPlaced", orderIdRule, traderRule, outcomeIdRule)
	if err != nil {
		return nil, err
	}
	return &LimitOrderMarketOrderPlacedIterator{contract: _LimitOrderMarket.contract, event: "OrderPlaced", logs: logs, sub: sub}, nil
}

// WatchOrderPlaced is a free log subscription operation binding the contract event 0xa8365d40fe5851a7e98a443de8467fe15bf002b7a850d4bcaf9b5fef1f0e0033.
//
// Solidity: event OrderPlaced(bytes32 indexed orderId, address indexed trader, uint8 indexed outcomeId, bool isBuy, uint256 price, uint256 amount, uint8 orderType)
func (_LimitOrderMarket *LimitOrderMarketFilterer) WatchOrderPlaced(opts *bind.WatchOpts, sink chan<- *LimitOrderMarketOrderPlaced, orderId [][32]byte, trader []common.Address, outcomeId []uint8) (event.Subscription, error) {

	var orderIdRule []interface{}
	for _, orderIdItem := range orderId {
		orderIdRule = append(orderIdRule, orderIdItem)
	}
	var traderRule []interface{}
	for _, traderItem := range trader {
		traderRule = append(traderRule, traderItem)
	}
	var outcomeIdRule []interface{}
	for _, outcomeIdItem := range outcomeId {
		outcomeIdRule = append(outcomeIdRule, outcomeIdItem)
	}

	logs, sub, err := _LimitOrderMarket.contract.WatchLogs(opts, "OrderPlaced", orderIdRule, traderRule, outcomeIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LimitOrderMarketOrderPlaced)
				if err := _LimitOrderMarket.contract.UnpackLog(event, "OrderPlaced", log); err != nil {
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

// ParseOrderPlaced is a log parse operation binding the contract event 0xa8365d40fe5851a7e98a443de8467fe15bf002b7a850d4bcaf9b5fef1f0e0033.
//
// Solidity: event OrderPlaced(bytes32 indexed orderId, address indexed trader, uint8 indexed outcomeId, bool isBuy, uint256 price, uint256 amount, uint8 orderType)
func (_LimitOrderMarket *LimitOrderMarketFilterer) ParseOrderPlaced(log types.Log) (*LimitOrderMarketOrderPlaced, error) {
	event := new(LimitOrderMarketOrderPlaced)
	if err := _LimitOrderMarket.contract.UnpackLog(event, "OrderPlaced", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LimitOrderMarketPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the LimitOrderMarket contract.
type LimitOrderMarketPausedIterator struct {
	Event *LimitOrderMarketPaused // Event containing the contract specifics and raw log

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
func (it *LimitOrderMarketPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LimitOrderMarketPaused)
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
		it.Event = new(LimitOrderMarketPaused)
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
func (it *LimitOrderMarketPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LimitOrderMarketPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LimitOrderMarketPaused represents a Paused event raised by the LimitOrderMarket contract.
type LimitOrderMarketPaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_LimitOrderMarket *LimitOrderMarketFilterer) FilterPaused(opts *bind.FilterOpts) (*LimitOrderMarketPausedIterator, error) {

	logs, sub, err := _LimitOrderMarket.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &LimitOrderMarketPausedIterator{contract: _LimitOrderMarket.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_LimitOrderMarket *LimitOrderMarketFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *LimitOrderMarketPaused) (event.Subscription, error) {

	logs, sub, err := _LimitOrderMarket.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LimitOrderMarketPaused)
				if err := _LimitOrderMarket.contract.UnpackLog(event, "Paused", log); err != nil {
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
func (_LimitOrderMarket *LimitOrderMarketFilterer) ParsePaused(log types.Log) (*LimitOrderMarketPaused, error) {
	event := new(LimitOrderMarketPaused)
	if err := _LimitOrderMarket.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LimitOrderMarketTradeIterator is returned from FilterTrade and is used to iterate over the raw logs and unpacked data for Trade events raised by the LimitOrderMarket contract.
type LimitOrderMarketTradeIterator struct {
	Event *LimitOrderMarketTrade // Event containing the contract specifics and raw log

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
func (it *LimitOrderMarketTradeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LimitOrderMarketTrade)
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
		it.Event = new(LimitOrderMarketTrade)
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
func (it *LimitOrderMarketTradeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LimitOrderMarketTradeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LimitOrderMarketTrade represents a Trade event raised by the LimitOrderMarket contract.
type LimitOrderMarketTrade struct {
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
func (_LimitOrderMarket *LimitOrderMarketFilterer) FilterTrade(opts *bind.FilterOpts, trader []common.Address, outcomeId []*big.Int) (*LimitOrderMarketTradeIterator, error) {

	var traderRule []interface{}
	for _, traderItem := range trader {
		traderRule = append(traderRule, traderItem)
	}
	var outcomeIdRule []interface{}
	for _, outcomeIdItem := range outcomeId {
		outcomeIdRule = append(outcomeIdRule, outcomeIdItem)
	}

	logs, sub, err := _LimitOrderMarket.contract.FilterLogs(opts, "Trade", traderRule, outcomeIdRule)
	if err != nil {
		return nil, err
	}
	return &LimitOrderMarketTradeIterator{contract: _LimitOrderMarket.contract, event: "Trade", logs: logs, sub: sub}, nil
}

// WatchTrade is a free log subscription operation binding the contract event 0xe34b2a81bbc1e1a545c34243f3cc283b6b6d1f4c2153be2a47b2612247e45865.
//
// Solidity: event Trade(address indexed trader, uint256 indexed outcomeId, uint256 amountIn, uint256 amountOut, uint256 fee, bool isBuy)
func (_LimitOrderMarket *LimitOrderMarketFilterer) WatchTrade(opts *bind.WatchOpts, sink chan<- *LimitOrderMarketTrade, trader []common.Address, outcomeId []*big.Int) (event.Subscription, error) {

	var traderRule []interface{}
	for _, traderItem := range trader {
		traderRule = append(traderRule, traderItem)
	}
	var outcomeIdRule []interface{}
	for _, outcomeIdItem := range outcomeId {
		outcomeIdRule = append(outcomeIdRule, outcomeIdItem)
	}

	logs, sub, err := _LimitOrderMarket.contract.WatchLogs(opts, "Trade", traderRule, outcomeIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LimitOrderMarketTrade)
				if err := _LimitOrderMarket.contract.UnpackLog(event, "Trade", log); err != nil {
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
func (_LimitOrderMarket *LimitOrderMarketFilterer) ParseTrade(log types.Log) (*LimitOrderMarketTrade, error) {
	event := new(LimitOrderMarketTrade)
	if err := _LimitOrderMarket.contract.UnpackLog(event, "Trade", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LimitOrderMarketTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the LimitOrderMarket contract.
type LimitOrderMarketTransferIterator struct {
	Event *LimitOrderMarketTransfer // Event containing the contract specifics and raw log

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
func (it *LimitOrderMarketTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LimitOrderMarketTransfer)
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
		it.Event = new(LimitOrderMarketTransfer)
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
func (it *LimitOrderMarketTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LimitOrderMarketTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LimitOrderMarketTransfer represents a Transfer event raised by the LimitOrderMarket contract.
type LimitOrderMarketTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_LimitOrderMarket *LimitOrderMarketFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*LimitOrderMarketTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _LimitOrderMarket.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &LimitOrderMarketTransferIterator{contract: _LimitOrderMarket.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_LimitOrderMarket *LimitOrderMarketFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *LimitOrderMarketTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _LimitOrderMarket.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LimitOrderMarketTransfer)
				if err := _LimitOrderMarket.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_LimitOrderMarket *LimitOrderMarketFilterer) ParseTransfer(log types.Log) (*LimitOrderMarketTransfer, error) {
	event := new(LimitOrderMarketTransfer)
	if err := _LimitOrderMarket.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LimitOrderMarketUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the LimitOrderMarket contract.
type LimitOrderMarketUnpausedIterator struct {
	Event *LimitOrderMarketUnpaused // Event containing the contract specifics and raw log

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
func (it *LimitOrderMarketUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LimitOrderMarketUnpaused)
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
		it.Event = new(LimitOrderMarketUnpaused)
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
func (it *LimitOrderMarketUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LimitOrderMarketUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LimitOrderMarketUnpaused represents a Unpaused event raised by the LimitOrderMarket contract.
type LimitOrderMarketUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_LimitOrderMarket *LimitOrderMarketFilterer) FilterUnpaused(opts *bind.FilterOpts) (*LimitOrderMarketUnpausedIterator, error) {

	logs, sub, err := _LimitOrderMarket.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &LimitOrderMarketUnpausedIterator{contract: _LimitOrderMarket.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_LimitOrderMarket *LimitOrderMarketFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *LimitOrderMarketUnpaused) (event.Subscription, error) {

	logs, sub, err := _LimitOrderMarket.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LimitOrderMarketUnpaused)
				if err := _LimitOrderMarket.contract.UnpackLog(event, "Unpaused", log); err != nil {
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
func (_LimitOrderMarket *LimitOrderMarketFilterer) ParseUnpaused(log types.Log) (*LimitOrderMarketUnpaused, error) {
	event := new(LimitOrderMarketUnpaused)
	if err := _LimitOrderMarket.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
