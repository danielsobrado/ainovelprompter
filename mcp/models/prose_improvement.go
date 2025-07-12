package models

import (
	"time"

	"github.com/google/uuid"
)

// ProsePromptVariant represents a variant of a prose improvement prompt
type ProsePromptVariant struct {
	VariantLabel         string   `json:"variantLabel,omitempty"`
	TargetModelFamilies  []string `json:"targetModelFamilies,omitempty"` // e.g., ["claude", "gpt", "ollama"]
	TargetModels         []string `json:"targetModels,omitempty"`        // e.g., ["anthropic/claude-3-opus", "openai/gpt-4o"]
	PromptText           string   `json:"promptText"`
}

// ProseImprovementPrompt represents a prompt definition for prose improvement
type ProseImprovementPrompt struct {
	ID                string              `json:"id"`
	Label             string              `json:"label"`
	Category          string              `json:"category"` // 'tropes' | 'style' | 'grammar' | 'custom'
	Order             int                 `json:"order"`
	Description       string              `json:"description,omitempty"`
	DefaultPromptText string              `json:"defaultPromptText"`
	Variants          []ProsePromptVariant `json:"variants"`
}

// ProseChange represents a suggested change to text
type ProseChange struct {
	ID            string     `json:"id"`
	Initial       string     `json:"initial"`
	Improved      string     `json:"improved"`
	Reason        string     `json:"reason"`
	TropeCategory string     `json:"trope_category,omitempty"`
	Status        string     `json:"status"` // 'pending' | 'accepted' | 'rejected'
	StartIndex    *int       `json:"startIndex,omitempty"`
	EndIndex      *int       `json:"endIndex,omitempty"`
	CreatedAt     time.Time  `json:"createdAt"`
	UpdatedAt     time.Time  `json:"updatedAt"`
}

// ProseImprovementSession represents an active prose improvement session
type ProseImprovementSession struct {
	ID                  string                    `json:"id"`
	OriginalText        string                    `json:"originalText"`
	CurrentText         string                    `json:"currentText"`
	Prompts             []ProseImprovementPrompt  `json:"prompts"`
	CurrentPromptIndex  int                       `json:"currentPromptIndex"`
	Changes             []ProseChange             `json:"changes"`
	CreatedAt           time.Time                 `json:"createdAt"`
	UpdatedAt           time.Time                 `json:"updatedAt"`
}

// ProseAnalysisRequest represents a request to analyze prose
type ProseAnalysisRequest struct {
	Text     string `json:"text"`
	PromptID string `json:"promptId"`
	Provider string `json:"provider,omitempty"`
}

// ProseAnalysisResult represents the result of prose analysis
type ProseAnalysisResult struct {
	OriginalText string        `json:"originalText"`
	Changes      []ProseChange `json:"changes"`
	PromptUsed   string        `json:"promptUsed"`
	Provider     string        `json:"provider"`
	ProcessedAt  time.Time     `json:"processedAt"`
}

// SearchResult represents a search result across all content
type SearchResult struct {
	ID          string                 `json:"id"`
	Type        string                 `json:"type"`        // 'character', 'location', 'codex', 'chapter', 'rule', etc.
	Title       string                 `json:"title"`
	Content     string                 `json:"content"`
	Snippet     string                 `json:"snippet"`     // Highlighted excerpt
	Score       float64                `json:"score"`       // Relevance score
	Metadata    map[string]interface{} `json:"metadata"`    // Additional type-specific data
	MatchedAt   time.Time              `json:"matchedAt"`
}

// GenerateProseChangeID creates a new UUID for prose changes
func GenerateProseChangeID() string {
	return uuid.New().String()
}

// GenerateSessionID creates a new UUID for prose improvement sessions
func GenerateSessionID() string {
	return uuid.New().String()
}

// NewProseChange creates a new prose change with current timestamp
func NewProseChange(initial, improved, reason, category string) *ProseChange {
	now := time.Now()
	return &ProseChange{
		ID:            GenerateProseChangeID(),
		Initial:       initial,
		Improved:      improved,
		Reason:        reason,
		TropeCategory: category,
		Status:        "pending",
		CreatedAt:     now,
		UpdatedAt:     now,
	}
}

// NewProseSession creates a new prose improvement session
func NewProseSession(originalText string, prompts []ProseImprovementPrompt) *ProseImprovementSession {
	now := time.Now()
	return &ProseImprovementSession{
		ID:                 GenerateSessionID(),
		OriginalText:       originalText,
		CurrentText:        originalText,
		Prompts:            prompts,
		CurrentPromptIndex: 0,
		Changes:            []ProseChange{},
		CreatedAt:          now,
		UpdatedAt:          now,
	}
}
