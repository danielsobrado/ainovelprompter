package storage

import (
	"github.com/danielsobrado/ainovelprompter/mcp/models"
)

type StoryContextStorage interface {
	// Characters
	GetCharacters() ([]models.Character, error)
	GetCharacterByID(id string) (*models.Character, error)
	CreateCharacter(character *models.Character) error
	UpdateCharacter(character *models.Character) error
	DeleteCharacter(id string) error
	SearchCharacters(query string) ([]models.Character, error)

	// Locations
	GetLocations() ([]models.Location, error)
	GetLocationByID(id string) (*models.Location, error)
	CreateLocation(location *models.Location) error
	UpdateLocation(location *models.Location) error
	DeleteLocation(id string) error
	SearchLocations(query string) ([]models.Location, error)

	// Codex
	GetCodexEntries() ([]models.CodexEntry, error)
	GetCodexEntryByID(id string) (*models.CodexEntry, error)
	CreateCodexEntry(entry *models.CodexEntry) error
	UpdateCodexEntry(entry *models.CodexEntry) error
	DeleteCodexEntry(id string) error
	SearchCodex(query string) ([]models.CodexEntry, error)

	// Rules
	GetRules() ([]models.Rule, error)
	GetRuleByID(id string) (*models.Rule, error)
	GetActiveRules() ([]models.Rule, error)
}

type ChapterStorage interface {
	// Chapters
	GetChapters() ([]models.Chapter, error)
	GetChapterByID(id string) (*models.Chapter, error)
	GetChapterByNumber(number int) (*models.Chapter, error)
	GetPreviousChapter(currentNumber int) (*models.Chapter, error)
	GetChapterRange(start, end int) ([]models.Chapter, error)
	CreateChapter(chapter *models.Chapter) error
	UpdateChapter(chapter *models.Chapter) error
	DeleteChapter(id string) error
	SearchChapters(query string) ([]models.Chapter, error)

	// Story Beats
	GetStoryBeats(chapterNumber int) (*models.StoryBeats, error)
	GetAllStoryBeats() ([]models.StoryBeats, error)
	SaveStoryBeats(beats *models.StoryBeats) error

	// Future Notes
	GetFutureNotes() ([]models.FutureNotes, error)
	GetFutureNotesByChapter(chapterNum int) ([]models.FutureNotes, error)
	CreateFutureNote(note *models.FutureNotes) error
	UpdateFutureNote(note *models.FutureNotes) error
	DeleteFutureNote(id string) error

	// Sample Chapters
	GetSampleChapters() ([]models.SampleChapter, error)
	GetSampleChapterByID(id string) (*models.SampleChapter, error)
	GetSampleChaptersByPurpose(purpose string) ([]models.SampleChapter, error)
	CreateSampleChapter(sample *models.SampleChapter) error

	// Task Types
	GetTaskTypes() ([]models.TaskType, error)
	GetTaskTypeByID(id string) (*models.TaskType, error)
	GetActiveTaskTypes() ([]models.TaskType, error)
}

// Storage provides basic file operations
type Storage interface {
	ReadFile(filename string) ([]byte, error)
	WriteFile(filename string, data []byte) error
	DeleteFile(filename string) error
	ListFiles(pattern string) ([]string, error)
}

// CombinedStorage provides both story context and chapter storage
type CombinedStorage interface {
	StoryContextStorage
	ChapterStorage
	Storage
}
