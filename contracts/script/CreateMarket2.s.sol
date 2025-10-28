// SPDX-License-Identifier: MIT
pragma solidity 0.8.24;

import "forge-std/Script.sol";
import "../src/MarketFactory.sol";
import "../src/HorizonToken.sol";

contract CreateMarket2 is Script {
    function run() external {
        address factoryAddr = vm.envAddress("MARKET_FACTORY_ADDR");
        address horizonTokenAddr = vm.envAddress("HORIZON_TOKEN_ADDR");
        uint256 deployerPrivateKey = vm.envUint("DEPLOYER_PRIVATE_KEY");
        
        vm.startBroadcast(deployerPrivateKey);
        
        MarketFactory factory = MarketFactory(factoryAddr);
        HorizonToken horizonToken = HorizonToken(horizonTokenAddr);
        
        // Create a new market that closes in 300 seconds (5 minutes)
        uint256 closeTime = block.timestamp + 300;
        
        // Create MarketParams struct
        MarketFactory.MarketParams memory params = MarketFactory.MarketParams({
            collateralToken: horizonTokenAddr,
            closeTime: closeTime,
            category: "sports",
            metadataURI: '{"question":"Will the Kansas City Chiefs win Super Bowl LIX?","tags":["NFL","superbowl"]}',
            creatorStake: 10000 ether
        });
        
        // Approve HORIZON token transfer
        horizonToken.approve(address(factory), params.creatorStake);
        
        uint256 marketId = factory.createMarket(params);
        
        console.log("Created Market ID:", marketId);
        console.log("Close Time:", closeTime);
        
        vm.stopBroadcast();
    }
}
