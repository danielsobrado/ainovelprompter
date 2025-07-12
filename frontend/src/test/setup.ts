import '@testing-library/jest-dom'
import { vi } from 'vitest'

// Mock Wails runtime
const mockWails = {
  runtime: {
    LogInfo: vi.fn(),
    LogError: vi.fn(),
    EventsOn: vi.fn(),
    EventsOff: vi.fn(),
    EventsEmit: vi.fn(),
    ClipboardSetText: vi.fn(),
    ClipboardGetText: vi.fn().mockResolvedValue(''),
  },
  go: {
    main: {
      App: {
        ReadSettingsFile: vi.fn().mockResolvedValue('{}'),
        WriteSettingsFile: vi.fn().mockResolvedValue(true),
        ReadCharactersFile: vi.fn().mockResolvedValue('[]'),
        WriteCharactersFile: vi.fn().mockResolvedValue(true),
        ReadLocationsFile: vi.fn().mockResolvedValue('[]'),
        WriteLocationsFile: vi.fn().mockResolvedValue(true),
        ReadRulesFile: vi.fn().mockResolvedValue('[]'),
        WriteRulesFile: vi.fn().mockResolvedValue(true),
        ReadCodexFile: vi.fn().mockResolvedValue('[]'),
        WriteCodexFile: vi.fn().mockResolvedValue(true),
        ReadSampleChaptersFile: vi.fn().mockResolvedValue('[]'),
        WriteSampleChaptersFile: vi.fn().mockResolvedValue(true),
        ReadTaskTypesFile: vi.fn().mockResolvedValue('[]'),
        WriteTaskTypesFile: vi.fn().mockResolvedValue(true),
        SelectFile: vi.fn().mockResolvedValue(''),
        SelectDirectory: vi.fn().mockResolvedValue(''),
        ReadFileContent: vi.fn().mockResolvedValue(''),
        HandleFileDrop: vi.fn().mockResolvedValue(''),
        ProcessFolder: vi.fn().mockResolvedValue(''),
        LogInfo: vi.fn(),
        GetCurrentDirectory: vi.fn().mockResolvedValue('/tmp'),
        ReadProsePromptsFile: vi.fn().mockResolvedValue('[]'),
        WriteProsePromptsFile: vi.fn().mockResolvedValue(true),
        ReadLLMSettingsFile: vi.fn().mockResolvedValue('{}'),
        WriteLLMSettingsFile: vi.fn().mockResolvedValue(true),
        GetInitialLLMSettings: vi.fn().mockResolvedValue({}),
        LoadPromptDefinitions: vi.fn().mockResolvedValue([]),
        SavePromptDefinitions: vi.fn().mockResolvedValue(true),
        FindPromptDefinitionByID: vi.fn().mockResolvedValue(null),
        UpdatePromptDefinition: vi.fn().mockResolvedValue(true),
        DeletePromptDefinition: vi.fn().mockResolvedValue(true),
        GetPromptVariantsForModel: vi.fn().mockResolvedValue([]),
        GetResolvedProsePrompt: vi.fn().mockResolvedValue(''),
      }
    }
  }
}

// Set up global window object
Object.defineProperty(window, 'wails', {
  value: mockWails,
  writable: true
})

// Mock crypto.randomUUID for ID generation
Object.defineProperty(global, 'crypto', {
  value: {
    randomUUID: vi.fn(() => 'mock-uuid-' + Math.random().toString(36).substr(2, 9))
  }
})

// Mock localStorage
const localStorageMock = {
  getItem: vi.fn(),
  setItem: vi.fn(),
  removeItem: vi.fn(),
  clear: vi.fn(),
}
Object.defineProperty(window, 'localStorage', {
  value: localStorageMock
})

// Mock ResizeObserver
global.ResizeObserver = vi.fn().mockImplementation(() => ({
  observe: vi.fn(),
  unobserve: vi.fn(),
  disconnect: vi.fn(),
}))

// Mock fetch for API calls
global.fetch = vi.fn()
