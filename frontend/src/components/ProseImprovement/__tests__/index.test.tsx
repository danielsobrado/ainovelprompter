import { describe, it, expect, vi, beforeEach } from 'vitest'
import { render, screen, fireEvent, waitFor } from '@testing-library/react'
import { ProseImprovementTab } from '../index'

// Mock the LLM provider hook
vi.mock('../../hooks/useLLMProvider', () => ({
  useLLMProvider: () => ({
    executePrompt: vi.fn().mockResolvedValue('Improved text response'),
    isLoading: false,
    error: null
  })
}))

// Mock the prose improvement hook
vi.mock('../../hooks/useProseImprovement', () => ({
  useProseImprovement: () => ({
    session: null,
    startSession: vi.fn(),
    processNextPrompt: vi.fn(),
    handleChangeDecision: vi.fn(),
    isProcessing: false
  })
}))

describe('ProseImprovementTab', () => {
  beforeEach(() => {
    vi.clearAllMocks()
  })

  it('should render prose improvement interface', () => {
    render(<ProseImprovementTab />)
    
    expect(screen.getByText('Prose Improvement')).toBeInTheDocument()
    expect(screen.getByPlaceholderText(/enter your text/i)).toBeInTheDocument()
    expect(screen.getByText('Start Improvement Session')).toBeInTheDocument()
  })

  it('should handle text input', () => {
    render(<ProseImprovementTab />)
    
    const textarea = screen.getByPlaceholderText(/enter your text/i)
    fireEvent.change(textarea, { target: { value: 'Test text to improve' } })
    
    expect(textarea.value).toBe('Test text to improve')
  })

  it('should display provider settings', () => {
    render(<ProseImprovementTab />)
    
    expect(screen.getByText('LLM Provider')).toBeInTheDocument()
    expect(screen.getByText('Manual')).toBeInTheDocument()
  })

  it('should show prompt manager', () => {
    render(<ProseImprovementTab />)
    
    expect(screen.getByText('Improvement Prompts')).toBeInTheDocument()
  })

  it('should handle provider selection', () => {
    render(<ProseImprovementTab />)
    
    const select = screen.getByDisplayValue('Manual')
    fireEvent.click(select)
    
    // Should show provider options
    expect(screen.getByText('LM Studio')).toBeInTheDocument()
    expect(screen.getByText('OpenRouter')).toBeInTheDocument()
  })

  it('should validate input before starting session', async () => {
    render(<ProseImprovementTab />)
    
    const startButton = screen.getByText('Start Improvement Session')
    fireEvent.click(startButton)
    
    // Should show validation message for empty input
    await waitFor(() => {
      expect(screen.getByText(/please enter some text/i)).toBeInTheDocument()
    })
  })

  it('should start improvement session with valid input', async () => {
    render(<ProseImprovementTab />)
    
    const textarea = screen.getByPlaceholderText(/enter your text/i)
    fireEvent.change(textarea, { target: { value: 'Text to improve' } })
    
    const startButton = screen.getByText('Start Improvement Session')
    fireEvent.click(startButton)
    
    // Should not show validation error
    expect(screen.queryByText(/please enter some text/i)).not.toBeInTheDocument()
  })

  it('should display character count', () => {
    render(<ProseImprovementTab />)
    
    const textarea = screen.getByPlaceholderText(/enter your text/i)
    fireEvent.change(textarea, { target: { value: 'Test' } })
    
    expect(screen.getByText('4 characters')).toBeInTheDocument()
  })

  it('should handle manual response input', () => {
    render(<ProseImprovementTab />)
    
    // First start a session
    const textarea = screen.getByPlaceholderText(/enter your text/i)
    fireEvent.change(textarea, { target: { value: 'Text to improve' } })
    
    const startButton = screen.getByText('Start Improvement Session')
    fireEvent.click(startButton)
    
    // Should show manual response input for Manual provider
    expect(screen.getByPlaceholderText(/enter the ai response/i)).toBeInTheDocument()
  })
})
