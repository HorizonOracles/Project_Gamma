// SPDX-License-Identifier: MIT
pragma solidity 0.8.24;

import "forge-std/Test.sol";

contract AmountTest is Test {
    uint256 constant Q96 = 2**96;
    
    function _sqrt(uint256 x) internal pure returns (uint256) {
        if (x == 0) return 0;
        uint256 z = (x + 1) / 2;
        uint256 y = x;
        while (z < y) {
            y = z;
            z = (x / z + z) / 2;
        }
        return y;
    }
    
    function _getSqrtRatioAtTick(int24 tick) internal pure returns (uint160) {
        int256 priceDeviation = (int256(tick) * 1e18) / 140000;
        uint256 price = uint256(int256(5e17) + priceDeviation);
        
        if (price < 1e16) price = 1e16;
        if (price > 99e16) price = 99e16;
        
        uint256 sqrtPrice = _sqrt(price);
        return uint160((sqrtPrice * Q96) / 1e9);
    }
    
    function test_AmountCalculation() public view {
        uint160 sqrtPriceX96 = 56022770974670905984299832681; // tick 0
        int24 tickLower = -69000;
        int24 tickUpper = 69000;
        
        uint160 sqrtRatioAX96 = _getSqrtRatioAtTick(tickLower);
        uint160 sqrtRatioBX96 = _getSqrtRatioAtTick(tickUpper);
        
        console.log("Current sqrtPrice:", sqrtPriceX96);
        console.log("SqrtRatioA:", sqrtRatioAX96);
        console.log("SqrtRatioB:", sqrtRatioBX96);
        console.log("");
        console.log("sqrtPrice <= sqrtRatioA?", sqrtPriceX96 <= sqrtRatioAX96);
        console.log("sqrtPrice < sqrtRatioB?", sqrtPriceX96 < sqrtRatioBX96);
        
        // Determine which branch
        if (sqrtPriceX96 <= sqrtRatioAX96) {
            console.log("Branch: Below range (amount0 only)");
        } else if (sqrtPriceX96 < sqrtRatioBX96) {
            console.log("Branch: In range (both amounts)");
        } else {
            console.log("Branch: Above range (amount1 only)");
        }
    }
}
