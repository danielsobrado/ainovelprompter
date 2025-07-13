package storage

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/danielsobrado/ainovelprompter/mcp/models"
)

// Migration handles converting from old JSON file format to new folder format
type Migration struct {
	oldPath string
	newPath string
	storage *FolderStorage
}

// NewMigration creates a new migration instance
func NewMigration(oldPath, newPath string) *Migration {
	return &Migration{
		oldPath: oldPath,
		newPath: newPath,
		storage: NewFolderStorage(newPath),
	}
}

// MigrateFromJSON migrates all data from old JSON format to new folder format
func (fs *FolderStorage) MigrateFromJSON(oldPath string) error {
	migration := NewMigration(oldPath, fs.basePath)
	return migration.MigrateAll()
}

// MigrateAll migrates all entity types
func (m *Migration) MigrateAll() error {
	migrationSteps := []struct {
		name     string
		jsonFile string
		migrate  func(string) error
	}{
		{"characters", "characters.json", m.migrateCharacters},
		{"locations", "locations.json", m.migrateLocations},
		{"codex", "codex.json", m.migrateCodex},
		{"rules", "rules.json", m.migrateRules},
		{"chapters", "chapters.json", m.migrateChapters},
		{"story_beats", "story_beats.json", m.migrateStoryBeats},
		{"future_notes", "future_notes.json", m.migrateFutureNotes},
		{"sample_chapters", "sample_chapters.json", m.migrateSampleChapters},
		{"task_types", "task_types.json", m.migrateTaskTypes},
		{"prose_prompts", "prose_prompts.json", m.migrateProsePrompts},
		{"prose_prompt_definitions", "prose_prompt_definitions.json", m.migrateProsePromptDefinitions},
	}

	migratedCount := 0
	skippedCount := 0

	for _, step := range migrationSteps {
		filePath := filepath.Join(m.oldPath, step.jsonFile)
		
		// Check if file exists
		if _, err := os.Stat(filePath); os.IsNotExist(err) {
			fmt.Printf("Skipping %s (file not found)\n", step.name)
			skippedCount++
			continue
		}

		fmt.Printf("Migrating %s...\n", step.name)
		if err := step.migrate(filePath); err != nil {
			fmt.Printf("Error migrating %s: %v\n", step.name, err)
			return fmt.Errorf("migration failed for %s: %v", step.name, err)
		}
		migratedCount++
	}

	fmt.Printf("Migration completed: %d migrated, %d skipped\n", migratedCount, skippedCount)
	return nil
}

// migrateCharacters migrates character data
func (m *Migration) migrateCharacters(filePath string) error {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return err
	}

	var characters []models.Character
	if err := json.Unmarshal(data, &characters); err != nil {
		return fmt.Errorf("failed to parse characters: %v", err)
	}

	for _, character := range characters {
		// Ensure timestamps are set
		if character.CreatedAt.IsZero() {
			character.CreatedAt = time.Now()
		}
		if character.UpdatedAt.IsZero() {
			character.UpdatedAt = character.CreatedAt
		}

		if _, err := m.storage.Create(EntityCharacters, &character); err != nil {
			return fmt.Errorf("failed to create character %s: %v", character.Name, err)
		}
	}

	fmt.Printf("  Migrated %d characters\n", len(characters))
	return nil
}

// migrateLocations migrates location data
func (m *Migration) migrateLocations(filePath string) error {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return err
	}

	var locations []models.Location
	if err := json.Unmarshal(data, &locations); err != nil {
		return fmt.Errorf("failed to parse locations: %v", err)
	}

	for _, location := range locations {
		// Ensure timestamps are set
		if location.CreatedAt.IsZero() {
			location.CreatedAt = time.Now()
		}
		if location.UpdatedAt.IsZero() {
			location.UpdatedAt = location.CreatedAt
		}

		if _, err := m.storage.Create(EntityLocations, &location); err != nil {
			return fmt.Errorf("failed to create location %s: %v", location.Name, err)
		}
	}

	fmt.Printf("  Migrated %d locations\n", len(locations))
	return nil
}

// migrateCodex migrates codex entries
func (m *Migration) migrateCodex(filePath string) error {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return err
	}

	var entries []models.CodexEntry
	if err := json.Unmarshal(data, &entries); err != nil {
		return fmt.Errorf("failed to parse codex entries: %v", err)
	}

	for _, entry := range entries {
		// Ensure timestamps are set
		if entry.CreatedAt.IsZero() {
			entry.CreatedAt = time.Now()
		}
		if entry.UpdatedAt.IsZero() {
			entry.UpdatedAt = entry.CreatedAt
		}

		if _, err := m.storage.Create(EntityCodex, &entry); err != nil {
			return fmt.Errorf("failed to create codex entry %s: %v", entry.Title, err)
		}
	}

	fmt.Printf("  Migrated %d codex entries\n", len(entries))
	return nil
}

// migrateRules migrates writing rules
func (m *Migration) migrateRules(filePath string) error {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return err
	}

	var rules []models.Rule
	if err := json.Unmarshal(data, &rules); err != nil {
		return fmt.Errorf("failed to parse rules: %v", err)
	}

	for _, rule := range rules {
		// Ensure timestamps are set
		if rule.CreatedAt.IsZero() {
			rule.CreatedAt = time.Now()
		}

		if _, err := m.storage.Create(EntityRules, &rule); err != nil {
			return fmt.Errorf("failed to create rule %s: %v", rule.Name, err)
		}
	}

	fmt.Printf("  Migrated %d rules\n", len(rules))
	return nil
}

// migrateChapters migrates chapter data
func (m *Migration) migrateChapters(filePath string) error {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return err
	}

	var chapters []models.Chapter
	if err := json.Unmarshal(data, &chapters); err != nil {
		return fmt.Errorf("failed to parse chapters: %v", err)
	}

	for _, chapter := range chapters {
		// Ensure timestamps are set
		if chapter.CreatedAt.IsZero() {
			chapter.CreatedAt = time.Now()
		}
		if chapter.UpdatedAt.IsZero() {
			chapter.UpdatedAt = chapter.CreatedAt
		}

		if _, err := m.storage.Create(EntityChapters, &chapter); err != nil {
			return fmt.Errorf("failed to create chapter %s: %v", chapter.Title, err)
		}
	}

	fmt.Printf("  Migrated %d chapters\n", len(chapters))
	return nil
}

// migrateStoryBeats migrates story beats
func (m *Migration) migrateStoryBeats(filePath string) error {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return err
	}

	var storyBeats []models.StoryBeats
	if err := json.Unmarshal(data, &storyBeats); err != nil {
		return fmt.Errorf("failed to parse story beats: %v", err)
	}

	for _, beats := range storyBeats {
		// Ensure timestamps are set
		if beats.UpdatedAt.IsZero() {
			beats.UpdatedAt = time.Now()
		}

		if _, err := m.storage.Create(EntityStoryBeats, &beats); err != nil {
			return fmt.Errorf("failed to create story beats for chapter %d: %v", beats.ChapterNumber, err)
		}
	}

	fmt.Printf("  Migrated %d story beats entries\n", len(storyBeats))
	return nil
}

// migrateFutureNotes migrates future notes
func (m *Migration) migrateFutureNotes(filePath string) error {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return err
	}

	var notes []models.FutureNotes
	if err := json.Unmarshal(data, &notes); err != nil {
		return fmt.Errorf("failed to parse future notes: %v", err)
	}

	for _, note := range notes {
		// Ensure timestamps are set
		if note.CreatedAt.IsZero() {
			note.CreatedAt = time.Now()
		}
		if note.UpdatedAt.IsZero() {
			note.UpdatedAt = note.CreatedAt
		}

		if _, err := m.storage.Create(EntityFutureNotes, &note); err != nil {
			return fmt.Errorf("failed to create future note: %v", err)
		}
	}

	fmt.Printf("  Migrated %d future notes\n", len(notes))
	return nil
}

// migrateSampleChapters migrates sample chapters
func (m *Migration) migrateSampleChapters(filePath string) error {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return err
	}

	var samples []models.SampleChapter
	if err := json.Unmarshal(data, &samples); err != nil {
		return fmt.Errorf("failed to parse sample chapters: %v", err)
	}

	for _, sample := range samples {
		// Ensure timestamps are set
		if sample.CreatedAt.IsZero() {
			sample.CreatedAt = time.Now()
		}

		if _, err := m.storage.Create(EntitySampleChapters, &sample); err != nil {
			return fmt.Errorf("failed to create sample chapter %s: %v", sample.Title, err)
		}
	}

	fmt.Printf("  Migrated %d sample chapters\n", len(samples))
	return nil
}

// migrateTaskTypes migrates task types
func (m *Migration) migrateTaskTypes(filePath string) error {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return err
	}

	var taskTypes []models.TaskType
	if err := json.Unmarshal(data, &taskTypes); err != nil {
		return fmt.Errorf("failed to parse task types: %v", err)
	}

	for _, taskType := range taskTypes {
		if _, err := m.storage.Create(EntityTaskTypes, &taskType); err != nil {
			return fmt.Errorf("failed to create task type %s: %v", taskType.Name, err)
		}
	}

	fmt.Printf("  Migrated %d task types\n", len(taskTypes))
	return nil
}

// migrateProsePrompts migrates basic prose prompts (legacy format)
func (m *Migration) migrateProsePrompts(filePath string) error {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return err
	}

	// Try to parse as basic prompt structure first
	var basicPrompts []map[string]interface{}
	if err := json.Unmarshal(data, &basicPrompts); err != nil {
		return fmt.Errorf("failed to parse prose prompts: %v", err)
	}

	for _, prompt := range basicPrompts {
		// Convert to ProseImprovementPrompt structure
		improvedPrompt := models.ProseImprovementPrompt{
			ID:       fmt.Sprintf("migrated_%d", time.Now().UnixNano()),
			Category: "custom",
			Order:    0,
		}

		// Extract fields with fallbacks
		if label, ok := prompt["label"].(string); ok {
			improvedPrompt.Label = label
		}
		if desc, ok := prompt["description"].(string); ok {
			improvedPrompt.Description = desc
		}
		if defaultText, ok := prompt["defaultPromptText"].(string); ok {
			improvedPrompt.DefaultPromptText = defaultText
		}
		if category, ok := prompt["category"].(string); ok {
			improvedPrompt.Category = category
		}

		if _, err := m.storage.Create(EntityProsePrompts, &improvedPrompt); err != nil {
			return fmt.Errorf("failed to create prose prompt: %v", err)
		}
	}

	fmt.Printf("  Migrated %d prose prompts\n", len(basicPrompts))
	return nil
}

// migrateProsePromptDefinitions migrates prose prompt definitions
func (m *Migration) migrateProsePromptDefinitions(filePath string) error {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return err
	}

	var prompts []models.ProseImprovementPrompt
	if err := json.Unmarshal(data, &prompts); err != nil {
		return fmt.Errorf("failed to parse prose prompt definitions: %v", err)
	}

	for _, prompt := range prompts {
		if _, err := m.storage.Create(EntityProsePrompts, &prompt); err != nil {
			return fmt.Errorf("failed to create prose prompt definition %s: %v", prompt.Label, err)
		}
	}

	fmt.Printf("  Migrated %d prose prompt definitions\n", len(prompts))
	return nil
}

// CreateBackup creates a backup of the old data directory
func (m *Migration) CreateBackup() error {
	backupPath := m.oldPath + "_backup_" + time.Now().Format("20060102_150405")
	
	// Create backup directory
	if err := os.MkdirAll(backupPath, 0755); err != nil {
		return fmt.Errorf("failed to create backup directory: %v", err)
	}

	// Copy all JSON files
	jsonFiles := []string{
		"characters.json", "locations.json", "codex.json", "rules.json",
		"chapters.json", "story_beats.json", "future_notes.json",
		"sample_chapters.json", "task_types.json", "prose_prompts.json",
		"prose_prompt_definitions.json", "settings.json", "llm_provider_settings.json",
	}

	for _, file := range jsonFiles {
		srcPath := filepath.Join(m.oldPath, file)
		dstPath := filepath.Join(backupPath, file)

		if _, err := os.Stat(srcPath); os.IsNotExist(err) {
			continue // Skip files that don't exist
		}

		data, err := os.ReadFile(srcPath)
		if err != nil {
			continue // Skip files that can't be read
		}

		if err := os.WriteFile(dstPath, data, 0644); err != nil {
			return fmt.Errorf("failed to backup %s: %v", file, err)
		}
	}

	fmt.Printf("Backup created at: %s\n", backupPath)
	return nil
}

// ValidateMigration checks if migration was successful
func (m *Migration) ValidateMigration() error {
	stats, err := m.storage.GetStorageStats()
	if err != nil {
		return fmt.Errorf("failed to get storage stats: %v", err)
	}

	totalEntities := 0
	for entityType, count := range stats.EntitiesByType {
		fmt.Printf("  %s: %d entities\n", entityType, count)
		totalEntities += count
	}

	fmt.Printf("Total entities migrated: %d\n", totalEntities)
	fmt.Printf("Total files created: %d\n", stats.TotalFiles)

	return nil
}
