// SPDX-License-Identifier: MIT
pragma solidity 0.8.24;

import "forge-std/Script.sol";
import "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import "../src/MarketFactory.sol";
import "../src/MarketAMM.sol";
import "../src/OutcomeToken.sol";

/**
 * @title CreateAIMarket
 * @notice Create a market that can be resolved by AI immediately
 */
contract CreateAIMarket is Script {
    function run() external {
        // Load contract addresses
        address horizonTokenAddr = vm.envAddress("HORIZON_TOKEN_ADDRESS");
        address marketFactoryAddr = vm.envAddress("MARKET_FACTORY_ADDRESS");
        
        IERC20 horizonToken = IERC20(horizonTokenAddr);
        MarketFactory marketFactory = MarketFactory(marketFactoryAddr);
        
        console.log("\n=== CREATING AI-RESOLVABLE MARKET ===\n");
        
        vm.startBroadcast();
        
        // Set close time to 2 minutes in the future (we'll advance time after creation)
        uint256 closeTime = block.timestamp + 120;
        
        // Market parameters - a simple factual question that AI can resolve
        MarketFactory.MarketParams memory params = MarketFactory.MarketParams({
            collateralToken: address(horizonToken),
            closeTime: closeTime,
            category: "crypto",
            metadataURI: "Did Bitcoin reach $100,000 in 2024?",
            creatorStake: 10_000 * 10**18  // 10k HORIZON
        });
        
        // Approve HORIZON for staking
        horizonToken.approve(address(marketFactory), params.creatorStake);
        
        // Create market
        uint256 marketId = marketFactory.createMarket(params);
        
        // Get market info
        MarketFactory.Market memory market = marketFactory.getMarket(marketId);
        
        vm.stopBroadcast();
        
        console.log("Market Created Successfully!");
        console.log("==========================================");
        console.log("Market ID:", marketId);
        console.log("AMM Address:", market.amm);
        console.log("Close Time:", closeTime, "(2 minutes from now)");
        console.log("Question: Did Bitcoin reach $100,000 in 2024?");
        console.log("Creator Stake: 10,000 HORIZON");
        console.log("==========================================\n");
        
        console.log("To resolve this market with AI, run:");
        console.log('curl -X POST http://localhost:8080/v1/propose \\');
        console.log('  -H "Content-Type: application/json" \\');
        console.log('  -d \'{"marketId":', marketId, ', "question": "Did Bitcoin reach $100,000 in 2024?"}\'');
        console.log("\n");
    }
}
