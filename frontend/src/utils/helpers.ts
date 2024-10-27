// utils/helpers.ts

/**
 * Generates a unique ID using a combination of timestamp and random numbers
 * More reliable than Math.random() and shorter than UUID
 */
export const generateUniqueId = (): string => {
    return `${Date.now().toString(36)}-${Math.random().toString(36).substring(2, 9)}`;
  };
  
  /**
   * Safely parses JSON with a default value
   */
  export function safeJSONParse<T>(str: string, defaultValue: T): T {
    try {
      return JSON.parse(str) as T;
    } catch {
      return defaultValue;
    }
  }
  
  /**
   * Truncates text to a specified length
   */
  export function truncateText(text: string, maxLength: number): string {
    if (text.length <= maxLength) return text;
    return text.slice(0, maxLength) + '...';
  }
  
  /**
   * Debounces a function
   */
  export function debounce<T extends (...args: any[]) => any>(
    fn: T,
    delay: number
  ): (...args: Parameters<T>) => void {
    let timeoutId: NodeJS.Timeout;
    return (...args: Parameters<T>) => {
      clearTimeout(timeoutId);
      timeoutId = setTimeout(() => fn(...args), delay);
    };
  }
  
  /**
   * Formats a date to a readable string
   */
  export function formatDate(date: Date): string {
    return new Intl.DateTimeFormat('en-US', {
      year: 'numeric',
      month: 'short',
      day: 'numeric',
      hour: '2-digit',
      minute: '2-digit',
    }).format(date);
  }
  
  /**
   * Validates an email address
   */
  export function isValidEmail(email: string): boolean {
    const emailRegex = /^[^\s@]+@[^\s@]+\.[^\s@]+$/;
    return emailRegex.test(email);
  }
  
  /**
   * Calculates approximate token count for a text string
   */
  export function calculateTokenCount(text: string): number {
    // This is a simple approximation. For production, use a proper tokenizer.
    return text.trim().split(/\s+/).length;
  }
  
  /**
   * Sanitizes HTML strings
   */
  export function sanitizeHTML(html: string): string {
    return html
      .replace(/&/g, '&amp;')
      .replace(/</g, '&lt;')
      .replace(/>/g, '&gt;')
      .replace(/"/g, '&quot;')
      .replace(/'/g, '&#039;');
  }
  
  /**
   * Creates a slug from a string
   */
  export function createSlug(text: string): string {
    return text
      .toLowerCase()
      .replace(/[^\w\s-]/g, '')
      .replace(/\s+/g, '-')
      .replace(/-+/g, '-')
      .trim();
  }
  
  /**
   * Deep clones an object
   */
  export function deepClone<T>(obj: T): T {
    return JSON.parse(JSON.stringify(obj));
  }
  
  /**
   * Groups an array of objects by a key
   */
  export function groupBy<T>(array: T[], key: keyof T): Record<string, T[]> {
    return array.reduce((groups, item) => {
      const groupKey = String(item[key]);
      return {
        ...groups,
        [groupKey]: [...(groups[groupKey] || []), item],
      };
    }, {} as Record<string, T[]>);
  }
  
  /**
   * Retries a function with exponential backoff
   */
  export async function retry<T>(
    fn: () => Promise<T>,
    maxAttempts: number = 3,
    baseDelay: number = 1000
  ): Promise<T> {
    let lastError: Error;
    
    for (let attempt = 1; attempt <= maxAttempts; attempt++) {
      try {
        return await fn();
      } catch (error) {
        lastError = error as Error;
        if (attempt === maxAttempts) break;
        
        const delay = baseDelay * Math.pow(2, attempt - 1);
        await new Promise(resolve => setTimeout(resolve, delay));
      }
    }
    
    throw lastError!;
  }
  
  export default {
    generateUniqueId,
    safeJSONParse,
    truncateText,
    debounce,
    formatDate,
    isValidEmail,
    calculateTokenCount,
    sanitizeHTML,
    createSlug,
    deepClone,
    groupBy,
    retry,
  };