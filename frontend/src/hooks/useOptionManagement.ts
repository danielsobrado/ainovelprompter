import { useState, useCallback } from 'react';
import { BaseOption } from '../types';
import { generateUniqueId } from '../utils/helpers';

interface UseOptionManagementProps<T extends BaseOption> {
  initialOptions: T[];
}

export function useOptionManagement<T extends BaseOption>({ 
  initialOptions = [] 
}: UseOptionManagementProps<T>) {
  const [options, setOptions] = useState<T[]>(initialOptions);
  const [selectedValues, setSelectedValues] = useState<string[]>([]);
  const [isEditModalOpen, setIsEditModalOpen] = useState(false);

  const addOption = useCallback((newOption: Omit<T, 'id'>) => {
    setOptions(prev => [...prev, { ...newOption, id: generateUniqueId() } as T]);
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