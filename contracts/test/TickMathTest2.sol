// SPDX-License-Identifier: MIT
pragma solidity 0.8.24;

import "forge-std/Test.sol";

contract TickMathTest2 is Test {
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
    
    function test_Tick0() public view {
        uint160 sqrtRatio0 = _getSqrtRatioAtTick(0);
        console.log("SqrtRatio at tick 0:", sqrtRatio0);
        console.log("Expected (2^96):", Q96);
        
        // Calculate actual price from sqrtRatio
        uint256 price = (uint256(sqrtRatio0) * uint256(sqrtRatio0) * 1e18) / (Q96 * Q96);
        console.log("Price at tick 0:", price);
    }
}
