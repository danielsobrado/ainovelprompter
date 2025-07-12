package handlers

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/danielsobrado/ainovelprompter/mcp/models"
	"github.com/danielsobrado/ainovelprompter/mcp/storage"
)

type ChapterHandler struct {
	storage storage.ChapterStorage
}

func NewChapterHandler(storage storage.ChapterStorage) *ChapterHandler {
	return &ChapterHandler{
		storage: storage,
	}
}

// Chapter handlers
func (h *ChapterHandler) GetChapters(params map[string]interface{}) (interface{}, error) {
	// Check for search parameter
	if searchQuery, ok := params["search"].(string); ok && searchQuery != "" {
		return h.storage.SearchChapters(searchQuery)
	}

	// Check for range parameters
	if startStr, ok := params["start"].(string); ok {
		start, _ := strconv.Atoi(startStr)
		end := start

		if endStr, ok := params["end"].(string); ok {
			end, _ = strconv.Atoi(endStr)
		}

		return h.storage.GetChapterRange(start, end)
	}

	// Check for status filter
	if status, ok := params["status"].(string); ok {
		chapters, err := h.storage.GetChapters()
		if err != nil {
			return nil, err
		}

		var filtered []models.Chapter
		for _, ch := range chapters {
			if ch.Status == status {
				filtered = append(filtered, ch)
			}
		}
		return filtered, nil
	}

	// Return all chapters
	return h.storage.GetChapters()
}

func (h *ChapterHandler) GetChapterContent(params map[string]interface{}) (interface{}, error) {
	// Support both ID and number
	if id, ok := params["id"].(string); ok {
		return h.storage.GetChapterByID(id)
	}

	if numberStr, ok := params["number"].(string); ok {
		number, err := strconv.Atoi(numberStr)
		if err != nil {
			return nil, fmt.Errorf("invalid chapter number")
		}
		return h.storage.GetChapterByNumber(number)
	}

	return nil, fmt.Errorf("chapter ID or number is required")
}

func (h *ChapterHandler) GetPreviousChapter(params map[string]interface{}) (interface{}, error) {
	currentStr, ok := params["currentChapter"].(string)
	if !ok {
		// If no current chapter specified, get the last chapter
		chapters, err := h.storage.GetChapters()
		if err != nil {
			return nil, err
		}

		if len(chapters) == 0 {
			return nil, fmt.Errorf("no chapters found")
		}

		// Chapters are already sorted by number
		return &chapters[len(chapters)-1], nil
	}

	current, err := strconv.Atoi(currentStr)
	if err != nil {
		return nil, fmt.Errorf("invalid chapter number")
	}

	return h.storage.GetPreviousChapter(current)
}

func (h *ChapterHandler) CreateChapter(params map[string]interface{}) (interface{}, error) {
	var chapter models.Chapter

	if title, ok := params["title"].(string); ok {
		chapter.Title = title
	}

	if content, ok := params["content"].(string); ok {
		chapter.Content = content
	} else {
		return nil, fmt.Errorf("chapter content is required")
	}

	if numberStr, ok := params["number"].(string); ok {
		chapter.Number, _ = strconv.Atoi(numberStr)
	}

	if summary, ok := params["summary"].(string); ok {
		chapter.Summary = summary
	}

	if status, ok := params["status"].(string); ok {
		chapter.Status = status
	} else {
		chapter.Status = "draft"
	}

	if beats, ok := params["storyBeats"].([]interface{}); ok {
		for _, beat := range beats {
			if beatStr, ok := beat.(string); ok {
				chapter.StoryBeats = append(chapter.StoryBeats, beatStr)
			}
		}
	}

	if charIDs, ok := params["characterIds"].([]interface{}); ok {
		for _, id := range charIDs {
			if idStr, ok := id.(string); ok {
				chapter.CharacterIDs = append(chapter.CharacterIDs, idStr)
			}
		}
	}

	if locIDs, ok := params["locationIds"].([]interface{}); ok {
		for _, id := range locIDs {
			if idStr, ok := id.(string); ok {
				chapter.LocationIDs = append(chapter.LocationIDs, idStr)
			}
		}
	}

	err := h.storage.CreateChapter(&chapter)
	if err != nil {
		return nil, err
	}

	return chapter, nil
}

func (h *ChapterHandler) UpdateChapter(params map[string]interface{}) (interface{}, error) {
	id, ok := params["id"].(string)
	if !ok {
		return nil, fmt.Errorf("chapter ID is required")
	}

	chapter, err := h.storage.GetChapterByID(id)
	if err != nil {
		return nil, err
	}

	if title, ok := params["title"].(string); ok {
		chapter.Title = title
	}

	if content, ok := params["content"].(string); ok {
		chapter.Content = content
	}

	if summary, ok := params["summary"].(string); ok {
		chapter.Summary = summary
	}

	if status, ok := params["status"].(string); ok {
		chapter.Status = status
	}

	err = h.storage.UpdateChapter(chapter)
	if err != nil {
		return nil, err
	}

	return chapter, nil
}

func (h *ChapterHandler) DeleteChapter(params map[string]interface{}) (interface{}, error) {
	id, ok := params["id"].(string)
	if !ok {
		return nil, fmt.Errorf("chapter ID is required")
	}

	err := h.storage.DeleteChapter(id)
	if err != nil {
		return nil, err
	}

	return map[string]string{"status": "deleted", "id": id}, nil
}

// Story beats handlers
func (h *ChapterHandler) GetStoryBeats(params map[string]interface{}) (interface{}, error) {
	if chapterStr, ok := params["chapter"].(string); ok {
		chapter, err := strconv.Atoi(chapterStr)
		if err != nil {
			return nil, fmt.Errorf("invalid chapter number")
		}
		return h.storage.GetStoryBeats(chapter)
	}

	// Return all story beats
	return h.storage.GetAllStoryBeats()
}

func (h *ChapterHandler) SaveStoryBeats(params map[string]interface{}) (interface{}, error) {
	var beats models.StoryBeats

	if chapterStr, ok := params["chapter"].(string); ok {
		beats.ChapterNumber, _ = strconv.Atoi(chapterStr)
	} else {
		return nil, fmt.Errorf("chapter number is required")
	}

	if beatsData, ok := params["beats"].([]interface{}); ok {
		for i, beatData := range beatsData {
			if beatMap, ok := beatData.(map[string]interface{}); ok {
				beat := models.Beat{
					Order: i,
				}

				if desc, ok := beatMap["description"].(string); ok {
					beat.Description = desc
				}
				if beatType, ok := beatMap["type"].(string); ok {
					beat.Type = beatType
				}
				if completed, ok := beatMap["completed"].(bool); ok {
					beat.Completed = completed
				}

				beats.Beats = append(beats.Beats, beat)
			}
		}
	}

	if notes, ok := params["notes"].(string); ok {
		beats.Notes = notes
	}

	err := h.storage.SaveStoryBeats(&beats)
	if err != nil {
		return nil, err
	}

	return beats, nil
}

// Future notes handlers
func (h *ChapterHandler) GetFutureNotes(params map[string]interface{}) (interface{}, error) {
	if chapterStr, ok := params["chapter"].(string); ok {
		chapter, err := strconv.Atoi(chapterStr)
		if err != nil {
			return nil, fmt.Errorf("invalid chapter number")
		}
		return h.storage.GetFutureNotesByChapter(chapter)
	}

	if priority, ok := params["priority"].(string); ok {
		notes, err := h.storage.GetFutureNotes()
		if err != nil {
			return nil, err
		}

		var filtered []models.FutureNotes
		for _, note := range notes {
			if note.Priority == priority {
				filtered = append(filtered, note)
			}
		}
		return filtered, nil
	}

	return h.storage.GetFutureNotes()
}

func (h *ChapterHandler) CreateFutureNote(params map[string]interface{}) (interface{}, error) {
	var note models.FutureNotes

	if chapterRange, ok := params["chapterRange"].(string); ok {
		note.ChapterRange = chapterRange
	} else {
		return nil, fmt.Errorf("chapter range is required")
	}

	if content, ok := params["content"].(string); ok {
		note.Content = content
	} else {
		return nil, fmt.Errorf("content is required")
	}

	if priority, ok := params["priority"].(string); ok {
		note.Priority = priority
	} else {
		note.Priority = "medium"
	}

	if category, ok := params["category"].(string); ok {
		note.Category = category
	}

	if tags, ok := params["tags"].([]interface{}); ok {
		for _, tag := range tags {
			if tagStr, ok := tag.(string); ok {
				note.Tags = append(note.Tags, tagStr)
			}
		}
	}

	err := h.storage.CreateFutureNote(&note)
	if err != nil {
		return nil, err
	}

	return note, nil
}

// Sample chapters handlers
func (h *ChapterHandler) GetSampleChapters(params map[string]interface{}) (interface{}, error) {
	if purpose, ok := params["purpose"].(string); ok {
		return h.storage.GetSampleChaptersByPurpose(purpose)
	}

	return h.storage.GetSampleChapters()
}

func (h *ChapterHandler) GetSampleChapterByID(params map[string]interface{}) (interface{}, error) {
	id, ok := params["id"].(string)
	if !ok {
		return nil, fmt.Errorf("sample chapter ID is required")
	}

	return h.storage.GetSampleChapterByID(id)
}

func (h *ChapterHandler) CreateSampleChapter(params map[string]interface{}) (interface{}, error) {
	var sample models.SampleChapter

	if title, ok := params["title"].(string); ok {
		sample.Title = title
	} else {
		return nil, fmt.Errorf("sample chapter title is required")
	}

	if content, ok := params["content"].(string); ok {
		sample.Content = content
	} else {
		return nil, fmt.Errorf("sample chapter content is required")
	}

	if author, ok := params["author"].(string); ok {
		sample.Author = author
	}

	if source, ok := params["source"].(string); ok {
		sample.Source = source
	}

	if purpose, ok := params["purpose"].(string); ok {
		sample.Purpose = purpose
	} else {
		sample.Purpose = "style_reference"
	}

	if tags, ok := params["tags"].([]interface{}); ok {
		for _, tag := range tags {
			if tagStr, ok := tag.(string); ok {
				sample.Tags = append(sample.Tags, tagStr)
			}
		}
	}

	err := h.storage.CreateSampleChapter(&sample)
	if err != nil {
		return nil, err
	}

	return sample, nil
}

// Task types handlers
func (h *ChapterHandler) GetTaskTypes(params map[string]interface{}) (interface{}, error) {
	if activeOnly, ok := params["activeOnly"].(bool); ok && activeOnly {
		return h.storage.GetActiveTaskTypes()
	}

	if category, ok := params["category"].(string); ok {
		types, err := h.storage.GetTaskTypes()
		if err != nil {
			return nil, err
		}

		var filtered []models.TaskType
		for _, tt := range types {
			if tt.Category == category {
				filtered = append(filtered, tt)
			}
		}
		return filtered, nil
	}

	return h.storage.GetTaskTypes()
}

func (h *ChapterHandler) GetTaskTypeByID(params map[string]interface{}) (interface{}, error) {
	id, ok := params["id"].(string)
	if !ok {
		return nil, fmt.Errorf("task type ID is required")
	}

	return h.storage.GetTaskTypeByID(id)
}

// Apply task template
func (h *ChapterHandler) ApplyTaskTemplate(params map[string]interface{}) (interface{}, error) {
	taskID, ok := params["taskId"].(string)
	if !ok {
		return nil, fmt.Errorf("task ID is required")
	}

	taskType, err := h.storage.GetTaskTypeByID(taskID)
	if err != nil {
		return nil, err
	}

	// Apply variables to template
	template := taskType.Template
	if variables, ok := params["variables"].(map[string]interface{}); ok {
		for key, value := range variables {
			if strVal, ok := value.(string); ok {
				placeholder := fmt.Sprintf("{{%s}}", key)
				template = strings.ReplaceAll(template, placeholder, strVal)
			}
		}
	}

	// Also replace default variables
	for key, value := range taskType.Variables {
		placeholder := fmt.Sprintf("{{%s}}", key)
		if !strings.Contains(template, placeholder) {
			continue
		}
		template = strings.ReplaceAll(template, placeholder, value)
	}

	return map[string]interface{}{
		"taskType": taskType,
		"prompt":   template,
	}, nil
}
