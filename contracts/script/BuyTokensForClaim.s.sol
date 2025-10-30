// SPDX-License-Identifier: MIT
pragma solidity 0.8.24;

import "forge-std/Script.sol";
import "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import "../src/MarketAMM.sol";

/**
 * @notice Script to setup test accounts with tokens and buy outcomes for Market 3
 * Account 1 (0x70997970C51812dc3A010C7d01b50e0d17dc79C8) will buy YES
 * Account 2 (0x3C44CdDdB6a900fa2b585dd299e03d12FA4293BC) will buy NO
 */
contract BuyTokensForClaim is Script {
    IERC20 horizonToken = IERC20(0x5FbDB2315678afecb367f032d93F642f64180aa3);
    MarketAMM amm = MarketAMM(0x94099942864EA81cCF197E9D71ac53310b1468D8);
    
    address constant account1 = 0x70997970C51812dc3A010C7d01b50e0d17dc79C8;
    address constant account2 = 0x3C44CdDdB6a900fa2b585dd299e03d12FA4293BC;
    
    function run() external {
        transferTokens();
        buyYesTokens();
        buyNoTokens();
        
        console.log("\n=== Setup Complete ===");
        console.log("Market ID: 3");
    }
    
    function transferTokens() internal {
        console.log("\n=== Transferring HORIZON tokens ===");
        vm.broadcast(0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80);
        horizonToken.transfer(account1, 100000 ether);
        
        vm.broadcast(0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80);
        horizonToken.transfer(account2, 100000 ether);
        
        console.log("Transferred HORIZON to accounts");
    }
    
    function buyYesTokens() internal {
        console.log("\n=== Account 1 buying YES tokens ===");
        
        vm.broadcast(0x59c6995e998f97a5a0044966f0945389dc9e86dae88c7a8412f4603b6b78690d);
        horizonToken.approve(address(amm), 10000 ether);
        
        vm.broadcast(0x59c6995e998f97a5a0044966f0945389dc9e86dae88c7a8412f4603b6b78690d);
        uint256 tokens = amm.buyYes(10000 ether, 0);
        
        console.log("Account1 bought YES tokens:", tokens / 1e18);
    }
    
    function buyNoTokens() internal {
        console.log("\n=== Account 2 buying NO tokens ===");
        
        vm.broadcast(0x5de4111afa1a4b94908f83103eb1f1706367c2e68ca870fc3fb9a804cdab365a);
        horizonToken.approve(address(amm), 10000 ether);
        
        vm.broadcast(0x5de4111afa1a4b94908f83103eb1f1706367c2e68ca870fc3fb9a804cdab365a);
        uint256 tokens = amm.buyNo(10000 ether, 0);
        
        console.log("Account2 bought NO tokens:", tokens / 1e18);
    }
}
