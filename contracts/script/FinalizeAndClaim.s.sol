// SPDX-License-Identifier: MIT
pragma solidity 0.8.24;

import "forge-std/Script.sol";
import "../src/ResolutionModule.sol";
import "../src/MarketAMM.sol";
import "../src/HorizonToken.sol";
import "../src/OutcomeToken.sol";

contract FinalizeAndClaim is Script {
    ResolutionModule resolution = ResolutionModule(0xDc64a140Aa3E981100a9becA4E685f962f0cF6C9);
    MarketAMM amm = MarketAMM(0x94099942864EA81cCF197E9D71ac53310b1468D8);
    HorizonToken horizonToken = HorizonToken(0x5FbDB2315678afecb367f032d93F642f64180aa3);
    OutcomeToken outcomeToken = OutcomeToken(0xe7f1725E7734CE288F8367e1Bb143E90bb3F0512);
    
    address constant account1 = 0x70997970C51812dc3A010C7d01b50e0d17dc79C8;
    address constant account2 = 0x3C44CdDdB6a900fa2b585dd299e03d12FA4293BC;
    
    function run() external {
        // Finalize resolution
        console.log("\n=== Finalizing Resolution ===");
        vm.broadcast(0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80);
        resolution.finalize(1);
        console.log("Market 1 resolution finalized");
        
        // Check balances before claims
        console.log("\n=== Balances Before Claims ===");
        console.log("Account1 HORIZON:", horizonToken.balanceOf(account1) / 1e18);
        console.log("Account2 HORIZON:", horizonToken.balanceOf(account2) / 1e18);
        
        // Account 1 claim (bought YES - should win)
        console.log("\n=== Account 1 Claiming (bought YES) ===");
        vm.broadcast(0x59c6995e998f97a5a0044966f0945389dc9e86dae88c7a8412f4603b6b78690d);
        uint256 payout1 = outcomeToken.redeem(1);
        console.log("Account1 payout:", payout1 / 1e18, "HORIZON");
        
        // Account 2 claim (bought NO - should not win)
        console.log("\n=== Account 2 Claiming (bought NO) ===");
        vm.broadcast(0x5de4111afa1a4b94908f83103eb1f1706367c2e68ca870fc3fb9a804cdab365a);
        uint256 payout2 = outcomeToken.redeem(1);
        console.log("Account2 payout:", payout2 / 1e18, "HORIZON");
        
        // Check balances after claims
        console.log("\n=== Balances After Claims ===");
        console.log("Account1 HORIZON:", horizonToken.balanceOf(account1) / 1e18);
        console.log("Account2 HORIZON:", horizonToken.balanceOf(account2) / 1e18);
    }
}
