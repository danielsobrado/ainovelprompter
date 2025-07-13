package storage

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/danielsobrado/ainovelprompter/mcp/models"
)

// Legacy interface implementation for FolderStorage
// This maintains compatibility with existing MCP handlers

// StoryContextStorage implementation
func (fs *FolderStorage) GetCharacters() ([]models.Character, error) {
	entities, err := fs.GetAll(EntityCharacters)
	if err != nil {
		return nil, err
	}
	
	var characters []models.Character
	for _, entity := range entities {
		if char, ok := entity.(*models.Character); ok {
			characters = append(characters, *char)
		}
	}
	
	return characters, nil
}

func (fs *FolderStorage) GetCharacterByID(id string) (*models.Character, error) {
	entity, err := fs.GetLatest(EntityCharacters, id)
	if err != nil {
		return nil, err
	}
	
	if char, ok := entity.(*models.Character); ok {
		return char, nil
	}
	
	return nil, fmt.Errorf("entity is not a character")
}

func (fs *FolderStorage) CreateCharacter(character *models.Character) error {
	_, err := fs.Create(EntityCharacters, character)
	return err
}

func (fs *FolderStorage) UpdateCharacter(character *models.Character) error {
	_, err := fs.Update(EntityCharacters, character.ID, character)
	return err
}

func (fs *FolderStorage) DeleteCharacter(id string) error {
	_, err := fs.Delete(EntityCharacters, id)
	return err
}

func (fs *FolderStorage) SearchCharacters(query string) ([]models.Character, error) {
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

// Locations
func (fs *FolderStorage) GetLocations() ([]models.Location, error) {
	entities, err := fs.GetAll(EntityLocations)
	if err != nil {
		return nil, err
	}
	
	var locations []models.Location
	for _, entity := range entities {
		if loc, ok := entity.(*models.Location); ok {
			locations = append(locations, *loc)
		}
	}
	
	return locations, nil
}

func (fs *FolderStorage) GetLocationByID(id string) (*models.Location, error) {
	entity, err := fs.GetLatest(EntityLocations, id)
	if err != nil {
		return nil, err
	}
	
	if loc, ok := entity.(*models.Location); ok {
		return loc, nil
	}
	
	return nil, fmt.Errorf("entity is not a location")
}

func (fs *FolderStorage) CreateLocation(location *models.Location) error {
	_, err := fs.Create(EntityLocations, location)
	return err
}

func (fs *FolderStorage) UpdateLocation(location *models.Location) error {
	_, err := fs.Update(EntityLocations, location.ID, location)
	return err
}

func (fs *FolderStorage) DeleteLocation(id string) error {
	_, err := fs.Delete(EntityLocations, id)
	return err
}

func (fs *FolderStorage) SearchLocations(query string) ([]models.Location, error) {
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

// Codex
func (fs *FolderStorage) GetCodexEntries() ([]models.CodexEntry, error) {
	entities, err := fs.GetAll(EntityCodex)
	if err != nil {
		return nil, err
	}
	
	var entries []models.CodexEntry
	for _, entity := range entities {
		if entry, ok := entity.(*models.CodexEntry); ok {
			entries = append(entries, *entry)
		}
	}
	
	return entries, nil
}

func (fs *FolderStorage) GetCodexEntryByID(id string) (*models.CodexEntry, error) {
	entity, err := fs.GetLatest(EntityCodex, id)
	if err != nil {
		return nil, err
	}
	
	if entry, ok := entity.(*models.CodexEntry); ok {
		return entry, nil
	}
	
	return nil, fmt.Errorf("entity is not a codex entry")
}

func (fs *FolderStorage) CreateCodexEntry(entry *models.CodexEntry) error {
	_, err := fs.Create(EntityCodex, entry)
	return err
}

func (fs *FolderStorage) UpdateCodexEntry(entry *models.CodexEntry) error {
	_, err := fs.Update(EntityCodex, entry.ID, entry)
	return err
}

func (fs *FolderStorage) DeleteCodexEntry(id string) error {
	_, err := fs.Delete(EntityCodex, id)
	return err
}

func (fs *FolderStorage) SearchCodex(query string) ([]models.CodexEntry, error) {
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

// Rules
func (fs *FolderStorage) GetRules() ([]models.Rule, error) {
	entities, err := fs.GetAll(EntityRules)
	if err != nil {
		return nil, err
	}
	
	var rules []models.Rule
	for _, entity := range entities {
		if rule, ok := entity.(*models.Rule); ok {
			rules = append(rules, *rule)
		}
	}
	
	return rules, nil
}

func (fs *FolderStorage) GetRuleByID(id string) (*models.Rule, error) {
	entity, err := fs.GetLatest(EntityRules, id)
	if err != nil {
		return nil, err
	}
	
	if rule, ok := entity.(*models.Rule); ok {
		return rule, nil
	}
	
	return nil, fmt.Errorf("entity is not a rule")
}

func (fs *FolderStorage) GetActiveRules() ([]models.Rule, error) {
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

// ChapterStorage implementation
func (fs *FolderStorage) GetChapters() ([]models.Chapter, error) {
	entities, err := fs.GetAll(EntityChapters)
	if err != nil {
		return nil, err
	}
	
	var chapters []models.Chapter
	for _, entity := range entities {
		if chapter, ok := entity.(*models.Chapter); ok {
			chapters = append(chapters, *chapter)
		}
	}
	
	return chapters, nil
}

func (fs *FolderStorage) GetChapterByID(id string) (*models.Chapter, error) {
	entity, err := fs.GetLatest(EntityChapters, id)
	if err != nil {
		return nil, err
	}
	
	if chapter, ok := entity.(*models.Chapter); ok {
		return chapter, nil
	}
	
	return nil, fmt.Errorf("entity is not a chapter")
}

func (fs *FolderStorage) GetChapterByNumber(number int) (*models.Chapter, error) {
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

func (fs *FolderStorage) GetPreviousChapter(currentNumber int) (*models.Chapter, error) {
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

func (fs *FolderStorage) GetChapterRange(start, end int) ([]models.Chapter, error) {
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

func (fs *FolderStorage) CreateChapter(chapter *models.Chapter) error {
	_, err := fs.Create(EntityChapters, chapter)
	return err
}

func (fs *FolderStorage) UpdateChapter(chapter *models.Chapter) error {
	_, err := fs.Update(EntityChapters, chapter.ID, chapter)
	return err
}

func (fs *FolderStorage) DeleteChapter(id string) error {
	_, err := fs.Delete(EntityChapters, id)
	return err
}

func (fs *FolderStorage) SearchChapters(query string) ([]models.Chapter, error) {
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

// Story Beats
func (fs *FolderStorage) GetStoryBeats(chapterNumber int) (*models.StoryBeats, error) {
	entities, err := fs.GetAll(EntityStoryBeats)
	if err != nil {
		return nil, err
	}
	
	for _, entity := range entities {
		if beats, ok := entity.(*models.StoryBeats); ok {
			if beats.ChapterNumber == chapterNumber {
				return beats, nil
			}
		}
	}
	
	return nil, fmt.Errorf("story beats for chapter %d not found", chapterNumber)
}

func (fs *FolderStorage) GetAllStoryBeats() ([]models.StoryBeats, error) {
	entities, err := fs.GetAll(EntityStoryBeats)
	if err != nil {
		return nil, err
	}
	
	var allBeats []models.StoryBeats
	for _, entity := range entities {
		if beats, ok := entity.(*models.StoryBeats); ok {
			allBeats = append(allBeats, *beats)
		}
	}
	
	return allBeats, nil
}

func (fs *FolderStorage) SaveStoryBeats(beats *models.StoryBeats) error {
	// Check if beats already exist
	if existing, err := fs.GetStoryBeats(beats.ChapterNumber); err == nil {
		beats.ID = existing.ID
		_, err := fs.Update(EntityStoryBeats, existing.ID, beats)
		return err
	} else {
		_, err := fs.Create(EntityStoryBeats, beats)
		return err
	}
}

// Future Notes
func (fs *FolderStorage) GetFutureNotes() ([]models.FutureNotes, error) {
	entities, err := fs.GetAll(EntityFutureNotes)
	if err != nil {
		return nil, err
	}
	
	var notes []models.FutureNotes
	for _, entity := range entities {
		if note, ok := entity.(*models.FutureNotes); ok {
			notes = append(notes, *note)
		}
	}
	
	return notes, nil
}

func (fs *FolderStorage) GetFutureNotesByChapter(chapterNum int) ([]models.FutureNotes, error) {
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

func (fs *FolderStorage) CreateFutureNote(note *models.FutureNotes) error {
	_, err := fs.Create(EntityFutureNotes, note)
	return err
}

func (fs *FolderStorage) UpdateFutureNote(note *models.FutureNotes) error {
	_, err := fs.Update(EntityFutureNotes, note.ID, note)
	return err
}

func (fs *FolderStorage) DeleteFutureNote(id string) error {
	_, err := fs.Delete(EntityFutureNotes, id)
	return err
}

// Sample Chapters
func (fs *FolderStorage) GetSampleChapters() ([]models.SampleChapter, error) {
	entities, err := fs.GetAll(EntitySampleChapters)
	if err != nil {
		return nil, err
	}
	
	var samples []models.SampleChapter
	for _, entity := range entities {
		if sample, ok := entity.(*models.SampleChapter); ok {
			samples = append(samples, *sample)
		}
	}
	
	return samples, nil
}

func (fs *FolderStorage) GetSampleChapterByID(id string) (*models.SampleChapter, error) {
	entity, err := fs.GetLatest(EntitySampleChapters, id)
	if err != nil {
		return nil, err
	}
	
	if sample, ok := entity.(*models.SampleChapter); ok {
		return sample, nil
	}
	
	return nil, fmt.Errorf("entity is not a sample chapter")
}

func (fs *FolderStorage) GetSampleChaptersByPurpose(purpose string) ([]models.SampleChapter, error) {
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

func (fs *FolderStorage) CreateSampleChapter(sample *models.SampleChapter) error {
	_, err := fs.Create(EntitySampleChapters, sample)
	return err
}

// Task Types
func (fs *FolderStorage) GetTaskTypes() ([]models.TaskType, error) {
	entities, err := fs.GetAll(EntityTaskTypes)
	if err != nil {
		return nil, err
	}
	
	var taskTypes []models.TaskType
	for _, entity := range entities {
		if taskType, ok := entity.(*models.TaskType); ok {
			taskTypes = append(taskTypes, *taskType)
		}
	}
	
	return taskTypes, nil
}

func (fs *FolderStorage) GetTaskTypeByID(id string) (*models.TaskType, error) {
	entity, err := fs.GetLatest(EntityTaskTypes, id)
	if err != nil {
		return nil, err
	}
	
	if taskType, ok := entity.(*models.TaskType); ok {
		return taskType, nil
	}
	
	return nil, fmt.Errorf("entity is not a task type")
}

func (fs *FolderStorage) GetActiveTaskTypes() ([]models.TaskType, error) {
	taskTypes, err := fs.GetTaskTypes()
	if err != nil {
		return nil, err
	}
	
	var active []models.TaskType
	for _, taskType := range taskTypes {
		if taskType.Active {
			active = append(active, taskType)
		}
	}
	
	return active, nil
}

// Basic Storage interface implementation
func (fs *FolderStorage) ReadFile(filename string) ([]byte, error) {
	return fs.ReadFile(filename)
}

func (fs *FolderStorage) WriteFile(filename string, data []byte) error {
	return fs.WriteFile(filename, data)
}

func (fs *FolderStorage) DeleteFile(filename string) error {
	return fs.DeleteFile(filename)
}

func (fs *FolderStorage) ListFiles(pattern string) ([]string, error) {
	return fs.ListFiles(pattern)
}
