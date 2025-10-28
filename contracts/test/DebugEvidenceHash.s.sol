// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import "forge-std/Script.sol";
import "forge-std/console.sol";

contract DebugEvidenceHash is Script {
    function run() external {
        string[] memory evidenceURIs = new string[](1);
        evidenceURIs[0] = "ipfs://test-evidence-uri";
        
        // Show what abi.encode produces
        bytes memory encoded = abi.encode(evidenceURIs);
        console.log("=== Evidence Hash Debug ===");
        console.log("Evidence URI:", evidenceURIs[0]);
        console.log("Encoded length:", encoded.length);
        console.logBytes(encoded);
        
        bytes32 hash = keccak256(encoded);
        console.log("Evidence hash:");
        console.logBytes32(hash);
        
        // Also test with the actual evidence we're using
        console.log("\n=== Testing with actual resolution evidence ===");
        
        // Try to retrieve the evidence from the last proposal
        // For now, just test with a known value
        string[] memory testEvidence = new string[](3);
        testEvidence[0] = "https://example.com/evidence1";
        testEvidence[1] = "https://example.com/evidence2";
        testEvidence[2] = "https://example.com/evidence3";
        
        bytes memory encoded2 = abi.encode(testEvidence);
        console.log("Test evidence encoded length:", encoded2.length);
        console.logBytes(encoded2);
        
        bytes32 hash2 = keccak256(encoded2);
        console.log("Test evidence hash:");
        console.logBytes32(hash2);
    }
}
