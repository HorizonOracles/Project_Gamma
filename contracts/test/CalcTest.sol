// SPDX-License-Identifier: MIT
pragma solidity 0.8.24;

import "forge-std/Test.sol";

contract CalcTest is Test {
    uint256 constant Q96 = 2**96;
    
    function test_Amount0Calc() public view {
        uint160 sqrtPriceX96 = 56022770974670905984299832681;
        uint160 sqrtRatioBX96 = 78831026358287349202703042891;
        uint128 liquidity = 100000000000000000000000;
        
        uint256 amount0 = (uint256(liquidity) * (sqrtRatioBX96 - sqrtPriceX96)) / sqrtRatioBX96 / Q96;
        
        console.log("Amount0:", amount0);
    }
}
