import { useState, useCallback, useEffect } from 'react';
import { BaseOption } from '../types';
import { generateUniqueId } from '../utils/helpers';

interface UseOptionManagementProps<T extends BaseOption> {
  initialOptions?: T[];
  storageKey?: string; // Optional for backwards compatibility
  readFile?: () => Promise<string>;
  writeFile?: (content: string) => Promise<void>;
}

export function useOptionManagement<T extends BaseOption>({ 
  initialOptions = [],
  storageKey,
  readFile,
  writeFile,
}: UseOptionManagementProps<T>) {
  const [options, setOptions] = useState<T[]>(initialOptions);
  const [isLoading, setIsLoading] = useState(false);
  const [error, setError] = useState<string | null>(null);
  
  // Load options on mount
  useEffect(() => {
    const loadOptions = async () => {
      if (readFile) {
        setIsLoading(true);
        setError(null);
        try {
          const content = await readFile();
          const parsedOptions = JSON.parse(content || '[]');
          setOptions(parsedOptions);
        } catch (err) {
          setError('Failed to load options');
          console.error('Failed to load options:', err);
        } finally {
          setIsLoading(false);
        }
      } else if (storageKey) {
        // Fallback to localStorage
        const savedOptions = localStorage.getItem(storageKey);
        if (savedOptions) {
          try {
            setOptions(JSON.parse(savedOptions));
          } catch (err) {
            console.error('Failed to parse stored options:', err);
          }
        }
      }
    };
    
    loadOptions();
  }, [readFile, storageKey]);
  
  const [selectedValues, setSelectedValues] = useState<string[]>([]);
  const [isEditModalOpen, setIsEditModalOpen] = useState(false);

  // Save options function
  const saveOptions = useCallback(async (newOptions: T[]) => {
    if (writeFile) {
      try {
        await writeFile(JSON.stringify(newOptions, null, 2));
        setOptions(newOptions);
      } catch (err) {
        console.error('Failed to save options:', err);
        throw err;
      }
    } else if (storageKey) {
      // Fallback to localStorage
      localStorage.setItem(storageKey, JSON.stringify(newOptions));
      setOptions(newOptions);
    }
  }, [writeFile, storageKey]);

  const addOption = useCallback(async (newOption: Omit<T, 'id'>) => {
    const updatedOptions = [...options, { ...newOption, id: generateUniqueId() } as T];
    await saveOptions(updatedOptions);
  }, [options, saveOptions]);

  const updateOption = useCallback(async (id: string, updatedOption: Partial<T>) => {
    const updatedOptions = options.map(option => 
      option.id === id ? { ...option, ...updatedOption } : option
    );
    await saveOptions(updatedOptions);
  }, [options, saveOptions]);

  const deleteOption = useCallback(async (id: string) => {
    const updatedOptions = options.filter(option => option.id !== id);
    await saveOptions(updatedOptions);
    setSelectedValues(prev => prev.filter(value => value !== id));
  }, [options, saveOptions]);

  // Helper to get selected options
  const selectedOptions = options.filter(option => selectedValues.includes(option.id));
  
  const setSelectedOptions = useCallback((newSelectedOptions: T[]) => {
    setSelectedValues(newSelectedOptions.map(option => option.id));
  }, []);

  return {
    options,
    setOptions,
    selectedValues,
    setSelectedValues,
    selectedOptions,
    setSelectedOptions,
    isEditModalOpen,
    setIsEditModalOpen,
    isLoading,
    error,
    addOption,
    updateOption,
    deleteOption,
    saveOptions
  };
}