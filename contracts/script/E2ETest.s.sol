// SPDX-License-Identifier: MIT
pragma solidity 0.8.24;

import "forge-std/Script.sol";
import "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import "../src/MarketFactory.sol";
import "../src/MarketAMM.sol";
import "../src/OutcomeToken.sol";

/**
 * @title E2ETest
 * @notice End-to-end test script for creating market, adding liquidity, and trading
 */
contract E2ETest is Script {
    // Contract addresses (will be loaded from environment)
    address public horizonTokenAddr;
    address public marketFactoryAddr;
    address public outcomeTokenAddr;
    
    // Contracts
    IERC20 public horizonToken;
    MarketFactory public marketFactory;
    OutcomeToken public outcomeToken;
    
    // Test accounts
    address public deployer = 0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266;
    address public trader1 = 0x70997970C51812dc3A010C7d01b50e0d17dc79C8;
    address public trader2 = 0x3C44CdDdB6a900fa2b585dd299e03d12FA4293BC;
    
    // Market parameters
    uint256 public marketId;
    address public ammAddress;
    
    function run() external {
        // Load contract addresses from environment
        horizonTokenAddr = vm.envAddress("HORIZON_TOKEN_ADDR");
        marketFactoryAddr = vm.envAddress("MARKET_FACTORY_ADDR");
        outcomeTokenAddr = vm.envAddress("OUTCOME_TOKEN_ADDR");
        
        horizonToken = IERC20(horizonTokenAddr);
        marketFactory = MarketFactory(marketFactoryAddr);
        outcomeToken = OutcomeToken(outcomeTokenAddr);
        
        console.log("\n=== E2E TEST: MARKET CREATION & TRADING ===\n");
        
        // Step 1: Create market
        console.log("STEP 1: Creating market...");
        createMarket();
        
        // Step 2: Add liquidity
        console.log("\nSTEP 2: Adding liquidity...");
        addLiquidity();
        
        // Step 3: Execute trades
        console.log("\nSTEP 3: Executing trades...");
        executeTrades();
        
        console.log("\n=== E2E TEST COMPLETE ===");
        outputSummary();
    }
    
    function createMarket() internal {
        vm.startBroadcast();
        
        // Calculate close time: 1 hour from now
        uint256 closeTime = block.timestamp + 3600;
        
        // Market parameters
        MarketFactory.MarketParams memory params = MarketFactory.MarketParams({
            collateralToken: address(horizonToken),
            closeTime: closeTime,
            category: "test",
            metadataURI: "Will Ethereum reach $5000 by end of 2024?",
            creatorStake: 10_000 * 10**18  // 10k HORIZON
        });
        
        // Approve HORIZON for staking
        horizonToken.approve(address(marketFactory), params.creatorStake);
        
        // Create market
        marketId = marketFactory.createMarket(params);
        
        // Get AMM address from market info
        MarketFactory.Market memory market = marketFactory.getMarket(marketId);
        ammAddress = market.amm;
        
        vm.stopBroadcast();
        
        console.log("  Market ID:", marketId);
        console.log("  AMM Address:", ammAddress);
        console.log("  Close Time:", closeTime);
        console.log("  Creator Stake: 10,000 HORIZON");
    }
    
    function addLiquidity() internal {
        vm.startBroadcast();
        
        MarketAMM amm = MarketAMM(ammAddress);
        uint256 liquidity = 100_000 * 10**18;  // 100k HORIZON
        
        // Approve HORIZON for liquidity
        horizonToken.approve(ammAddress, liquidity);
        
        // Add liquidity
        amm.addLiquidity(liquidity);
        
        vm.stopBroadcast();
        
        console.log("  Liquidity Added: 100,000 HORIZON");
        console.log("  LP Tokens Received:", amm.balanceOf(deployer) / 10**18);
    }
    
    function executeTrades() internal {
        MarketAMM amm = MarketAMM(ammAddress);
        
        // Trade 1: Trader1 buys YES
        console.log("\n  Trade 1: Trader1 buys YES tokens");
        vm.startBroadcast(0x59c6995e998f97a5a0044966f0945389dc9e86dae88c7a8412f4603b6b78690d); // trader1 private key
        
        // Get HORIZON tokens for trader1
        vm.stopBroadcast();
        vm.startBroadcast();
        horizonToken.transfer(trader1, 10_000 * 10**18);
        vm.stopBroadcast();
        
        vm.startBroadcast(0x59c6995e998f97a5a0044966f0945389dc9e86dae88c7a8412f4603b6b78690d);
        IERC20(horizonTokenAddr).approve(ammAddress, 5_000 * 10**18);
        amm.buyYes(5_000 * 10**18, 1);  // Buy YES with 5k HORIZON
        vm.stopBroadcast();
        console.log("    Trader1 spent: 5,000 HORIZON for YES tokens");
        
        // Trade 2: Trader2 buys NO
        console.log("\n  Trade 2: Trader2 buys NO tokens");
        vm.startBroadcast();
        horizonToken.transfer(trader2, 10_000 * 10**18);
        vm.stopBroadcast();
        
        vm.startBroadcast(0x5de4111afa1a4b94908f83103eb1f1706367c2e68ca870fc3fb9a804cdab365a); // trader2 private key
        IERC20(horizonTokenAddr).approve(ammAddress, 5_000 * 10**18);
        amm.buyNo(5_000 * 10**18, 1);  // Buy NO with 5k HORIZON
        vm.stopBroadcast();
        console.log("    Trader2 spent: 5,000 HORIZON for NO tokens");
    }
    
    function outputSummary() internal view {
        MarketAMM amm = MarketAMM(ammAddress);
        
        console.log("\n==========================================");
        console.log("          TEST SUMMARY");
        console.log("==========================================\n");
        
        console.log("Market ID:", marketId);
        console.log("AMM Address:", ammAddress);
        
        // Get YES and NO token IDs
        uint256 yesTokenId = outcomeToken.encodeTokenId(marketId, 1);  // YES = outcome 1
        uint256 noTokenId = outcomeToken.encodeTokenId(marketId, 0);   // NO = outcome 0
        
        console.log("\nToken Balances:");
        console.log("  Trader1 YES:", outcomeToken.balanceOf(trader1, yesTokenId) / 10**18);
        console.log("  Trader2 NO:", outcomeToken.balanceOf(trader2, noTokenId) / 10**18);
        
        console.log("\nAMM Reserves:");
        console.log("  YES Reserve:", amm.reserveYes() / 10**18);
        console.log("  NO Reserve:", amm.reserveNo() / 10**18);
        
        console.log("\n==========================================\n");
    }
}
