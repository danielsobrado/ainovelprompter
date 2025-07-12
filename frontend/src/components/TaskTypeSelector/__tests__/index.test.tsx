import { describe, it, expect, vi, beforeEach } from 'vitest'
import { render, screen, fireEvent } from '@testing-library/react'
import { TaskTypeSelector } from '../index'
import { mockTaskTypes, createMockEvent } from '../../test/mocks'

// Mock the edit modal
vi.mock('./TaskTypeEditModal', () => ({
  TaskTypeEditModal: ({ isOpen, onClose }: any) => 
    isOpen ? <div data-testid="task-edit-modal">Task Edit Modal<button onClick={onClose}>Close</button></div> : null
}))

describe('TaskTypeSelector', () => {
  const defaultProps = {
    options: mockTaskTypes,
    selectedOptions: [],
    onSelectionChange: vi.fn(),
    onOptionsChange: vi.fn()
  }

  beforeEach(() => {
    vi.clearAllMocks()
  })

  it('should render component with title', () => {
    render(<TaskTypeSelector {...defaultProps} />)
    expect(screen.getByText('Task Types')).toBeInTheDocument()
  })

  it('should display task type options', () => {
    render(<TaskTypeSelector {...defaultProps} />)
    expect(screen.getByText('Continue Story')).toBeInTheDocument()
    expect(screen.getByText('Character Development')).toBeInTheDocument()
  })

  it('should handle task type selection', () => {
    const onSelectionChange = vi.fn()
    render(<TaskTypeSelector {...defaultProps} onSelectionChange={onSelectionChange} />)
    
    const checkbox = screen.getByRole('checkbox', { name: /continue story/i })
    fireEvent.click(checkbox)
    
    expect(onSelectionChange).toHaveBeenCalledWith([mockTaskTypes[0]])
  })

  it('should show selected task types as checked', () => {
    render(<TaskTypeSelector {...defaultProps} selectedOptions={[mockTaskTypes[0]]} />)
    
    const checkbox = screen.getByRole('checkbox', { name: /continue story/i })
    expect(checkbox).toBeChecked()
  })

  it('should open edit modal when edit button is clicked', async () => {
    render(<TaskTypeSelector {...defaultProps} />)
    
    const editButton = screen.getByText('Edit Task Types')
    fireEvent.click(editButton)
    
    expect(screen.getByTestId('task-edit-modal')).toBeInTheDocument()
  })

  it('should display task type descriptions', () => {
    render(<TaskTypeSelector {...defaultProps} />)
    expect(screen.getByText(/write the next chapter/i)).toBeInTheDocument()
    expect(screen.getByText(/focus on developing character/i)).toBeInTheDocument()
  })

  it('should handle multiple selections', () => {
    const onSelectionChange = vi.fn()
    render(<TaskTypeSelector {...defaultProps} onSelectionChange={onSelectionChange} />)
    
    const checkbox1 = screen.getByRole('checkbox', { name: /continue story/i })
    const checkbox2 = screen.getByRole('checkbox', { name: /character development/i })
    
    fireEvent.click(checkbox1)
    expect(onSelectionChange).toHaveBeenCalledWith([mockTaskTypes[0]])
    
    // Simulate having first item selected
    render(<TaskTypeSelector 
      {...defaultProps} 
      selectedOptions={[mockTaskTypes[0]]} 
      onSelectionChange={onSelectionChange} 
    />)
    
    const checkbox2Updated = screen.getByRole('checkbox', { name: /character development/i })
    fireEvent.click(checkbox2Updated)
    expect(onSelectionChange).toHaveBeenCalledWith([mockTaskTypes[0], mockTaskTypes[1]])
  })

  it('should handle deselection', () => {
    const onSelectionChange = vi.fn()
    render(<TaskTypeSelector 
      {...defaultProps} 
      selectedOptions={[mockTaskTypes[0]]} 
      onSelectionChange={onSelectionChange} 
    />)
    
    const checkbox = screen.getByRole('checkbox', { name: /continue story/i })
    fireEvent.click(checkbox)
    
    expect(onSelectionChange).toHaveBeenCalledWith([])
  })
})
