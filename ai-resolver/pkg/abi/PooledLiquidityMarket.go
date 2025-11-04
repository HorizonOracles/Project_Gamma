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


// PooledLiquidityMarketMetaData contains all meta data concerning the PooledLiquidityMarket contract.
var PooledLiquidityMarketMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_marketId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_collateralToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_outcomeToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_feeSplitter\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_horizonPerks\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_closeTime\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_lpTokenName\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"_lpTokenSymbol\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"FEE_TIER\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint24\",\"internalType\":\"uint24\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"MAX_TICK\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"int24\",\"internalType\":\"int24\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"MINIMUM_LIQUIDITY\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"MIN_TICK\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"int24\",\"internalType\":\"int24\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"PRICE_PRECISION\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"TICK_SPACING\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"int24\",\"internalType\":\"int24\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"addLiquidity\",\"inputs\":[{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"lpTokens\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"admin\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"allowance\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"spender\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"approve\",\"inputs\":[{\"name\":\"spender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"balanceOf\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"burnPosition\",\"inputs\":[{\"name\":\"tickLower\",\"type\":\"int24\",\"internalType\":\"int24\"},{\"name\":\"tickUpper\",\"type\":\"int24\",\"internalType\":\"int24\"},{\"name\":\"liquidityToBurn\",\"type\":\"uint128\",\"internalType\":\"uint128\"}],\"outputs\":[{\"name\":\"amount0\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"amount1\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"buy\",\"inputs\":[{\"name\":\"outcomeId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"collateralIn\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"minTokensOut\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"tokensOut\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"closeTime\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"collateralToken\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIERC20\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"collectFees\",\"inputs\":[{\"name\":\"tickLower\",\"type\":\"int24\",\"internalType\":\"int24\"},{\"name\":\"tickUpper\",\"type\":\"int24\",\"internalType\":\"int24\"}],\"outputs\":[{\"name\":\"amount0\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"amount1\",\"type\":\"uint128\",\"internalType\":\"uint128\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"decimals\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"feeSplitter\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractFeeSplitter\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"fundRedemptions\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getMarketInfo\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structIMarket.MarketInfo\",\"components\":[{\"name\":\"marketId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"marketType\",\"type\":\"uint8\",\"internalType\":\"enumIMarket.MarketType\"},{\"name\":\"collateralToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"closeTime\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"outcomeCount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"isResolved\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"isPaused\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getMarketType\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"enumIMarket.MarketType\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getOutcomeCount\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getPrice\",\"inputs\":[{\"name\":\"outcomeId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"price\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getQuoteBuy\",\"inputs\":[{\"name\":\"outcomeId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"collateralIn\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"user\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"tokensOut\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"fee\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getQuoteSell\",\"inputs\":[{\"name\":\"outcomeId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"tokensIn\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"user\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"collateralOut\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"fee\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"horizonPerks\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractHorizonPerks\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"marketId\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"marketType\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"enumIMarket.MarketType\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"mintPosition\",\"inputs\":[{\"name\":\"tickLower\",\"type\":\"int24\",\"internalType\":\"int24\"},{\"name\":\"tickUpper\",\"type\":\"int24\",\"internalType\":\"int24\"},{\"name\":\"liquidityDesired\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"amount0Max\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"amount1Max\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"liquidity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"amount0\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"amount1\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"name\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"onERC1155BatchReceived\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"},{\"name\":\"\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"},{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"onERC1155Received\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"outcomeCount\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"outcomeToken\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractOutcomeToken\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pause\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"paused\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"poolState\",\"inputs\":[],\"outputs\":[{\"name\":\"sqrtPriceX96\",\"type\":\"uint160\",\"internalType\":\"uint160\"},{\"name\":\"tick\",\"type\":\"int24\",\"internalType\":\"int24\"},{\"name\":\"liquidity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"feeGrowthGlobal0X128\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"feeGrowthGlobal1X128\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"positions\",\"inputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"liquidity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"feeGrowthInside0LastX128\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"feeGrowthInside1LastX128\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"tokensOwed0\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"tokensOwed1\",\"type\":\"uint128\",\"internalType\":\"uint128\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"removeLiquidity\",\"inputs\":[{\"name\":\"lpTokens\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"collateralOut\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"reserves\",\"inputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"sell\",\"inputs\":[{\"name\":\"outcomeId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"tokensIn\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"minCollateralOut\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"collateralOut\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setAdmin\",\"inputs\":[{\"name\":\"newAdmin\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"supportsInterface\",\"inputs\":[{\"name\":\"interfaceId\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"symbol\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"ticks\",\"inputs\":[{\"name\":\"\",\"type\":\"int24\",\"internalType\":\"int24\"}],\"outputs\":[{\"name\":\"liquidityGross\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"liquidityNet\",\"type\":\"int128\",\"internalType\":\"int128\"},{\"name\":\"feeGrowthOutside0X128\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"feeGrowthOutside1X128\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"initialized\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"totalCollateral\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"totalSupply\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"transfer\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferFrom\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"unpause\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"Approval\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"spender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"FeesCollected\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"tickLower\",\"type\":\"int24\",\"indexed\":true,\"internalType\":\"int24\"},{\"name\":\"tickUpper\",\"type\":\"int24\",\"indexed\":true,\"internalType\":\"int24\"},{\"name\":\"amount0\",\"type\":\"uint128\",\"indexed\":false,\"internalType\":\"uint128\"},{\"name\":\"amount1\",\"type\":\"uint128\",\"indexed\":false,\"internalType\":\"uint128\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"LiquidityChanged\",\"inputs\":[{\"name\":\"provider\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"isAddition\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Paused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"PositionBurned\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"tickLower\",\"type\":\"int24\",\"indexed\":true,\"internalType\":\"int24\"},{\"name\":\"tickUpper\",\"type\":\"int24\",\"indexed\":true,\"internalType\":\"int24\"},{\"name\":\"liquidity\",\"type\":\"uint128\",\"indexed\":false,\"internalType\":\"uint128\"},{\"name\":\"amount0\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"amount1\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"PositionMinted\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"tickLower\",\"type\":\"int24\",\"indexed\":true,\"internalType\":\"int24\"},{\"name\":\"tickUpper\",\"type\":\"int24\",\"indexed\":true,\"internalType\":\"int24\"},{\"name\":\"liquidity\",\"type\":\"uint128\",\"indexed\":false,\"internalType\":\"uint128\"},{\"name\":\"amount0\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"amount1\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Swap\",\"inputs\":[{\"name\":\"trader\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"outcomeId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"amountIn\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"amountOut\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"sqrtPriceX96\",\"type\":\"uint160\",\"indexed\":false,\"internalType\":\"uint160\"},{\"name\":\"tick\",\"type\":\"int24\",\"indexed\":false,\"internalType\":\"int24\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Trade\",\"inputs\":[{\"name\":\"trader\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"outcomeId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"amountIn\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"amountOut\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"fee\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"isBuy\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Transfer\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Unpaused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"ERC20InsufficientAllowance\",\"inputs\":[{\"name\":\"spender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"allowance\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"needed\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"ERC20InsufficientBalance\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"balance\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"needed\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"ERC20InvalidApprover\",\"inputs\":[{\"name\":\"approver\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ERC20InvalidReceiver\",\"inputs\":[{\"name\":\"receiver\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ERC20InvalidSender\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ERC20InvalidSpender\",\"inputs\":[{\"name\":\"spender\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"EnforcedPause\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ExpectedPause\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InsufficientLPTokens\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InsufficientLiquidity\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InsufficientPosition\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidAmount\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidOutcomeId\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidState\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidTick\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidTickRange\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"MarketClosed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"MarketResolved\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"MinimumLiquidityRequired\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"PositionNotFound\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ReentrancyGuardReentrantCall\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"SafeERC20FailedOperation\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"SlippageExceeded\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"TickSpacingError\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"Unauthorized\",\"inputs\":[]}]",
}

// PooledLiquidityMarketABI is the input ABI used to generate the binding from.
// Deprecated: Use PooledLiquidityMarketMetaData.ABI instead.
var PooledLiquidityMarketABI = PooledLiquidityMarketMetaData.ABI

// PooledLiquidityMarket is an auto generated Go binding around an Ethereum contract.
type PooledLiquidityMarket struct {
	PooledLiquidityMarketCaller     // Read-only binding to the contract
	PooledLiquidityMarketTransactor // Write-only binding to the contract
	PooledLiquidityMarketFilterer   // Log filterer for contract events
}

// PooledLiquidityMarketCaller is an auto generated read-only Go binding around an Ethereum contract.
type PooledLiquidityMarketCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PooledLiquidityMarketTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PooledLiquidityMarketTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PooledLiquidityMarketFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PooledLiquidityMarketFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PooledLiquidityMarketSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PooledLiquidityMarketSession struct {
	Contract     *PooledLiquidityMarket // Generic contract binding to set the session for
	CallOpts     bind.CallOpts          // Call options to use throughout this session
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// PooledLiquidityMarketCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PooledLiquidityMarketCallerSession struct {
	Contract *PooledLiquidityMarketCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                // Call options to use throughout this session
}

// PooledLiquidityMarketTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PooledLiquidityMarketTransactorSession struct {
	Contract     *PooledLiquidityMarketTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                // Transaction auth options to use throughout this session
}

// PooledLiquidityMarketRaw is an auto generated low-level Go binding around an Ethereum contract.
type PooledLiquidityMarketRaw struct {
	Contract *PooledLiquidityMarket // Generic contract binding to access the raw methods on
}

// PooledLiquidityMarketCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PooledLiquidityMarketCallerRaw struct {
	Contract *PooledLiquidityMarketCaller // Generic read-only contract binding to access the raw methods on
}

// PooledLiquidityMarketTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PooledLiquidityMarketTransactorRaw struct {
	Contract *PooledLiquidityMarketTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPooledLiquidityMarket creates a new instance of PooledLiquidityMarket, bound to a specific deployed contract.
func NewPooledLiquidityMarket(address common.Address, backend bind.ContractBackend) (*PooledLiquidityMarket, error) {
	contract, err := bindPooledLiquidityMarket(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &PooledLiquidityMarket{PooledLiquidityMarketCaller: PooledLiquidityMarketCaller{contract: contract}, PooledLiquidityMarketTransactor: PooledLiquidityMarketTransactor{contract: contract}, PooledLiquidityMarketFilterer: PooledLiquidityMarketFilterer{contract: contract}}, nil
}

// NewPooledLiquidityMarketCaller creates a new read-only instance of PooledLiquidityMarket, bound to a specific deployed contract.
func NewPooledLiquidityMarketCaller(address common.Address, caller bind.ContractCaller) (*PooledLiquidityMarketCaller, error) {
	contract, err := bindPooledLiquidityMarket(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PooledLiquidityMarketCaller{contract: contract}, nil
}

// NewPooledLiquidityMarketTransactor creates a new write-only instance of PooledLiquidityMarket, bound to a specific deployed contract.
func NewPooledLiquidityMarketTransactor(address common.Address, transactor bind.ContractTransactor) (*PooledLiquidityMarketTransactor, error) {
	contract, err := bindPooledLiquidityMarket(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PooledLiquidityMarketTransactor{contract: contract}, nil
}

// NewPooledLiquidityMarketFilterer creates a new log filterer instance of PooledLiquidityMarket, bound to a specific deployed contract.
func NewPooledLiquidityMarketFilterer(address common.Address, filterer bind.ContractFilterer) (*PooledLiquidityMarketFilterer, error) {
	contract, err := bindPooledLiquidityMarket(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PooledLiquidityMarketFilterer{contract: contract}, nil
}

// bindPooledLiquidityMarket binds a generic wrapper to an already deployed contract.
func bindPooledLiquidityMarket(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := PooledLiquidityMarketMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PooledLiquidityMarket *PooledLiquidityMarketRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PooledLiquidityMarket.Contract.PooledLiquidityMarketCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PooledLiquidityMarket *PooledLiquidityMarketRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PooledLiquidityMarket.Contract.PooledLiquidityMarketTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PooledLiquidityMarket *PooledLiquidityMarketRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PooledLiquidityMarket.Contract.PooledLiquidityMarketTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PooledLiquidityMarket *PooledLiquidityMarketCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PooledLiquidityMarket.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PooledLiquidityMarket *PooledLiquidityMarketTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PooledLiquidityMarket.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PooledLiquidityMarket *PooledLiquidityMarketTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PooledLiquidityMarket.Contract.contract.Transact(opts, method, params...)
}

// FEETIER is a free data retrieval call binding the contract method 0x4c69a6c9.
//
// Solidity: function FEE_TIER() view returns(uint24)
func (_PooledLiquidityMarket *PooledLiquidityMarketCaller) FEETIER(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PooledLiquidityMarket.contract.Call(opts, &out, "FEE_TIER")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FEETIER is a free data retrieval call binding the contract method 0x4c69a6c9.
//
// Solidity: function FEE_TIER() view returns(uint24)
func (_PooledLiquidityMarket *PooledLiquidityMarketSession) FEETIER() (*big.Int, error) {
	return _PooledLiquidityMarket.Contract.FEETIER(&_PooledLiquidityMarket.CallOpts)
}

// FEETIER is a free data retrieval call binding the contract method 0x4c69a6c9.
//
// Solidity: function FEE_TIER() view returns(uint24)
func (_PooledLiquidityMarket *PooledLiquidityMarketCallerSession) FEETIER() (*big.Int, error) {
	return _PooledLiquidityMarket.Contract.FEETIER(&_PooledLiquidityMarket.CallOpts)
}

// MAXTICK is a free data retrieval call binding the contract method 0x6882a888.
//
// Solidity: function MAX_TICK() view returns(int24)
func (_PooledLiquidityMarket *PooledLiquidityMarketCaller) MAXTICK(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PooledLiquidityMarket.contract.Call(opts, &out, "MAX_TICK")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXTICK is a free data retrieval call binding the contract method 0x6882a888.
//
// Solidity: function MAX_TICK() view returns(int24)
func (_PooledLiquidityMarket *PooledLiquidityMarketSession) MAXTICK() (*big.Int, error) {
	return _PooledLiquidityMarket.Contract.MAXTICK(&_PooledLiquidityMarket.CallOpts)
}

// MAXTICK is a free data retrieval call binding the contract method 0x6882a888.
//
// Solidity: function MAX_TICK() view returns(int24)
func (_PooledLiquidityMarket *PooledLiquidityMarketCallerSession) MAXTICK() (*big.Int, error) {
	return _PooledLiquidityMarket.Contract.MAXTICK(&_PooledLiquidityMarket.CallOpts)
}

// MINIMUMLIQUIDITY is a free data retrieval call binding the contract method 0xba9a7a56.
//
// Solidity: function MINIMUM_LIQUIDITY() view returns(uint256)
func (_PooledLiquidityMarket *PooledLiquidityMarketCaller) MINIMUMLIQUIDITY(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PooledLiquidityMarket.contract.Call(opts, &out, "MINIMUM_LIQUIDITY")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MINIMUMLIQUIDITY is a free data retrieval call binding the contract method 0xba9a7a56.
//
// Solidity: function MINIMUM_LIQUIDITY() view returns(uint256)
func (_PooledLiquidityMarket *PooledLiquidityMarketSession) MINIMUMLIQUIDITY() (*big.Int, error) {
	return _PooledLiquidityMarket.Contract.MINIMUMLIQUIDITY(&_PooledLiquidityMarket.CallOpts)
}

// MINIMUMLIQUIDITY is a free data retrieval call binding the contract method 0xba9a7a56.
//
// Solidity: function MINIMUM_LIQUIDITY() view returns(uint256)
func (_PooledLiquidityMarket *PooledLiquidityMarketCallerSession) MINIMUMLIQUIDITY() (*big.Int, error) {
	return _PooledLiquidityMarket.Contract.MINIMUMLIQUIDITY(&_PooledLiquidityMarket.CallOpts)
}

// MINTICK is a free data retrieval call binding the contract method 0xa1634b14.
//
// Solidity: function MIN_TICK() view returns(int24)
func (_PooledLiquidityMarket *PooledLiquidityMarketCaller) MINTICK(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PooledLiquidityMarket.contract.Call(opts, &out, "MIN_TICK")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MINTICK is a free data retrieval call binding the contract method 0xa1634b14.
//
// Solidity: function MIN_TICK() view returns(int24)
func (_PooledLiquidityMarket *PooledLiquidityMarketSession) MINTICK() (*big.Int, error) {
	return _PooledLiquidityMarket.Contract.MINTICK(&_PooledLiquidityMarket.CallOpts)
}

// MINTICK is a free data retrieval call binding the contract method 0xa1634b14.
//
// Solidity: function MIN_TICK() view returns(int24)
func (_PooledLiquidityMarket *PooledLiquidityMarketCallerSession) MINTICK() (*big.Int, error) {
	return _PooledLiquidityMarket.Contract.MINTICK(&_PooledLiquidityMarket.CallOpts)
}

// PRICEPRECISION is a free data retrieval call binding the contract method 0x95082d25.
//
// Solidity: function PRICE_PRECISION() view returns(uint256)
func (_PooledLiquidityMarket *PooledLiquidityMarketCaller) PRICEPRECISION(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PooledLiquidityMarket.contract.Call(opts, &out, "PRICE_PRECISION")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PRICEPRECISION is a free data retrieval call binding the contract method 0x95082d25.
//
// Solidity: function PRICE_PRECISION() view returns(uint256)
func (_PooledLiquidityMarket *PooledLiquidityMarketSession) PRICEPRECISION() (*big.Int, error) {
	return _PooledLiquidityMarket.Contract.PRICEPRECISION(&_PooledLiquidityMarket.CallOpts)
}

// PRICEPRECISION is a free data retrieval call binding the contract method 0x95082d25.
//
// Solidity: function PRICE_PRECISION() view returns(uint256)
func (_PooledLiquidityMarket *PooledLiquidityMarketCallerSession) PRICEPRECISION() (*big.Int, error) {
	return _PooledLiquidityMarket.Contract.PRICEPRECISION(&_PooledLiquidityMarket.CallOpts)
}

// TICKSPACING is a free data retrieval call binding the contract method 0x46ca626b.
//
// Solidity: function TICK_SPACING() view returns(int24)
func (_PooledLiquidityMarket *PooledLiquidityMarketCaller) TICKSPACING(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PooledLiquidityMarket.contract.Call(opts, &out, "TICK_SPACING")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TICKSPACING is a free data retrieval call binding the contract method 0x46ca626b.
//
// Solidity: function TICK_SPACING() view returns(int24)
func (_PooledLiquidityMarket *PooledLiquidityMarketSession) TICKSPACING() (*big.Int, error) {
	return _PooledLiquidityMarket.Contract.TICKSPACING(&_PooledLiquidityMarket.CallOpts)
}

// TICKSPACING is a free data retrieval call binding the contract method 0x46ca626b.
//
// Solidity: function TICK_SPACING() view returns(int24)
func (_PooledLiquidityMarket *PooledLiquidityMarketCallerSession) TICKSPACING() (*big.Int, error) {
	return _PooledLiquidityMarket.Contract.TICKSPACING(&_PooledLiquidityMarket.CallOpts)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_PooledLiquidityMarket *PooledLiquidityMarketCaller) Admin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PooledLiquidityMarket.contract.Call(opts, &out, "admin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_PooledLiquidityMarket *PooledLiquidityMarketSession) Admin() (common.Address, error) {
	return _PooledLiquidityMarket.Contract.Admin(&_PooledLiquidityMarket.CallOpts)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_PooledLiquidityMarket *PooledLiquidityMarketCallerSession) Admin() (common.Address, error) {
	return _PooledLiquidityMarket.Contract.Admin(&_PooledLiquidityMarket.CallOpts)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_PooledLiquidityMarket *PooledLiquidityMarketCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _PooledLiquidityMarket.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_PooledLiquidityMarket *PooledLiquidityMarketSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _PooledLiquidityMarket.Contract.Allowance(&_PooledLiquidityMarket.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_PooledLiquidityMarket *PooledLiquidityMarketCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _PooledLiquidityMarket.Contract.Allowance(&_PooledLiquidityMarket.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_PooledLiquidityMarket *PooledLiquidityMarketCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _PooledLiquidityMarket.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_PooledLiquidityMarket *PooledLiquidityMarketSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _PooledLiquidityMarket.Contract.BalanceOf(&_PooledLiquidityMarket.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_PooledLiquidityMarket *PooledLiquidityMarketCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _PooledLiquidityMarket.Contract.BalanceOf(&_PooledLiquidityMarket.CallOpts, account)
}

// CloseTime is a free data retrieval call binding the contract method 0x627749e6.
//
// Solidity: function closeTime() view returns(uint256)
func (_PooledLiquidityMarket *PooledLiquidityMarketCaller) CloseTime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PooledLiquidityMarket.contract.Call(opts, &out, "closeTime")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CloseTime is a free data retrieval call binding the contract method 0x627749e6.
//
// Solidity: function closeTime() view returns(uint256)
func (_PooledLiquidityMarket *PooledLiquidityMarketSession) CloseTime() (*big.Int, error) {
	return _PooledLiquidityMarket.Contract.CloseTime(&_PooledLiquidityMarket.CallOpts)
}

// CloseTime is a free data retrieval call binding the contract method 0x627749e6.
//
// Solidity: function closeTime() view returns(uint256)
func (_PooledLiquidityMarket *PooledLiquidityMarketCallerSession) CloseTime() (*big.Int, error) {
	return _PooledLiquidityMarket.Contract.CloseTime(&_PooledLiquidityMarket.CallOpts)
}

// CollateralToken is a free data retrieval call binding the contract method 0xb2016bd4.
//
// Solidity: function collateralToken() view returns(address)
func (_PooledLiquidityMarket *PooledLiquidityMarketCaller) CollateralToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PooledLiquidityMarket.contract.Call(opts, &out, "collateralToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CollateralToken is a free data retrieval call binding the contract method 0xb2016bd4.
//
// Solidity: function collateralToken() view returns(address)
func (_PooledLiquidityMarket *PooledLiquidityMarketSession) CollateralToken() (common.Address, error) {
	return _PooledLiquidityMarket.Contract.CollateralToken(&_PooledLiquidityMarket.CallOpts)
}

// CollateralToken is a free data retrieval call binding the contract method 0xb2016bd4.
//
// Solidity: function collateralToken() view returns(address)
func (_PooledLiquidityMarket *PooledLiquidityMarketCallerSession) CollateralToken() (common.Address, error) {
	return _PooledLiquidityMarket.Contract.CollateralToken(&_PooledLiquidityMarket.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_PooledLiquidityMarket *PooledLiquidityMarketCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _PooledLiquidityMarket.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_PooledLiquidityMarket *PooledLiquidityMarketSession) Decimals() (uint8, error) {
	return _PooledLiquidityMarket.Contract.Decimals(&_PooledLiquidityMarket.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_PooledLiquidityMarket *PooledLiquidityMarketCallerSession) Decimals() (uint8, error) {
	return _PooledLiquidityMarket.Contract.Decimals(&_PooledLiquidityMarket.CallOpts)
}

// FeeSplitter is a free data retrieval call binding the contract method 0x6052970c.
//
// Solidity: function feeSplitter() view returns(address)
func (_PooledLiquidityMarket *PooledLiquidityMarketCaller) FeeSplitter(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PooledLiquidityMarket.contract.Call(opts, &out, "feeSplitter")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FeeSplitter is a free data retrieval call binding the contract method 0x6052970c.
//
// Solidity: function feeSplitter() view returns(address)
func (_PooledLiquidityMarket *PooledLiquidityMarketSession) FeeSplitter() (common.Address, error) {
	return _PooledLiquidityMarket.Contract.FeeSplitter(&_PooledLiquidityMarket.CallOpts)
}

// FeeSplitter is a free data retrieval call binding the contract method 0x6052970c.
//
// Solidity: function feeSplitter() view returns(address)
func (_PooledLiquidityMarket *PooledLiquidityMarketCallerSession) FeeSplitter() (common.Address, error) {
	return _PooledLiquidityMarket.Contract.FeeSplitter(&_PooledLiquidityMarket.CallOpts)
}

// GetMarketInfo is a free data retrieval call binding the contract method 0x23341a05.
//
// Solidity: function getMarketInfo() view returns((uint256,uint8,address,uint256,uint256,bool,bool))
func (_PooledLiquidityMarket *PooledLiquidityMarketCaller) GetMarketInfo(opts *bind.CallOpts) (IMarketMarketInfo, error) {
	var out []interface{}
	err := _PooledLiquidityMarket.contract.Call(opts, &out, "getMarketInfo")

	if err != nil {
		return *new(IMarketMarketInfo), err
	}

	out0 := *abi.ConvertType(out[0], new(IMarketMarketInfo)).(*IMarketMarketInfo)

	return out0, err

}

// GetMarketInfo is a free data retrieval call binding the contract method 0x23341a05.
//
// Solidity: function getMarketInfo() view returns((uint256,uint8,address,uint256,uint256,bool,bool))
func (_PooledLiquidityMarket *PooledLiquidityMarketSession) GetMarketInfo() (IMarketMarketInfo, error) {
	return _PooledLiquidityMarket.Contract.GetMarketInfo(&_PooledLiquidityMarket.CallOpts)
}

// GetMarketInfo is a free data retrieval call binding the contract method 0x23341a05.
//
// Solidity: function getMarketInfo() view returns((uint256,uint8,address,uint256,uint256,bool,bool))
func (_PooledLiquidityMarket *PooledLiquidityMarketCallerSession) GetMarketInfo() (IMarketMarketInfo, error) {
	return _PooledLiquidityMarket.Contract.GetMarketInfo(&_PooledLiquidityMarket.CallOpts)
}

// GetMarketType is a free data retrieval call binding the contract method 0x33e7a1d0.
//
// Solidity: function getMarketType() view returns(uint8)
func (_PooledLiquidityMarket *PooledLiquidityMarketCaller) GetMarketType(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _PooledLiquidityMarket.contract.Call(opts, &out, "getMarketType")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// GetMarketType is a free data retrieval call binding the contract method 0x33e7a1d0.
//
// Solidity: function getMarketType() view returns(uint8)
func (_PooledLiquidityMarket *PooledLiquidityMarketSession) GetMarketType() (uint8, error) {
	return _PooledLiquidityMarket.Contract.GetMarketType(&_PooledLiquidityMarket.CallOpts)
}

// GetMarketType is a free data retrieval call binding the contract method 0x33e7a1d0.
//
// Solidity: function getMarketType() view returns(uint8)
func (_PooledLiquidityMarket *PooledLiquidityMarketCallerSession) GetMarketType() (uint8, error) {
	return _PooledLiquidityMarket.Contract.GetMarketType(&_PooledLiquidityMarket.CallOpts)
}

// GetOutcomeCount is a free data retrieval call binding the contract method 0x7dc8f086.
//
// Solidity: function getOutcomeCount() view returns(uint256)
func (_PooledLiquidityMarket *PooledLiquidityMarketCaller) GetOutcomeCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PooledLiquidityMarket.contract.Call(opts, &out, "getOutcomeCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetOutcomeCount is a free data retrieval call binding the contract method 0x7dc8f086.
//
// Solidity: function getOutcomeCount() view returns(uint256)
func (_PooledLiquidityMarket *PooledLiquidityMarketSession) GetOutcomeCount() (*big.Int, error) {
	return _PooledLiquidityMarket.Contract.GetOutcomeCount(&_PooledLiquidityMarket.CallOpts)
}

// GetOutcomeCount is a free data retrieval call binding the contract method 0x7dc8f086.
//
// Solidity: function getOutcomeCount() view returns(uint256)
func (_PooledLiquidityMarket *PooledLiquidityMarketCallerSession) GetOutcomeCount() (*big.Int, error) {
	return _PooledLiquidityMarket.Contract.GetOutcomeCount(&_PooledLiquidityMarket.CallOpts)
}

// GetPrice is a free data retrieval call binding the contract method 0xe7572230.
//
// Solidity: function getPrice(uint256 outcomeId) view returns(uint256 price)
func (_PooledLiquidityMarket *PooledLiquidityMarketCaller) GetPrice(opts *bind.CallOpts, outcomeId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _PooledLiquidityMarket.contract.Call(opts, &out, "getPrice", outcomeId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetPrice is a free data retrieval call binding the contract method 0xe7572230.
//
// Solidity: function getPrice(uint256 outcomeId) view returns(uint256 price)
func (_PooledLiquidityMarket *PooledLiquidityMarketSession) GetPrice(outcomeId *big.Int) (*big.Int, error) {
	return _PooledLiquidityMarket.Contract.GetPrice(&_PooledLiquidityMarket.CallOpts, outcomeId)
}

// GetPrice is a free data retrieval call binding the contract method 0xe7572230.
//
// Solidity: function getPrice(uint256 outcomeId) view returns(uint256 price)
func (_PooledLiquidityMarket *PooledLiquidityMarketCallerSession) GetPrice(outcomeId *big.Int) (*big.Int, error) {
	return _PooledLiquidityMarket.Contract.GetPrice(&_PooledLiquidityMarket.CallOpts, outcomeId)
}

// GetQuoteBuy is a free data retrieval call binding the contract method 0xca6d5811.
//
// Solidity: function getQuoteBuy(uint256 outcomeId, uint256 collateralIn, address user) view returns(uint256 tokensOut, uint256 fee)
func (_PooledLiquidityMarket *PooledLiquidityMarketCaller) GetQuoteBuy(opts *bind.CallOpts, outcomeId *big.Int, collateralIn *big.Int, user common.Address) (struct {
	TokensOut *big.Int
	Fee       *big.Int
}, error) {
	var out []interface{}
	err := _PooledLiquidityMarket.contract.Call(opts, &out, "getQuoteBuy", outcomeId, collateralIn, user)

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
func (_PooledLiquidityMarket *PooledLiquidityMarketSession) GetQuoteBuy(outcomeId *big.Int, collateralIn *big.Int, user common.Address) (struct {
	TokensOut *big.Int
	Fee       *big.Int
}, error) {
	return _PooledLiquidityMarket.Contract.GetQuoteBuy(&_PooledLiquidityMarket.CallOpts, outcomeId, collateralIn, user)
}

// GetQuoteBuy is a free data retrieval call binding the contract method 0xca6d5811.
//
// Solidity: function getQuoteBuy(uint256 outcomeId, uint256 collateralIn, address user) view returns(uint256 tokensOut, uint256 fee)
func (_PooledLiquidityMarket *PooledLiquidityMarketCallerSession) GetQuoteBuy(outcomeId *big.Int, collateralIn *big.Int, user common.Address) (struct {
	TokensOut *big.Int
	Fee       *big.Int
}, error) {
	return _PooledLiquidityMarket.Contract.GetQuoteBuy(&_PooledLiquidityMarket.CallOpts, outcomeId, collateralIn, user)
}

// GetQuoteSell is a free data retrieval call binding the contract method 0x8b5e8a24.
//
// Solidity: function getQuoteSell(uint256 outcomeId, uint256 tokensIn, address user) view returns(uint256 collateralOut, uint256 fee)
func (_PooledLiquidityMarket *PooledLiquidityMarketCaller) GetQuoteSell(opts *bind.CallOpts, outcomeId *big.Int, tokensIn *big.Int, user common.Address) (struct {
	CollateralOut *big.Int
	Fee           *big.Int
}, error) {
	var out []interface{}
	err := _PooledLiquidityMarket.contract.Call(opts, &out, "getQuoteSell", outcomeId, tokensIn, user)

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
func (_PooledLiquidityMarket *PooledLiquidityMarketSession) GetQuoteSell(outcomeId *big.Int, tokensIn *big.Int, user common.Address) (struct {
	CollateralOut *big.Int
	Fee           *big.Int
}, error) {
	return _PooledLiquidityMarket.Contract.GetQuoteSell(&_PooledLiquidityMarket.CallOpts, outcomeId, tokensIn, user)
}

// GetQuoteSell is a free data retrieval call binding the contract method 0x8b5e8a24.
//
// Solidity: function getQuoteSell(uint256 outcomeId, uint256 tokensIn, address user) view returns(uint256 collateralOut, uint256 fee)
func (_PooledLiquidityMarket *PooledLiquidityMarketCallerSession) GetQuoteSell(outcomeId *big.Int, tokensIn *big.Int, user common.Address) (struct {
	CollateralOut *big.Int
	Fee           *big.Int
}, error) {
	return _PooledLiquidityMarket.Contract.GetQuoteSell(&_PooledLiquidityMarket.CallOpts, outcomeId, tokensIn, user)
}

// HorizonPerks is a free data retrieval call binding the contract method 0xffe02e34.
//
// Solidity: function horizonPerks() view returns(address)
func (_PooledLiquidityMarket *PooledLiquidityMarketCaller) HorizonPerks(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PooledLiquidityMarket.contract.Call(opts, &out, "horizonPerks")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// HorizonPerks is a free data retrieval call binding the contract method 0xffe02e34.
//
// Solidity: function horizonPerks() view returns(address)
func (_PooledLiquidityMarket *PooledLiquidityMarketSession) HorizonPerks() (common.Address, error) {
	return _PooledLiquidityMarket.Contract.HorizonPerks(&_PooledLiquidityMarket.CallOpts)
}

// HorizonPerks is a free data retrieval call binding the contract method 0xffe02e34.
//
// Solidity: function horizonPerks() view returns(address)
func (_PooledLiquidityMarket *PooledLiquidityMarketCallerSession) HorizonPerks() (common.Address, error) {
	return _PooledLiquidityMarket.Contract.HorizonPerks(&_PooledLiquidityMarket.CallOpts)
}

// MarketId is a free data retrieval call binding the contract method 0x6ed71ede.
//
// Solidity: function marketId() view returns(uint256)
func (_PooledLiquidityMarket *PooledLiquidityMarketCaller) MarketId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PooledLiquidityMarket.contract.Call(opts, &out, "marketId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MarketId is a free data retrieval call binding the contract method 0x6ed71ede.
//
// Solidity: function marketId() view returns(uint256)
func (_PooledLiquidityMarket *PooledLiquidityMarketSession) MarketId() (*big.Int, error) {
	return _PooledLiquidityMarket.Contract.MarketId(&_PooledLiquidityMarket.CallOpts)
}

// MarketId is a free data retrieval call binding the contract method 0x6ed71ede.
//
// Solidity: function marketId() view returns(uint256)
func (_PooledLiquidityMarket *PooledLiquidityMarketCallerSession) MarketId() (*big.Int, error) {
	return _PooledLiquidityMarket.Contract.MarketId(&_PooledLiquidityMarket.CallOpts)
}

// MarketType is a free data retrieval call binding the contract method 0x2dd48909.
//
// Solidity: function marketType() view returns(uint8)
func (_PooledLiquidityMarket *PooledLiquidityMarketCaller) MarketType(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _PooledLiquidityMarket.contract.Call(opts, &out, "marketType")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// MarketType is a free data retrieval call binding the contract method 0x2dd48909.
//
// Solidity: function marketType() view returns(uint8)
func (_PooledLiquidityMarket *PooledLiquidityMarketSession) MarketType() (uint8, error) {
	return _PooledLiquidityMarket.Contract.MarketType(&_PooledLiquidityMarket.CallOpts)
}

// MarketType is a free data retrieval call binding the contract method 0x2dd48909.
//
// Solidity: function marketType() view returns(uint8)
func (_PooledLiquidityMarket *PooledLiquidityMarketCallerSession) MarketType() (uint8, error) {
	return _PooledLiquidityMarket.Contract.MarketType(&_PooledLiquidityMarket.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_PooledLiquidityMarket *PooledLiquidityMarketCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _PooledLiquidityMarket.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_PooledLiquidityMarket *PooledLiquidityMarketSession) Name() (string, error) {
	return _PooledLiquidityMarket.Contract.Name(&_PooledLiquidityMarket.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_PooledLiquidityMarket *PooledLiquidityMarketCallerSession) Name() (string, error) {
	return _PooledLiquidityMarket.Contract.Name(&_PooledLiquidityMarket.CallOpts)
}

// OnERC1155BatchReceived is a free data retrieval call binding the contract method 0xbc197c81.
//
// Solidity: function onERC1155BatchReceived(address , address , uint256[] , uint256[] , bytes ) pure returns(bytes4)
func (_PooledLiquidityMarket *PooledLiquidityMarketCaller) OnERC1155BatchReceived(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address, arg2 []*big.Int, arg3 []*big.Int, arg4 []byte) ([4]byte, error) {
	var out []interface{}
	err := _PooledLiquidityMarket.contract.Call(opts, &out, "onERC1155BatchReceived", arg0, arg1, arg2, arg3, arg4)

	if err != nil {
		return *new([4]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([4]byte)).(*[4]byte)

	return out0, err

}

// OnERC1155BatchReceived is a free data retrieval call binding the contract method 0xbc197c81.
//
// Solidity: function onERC1155BatchReceived(address , address , uint256[] , uint256[] , bytes ) pure returns(bytes4)
func (_PooledLiquidityMarket *PooledLiquidityMarketSession) OnERC1155BatchReceived(arg0 common.Address, arg1 common.Address, arg2 []*big.Int, arg3 []*big.Int, arg4 []byte) ([4]byte, error) {
	return _PooledLiquidityMarket.Contract.OnERC1155BatchReceived(&_PooledLiquidityMarket.CallOpts, arg0, arg1, arg2, arg3, arg4)
}

// OnERC1155BatchReceived is a free data retrieval call binding the contract method 0xbc197c81.
//
// Solidity: function onERC1155BatchReceived(address , address , uint256[] , uint256[] , bytes ) pure returns(bytes4)
func (_PooledLiquidityMarket *PooledLiquidityMarketCallerSession) OnERC1155BatchReceived(arg0 common.Address, arg1 common.Address, arg2 []*big.Int, arg3 []*big.Int, arg4 []byte) ([4]byte, error) {
	return _PooledLiquidityMarket.Contract.OnERC1155BatchReceived(&_PooledLiquidityMarket.CallOpts, arg0, arg1, arg2, arg3, arg4)
}

// OnERC1155Received is a free data retrieval call binding the contract method 0xf23a6e61.
//
// Solidity: function onERC1155Received(address , address , uint256 , uint256 , bytes ) pure returns(bytes4)
func (_PooledLiquidityMarket *PooledLiquidityMarketCaller) OnERC1155Received(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 *big.Int, arg4 []byte) ([4]byte, error) {
	var out []interface{}
	err := _PooledLiquidityMarket.contract.Call(opts, &out, "onERC1155Received", arg0, arg1, arg2, arg3, arg4)

	if err != nil {
		return *new([4]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([4]byte)).(*[4]byte)

	return out0, err

}

// OnERC1155Received is a free data retrieval call binding the contract method 0xf23a6e61.
//
// Solidity: function onERC1155Received(address , address , uint256 , uint256 , bytes ) pure returns(bytes4)
func (_PooledLiquidityMarket *PooledLiquidityMarketSession) OnERC1155Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 *big.Int, arg4 []byte) ([4]byte, error) {
	return _PooledLiquidityMarket.Contract.OnERC1155Received(&_PooledLiquidityMarket.CallOpts, arg0, arg1, arg2, arg3, arg4)
}

// OnERC1155Received is a free data retrieval call binding the contract method 0xf23a6e61.
//
// Solidity: function onERC1155Received(address , address , uint256 , uint256 , bytes ) pure returns(bytes4)
func (_PooledLiquidityMarket *PooledLiquidityMarketCallerSession) OnERC1155Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 *big.Int, arg4 []byte) ([4]byte, error) {
	return _PooledLiquidityMarket.Contract.OnERC1155Received(&_PooledLiquidityMarket.CallOpts, arg0, arg1, arg2, arg3, arg4)
}

// OutcomeCount is a free data retrieval call binding the contract method 0xd300cb31.
//
// Solidity: function outcomeCount() view returns(uint256)
func (_PooledLiquidityMarket *PooledLiquidityMarketCaller) OutcomeCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PooledLiquidityMarket.contract.Call(opts, &out, "outcomeCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// OutcomeCount is a free data retrieval call binding the contract method 0xd300cb31.
//
// Solidity: function outcomeCount() view returns(uint256)
func (_PooledLiquidityMarket *PooledLiquidityMarketSession) OutcomeCount() (*big.Int, error) {
	return _PooledLiquidityMarket.Contract.OutcomeCount(&_PooledLiquidityMarket.CallOpts)
}

// OutcomeCount is a free data retrieval call binding the contract method 0xd300cb31.
//
// Solidity: function outcomeCount() view returns(uint256)
func (_PooledLiquidityMarket *PooledLiquidityMarketCallerSession) OutcomeCount() (*big.Int, error) {
	return _PooledLiquidityMarket.Contract.OutcomeCount(&_PooledLiquidityMarket.CallOpts)
}

// OutcomeToken is a free data retrieval call binding the contract method 0xa998d6d8.
//
// Solidity: function outcomeToken() view returns(address)
func (_PooledLiquidityMarket *PooledLiquidityMarketCaller) OutcomeToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PooledLiquidityMarket.contract.Call(opts, &out, "outcomeToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OutcomeToken is a free data retrieval call binding the contract method 0xa998d6d8.
//
// Solidity: function outcomeToken() view returns(address)
func (_PooledLiquidityMarket *PooledLiquidityMarketSession) OutcomeToken() (common.Address, error) {
	return _PooledLiquidityMarket.Contract.OutcomeToken(&_PooledLiquidityMarket.CallOpts)
}

// OutcomeToken is a free data retrieval call binding the contract method 0xa998d6d8.
//
// Solidity: function outcomeToken() view returns(address)
func (_PooledLiquidityMarket *PooledLiquidityMarketCallerSession) OutcomeToken() (common.Address, error) {
	return _PooledLiquidityMarket.Contract.OutcomeToken(&_PooledLiquidityMarket.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_PooledLiquidityMarket *PooledLiquidityMarketCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _PooledLiquidityMarket.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_PooledLiquidityMarket *PooledLiquidityMarketSession) Paused() (bool, error) {
	return _PooledLiquidityMarket.Contract.Paused(&_PooledLiquidityMarket.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_PooledLiquidityMarket *PooledLiquidityMarketCallerSession) Paused() (bool, error) {
	return _PooledLiquidityMarket.Contract.Paused(&_PooledLiquidityMarket.CallOpts)
}

// PoolState is a free data retrieval call binding the contract method 0x641ad8a9.
//
// Solidity: function poolState() view returns(uint160 sqrtPriceX96, int24 tick, uint128 liquidity, uint256 feeGrowthGlobal0X128, uint256 feeGrowthGlobal1X128)
func (_PooledLiquidityMarket *PooledLiquidityMarketCaller) PoolState(opts *bind.CallOpts) (struct {
	SqrtPriceX96         *big.Int
	Tick                 *big.Int
	Liquidity            *big.Int
	FeeGrowthGlobal0X128 *big.Int
	FeeGrowthGlobal1X128 *big.Int
}, error) {
	var out []interface{}
	err := _PooledLiquidityMarket.contract.Call(opts, &out, "poolState")

	outstruct := new(struct {
		SqrtPriceX96         *big.Int
		Tick                 *big.Int
		Liquidity            *big.Int
		FeeGrowthGlobal0X128 *big.Int
		FeeGrowthGlobal1X128 *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.SqrtPriceX96 = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Tick = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.Liquidity = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.FeeGrowthGlobal0X128 = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.FeeGrowthGlobal1X128 = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// PoolState is a free data retrieval call binding the contract method 0x641ad8a9.
//
// Solidity: function poolState() view returns(uint160 sqrtPriceX96, int24 tick, uint128 liquidity, uint256 feeGrowthGlobal0X128, uint256 feeGrowthGlobal1X128)
func (_PooledLiquidityMarket *PooledLiquidityMarketSession) PoolState() (struct {
	SqrtPriceX96         *big.Int
	Tick                 *big.Int
	Liquidity            *big.Int
	FeeGrowthGlobal0X128 *big.Int
	FeeGrowthGlobal1X128 *big.Int
}, error) {
	return _PooledLiquidityMarket.Contract.PoolState(&_PooledLiquidityMarket.CallOpts)
}

// PoolState is a free data retrieval call binding the contract method 0x641ad8a9.
//
// Solidity: function poolState() view returns(uint160 sqrtPriceX96, int24 tick, uint128 liquidity, uint256 feeGrowthGlobal0X128, uint256 feeGrowthGlobal1X128)
func (_PooledLiquidityMarket *PooledLiquidityMarketCallerSession) PoolState() (struct {
	SqrtPriceX96         *big.Int
	Tick                 *big.Int
	Liquidity            *big.Int
	FeeGrowthGlobal0X128 *big.Int
	FeeGrowthGlobal1X128 *big.Int
}, error) {
	return _PooledLiquidityMarket.Contract.PoolState(&_PooledLiquidityMarket.CallOpts)
}

// Positions is a free data retrieval call binding the contract method 0x514ea4bf.
//
// Solidity: function positions(bytes32 ) view returns(uint128 liquidity, uint256 feeGrowthInside0LastX128, uint256 feeGrowthInside1LastX128, uint128 tokensOwed0, uint128 tokensOwed1)
func (_PooledLiquidityMarket *PooledLiquidityMarketCaller) Positions(opts *bind.CallOpts, arg0 [32]byte) (struct {
	Liquidity                *big.Int
	FeeGrowthInside0LastX128 *big.Int
	FeeGrowthInside1LastX128 *big.Int
	TokensOwed0              *big.Int
	TokensOwed1              *big.Int
}, error) {
	var out []interface{}
	err := _PooledLiquidityMarket.contract.Call(opts, &out, "positions", arg0)

	outstruct := new(struct {
		Liquidity                *big.Int
		FeeGrowthInside0LastX128 *big.Int
		FeeGrowthInside1LastX128 *big.Int
		TokensOwed0              *big.Int
		TokensOwed1              *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Liquidity = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.FeeGrowthInside0LastX128 = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.FeeGrowthInside1LastX128 = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.TokensOwed0 = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.TokensOwed1 = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Positions is a free data retrieval call binding the contract method 0x514ea4bf.
//
// Solidity: function positions(bytes32 ) view returns(uint128 liquidity, uint256 feeGrowthInside0LastX128, uint256 feeGrowthInside1LastX128, uint128 tokensOwed0, uint128 tokensOwed1)
func (_PooledLiquidityMarket *PooledLiquidityMarketSession) Positions(arg0 [32]byte) (struct {
	Liquidity                *big.Int
	FeeGrowthInside0LastX128 *big.Int
	FeeGrowthInside1LastX128 *big.Int
	TokensOwed0              *big.Int
	TokensOwed1              *big.Int
}, error) {
	return _PooledLiquidityMarket.Contract.Positions(&_PooledLiquidityMarket.CallOpts, arg0)
}

// Positions is a free data retrieval call binding the contract method 0x514ea4bf.
//
// Solidity: function positions(bytes32 ) view returns(uint128 liquidity, uint256 feeGrowthInside0LastX128, uint256 feeGrowthInside1LastX128, uint128 tokensOwed0, uint128 tokensOwed1)
func (_PooledLiquidityMarket *PooledLiquidityMarketCallerSession) Positions(arg0 [32]byte) (struct {
	Liquidity                *big.Int
	FeeGrowthInside0LastX128 *big.Int
	FeeGrowthInside1LastX128 *big.Int
	TokensOwed0              *big.Int
	TokensOwed1              *big.Int
}, error) {
	return _PooledLiquidityMarket.Contract.Positions(&_PooledLiquidityMarket.CallOpts, arg0)
}

// Reserves is a free data retrieval call binding the contract method 0x8334278d.
//
// Solidity: function reserves(uint256 ) view returns(uint256)
func (_PooledLiquidityMarket *PooledLiquidityMarketCaller) Reserves(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _PooledLiquidityMarket.contract.Call(opts, &out, "reserves", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Reserves is a free data retrieval call binding the contract method 0x8334278d.
//
// Solidity: function reserves(uint256 ) view returns(uint256)
func (_PooledLiquidityMarket *PooledLiquidityMarketSession) Reserves(arg0 *big.Int) (*big.Int, error) {
	return _PooledLiquidityMarket.Contract.Reserves(&_PooledLiquidityMarket.CallOpts, arg0)
}

// Reserves is a free data retrieval call binding the contract method 0x8334278d.
//
// Solidity: function reserves(uint256 ) view returns(uint256)
func (_PooledLiquidityMarket *PooledLiquidityMarketCallerSession) Reserves(arg0 *big.Int) (*big.Int, error) {
	return _PooledLiquidityMarket.Contract.Reserves(&_PooledLiquidityMarket.CallOpts, arg0)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) pure returns(bool)
func (_PooledLiquidityMarket *PooledLiquidityMarketCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _PooledLiquidityMarket.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) pure returns(bool)
func (_PooledLiquidityMarket *PooledLiquidityMarketSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _PooledLiquidityMarket.Contract.SupportsInterface(&_PooledLiquidityMarket.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) pure returns(bool)
func (_PooledLiquidityMarket *PooledLiquidityMarketCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _PooledLiquidityMarket.Contract.SupportsInterface(&_PooledLiquidityMarket.CallOpts, interfaceId)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_PooledLiquidityMarket *PooledLiquidityMarketCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _PooledLiquidityMarket.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_PooledLiquidityMarket *PooledLiquidityMarketSession) Symbol() (string, error) {
	return _PooledLiquidityMarket.Contract.Symbol(&_PooledLiquidityMarket.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_PooledLiquidityMarket *PooledLiquidityMarketCallerSession) Symbol() (string, error) {
	return _PooledLiquidityMarket.Contract.Symbol(&_PooledLiquidityMarket.CallOpts)
}

// Ticks is a free data retrieval call binding the contract method 0xf30dba93.
//
// Solidity: function ticks(int24 ) view returns(uint128 liquidityGross, int128 liquidityNet, uint256 feeGrowthOutside0X128, uint256 feeGrowthOutside1X128, bool initialized)
func (_PooledLiquidityMarket *PooledLiquidityMarketCaller) Ticks(opts *bind.CallOpts, arg0 *big.Int) (struct {
	LiquidityGross        *big.Int
	LiquidityNet          *big.Int
	FeeGrowthOutside0X128 *big.Int
	FeeGrowthOutside1X128 *big.Int
	Initialized           bool
}, error) {
	var out []interface{}
	err := _PooledLiquidityMarket.contract.Call(opts, &out, "ticks", arg0)

	outstruct := new(struct {
		LiquidityGross        *big.Int
		LiquidityNet          *big.Int
		FeeGrowthOutside0X128 *big.Int
		FeeGrowthOutside1X128 *big.Int
		Initialized           bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.LiquidityGross = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.LiquidityNet = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.FeeGrowthOutside0X128 = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.FeeGrowthOutside1X128 = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.Initialized = *abi.ConvertType(out[4], new(bool)).(*bool)

	return *outstruct, err

}

// Ticks is a free data retrieval call binding the contract method 0xf30dba93.
//
// Solidity: function ticks(int24 ) view returns(uint128 liquidityGross, int128 liquidityNet, uint256 feeGrowthOutside0X128, uint256 feeGrowthOutside1X128, bool initialized)
func (_PooledLiquidityMarket *PooledLiquidityMarketSession) Ticks(arg0 *big.Int) (struct {
	LiquidityGross        *big.Int
	LiquidityNet          *big.Int
	FeeGrowthOutside0X128 *big.Int
	FeeGrowthOutside1X128 *big.Int
	Initialized           bool
}, error) {
	return _PooledLiquidityMarket.Contract.Ticks(&_PooledLiquidityMarket.CallOpts, arg0)
}

// Ticks is a free data retrieval call binding the contract method 0xf30dba93.
//
// Solidity: function ticks(int24 ) view returns(uint128 liquidityGross, int128 liquidityNet, uint256 feeGrowthOutside0X128, uint256 feeGrowthOutside1X128, bool initialized)
func (_PooledLiquidityMarket *PooledLiquidityMarketCallerSession) Ticks(arg0 *big.Int) (struct {
	LiquidityGross        *big.Int
	LiquidityNet          *big.Int
	FeeGrowthOutside0X128 *big.Int
	FeeGrowthOutside1X128 *big.Int
	Initialized           bool
}, error) {
	return _PooledLiquidityMarket.Contract.Ticks(&_PooledLiquidityMarket.CallOpts, arg0)
}

// TotalCollateral is a free data retrieval call binding the contract method 0x4ac8eb5f.
//
// Solidity: function totalCollateral() view returns(uint256)
func (_PooledLiquidityMarket *PooledLiquidityMarketCaller) TotalCollateral(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PooledLiquidityMarket.contract.Call(opts, &out, "totalCollateral")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalCollateral is a free data retrieval call binding the contract method 0x4ac8eb5f.
//
// Solidity: function totalCollateral() view returns(uint256)
func (_PooledLiquidityMarket *PooledLiquidityMarketSession) TotalCollateral() (*big.Int, error) {
	return _PooledLiquidityMarket.Contract.TotalCollateral(&_PooledLiquidityMarket.CallOpts)
}

// TotalCollateral is a free data retrieval call binding the contract method 0x4ac8eb5f.
//
// Solidity: function totalCollateral() view returns(uint256)
func (_PooledLiquidityMarket *PooledLiquidityMarketCallerSession) TotalCollateral() (*big.Int, error) {
	return _PooledLiquidityMarket.Contract.TotalCollateral(&_PooledLiquidityMarket.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_PooledLiquidityMarket *PooledLiquidityMarketCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PooledLiquidityMarket.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_PooledLiquidityMarket *PooledLiquidityMarketSession) TotalSupply() (*big.Int, error) {
	return _PooledLiquidityMarket.Contract.TotalSupply(&_PooledLiquidityMarket.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_PooledLiquidityMarket *PooledLiquidityMarketCallerSession) TotalSupply() (*big.Int, error) {
	return _PooledLiquidityMarket.Contract.TotalSupply(&_PooledLiquidityMarket.CallOpts)
}

// AddLiquidity is a paid mutator transaction binding the contract method 0x51c6590a.
//
// Solidity: function addLiquidity(uint256 amount) returns(uint256 lpTokens)
func (_PooledLiquidityMarket *PooledLiquidityMarketTransactor) AddLiquidity(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _PooledLiquidityMarket.contract.Transact(opts, "addLiquidity", amount)
}

// AddLiquidity is a paid mutator transaction binding the contract method 0x51c6590a.
//
// Solidity: function addLiquidity(uint256 amount) returns(uint256 lpTokens)
func (_PooledLiquidityMarket *PooledLiquidityMarketSession) AddLiquidity(amount *big.Int) (*types.Transaction, error) {
	return _PooledLiquidityMarket.Contract.AddLiquidity(&_PooledLiquidityMarket.TransactOpts, amount)
}

// AddLiquidity is a paid mutator transaction binding the contract method 0x51c6590a.
//
// Solidity: function addLiquidity(uint256 amount) returns(uint256 lpTokens)
func (_PooledLiquidityMarket *PooledLiquidityMarketTransactorSession) AddLiquidity(amount *big.Int) (*types.Transaction, error) {
	return _PooledLiquidityMarket.Contract.AddLiquidity(&_PooledLiquidityMarket.TransactOpts, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_PooledLiquidityMarket *PooledLiquidityMarketTransactor) Approve(opts *bind.TransactOpts, spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _PooledLiquidityMarket.contract.Transact(opts, "approve", spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_PooledLiquidityMarket *PooledLiquidityMarketSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _PooledLiquidityMarket.Contract.Approve(&_PooledLiquidityMarket.TransactOpts, spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_PooledLiquidityMarket *PooledLiquidityMarketTransactorSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _PooledLiquidityMarket.Contract.Approve(&_PooledLiquidityMarket.TransactOpts, spender, value)
}

// BurnPosition is a paid mutator transaction binding the contract method 0x17fc496e.
//
// Solidity: function burnPosition(int24 tickLower, int24 tickUpper, uint128 liquidityToBurn) returns(uint256 amount0, uint256 amount1)
func (_PooledLiquidityMarket *PooledLiquidityMarketTransactor) BurnPosition(opts *bind.TransactOpts, tickLower *big.Int, tickUpper *big.Int, liquidityToBurn *big.Int) (*types.Transaction, error) {
	return _PooledLiquidityMarket.contract.Transact(opts, "burnPosition", tickLower, tickUpper, liquidityToBurn)
}

// BurnPosition is a paid mutator transaction binding the contract method 0x17fc496e.
//
// Solidity: function burnPosition(int24 tickLower, int24 tickUpper, uint128 liquidityToBurn) returns(uint256 amount0, uint256 amount1)
func (_PooledLiquidityMarket *PooledLiquidityMarketSession) BurnPosition(tickLower *big.Int, tickUpper *big.Int, liquidityToBurn *big.Int) (*types.Transaction, error) {
	return _PooledLiquidityMarket.Contract.BurnPosition(&_PooledLiquidityMarket.TransactOpts, tickLower, tickUpper, liquidityToBurn)
}

// BurnPosition is a paid mutator transaction binding the contract method 0x17fc496e.
//
// Solidity: function burnPosition(int24 tickLower, int24 tickUpper, uint128 liquidityToBurn) returns(uint256 amount0, uint256 amount1)
func (_PooledLiquidityMarket *PooledLiquidityMarketTransactorSession) BurnPosition(tickLower *big.Int, tickUpper *big.Int, liquidityToBurn *big.Int) (*types.Transaction, error) {
	return _PooledLiquidityMarket.Contract.BurnPosition(&_PooledLiquidityMarket.TransactOpts, tickLower, tickUpper, liquidityToBurn)
}

// Buy is a paid mutator transaction binding the contract method 0x40993b26.
//
// Solidity: function buy(uint256 outcomeId, uint256 collateralIn, uint256 minTokensOut) returns(uint256 tokensOut)
func (_PooledLiquidityMarket *PooledLiquidityMarketTransactor) Buy(opts *bind.TransactOpts, outcomeId *big.Int, collateralIn *big.Int, minTokensOut *big.Int) (*types.Transaction, error) {
	return _PooledLiquidityMarket.contract.Transact(opts, "buy", outcomeId, collateralIn, minTokensOut)
}

// Buy is a paid mutator transaction binding the contract method 0x40993b26.
//
// Solidity: function buy(uint256 outcomeId, uint256 collateralIn, uint256 minTokensOut) returns(uint256 tokensOut)
func (_PooledLiquidityMarket *PooledLiquidityMarketSession) Buy(outcomeId *big.Int, collateralIn *big.Int, minTokensOut *big.Int) (*types.Transaction, error) {
	return _PooledLiquidityMarket.Contract.Buy(&_PooledLiquidityMarket.TransactOpts, outcomeId, collateralIn, minTokensOut)
}

// Buy is a paid mutator transaction binding the contract method 0x40993b26.
//
// Solidity: function buy(uint256 outcomeId, uint256 collateralIn, uint256 minTokensOut) returns(uint256 tokensOut)
func (_PooledLiquidityMarket *PooledLiquidityMarketTransactorSession) Buy(outcomeId *big.Int, collateralIn *big.Int, minTokensOut *big.Int) (*types.Transaction, error) {
	return _PooledLiquidityMarket.Contract.Buy(&_PooledLiquidityMarket.TransactOpts, outcomeId, collateralIn, minTokensOut)
}

// CollectFees is a paid mutator transaction binding the contract method 0xf4409308.
//
// Solidity: function collectFees(int24 tickLower, int24 tickUpper) returns(uint128 amount0, uint128 amount1)
func (_PooledLiquidityMarket *PooledLiquidityMarketTransactor) CollectFees(opts *bind.TransactOpts, tickLower *big.Int, tickUpper *big.Int) (*types.Transaction, error) {
	return _PooledLiquidityMarket.contract.Transact(opts, "collectFees", tickLower, tickUpper)
}

// CollectFees is a paid mutator transaction binding the contract method 0xf4409308.
//
// Solidity: function collectFees(int24 tickLower, int24 tickUpper) returns(uint128 amount0, uint128 amount1)
func (_PooledLiquidityMarket *PooledLiquidityMarketSession) CollectFees(tickLower *big.Int, tickUpper *big.Int) (*types.Transaction, error) {
	return _PooledLiquidityMarket.Contract.CollectFees(&_PooledLiquidityMarket.TransactOpts, tickLower, tickUpper)
}

// CollectFees is a paid mutator transaction binding the contract method 0xf4409308.
//
// Solidity: function collectFees(int24 tickLower, int24 tickUpper) returns(uint128 amount0, uint128 amount1)
func (_PooledLiquidityMarket *PooledLiquidityMarketTransactorSession) CollectFees(tickLower *big.Int, tickUpper *big.Int) (*types.Transaction, error) {
	return _PooledLiquidityMarket.Contract.CollectFees(&_PooledLiquidityMarket.TransactOpts, tickLower, tickUpper)
}

// FundRedemptions is a paid mutator transaction binding the contract method 0x281155ba.
//
// Solidity: function fundRedemptions() returns()
func (_PooledLiquidityMarket *PooledLiquidityMarketTransactor) FundRedemptions(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PooledLiquidityMarket.contract.Transact(opts, "fundRedemptions")
}

// FundRedemptions is a paid mutator transaction binding the contract method 0x281155ba.
//
// Solidity: function fundRedemptions() returns()
func (_PooledLiquidityMarket *PooledLiquidityMarketSession) FundRedemptions() (*types.Transaction, error) {
	return _PooledLiquidityMarket.Contract.FundRedemptions(&_PooledLiquidityMarket.TransactOpts)
}

// FundRedemptions is a paid mutator transaction binding the contract method 0x281155ba.
//
// Solidity: function fundRedemptions() returns()
func (_PooledLiquidityMarket *PooledLiquidityMarketTransactorSession) FundRedemptions() (*types.Transaction, error) {
	return _PooledLiquidityMarket.Contract.FundRedemptions(&_PooledLiquidityMarket.TransactOpts)
}

// MintPosition is a paid mutator transaction binding the contract method 0x24c274fc.
//
// Solidity: function mintPosition(int24 tickLower, int24 tickUpper, uint128 liquidityDesired, uint256 amount0Max, uint256 amount1Max) returns(uint128 liquidity, uint256 amount0, uint256 amount1)
func (_PooledLiquidityMarket *PooledLiquidityMarketTransactor) MintPosition(opts *bind.TransactOpts, tickLower *big.Int, tickUpper *big.Int, liquidityDesired *big.Int, amount0Max *big.Int, amount1Max *big.Int) (*types.Transaction, error) {
	return _PooledLiquidityMarket.contract.Transact(opts, "mintPosition", tickLower, tickUpper, liquidityDesired, amount0Max, amount1Max)
}

// MintPosition is a paid mutator transaction binding the contract method 0x24c274fc.
//
// Solidity: function mintPosition(int24 tickLower, int24 tickUpper, uint128 liquidityDesired, uint256 amount0Max, uint256 amount1Max) returns(uint128 liquidity, uint256 amount0, uint256 amount1)
func (_PooledLiquidityMarket *PooledLiquidityMarketSession) MintPosition(tickLower *big.Int, tickUpper *big.Int, liquidityDesired *big.Int, amount0Max *big.Int, amount1Max *big.Int) (*types.Transaction, error) {
	return _PooledLiquidityMarket.Contract.MintPosition(&_PooledLiquidityMarket.TransactOpts, tickLower, tickUpper, liquidityDesired, amount0Max, amount1Max)
}

// MintPosition is a paid mutator transaction binding the contract method 0x24c274fc.
//
// Solidity: function mintPosition(int24 tickLower, int24 tickUpper, uint128 liquidityDesired, uint256 amount0Max, uint256 amount1Max) returns(uint128 liquidity, uint256 amount0, uint256 amount1)
func (_PooledLiquidityMarket *PooledLiquidityMarketTransactorSession) MintPosition(tickLower *big.Int, tickUpper *big.Int, liquidityDesired *big.Int, amount0Max *big.Int, amount1Max *big.Int) (*types.Transaction, error) {
	return _PooledLiquidityMarket.Contract.MintPosition(&_PooledLiquidityMarket.TransactOpts, tickLower, tickUpper, liquidityDesired, amount0Max, amount1Max)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_PooledLiquidityMarket *PooledLiquidityMarketTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PooledLiquidityMarket.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_PooledLiquidityMarket *PooledLiquidityMarketSession) Pause() (*types.Transaction, error) {
	return _PooledLiquidityMarket.Contract.Pause(&_PooledLiquidityMarket.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_PooledLiquidityMarket *PooledLiquidityMarketTransactorSession) Pause() (*types.Transaction, error) {
	return _PooledLiquidityMarket.Contract.Pause(&_PooledLiquidityMarket.TransactOpts)
}

// RemoveLiquidity is a paid mutator transaction binding the contract method 0x9c8f9f23.
//
// Solidity: function removeLiquidity(uint256 lpTokens) returns(uint256 collateralOut)
func (_PooledLiquidityMarket *PooledLiquidityMarketTransactor) RemoveLiquidity(opts *bind.TransactOpts, lpTokens *big.Int) (*types.Transaction, error) {
	return _PooledLiquidityMarket.contract.Transact(opts, "removeLiquidity", lpTokens)
}

// RemoveLiquidity is a paid mutator transaction binding the contract method 0x9c8f9f23.
//
// Solidity: function removeLiquidity(uint256 lpTokens) returns(uint256 collateralOut)
func (_PooledLiquidityMarket *PooledLiquidityMarketSession) RemoveLiquidity(lpTokens *big.Int) (*types.Transaction, error) {
	return _PooledLiquidityMarket.Contract.RemoveLiquidity(&_PooledLiquidityMarket.TransactOpts, lpTokens)
}

// RemoveLiquidity is a paid mutator transaction binding the contract method 0x9c8f9f23.
//
// Solidity: function removeLiquidity(uint256 lpTokens) returns(uint256 collateralOut)
func (_PooledLiquidityMarket *PooledLiquidityMarketTransactorSession) RemoveLiquidity(lpTokens *big.Int) (*types.Transaction, error) {
	return _PooledLiquidityMarket.Contract.RemoveLiquidity(&_PooledLiquidityMarket.TransactOpts, lpTokens)
}

// Sell is a paid mutator transaction binding the contract method 0xd3c9727c.
//
// Solidity: function sell(uint256 outcomeId, uint256 tokensIn, uint256 minCollateralOut) returns(uint256 collateralOut)
func (_PooledLiquidityMarket *PooledLiquidityMarketTransactor) Sell(opts *bind.TransactOpts, outcomeId *big.Int, tokensIn *big.Int, minCollateralOut *big.Int) (*types.Transaction, error) {
	return _PooledLiquidityMarket.contract.Transact(opts, "sell", outcomeId, tokensIn, minCollateralOut)
}

// Sell is a paid mutator transaction binding the contract method 0xd3c9727c.
//
// Solidity: function sell(uint256 outcomeId, uint256 tokensIn, uint256 minCollateralOut) returns(uint256 collateralOut)
func (_PooledLiquidityMarket *PooledLiquidityMarketSession) Sell(outcomeId *big.Int, tokensIn *big.Int, minCollateralOut *big.Int) (*types.Transaction, error) {
	return _PooledLiquidityMarket.Contract.Sell(&_PooledLiquidityMarket.TransactOpts, outcomeId, tokensIn, minCollateralOut)
}

// Sell is a paid mutator transaction binding the contract method 0xd3c9727c.
//
// Solidity: function sell(uint256 outcomeId, uint256 tokensIn, uint256 minCollateralOut) returns(uint256 collateralOut)
func (_PooledLiquidityMarket *PooledLiquidityMarketTransactorSession) Sell(outcomeId *big.Int, tokensIn *big.Int, minCollateralOut *big.Int) (*types.Transaction, error) {
	return _PooledLiquidityMarket.Contract.Sell(&_PooledLiquidityMarket.TransactOpts, outcomeId, tokensIn, minCollateralOut)
}

// SetAdmin is a paid mutator transaction binding the contract method 0x704b6c02.
//
// Solidity: function setAdmin(address newAdmin) returns()
func (_PooledLiquidityMarket *PooledLiquidityMarketTransactor) SetAdmin(opts *bind.TransactOpts, newAdmin common.Address) (*types.Transaction, error) {
	return _PooledLiquidityMarket.contract.Transact(opts, "setAdmin", newAdmin)
}

// SetAdmin is a paid mutator transaction binding the contract method 0x704b6c02.
//
// Solidity: function setAdmin(address newAdmin) returns()
func (_PooledLiquidityMarket *PooledLiquidityMarketSession) SetAdmin(newAdmin common.Address) (*types.Transaction, error) {
	return _PooledLiquidityMarket.Contract.SetAdmin(&_PooledLiquidityMarket.TransactOpts, newAdmin)
}

// SetAdmin is a paid mutator transaction binding the contract method 0x704b6c02.
//
// Solidity: function setAdmin(address newAdmin) returns()
func (_PooledLiquidityMarket *PooledLiquidityMarketTransactorSession) SetAdmin(newAdmin common.Address) (*types.Transaction, error) {
	return _PooledLiquidityMarket.Contract.SetAdmin(&_PooledLiquidityMarket.TransactOpts, newAdmin)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_PooledLiquidityMarket *PooledLiquidityMarketTransactor) Transfer(opts *bind.TransactOpts, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _PooledLiquidityMarket.contract.Transact(opts, "transfer", to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_PooledLiquidityMarket *PooledLiquidityMarketSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _PooledLiquidityMarket.Contract.Transfer(&_PooledLiquidityMarket.TransactOpts, to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_PooledLiquidityMarket *PooledLiquidityMarketTransactorSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _PooledLiquidityMarket.Contract.Transfer(&_PooledLiquidityMarket.TransactOpts, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_PooledLiquidityMarket *PooledLiquidityMarketTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _PooledLiquidityMarket.contract.Transact(opts, "transferFrom", from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_PooledLiquidityMarket *PooledLiquidityMarketSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _PooledLiquidityMarket.Contract.TransferFrom(&_PooledLiquidityMarket.TransactOpts, from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_PooledLiquidityMarket *PooledLiquidityMarketTransactorSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _PooledLiquidityMarket.Contract.TransferFrom(&_PooledLiquidityMarket.TransactOpts, from, to, value)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_PooledLiquidityMarket *PooledLiquidityMarketTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PooledLiquidityMarket.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_PooledLiquidityMarket *PooledLiquidityMarketSession) Unpause() (*types.Transaction, error) {
	return _PooledLiquidityMarket.Contract.Unpause(&_PooledLiquidityMarket.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_PooledLiquidityMarket *PooledLiquidityMarketTransactorSession) Unpause() (*types.Transaction, error) {
	return _PooledLiquidityMarket.Contract.Unpause(&_PooledLiquidityMarket.TransactOpts)
}

// PooledLiquidityMarketApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the PooledLiquidityMarket contract.
type PooledLiquidityMarketApprovalIterator struct {
	Event *PooledLiquidityMarketApproval // Event containing the contract specifics and raw log

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
func (it *PooledLiquidityMarketApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PooledLiquidityMarketApproval)
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
		it.Event = new(PooledLiquidityMarketApproval)
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
func (it *PooledLiquidityMarketApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PooledLiquidityMarketApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PooledLiquidityMarketApproval represents a Approval event raised by the PooledLiquidityMarket contract.
type PooledLiquidityMarketApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_PooledLiquidityMarket *PooledLiquidityMarketFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*PooledLiquidityMarketApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _PooledLiquidityMarket.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &PooledLiquidityMarketApprovalIterator{contract: _PooledLiquidityMarket.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_PooledLiquidityMarket *PooledLiquidityMarketFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *PooledLiquidityMarketApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _PooledLiquidityMarket.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PooledLiquidityMarketApproval)
				if err := _PooledLiquidityMarket.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_PooledLiquidityMarket *PooledLiquidityMarketFilterer) ParseApproval(log types.Log) (*PooledLiquidityMarketApproval, error) {
	event := new(PooledLiquidityMarketApproval)
	if err := _PooledLiquidityMarket.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PooledLiquidityMarketFeesCollectedIterator is returned from FilterFeesCollected and is used to iterate over the raw logs and unpacked data for FeesCollected events raised by the PooledLiquidityMarket contract.
type PooledLiquidityMarketFeesCollectedIterator struct {
	Event *PooledLiquidityMarketFeesCollected // Event containing the contract specifics and raw log

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
func (it *PooledLiquidityMarketFeesCollectedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PooledLiquidityMarketFeesCollected)
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
		it.Event = new(PooledLiquidityMarketFeesCollected)
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
func (it *PooledLiquidityMarketFeesCollectedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PooledLiquidityMarketFeesCollectedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PooledLiquidityMarketFeesCollected represents a FeesCollected event raised by the PooledLiquidityMarket contract.
type PooledLiquidityMarketFeesCollected struct {
	Owner     common.Address
	TickLower *big.Int
	TickUpper *big.Int
	Amount0   *big.Int
	Amount1   *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterFeesCollected is a free log retrieval operation binding the contract event 0xadd618188d46b342ce3b418805cd09ce9047cd753595a117c0f2764e98637d69.
//
// Solidity: event FeesCollected(address indexed owner, int24 indexed tickLower, int24 indexed tickUpper, uint128 amount0, uint128 amount1)
func (_PooledLiquidityMarket *PooledLiquidityMarketFilterer) FilterFeesCollected(opts *bind.FilterOpts, owner []common.Address, tickLower []*big.Int, tickUpper []*big.Int) (*PooledLiquidityMarketFeesCollectedIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var tickLowerRule []interface{}
	for _, tickLowerItem := range tickLower {
		tickLowerRule = append(tickLowerRule, tickLowerItem)
	}
	var tickUpperRule []interface{}
	for _, tickUpperItem := range tickUpper {
		tickUpperRule = append(tickUpperRule, tickUpperItem)
	}

	logs, sub, err := _PooledLiquidityMarket.contract.FilterLogs(opts, "FeesCollected", ownerRule, tickLowerRule, tickUpperRule)
	if err != nil {
		return nil, err
	}
	return &PooledLiquidityMarketFeesCollectedIterator{contract: _PooledLiquidityMarket.contract, event: "FeesCollected", logs: logs, sub: sub}, nil
}

// WatchFeesCollected is a free log subscription operation binding the contract event 0xadd618188d46b342ce3b418805cd09ce9047cd753595a117c0f2764e98637d69.
//
// Solidity: event FeesCollected(address indexed owner, int24 indexed tickLower, int24 indexed tickUpper, uint128 amount0, uint128 amount1)
func (_PooledLiquidityMarket *PooledLiquidityMarketFilterer) WatchFeesCollected(opts *bind.WatchOpts, sink chan<- *PooledLiquidityMarketFeesCollected, owner []common.Address, tickLower []*big.Int, tickUpper []*big.Int) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var tickLowerRule []interface{}
	for _, tickLowerItem := range tickLower {
		tickLowerRule = append(tickLowerRule, tickLowerItem)
	}
	var tickUpperRule []interface{}
	for _, tickUpperItem := range tickUpper {
		tickUpperRule = append(tickUpperRule, tickUpperItem)
	}

	logs, sub, err := _PooledLiquidityMarket.contract.WatchLogs(opts, "FeesCollected", ownerRule, tickLowerRule, tickUpperRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PooledLiquidityMarketFeesCollected)
				if err := _PooledLiquidityMarket.contract.UnpackLog(event, "FeesCollected", log); err != nil {
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

// ParseFeesCollected is a log parse operation binding the contract event 0xadd618188d46b342ce3b418805cd09ce9047cd753595a117c0f2764e98637d69.
//
// Solidity: event FeesCollected(address indexed owner, int24 indexed tickLower, int24 indexed tickUpper, uint128 amount0, uint128 amount1)
func (_PooledLiquidityMarket *PooledLiquidityMarketFilterer) ParseFeesCollected(log types.Log) (*PooledLiquidityMarketFeesCollected, error) {
	event := new(PooledLiquidityMarketFeesCollected)
	if err := _PooledLiquidityMarket.contract.UnpackLog(event, "FeesCollected", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PooledLiquidityMarketLiquidityChangedIterator is returned from FilterLiquidityChanged and is used to iterate over the raw logs and unpacked data for LiquidityChanged events raised by the PooledLiquidityMarket contract.
type PooledLiquidityMarketLiquidityChangedIterator struct {
	Event *PooledLiquidityMarketLiquidityChanged // Event containing the contract specifics and raw log

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
func (it *PooledLiquidityMarketLiquidityChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PooledLiquidityMarketLiquidityChanged)
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
		it.Event = new(PooledLiquidityMarketLiquidityChanged)
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
func (it *PooledLiquidityMarketLiquidityChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PooledLiquidityMarketLiquidityChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PooledLiquidityMarketLiquidityChanged represents a LiquidityChanged event raised by the PooledLiquidityMarket contract.
type PooledLiquidityMarketLiquidityChanged struct {
	Provider   common.Address
	Amount     *big.Int
	IsAddition bool
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterLiquidityChanged is a free log retrieval operation binding the contract event 0xb029a6414a0c6d2e4fa2e5287326aa8a8c7191f9f5ced9799754a380471458d4.
//
// Solidity: event LiquidityChanged(address indexed provider, uint256 amount, bool isAddition)
func (_PooledLiquidityMarket *PooledLiquidityMarketFilterer) FilterLiquidityChanged(opts *bind.FilterOpts, provider []common.Address) (*PooledLiquidityMarketLiquidityChangedIterator, error) {

	var providerRule []interface{}
	for _, providerItem := range provider {
		providerRule = append(providerRule, providerItem)
	}

	logs, sub, err := _PooledLiquidityMarket.contract.FilterLogs(opts, "LiquidityChanged", providerRule)
	if err != nil {
		return nil, err
	}
	return &PooledLiquidityMarketLiquidityChangedIterator{contract: _PooledLiquidityMarket.contract, event: "LiquidityChanged", logs: logs, sub: sub}, nil
}

// WatchLiquidityChanged is a free log subscription operation binding the contract event 0xb029a6414a0c6d2e4fa2e5287326aa8a8c7191f9f5ced9799754a380471458d4.
//
// Solidity: event LiquidityChanged(address indexed provider, uint256 amount, bool isAddition)
func (_PooledLiquidityMarket *PooledLiquidityMarketFilterer) WatchLiquidityChanged(opts *bind.WatchOpts, sink chan<- *PooledLiquidityMarketLiquidityChanged, provider []common.Address) (event.Subscription, error) {

	var providerRule []interface{}
	for _, providerItem := range provider {
		providerRule = append(providerRule, providerItem)
	}

	logs, sub, err := _PooledLiquidityMarket.contract.WatchLogs(opts, "LiquidityChanged", providerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PooledLiquidityMarketLiquidityChanged)
				if err := _PooledLiquidityMarket.contract.UnpackLog(event, "LiquidityChanged", log); err != nil {
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
func (_PooledLiquidityMarket *PooledLiquidityMarketFilterer) ParseLiquidityChanged(log types.Log) (*PooledLiquidityMarketLiquidityChanged, error) {
	event := new(PooledLiquidityMarketLiquidityChanged)
	if err := _PooledLiquidityMarket.contract.UnpackLog(event, "LiquidityChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PooledLiquidityMarketPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the PooledLiquidityMarket contract.
type PooledLiquidityMarketPausedIterator struct {
	Event *PooledLiquidityMarketPaused // Event containing the contract specifics and raw log

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
func (it *PooledLiquidityMarketPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PooledLiquidityMarketPaused)
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
		it.Event = new(PooledLiquidityMarketPaused)
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
func (it *PooledLiquidityMarketPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PooledLiquidityMarketPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PooledLiquidityMarketPaused represents a Paused event raised by the PooledLiquidityMarket contract.
type PooledLiquidityMarketPaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_PooledLiquidityMarket *PooledLiquidityMarketFilterer) FilterPaused(opts *bind.FilterOpts) (*PooledLiquidityMarketPausedIterator, error) {

	logs, sub, err := _PooledLiquidityMarket.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &PooledLiquidityMarketPausedIterator{contract: _PooledLiquidityMarket.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_PooledLiquidityMarket *PooledLiquidityMarketFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *PooledLiquidityMarketPaused) (event.Subscription, error) {

	logs, sub, err := _PooledLiquidityMarket.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PooledLiquidityMarketPaused)
				if err := _PooledLiquidityMarket.contract.UnpackLog(event, "Paused", log); err != nil {
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
func (_PooledLiquidityMarket *PooledLiquidityMarketFilterer) ParsePaused(log types.Log) (*PooledLiquidityMarketPaused, error) {
	event := new(PooledLiquidityMarketPaused)
	if err := _PooledLiquidityMarket.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PooledLiquidityMarketPositionBurnedIterator is returned from FilterPositionBurned and is used to iterate over the raw logs and unpacked data for PositionBurned events raised by the PooledLiquidityMarket contract.
type PooledLiquidityMarketPositionBurnedIterator struct {
	Event *PooledLiquidityMarketPositionBurned // Event containing the contract specifics and raw log

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
func (it *PooledLiquidityMarketPositionBurnedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PooledLiquidityMarketPositionBurned)
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
		it.Event = new(PooledLiquidityMarketPositionBurned)
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
func (it *PooledLiquidityMarketPositionBurnedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PooledLiquidityMarketPositionBurnedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PooledLiquidityMarketPositionBurned represents a PositionBurned event raised by the PooledLiquidityMarket contract.
type PooledLiquidityMarketPositionBurned struct {
	Owner     common.Address
	TickLower *big.Int
	TickUpper *big.Int
	Liquidity *big.Int
	Amount0   *big.Int
	Amount1   *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterPositionBurned is a free log retrieval operation binding the contract event 0x19e1110dbc679245edfedb549b03aa54ed30d802c400c53f91ec74d395360826.
//
// Solidity: event PositionBurned(address indexed owner, int24 indexed tickLower, int24 indexed tickUpper, uint128 liquidity, uint256 amount0, uint256 amount1)
func (_PooledLiquidityMarket *PooledLiquidityMarketFilterer) FilterPositionBurned(opts *bind.FilterOpts, owner []common.Address, tickLower []*big.Int, tickUpper []*big.Int) (*PooledLiquidityMarketPositionBurnedIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var tickLowerRule []interface{}
	for _, tickLowerItem := range tickLower {
		tickLowerRule = append(tickLowerRule, tickLowerItem)
	}
	var tickUpperRule []interface{}
	for _, tickUpperItem := range tickUpper {
		tickUpperRule = append(tickUpperRule, tickUpperItem)
	}

	logs, sub, err := _PooledLiquidityMarket.contract.FilterLogs(opts, "PositionBurned", ownerRule, tickLowerRule, tickUpperRule)
	if err != nil {
		return nil, err
	}
	return &PooledLiquidityMarketPositionBurnedIterator{contract: _PooledLiquidityMarket.contract, event: "PositionBurned", logs: logs, sub: sub}, nil
}

// WatchPositionBurned is a free log subscription operation binding the contract event 0x19e1110dbc679245edfedb549b03aa54ed30d802c400c53f91ec74d395360826.
//
// Solidity: event PositionBurned(address indexed owner, int24 indexed tickLower, int24 indexed tickUpper, uint128 liquidity, uint256 amount0, uint256 amount1)
func (_PooledLiquidityMarket *PooledLiquidityMarketFilterer) WatchPositionBurned(opts *bind.WatchOpts, sink chan<- *PooledLiquidityMarketPositionBurned, owner []common.Address, tickLower []*big.Int, tickUpper []*big.Int) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var tickLowerRule []interface{}
	for _, tickLowerItem := range tickLower {
		tickLowerRule = append(tickLowerRule, tickLowerItem)
	}
	var tickUpperRule []interface{}
	for _, tickUpperItem := range tickUpper {
		tickUpperRule = append(tickUpperRule, tickUpperItem)
	}

	logs, sub, err := _PooledLiquidityMarket.contract.WatchLogs(opts, "PositionBurned", ownerRule, tickLowerRule, tickUpperRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PooledLiquidityMarketPositionBurned)
				if err := _PooledLiquidityMarket.contract.UnpackLog(event, "PositionBurned", log); err != nil {
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

// ParsePositionBurned is a log parse operation binding the contract event 0x19e1110dbc679245edfedb549b03aa54ed30d802c400c53f91ec74d395360826.
//
// Solidity: event PositionBurned(address indexed owner, int24 indexed tickLower, int24 indexed tickUpper, uint128 liquidity, uint256 amount0, uint256 amount1)
func (_PooledLiquidityMarket *PooledLiquidityMarketFilterer) ParsePositionBurned(log types.Log) (*PooledLiquidityMarketPositionBurned, error) {
	event := new(PooledLiquidityMarketPositionBurned)
	if err := _PooledLiquidityMarket.contract.UnpackLog(event, "PositionBurned", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PooledLiquidityMarketPositionMintedIterator is returned from FilterPositionMinted and is used to iterate over the raw logs and unpacked data for PositionMinted events raised by the PooledLiquidityMarket contract.
type PooledLiquidityMarketPositionMintedIterator struct {
	Event *PooledLiquidityMarketPositionMinted // Event containing the contract specifics and raw log

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
func (it *PooledLiquidityMarketPositionMintedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PooledLiquidityMarketPositionMinted)
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
		it.Event = new(PooledLiquidityMarketPositionMinted)
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
func (it *PooledLiquidityMarketPositionMintedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PooledLiquidityMarketPositionMintedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PooledLiquidityMarketPositionMinted represents a PositionMinted event raised by the PooledLiquidityMarket contract.
type PooledLiquidityMarketPositionMinted struct {
	Owner     common.Address
	TickLower *big.Int
	TickUpper *big.Int
	Liquidity *big.Int
	Amount0   *big.Int
	Amount1   *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterPositionMinted is a free log retrieval operation binding the contract event 0x93203984c4f6aa9fe36590799fd75993c24a83e88f62a1c3d28b8362577b6edc.
//
// Solidity: event PositionMinted(address indexed owner, int24 indexed tickLower, int24 indexed tickUpper, uint128 liquidity, uint256 amount0, uint256 amount1)
func (_PooledLiquidityMarket *PooledLiquidityMarketFilterer) FilterPositionMinted(opts *bind.FilterOpts, owner []common.Address, tickLower []*big.Int, tickUpper []*big.Int) (*PooledLiquidityMarketPositionMintedIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var tickLowerRule []interface{}
	for _, tickLowerItem := range tickLower {
		tickLowerRule = append(tickLowerRule, tickLowerItem)
	}
	var tickUpperRule []interface{}
	for _, tickUpperItem := range tickUpper {
		tickUpperRule = append(tickUpperRule, tickUpperItem)
	}

	logs, sub, err := _PooledLiquidityMarket.contract.FilterLogs(opts, "PositionMinted", ownerRule, tickLowerRule, tickUpperRule)
	if err != nil {
		return nil, err
	}
	return &PooledLiquidityMarketPositionMintedIterator{contract: _PooledLiquidityMarket.contract, event: "PositionMinted", logs: logs, sub: sub}, nil
}

// WatchPositionMinted is a free log subscription operation binding the contract event 0x93203984c4f6aa9fe36590799fd75993c24a83e88f62a1c3d28b8362577b6edc.
//
// Solidity: event PositionMinted(address indexed owner, int24 indexed tickLower, int24 indexed tickUpper, uint128 liquidity, uint256 amount0, uint256 amount1)
func (_PooledLiquidityMarket *PooledLiquidityMarketFilterer) WatchPositionMinted(opts *bind.WatchOpts, sink chan<- *PooledLiquidityMarketPositionMinted, owner []common.Address, tickLower []*big.Int, tickUpper []*big.Int) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var tickLowerRule []interface{}
	for _, tickLowerItem := range tickLower {
		tickLowerRule = append(tickLowerRule, tickLowerItem)
	}
	var tickUpperRule []interface{}
	for _, tickUpperItem := range tickUpper {
		tickUpperRule = append(tickUpperRule, tickUpperItem)
	}

	logs, sub, err := _PooledLiquidityMarket.contract.WatchLogs(opts, "PositionMinted", ownerRule, tickLowerRule, tickUpperRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PooledLiquidityMarketPositionMinted)
				if err := _PooledLiquidityMarket.contract.UnpackLog(event, "PositionMinted", log); err != nil {
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

// ParsePositionMinted is a log parse operation binding the contract event 0x93203984c4f6aa9fe36590799fd75993c24a83e88f62a1c3d28b8362577b6edc.
//
// Solidity: event PositionMinted(address indexed owner, int24 indexed tickLower, int24 indexed tickUpper, uint128 liquidity, uint256 amount0, uint256 amount1)
func (_PooledLiquidityMarket *PooledLiquidityMarketFilterer) ParsePositionMinted(log types.Log) (*PooledLiquidityMarketPositionMinted, error) {
	event := new(PooledLiquidityMarketPositionMinted)
	if err := _PooledLiquidityMarket.contract.UnpackLog(event, "PositionMinted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PooledLiquidityMarketSwapIterator is returned from FilterSwap and is used to iterate over the raw logs and unpacked data for Swap events raised by the PooledLiquidityMarket contract.
type PooledLiquidityMarketSwapIterator struct {
	Event *PooledLiquidityMarketSwap // Event containing the contract specifics and raw log

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
func (it *PooledLiquidityMarketSwapIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PooledLiquidityMarketSwap)
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
		it.Event = new(PooledLiquidityMarketSwap)
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
func (it *PooledLiquidityMarketSwapIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PooledLiquidityMarketSwapIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PooledLiquidityMarketSwap represents a Swap event raised by the PooledLiquidityMarket contract.
type PooledLiquidityMarketSwap struct {
	Trader       common.Address
	OutcomeId    *big.Int
	AmountIn     *big.Int
	AmountOut    *big.Int
	SqrtPriceX96 *big.Int
	Tick         *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterSwap is a free log retrieval operation binding the contract event 0xba8b761f8d89e7baa33e0b9382ac49687de943616fcf6f7d954896506aad64ac.
//
// Solidity: event Swap(address indexed trader, uint256 indexed outcomeId, uint256 amountIn, uint256 amountOut, uint160 sqrtPriceX96, int24 tick)
func (_PooledLiquidityMarket *PooledLiquidityMarketFilterer) FilterSwap(opts *bind.FilterOpts, trader []common.Address, outcomeId []*big.Int) (*PooledLiquidityMarketSwapIterator, error) {

	var traderRule []interface{}
	for _, traderItem := range trader {
		traderRule = append(traderRule, traderItem)
	}
	var outcomeIdRule []interface{}
	for _, outcomeIdItem := range outcomeId {
		outcomeIdRule = append(outcomeIdRule, outcomeIdItem)
	}

	logs, sub, err := _PooledLiquidityMarket.contract.FilterLogs(opts, "Swap", traderRule, outcomeIdRule)
	if err != nil {
		return nil, err
	}
	return &PooledLiquidityMarketSwapIterator{contract: _PooledLiquidityMarket.contract, event: "Swap", logs: logs, sub: sub}, nil
}

// WatchSwap is a free log subscription operation binding the contract event 0xba8b761f8d89e7baa33e0b9382ac49687de943616fcf6f7d954896506aad64ac.
//
// Solidity: event Swap(address indexed trader, uint256 indexed outcomeId, uint256 amountIn, uint256 amountOut, uint160 sqrtPriceX96, int24 tick)
func (_PooledLiquidityMarket *PooledLiquidityMarketFilterer) WatchSwap(opts *bind.WatchOpts, sink chan<- *PooledLiquidityMarketSwap, trader []common.Address, outcomeId []*big.Int) (event.Subscription, error) {

	var traderRule []interface{}
	for _, traderItem := range trader {
		traderRule = append(traderRule, traderItem)
	}
	var outcomeIdRule []interface{}
	for _, outcomeIdItem := range outcomeId {
		outcomeIdRule = append(outcomeIdRule, outcomeIdItem)
	}

	logs, sub, err := _PooledLiquidityMarket.contract.WatchLogs(opts, "Swap", traderRule, outcomeIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PooledLiquidityMarketSwap)
				if err := _PooledLiquidityMarket.contract.UnpackLog(event, "Swap", log); err != nil {
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

// ParseSwap is a log parse operation binding the contract event 0xba8b761f8d89e7baa33e0b9382ac49687de943616fcf6f7d954896506aad64ac.
//
// Solidity: event Swap(address indexed trader, uint256 indexed outcomeId, uint256 amountIn, uint256 amountOut, uint160 sqrtPriceX96, int24 tick)
func (_PooledLiquidityMarket *PooledLiquidityMarketFilterer) ParseSwap(log types.Log) (*PooledLiquidityMarketSwap, error) {
	event := new(PooledLiquidityMarketSwap)
	if err := _PooledLiquidityMarket.contract.UnpackLog(event, "Swap", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PooledLiquidityMarketTradeIterator is returned from FilterTrade and is used to iterate over the raw logs and unpacked data for Trade events raised by the PooledLiquidityMarket contract.
type PooledLiquidityMarketTradeIterator struct {
	Event *PooledLiquidityMarketTrade // Event containing the contract specifics and raw log

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
func (it *PooledLiquidityMarketTradeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PooledLiquidityMarketTrade)
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
		it.Event = new(PooledLiquidityMarketTrade)
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
func (it *PooledLiquidityMarketTradeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PooledLiquidityMarketTradeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PooledLiquidityMarketTrade represents a Trade event raised by the PooledLiquidityMarket contract.
type PooledLiquidityMarketTrade struct {
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
func (_PooledLiquidityMarket *PooledLiquidityMarketFilterer) FilterTrade(opts *bind.FilterOpts, trader []common.Address, outcomeId []*big.Int) (*PooledLiquidityMarketTradeIterator, error) {

	var traderRule []interface{}
	for _, traderItem := range trader {
		traderRule = append(traderRule, traderItem)
	}
	var outcomeIdRule []interface{}
	for _, outcomeIdItem := range outcomeId {
		outcomeIdRule = append(outcomeIdRule, outcomeIdItem)
	}

	logs, sub, err := _PooledLiquidityMarket.contract.FilterLogs(opts, "Trade", traderRule, outcomeIdRule)
	if err != nil {
		return nil, err
	}
	return &PooledLiquidityMarketTradeIterator{contract: _PooledLiquidityMarket.contract, event: "Trade", logs: logs, sub: sub}, nil
}

// WatchTrade is a free log subscription operation binding the contract event 0xe34b2a81bbc1e1a545c34243f3cc283b6b6d1f4c2153be2a47b2612247e45865.
//
// Solidity: event Trade(address indexed trader, uint256 indexed outcomeId, uint256 amountIn, uint256 amountOut, uint256 fee, bool isBuy)
func (_PooledLiquidityMarket *PooledLiquidityMarketFilterer) WatchTrade(opts *bind.WatchOpts, sink chan<- *PooledLiquidityMarketTrade, trader []common.Address, outcomeId []*big.Int) (event.Subscription, error) {

	var traderRule []interface{}
	for _, traderItem := range trader {
		traderRule = append(traderRule, traderItem)
	}
	var outcomeIdRule []interface{}
	for _, outcomeIdItem := range outcomeId {
		outcomeIdRule = append(outcomeIdRule, outcomeIdItem)
	}

	logs, sub, err := _PooledLiquidityMarket.contract.WatchLogs(opts, "Trade", traderRule, outcomeIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PooledLiquidityMarketTrade)
				if err := _PooledLiquidityMarket.contract.UnpackLog(event, "Trade", log); err != nil {
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
func (_PooledLiquidityMarket *PooledLiquidityMarketFilterer) ParseTrade(log types.Log) (*PooledLiquidityMarketTrade, error) {
	event := new(PooledLiquidityMarketTrade)
	if err := _PooledLiquidityMarket.contract.UnpackLog(event, "Trade", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PooledLiquidityMarketTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the PooledLiquidityMarket contract.
type PooledLiquidityMarketTransferIterator struct {
	Event *PooledLiquidityMarketTransfer // Event containing the contract specifics and raw log

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
func (it *PooledLiquidityMarketTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PooledLiquidityMarketTransfer)
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
		it.Event = new(PooledLiquidityMarketTransfer)
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
func (it *PooledLiquidityMarketTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PooledLiquidityMarketTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PooledLiquidityMarketTransfer represents a Transfer event raised by the PooledLiquidityMarket contract.
type PooledLiquidityMarketTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_PooledLiquidityMarket *PooledLiquidityMarketFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*PooledLiquidityMarketTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _PooledLiquidityMarket.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &PooledLiquidityMarketTransferIterator{contract: _PooledLiquidityMarket.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_PooledLiquidityMarket *PooledLiquidityMarketFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *PooledLiquidityMarketTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _PooledLiquidityMarket.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PooledLiquidityMarketTransfer)
				if err := _PooledLiquidityMarket.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_PooledLiquidityMarket *PooledLiquidityMarketFilterer) ParseTransfer(log types.Log) (*PooledLiquidityMarketTransfer, error) {
	event := new(PooledLiquidityMarketTransfer)
	if err := _PooledLiquidityMarket.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PooledLiquidityMarketUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the PooledLiquidityMarket contract.
type PooledLiquidityMarketUnpausedIterator struct {
	Event *PooledLiquidityMarketUnpaused // Event containing the contract specifics and raw log

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
func (it *PooledLiquidityMarketUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PooledLiquidityMarketUnpaused)
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
		it.Event = new(PooledLiquidityMarketUnpaused)
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
func (it *PooledLiquidityMarketUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PooledLiquidityMarketUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PooledLiquidityMarketUnpaused represents a Unpaused event raised by the PooledLiquidityMarket contract.
type PooledLiquidityMarketUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_PooledLiquidityMarket *PooledLiquidityMarketFilterer) FilterUnpaused(opts *bind.FilterOpts) (*PooledLiquidityMarketUnpausedIterator, error) {

	logs, sub, err := _PooledLiquidityMarket.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &PooledLiquidityMarketUnpausedIterator{contract: _PooledLiquidityMarket.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_PooledLiquidityMarket *PooledLiquidityMarketFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *PooledLiquidityMarketUnpaused) (event.Subscription, error) {

	logs, sub, err := _PooledLiquidityMarket.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PooledLiquidityMarketUnpaused)
				if err := _PooledLiquidityMarket.contract.UnpackLog(event, "Unpaused", log); err != nil {
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
func (_PooledLiquidityMarket *PooledLiquidityMarketFilterer) ParseUnpaused(log types.Log) (*PooledLiquidityMarketUnpaused, error) {
	event := new(PooledLiquidityMarketUnpaused)
	if err := _PooledLiquidityMarket.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
