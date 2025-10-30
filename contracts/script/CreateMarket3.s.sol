// SPDX-License-Identifier: MIT
pragma solidity 0.8.24;

import "forge-std/Script.sol";
import "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import "../src/MarketFactory.sol";

contract CreateMarket3 is Script {
    function run() external {
        address factoryAddr = vm.envAddress("MARKET_FACTORY_ADDR");
        address horizonTokenAddr = vm.envAddress("HORIZON_TOKEN_ADDR");
        uint256 deployerPrivateKey = vm.envUint("DEPLOYER_PRIVATE_KEY");
        
        vm.startBroadcast(deployerPrivateKey);
        
        MarketFactory factory = MarketFactory(factoryAddr);
        IERC20 horizonToken = IERC20(horizonTokenAddr);
        
        // Create a new market that closes in 600 seconds (10 minutes)
        uint256 closeTime = block.timestamp + 600;
        
        // Create MarketParams struct
        MarketFactory.MarketParams memory params = MarketFactory.MarketParams({
            collateralToken: horizonTokenAddr,
            closeTime: closeTime,
            category: "history",
            metadataURI: '{"question":"Did humans land on the Moon in 1969?","tags":["space","history"]}',
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
