/**
 * Storage utility functions for managing application data persistence
 */

import type { 
    TaskTypeOption, 
    RuleOption, 
    CharacterOption, 
    LocationOption, 
    CodexOption,
    SampleChapterOption 
  } from '../types';
  
  // Constants
  export const STORAGE_VERSION = '1.0';
  export const STORAGE_PREFIX = 'ai-novel-prompter';
  
  export const STORAGE_KEYS = {
    TASK_TYPES: 'task-types',
    RULES: 'rules',
    CHARACTERS: 'characters',
    LOCATIONS: 'locations',
    CODEX: 'codex',
    SAMPLE_CHAPTERS: 'sample-chapters',
    SETTINGS: 'settings',
  } as const;
  
  /**
   * Generates a versioned storage key with prefix
   */
  export function getStorageKey(key: string): string {
    return `${STORAGE_PREFIX}-${key}-v${STORAGE_VERSION}`;
  }
  
  /**
   * Clears all application data from localStorage
   */
  export function clearAllStoredData(): void {
    Object.keys(localStorage).forEach(key => {
      if (key.startsWith(STORAGE_PREFIX)) {
        localStorage.removeItem(key);
      }
    });
  }
  
  /**
   * Safely parses JSON with type checking and error handling
   */
  export function safelyParseJSON<T>(json: string | null, fallback: T): T {
    if (!json) return fallback;
    try {
      return JSON.parse(json) as T;
    } catch (e) {
      console.error('Error parsing stored data:', e);
      return fallback;
    }
  }
  
  /**
   * Type-safe storage operations for different data types
   */
  export const storage = {
    taskTypes: {
      save: (data: TaskTypeOption[]) => 
        localStorage.setItem(getStorageKey(STORAGE_KEYS.TASK_TYPES), JSON.stringify(data)),
      load: (): TaskTypeOption[] => 
        safelyParseJSON(localStorage.getItem(getStorageKey(STORAGE_KEYS.TASK_TYPES)), []),
    },
    rules: {
      save: (data: RuleOption[]) => 
        localStorage.setItem(getStorageKey(STORAGE_KEYS.RULES), JSON.stringify(data)),
      load: (): RuleOption[] => 
        safelyParseJSON(localStorage.getItem(getStorageKey(STORAGE_KEYS.RULES)), []),
    },
    characters: {
      save: (data: CharacterOption[]) => 
        localStorage.setItem(getStorageKey(STORAGE_KEYS.CHARACTERS), JSON.stringify(data)),
      load: (): CharacterOption[] => 
        safelyParseJSON(localStorage.getItem(getStorageKey(STORAGE_KEYS.CHARACTERS)), []),
    },
    locations: {
      save: (data: LocationOption[]) => 
        localStorage.setItem(getStorageKey(STORAGE_KEYS.LOCATIONS), JSON.stringify(data)),
      load: (): LocationOption[] => 
        safelyParseJSON(localStorage.getItem(getStorageKey(STORAGE_KEYS.LOCATIONS)), []),
    },
    codex: {
      save: (data: CodexOption[]) => 
        localStorage.setItem(getStorageKey(STORAGE_KEYS.CODEX), JSON.stringify(data)),
      load: (): CodexOption[] => 
        safelyParseJSON(localStorage.getItem(getStorageKey(STORAGE_KEYS.CODEX)), []),
    },
    sampleChapters: {
      save: (data: SampleChapterOption[]) => 
        localStorage.setItem(getStorageKey(STORAGE_KEYS.SAMPLE_CHAPTERS), JSON.stringify(data)),
      load: (): SampleChapterOption[] => 
        safelyParseJSON(localStorage.getItem(getStorageKey(STORAGE_KEYS.SAMPLE_CHAPTERS)), []),
    },
  };
  
  /**
   * Exports all stored data for backup purposes
   */
  export function exportStoredData(): string {
    const data = {
      version: STORAGE_VERSION,
      timestamp: new Date().toISOString(),
      data: {
        taskTypes: storage.taskTypes.load(),
        rules: storage.rules.load(),
        characters: storage.characters.load(),
        locations: storage.locations.load(),
        codex: storage.codex.load(),
        sampleChapters: storage.sampleChapters.load(),
      }
    };
    return JSON.stringify(data, null, 2);
  }
  
  /**
   * Imports data from a backup
   * @returns success status and error message if applicable
   */
  export function importStoredData(jsonString: string): { success: boolean; error?: string } {
    try {
      const data = JSON.parse(jsonString);
      
      // Version check
      if (data.version !== STORAGE_VERSION) {
        return { 
          success: false, 
          error: `Version mismatch. Expected ${STORAGE_VERSION}, got ${data.version}` 
        };
      }
  
      // Import each data type
      storage.taskTypes.save(data.data.taskTypes ?? []);
      storage.rules.save(data.data.rules ?? []);
      storage.characters.save(data.data.characters ?? []);
      storage.locations.save(data.data.locations ?? []);
      storage.codex.save(data.data.codex ?? []);
      storage.sampleChapters.save(data.data.sampleChapters ?? []);
  
      return { success: true };
    } catch (e) {
      return { 
        success: false, 
        error: e instanceof Error ? e.message : 'Unknown error during import' 
      };
    }
  }
  
  /**
   * Gets the total size of stored data in bytes
   */
  export function getStorageSize(): number {
    let size = 0;
    Object.keys(localStorage).forEach(key => {
      if (key.startsWith(STORAGE_PREFIX)) {
        size += localStorage.getItem(key)?.length ?? 0;
      }
    });
    return size;
  }
  
  export default {
    getStorageKey,
    clearAllStoredData,
    safelyParseJSON,
    storage,
    exportStoredData,
    importStoredData,
    getStorageSize,
    STORAGE_VERSION,
    STORAGE_PREFIX,
    STORAGE_KEYS,
  };