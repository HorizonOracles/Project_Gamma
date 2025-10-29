// SPDX-License-Identifier: MIT
pragma solidity 0.8.24;

import "forge-std/Test.sol";

contract TickMathTest is Test {
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
        uint256 absTick = tick < 0 ? uint256(-int256(tick)) : uint256(int256(tick));
        
        int256 priceDeviation = (int256(tick) * 1e18) / 140000;
        uint256 price = uint256(int256(5e17) + priceDeviation);
        
        if (price < 1e16) price = 1e16;
        if (price > 99e16) price = 99e16;
        
        uint256 sqrtPrice = _sqrt(price);
        return uint160((sqrtPrice * Q96) / 1e9);
    }
    
    function test_TickMath() public view {
        uint160 currentPrice = uint160(Q96); // 2^96
        int24 tickLower = -69000;
        int24 tickUpper = 69000;
        
        uint160 sqrtRatioA = _getSqrtRatioAtTick(tickLower);
        uint160 sqrtRatioB = _getSqrtRatioAtTick(tickUpper);
        
        console.log("Current sqrtPrice:", currentPrice);
        console.log("SqrtRatioA (tick -69000):", sqrtRatioA);
        console.log("SqrtRatioB (tick 69000):", sqrtRatioB);
        console.log("currentPrice <= sqrtRatioA?", currentPrice <= sqrtRatioA);
        console.log("currentPrice < sqrtRatioB?", currentPrice < sqrtRatioB);
    }
}
