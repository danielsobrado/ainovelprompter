import { describe, it, expect, vi } from 'vitest'
import { render, screen, fireEvent, waitFor } from '@testing-library/react'
import { CharactersSelector } from '../CharactersSelector'
import { mockCharacters, createMockEvent } from '../../test/mocks'

// Mock the edit modal
vi.mock('./CharactersEditModal', () => ({
  CharactersEditModal: ({ isOpen, onClose }: any) => 
    isOpen ? <div data-testid="edit-modal">Edit Modal<button onClick={onClose}>Close</button></div> : null
}))

describe('CharactersSelector', () => {
  const defaultProps = {
    options: mockCharacters,
    selectedOptions: [],
    onSelectionChange: vi.fn(),
    onOptionsChange: vi.fn()
  }

  beforeEach(() => {
    vi.clearAllMocks()
  })

  it('should render component with title', () => {
    render(<CharactersSelector {...defaultProps} />)
    expect(screen.getByText('Characters')).toBeInTheDocument()
  })

  it('should display available character options', () => {
    render(<CharactersSelector {...defaultProps} />)
    expect(screen.getByText('John Doe')).toBeInTheDocument()
    expect(screen.getByText('Jane Smith')).toBeInTheDocument()
  })

  it('should handle character selection', () => {
    const onSelectionChange = vi.fn()
    render(<CharactersSelector {...defaultProps} onSelectionChange={onSelectionChange} />)
    
    const checkbox = screen.getByRole('checkbox', { name: /john doe/i })
    fireEvent.click(checkbox)
    
    expect(onSelectionChange).toHaveBeenCalledWith([mockCharacters[0]])
  })

  it('should show selected characters as checked', () => {
    render(<CharactersSelector {...defaultProps} selectedOptions={[mockCharacters[0]]} />)
    
    const checkbox = screen.getByRole('checkbox', { name: /john doe/i })
    expect(checkbox).toBeChecked()
  })

  it('should open edit modal when edit button is clicked', async () => {
    render(<CharactersSelector {...defaultProps} />)
    
    const editButton = screen.getByText('Edit Characters')
    fireEvent.click(editButton)
    
    await waitFor(() => {
      expect(screen.getByTestId('edit-modal')).toBeInTheDocument()
    })
  })

  it('should close edit modal', async () => {
    render(<CharactersSelector {...defaultProps} />)
    
    const editButton = screen.getByText('Edit Characters')
    fireEvent.click(editButton)
    
    await waitFor(() => {
      expect(screen.getByTestId('edit-modal')).toBeInTheDocument()
    })
    
    const closeButton = screen.getByText('Close')
    fireEvent.click(closeButton)
    
    await waitFor(() => {
      expect(screen.queryByTestId('edit-modal')).not.toBeInTheDocument()
    })
  })

  it('should handle deselection of characters', () => {
    const onSelectionChange = vi.fn()
    render(<CharactersSelector 
      {...defaultProps} 
      selectedOptions={[mockCharacters[0]]} 
      onSelectionChange={onSelectionChange} 
    />)
    
    const checkbox = screen.getByRole('checkbox', { name: /john doe/i })
    fireEvent.click(checkbox)
    
    expect(onSelectionChange).toHaveBeenCalledWith([])
  })

  it('should display character descriptions', () => {
    render(<CharactersSelector {...defaultProps} />)
    expect(screen.getByText(/protagonist.*30s.*detective/i)).toBeInTheDocument()
  })

  it('should handle empty options array', () => {
    render(<CharactersSelector {...defaultProps} options={[]} />)
    expect(screen.getByText('Characters')).toBeInTheDocument()
    expect(screen.queryByRole('checkbox')).not.toBeInTheDocument()
  })
})
