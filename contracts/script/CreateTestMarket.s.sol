// SPDX-License-Identifier: MIT
pragma solidity 0.8.24;

import "forge-std/Script.sol";
import "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import "../src/MarketFactory.sol";

/**
 * @title CreateTestMarket
 * @notice Creates a test market on existing deployed contracts
 */
contract CreateTestMarket is Script {
    function run() external {
        // Use existing deployed addresses
        address horizonTokenAddr = 0x68B1D87F95878fE05B998F19b66F4baba5De1aed;
        address marketFactoryAddr = 0x7a2088a1bFc9d81c55368AE168C2C02570cB814F;
        
        IERC20 horizonToken = IERC20(horizonTokenAddr);
        MarketFactory marketFactory = MarketFactory(marketFactoryAddr);
        
        vm.startBroadcast();
        
        console.log("Creating market...");
        console.log("Deployer:", msg.sender);
        console.log("MarketFactory:", address(marketFactory));
        
        // Create a test market with closeTime 2 minutes in the future
        uint256 closeTime = block.timestamp + 120;
        
        MarketFactory.MarketParams memory params = MarketFactory.MarketParams({
            collateralToken: address(horizonToken),
            closeTime: closeTime,
            category: "politics",
            metadataURI: "Will Donald Trump win the 2024 US Presidential election?",
            creatorStake: 10_000 * 10**18
        });
        
        console.log("Approving HORIZON tokens...");
        horizonToken.approve(address(marketFactory), params.creatorStake);
        
        console.log("Creating market with closeTime:", closeTime);
        uint256 marketId = marketFactory.createMarket(params);
        
        console.log("\n=== SUCCESS ===");
        console.log("Market ID:", marketId);
        console.log("Close Time:", closeTime);
        
        vm.stopBroadcast();
    }
}
