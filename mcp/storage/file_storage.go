package storage

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/danielsobrado/ainovelprompter/mcp/models"
	"github.com/google/uuid"
)

type FileStorage struct {
	basePath string
	mu       sync.RWMutex
}

func NewFileStorage(basePath string) *FileStorage {
	return &FileStorage{
		basePath: basePath,
	}
}

// Characters implementation
func (fs *FileStorage) GetCharacters() ([]models.Character, error) {
	fs.mu.RLock()
	defer fs.mu.RUnlock()

	data, err := fs.readJSONFile("characters.json")
	if err != nil {
		return nil, err
	}

	var characters []models.Character
	if err := json.Unmarshal(data, &characters); err != nil {
		return nil, fmt.Errorf("error parsing characters: %v", err)
	}

	return characters, nil
}

func (fs *FileStorage) GetCharacterByID(id string) (*models.Character, error) {
	characters, err := fs.GetCharacters()
	if err != nil {
		return nil, err
	}

	for _, char := range characters {
		if char.ID == id {
			return &char, nil
		}
	}

	return nil, fmt.Errorf("character with ID %s not found", id)
}

func (fs *FileStorage) SearchCharacters(query string) ([]models.Character, error) {
	characters, err := fs.GetCharacters()
	if err != nil {
		return nil, err
	}

	query = strings.ToLower(query)
	var results []models.Character

	for _, char := range characters {
		if strings.Contains(strings.ToLower(char.Name), query) ||
			strings.Contains(strings.ToLower(char.Description), query) ||
			strings.Contains(strings.ToLower(char.Notes), query) {
			results = append(results, char)
		}
	}

	return results, nil
}

func (fs *FileStorage) CreateCharacter(character *models.Character) error {
	fs.mu.Lock()
	defer fs.mu.Unlock()

	characters, err := fs.GetCharacters()
	if err != nil && !os.IsNotExist(err) {
		return err
	}

	if character.ID == "" {
		character.ID = uuid.New().String()
	}
	character.CreatedAt = time.Now()
	character.UpdatedAt = time.Now()

	characters = append(characters, *character)

	return fs.writeJSONFile("characters.json", characters)
}

func (fs *FileStorage) UpdateCharacter(character *models.Character) error {
	fs.mu.Lock()
	defer fs.mu.Unlock()

	characters, err := fs.GetCharacters()
	if err != nil {
		return err
	}

	for i, char := range characters {
		if char.ID == character.ID {
			character.UpdatedAt = time.Now()
			characters[i] = *character
			return fs.writeJSONFile("characters.json", characters)
		}
	}

	return fmt.Errorf("character with ID %s not found", character.ID)
}

func (fs *FileStorage) DeleteCharacter(id string) error {
	fs.mu.Lock()
	defer fs.mu.Unlock()

	characters, err := fs.GetCharacters()
	if err != nil {
		return err
	}

	for i, char := range characters {
		if char.ID == id {
			characters = append(characters[:i], characters[i+1:]...)
			return fs.writeJSONFile("characters.json", characters)
		}
	}

	return fmt.Errorf("character with ID %s not found", id)
}

// Locations implementation
func (fs *FileStorage) GetLocations() ([]models.Location, error) {
	fs.mu.RLock()
	defer fs.mu.RUnlock()

	data, err := fs.readJSONFile("locations.json")
	if err != nil {
		return nil, err
	}

	var locations []models.Location
	if err := json.Unmarshal(data, &locations); err != nil {
		return nil, fmt.Errorf("error parsing locations: %v", err)
	}

	return locations, nil
}

func (fs *FileStorage) GetLocationByID(id string) (*models.Location, error) {
	locations, err := fs.GetLocations()
	if err != nil {
		return nil, err
	}

	for _, loc := range locations {
		if loc.ID == id {
			return &loc, nil
		}
	}

	return nil, fmt.Errorf("location with ID %s not found", id)
}

func (fs *FileStorage) SearchLocations(query string) ([]models.Location, error) {
	locations, err := fs.GetLocations()
	if err != nil {
		return nil, err
	}

	query = strings.ToLower(query)
	var results []models.Location

	for _, loc := range locations {
		if strings.Contains(strings.ToLower(loc.Name), query) ||
			strings.Contains(strings.ToLower(loc.Description), query) ||
			strings.Contains(strings.ToLower(loc.Details), query) {
			results = append(results, loc)
		}
	}

	return results, nil
}

func (fs *FileStorage) CreateLocation(location *models.Location) error {
	fs.mu.Lock()
	defer fs.mu.Unlock()

	locations, err := fs.GetLocations()
	if err != nil && !os.IsNotExist(err) {
		return err
	}

	if location.ID == "" {
		location.ID = uuid.New().String()
	}
	location.CreatedAt = time.Now()
	location.UpdatedAt = time.Now()

	locations = append(locations, *location)

	return fs.writeJSONFile("locations.json", locations)
}

func (fs *FileStorage) UpdateLocation(location *models.Location) error {
	fs.mu.Lock()
	defer fs.mu.Unlock()

	locations, err := fs.GetLocations()
	if err != nil {
		return err
	}

	for i, loc := range locations {
		if loc.ID == location.ID {
			location.UpdatedAt = time.Now()
			locations[i] = *location
			return fs.writeJSONFile("locations.json", locations)
		}
	}

	return fmt.Errorf("location with ID %s not found", location.ID)
}

func (fs *FileStorage) DeleteLocation(id string) error {
	fs.mu.Lock()
	defer fs.mu.Unlock()

	locations, err := fs.GetLocations()
	if err != nil {
		return err
	}

	for i, loc := range locations {
		if loc.ID == id {
			locations = append(locations[:i], locations[i+1:]...)
			return fs.writeJSONFile("locations.json", locations)
		}
	}

	return fmt.Errorf("location with ID %s not found", id)
}

// Codex implementation
func (fs *FileStorage) GetCodexEntries() ([]models.CodexEntry, error) {
	fs.mu.RLock()
	defer fs.mu.RUnlock()

	data, err := fs.readJSONFile("codex.json")
	if err != nil {
		return nil, err
	}

	var entries []models.CodexEntry
	if err := json.Unmarshal(data, &entries); err != nil {
		return nil, fmt.Errorf("error parsing codex: %v", err)
	}

	return entries, nil
}

func (fs *FileStorage) GetCodexEntryByID(id string) (*models.CodexEntry, error) {
	entries, err := fs.GetCodexEntries()
	if err != nil {
		return nil, err
	}

	for _, entry := range entries {
		if entry.ID == id {
			return &entry, nil
		}
	}

	return nil, fmt.Errorf("codex entry with ID %s not found", id)
}

func (fs *FileStorage) SearchCodex(query string) ([]models.CodexEntry, error) {
	entries, err := fs.GetCodexEntries()
	if err != nil {
		return nil, err
	}

	query = strings.ToLower(query)
	var results []models.CodexEntry

	for _, entry := range entries {
		if strings.Contains(strings.ToLower(entry.Title), query) ||
			strings.Contains(strings.ToLower(entry.Content), query) ||
			strings.Contains(strings.ToLower(entry.Category), query) {
			results = append(results, entry)
		}
	}

	return results, nil
}

func (fs *FileStorage) CreateCodexEntry(entry *models.CodexEntry) error {
	fs.mu.Lock()
	defer fs.mu.Unlock()

	entries, err := fs.GetCodexEntries()
	if err != nil && !os.IsNotExist(err) {
		return err
	}

	if entry.ID == "" {
		entry.ID = uuid.New().String()
	}
	entry.CreatedAt = time.Now()
	entry.UpdatedAt = time.Now()

	entries = append(entries, *entry)

	return fs.writeJSONFile("codex.json", entries)
}

func (fs *FileStorage) UpdateCodexEntry(entry *models.CodexEntry) error {
	fs.mu.Lock()
	defer fs.mu.Unlock()

	entries, err := fs.GetCodexEntries()
	if err != nil {
		return err
	}

	for i, e := range entries {
		if e.ID == entry.ID {
			entry.UpdatedAt = time.Now()
			entries[i] = *entry
			return fs.writeJSONFile("codex.json", entries)
		}
	}

	return fmt.Errorf("codex entry with ID %s not found", entry.ID)
}

func (fs *FileStorage) DeleteCodexEntry(id string) error {
	fs.mu.Lock()
	defer fs.mu.Unlock()

	entries, err := fs.GetCodexEntries()
	if err != nil {
		return err
	}

	for i, entry := range entries {
		if entry.ID == id {
			entries = append(entries[:i], entries[i+1:]...)
			return fs.writeJSONFile("codex.json", entries)
		}
	}

	return fmt.Errorf("codex entry with ID %s not found", id)
}

// Rules implementation
func (fs *FileStorage) GetRules() ([]models.Rule, error) {
	fs.mu.RLock()
	defer fs.mu.RUnlock()

	data, err := fs.readJSONFile("rules.json")
	if err != nil {
		return nil, err
	}

	var rules []models.Rule
	if err := json.Unmarshal(data, &rules); err != nil {
		return nil, fmt.Errorf("error parsing rules: %v", err)
	}

	return rules, nil
}

func (fs *FileStorage) GetRuleByID(id string) (*models.Rule, error) {
	rules, err := fs.GetRules()
	if err != nil {
		return nil, err
	}

	for _, rule := range rules {
		if rule.ID == id {
			return &rule, nil
		}
	}

	return nil, fmt.Errorf("rule with ID %s not found", id)
}

func (fs *FileStorage) GetActiveRules() ([]models.Rule, error) {
	rules, err := fs.GetRules()
	if err != nil {
		return nil, err
	}

	var active []models.Rule
	for _, rule := range rules {
		if rule.Active {
			active = append(active, rule)
		}
	}

	return active, nil
}

// Chapters implementation
func (fs *FileStorage) GetChapters() ([]models.Chapter, error) {
	fs.mu.RLock()
	defer fs.mu.RUnlock()

	data, err := fs.readJSONFile("chapters.json")
	if err != nil {
		return nil, err
	}

	var chapters []models.Chapter
	if err := json.Unmarshal(data, &chapters); err != nil {
		return nil, fmt.Errorf("error parsing chapters: %v", err)
	}

	// Sort by chapter number
	sort.Slice(chapters, func(i, j int) bool {
		return chapters[i].Number < chapters[j].Number
	})

	return chapters, nil
}

func (fs *FileStorage) GetChapterByID(id string) (*models.Chapter, error) {
	chapters, err := fs.GetChapters()
	if err != nil {
		return nil, err
	}

	for _, chapter := range chapters {
		if chapter.ID == id {
			return &chapter, nil
		}
	}

	return nil, fmt.Errorf("chapter with ID %s not found", id)
}

func (fs *FileStorage) GetChapterByNumber(number int) (*models.Chapter, error) {
	chapters, err := fs.GetChapters()
	if err != nil {
		return nil, err
	}

	for _, chapter := range chapters {
		if chapter.Number == number {
			return &chapter, nil
		}
	}

	return nil, fmt.Errorf("chapter %d not found", number)
}

func (fs *FileStorage) GetPreviousChapter(currentNumber int) (*models.Chapter, error) {
	chapters, err := fs.GetChapters()
	if err != nil {
		return nil, err
	}

	var previousChapter *models.Chapter
	for i := range chapters {
		if chapters[i].Number < currentNumber {
			if previousChapter == nil || chapters[i].Number > previousChapter.Number {
				previousChapter = &chapters[i]
			}
		}
	}

	if previousChapter == nil {
		return nil, fmt.Errorf("no previous chapter found")
	}

	return previousChapter, nil
}

func (fs *FileStorage) GetChapterRange(start, end int) ([]models.Chapter, error) {
	chapters, err := fs.GetChapters()
	if err != nil {
		return nil, err
	}

	var result []models.Chapter
	for _, chapter := range chapters {
		if chapter.Number >= start && chapter.Number <= end {
			result = append(result, chapter)
		}
	}

	return result, nil
}

func (fs *FileStorage) CreateChapter(chapter *models.Chapter) error {
	fs.mu.Lock()
	defer fs.mu.Unlock()

	chapters, err := fs.GetChapters()
	if err != nil && !os.IsNotExist(err) {
		return err
	}

	if chapter.ID == "" {
		chapter.ID = uuid.New().String()
	}

	// Auto-assign chapter number if not provided
	if chapter.Number == 0 {
		maxNumber := 0
		for _, ch := range chapters {
			if ch.Number > maxNumber {
				maxNumber = ch.Number
			}
		}
		chapter.Number = maxNumber + 1
	}

	// Calculate word count
	chapter.WordCount = len(strings.Fields(chapter.Content))
	chapter.CreatedAt = time.Now()
	chapter.UpdatedAt = time.Now()

	chapters = append(chapters, *chapter)

	return fs.writeJSONFile("chapters.json", chapters)
}

func (fs *FileStorage) UpdateChapter(chapter *models.Chapter) error {
	fs.mu.Lock()
	defer fs.mu.Unlock()

	chapters, err := fs.GetChapters()
	if err != nil {
		return err
	}

	for i, ch := range chapters {
		if ch.ID == chapter.ID {
			chapter.UpdatedAt = time.Now()
			chapter.WordCount = len(strings.Fields(chapter.Content))
			chapters[i] = *chapter
			return fs.writeJSONFile("chapters.json", chapters)
		}
	}

	return fmt.Errorf("chapter with ID %s not found", chapter.ID)
}

func (fs *FileStorage) DeleteChapter(id string) error {
	fs.mu.Lock()
	defer fs.mu.Unlock()

	chapters, err := fs.GetChapters()
	if err != nil {
		return err
	}

	for i, chapter := range chapters {
		if chapter.ID == id {
			chapters = append(chapters[:i], chapters[i+1:]...)
			return fs.writeJSONFile("chapters.json", chapters)
		}
	}

	return fmt.Errorf("chapter with ID %s not found", id)
}

func (fs *FileStorage) SearchChapters(query string) ([]models.Chapter, error) {
	chapters, err := fs.GetChapters()
	if err != nil {
		return nil, err
	}

	query = strings.ToLower(query)
	var results []models.Chapter

	for _, chapter := range chapters {
		if strings.Contains(strings.ToLower(chapter.Title), query) ||
			strings.Contains(strings.ToLower(chapter.Content), query) ||
			strings.Contains(strings.ToLower(chapter.Summary), query) {
			results = append(results, chapter)
		}
	}

	return results, nil
}

// Story Beats implementation
func (fs *FileStorage) GetStoryBeats(chapterNumber int) (*models.StoryBeats, error) {
	fs.mu.RLock()
	defer fs.mu.RUnlock()

	data, err := fs.readJSONFile("story_beats.json")
	if err != nil {
		return nil, err
	}

	var allBeats []models.StoryBeats
	if err := json.Unmarshal(data, &allBeats); err != nil {
		return nil, fmt.Errorf("error parsing story beats: %v", err)
	}

	for _, beats := range allBeats {
		if beats.ChapterNumber == chapterNumber {
			return &beats, nil
		}
	}

	return nil, fmt.Errorf("story beats for chapter %d not found", chapterNumber)
}

func (fs *FileStorage) GetAllStoryBeats() ([]models.StoryBeats, error) {
	fs.mu.RLock()
	defer fs.mu.RUnlock()

	data, err := fs.readJSONFile("story_beats.json")
	if err != nil {
		return nil, err
	}

	var allBeats []models.StoryBeats
	if err := json.Unmarshal(data, &allBeats); err != nil {
		return nil, fmt.Errorf("error parsing story beats: %v", err)
	}

	return allBeats, nil
}

func (fs *FileStorage) SaveStoryBeats(beats *models.StoryBeats) error {
	fs.mu.Lock()
	defer fs.mu.Unlock()

	allBeats, err := fs.getAllStoryBeats()
	if err != nil && !os.IsNotExist(err) {
		return err
	}

	beats.UpdatedAt = time.Now()

	// Update or append
	found := false
	for i, existing := range allBeats {
		if existing.ChapterNumber == beats.ChapterNumber {
			allBeats[i] = *beats
			found = true
			break
		}
	}

	if !found {
		allBeats = append(allBeats, *beats)
	}

	return fs.writeJSONFile("story_beats.json", allBeats)
}

// Future Notes implementation
func (fs *FileStorage) GetFutureNotes() ([]models.FutureNotes, error) {
	fs.mu.RLock()
	defer fs.mu.RUnlock()

	data, err := fs.readJSONFile("future_notes.json")
	if err != nil {
		return nil, err
	}

	var notes []models.FutureNotes
	if err := json.Unmarshal(data, &notes); err != nil {
		return nil, fmt.Errorf("error parsing future notes: %v", err)
	}

	return notes, nil
}

func (fs *FileStorage) GetFutureNotesByChapter(chapterNum int) ([]models.FutureNotes, error) {
	notes, err := fs.GetFutureNotes()
	if err != nil {
		return nil, err
	}

	var relevant []models.FutureNotes
	chapterStr := strconv.Itoa(chapterNum)

	for _, note := range notes {
		// Parse chapter range (e.g., "10-15")
		if strings.Contains(note.ChapterRange, "-") {
			parts := strings.Split(note.ChapterRange, "-")
			if len(parts) == 2 {
				start, _ := strconv.Atoi(strings.TrimSpace(parts[0]))
				end, _ := strconv.Atoi(strings.TrimSpace(parts[1]))
				if chapterNum >= start && chapterNum <= end {
					relevant = append(relevant, note)
				}
			}
		} else if note.ChapterRange == chapterStr {
			relevant = append(relevant, note)
		}
	}

	return relevant, nil
}

func (fs *FileStorage) CreateFutureNote(note *models.FutureNotes) error {
	fs.mu.Lock()
	defer fs.mu.Unlock()

	notes, err := fs.GetFutureNotes()
	if err != nil && !os.IsNotExist(err) {
		return err
	}

	if note.ID == "" {
		note.ID = uuid.New().String()
	}
	note.CreatedAt = time.Now()
	note.UpdatedAt = time.Now()

	notes = append(notes, *note)

	return fs.writeJSONFile("future_notes.json", notes)
}

func (fs *FileStorage) UpdateFutureNote(note *models.FutureNotes) error {
	fs.mu.Lock()
	defer fs.mu.Unlock()

	notes, err := fs.GetFutureNotes()
	if err != nil {
		return err
	}

	for i, n := range notes {
		if n.ID == note.ID {
			note.UpdatedAt = time.Now()
			notes[i] = *note
			return fs.writeJSONFile("future_notes.json", notes)
		}
	}

	return fmt.Errorf("future note with ID %s not found", note.ID)
}

func (fs *FileStorage) DeleteFutureNote(id string) error {
	fs.mu.Lock()
	defer fs.mu.Unlock()

	notes, err := fs.GetFutureNotes()
	if err != nil {
		return err
	}

	for i, note := range notes {
		if note.ID == id {
			notes = append(notes[:i], notes[i+1:]...)
			return fs.writeJSONFile("future_notes.json", notes)
		}
	}

	return fmt.Errorf("future note with ID %s not found", id)
}

// Sample Chapters implementation
func (fs *FileStorage) GetSampleChapters() ([]models.SampleChapter, error) {
	fs.mu.RLock()
	defer fs.mu.RUnlock()

	data, err := fs.readJSONFile("sample_chapters.json")
	if err != nil {
		return nil, err
	}

	var samples []models.SampleChapter
	if err := json.Unmarshal(data, &samples); err != nil {
		return nil, fmt.Errorf("error parsing sample chapters: %v", err)
	}

	return samples, nil
}

func (fs *FileStorage) GetSampleChapterByID(id string) (*models.SampleChapter, error) {
	samples, err := fs.GetSampleChapters()
	if err != nil {
		return nil, err
	}

	for _, sample := range samples {
		if sample.ID == id {
			return &sample, nil
		}
	}

	return nil, fmt.Errorf("sample chapter with ID %s not found", id)
}

func (fs *FileStorage) GetSampleChaptersByPurpose(purpose string) ([]models.SampleChapter, error) {
	samples, err := fs.GetSampleChapters()
	if err != nil {
		return nil, err
	}

	var filtered []models.SampleChapter
	for _, sample := range samples {
		if sample.Purpose == purpose {
			filtered = append(filtered, sample)
		}
	}

	return filtered, nil
}

func (fs *FileStorage) CreateSampleChapter(sample *models.SampleChapter) error {
	fs.mu.Lock()
	defer fs.mu.Unlock()

	samples, err := fs.GetSampleChapters()
	if err != nil && !os.IsNotExist(err) {
		return err
	}

	if sample.ID == "" {
		sample.ID = uuid.New().String()
	}
	sample.CreatedAt = time.Now()

	samples = append(samples, *sample)

	return fs.writeJSONFile("sample_chapters.json", samples)
}

// Task Types implementation
func (fs *FileStorage) GetTaskTypes() ([]models.TaskType, error) {
	fs.mu.RLock()
	defer fs.mu.RUnlock()

	data, err := fs.readJSONFile("task_types.json")
	if err != nil {
		return nil, err
	}

	var taskTypes []models.TaskType
	if err := json.Unmarshal(data, &taskTypes); err != nil {
		return nil, fmt.Errorf("error parsing task types: %v", err)
	}

	return taskTypes, nil
}

func (fs *FileStorage) GetTaskTypeByID(id string) (*models.TaskType, error) {
	taskTypes, err := fs.GetTaskTypes()
	if err != nil {
		return nil, err
	}

	for _, tt := range taskTypes {
		if tt.ID == id {
			return &tt, nil
		}
	}

	return nil, fmt.Errorf("task type with ID %s not found", id)
}

func (fs *FileStorage) GetActiveTaskTypes() ([]models.TaskType, error) {
	taskTypes, err := fs.GetTaskTypes()
	if err != nil {
		return nil, err
	}

	var active []models.TaskType
	for _, tt := range taskTypes {
		if tt.Active {
			active = append(active, tt)
		}
	}

	return active, nil
}

// Helper methods
func (fs *FileStorage) readJSONFile(filename string) ([]byte, error) {
	path := filepath.Join(fs.basePath, filename)
	data, err := os.ReadFile(path)
	if err != nil {
		if os.IsNotExist(err) {
			return []byte("[]"), nil
		}
		return nil, err
	}
	return data, nil
}

func (fs *FileStorage) writeJSONFile(filename string, data interface{}) error {
	path := filepath.Join(fs.basePath, filename)

	jsonData, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return fmt.Errorf("error marshaling data: %v", err)
	}

	return os.WriteFile(path, jsonData, 0644)
}

func (fs *FileStorage) getAllStoryBeats() ([]models.StoryBeats, error) {
	data, err := fs.readJSONFile("story_beats.json")
	if err != nil {
		if os.IsNotExist(err) {
			return []models.StoryBeats{}, nil
		}
		return nil, err
	}

	var beats []models.StoryBeats
	if err := json.Unmarshal(data, &beats); err != nil {
		return nil, err
	}

	return beats, nil
}

// Basic Storage interface implementation

// ReadFile reads a file from the storage directory
func (fs *FileStorage) ReadFile(filename string) ([]byte, error) {
	fs.mu.RLock()
	defer fs.mu.RUnlock()
	
	path := filepath.Join(fs.basePath, filename)
	return os.ReadFile(path)
}

// WriteFile writes data to a file in the storage directory
func (fs *FileStorage) WriteFile(filename string, data []byte) error {
	fs.mu.Lock()
	defer fs.mu.Unlock()
	
	path := filepath.Join(fs.basePath, filename)
	return os.WriteFile(path, data, 0644)
}

// DeleteFile deletes a file from the storage directory
func (fs *FileStorage) DeleteFile(filename string) error {
	fs.mu.Lock()
	defer fs.mu.Unlock()
	
	path := filepath.Join(fs.basePath, filename)
	return os.Remove(path)
}

// ListFiles lists files matching a pattern in the storage directory
func (fs *FileStorage) ListFiles(pattern string) ([]string, error) {
	fs.mu.RLock()
	defer fs.mu.RUnlock()
	
	fullPattern := filepath.Join(fs.basePath, pattern)
	matches, err := filepath.Glob(fullPattern)
	if err != nil {
		return nil, err
	}
	
	// Convert full paths back to relative filenames
	var filenames []string
	for _, match := range matches {
		filename, err := filepath.Rel(fs.basePath, match)
		if err != nil {
			continue
		}
		filenames = append(filenames, filename)
	}
	
	return filenames, nil
}
