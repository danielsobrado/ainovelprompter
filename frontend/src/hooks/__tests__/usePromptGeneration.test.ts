import { describe, it, expect, vi, beforeEach } from 'vitest'
import { renderHook, act } from '@testing-library/react'
import { usePromptGeneration } from '../usePromptGeneration'
import { PromptType } from '../../types'
import { mockCharacters, mockLocations, mockRules, mockCodex, mockSampleChapters, mockTaskTypes } from '../../test/mocks'

describe('usePromptGeneration', () => {
  beforeEach(() => {
    vi.clearAllMocks()
  })

  const mockOptions = {
    selectedCharacters: mockCharacters,
    selectedLocations: mockLocations,
    selectedRules: mockRules,
    selectedCodex: mockCodex,
    selectedSampleChapters: mockSampleChapters,
    selectedTaskTypes: mockTaskTypes,
    previousChapter: 'Previous chapter content',
    futureChapterNotes: 'Future chapter notes',
    beats: 'Story beats'
  }

  it('should initialize with empty prompt', () => {
    const { result } = renderHook(() => usePromptGeneration())

    expect(result.current.finalPrompt).toBe('')
    expect(result.current.tokenCount).toBe(0)
  })

  it('should generate ChatGPT prompt', () => {
    const { result } = renderHook(() => usePromptGeneration())

    act(() => {
      result.current.generatePrompt(PromptType.CHATGPT, mockOptions)
    })

    expect(result.current.finalPrompt).toContain('You are a skilled creative writer')
    expect(result.current.finalPrompt).toContain('John Doe')
    expect(result.current.tokenCount).toBeGreaterThan(0)
  })

  it('should generate Claude prompt', () => {
    const { result } = renderHook(() => usePromptGeneration())

    act(() => {
      result.current.generatePrompt(PromptType.CLAUDE, mockOptions)
    })

    expect(result.current.finalPrompt).toContain('<characters>')
    expect(result.current.finalPrompt).toContain('John Doe')
    expect(result.current.finalPrompt).toContain('</characters>')
    expect(result.current.tokenCount).toBeGreaterThan(0)
  })

  it('should clear prompt', () => {
    const { result } = renderHook(() => usePromptGeneration())

    act(() => {
      result.current.generatePrompt(PromptType.CHATGPT, mockOptions)
    })

    expect(result.current.finalPrompt).not.toBe('')

    act(() => {
      result.current.clearPrompt()
    })

    expect(result.current.finalPrompt).toBe('')
    expect(result.current.tokenCount).toBe(0)
  })

  it('should handle empty options', () => {
    const { result } = renderHook(() => usePromptGeneration())

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

    act(() => {
      result.current.generatePrompt(PromptType.CHATGPT, emptyOptions)
    })

    expect(result.current.finalPrompt).toContain('You are a skilled creative writer')
    expect(result.current.tokenCount).toBeGreaterThan(0)
  })

  it('should calculate token count correctly', () => {
    const { result } = renderHook(() => usePromptGeneration())

    act(() => {
      result.current.generatePrompt(PromptType.CHATGPT, mockOptions)
    })

    const expectedTokens = Math.ceil(result.current.finalPrompt.length / 4)
    expect(result.current.tokenCount).toBeCloseTo(expectedTokens, -1)
  })

  it('should update prompt when options change', () => {
    const { result } = renderHook(() => usePromptGeneration())

    act(() => {
      result.current.generatePrompt(PromptType.CHATGPT, mockOptions)
    })

    const firstPrompt = result.current.finalPrompt

    const updatedOptions = {
      ...mockOptions,
      previousChapter: 'Updated previous chapter content'
    }

    act(() => {
      result.current.generatePrompt(PromptType.CHATGPT, updatedOptions)
    })

    expect(result.current.finalPrompt).not.toBe(firstPrompt)
    expect(result.current.finalPrompt).toContain('Updated previous chapter content')
  })
})
