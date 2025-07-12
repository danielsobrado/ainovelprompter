import { describe, it, expect, vi, beforeEach } from 'vitest'
import { render, screen, fireEvent, waitFor } from '@testing-library/react'
import { SettingsModal } from '../SettingsModal'
import { mockSettings } from '../test/mocks'

describe('SettingsModal', () => {
  const defaultProps = {
    isOpen: true,
    onClose: vi.fn(),
    settings: mockSettings,
    onSave: vi.fn()
  }

  beforeEach(() => {
    vi.clearAllMocks()
  })

  it('should render when open', () => {
    render(<SettingsModal {...defaultProps} />)
    expect(screen.getByText('Settings')).toBeInTheDocument()
  })

  it('should not render when closed', () => {
    render(<SettingsModal {...defaultProps} isOpen={false} />)
    expect(screen.queryByText('Settings')).not.toBeInTheDocument()
  })

  it('should display current settings', () => {
    render(<SettingsModal {...defaultProps} />)
    
    // Should show theme setting
    expect(screen.getByDisplayValue('dark')).toBeInTheDocument()
    
    // Should show language setting
    expect(screen.getByDisplayValue('en')).toBeInTheDocument()
    
    // Should show auto-save checkbox
    const autoSaveCheckbox = screen.getByRole('checkbox', { name: /auto.*save/i })
    expect(autoSaveCheckbox).toBeChecked()
  })

  it('should handle theme change', () => {
    render(<SettingsModal {...defaultProps} />)
    
    const themeSelect = screen.getByDisplayValue('dark')
    fireEvent.change(themeSelect, { target: { value: 'light' } })
    
    expect(themeSelect).toHaveValue('light')
  })

  it('should handle language change', () => {
    render(<SettingsModal {...defaultProps} />)
    
    const languageSelect = screen.getByDisplayValue('en')
    fireEvent.change(languageSelect, { target: { value: 'es' } })
    
    expect(languageSelect).toHaveValue('es')
  })

  it('should handle auto-save toggle', () => {
    render(<SettingsModal {...defaultProps} />)
    
    const autoSaveCheckbox = screen.getByRole('checkbox', { name: /auto.*save/i })
    fireEvent.click(autoSaveCheckbox)
    
    expect(autoSaveCheckbox).not.toBeChecked()
  })

  it('should save settings', async () => {
    const onSave = vi.fn()
    render(<SettingsModal {...defaultProps} onSave={onSave} />)
    
    // Change a setting
    const themeSelect = screen.getByDisplayValue('dark')
    fireEvent.change(themeSelect, { target: { value: 'light' } })
    
    // Click save
    const saveButton = screen.getByText('Save')
    fireEvent.click(saveButton)
    
    expect(onSave).toHaveBeenCalledWith(
      expect.objectContaining({
        theme: 'light',
        defaultLanguage: 'en',
        enableAutoSave: true
      })
    )
  })

  it('should close modal when save is clicked', async () => {
    const onClose = vi.fn()
    render(<SettingsModal {...defaultProps} onClose={onClose} />)
    
    const saveButton = screen.getByText('Save')
    fireEvent.click(saveButton)
    
    expect(onClose).toHaveBeenCalled()
  })

  it('should close modal when cancel is clicked', () => {
    const onClose = vi.fn()
    render(<SettingsModal {...defaultProps} onClose={onClose} />)
    
    const cancelButton = screen.getByText('Cancel')
    fireEvent.click(cancelButton)
    
    expect(onClose).toHaveBeenCalled()
  })

  it('should reset settings when cancel is clicked', () => {
    render(<SettingsModal {...defaultProps} />)
    
    // Change a setting
    const themeSelect = screen.getByDisplayValue('dark')
    fireEvent.change(themeSelect, { target: { value: 'light' } })
    
    // Click cancel
    const cancelButton = screen.getByText('Cancel')
    fireEvent.click(cancelButton)
    
    // Re-open modal
    render(<SettingsModal {...defaultProps} />)
    
    // Should revert to original value
    expect(screen.getByDisplayValue('dark')).toBeInTheDocument()
  })

  it('should show available theme options', () => {
    render(<SettingsModal {...defaultProps} />)
    
    const themeSelect = screen.getByDisplayValue('dark')
    fireEvent.click(themeSelect)
    
    expect(screen.getByText('Light')).toBeInTheDocument()
    expect(screen.getByText('Dark')).toBeInTheDocument()
    expect(screen.getByText('System')).toBeInTheDocument()
  })

  it('should show available language options', () => {
    render(<SettingsModal {...defaultProps} />)
    
    const languageSelect = screen.getByDisplayValue('en')
    fireEvent.click(languageSelect)
    
    expect(screen.getByText('English')).toBeInTheDocument()
    expect(screen.getByText('Spanish')).toBeInTheDocument()
    expect(screen.getByText('French')).toBeInTheDocument()
  })
})
