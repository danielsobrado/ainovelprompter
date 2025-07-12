import { describe, it, expect, vi, beforeEach } from 'vitest'
import { render, screen, fireEvent, waitFor } from '@testing-library/react'
import { calculateTokenCount, truncateText, formatDate, isValidEmail, createSlug, generateUniqueId } from '../utils/helpers'

describe('Helper Functions', () => {
  describe('calculateTokenCount', () => {
    it('should calculate approximate token count', () => {
      const text = 'Hello world this is a test'
      const result = calculateTokenCount(text)
      expect(result).toBeGreaterThan(0)
      expect(typeof result).toBe('number')
    })

    it('should handle empty text', () => {
      expect(calculateTokenCount('')).toBe(0)
    })

    it('should handle null and undefined', () => {
      expect(calculateTokenCount(null as any)).toBe(0)
      expect(calculateTokenCount(undefined as any)).toBe(0)
    })
  })

  describe('truncateText', () => {
    it('should truncate text to specified length', () => {
      const text = 'This is a very long text that should be truncated'
      const result = truncateText(text, 20)
      expect(result).toHaveLength(23) // 20 + '...'
      expect(result.endsWith('...')).toBe(true)
    })

    it('should not truncate if text is shorter than limit', () => {
      const text = 'Short text'
      const result = truncateText(text, 20)
      expect(result).toBe(text)
    })

    it('should handle empty text', () => {
      expect(truncateText('', 10)).toBe('')
    })
  })

  describe('formatDate', () => {
    it('should format date correctly', () => {
      const date = new Date('2024-01-15T10:30:00Z')
      const result = formatDate(date)
      expect(result).toMatch(/\d{1,2}\/\d{1,2}\/\d{4}/)
    })

    it('should handle date string input', () => {
      const result = formatDate('2024-01-15')
      expect(result).toMatch(/\d{1,2}\/\d{1,2}\/\d{4}/)
    })
  })

  describe('isValidEmail', () => {
    it('should validate correct email addresses', () => {
      expect(isValidEmail('test@example.com')).toBe(true)
      expect(isValidEmail('user.name+tag@domain.co.uk')).toBe(true)
    })

    it('should reject invalid email addresses', () => {
      expect(isValidEmail('invalid-email')).toBe(false)
      expect(isValidEmail('test@')).toBe(false)
      expect(isValidEmail('@domain.com')).toBe(false)
      expect(isValidEmail('')).toBe(false)
    })
  })

  describe('createSlug', () => {
    it('should create valid slugs', () => {
      expect(createSlug('Hello World')).toBe('hello-world')
      expect(createSlug('Test With Special Characters!')).toBe('test-with-special-characters')
      expect(createSlug('Multiple   Spaces')).toBe('multiple-spaces')
    })

    it('should handle edge cases', () => {
      expect(createSlug('')).toBe('')
      expect(createSlug('   ')).toBe('')
      expect(createSlug('123')).toBe('123')
    })
  })

  describe('generateUniqueId', () => {
    it('should generate unique IDs', () => {
      const id1 = generateUniqueId()
      const id2 = generateUniqueId()
      expect(id1).not.toBe(id2)
      expect(typeof id1).toBe('string')
      expect(id1.length).toBeGreaterThan(0)
    })
  })
})
