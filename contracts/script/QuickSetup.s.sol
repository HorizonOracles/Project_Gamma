// SPDX-License-Identifier: MIT
pragma solidity 0.8.24;

import "forge-std/Script.sol";
import "../src/HorizonToken.sol";
import "../src/OutcomeToken.sol";
import "../src/HorizonPerks.sol";
import "../src/FeeSplitter.sol";
import "../src/ResolutionModule.sol";
import "../src/AIOracleAdapter.sol";
import "../src/MarketFactory.sol";
import "../src/MarketAMM.sol";

/**
 * @title QuickSetup
 * @notice Quick deployment + test market creation for local testing
 */
contract QuickSetup is Script {
    function run() external {
        vm.startBroadcast();
        
        address deployer = msg.sender;
        address aiSigner = deployer; // Same as deployer for testing
        
        console.log("\n=== DEPLOYING CONTRACTS ===");
        console.log("Deployer:", deployer);
        console.log("AI Signer:", aiSigner);
        
        // 1. Deploy tokens
        HorizonToken horizonToken = new HorizonToken(100_000_000 * 10**18);
        OutcomeToken outcomeToken = new OutcomeToken("https://horizon.markets/api/metadata/{id}.json");
        HorizonPerks horizonPerks = new HorizonPerks(address(horizonToken));
        
        console.log("\nHorizonToken:", address(horizonToken));
        console.log("OutcomeToken:", address(outcomeToken));
        console.log("HorizonPerks:", address(horizonPerks));
        
        // 2. Deploy infrastructure
        FeeSplitter feeSplitter = new FeeSplitter(deployer);
        ResolutionModule resolutionModule = new ResolutionModule(
            address(outcomeToken),
            address(horizonToken),
            deployer // arbitrator
        );
        resolutionModule.setMinBond(1_000 * 10**18);
        resolutionModule.setDisputeWindow(172800);
        
        AIOracleAdapter aiOracleAdapter = new AIOracleAdapter(
            address(resolutionModule),
            address(horizonToken),
            aiSigner
        );
        
        console.log("\nFeeSplitter:", address(feeSplitter));
        console.log("ResolutionModule:", address(resolutionModule));
        console.log("AIOracleAdapter:", address(aiOracleAdapter));
        
        // 3. Deploy market factory
        MarketFactory marketFactory = new MarketFactory(
            address(outcomeToken),
            address(feeSplitter),
            address(horizonPerks),
            address(horizonToken)
        );
        marketFactory.setMinCreatorStake(10_000 * 10**18);
        
        console.log("\nMarketFactory:", address(marketFactory));
        
        // 4. Setup authorizations
        horizonToken.addMinter(address(marketFactory));
        outcomeToken.setResolutionAuthorization(address(resolutionModule), true);
        outcomeToken.transferOwnership(address(marketFactory));
        feeSplitter.transferOwnership(address(marketFactory));
        
        console.log("\n=== CREATING TEST MARKET ===");
        
        // 5. Create a test market with closeTime very soon (2 minutes from now)
        uint256 closeTime = block.timestamp + 120; // 2 minutes from now
        
        MarketFactory.MarketParams memory params = MarketFactory.MarketParams({
            collateralToken: address(horizonToken),
            closeTime: closeTime,
            category: "politics",
            metadataURI: "Will Donald Trump win the 2024 US Presidential election?",
            creatorStake: 10_000 * 10**18
        });
        
        horizonToken.approve(address(marketFactory), params.creatorStake);
        uint256 marketId = marketFactory.createMarket(params);
        
        MarketFactory.Market memory market = marketFactory.getMarket(marketId);
        
        console.log("\nMarket ID:", marketId);
        console.log("AMM Address:", market.amm);
        console.log("Close Time:", closeTime);
        console.log("Status: Active (ready for resolution)");
        
        // 6. Add some liquidity
        MarketAMM amm = MarketAMM(market.amm);
        uint256 liquidity = 100_000 * 10**18;
        horizonToken.approve(address(amm), liquidity);
        amm.addLiquidity(liquidity);
        
        console.log("\nLiquidity added: 100,000 HORIZON");
        
        vm.stopBroadcast();
        
        // Output for .env file
        console.log("\n=== COPY TO .ENV FILE ===");
        console.log("HORIZON_TOKEN_ADDR=", address(horizonToken));
        console.log("MARKET_FACTORY_ADDR=", address(marketFactory));
        console.log("RESOLUTION_MODULE_ADDR=", address(resolutionModule));
        console.log("AI_ORACLE_ADAPTER_ADDR=", address(aiOracleAdapter));
        console.log("\n========================");
    }
}
