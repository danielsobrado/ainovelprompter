import { describe, it, expect, vi } from 'vitest'
import { generateChatGPTPrompt, generateClaudePrompt } from '../promptGenerators'
import { mockCharacters, mockLocations, mockRules, mockCodex, mockSampleChapters, mockTaskTypes } from '../../test/mocks'

describe('Prompt Generators', () => {
  const mockOptions = {
    selectedCharacters: mockCharacters,
    selectedLocations: mockLocations,
    selectedRules: mockRules,
    selectedCodex: mockCodex,
    selectedSampleChapters: mockSampleChapters,
    selectedTaskTypes: mockTaskTypes,
    previousChapter: 'Previous chapter content...',
    futureChapterNotes: 'Future chapter notes...',
    beats: 'Story beats...'
  }

  describe('generateChatGPTPrompt', () => {
    it('should generate a valid ChatGPT prompt', () => {
      const result = generateChatGPTPrompt(mockOptions)
      
      expect(result).toContain('You are a skilled creative writer')
      expect(result).toContain('Characters:')
      expect(result).toContain('John Doe')
      expect(result).toContain('Locations:')
      expect(result).toContain('Dark Alley')
      expect(result).toContain('Rules:')
      expect(result).toContain('Show, Don\'t Tell')
    })

    it('should include previous chapter when provided', () => {
      const result = generateChatGPTPrompt(mockOptions)
      expect(result).toContain('Previous chapter content...')
    })

    it('should include future notes when provided', () => {
      const result = generateChatGPTPrompt(mockOptions)
      expect(result).toContain('Future chapter notes...')
    })

    it('should include story beats when provided', () => {
      const result = generateChatGPTPrompt(mockOptions)
      expect(result).toContain('Story beats...')
    })

    it('should handle empty selections gracefully', () => {
      const emptyOptions = {
        selectedCharacters: [],
        selectedLocations: [],
        selectedRules: [],
        selectedCodex: [],
        selectedSampleChapters: [],
        selectedTaskTypes: [],
        previousChapter: '',
        futureChapterNotes: '',
        beats: ''
      }
      
      const result = generateChatGPTPrompt(emptyOptions)
      expect(result).toContain('You are a skilled creative writer')
      expect(typeof result).toBe('string')
      expect(result.length).toBeGreaterThan(0)
    })
  })

  describe('generateClaudePrompt', () => {
    it('should generate a valid Claude prompt with XML tags', () => {
      const result = generateClaudePrompt(mockOptions)
      
      expect(result).toContain('<characters>')
      expect(result).toContain('</characters>')
      expect(result).toContain('<locations>')
      expect(result).toContain('</locations>')
      expect(result).toContain('<rules>')
      expect(result).toContain('</rules>')
      expect(result).toContain('John Doe')
      expect(result).toContain('Dark Alley')
    })

    it('should include task types with proper formatting', () => {
      const result = generateClaudePrompt(mockOptions)
      expect(result).toContain('<task>')
      expect(result).toContain('Continue Story')
      expect(result).toContain('</task>')
    })

    it('should include codex information', () => {
      const result = generateClaudePrompt(mockOptions)
      expect(result).toContain('<codex>')
      expect(result).toContain('Magic System')
      expect(result).toContain('</codex>')
    })

    it('should handle empty selections with proper XML structure', () => {
      const emptyOptions = {
        selectedCharacters: [],
        selectedLocations: [],
        selectedRules: [],
        selectedCodex: [],
        selectedSampleChapters: [],
        selectedTaskTypes: [],
        previousChapter: '',
        futureChapterNotes: '',
        beats: ''
      }
      
      const result = generateClaudePrompt(emptyOptions)
      expect(result).toContain('<characters>')
      expect(result).toContain('</characters>')
      expect(typeof result).toBe('string')
    })
  })
})
