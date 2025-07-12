import { describe, it, expect, vi, beforeEach } from 'vitest'
import { render, screen, fireEvent, waitFor } from '@testing-library/react'
import App from '../App'

// Mock all the Wails functions
vi.mock('@/wailsjs/go/main/App', () => ({
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
}))

// Mock the WailsReady context
vi.mock('@/contexts/WailsReadyContext', () => ({
  useWailsReady: () => ({ isReady: true }),
  WailsReadyProvider: ({ children }: any) => children
}))

describe('App Integration Tests', () => {
  beforeEach(() => {
    vi.clearAllMocks()
  })

  it('should render main application components', async () => {
    render(<App />)
    
    // Should render the main tabs
    expect(screen.getByText('Prompt Generator')).toBeInTheDocument()
    expect(screen.getByText('Prose Improvement')).toBeInTheDocument()
    
    // Should render file operations
    expect(screen.getByText('File Operations')).toBeInTheDocument()
  })

  it('should switch between tabs', async () => {
    render(<App />)
    
    // Click on Prose Improvement tab
    const proseTab = screen.getByText('Prose Improvement')
    fireEvent.click(proseTab)
    
    // Should show prose improvement content
    await waitFor(() => {
      expect(screen.getByText(/enter your text/i)).toBeInTheDocument()
    })
  })

  it('should generate prompts', async () => {
    render(<App />)
    
    // Should have prompt generation buttons
    expect(screen.getByText('Generate ChatGPT Prompt')).toBeInTheDocument()
    expect(screen.getByText('Generate Claude Prompt')).toBeInTheDocument()
  })

  it('should handle character selection', async () => {
    render(<App />)
    
    // Should render characters section
    expect(screen.getByText('Characters')).toBeInTheDocument()
    expect(screen.getByText('Edit Characters')).toBeInTheDocument()
  })

  it('should handle task type selection', async () => {
    render(<App />)
    
    // Should render task types section
    expect(screen.getByText('Task Types')).toBeInTheDocument()
  })

  it('should copy prompt to clipboard', async () => {
    // Mock clipboard API
    const mockWriteText = vi.fn().mockResolvedValue(undefined)
    Object.assign(navigator, {
      clipboard: {
        writeText: mockWriteText,
      },
    })

    render(<App />)
    
    // Generate a prompt first
    const generateButton = screen.getByText('Generate ChatGPT Prompt')
    fireEvent.click(generateButton)
    
    // Find and click copy button
    const copyButton = screen.getByText('Copy to Clipboard')
    fireEvent.click(copyButton)
    
    // Should show success message
    await waitFor(() => {
      expect(screen.getByText(/copied to clipboard/i)).toBeInTheDocument()
    })
  })

  it('should clear all fields', async () => {
    render(<App />)
    
    const clearButton = screen.getByText('Clear All')
    fireEvent.click(clearButton)
    
    // All text areas should be cleared
    const textareas = screen.getAllByRole('textbox')
    textareas.forEach(textarea => {
      if (textarea.tagName === 'TEXTAREA') {
        expect(textarea.value).toBe('')
      }
    })
  })

  it('should show token count', async () => {
    render(<App />)
    
    // Should display token count
    expect(screen.getByText(/tokens/i)).toBeInTheDocument()
  })

  it('should handle settings modal', async () => {
    render(<App />)
    
    // Find settings button (usually a gear icon or settings text)
    const settingsButton = screen.getByLabelText(/settings/i) || screen.getByText(/settings/i)
    fireEvent.click(settingsButton)
    
    // Should open settings modal
    await waitFor(() => {
      expect(screen.getByText(/theme/i)).toBeInTheDocument()
    })
  })
})
