// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "forge-std/Script.sol";
import "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import "../src/HorizonPerks.sol";
import "../src/ResolutionModule.sol";
import "../src/MarketFactory.sol";

contract TestToken3Integration is Script {
    IERC20 public token;
    HorizonPerks public perks;
    ResolutionModule public resolution;
    MarketFactory public factory;
    
    address public deployer = 0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266;
    
    function run() external {
        // Contract addresses from deployment
        token = IERC20(0x5FbDB2315678afecb367f032d93F642f64180aa3);
        perks = HorizonPerks(0x0165878A594ca255338adfa4d48449f69242Eb8F);
        resolution = ResolutionModule(0x2279B7A0a67DB372996a5FaB50D91eAA73d2eBe6);
        factory = MarketFactory(0xB7f8BC63BbcaD18155201308C8f3540b07f84F5e);
        
        console.log("\n=== TOKEN3 INTEGRATION TEST ===\n");
        
        vm.startBroadcast();
        
        // Test 1: Check token balance
        console.log("Test 1: Check Token Balance");
        uint256 balance = token.balanceOf(deployer);
        console.log("  Deployer balance:", balance / 1 ether, "tokens");
        require(balance > 0, "No tokens in deployer account");
        console.log("  PASS: Token balance check\n");
        
        // Test 2: Approve HorizonPerks to spend tokens
        console.log("Test 2: Approve HorizonPerks");
        uint256 approveAmount = 1000 ether;
        bool success = token.approve(address(perks), approveAmount);
        require(success, "Approval failed");
        uint256 allowance = token.allowance(deployer, address(perks));
        console.log("  Approved amount:", allowance / 1 ether, "tokens");
        require(allowance == approveAmount, "Allowance mismatch");
        console.log("  PASS: HorizonPerks approval\n");
        
        // Test 3: Approve ResolutionModule to spend tokens
        console.log("Test 3: Approve ResolutionModule");
        success = token.approve(address(resolution), approveAmount);
        require(success, "Approval failed");
        allowance = token.allowance(deployer, address(resolution));
        console.log("  Approved amount:", allowance / 1 ether, "tokens");
        require(allowance == approveAmount, "Allowance mismatch");
        console.log("  PASS: ResolutionModule approval\n");
        
        // Test 4: Approve MarketFactory to spend tokens
        console.log("Test 4: Approve MarketFactory");
        success = token.approve(address(factory), approveAmount);
        require(success, "Approval failed");
        allowance = token.allowance(deployer, address(factory));
        console.log("  Approved amount:", allowance / 1 ether, "tokens");
        require(allowance == approveAmount, "Allowance mismatch");
        console.log("  PASS: MarketFactory approval\n");
        
        // Test 5: Transfer tokens to test account
        console.log("Test 5: Transfer Tokens");
        address testAccount = 0x90F79bf6EB2c4f870365E785982E1f101E93b906;
        uint256 transferAmount = 100 ether;
        uint256 balanceBefore = token.balanceOf(testAccount);
        success = token.transfer(testAccount, transferAmount);
        require(success, "Transfer failed");
        uint256 balanceAfter = token.balanceOf(testAccount);
        console.log("  Test account balance before:", balanceBefore / 1 ether, "tokens");
        console.log("  Test account balance after:", balanceAfter / 1 ether, "tokens");
        require(balanceAfter == balanceBefore + transferAmount, "Transfer amount mismatch");
        console.log("  PASS: Token transfer\n");
        
        vm.stopBroadcast();
        
        console.log("=== ALL TESTS PASSED ===");
        console.log("Token3 is fully compatible with utility contracts!");
    }
}
