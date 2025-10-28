package adapter

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/project-gamma/ai-resolver/pkg/abi"
)

// Client wraps Ethereum client and contract bindings
type Client struct {
	eth        *ethclient.Client
	chainID    *big.Int
	signer     *ecdsa.PrivateKey
	signerAddr common.Address

	// Contract instances
	adapter       *abi.AIOracleAdapter
	factory       *abi.MarketFactory
	resolutionMod *abi.ResolutionModule
	horizonToken  *abi.HorizonToken

	// Contract addresses
	adapterAddr      common.Address
	factoryAddr      common.Address
	resolutionAddr   common.Address
	horizonTokenAddr common.Address
}

// Config holds client configuration
type Config struct {
	RPCURL              string
	ChainID             int64
	SignerPrivateKey    string
	AdapterAddress      string
	FactoryAddress      string
	ResolutionAddress   string
	HorizonTokenAddress string
}

// NewClient creates a new contract client
func NewClient(ctx context.Context, cfg Config) (*Client, error) {
	// Connect to Ethereum node
	eth, err := ethclient.DialContext(ctx, cfg.RPCURL)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to Ethereum node: %w", err)
	}

	// Parse private key
	privateKey, err := crypto.HexToECDSA(cfg.SignerPrivateKey)
	if err != nil {
		return nil, fmt.Errorf("failed to parse private key: %w", err)
	}

	signerAddr := crypto.PubkeyToAddress(privateKey.PublicKey)
	chainID := big.NewInt(cfg.ChainID)

	// Parse contract addresses
	adapterAddr := common.HexToAddress(cfg.AdapterAddress)
	factoryAddr := common.HexToAddress(cfg.FactoryAddress)
	resolutionAddr := common.HexToAddress(cfg.ResolutionAddress)
	horizonTokenAddr := common.HexToAddress(cfg.HorizonTokenAddress)

	// Create contract instances
	adapter, err := abi.NewAIOracleAdapter(adapterAddr, eth)
	if err != nil {
		return nil, fmt.Errorf("failed to create adapter instance: %w", err)
	}

	factory, err := abi.NewMarketFactory(factoryAddr, eth)
	if err != nil {
		return nil, fmt.Errorf("failed to create factory instance: %w", err)
	}

	resolutionMod, err := abi.NewResolutionModule(resolutionAddr, eth)
	if err != nil {
		return nil, fmt.Errorf("failed to create resolution module instance: %w", err)
	}

	horizonToken, err := abi.NewHorizonToken(horizonTokenAddr, eth)
	if err != nil {
		return nil, fmt.Errorf("failed to create horizon token instance: %w", err)
	}

	return &Client{
		eth:              eth,
		chainID:          chainID,
		signer:           privateKey,
		signerAddr:       signerAddr,
		adapter:          adapter,
		factory:          factory,
		resolutionMod:    resolutionMod,
		horizonToken:     horizonToken,
		adapterAddr:      adapterAddr,
		factoryAddr:      factoryAddr,
		resolutionAddr:   resolutionAddr,
		horizonTokenAddr: horizonTokenAddr,
	}, nil
}

// Close closes the Ethereum client connection
func (c *Client) Close() {
	c.eth.Close()
}

// MarketInfo holds market details
type MarketInfo struct {
	ID              *big.Int
	Creator         common.Address
	AMM             common.Address
	CollateralToken common.Address
	CloseTime       *big.Int
	Category        string
	MetadataURI     string
	CreatorStake    *big.Int
	StakeRefunded   bool
	Status          uint8
}

// GetMarket fetches market details from the factory
func (c *Client) GetMarket(ctx context.Context, marketID *big.Int) (*MarketInfo, error) {
	market, err := c.factory.GetMarket(&bind.CallOpts{Context: ctx}, marketID)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch market: %w", err)
	}

	return &MarketInfo{
		ID:              market.Id,
		Creator:         market.Creator,
		AMM:             market.Amm,
		CollateralToken: market.CollateralToken,
		CloseTime:       market.CloseTime,
		Category:        market.Category,
		MetadataURI:     market.MetadataURI,
		CreatorStake:    market.CreatorStake,
		StakeRefunded:   market.StakeRefunded,
		Status:          market.Status,
	}, nil
}

// CheckAllowance checks if the adapter has sufficient token allowance
func (c *Client) CheckAllowance(ctx context.Context) (*big.Int, error) {
	allowance, err := c.horizonToken.Allowance(&bind.CallOpts{Context: ctx}, c.signerAddr, c.adapterAddr)
	if err != nil {
		return nil, fmt.Errorf("failed to check allowance: %w", err)
	}
	return allowance, nil
}

// ApproveBond approves the adapter to spend HORIZON tokens for bonding
func (c *Client) ApproveBond(ctx context.Context, amount *big.Int) (*types.Transaction, error) {
	auth, err := c.newTransactor(ctx)
	if err != nil {
		return nil, err
	}

	tx, err := c.horizonToken.Approve(auth, c.adapterAddr, amount)
	if err != nil {
		return nil, fmt.Errorf("failed to approve bond: %w", err)
	}

	return tx, nil
}

// ProposeOutcome submits an AI-generated proposal to the adapter
// The proposal parameter must be the exact same struct that was signed
func (c *Client) ProposeOutcome(ctx context.Context, proposal abi.AIOracleAdapterProposedOutcome, signature []byte, bondAmount *big.Int, evidenceURIs []string) (*types.Transaction, error) {
	auth, err := c.newTransactor(ctx)
	if err != nil {
		return nil, err
	}

	tx, err := c.adapter.ProposeAI(auth, proposal, signature, bondAmount, evidenceURIs)
	if err != nil {
		return nil, fmt.Errorf("failed to propose outcome: %w", err)
	}

	return tx, nil
}

// WaitForTransaction waits for a transaction to be mined
func (c *Client) WaitForTransaction(ctx context.Context, tx *types.Transaction) (*types.Receipt, error) {
	timeout := 2 * time.Minute
	deadline := time.Now().Add(timeout)

	for time.Now().Before(deadline) {
		receipt, err := c.eth.TransactionReceipt(ctx, tx.Hash())
		if err == nil {
			if receipt.Status == types.ReceiptStatusSuccessful {
				return receipt, nil
			}
			return nil, fmt.Errorf("transaction failed: %s", tx.Hash().Hex())
		}

		// Wait before retrying
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		case <-time.After(2 * time.Second):
		}
	}

	return nil, fmt.Errorf("transaction timeout: %s", tx.Hash().Hex())
}

// GetBalance gets the signer's HORIZON token balance
func (c *Client) GetBalance(ctx context.Context) (*big.Int, error) {
	balance, err := c.horizonToken.BalanceOf(&bind.CallOpts{Context: ctx}, c.signerAddr)
	if err != nil {
		return nil, fmt.Errorf("failed to get balance: %w", err)
	}
	return balance, nil
}

// GetSignerAddress returns the signer's Ethereum address
func (c *Client) GetSignerAddress() common.Address {
	return c.signerAddr
}

// GetChainID returns the configured chain ID
func (c *Client) GetChainID() *big.Int {
	return c.chainID
}

// GetCurrentBlockTimestamp fetches the current blockchain timestamp
func (c *Client) GetCurrentBlockTimestamp(ctx context.Context) (int64, error) {
	header, err := c.eth.HeaderByNumber(ctx, nil)
	if err != nil {
		return 0, fmt.Errorf("failed to get latest block header: %w", err)
	}
	return int64(header.Time), nil
}

// newTransactor creates a new transaction signer
func (c *Client) newTransactor(ctx context.Context) (*bind.TransactOpts, error) {
	nonce, err := c.eth.PendingNonceAt(ctx, c.signerAddr)
	if err != nil {
		return nil, fmt.Errorf("failed to get nonce: %w", err)
	}

	gasPrice, err := c.eth.SuggestGasPrice(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to suggest gas price: %w", err)
	}

	auth, err := bind.NewKeyedTransactorWithChainID(c.signer, c.chainID)
	if err != nil {
		return nil, fmt.Errorf("failed to create transactor: %w", err)
	}

	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)
	auth.GasLimit = uint64(1000000) // Increased gas limit for proposeAI calls
	auth.GasPrice = gasPrice
	auth.Context = ctx

	return auth, nil
}
