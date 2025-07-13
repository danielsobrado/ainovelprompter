package models

import (
	"time"
)

// Chapter represents a story chapter
type Chapter struct {
	ID           string    `json:"id"`
	Number       int       `json:"number"`
	Title        string    `json:"title"`
	Content      string    `json:"content"`
	StoryBeats   []string  `json:"storyBeats,omitempty"`
	Summary      string    `json:"summary,omitempty"`
	WordCount    int       `json:"wordCount"`
	Status       string    `json:"status"` // draft, review, final
	Tags         []string  `json:"tags,omitempty"`
	CharacterIDs []string  `json:"characterIds,omitempty"`
	LocationIDs  []string  `json:"locationIds,omitempty"`
	CreatedAt    time.Time `json:"createdAt"`
	UpdatedAt    time.Time `json:"updatedAt"`
}

// StoryBeats represents planned beats for future chapters
type StoryBeats struct {
	ID            string    `json:"id"`
	ChapterNumber int       `json:"chapterNumber"`
	Beats         []Beat    `json:"beats"`
	Notes         string    `json:"notes,omitempty"`
	UpdatedAt     time.Time `json:"updatedAt"`
}

// Beat represents a single story beat
type Beat struct {
	ID          string `json:"id"`
	Description string `json:"description"`
	Type        string `json:"type"` // scene, transition, revelation, conflict
	Completed   bool   `json:"completed"`
	Order       int    `json:"order"`
}

// FutureNotes represents planned developments
type FutureNotes struct {
	ID           string    `json:"id"`
	ChapterRange string    `json:"chapterRange"` // e.g., "10-15"
	Content      string    `json:"content"`
	Priority     string    `json:"priority"` // high, medium, low
	Category     string    `json:"category"` // plot, character, world
	Tags         []string  `json:"tags,omitempty"`
	CreatedAt    time.Time `json:"createdAt"`
	UpdatedAt    time.Time `json:"updatedAt"`
}

// SampleChapter represents reference chapters for style
type SampleChapter struct {
	ID        string    `json:"id"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	Author    string    `json:"author,omitempty"`
	Source    string    `json:"source,omitempty"`
	Tags      []string  `json:"tags,omitempty"`
	Purpose   string    `json:"purpose"` // style_reference, tone_example, pacing_guide
	CreatedAt time.Time `json:"createdAt"`
}

// TaskType represents writing task templates
type TaskType struct {
	ID          string            `json:"id"`
	Name        string            `json:"name"`
	Description string            `json:"description"`
	Template    string            `json:"template"`
	Variables   map[string]string `json:"variables,omitempty"`
	Category    string            `json:"category"`
	Active      bool              `json:"active"`
}
