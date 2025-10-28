// SPDX-License-Identifier: MIT
pragma solidity 0.8.24;

import "forge-std/Script.sol";
import "forge-std/console.sol";

contract TestSigRecovery is Script {
    bytes32 constant PROPOSED_OUTCOME_TYPEHASH = keccak256(
        "ProposedOutcome(uint256 marketId,uint256 outcomeId,uint256 closeTime,bytes32 evidenceHash,uint256 notBefore,uint256 deadline)"
    );
    
    bytes32 constant DOMAIN_SEPARATOR = 0xd1a4b0731cca76c8670766aa2d444034f5f2e2d560875ad291352366f4198abc;
    
    function run() external view {
        // Data from the logs
        uint256 marketId = 1;
        uint256 outcomeId = 1;
        uint256 closeTime = 1761668396;
        bytes32 evidenceHash = 0x7df2850b26654bf1ecdf6b52891fac52c59691201e7532d9eef9226b77ca9db1;
        uint256 notBefore = 0x6900edad;
        uint256 deadline = 0x690109cd;
        
        // Signature from logs
        bytes memory sig = hex"db1a0e925a13879bf03fa91eac6c78a5603e2a4e37773a0a7f7eec9cfebd03f9779a19b35e5f80a076ad7ed40082d475e542961c9bc23b308da0146d5be632e61c";
        
        // Compute struct hash
        bytes32 structHash = keccak256(
            abi.encode(
                PROPOSED_OUTCOME_TYPEHASH,
                marketId,
                outcomeId,
                closeTime,
                evidenceHash,
                notBefore,
                deadline
            )
        );
        
        // Compute digest
        bytes32 digest = keccak256(abi.encodePacked("\x19\x01", DOMAIN_SEPARATOR, structHash));
        
        // Split signature
        bytes32 r;
        bytes32 s;
        uint8 v;
        assembly {
            r := mload(add(sig, 32))
            s := mload(add(sig, 64))
            v := byte(0, mload(add(sig, 96)))
        }
        
        // Recover signer
        address signer = ecrecover(digest, v, r, s);
        
        console.log("Struct Hash:", vm.toString(structHash));
        console.log("Digest:", vm.toString(digest));
        console.log("Recovered Signer:", signer);
        console.log("Expected Signer:  0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266");
    }
}
