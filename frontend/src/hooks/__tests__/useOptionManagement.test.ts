import { describe, it, expect, vi, beforeEach } from 'vitest'
import { render, screen, fireEvent, waitFor } from '@testing-library/react'
import { useOptionManagement } from '../useOptionManagement'
import { renderHook, act } from '@testing-library/react'
import { mockCharacters, mockWailsSuccess, mockWailsError } from '../../test/mocks'

describe('useOptionManagement', () => {
  beforeEach(() => {
    vi.clearAllMocks()
    // Reset Wails mocks
    window.wails.go.main.App.ReadCharactersFile.mockResolvedValue(JSON.stringify(mockCharacters))
    window.wails.go.main.App.WriteCharactersFile.mockResolvedValue(true)
  })

  it('should initialize with empty options', () => {
    const { result } = renderHook(() => useOptionManagement({
      readFile: window.wails.go.main.App.ReadCharactersFile,
      writeFile: window.wails.go.main.App.WriteCharactersFile,
      storageKey: 'test-key'
    }))

    expect(result.current.options).toEqual([])
    expect(result.current.selectedOptions).toEqual([])
    expect(result.current.isLoading).toBe(false)
    expect(result.current.error).toBeNull()
  })

  it('should load options on mount', async () => {
    const { result } = renderHook(() => useOptionManagement({
      readFile: window.wails.go.main.App.ReadCharactersFile,
      writeFile: window.wails.go.main.App.WriteCharactersFile,
      storageKey: 'test-key'
    }))

    // Wait for the effect to run
    await waitFor(() => {
      expect(result.current.options).toEqual(mockCharacters)
    })

    expect(window.wails.go.main.App.ReadCharactersFile).toHaveBeenCalled()
  })

  it('should handle save options', async () => {
    const { result } = renderHook(() => useOptionManagement({
      readFile: window.wails.go.main.App.ReadCharactersFile,
      writeFile: window.wails.go.main.App.WriteCharactersFile,
      storageKey: 'test-key'
    }))

    const newOptions = [...mockCharacters, { id: '3', label: 'New Character', description: 'New character description' }]

    await act(async () => {
      await result.current.saveOptions(newOptions)
    })

    expect(window.wails.go.main.App.WriteCharactersFile).toHaveBeenCalledWith(JSON.stringify(newOptions, null, 2))
    expect(result.current.options).toEqual(newOptions)
  })

  it('should handle selection changes', () => {
    const { result } = renderHook(() => useOptionManagement({
      readFile: window.wails.go.main.App.ReadCharactersFile,
      writeFile: window.wails.go.main.App.WriteCharactersFile,
      storageKey: 'test-key'
    }))

    act(() => {
      result.current.setSelectedOptions([mockCharacters[0]])
    })

    expect(result.current.selectedOptions).toEqual([mockCharacters[0]])
  })

  it('should handle load error', async () => {
    window.wails.go.main.App.ReadCharactersFile.mockRejectedValue(new Error('Failed to read'))

    const { result } = renderHook(() => useOptionManagement({
      readFile: window.wails.go.main.App.ReadCharactersFile,
      writeFile: window.wails.go.main.App.WriteCharactersFile,
      storageKey: 'test-key'
    }))

    await waitFor(() => {
      expect(result.current.error).toBe('Failed to load options')
    })
  })

  it('should handle save error', async () => {
    window.wails.go.main.App.WriteCharactersFile.mockRejectedValue(new Error('Failed to write'))

    const { result } = renderHook(() => useOptionManagement({
      readFile: window.wails.go.main.App.ReadCharactersFile,
      writeFile: window.wails.go.main.App.WriteCharactersFile,
      storageKey: 'test-key'
    }))

    await act(async () => {
      try {
        await result.current.saveOptions(mockCharacters)
      } catch (error) {
        expect(error).toBeInstanceOf(Error)
      }
    })
  })

  it('should set loading state correctly', async () => {
    // Make the read operation slow
    window.wails.go.main.App.ReadCharactersFile.mockImplementation(
      () => new Promise(resolve => setTimeout(() => resolve(JSON.stringify(mockCharacters)), 100))
    )

    const { result } = renderHook(() => useOptionManagement({
      readFile: window.wails.go.main.App.ReadCharactersFile,
      writeFile: window.wails.go.main.App.WriteCharactersFile,
      storageKey: 'test-key'
    }))

    // Initially should be loading
    expect(result.current.isLoading).toBe(true)

    // Wait for loading to complete
    await waitFor(() => {
      expect(result.current.isLoading).toBe(false)
    })
  })

  it('should handle malformed JSON gracefully', async () => {
    window.wails.go.main.App.ReadCharactersFile.mockResolvedValue('invalid json')

    const { result } = renderHook(() => useOptionManagement({
      readFile: window.wails.go.main.App.ReadCharactersFile,
      writeFile: window.wails.go.main.App.WriteCharactersFile,
      storageKey: 'test-key'
    }))

    await waitFor(() => {
      expect(result.current.options).toEqual([])
      expect(result.current.error).toBe('Failed to load options')
    })
  })
})