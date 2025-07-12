import { vi } from 'vitest'
import { BaseOption, CharacterOption, LocationOption, RuleOption, CodexOption, SampleChapterOption, TaskTypeOption } from '../types'

// Test data fixtures
export const mockCharacters: CharacterOption[] = [
  {
    id: '1',
    label: 'John Doe',
    description: 'Protagonist, 30s, detective with a troubled past'
  },
  {
    id: '2',
    label: 'Jane Smith',
    description: 'Supporting character, mysterious woman with secrets'
  }
]

export const mockLocations: LocationOption[] = [
  {
    id: '1',
    label: 'Dark Alley',
    description: 'Narrow, dimly lit alley in downtown district'
  },
  {
    id: '2',
    label: 'Police Station',
    description: 'Busy metropolitan police headquarters'
  }
]

export const mockRules: RuleOption[] = [
  {
    id: '1',
    label: 'Show, Don\'t Tell',
    description: 'Use actions and dialogue instead of exposition'
  },
  {
    id: '2',
    label: 'Third Person Limited',
    description: 'Maintain consistent point of view'
  }
]

export const mockCodex: CodexOption[] = [
  {
    id: '1',
    label: 'Magic System',
    description: 'Hard magic system based on elemental manipulation'
  }
]

export const mockSampleChapters: SampleChapterOption[] = [
  {
    id: '1',
    label: 'Chapter 1 Opening',
    content: 'It was a dark and stormy night when Detective John Doe received the call that would change everything.'
  }
]

export const mockTaskTypes: TaskTypeOption[] = [
  {
    id: '1',
    label: 'Continue Story',
    description: 'Write the next chapter continuing the story'
  },
  {
    id: '2',
    label: 'Character Development',
    description: 'Focus on developing character relationships'
  }
]

export const mockSettings = {
  theme: 'dark',
  defaultLanguage: 'en',
  enableAutoSave: true
}

// Utility functions for tests
export const createMockOption = (type: string, overrides: Partial<BaseOption> = {}): BaseOption => ({
  id: `mock-${type}-${Math.random().toString(36).substr(2, 9)}`,
  label: `Mock ${type}`,
  description: `Mock ${type} description`,
  ...overrides
})

export const createMockFileContent = (content: any) => JSON.stringify(content, null, 2)

// Mock Wails API responses
export const mockWailsSuccess = (data: any = true) => Promise.resolve(data)
export const mockWailsError = (message: string = 'Mock error') => Promise.reject(new Error(message))

// Test component wrapper helpers
export const mockWailsContext = {
  isReady: true,
  runtime: {
    LogInfo: vi.fn(),
    LogError: vi.fn(),
    EventsOn: vi.fn(),
    EventsOff: vi.fn(),
    EventsEmit: vi.fn(),
    ClipboardSetText: vi.fn(),
    ClipboardGetText: vi.fn().mockResolvedValue(''),
  }
}

// Helper to wait for async operations
export const waitFor = (ms: number = 100) => new Promise(resolve => setTimeout(resolve, ms))

// Helper to create mock events
export const createMockEvent = (eventData: any = {}) => ({
  preventDefault: vi.fn(),
  stopPropagation: vi.fn(),
  target: { value: '' },
  ...eventData
})
