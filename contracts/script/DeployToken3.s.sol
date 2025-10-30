// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "forge-std/Script.sol";
import "../src/token/Token3.sol";

contract DeployToken3 is Script {
    function run() external returns (address) {
        vm.startBroadcast();
        
        Token token = new Token();
        console.log("Token3 deployed at:", address(token));
        
        // Initialize the token
        token.init("Horizon Token", "HORIZON", 1_000_000_000 ether);
        console.log("Token initialized with 1B tokens");
        
        // Set mode to normal to enable transfers
        token.setMode(token.MODE_NORMAL());
        console.log("Transfer mode set to NORMAL");
        
        vm.stopBroadcast();
        
        return address(token);
    }
}
