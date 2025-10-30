// SPDX-License-Identifier: MIT
pragma solidity 0.8.24;

import "forge-std/Script.sol";
import "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import "../src/MarketAMM.sol";

/**
 * @title AddLiquidityMarket3
 * @notice Script to add liquidity to Market ID 3 AMM
 */
contract AddLiquidityMarket3 is Script {
    function run() external {
        uint256 deployerPrivateKey = vm.envUint("DEPLOYER_PRIVATE_KEY");
        
        // Contract addresses (from deployment)
        address ammAddress = 0x94099942864EA81cCF197E9D71ac53310b1468D8;
        address horizonTokenAddress = 0x5FbDB2315678afecb367f032d93F642f64180aa3;
        
        MarketAMM amm = MarketAMM(ammAddress);
        IERC20 horizonToken = IERC20(horizonTokenAddress);
        
        // Amount to add: 50,000 HORIZON tokens
        uint256 liquidityAmount = 50_000 * 10**18;
        
        vm.startBroadcast(deployerPrivateKey);
        
        // 1. Approve AMM to spend HORIZON tokens
        console.log("Approving AMM to spend HORIZON tokens...");
        horizonToken.approve(ammAddress, liquidityAmount);
        console.log("Approved:", liquidityAmount / 10**18, "HORIZON");
        
        // 2. Add liquidity
        console.log("\nAdding liquidity to Market ID 3...");
        uint256 lpTokens = amm.addLiquidity(liquidityAmount);
        console.log("Liquidity added:", liquidityAmount / 10**18, "HORIZON");
        console.log("LP tokens received:", lpTokens / 10**18);
        
        // 3. Check reserves
        console.log("\n=== Market AMM Status ===");
        console.log("Reserve YES:", amm.reserveYes() / 10**18);
        console.log("Reserve NO:", amm.reserveNo() / 10**18);
        console.log("Total Collateral:", amm.totalCollateral() / 10**18);
        console.log("LP Token Balance:", amm.balanceOf(vm.addr(deployerPrivateKey)) / 10**18);
        
        vm.stopBroadcast();
        
        console.log("\n[SUCCESS] Liquidity successfully added to Market ID 3!");
        console.log("AMM is now ready for trading");
    }
}
