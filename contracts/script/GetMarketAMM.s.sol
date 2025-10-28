// SPDX-License-Identifier: MIT
pragma solidity 0.8.24;

import "forge-std/Script.sol";
import "../src/MarketFactory.sol";

contract GetMarketAMM is Script {
    function run() external view {
        address factoryAddr = 0x2279B7A0a67DB372996a5FaB50D91eAA73d2eBe6;
        MarketFactory factory = MarketFactory(factoryAddr);
        
        (
            uint256 id,
            address creator,
            address amm,
            address collateralToken,
            uint256 closeTime,
            string memory category,
            string memory metadataURI,
            uint256 creatorStake,
            ,
            MarketFactory.MarketStatus status
        ) = factory.markets(1);
        
        console.log("Market ID:", id);
        console.log("AMM Address:", amm);
        console.log("Close Time:", closeTime);
        console.log("Status:", uint8(status));
    }
}
