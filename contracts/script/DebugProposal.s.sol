// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import "forge-std/Script.sol";
import "../src/MarketFactory.sol";
import "../src/ResolutionModule.sol";

contract DebugProposal is Script {
    function run() external view {
        // Contract addresses from deployment
        address factoryAddr = 0x2279B7A0a67DB372996a5FaB50D91eAA73d2eBe6;
        address resolutionAddr = 0xDc64a140Aa3E981100a9becA4E685f962f0cF6C9;
        
        MarketFactory factory = MarketFactory(factoryAddr);
        ResolutionModule resolutionModule = ResolutionModule(resolutionAddr);
        
        // Get market 1 details
        MarketFactory.Market memory market = factory.getMarket(1);
        
        console.log("=== Market 1 Details ===");
        console.log("ID:", market.id);
        console.log("Creator:", market.creator);
        console.log("AMM:", market.amm);
        console.log("Close Time:", market.closeTime);
        console.log("Current Time:", block.timestamp);
        console.log("Time past close:", block.timestamp > market.closeTime ? block.timestamp - market.closeTime : 0);
        console.log("Status:", uint(market.status));
        
        // Check if resolution already exists
        (
            ResolutionModule.ResolutionState state,
            uint256 proposedOutcome,
            uint256 proposalTime,
            address proposer,
            uint256 proposerBond,
            address disputer,
            uint256 disputerBond,
            string memory evidenceURI
        ) = resolutionModule.resolutions(1);
        
        console.log("\n=== Resolution State ===");
        console.log("State:", uint(state));
        console.log("Proposed Outcome:", proposedOutcome);
        console.log("Proposal Time:", proposalTime);
        console.log("Proposer:", proposer);
        console.log("Proposer Bond:", proposerBond);
        
        console.log("\n=== Computed Error Selectors ===");
        console.logBytes4(MarketFactory.InvalidCloseTime.selector);
        console.logBytes4(ResolutionModule.InvalidState.selector);
        console.logBytes4(ResolutionModule.MarketAlreadyResolved.selector);
    }
}
