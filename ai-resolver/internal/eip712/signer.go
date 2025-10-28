package eip712

import (
	"crypto/ecdsa"
	"encoding/hex"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

// Domain represents the EIP-712 domain
type Domain struct {
	Name              string
	Version           string
	ChainID           *big.Int
	VerifyingContract common.Address
}

// ProposedOutcome represents the EIP-712 typed data for AI proposals
type ProposedOutcome struct {
	MarketID     *big.Int
	OutcomeID    *big.Int
	CloseTime    *big.Int
	EvidenceHash [32]byte
	NotBefore    *big.Int
	Deadline     *big.Int
}

// Signer handles EIP-712 signing for AI proposals
type Signer struct {
	domain Domain
}

// NewSigner creates a new EIP-712 signer
func NewSigner(chainID *big.Int, verifyingContract common.Address) *Signer {
	return &Signer{
		domain: Domain{
			Name:              "AIOracleAdapter",
			Version:           "1",
			ChainID:           chainID,
			VerifyingContract: verifyingContract,
		},
	}
}

// SignProposal signs a ProposedOutcome with the given private key
func (s *Signer) SignProposal(proposal ProposedOutcome, privateKey *ecdsa.PrivateKey) ([]byte, error) {
	// Log proposal details
	fmt.Printf("\n=== SIGNING PROPOSAL ===\n")
	fmt.Printf("MarketID:     %s\n", proposal.MarketID.String())
	fmt.Printf("OutcomeID:    %s\n", proposal.OutcomeID.String())
	fmt.Printf("CloseTime:    %s (%d)\n", proposal.CloseTime.String(), proposal.CloseTime.Int64())
	fmt.Printf("EvidenceHash: %x\n", proposal.EvidenceHash)
	fmt.Printf("NotBefore:    %s (%d)\n", proposal.NotBefore.String(), proposal.NotBefore.Int64())
	fmt.Printf("Deadline:     %s (%d)\n", proposal.Deadline.String(), proposal.Deadline.Int64())

	// Log domain details
	fmt.Printf("\n=== DOMAIN ===\n")
	fmt.Printf("Name:              %s\n", s.domain.Name)
	fmt.Printf("Version:           %s\n", s.domain.Version)
	fmt.Printf("ChainID:           %s\n", s.domain.ChainID.String())
	fmt.Printf("VerifyingContract: %s\n", s.domain.VerifyingContract.Hex())

	// Compute the digest
	digest := s.computeDigest(proposal)
	fmt.Printf("\n=== DIGEST ===\n")
	fmt.Printf("Digest: %x\n", digest)

	// Sign the digest
	signature, err := crypto.Sign(digest[:], privateKey)
	if err != nil {
		return nil, fmt.Errorf("failed to sign: %w", err)
	}

	// Ethereum signatures: adjust v from [0,1] to [27,28]
	if signature[64] < 27 {
		signature[64] += 27
	}

	fmt.Printf("\n=== SIGNATURE ===\n")
	fmt.Printf("r: %x\n", signature[:32])
	fmt.Printf("s: %x\n", signature[32:64])
	fmt.Printf("v: %d\n", signature[64])
	fmt.Printf("Full signature: %x\n", signature)
	fmt.Printf("========================\n\n")

	return signature, nil
}

// VerifySignature verifies an EIP-712 signature
func (s *Signer) VerifySignature(proposal ProposedOutcome, signature []byte, expectedSigner common.Address) (bool, error) {
	if len(signature) != 65 {
		return false, fmt.Errorf("invalid signature length: %d", len(signature))
	}

	// Compute digest
	digest := s.computeDigest(proposal)

	// Adjust v for recovery
	sig := make([]byte, 65)
	copy(sig, signature)
	if sig[64] >= 27 {
		sig[64] -= 27
	}

	// Recover public key
	pubKey, err := crypto.SigToPub(digest[:], sig)
	if err != nil {
		return false, fmt.Errorf("failed to recover public key: %w", err)
	}

	// Get address from public key
	recoveredAddr := crypto.PubkeyToAddress(*pubKey)

	return recoveredAddr == expectedSigner, nil
}

// computeDigest computes the EIP-712 digest for a proposal
func (s *Signer) computeDigest(proposal ProposedOutcome) common.Hash {
	// Domain separator
	domainSeparator := s.computeDomainSeparator()

	// Typed data hash
	typedDataHash := s.hashProposal(proposal)

	// Final digest: keccak256("\x19\x01" || domainSeparator || typedDataHash)
	data := make([]byte, 0, 66)
	data = append(data, []byte("\x19\x01")...)
	data = append(data, domainSeparator[:]...)
	data = append(data, typedDataHash[:]...)

	return crypto.Keccak256Hash(data)
}

// computeDomainSeparator computes the EIP-712 domain separator
func (s *Signer) computeDomainSeparator() common.Hash {
	// EIP712Domain(string name,string version,uint256 chainId,address verifyingContract)
	domainTypeHash := crypto.Keccak256Hash([]byte(
		"EIP712Domain(string name,string version,uint256 chainId,address verifyingContract)",
	))

	nameHash := crypto.Keccak256Hash([]byte(s.domain.Name))
	versionHash := crypto.Keccak256Hash([]byte(s.domain.Version))

	// Encode domain
	data := make([]byte, 0, 160)
	data = append(data, domainTypeHash[:]...)
	data = append(data, nameHash[:]...)
	data = append(data, versionHash[:]...)
	data = append(data, common.LeftPadBytes(s.domain.ChainID.Bytes(), 32)...)
	data = append(data, common.LeftPadBytes(s.domain.VerifyingContract.Bytes(), 32)...)

	return crypto.Keccak256Hash(data)
}

// hashProposal computes the typed data hash for a proposal
func (s *Signer) hashProposal(proposal ProposedOutcome) common.Hash {
	// ProposedOutcome(uint256 marketId,uint256 outcomeId,uint256 closeTime,bytes32 evidenceHash,uint256 notBefore,uint256 deadline)
	proposalTypeHash := crypto.Keccak256Hash([]byte(
		"ProposedOutcome(uint256 marketId,uint256 outcomeId,uint256 closeTime,bytes32 evidenceHash,uint256 notBefore,uint256 deadline)",
	))

	// Encode proposal (evidenceHash is already bytes32, no conversion needed)
	data := make([]byte, 0, 224)
	data = append(data, proposalTypeHash[:]...)
	data = append(data, common.LeftPadBytes(proposal.MarketID.Bytes(), 32)...)
	data = append(data, common.LeftPadBytes(proposal.OutcomeID.Bytes(), 32)...)
	data = append(data, common.LeftPadBytes(proposal.CloseTime.Bytes(), 32)...)
	data = append(data, proposal.EvidenceHash[:]...)
	data = append(data, common.LeftPadBytes(proposal.NotBefore.Bytes(), 32)...)
	data = append(data, common.LeftPadBytes(proposal.Deadline.Bytes(), 32)...)

	return crypto.Keccak256Hash(data)
}

// ParsePrivateKey parses a hex-encoded private key
func ParsePrivateKey(hexKey string) (*ecdsa.PrivateKey, error) {
	// Remove 0x prefix if present
	if len(hexKey) >= 2 && hexKey[:2] == "0x" {
		hexKey = hexKey[2:]
	}

	keyBytes, err := hex.DecodeString(hexKey)
	if err != nil {
		return nil, fmt.Errorf("failed to decode hex: %w", err)
	}

	privateKey, err := crypto.ToECDSA(keyBytes)
	if err != nil {
		return nil, fmt.Errorf("failed to parse private key: %w", err)
	}

	return privateKey, nil
}

// GetAddress returns the Ethereum address for a private key
func GetAddress(privateKey *ecdsa.PrivateKey) common.Address {
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return common.Address{}
	}
	return crypto.PubkeyToAddress(*publicKeyECDSA)
}

// ComputeEvidenceHash computes the evidence hash from URIs
func ComputeEvidenceHash(evidenceURIs []string) [32]byte {
	// Must match Solidity's keccak256(abi.encode(string[]))
	// abi.encode for dynamic arrays starts with offset to array data

	var encoded []byte

	// Initial offset to array data (always 0x20 = 32 bytes for a single dynamic array)
	encoded = append(encoded, common.LeftPadBytes(big.NewInt(32).Bytes(), 32)...)

	// Array length
	arrayLen := big.NewInt(int64(len(evidenceURIs)))
	encoded = append(encoded, common.LeftPadBytes(arrayLen.Bytes(), 32)...)

	// Calculate offsets for each string (relative to start of string offset list)
	currentOffset := len(evidenceURIs) * 32 // Start after all string offsets
	offsets := make([]int, len(evidenceURIs))

	for i, uri := range evidenceURIs {
		offsets[i] = currentOffset
		dataLen := len(uri)
		paddedLen := ((dataLen + 31) / 32) * 32
		currentOffset += 32 + paddedLen // length field + padded data
	}

	// Write offsets
	for _, offset := range offsets {
		offsetBig := big.NewInt(int64(offset))
		encoded = append(encoded, common.LeftPadBytes(offsetBig.Bytes(), 32)...)
	}

	// Write each string (length + data)
	for _, uri := range evidenceURIs {
		uriBytes := []byte(uri)
		uriLen := big.NewInt(int64(len(uriBytes)))

		// String length
		encoded = append(encoded, common.LeftPadBytes(uriLen.Bytes(), 32)...)

		// String data (right-padded to 32-byte boundary)
		encoded = append(encoded, uriBytes...)
		padding := ((len(uriBytes)+31)/32)*32 - len(uriBytes)
		if padding > 0 {
			encoded = append(encoded, make([]byte, padding)...)
		}
	}

	// Log the encoded bytes for debugging
	fmt.Printf("DEBUG: Evidence URIs: %v\n", evidenceURIs)
	fmt.Printf("DEBUG: Encoded bytes (hex): %s\n", hex.EncodeToString(encoded))
	fmt.Printf("DEBUG: Encoded bytes length: %d\n", len(encoded))

	hash := crypto.Keccak256Hash(encoded)
	fmt.Printf("DEBUG: Evidence hash: %s\n", hash.Hex())

	var result [32]byte
	copy(result[:], hash[:])
	return result
}
