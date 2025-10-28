package llm

import (
	"context"
)

// Pipeline defines the multi-pass LLM analysis pipeline
type Pipeline interface {
	// AnalyzeMarket performs the complete analysis pipeline using integrated web search
	AnalyzeMarket(ctx context.Context, market MarketInfo) (*Decision, error)
}

// MarketInfo contains information about the market to be resolved
type MarketInfo struct {
	MarketID     uint64 `json:"marketId"`
	Question     string `json:"question"`
	Description  string `json:"description"`
	Category     string `json:"category"`
	CloseTime    int64  `json:"closeTime"`
	MetadataURI  string `json:"metadataUri"`
	OutcomeCount int    `json:"outcomeCount"` // 2 for binary (YES/NO)
}

// Decision represents the final outcome decision with evidence
type Decision struct {
	OutcomeID  uint64     `json:"outcomeId"`  // 0 = NO, 1 = YES (for binary markets)
	Confidence float64    `json:"confidence"` // 0-1 confidence score
	Reasoning  string     `json:"reasoning"`  // Explanation of decision
	Citations  []Citation `json:"citations"`  // Evidence citations
	Facts      []Fact     `json:"facts"`      // Extracted facts
	Timestamp  int64      `json:"timestamp"`  // Unix timestamp of decision
}

// Citation represents a source citation
type Citation struct {
	URL     string  `json:"url"`
	Title   string  `json:"title"`
	Snippet string  `json:"snippet"`
	Weight  float64 `json:"weight"` // Importance weight (0-1)
}

// Fact represents an extracted fact from sources
type Fact struct {
	Statement          string   `json:"statement"`
	Sources            []string `json:"sources"`     // URLs supporting this fact
	Confidence         float64  `json:"confidence"`  // 0-1 confidence in fact
	Contradicts        bool     `json:"contradicts"` // Whether this contradicts other facts
	SupportingEvidence string   `json:"supportingEvidence"`
}

// AnalysisStep represents a step in the multi-pass pipeline
type AnalysisStep string

const (
	StepExtractFacts        AnalysisStep = "extract_facts"
	StepCheckContradictions AnalysisStep = "check_contradictions"
	StepDecideOutcome       AnalysisStep = "decide_outcome"
	StepBuildCitations      AnalysisStep = "build_citations"
)
