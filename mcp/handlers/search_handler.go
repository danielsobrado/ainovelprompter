package handlers

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/danielsobrado/ainovelprompter/mcp/models"
	"github.com/danielsobrado/ainovelprompter/mcp/storage"
)

// SearchHandler handles search and analysis operations across all content
type SearchHandler struct {
	storage        storage.Storage
	storyHandler   *StoryContextHandler
	chapterHandler *ChapterHandler
	proseHandler   *ProseImprovementHandler
}

// NewSearchHandler creates a new search handler
func NewSearchHandler(storage storage.Storage, storyHandler *StoryContextHandler, chapterHandler *ChapterHandler, proseHandler *ProseImprovementHandler) *SearchHandler {
	return &SearchHandler{
		storage:        storage,
		storyHandler:   storyHandler,
		chapterHandler: chapterHandler,
		proseHandler:   proseHandler,
	}
}

// SearchAllContent performs a global search across all story elements
func (h *SearchHandler) SearchAllContent(params map[string]interface{}) (interface{}, error) {
	query, ok := params["query"].(string)
	if !ok {
		return nil, fmt.Errorf("query parameter is required")
	}

	contentTypes, _ := params["contentTypes"].([]interface{})
	limit, _ := params["limit"].(float64)
	if limit == 0 {
		limit = 50 // Default limit
	}

	var allResults []models.SearchResult

	// Search characters
	if h.shouldSearchType(contentTypes, "characters") {
		if results, err := h.searchCharacters(query); err == nil {
			allResults = append(allResults, results...)
		}
	}

	// Search locations
	if h.shouldSearchType(contentTypes, "locations") {
		if results, err := h.searchLocations(query); err == nil {
			allResults = append(allResults, results...)
		}
	}

	// Search codex entries
	if h.shouldSearchType(contentTypes, "codex") {
		if results, err := h.searchCodex(query); err == nil {
			allResults = append(allResults, results...)
		}
	}

	// Search chapters
	if h.shouldSearchType(contentTypes, "chapters") {
		if results, err := h.searchChapters(query); err == nil {
			allResults = append(allResults, results...)
		}
	}

	// Search rules
	if h.shouldSearchType(contentTypes, "rules") {
		if results, err := h.searchRules(query); err == nil {
			allResults = append(allResults, results...)
		}
	}

	// Search prose content
	if h.shouldSearchType(contentTypes, "prose") {
		if results, err := h.searchProseContent(query); err == nil {
			allResults = append(allResults, results...)
		}
	}

	// Sort by score (highest first)
	h.sortResultsByScore(allResults)

	// Apply limit
	if int(limit) < len(allResults) {
		allResults = allResults[:int(limit)]
	}

	return allResults, nil
}

// AnalyzeTextTraits extracts style and tone information from text
func (h *SearchHandler) AnalyzeTextTraits(params map[string]interface{}) (interface{}, error) {
	text, ok := params["text"].(string)
	if !ok {
		return nil, fmt.Errorf("text parameter is required")
	}

	// Simple text analysis - in a real implementation, this could use NLP libraries
	analysis := map[string]interface{}{
		"wordCount":                  len(strings.Fields(text)),
		"characterCount":             len(text),
		"sentenceCount":              strings.Count(text, ".") + strings.Count(text, "!") + strings.Count(text, "?"),
		"paragraphCount":             len(strings.Split(text, "\n\n")),
		"averageWordsPerSentence":    h.calculateAverageWordsPerSentence(text),
		"readabilityLevel":           h.estimateReadabilityLevel(text),
		"tone":                       h.analyzeTone(text),
		"commonWords":                h.getCommonWords(text, 10),
		"writingStyle":               h.analyzeWritingStyle(text),
		"analyzedAt":                 time.Now(),
	}

	return analysis, nil
}

// GetCharacterMentions finds character references in chapters
func (h *SearchHandler) GetCharacterMentions(params map[string]interface{}) (interface{}, error) {
	characterID, _ := params["characterId"].(string)
	characterName, _ := params["characterName"].(string)
	chapterRange, _ := params["chapterRange"].(string)

	if characterID == "" && characterName == "" {
		return nil, fmt.Errorf("either characterId or characterName parameter is required")
	}

	// Get character name if ID provided
	if characterID != "" && characterName == "" {
		charResult, err := h.storyHandler.GetCharacterByID(map[string]interface{}{"id": characterID})
		if err != nil {
			return nil, fmt.Errorf("character not found: %s", characterID)
		}
		if char, ok := charResult.(models.Character); ok {
			characterName = char.Name
		}
	}

	// Get all chapters
	chaptersResult, err := h.chapterHandler.GetChapters(map[string]interface{}{})
	if err != nil {
		return nil, err
	}

	chapters, ok := chaptersResult.([]models.Chapter)
	if !ok {
		return nil, fmt.Errorf("failed to get chapters")
	}

	var mentions []map[string]interface{}

	for _, chapter := range chapters {
		// Skip if chapter range specified and this chapter is outside range
		if chapterRange != "" && !h.isChapterInRange(chapter.Number, chapterRange) {
			continue
		}

		// Search for character mentions in content
		content := strings.ToLower(chapter.Content)
		name := strings.ToLower(characterName)
		
		if strings.Contains(content, name) {
			// Find all occurrences
			mentionCount := strings.Count(content, name)
			
			// Create context snippets
			snippets := h.findMentionSnippets(chapter.Content, characterName, 3)
			
			mentions = append(mentions, map[string]interface{}{
				"chapterId":     chapter.ID,
				"chapterNumber": chapter.Number,
				"chapterTitle":  chapter.Title,
				"mentionCount":  mentionCount,
				"snippets":      snippets,
				"lastMention":   time.Now(), // In real implementation, would track actual timestamps
			})
		}
	}

	return map[string]interface{}{
		"characterName": characterName,
		"totalMentions": len(mentions),
		"chapters":      mentions,
		"analyzedAt":    time.Now(),
	}, nil
}

// GetTimelineEvents retrieves story chronology and events
func (h *SearchHandler) GetTimelineEvents(params map[string]interface{}) (interface{}, error) {
	startChapter, _ := params["startChapter"].(string)
	endChapter, _ := params["endChapter"].(string)
	eventType, _ := params["eventType"].(string)

	// Get all chapters
	chaptersResult, err := h.chapterHandler.GetChapters(map[string]interface{}{
		"start": startChapter,
		"end":   endChapter,
	})
	if err != nil {
		return nil, err
	}

	chapters, ok := chaptersResult.([]models.Chapter)
	if !ok {
		return nil, fmt.Errorf("failed to get chapters")
	}

	// Get story beats for timeline context
	beatsResult, err := h.chapterHandler.GetStoryBeats(map[string]interface{}{})
	if err != nil {
		return nil, err
	}

	var events []map[string]interface{}

	// Extract events from chapters
	for _, chapter := range chapters {
		chapterEvents := h.extractEventsFromChapter(chapter, eventType)
		events = append(events, chapterEvents...)
	}

	// Add story beats as planned events
	if beats, ok := beatsResult.([]models.StoryBeat); ok {
		for _, beat := range beats {
			if eventType == "" || eventType == "planned" {
				events = append(events, map[string]interface{}{
					"id":          beat.ID,
					"type":        "planned",
					"title":       beat.Description,
					"chapter":     beat.ChapterNumber,
					"description": beat.Notes,
					"status":      "planned",
					"createdAt":   beat.CreatedAt,
				})
			}
		}
	}

	// Sort events by chapter order
	h.sortEventsByChapter(events)

	return map[string]interface{}{
		"events":           events,
		"totalEvents":      len(events),
		"timelineGenerated": time.Now(),
	}, nil
}

// Helper methods for search operations

func (h *SearchHandler) shouldSearchType(contentTypes []interface{}, searchType string) bool {
	if len(contentTypes) == 0 {
		return true // Search all types if none specified
	}
	
	for _, ct := range contentTypes {
		if ct.(string) == searchType {
			return true
		}
	}
	return false
}

func (h *SearchHandler) searchCharacters(query string) ([]models.SearchResult, error) {
	charactersResult, err := h.storyHandler.GetCharacters(map[string]interface{}{"search": query})
	if err != nil {
		return nil, err
	}

	characters, ok := charactersResult.([]models.Character)
	if !ok {
		return nil, fmt.Errorf("failed to get characters")
	}

	var results []models.SearchResult
	for _, char := range characters {
		if h.matchesQuery(strings.ToLower(query), strings.ToLower(char.Name), strings.ToLower(char.Description), strings.ToLower(char.Notes)) {
			results = append(results, models.SearchResult{
				ID:      char.ID,
				Type:    "character",
				Title:   char.Name,
				Content: char.Description,
				Snippet: h.createSnippet(char.Description, query),
				Score:   h.calculateScore(query, char.Name, char.Description),
				Metadata: map[string]interface{}{
					"traits": char.Traits,
					"notes":  char.Notes,
				},
				MatchedAt: time.Now(),
			})
		}
	}
	return results, nil
}

func (h *SearchHandler) searchLocations(query string) ([]models.SearchResult, error) {
	locationsResult, err := h.storyHandler.GetLocations(map[string]interface{}{"search": query})
	if err != nil {
		return nil, err
	}

	locations, ok := locationsResult.([]models.Location)
	if !ok {
		return nil, fmt.Errorf("failed to get locations")
	}

	var results []models.SearchResult
	for _, loc := range locations {
		if h.matchesQuery(strings.ToLower(query), strings.ToLower(loc.Name), strings.ToLower(loc.Description), strings.ToLower(loc.Details), strings.ToLower(loc.Notes)) {
			results = append(results, models.SearchResult{
				ID:      loc.ID,
				Type:    "location",
				Title:   loc.Name,
				Content: loc.Description,
				Snippet: h.createSnippet(loc.Description, query),
				Score:   h.calculateScore(query, loc.Name, loc.Description),
				Metadata: map[string]interface{}{
					"details": loc.Details,
					"notes":   loc.Notes,
				},
				MatchedAt: time.Now(),
			})
		}
	}
	return results, nil
}

func (h *SearchHandler) searchCodex(query string) ([]models.SearchResult, error) {
	codexResult, err := h.storyHandler.GetCodexEntries(map[string]interface{}{"search": query})
	if err != nil {
		return nil, err
	}

	entries, ok := codexResult.([]models.CodexEntry)
	if !ok {
		return nil, fmt.Errorf("failed to get codex entries")
	}

	var results []models.SearchResult
	for _, entry := range entries {
		if h.matchesQuery(strings.ToLower(query), strings.ToLower(entry.Title), strings.ToLower(entry.Content)) {
			results = append(results, models.SearchResult{
				ID:      entry.ID,
				Type:    "codex",
				Title:   entry.Title,
				Content: entry.Content,
				Snippet: h.createSnippet(entry.Content, query),
				Score:   h.calculateScore(query, entry.Title, entry.Content),
				Metadata: map[string]interface{}{
					"category": entry.Category,
					"tags":     entry.Tags,
				},
				MatchedAt: time.Now(),
			})
		}
	}
	return results, nil
}

func (h *SearchHandler) searchChapters(query string) ([]models.SearchResult, error) {
	chaptersResult, err := h.chapterHandler.GetChapters(map[string]interface{}{"search": query})
	if err != nil {
		return nil, err
	}

	chapters, ok := chaptersResult.([]models.Chapter)
	if !ok {
		return nil, fmt.Errorf("failed to get chapters")
	}

	var results []models.SearchResult
	for _, chapter := range chapters {
		if h.matchesQuery(strings.ToLower(query), strings.ToLower(chapter.Title), strings.ToLower(chapter.Content), strings.ToLower(chapter.Summary)) {
			results = append(results, models.SearchResult{
				ID:      chapter.ID,
				Type:    "chapter",
				Title:   fmt.Sprintf("Chapter %s: %s", chapter.Number, chapter.Title),
				Content: chapter.Content,
				Snippet: h.createSnippet(chapter.Content, query),
				Score:   h.calculateScore(query, chapter.Title, chapter.Content),
				Metadata: map[string]interface{}{
					"number":    chapter.Number,
					"summary":   chapter.Summary,
					"status":    chapter.Status,
					"wordCount": chapter.WordCount,
				},
				MatchedAt: time.Now(),
			})
		}
	}
	return results, nil
}

func (h *SearchHandler) searchRules(query string) ([]models.SearchResult, error) {
	rulesResult, err := h.storyHandler.GetRules(map[string]interface{}{})
	if err != nil {
		return nil, err
	}

	rules, ok := rulesResult.([]models.Rule)
	if !ok {
		return nil, fmt.Errorf("failed to get rules")
	}

	var results []models.SearchResult
	for _, rule := range rules {
		if h.matchesQuery(strings.ToLower(query), strings.ToLower(rule.Title), strings.ToLower(rule.Description)) {
			results = append(results, models.SearchResult{
				ID:      rule.ID,
				Type:    "rule",
				Title:   rule.Title,
				Content: rule.Description,
				Snippet: h.createSnippet(rule.Description, query),
				Score:   h.calculateScore(query, rule.Title, rule.Description),
				Metadata: map[string]interface{}{
					"category": rule.Category,
					"active":   rule.Active,
				},
				MatchedAt: time.Now(),
			})
		}
	}
	return results, nil
}

func (h *SearchHandler) searchProseContent(query string) ([]models.SearchResult, error) {
	if h.proseHandler != nil {
		proseResult, err := h.proseHandler.SearchProseContent(map[string]interface{}{"query": query})
		if err != nil {
			return nil, err
		}

		if results, ok := proseResult.([]models.SearchResult); ok {
			return results, nil
		}
	}
	return []models.SearchResult{}, nil
}

// Text analysis helper methods

func (h *SearchHandler) calculateAverageWordsPerSentence(text string) float64 {
	sentences := strings.Count(text, ".") + strings.Count(text, "!") + strings.Count(text, "?")
	if sentences == 0 {
		return 0
	}
	words := len(strings.Fields(text))
	return float64(words) / float64(sentences)
}

func (h *SearchHandler) estimateReadabilityLevel(text string) string {
	avgWordsPerSentence := h.calculateAverageWordsPerSentence(text)
	
	if avgWordsPerSentence < 10 {
		return "elementary"
	} else if avgWordsPerSentence < 15 {
		return "middle_school"
	} else if avgWordsPerSentence < 20 {
		return "high_school"
	} else {
		return "college"
	}
}

func (h *SearchHandler) analyzeTone(text string) string {
	text = strings.ToLower(text)
	
	// Simple keyword-based tone analysis
	positiveWords := []string{"happy", "joy", "wonderful", "amazing", "beautiful", "love", "peaceful"}
	negativeWords := []string{"sad", "angry", "terrible", "awful", "hate", "dark", "fear", "death"}
	formalWords := []string{"therefore", "furthermore", "consequently", "moreover", "however"}
	informalWords := []string{"gonna", "wanna", "yeah", "cool", "awesome", "hey"}
	
	positive := h.countWords(text, positiveWords)
	negative := h.countWords(text, negativeWords)
	formal := h.countWords(text, formalWords)
	informal := h.countWords(text, informalWords)
	
	if positive > negative && positive > 2 {
		return "positive"
	} else if negative > positive && negative > 2 {
		return "negative"
	} else if formal > informal {
		return "formal"
	} else if informal > formal {
		return "informal"
	} else {
		return "neutral"
	}
}

func (h *SearchHandler) getCommonWords(text string, limit int) []map[string]interface{} {
	words := strings.Fields(strings.ToLower(text))
	wordCount := make(map[string]int)
	
	// Common stop words to exclude
	stopWords := map[string]bool{
		"the": true, "a": true, "an": true, "and": true, "or": true, "but": true,
		"in": true, "on": true, "at": true, "to": true, "for": true, "of": true,
		"with": true, "by": true, "is": true, "was": true, "are": true, "were": true,
		"be": true, "been": true, "have": true, "has": true, "had": true, "do": true,
		"does": true, "did": true, "will": true, "would": true, "could": true, "should": true,
		"i": true, "you": true, "he": true, "she": true, "it": true, "we": true, "they": true,
	}
	
	for _, word := range words {
		word = strings.Trim(word, ".,!?;:")
		if len(word) > 2 && !stopWords[word] {
			wordCount[word]++
		}
	}
	
	// Convert to sorted slice
	type wordFreq struct {
		Word  string
		Count int
	}
	
	var frequencies []wordFreq
	for word, count := range wordCount {
		frequencies = append(frequencies, wordFreq{Word: word, Count: count})
	}
	
	// Sort by count (descending)
	for i := 0; i < len(frequencies)-1; i++ {
		for j := i + 1; j < len(frequencies); j++ {
			if frequencies[j].Count > frequencies[i].Count {
				frequencies[i], frequencies[j] = frequencies[j], frequencies[i]
			}
		}
	}
	
	// Return top words
	var result []map[string]interface{}
	maxResults := limit
	if maxResults > len(frequencies) {
		maxResults = len(frequencies)
	}
	
	for i := 0; i < maxResults; i++ {
		result = append(result, map[string]interface{}{
			"word":  frequencies[i].Word,
			"count": frequencies[i].Count,
		})
	}
	
	return result
}

func (h *SearchHandler) analyzeWritingStyle(text string) map[string]interface{} {
	return map[string]interface{}{
		"hasDialogue":             strings.Contains(text, "\""),
		"hasQuestions":            strings.Contains(text, "?"),
		"hasExclamations":         strings.Contains(text, "!"),
		"averageSentenceLength":   h.calculateAverageWordsPerSentence(text),
		"complexSentences":        h.countComplexSentences(text),
		"narrativeStyle":          h.detectNarrativeStyle(text),
	}
}

func (h *SearchHandler) countWords(text string, words []string) int {
	count := 0
	for _, word := range words {
		count += strings.Count(text, word)
	}
	return count
}

func (h *SearchHandler) countComplexSentences(text string) int {
	// Simple heuristic: sentences with semicolons, em-dashes, or multiple clauses
	return strings.Count(text, ";") + strings.Count(text, "—") + strings.Count(text, " which ") + strings.Count(text, " that ")
}

func (h *SearchHandler) detectNarrativeStyle(text string) string {
	text = strings.ToLower(text)
	
	firstPerson := strings.Count(text, " i ") + strings.Count(text, " me ") + strings.Count(text, " my ")
	secondPerson := strings.Count(text, " you ") + strings.Count(text, " your ")
	thirdPerson := strings.Count(text, " he ") + strings.Count(text, " she ") + strings.Count(text, " they ")
	
	if firstPerson > secondPerson && firstPerson > thirdPerson {
		return "first_person"
	} else if secondPerson > firstPerson && secondPerson > thirdPerson {
		return "second_person"
	} else {
		return "third_person"
	}
}

// Utility methods

func (h *SearchHandler) matchesQuery(query string, texts ...string) bool {
	for _, text := range texts {
		if strings.Contains(text, query) {
			return true
		}
	}
	return false
}

func (h *SearchHandler) createSnippet(text, query string) string {
	text = strings.ToLower(text)
	query = strings.ToLower(query)
	
	index := strings.Index(text, query)
	if index == -1 {
		if len(text) > 100 {
			return text[:100] + "..."
		}
		return text
	}

	start := index - 50
	if start < 0 {
		start = 0
	}

	end := index + len(query) + 50
	if end > len(text) {
		end = len(text)
	}

	snippet := text[start:end]
	if start > 0 {
		snippet = "..." + snippet
	}
	if end < len(text) {
		snippet = snippet + "..."
	}

	return snippet
}

func (h *SearchHandler) calculateScore(query string, texts ...string) float64 {
	score := 0.0
	query = strings.ToLower(query)

	for _, text := range texts {
		text = strings.ToLower(text)
		
		if strings.Contains(text, query) {
			score += 1.0
		}
		
		queryWords := strings.Fields(query)
		textWords := strings.Fields(text)
		matches := 0
		
		for _, qWord := range queryWords {
			for _, tWord := range textWords {
				if strings.Contains(tWord, qWord) {
					matches++
					break
				}
			}
		}
		
		if len(queryWords) > 0 {
			score += float64(matches) / float64(len(queryWords)) * 0.5
		}
	}

	return score
}

func (h *SearchHandler) sortResultsByScore(results []models.SearchResult) {
	for i := 0; i < len(results)-1; i++ {
		for j := i + 1; j < len(results); j++ {
			if results[j].Score > results[i].Score {
				results[i], results[j] = results[j], results[i]
			}
		}
	}
}

func (h *SearchHandler) isChapterInRange(chapterNumber, chapterRange string) bool {
	// Simple range check - could be enhanced
	if strings.Contains(chapterRange, "-") {
		// Range format: "1-5"
		parts := strings.Split(chapterRange, "-")
		if len(parts) == 2 {
			return chapterNumber >= parts[0] && chapterNumber <= parts[1]
		}
	} else {
		// Single chapter
		return chapterNumber == chapterRange
	}
	return true
}

func (h *SearchHandler) findMentionSnippets(text, characterName string, maxSnippets int) []string {
	var snippets []string
	textLower := strings.ToLower(text)
	nameLower := strings.ToLower(characterName)
	
	startIndex := 0
	for len(snippets) < maxSnippets {
		index := strings.Index(textLower[startIndex:], nameLower)
		if index == -1 {
			break
		}
		
		actualIndex := startIndex + index
		start := actualIndex - 50
		if start < 0 {
			start = 0
		}
		
		end := actualIndex + len(nameLower) + 50
		if end > len(text) {
			end = len(text)
		}
		
		snippet := text[start:end]
		if start > 0 {
			snippet = "..." + snippet
		}
		if end < len(text) {
			snippet = snippet + "..."
		}
		
		snippets = append(snippets, snippet)
		startIndex = actualIndex + len(nameLower)
	}
	
	return snippets
}

func (h *SearchHandler) extractEventsFromChapter(chapter models.Chapter, eventType string) []map[string]interface{} {
	var events []map[string]interface{}
	
	// Simple event extraction based on keywords and patterns
	content := strings.ToLower(chapter.Content)
	
	// Look for action keywords
	actionKeywords := []string{"arrived", "left", "fought", "died", "married", "discovered", "revealed", "decided"}
	
	for _, keyword := range actionKeywords {
		if eventType == "" || eventType == "action" {
			if strings.Contains(content, keyword) {
				events = append(events, map[string]interface{}{
					"id":          fmt.Sprintf("%s_%s", chapter.ID, keyword),
					"type":        "action",
					"title":       fmt.Sprintf("Action: %s", keyword),
					"chapter":     chapter.Number,
					"description": h.createSnippet(chapter.Content, keyword),
					"status":      "occurred",
					"extractedAt": time.Now(),
				})
			}
		}
	}
	
	return events
}

func (h *SearchHandler) sortEventsByChapter(events []map[string]interface{}) {
	for i := 0; i < len(events)-1; i++ {
		for j := i + 1; j < len(events); j++ {
			chapterI, okI := events[i]["chapter"].(string)
			chapterJ, okJ := events[j]["chapter"].(string)
			
			if okI && okJ && chapterJ < chapterI {
				events[i], events[j] = events[j], events[i]
			}
		}
	}
}
