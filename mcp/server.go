package mcp

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/danielsobrado/ainovelprompter/mcp/handlers"
	"github.com/danielsobrado/ainovelprompter/mcp/models"
	"github.com/danielsobrado/ainovelprompter/mcp/storage"
)

type MCPServer struct {
	storyHandler   *handlers.StoryContextHandler
	chapterHandler *handlers.ChapterHandler
	proseHandler   *handlers.ProseImprovementHandler
	searchHandler  *handlers.SearchHandler
}

func NewMCPServer() (*MCPServer, error) {
	return NewMCPServerWithDataDir("")
}

func NewMCPServerWithDataDir(dataDir string) (*MCPServer, error) {
	var appDataDir string
	var err error

	if dataDir != "" {
		// Use provided data directory
		appDataDir = dataDir
	} else {
		// Fall back to default user home directory
		homeDir, err := os.UserHomeDir()
		if err != nil {
			return nil, err
		}
		appDataDir = filepath.Join(homeDir, ".ai-novel-prompter")
	}

	// Create directory if it doesn't exist
	if err := os.MkdirAll(appDataDir, 0755); err != nil {
		return nil, fmt.Errorf("failed to create app data directory: %v", err)
	}

	// Initialize storage
	fileStorage := storage.NewFileStorage(appDataDir)

	// Initialize handlers
	storyHandler := handlers.NewStoryContextHandler(fileStorage)
	chapterHandler := handlers.NewChapterHandler(fileStorage)
	proseHandler := handlers.NewProseImprovementHandler(fileStorage)
	searchHandler := handlers.NewSearchHandler(fileStorage, storyHandler, chapterHandler, proseHandler)

	return &MCPServer{
		storyHandler:   storyHandler,
		chapterHandler: chapterHandler,
		proseHandler:   proseHandler,
		searchHandler:  searchHandler,
	}, nil
}

// Tool definitions for MCP
func (s *MCPServer) GetTools() []Tool {
	return []Tool{
		// Story Context Management Tools
		{
			Name:        "get_characters",
			Description: "Get story characters. Supports search parameter and filtering by IDs.",
			Parameters: map[string]ParameterDef{
				"search": {Type: "string", Description: "Optional search query", Required: false},
				"ids":    {Type: "array", Description: "Optional array of character IDs", Required: false},
			},
		},
		{
			Name:        "get_character_by_id",
			Description: "Get a specific character by ID",
			Parameters: map[string]ParameterDef{
				"id": {Type: "string", Description: "Character ID", Required: true},
			},
		},
		{
			Name:        "create_character",
			Description: "Create a new character",
			Parameters: map[string]ParameterDef{
				"name":        {Type: "string", Description: "Character name", Required: true},
				"description": {Type: "string", Description: "Character description", Required: false},
				"traits":      {Type: "object", Description: "Character traits as key-value pairs", Required: false},
				"notes":       {Type: "string", Description: "Additional notes", Required: false},
			},
		},
		{
			Name:        "update_character",
			Description: "Update an existing character",
			Parameters: map[string]ParameterDef{
				"id":          {Type: "string", Description: "Character ID", Required: true},
				"name":        {Type: "string", Description: "Character name", Required: false},
				"description": {Type: "string", Description: "Character description", Required: false},
				"traits":      {Type: "object", Description: "Character traits as key-value pairs", Required: false},
				"notes":       {Type: "string", Description: "Additional notes", Required: false},
			},
		},
		{
			Name:        "delete_character",
			Description: "Delete a character",
			Parameters: map[string]ParameterDef{
				"id": {Type: "string", Description: "Character ID", Required: true},
			},
		},
		{
			Name:        "get_locations",
			Description: "Get story locations. Supports search parameter.",
			Parameters: map[string]ParameterDef{
				"search": {Type: "string", Description: "Optional search query", Required: false},
			},
		},
		{
			Name:        "get_location_by_id",
			Description: "Get a specific location by ID",
			Parameters: map[string]ParameterDef{
				"id": {Type: "string", Description: "Location ID", Required: true},
			},
		},
		{
			Name:        "create_location",
			Description: "Create a new location",
			Parameters: map[string]ParameterDef{
				"name":        {Type: "string", Description: "Location name", Required: true},
				"description": {Type: "string", Description: "Location description", Required: false},
				"details":     {Type: "string", Description: "Additional details", Required: false},
				"notes":       {Type: "string", Description: "Additional notes", Required: false},
			},
		},
		{
			Name:        "get_codex_entries",
			Description: "Get codex/lore entries. Supports search and category filtering.",
			Parameters: map[string]ParameterDef{
				"search":   {Type: "string", Description: "Optional search query", Required: false},
				"category": {Type: "string", Description: "Filter by category", Required: false},
			},
		},
		{
			Name:        "get_codex_entry_by_id",
			Description: "Get a specific codex entry by ID",
			Parameters: map[string]ParameterDef{
				"id": {Type: "string", Description: "Codex entry ID", Required: true},
			},
		},
		{
			Name:        "create_codex_entry",
			Description: "Create a new codex entry",
			Parameters: map[string]ParameterDef{
				"title":    {Type: "string", Description: "Entry title", Required: true},
				"content":  {Type: "string", Description: "Entry content", Required: true},
				"category": {Type: "string", Description: "Entry category", Required: false},
				"tags":     {Type: "array", Description: "Tags for the entry", Required: false},
			},
		},
		{
			Name:        "get_rules",
			Description: "Get writing rules and guidelines",
			Parameters: map[string]ParameterDef{
				"activeOnly": {Type: "boolean", Description: "Only return active rules", Required: false},
				"category":   {Type: "string", Description: "Filter by category", Required: false},
			},
		},
		{
			Name:        "get_rule_by_id",
			Description: "Get a specific rule by ID",
			Parameters: map[string]ParameterDef{
				"id": {Type: "string", Description: "Rule ID", Required: true},
			},
		},
		{
			Name:        "build_writing_context",
			Description: "Build a comprehensive writing context with selected story elements",
			Parameters: map[string]ParameterDef{
				"characterIds":    {Type: "array", Description: "Character IDs to include", Required: false},
				"locationIds":     {Type: "array", Description: "Location IDs to include", Required: false},
				"includeRules":    {Type: "boolean", Description: "Include active writing rules", Required: false},
				"codexCategories": {Type: "array", Description: "Codex categories to include", Required: false},
			},
		},

		// Chapter Management Tools
		{
			Name:        "get_chapters",
			Description: "Get story chapters. Supports search, range, and status filtering.",
			Parameters: map[string]ParameterDef{
				"search": {Type: "string", Description: "Search in title, content, or summary", Required: false},
				"start":  {Type: "string", Description: "Start chapter number for range", Required: false},
				"end":    {Type: "string", Description: "End chapter number for range", Required: false},
				"status": {Type: "string", Description: "Filter by status (draft, review, final)", Required: false},
			},
		},
		{
			Name:        "get_chapter_content",
			Description: "Get a specific chapter's full content",
			Parameters: map[string]ParameterDef{
				"id":     {Type: "string", Description: "Chapter ID", Required: false},
				"number": {Type: "string", Description: "Chapter number", Required: false},
			},
		},
		{
			Name:        "get_previous_chapter",
			Description: "Get the previous chapter content",
			Parameters: map[string]ParameterDef{
				"currentChapter": {Type: "string", Description: "Current chapter number (optional)", Required: false},
			},
		},
		{
			Name:        "create_chapter",
			Description: "Create a new chapter",
			Parameters: map[string]ParameterDef{
				"title":        {Type: "string", Description: "Chapter title", Required: false},
				"content":      {Type: "string", Description: "Chapter content", Required: true},
				"number":       {Type: "string", Description: "Chapter number (auto-assigned if not provided)", Required: false},
				"summary":      {Type: "string", Description: "Chapter summary", Required: false},
				"status":       {Type: "string", Description: "Status (draft, review, final)", Required: false},
				"storyBeats":   {Type: "array", Description: "Story beats covered", Required: false},
				"characterIds": {Type: "array", Description: "Character IDs in chapter", Required: false},
				"locationIds":  {Type: "array", Description: "Location IDs in chapter", Required: false},
			},
		},
		{
			Name:        "update_chapter",
			Description: "Update an existing chapter",
			Parameters: map[string]ParameterDef{
				"id":      {Type: "string", Description: "Chapter ID", Required: true},
				"title":   {Type: "string", Description: "Chapter title", Required: false},
				"content": {Type: "string", Description: "Chapter content", Required: false},
				"summary": {Type: "string", Description: "Chapter summary", Required: false},
				"status":  {Type: "string", Description: "Status (draft, review, final)", Required: false},
			},
		},
		{
			Name:        "delete_chapter",
			Description: "Delete a chapter",
			Parameters: map[string]ParameterDef{
				"id": {Type: "string", Description: "Chapter ID", Required: true},
			},
		},
		{
			Name:        "get_story_beats",
			Description: "Get story beats for planning",
			Parameters: map[string]ParameterDef{
				"chapter": {Type: "string", Description: "Chapter number (returns all if not specified)", Required: false},
			},
		},
		{
			Name:        "save_story_beats",
			Description: "Save story beats for a chapter",
			Parameters: map[string]ParameterDef{
				"chapter": {Type: "string", Description: "Chapter number", Required: true},
				"beats":   {Type: "array", Description: "Array of beat objects", Required: true},
				"notes":   {Type: "string", Description: "Additional notes", Required: false},
			},
		},
		{
			Name:        "get_future_notes",
			Description: "Get planned future developments",
			Parameters: map[string]ParameterDef{
				"chapter":  {Type: "string", Description: "Filter by chapter number", Required: false},
				"priority": {Type: "string", Description: "Filter by priority (high, medium, low)", Required: false},
			},
		},
		{
			Name:        "create_future_note",
			Description: "Create a new future note",
			Parameters: map[string]ParameterDef{
				"chapterRange": {Type: "string", Description: "Chapter range (e.g., '10-15' or '12')", Required: true},
				"content":      {Type: "string", Description: "Note content", Required: true},
				"priority":     {Type: "string", Description: "Priority (high, medium, low)", Required: false},
				"category":     {Type: "string", Description: "Category (plot, character, world)", Required: false},
				"tags":         {Type: "array", Description: "Tags for the note", Required: false},
			},
		},
		{
			Name:        "get_sample_chapters",
			Description: "Get reference chapters for style consistency",
			Parameters: map[string]ParameterDef{
				"purpose": {Type: "string", Description: "Filter by purpose (style_reference, tone_example, pacing_guide)", Required: false},
			},
		},
		{
			Name:        "get_sample_chapter_by_id",
			Description: "Get a specific sample chapter by ID",
			Parameters: map[string]ParameterDef{
				"id": {Type: "string", Description: "Sample chapter ID", Required: true},
			},
		},
		{
			Name:        "create_sample_chapter",
			Description: "Create a new sample chapter",
			Parameters: map[string]ParameterDef{
				"title":   {Type: "string", Description: "Chapter title", Required: true},
				"content": {Type: "string", Description: "Chapter content", Required: true},
				"author":  {Type: "string", Description: "Author name", Required: false},
				"source":  {Type: "string", Description: "Source material", Required: false},
				"purpose": {Type: "string", Description: "Purpose (style_reference, tone_example, pacing_guide)", Required: false},
				"tags":    {Type: "array", Description: "Tags for the chapter", Required: false},
			},
		},
		{
			Name:        "get_task_types",
			Description: "Get writing task templates",
			Parameters: map[string]ParameterDef{
				"activeOnly": {Type: "boolean", Description: "Only return active task types", Required: false},
				"category":   {Type: "string", Description: "Filter by category", Required: false},
			},
		},
		{
			Name:        "get_task_type_by_id",
			Description: "Get a specific task type by ID",
			Parameters: map[string]ParameterDef{
				"id": {Type: "string", Description: "Task type ID", Required: true},
			},
		},
		{
			Name:        "apply_task_template",
			Description: "Generate prompt from task template",
			Parameters: map[string]ParameterDef{
				"taskId":    {Type: "string", Description: "Task type ID", Required: true},
				"variables": {Type: "object", Description: "Variables to fill in template", Required: false},
			},
		},

		// Prose Improvement Tools
		{
			Name:        "get_prose_prompts",
			Description: "Get prose improvement prompts. Supports category filtering.",
			Parameters: map[string]ParameterDef{
				"category":   {Type: "string", Description: "Filter by category (tropes, style, grammar, custom)", Required: false},
				"activeOnly": {Type: "boolean", Description: "Only return active prompts", Required: false},
			},
		},
		{
			Name:        "get_prose_prompt_by_id",
			Description: "Get a specific prose improvement prompt by ID",
			Parameters: map[string]ParameterDef{
				"id": {Type: "string", Description: "Prose prompt ID", Required: true},
			},
		},
		{
			Name:        "create_prose_prompt",
			Description: "Create a new prose improvement prompt",
			Parameters: map[string]ParameterDef{
				"label":             {Type: "string", Description: "Prompt label", Required: true},
				"defaultPromptText": {Type: "string", Description: "Default prompt text", Required: true},
				"category":          {Type: "string", Description: "Prompt category", Required: false},
				"description":       {Type: "string", Description: "Prompt description", Required: false},
			},
		},
		{
			Name:        "analyze_prose",
			Description: "Apply a specific prose improvement prompt to text",
			Parameters: map[string]ParameterDef{
				"text":     {Type: "string", Description: "Text to analyze", Required: true},
				"promptId": {Type: "string", Description: "Prose prompt ID to apply", Required: true},
				"provider": {Type: "string", Description: "LLM provider (manual, openrouter, lmstudio)", Required: false},
			},
		},
		{
			Name:        "get_prose_prompt_by_category",
			Description: "Get prose prompts filtered by category",
			Parameters: map[string]ParameterDef{
				"category": {Type: "string", Description: "Prompt category", Required: true},
			},
		},
		{
			Name:        "create_prose_session",
			Description: "Create a new prose improvement session",
			Parameters: map[string]ParameterDef{
				"text": {Type: "string", Description: "Text to improve", Required: true},
			},
		},
		{
			Name:        "get_prose_session",
			Description: "Get a prose improvement session by ID",
			Parameters: map[string]ParameterDef{
				"sessionId": {Type: "string", Description: "Session ID", Required: true},
			},
		},
		{
			Name:        "update_prose_session",
			Description: "Update a prose improvement session",
			Parameters: map[string]ParameterDef{
				"sessionId":           {Type: "string", Description: "Session ID", Required: true},
				"currentText":         {Type: "string", Description: "Updated current text", Required: false},
				"currentPromptIndex":  {Type: "number", Description: "Current prompt index", Required: false},
				"changes":             {Type: "array", Description: "Array of prose changes", Required: false},
			},
		},

		// Search & Analysis Tools
		{
			Name:        "search_all_content",
			Description: "Global search across all story elements",
			Parameters: map[string]ParameterDef{
				"query":        {Type: "string", Description: "Search query", Required: true},
				"contentTypes": {Type: "array", Description: "Content types to search (characters, locations, codex, chapters, rules, prose)", Required: false},
				"limit":        {Type: "number", Description: "Maximum number of results (default 50)", Required: false},
			},
		},
		{
			Name:        "analyze_text_traits",
			Description: "Extract style and tone information from text",
			Parameters: map[string]ParameterDef{
				"text": {Type: "string", Description: "Text to analyze", Required: true},
			},
		},
		{
			Name:        "get_character_mentions",
			Description: "Find character references in chapters",
			Parameters: map[string]ParameterDef{
				"characterId":   {Type: "string", Description: "Character ID", Required: false},
				"characterName": {Type: "string", Description: "Character name", Required: false},
				"chapterRange":  {Type: "string", Description: "Chapter range (e.g., '1-5' or '3')", Required: false},
			},
		},
		{
			Name:        "get_timeline_events",
			Description: "Retrieve story chronology and events",
			Parameters: map[string]ParameterDef{
				"startChapter": {Type: "string", Description: "Start chapter number", Required: false},
				"endChapter":   {Type: "string", Description: "End chapter number", Required: false},
				"eventType":    {Type: "string", Description: "Event type filter (action, planned, etc.)", Required: false},
			},
		},

		// Prompt Generation Tools
		{
			Name:        "generate_chapter_prompt",
			Description: "Generate AI prompt with story context for chapter writing",
			Parameters: map[string]ParameterDef{
				"promptType":       {Type: "string", Description: "Prompt format (ChatGPT, Claude)", Required: true},
				"characterIds":     {Type: "array", Description: "Character IDs to include", Required: false},
				"locationIds":      {Type: "array", Description: "Location IDs to include", Required: false},
				"ruleIds":          {Type: "array", Description: "Rule IDs to include", Required: false},
				"codexIds":         {Type: "array", Description: "Codex entry IDs to include", Required: false},
				"previousChapter":  {Type: "string", Description: "Previous chapter content", Required: false},
				"nextChapterBeats": {Type: "string", Description: "Story beats for next chapter", Required: false},
				"futureNotes":      {Type: "string", Description: "Future chapter notes", Required: false},
				"taskType":         {Type: "string", Description: "Task type/instruction", Required: false},
				"sampleChapter":    {Type: "string", Description: "Sample chapter for style reference", Required: false},
			},
		},
		{
			Name:        "get_prompt_template",
			Description: "Get a specific prompt template by format",
			Parameters: map[string]ParameterDef{
				"format": {Type: "string", Description: "Prompt format (ChatGPT, Claude)", Required: true},
			},
		},
	}
}

// Execute tool
func (s *MCPServer) ExecuteTool(toolName string, params map[string]interface{}) (interface{}, error) {
	switch toolName {
	// Story Context Management
	case "get_characters":
		return s.storyHandler.GetCharacters(params)
	case "get_character_by_id":
		return s.storyHandler.GetCharacterByID(params)
	case "create_character":
		return s.storyHandler.CreateCharacter(params)
	case "update_character":
		return s.storyHandler.UpdateCharacter(params)
	case "delete_character":
		return s.storyHandler.DeleteCharacter(params)
	case "get_locations":
		return s.storyHandler.GetLocations(params)
	case "get_location_by_id":
		return s.storyHandler.GetLocationByID(params)
	case "create_location":
		return s.storyHandler.CreateLocation(params)
	case "get_codex_entries":
		return s.storyHandler.GetCodexEntries(params)
	case "get_codex_entry_by_id":
		return s.storyHandler.GetCodexEntryByID(params)
	case "create_codex_entry":
		return s.storyHandler.CreateCodexEntry(params)
	case "get_rules":
		return s.storyHandler.GetRules(params)
	case "get_rule_by_id":
		return s.storyHandler.GetRuleByID(params)
	case "build_writing_context":
		return s.storyHandler.BuildWritingContext(params)

	// Chapter Management
	case "get_chapters":
		return s.chapterHandler.GetChapters(params)
	case "get_chapter_content":
		return s.chapterHandler.GetChapterContent(params)
	case "get_previous_chapter":
		return s.chapterHandler.GetPreviousChapter(params)
	case "create_chapter":
		return s.chapterHandler.CreateChapter(params)
	case "update_chapter":
		return s.chapterHandler.UpdateChapter(params)
	case "delete_chapter":
		return s.chapterHandler.DeleteChapter(params)
	case "get_story_beats":
		return s.chapterHandler.GetStoryBeats(params)
	case "save_story_beats":
		return s.chapterHandler.SaveStoryBeats(params)
	case "get_future_notes":
		return s.chapterHandler.GetFutureNotes(params)
	case "create_future_note":
		return s.chapterHandler.CreateFutureNote(params)
	case "get_sample_chapters":
		return s.chapterHandler.GetSampleChapters(params)
	case "get_sample_chapter_by_id":
		return s.chapterHandler.GetSampleChapterByID(params)
	case "create_sample_chapter":
		return s.chapterHandler.CreateSampleChapter(params)
	case "get_task_types":
		return s.chapterHandler.GetTaskTypes(params)
	case "get_task_type_by_id":
		return s.chapterHandler.GetTaskTypeByID(params)
	case "apply_task_template":
		return s.chapterHandler.ApplyTaskTemplate(params)

	// Prose Improvement Tools
	case "get_prose_prompts":
		return s.proseHandler.GetProsePrompts(params)
	case "get_prose_prompt_by_id":
		return s.proseHandler.GetProsePromptByID(params)
	case "create_prose_prompt":
		return s.proseHandler.CreateProsePrompt(params)
	case "analyze_prose":
		return s.proseHandler.AnalyzeProse(params)
	case "get_prose_prompt_by_category":
		return s.proseHandler.GetProsePromptByCategory(params)
	case "create_prose_session":
		return s.proseHandler.CreateProseSession(params)
	case "get_prose_session":
		return s.proseHandler.GetProseSession(params)
	case "update_prose_session":
		return s.proseHandler.UpdateProseSession(params)

	// Search & Analysis Tools
	case "search_all_content":
		return s.searchHandler.SearchAllContent(params)
	case "analyze_text_traits":
		return s.searchHandler.AnalyzeTextTraits(params)
	case "get_character_mentions":
		return s.searchHandler.GetCharacterMentions(params)
	case "get_timeline_events":
		return s.searchHandler.GetTimelineEvents(params)

	// Prompt Generation Tools
	case "generate_chapter_prompt":
		return s.generateChapterPrompt(params)
	case "get_prompt_template":
		return s.getPromptTemplate(params)

	default:
		return nil, fmt.Errorf("unknown tool: %s", toolName)
	}
}

// Prompt Generation Methods

// generateChapterPrompt generates a chapter writing prompt with story context
func (s *MCPServer) generateChapterPrompt(params map[string]interface{}) (interface{}, error) {
	promptType, ok := params["promptType"].(string)
	if !ok {
		promptType = "ChatGPT" // Default
	}

	var prompt string

	// Build the prompt based on type
	if promptType == "Claude" {
		prompt = s.buildClaudePrompt(params)
	} else {
		prompt = s.buildChatGPTPrompt(params)
	}

	// Calculate approximate token count (rough estimate)
	tokenCount := len(strings.Fields(prompt))

	return map[string]interface{}{
		"prompt":     prompt,
		"promptType": promptType,
		"tokenCount": tokenCount,
		"generatedAt": time.Now(),
	}, nil
}

// getPromptTemplate returns a basic prompt template for the specified format
func (s *MCPServer) getPromptTemplate(params map[string]interface{}) (interface{}, error) {
	format, ok := params["format"].(string)
	if !ok {
		return nil, fmt.Errorf("format parameter is required")
	}

	var template string
	switch format {
	case "ChatGPT":
		template = `You are a creative writer tasked with composing the next chapter based on the provided context and requirements. Please follow these guidelines:

1. Maintain consistency with previous chapters
2. Follow the provided story beats
3. Stay true to character voices and personalities
4. Keep the established tone and style
5. Incorporate relevant worldbuilding elements from the codex
6. Respect all provided rules and constraints

Your task is to write the next chapter that seamlessly continues the story while incorporating all the elements specified above.`

	case "Claude":
		template = `<instructions>
You are a creative writer tasked with composing the next chapter based on the provided context and requirements.

<guidelines>
1. Maintain consistency with the <previous_chapter>
2. Follow the provided <beats> precisely
3. Stay true to character voices defined in <characters>
4. Keep the established tone and style
5. Incorporate relevant worldbuilding from <codex>
6. Respect all <rules> strictly
</guidelines>

Your task is to write the next chapter that seamlessly continues the story while incorporating all the elements above.
</instructions>`

	default:
		return nil, fmt.Errorf("unsupported format: %s", format)
	}

	return map[string]interface{}{
		"template": template,
		"format":   format,
	}, nil
}

// buildChatGPTPrompt builds a ChatGPT-formatted prompt
func (s *MCPServer) buildChatGPTPrompt(params map[string]interface{}) string {
	var prompt strings.Builder

	// Add Previous Chapter
	if previousChapter, ok := params["previousChapter"].(string); ok && previousChapter != "" {
		prompt.WriteString("$1\n$2")
		prompt.WriteString(previousChapter)
		prompt.WriteString("$1\n$2")
	}

	// Add Sample Chapter
	if sampleChapter, ok := params["sampleChapter"].(string); ok && sampleChapter != "" {
		prompt.WriteString("$1\n$2")
		prompt.WriteString(sampleChapter)
		prompt.WriteString("$1\n$2")
	}

	// Add Future Notes
	if futureNotes, ok := params["futureNotes"].(string); ok && futureNotes != "" {
		prompt.WriteString("$1\n$2")
		prompt.WriteString(futureNotes)
		prompt.WriteString("$1\n$2")
	}

	// Add Story Beats
	if beats, ok := params["nextChapterBeats"].(string); ok && beats != "" {
		prompt.WriteString("$1\n$2")
		prompt.WriteString(beats)
		prompt.WriteString("$1\n$2")
	}

	// Add Rules
	if ruleIds, ok := params["ruleIds"].([]interface{}); ok && len(ruleIds) > 0 {
		s.addRulesToPrompt(&prompt, ruleIds)
	}

	// Add Characters
	if characterIds, ok := params["characterIds"].([]interface{}); ok && len(characterIds) > 0 {
		s.addCharactersToPrompt(&prompt, characterIds)
	}

	// Add Locations
	if locationIds, ok := params["locationIds"].([]interface{}); ok && len(locationIds) > 0 {
		s.addLocationsToPrompt(&prompt, locationIds)
	}

	// Add Codex Entries
	if codexIds, ok := params["codexIds"].([]interface{}); ok && len(codexIds) > 0 {
		s.addCodexToPrompt(&prompt, codexIds)
	}

	// Add Task Instructions
	if taskType, ok := params["taskType"].(string); ok && taskType != "" {
		prompt.WriteString("$1\n$2")
		prompt.WriteString(taskType)
		prompt.WriteString("$1\n$2")
	}

	return prompt.String()
}

// buildClaudePrompt builds a Claude-formatted prompt with XML tags
func (s *MCPServer) buildClaudePrompt(params map[string]interface{}) string {
	var prompt strings.Builder

	prompt.WriteString("$1\n$2")

	// Add Previous Chapter
	if previousChapter, ok := params["previousChapter"].(string); ok && previousChapter != "" {
		prompt.WriteString("$1\n$2")
		prompt.WriteString(previousChapter)
		prompt.WriteString("$1\n$2")
	}

	// Add Sample Chapter
	if sampleChapter, ok := params["sampleChapter"].(string); ok && sampleChapter != "" {
		prompt.WriteString("$1\n$2")
		prompt.WriteString(sampleChapter)
		prompt.WriteString("$1\n$2")
	}

	// Add Future Notes
	if futureNotes, ok := params["futureNotes"].(string); ok && futureNotes != "" {
		prompt.WriteString("$1\n$2")
		prompt.WriteString(futureNotes)
		prompt.WriteString("$1\n$2")
	}

	// Add Story Beats
	if beats, ok := params["nextChapterBeats"].(string); ok && beats != "" {
		prompt.WriteString("$1\n$2")
		prompt.WriteString(beats)
		prompt.WriteString("$1\n$2")
	}

	// Add Rules with XML tags
	if ruleIds, ok := params["ruleIds"].([]interface{}); ok && len(ruleIds) > 0 {
		s.addRulesToClaudePrompt(&prompt, ruleIds)
	}

	// Add Characters with XML tags
	if characterIds, ok := params["characterIds"].([]interface{}); ok && len(characterIds) > 0 {
		s.addCharactersToClaudePrompt(&prompt, characterIds)
	}

	// Add Locations with XML tags
	if locationIds, ok := params["locationIds"].([]interface{}); ok && len(locationIds) > 0 {
		s.addLocationsToClaudePrompt(&prompt, locationIds)
	}

	// Add Codex Entries with XML tags
	if codexIds, ok := params["codexIds"].([]interface{}); ok && len(codexIds) > 0 {
		s.addCodexToClaudePrompt(&prompt, codexIds)
	}

	// Add Task Instructions
	if taskType, ok := params["taskType"].(string); ok && taskType != "" {
		prompt.WriteString("$1\n$2")
		prompt.WriteString(taskType)
		prompt.WriteString("$1\n$2")
	}

	prompt.WriteString("</instructions>")

	return prompt.String()
}

// Helper methods for building prompts

func (s *MCPServer) addRulesToPrompt(prompt *strings.Builder, ruleIds []interface{}) {
	prompt.WriteString("$1\n$2")
	for _, ruleId := range ruleIds {
		if ruleIdStr, ok := ruleId.(string); ok {
			ruleResult, err := s.storyHandler.GetRuleByID(map[string]interface{}{"id": ruleIdStr})
			if err == nil {
				if rule, ok := ruleResult.(models.Rule); ok {
					prompt.WriteString(rule.Description)
					prompt.WriteString("$1\n$2")
				}
			}
		}
	}
	prompt.WriteString("$1\n$2")
}

func (s *MCPServer) addCharactersToPrompt(prompt *strings.Builder, characterIds []interface{}) {
	prompt.WriteString("$1\n$2")
	for _, characterId := range characterIds {
		if characterIdStr, ok := characterId.(string); ok {
			charResult, err := s.storyHandler.GetCharacterByID(map[string]interface{}{"id": characterIdStr})
			if err == nil {
				if character, ok := charResult.(models.Character); ok {
					prompt.WriteString(character.Name)
					prompt.WriteString(": ")
					prompt.WriteString(character.Description)
					prompt.WriteString("$1\n$2")
				}
			}
		}
	}
	prompt.WriteString("$1\n$2")
}

func (s *MCPServer) addLocationsToPrompt(prompt *strings.Builder, locationIds []interface{}) {
	prompt.WriteString("$1\n$2")
	for _, locationId := range locationIds {
		if locationIdStr, ok := locationId.(string); ok {
			locResult, err := s.storyHandler.GetLocationByID(map[string]interface{}{"id": locationIdStr})
			if err == nil {
				if location, ok := locResult.(models.Location); ok {
					prompt.WriteString(location.Name)
					prompt.WriteString(": ")
					prompt.WriteString(location.Description)
					prompt.WriteString("$1\n$2")
				}
			}
		}
	}
	prompt.WriteString("$1\n$2")
}

func (s *MCPServer) addCodexToPrompt(prompt *strings.Builder, codexIds []interface{}) {
	prompt.WriteString("$1\n$2")
	for _, codexId := range codexIds {
		if codexIdStr, ok := codexId.(string); ok {
			codexResult, err := s.storyHandler.GetCodexEntryByID(map[string]interface{}{"id": codexIdStr})
			if err == nil {
				if codex, ok := codexResult.(models.CodexEntry); ok {
					prompt.WriteString(codex.Title)
					prompt.WriteString(": ")
					prompt.WriteString(codex.Content)
					prompt.WriteString("$1\n$2")
				}
			}
		}
	}
	prompt.WriteString("$1\n$2")
}

func (s *MCPServer) addRulesToClaudePrompt(prompt *strings.Builder, ruleIds []interface{}) {
	prompt.WriteString("$1\n$2")
	for _, ruleId := range ruleIds {
		if ruleIdStr, ok := ruleId.(string); ok {
			ruleResult, err := s.storyHandler.GetRuleByID(map[string]interface{}{"id": ruleIdStr})
			if err == nil {
				if rule, ok := ruleResult.(models.Rule); ok {
					prompt.WriteString(rule.Description)
					prompt.WriteString("$1\n$2")
				}
			}
		}
	}
	prompt.WriteString("$1\n$2")
}

func (s *MCPServer) addCharactersToClaudePrompt(prompt *strings.Builder, characterIds []interface{}) {
	prompt.WriteString("$1\n$2")
	for _, characterId := range characterIds {
		if characterIdStr, ok := characterId.(string); ok {
			charResult, err := s.storyHandler.GetCharacterByID(map[string]interface{}{"id": characterIdStr})
			if err == nil {
				if character, ok := charResult.(models.Character); ok {
					prompt.WriteString(character.Name)
					prompt.WriteString(": ")
					prompt.WriteString(character.Description)
					prompt.WriteString("$1\n$2")
				}
			}
		}
	}
	prompt.WriteString("$1\n$2")
}

func (s *MCPServer) addLocationsToClaudePrompt(prompt *strings.Builder, locationIds []interface{}) {
	prompt.WriteString("$1\n$2")
	for _, locationId := range locationIds {
		if locationIdStr, ok := locationId.(string); ok {
			locResult, err := s.storyHandler.GetLocationByID(map[string]interface{}{"id": locationIdStr})
			if err == nil {
				if location, ok := locResult.(models.Location); ok {
					prompt.WriteString(location.Name)
					prompt.WriteString(": ")
					prompt.WriteString(location.Description)
					prompt.WriteString("$1\n$2")
				}
			}
		}
	}
	prompt.WriteString("$1\n$2")
}

func (s *MCPServer) addCodexToClaudePrompt(prompt *strings.Builder, codexIds []interface{}) {
	prompt.WriteString("$1\n$2")
	for _, codexId := range codexIds {
		if codexIdStr, ok := codexId.(string); ok {
			codexResult, err := s.storyHandler.GetCodexEntryByID(map[string]interface{}{"id": codexIdStr})
			if err == nil {
				if codex, ok := codexResult.(models.CodexEntry); ok {
					prompt.WriteString(codex.Title)
					prompt.WriteString(": ")
					prompt.WriteString(codex.Content)
					prompt.WriteString("$1\n$2")
				}
			}
		}
	}
	prompt.WriteString("$1\n$2")
}

// MCP protocol types
type Tool struct {
	Name        string                  `json:"name"`
	Description string                  `json:"description"`
	Parameters  map[string]ParameterDef `json:"parameters"`
}

type ParameterDef struct {
	Type        string `json:"type"`
	Description string `json:"description"`
	Required    bool   `json:"required"`
}
