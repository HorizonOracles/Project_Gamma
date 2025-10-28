package llm

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"
)

// OpenAIPipeline implements the multi-pass analysis pipeline using OpenAI with web search
type OpenAIPipeline struct {
	apiKey     string
	model      string
	httpClient *http.Client
	baseURL    string
}

// NewOpenAIPipeline creates a new OpenAI-based pipeline
func NewOpenAIPipeline(apiKey, model string) *OpenAIPipeline {
	return &OpenAIPipeline{
		apiKey: apiKey,
		model:  model,
		httpClient: &http.Client{
			Timeout: 120 * time.Second, // Increased for web search
		},
		baseURL: "https://api.openai.com/v1/chat/completions",
	}
}

// AnalyzeMarket performs the complete multi-pass analysis with integrated web search
func (p *OpenAIPipeline) AnalyzeMarket(ctx context.Context, market MarketInfo) (*Decision, error) {
	// Build search query from market information
	searchQuery := p.buildSearchQuery(market)

	// Step 1: Search and extract facts using OpenAI web search
	facts, webSources, err := p.searchAndExtractFacts(ctx, market, searchQuery)
	if err != nil {
		return nil, fmt.Errorf("failed to search and extract facts: %w", err)
	}

	// Step 2: Check for contradictions
	facts, err = p.checkContradictions(ctx, market, facts)
	if err != nil {
		return nil, fmt.Errorf("failed to check contradictions: %w", err)
	}

	// Step 3: Decide outcome based on facts
	decision, err := p.decideOutcome(ctx, market, facts)
	if err != nil {
		return nil, fmt.Errorf("failed to decide outcome: %w", err)
	}

	// Step 4: Build citations from web sources
	decision.Citations = p.buildCitationsFromSources(webSources, facts)
	decision.Timestamp = time.Now().Unix()

	return decision, nil
}

// buildSearchQuery creates an effective search query from market info
func (p *OpenAIPipeline) buildSearchQuery(market MarketInfo) string {
	// Use the question as the primary search query
	query := market.Question

	// If we have additional context from description, include key terms
	if market.Description != "" && len(market.Description) < 200 {
		query = market.Question + " " + market.Description
	}

	return query
}

// searchAndExtractFacts uses OpenAI web search to find and extract facts
func (p *OpenAIPipeline) searchAndExtractFacts(ctx context.Context, market MarketInfo, searchQuery string) ([]Fact, []WebSource, error) {
	prompt := fmt.Sprintf(`You are analyzing evidence to resolve a prediction market question. Use web search to find current information.

Question: %s
Description: %s
Category: %s

Task: Search the web for information about this question, then extract key facts that are relevant to answering it. For each fact:
1. State the fact clearly
2. List the sources (URLs) that support it
3. Rate your confidence (0-1)
4. Provide supporting evidence (brief quote or summary)

IMPORTANT: Output ONLY a valid JSON object, no other text before or after. Use this exact format:
{
  "facts": [
    {
      "statement": "clear factual statement",
      "sources": ["url1", "url2"],
      "confidence": 0.95,
      "supportingEvidence": "brief quote or summary"
    }
  ],
  "sources": [
    {
      "url": "https://example.com/article",
      "title": "Article Title",
      "snippet": "Relevant excerpt from the article"
    }
  ]
}

Focus on facts that are:
- Verifiable and specific
- Directly relevant to the question
- From credible sources
- Recent and timely

Search query to use: %s`, market.Question, market.Description, market.Category, searchQuery)

	response, err := p.callOpenAIWithWebSearch(ctx, prompt, 0.3)
	if err != nil {
		return nil, nil, err
	}

	// Extract JSON from response (may have extra text)
	jsonStr := extractJSON(response)

	// Parse JSON response
	var result struct {
		Facts   []Fact      `json:"facts"`
		Sources []WebSource `json:"sources"`
	}
	if err := json.Unmarshal([]byte(jsonStr), &result); err != nil {
		return nil, nil, fmt.Errorf("failed to parse facts JSON: %w\nResponse: %s", err, response)
	}

	return result.Facts, result.Sources, nil
}

// checkContradictions checks for contradictory facts using standard chat API
func (p *OpenAIPipeline) checkContradictions(ctx context.Context, market MarketInfo, facts []Fact) ([]Fact, error) {
	if len(facts) == 0 {
		return facts, nil
	}

	factsJSON, _ := json.MarshalIndent(facts, "", "  ")
	prompt := fmt.Sprintf(`You are reviewing extracted facts for contradictions.

Question: %s

Extracted Facts:
%s

Task: Identify any facts that contradict each other. Return the same JSON array but with "contradicts" set to true for any contradictory facts.

Consider facts contradictory if they make opposing claims about the same aspect of the question.

Return the JSON array with the contradicts field updated.`, market.Question, string(factsJSON))

	response, err := p.callOpenAIChat(ctx, prompt, 0.2)
	if err != nil {
		return nil, err
	}

	jsonStr := extractJSON(response)
	var updatedFacts []Fact
	if err := json.Unmarshal([]byte(jsonStr), &updatedFacts); err != nil {
		return facts, nil // Return original on parse error
	}

	return updatedFacts, nil
}

// decideOutcome makes the final decision based on facts using standard chat API
func (p *OpenAIPipeline) decideOutcome(ctx context.Context, market MarketInfo, facts []Fact) (*Decision, error) {
	factsJSON, _ := json.MarshalIndent(facts, "", "  ")
	prompt := fmt.Sprintf(`You are making a final decision on a prediction market question.

Question: %s
Description: %s

Analyzed Facts:
%s

Task: Decide the outcome and provide reasoning.

For binary markets:
- outcomeId: 0 = NO (did not happen, false)
- outcomeId: 1 = YES (did happen, true)

Return JSON in this exact format:
{
  "outcomeId": 0 or 1,
  "confidence": 0.0 to 1.0,
  "reasoning": "clear explanation of why this outcome is correct",
  "facts": [copy the facts array here]
}

Base your decision on:
1. Weight of evidence
2. Source credibility
3. Fact confidence scores
4. Resolution of contradictions
5. Completeness of information

Be conservative - if evidence is insufficient or contradictory, reduce confidence accordingly.`,
		market.Question, market.Description, string(factsJSON))

	response, err := p.callOpenAIChat(ctx, prompt, 0.4)
	if err != nil {
		return nil, err
	}

	jsonStr := extractJSON(response)
	var decision Decision
	if err := json.Unmarshal([]byte(jsonStr), &decision); err != nil {
		return nil, fmt.Errorf("failed to parse decision JSON: %w", err)
	}

	// Validation
	if decision.OutcomeID > 1 {
		return nil, fmt.Errorf("invalid outcome ID: %d (must be 0 or 1)", decision.OutcomeID)
	}
	if decision.Confidence < 0 || decision.Confidence > 1 {
		return nil, fmt.Errorf("invalid confidence: %f (must be 0-1)", decision.Confidence)
	}

	return &decision, nil
}

// buildCitationsFromSources creates citations from web sources and facts
func (p *OpenAIPipeline) buildCitationsFromSources(sources []WebSource, facts []Fact) []Citation {
	// Build a map of URLs mentioned in facts with their weights
	urlWeight := make(map[string]float64)
	for _, fact := range facts {
		weight := fact.Confidence / float64(len(fact.Sources))
		for _, url := range fact.Sources {
			urlWeight[url] += weight
		}
	}

	// Create citations from sources that are referenced in facts
	citations := make([]Citation, 0)
	for _, source := range sources {
		if weight, ok := urlWeight[source.URL]; ok && weight > 0 {
			citations = append(citations, Citation{
				URL:     source.URL,
				Title:   source.Title,
				Snippet: source.Snippet,
				Weight:  weight,
			})
		}
	}

	// If no citations matched, include all sources with default weight
	if len(citations) == 0 && len(sources) > 0 {
		for _, source := range sources {
			citations = append(citations, Citation{
				URL:     source.URL,
				Title:   source.Title,
				Snippet: source.Snippet,
				Weight:  0.5,
			})
		}
	}

	return citations
}

// callOpenAIWithWebSearch makes a request to OpenAI Responses API with web search enabled
func (p *OpenAIPipeline) callOpenAIWithWebSearch(ctx context.Context, prompt string, temperature float64) (string, error) {
	// Use the responses API endpoint
	responsesURL := "https://api.openai.com/v1/responses"

	reqBody := map[string]any{
		"model": p.model,
		"input": prompt,
		"tools": []map[string]string{
			{"type": "web_search_preview"},
		},
		"temperature": temperature,
	}

	jsonData, err := json.Marshal(reqBody)
	if err != nil {
		return "", fmt.Errorf("failed to marshal request: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, "POST", responsesURL, bytes.NewBuffer(jsonData))
	if err != nil {
		return "", fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Authorization", "Bearer "+p.apiKey)
	req.Header.Set("Content-Type", "application/json")

	resp, err := p.httpClient.Do(req)
	if err != nil {
		return "", fmt.Errorf("request failed: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("failed to read response: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("API returned status %d: %s", resp.StatusCode, string(body))
	}

	var apiResp openAIResponsesAPIResponse
	if err := json.Unmarshal(body, &apiResp); err != nil {
		return "", fmt.Errorf("failed to parse response: %w", err)
	}

	if len(apiResp.Output) == 0 {
		return "", fmt.Errorf("no output in response")
	}

	// Extract text from message outputs
	var resultText string
	for _, output := range apiResp.Output {
		if output.Type == "message" && len(output.Content) > 0 {
			for _, content := range output.Content {
				if content.Type == "output_text" {
					resultText = content.Text
					break
				}
			}
			if resultText != "" {
				break
			}
		}
	}

	if resultText == "" {
		return "", fmt.Errorf("no text content found in response")
	}

	return resultText, nil
}

// callOpenAIChat makes a standard chat completion request (for non-search steps)
func (p *OpenAIPipeline) callOpenAIChat(ctx context.Context, prompt string, temperature float64) (string, error) {
	chatURL := strings.Replace(p.baseURL, "/responses", "/chat/completions", 1)

	reqBody := map[string]any{
		"model": p.model,
		"messages": []map[string]string{
			{
				"role":    "system",
				"content": "You are a precise, factual AI assistant analyzing evidence for prediction markets. Always respond with valid JSON.",
			},
			{
				"role":    "user",
				"content": prompt,
			},
		},
		"temperature": temperature,
		"max_tokens":  2000,
	}

	jsonData, err := json.Marshal(reqBody)
	if err != nil {
		return "", fmt.Errorf("failed to marshal request: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, "POST", chatURL, bytes.NewBuffer(jsonData))
	if err != nil {
		return "", fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Authorization", "Bearer "+p.apiKey)
	req.Header.Set("Content-Type", "application/json")

	resp, err := p.httpClient.Do(req)
	if err != nil {
		return "", fmt.Errorf("request failed: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("failed to read response: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("API returned status %d: %s", resp.StatusCode, string(body))
	}

	var apiResp openAIChatResponse
	if err := json.Unmarshal(body, &apiResp); err != nil {
		return "", fmt.Errorf("failed to parse response: %w", err)
	}

	if len(apiResp.Choices) == 0 {
		return "", fmt.Errorf("no choices in response")
	}

	return apiResp.Choices[0].Message.Content, nil
}

// openAIResponsesAPIResponse represents OpenAI Responses API response
type openAIResponsesAPIResponse struct {
	ID     string                     `json:"id"`
	Output []openAIResponsesAPIOutput `json:"output"`
	Model  string                     `json:"model"`
	Status string                     `json:"status"`
}

type openAIResponsesAPIOutput struct {
	ID      string                      `json:"id"`
	Type    string                      `json:"type"`
	Status  string                      `json:"status"`
	Content []openAIResponsesAPIContent `json:"content,omitempty"`
	Role    string                      `json:"role,omitempty"`
}

type openAIResponsesAPIContent struct {
	Type        string                         `json:"type"`
	Text        string                         `json:"text"`
	Annotations []openAIResponsesAPIAnnotation `json:"annotations,omitempty"`
}

type openAIResponsesAPIAnnotation struct {
	Type       string `json:"type"`
	StartIndex int    `json:"start_index"`
	EndIndex   int    `json:"end_index"`
	Title      string `json:"title"`
	URL        string `json:"url"`
}

// openAIChatResponse represents OpenAI Chat API response
type openAIChatResponse struct {
	Choices []struct {
		Message struct {
			Content string `json:"content"`
		} `json:"message"`
	} `json:"choices"`
}

// WebSource represents a web source found during search
type WebSource struct {
	URL     string `json:"url"`
	Title   string `json:"title"`
	Snippet string `json:"snippet"`
}

// extractJSON attempts to extract a JSON object from text that may contain additional content
func extractJSON(text string) string {
	// Try to find JSON object boundaries
	start := strings.Index(text, "{")
	if start == -1 {
		return text
	}

	// Find matching closing brace
	depth := 0
	for i := start; i < len(text); i++ {
		switch text[i] {
		case '{':
			depth++
		case '}':
			depth--
			if depth == 0 {
				return text[start : i+1]
			}
		}
	}

	return text
}
