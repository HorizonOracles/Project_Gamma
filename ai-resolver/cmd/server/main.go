package main

import (
	"context"
	"crypto/ecdsa"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"log"
	"math/big"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/project-gamma/ai-resolver/internal/adapter"
	"github.com/project-gamma/ai-resolver/internal/config"
	"github.com/project-gamma/ai-resolver/internal/eip712"
	"github.com/project-gamma/ai-resolver/internal/llm"
	"github.com/project-gamma/ai-resolver/pkg/abi"
)

func main() {
	// Load configuration
	cfg, err := config.LoadFromEnv()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	// Initialize components
	ctx := context.Background()

	// Initialize blockchain client
	client, err := adapter.NewClient(ctx, adapter.Config{
		RPCURL:              cfg.RPCEndpoint,
		ChainID:             cfg.ChainID,
		SignerPrivateKey:    cfg.SignerPrivateKey,
		AdapterAddress:      cfg.AIOracleAdapterAddr,
		FactoryAddress:      cfg.MarketFactoryAddr,
		ResolutionAddress:   cfg.ResolutionModuleAddr,
		HorizonTokenAddress: cfg.HorizonTokenAddr,
	})
	if err != nil {
		log.Fatalf("Failed to initialize blockchain client: %v", err)
	}
	defer client.Close()

	// Initialize LLM pipeline with integrated web search
	llmPipeline := llm.NewOpenAIPipeline(cfg.OpenAIAPIKey, cfg.OpenAIModel)

	// Parse private key for signing
	privateKey, err := crypto.HexToECDSA(cfg.SignerPrivateKey)
	if err != nil {
		log.Fatalf("Failed to parse private key: %v", err)
	}

	// Initialize EIP-712 signer
	eip712Signer := eip712.NewSigner(big.NewInt(cfg.ChainID), common.HexToAddress(cfg.AIOracleAdapterAddr))

	// Initialize server
	srv := &Server{
		config:     cfg,
		client:     client,
		llm:        llmPipeline,
		signer:     eip712Signer,
		privateKey: privateKey,
	}

	// Create HTTP server
	httpServer := &http.Server{
		Addr:         fmt.Sprintf("%s:%s", cfg.ServerHost, cfg.ServerPort),
		Handler:      srv.routes(),
		ReadTimeout:  60 * time.Second,
		WriteTimeout: 60 * time.Second,
		IdleTimeout:  120 * time.Second,
	}

	// Start server in a goroutine
	go func() {
		log.Printf("Starting AI Resolver server on %s", httpServer.Addr)
		log.Printf("Chain ID: %d, Signer: %s", cfg.ChainID, client.GetSignerAddress().Hex())
		if err := httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Server failed: %v", err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Println("Shutting down server...")

	// Graceful shutdown with timeout
	shutdownCtx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	if err := httpServer.Shutdown(shutdownCtx); err != nil {
		log.Fatalf("Server forced to shutdown: %v", err)
	}

	log.Println("Server stopped")
}

// Server holds the application state
type Server struct {
	config     *config.Config
	client     *adapter.Client
	llm        llm.Pipeline
	signer     *eip712.Signer
	privateKey *ecdsa.PrivateKey
}

// routes sets up the HTTP routes
func (s *Server) routes() http.Handler {
	mux := http.NewServeMux()

	// Health check
	mux.HandleFunc("/healthz", s.handleHealth)
	mux.HandleFunc("/v1/healthz", s.handleHealth)

	// API endpoints
	mux.HandleFunc("/v1/propose", s.handlePropose)
	mux.HandleFunc("/v1/markets", s.handleMarkets)

	// Wrap with middleware
	return s.corsMiddleware(s.loggingMiddleware(mux))
}

// handleHealth returns the health status
func (s *Server) handleHealth(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	response := map[string]any{
		"status":  "healthy",
		"version": "1.0.0",
		"time":    time.Now().Unix(),
		"signer":  s.client.GetSignerAddress().Hex(),
		"chainId": s.client.GetChainID().Int64(),
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// handlePropose handles market proposal requests
func (s *Server) handlePropose(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req struct {
		MarketID      uint64   `json:"marketId"`
		CloseTime     int64    `json:"closeTime"`
		Question      string   `json:"question"`
		OutcomeTokens []string `json:"outcomeTokens"`
		Metadata      string   `json:"metadata"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if req.MarketID == 0 {
		http.Error(w, "marketId is required", http.StatusBadRequest)
		return
	}

	if req.Question == "" {
		http.Error(w, "question is required", http.StatusBadRequest)
		return
	}

	ctx, cancel := context.WithTimeout(r.Context(), s.config.ProposalTimeout)
	defer cancel()

	// Execute proposal pipeline
	result, err := s.processProposal(ctx, req.MarketID, req.Question)
	if err != nil {
		log.Printf("Failed to process proposal for market %d: %v", req.MarketID, err)
		http.Error(w, fmt.Sprintf("Failed to process proposal: %v", err), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(result)
}

// processProposal executes the full AI resolution pipeline
// Updated: 2025-10-28 to accept question parameter
func (s *Server) processProposal(ctx context.Context, marketID uint64, question string) (map[string]any, error) {
	marketIDBig := big.NewInt(int64(marketID))

	// Step 1: Fetch market details
	log.Printf("Fetching market %d details...", marketID)
	market, err := s.client.GetMarket(ctx, marketIDBig)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch market: %w", err)
	}

	// Log the closeTime from the contract
	log.Printf("Market closeTime from contract: %s (%d)", market.CloseTime.String(), market.CloseTime.Int64())
	log.Printf("Current time: %d", time.Now().Unix())

	// Build market info for LLM
	marketInfo := llm.MarketInfo{
		MarketID:     marketID,
		Question:     question, // Use the question from the request
		Description:  "",
		Category:     market.Category,
		CloseTime:    market.CloseTime.Int64(),
		MetadataURI:  market.MetadataURI,
		OutcomeCount: 2,
	}
	log.Printf("Market: %s (Category: %s)", marketInfo.Question, marketInfo.Category)

	// Step 2: Run LLM analysis with integrated web search
	log.Printf("Running LLM multi-pass analysis with web search...")
	decision, err := s.llm.AnalyzeMarket(ctx, marketInfo)
	if err != nil {
		return nil, fmt.Errorf("failed to analyze: %w", err)
	}
	log.Printf("LLM decision: outcomeId=%d, confidence=%.2f", decision.OutcomeID, decision.Confidence)

	// Step 3: Prepare evidence hash and URIs
	evidenceURIs := make([]string, 0, len(decision.Citations))
	for _, citation := range decision.Citations {
		evidenceURIs = append(evidenceURIs, citation.URL)
	}

	evidenceHash := eip712.ComputeEvidenceHash(evidenceURIs)
	log.Printf("Evidence hash: %x", evidenceHash)

	// Step 4: Create proposal and sign
	// IMPORTANT: Use blockchain timestamp instead of system time
	blockchainTime, err := s.client.GetCurrentBlockTimestamp(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get blockchain timestamp: %w", err)
	}
	log.Printf("Blockchain time: %d, System time: %d", blockchainTime, time.Now().Unix())

	now := blockchainTime
	outcomeID := big.NewInt(int64(decision.OutcomeID))

	proposal := eip712.ProposedOutcome{
		MarketID:     marketIDBig,
		OutcomeID:    outcomeID,
		CloseTime:    market.CloseTime,
		EvidenceHash: evidenceHash,
		NotBefore:    big.NewInt(now),
		Deadline:     big.NewInt(now + 7200), // 2 hour validity to account for LLM processing time
	}

	signature, err := s.signer.SignProposal(proposal, s.privateKey)
	if err != nil {
		return nil, fmt.Errorf("failed to sign proposal: %w", err)
	}
	log.Printf("Signature: %x", signature)

	// Step 5: Check allowance and approve if needed
	bondAmountBig := new(big.Int)
	bondAmountBig.SetString(s.config.DefaultBondAmount, 10)

	allowance, err := s.client.CheckAllowance(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to check allowance: %w", err)
	}

	if allowance.Cmp(bondAmountBig) < 0 {
		log.Printf("Approving bond amount: %s", bondAmountBig.String())
		approveTx, err := s.client.ApproveBond(ctx, bondAmountBig)
		if err != nil {
			return nil, fmt.Errorf("failed to approve bond: %w", err)
		}
		log.Printf("Approve tx: %s", approveTx.Hash().Hex())

		// Wait for approval
		_, err = s.client.WaitForTransaction(ctx, approveTx)
		if err != nil {
			return nil, fmt.Errorf("approval transaction failed: %w", err)
		}
		log.Printf("Approval confirmed")
	}

	// Step 6: Submit proposal
	log.Printf("Submitting proposal to blockchain...")

	// Convert eip712.ProposedOutcome to abi.AIOracleAdapterProposedOutcome
	abiProposal := abi.AIOracleAdapterProposedOutcome{
		MarketId:     proposal.MarketID,
		OutcomeId:    proposal.OutcomeID,
		CloseTime:    proposal.CloseTime,
		EvidenceHash: proposal.EvidenceHash,
		NotBefore:    proposal.NotBefore,
		Deadline:     proposal.Deadline,
	}

	// Log the proposal being submitted
	log.Printf("\n=== SUBMITTING PROPOSAL TO BLOCKCHAIN ===")
	log.Printf("MarketID:     %s", abiProposal.MarketId.String())
	log.Printf("OutcomeID:    %s", abiProposal.OutcomeId.String())
	log.Printf("CloseTime:    %s (%d)", abiProposal.CloseTime.String(), abiProposal.CloseTime.Int64())
	log.Printf("EvidenceHash: %x", abiProposal.EvidenceHash)
	log.Printf("NotBefore:    %s (%d)", abiProposal.NotBefore.String(), abiProposal.NotBefore.Int64())
	log.Printf("Deadline:     %s (%d)", abiProposal.Deadline.String(), abiProposal.Deadline.Int64())
	log.Printf("Signature:    %x", signature)
	log.Printf("BondAmount:   %s", bondAmountBig.String())
	log.Printf("EvidenceURIs: %v", evidenceURIs)
	log.Printf("=========================================\n")

	tx, err := s.client.ProposeOutcome(ctx, abiProposal, signature, bondAmountBig, evidenceURIs)
	if err != nil {
		return nil, fmt.Errorf("failed to submit proposal: %w", err)
	}
	log.Printf("Proposal tx: %s", tx.Hash().Hex())

	// Wait for confirmation (optional - could be async)
	receipt, err := s.client.WaitForTransaction(ctx, tx)
	if err != nil {
		log.Printf("Warning: failed to wait for transaction: %v", err)
		// Continue anyway - tx might still succeed
	} else {
		log.Printf("Proposal confirmed in block %d", receipt.BlockNumber.Uint64())
	}

	// Return result
	return map[string]any{
		"status":       "submitted",
		"marketId":     marketID,
		"outcomeId":    decision.OutcomeID,
		"confidence":   decision.Confidence,
		"reasoning":    decision.Reasoning,
		"txHash":       tx.Hash().Hex(),
		"evidenceHash": hex.EncodeToString(evidenceHash[:]),
		"citations":    len(decision.Citations),
		"facts":        len(decision.Facts),
	}, nil
}

// handleMarkets returns pending markets
func (s *Server) handleMarkets(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// TODO: Query pending markets from blockchain
	// This would require additional contract methods to list markets by status
	response := map[string]any{
		"markets": []any{},
		"count":   0,
		"message": "Market listing not yet implemented",
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// loggingMiddleware logs HTTP requests
func (s *Server) loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		log.Printf("%s %s from %s", r.Method, r.URL.Path, r.RemoteAddr)
		next.ServeHTTP(w, r)
		log.Printf("Completed in %v", time.Since(start))
	})
}

// corsMiddleware adds CORS headers
func (s *Server) corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}
