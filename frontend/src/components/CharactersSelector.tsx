import React from 'react';
import { Label } from '@/components/ui/label';
import { Button } from '@/components/ui/button';
import { Edit } from 'lucide-react';
import { CharacterOption } from '../types';
import { FancyMultiSelect } from '@/components/ui/fancy-multi-select';

interface CharactersSelectorProps {
  values: string[];
  onChange: (values: string[]) => void;
  onEditClick: () => void;
  options: CharacterOption[];
}

export default function CharactersSelector({
  values,
  onChange,
  onEditClick,
  options,
}: CharactersSelectorProps) {
  const selectOptions = options.map(option => ({
    label: option.label,
    value: option.label
  }));

  return (
    <div className="flex items-center space-x-2">
      <Label htmlFor="characters" className="whitespace-nowrap">
        Characters
      </Label>
      <div className="flex-1">
        <FancyMultiSelect
          options={selectOptions}
          selected={values}
          onChange={onChange}
          placeholder="Select characters..."
        />
      </div>
      <Button variant="ghost" size="icon" onClick={onEditClick}>
        <Edit className="h-4 w-4" />
      </Button>
    </div>
  );
}