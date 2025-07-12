import { describe, it, expect, vi, beforeEach } from 'vitest'
import { render, screen, fireEvent, waitFor } from '@testing-library/react'
import { FinalPrompt } from '../FinalPrompt'

describe('FinalPrompt', () => {
  const defaultProps = {
    finalPrompt: 'This is a sample prompt for testing',
    tokenCount: 42,
    onCopy: vi.fn(),
    onClear: vi.fn()
  }

  beforeEach(() => {
    vi.clearAllMocks()
  })

  it('should render prompt text', () => {
    render(<FinalPrompt {...defaultProps} />)
    expect(screen.getByDisplayValue('This is a sample prompt for testing')).toBeInTheDocument()
  })

  it('should display token count', () => {
    render(<FinalPrompt {...defaultProps} />)
    expect(screen.getByText(/42 tokens/i)).toBeInTheDocument()
  })

  it('should handle copy action', () => {
    const onCopy = vi.fn()
    render(<FinalPrompt {...defaultProps} onCopy={onCopy} />)
    
    const copyButton = screen.getByText(/copy/i)
    fireEvent.click(copyButton)
    
    expect(onCopy).toHaveBeenCalled()
  })

  it('should handle clear action', () => {
    const onClear = vi.fn()
    render(<FinalPrompt {...defaultProps} onClear={onClear} />)
    
    const clearButton = screen.getByText(/clear/i)
    fireEvent.click(clearButton)
    
    expect(onClear).toHaveBeenCalled()
  })

  it('should show empty state when no prompt', () => {
    render(<FinalPrompt {...defaultProps} finalPrompt="" tokenCount={0} />)
    
    const textarea = screen.getByRole('textbox')
    expect(textarea).toHaveValue('')
    expect(screen.getByText(/0 tokens/i)).toBeInTheDocument()
  })

  it('should be readonly', () => {
    render(<FinalPrompt {...defaultProps} />)
    
    const textarea = screen.getByRole('textbox')
    expect(textarea).toHaveAttribute('readonly')
  })

  it('should format large token counts', () => {
    render(<FinalPrompt {...defaultProps} tokenCount={1500} />)
    expect(screen.getByText(/1,500 tokens/i)).toBeInTheDocument()
  })

  it('should disable copy button when prompt is empty', () => {
    render(<FinalPrompt {...defaultProps} finalPrompt="" />)
    
    const copyButton = screen.getByText(/copy/i)
    expect(copyButton).toBeDisabled()
  })

  it('should disable clear button when prompt is empty', () => {
    render(<FinalPrompt {...defaultProps} finalPrompt="" />)
    
    const clearButton = screen.getByText(/clear/i)
    expect(clearButton).toBeDisabled()
  })
})
