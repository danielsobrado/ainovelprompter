import { useState, useCallback, useEffect } from 'react';
import { BaseOption } from '../types';
import { generateUniqueId } from '../utils/helpers';

interface UseOptionManagementProps<T extends BaseOption> {
  initialOptions?: T[];
  storageKey: string; // Add this to identify different option types
}

export function useOptionManagement<T extends BaseOption>({ 
  initialOptions = [],
  storageKey,
}: UseOptionManagementProps<T>) {
  // Initialize state from localStorage if available
  const [options, setOptions] = useState<T[]>(() => {
    const savedOptions = localStorage.getItem(storageKey);
    return savedOptions ? JSON.parse(savedOptions) : initialOptions;
  });
  
  const [selectedValues, setSelectedValues] = useState<string[]>([]);
  const [isEditModalOpen, setIsEditModalOpen] = useState(false);

  // Save to localStorage whenever options change
  useEffect(() => {
    localStorage.setItem(storageKey, JSON.stringify(options));
  }, [options, storageKey]);

  const addOption = useCallback((newOption: Omit<T, 'id'>) => {
    setOptions(prev => {
      const updatedOptions = [...prev, { ...newOption, id: generateUniqueId() } as T];
      return updatedOptions;
    });
  }, []);

  const updateOption = useCallback((id: string, updatedOption: Partial<T>) => {
    setOptions(prev => prev.map(option => 
      option.id === id ? { ...option, ...updatedOption } : option
    ));
  }, []);

  const deleteOption = useCallback((id: string) => {
    setOptions(prev => prev.filter(option => option.id !== id));
    setSelectedValues(prev => prev.filter(value => value !== id));
  }, []);

  return {
    options,
    setOptions,
    selectedValues,
    setSelectedValues,
    isEditModalOpen,
    setIsEditModalOpen,
    addOption,
    updateOption,
    deleteOption
  };
}