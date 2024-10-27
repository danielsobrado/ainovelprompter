// CharactersSelector.tsx

import React from 'react';
import {
  Select,
  SelectContent,
  SelectItem,
  SelectTrigger,
  SelectValue,
} from '@/components/ui/select';
import { Label } from '@/components/ui/label';
import { Button } from '@/components/ui/button';
import { Edit } from 'lucide-react';
import { CharacterOption } from '../types';

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
  return (
    <div className="flex items-center space-x-2">
      <Label htmlFor="characters" className="whitespace-nowrap">
        Characters
      </Label>
      <Select
        value=""
        onValueChange={(value) => {
          const isSelected = values.includes(value);
          if (isSelected) {
            onChange(values.filter((v) => v !== value));
          } else {
            onChange([...values, value]);
          }
        }}
      >
        <SelectTrigger id="characters" className="w-full">
          <SelectValue placeholder="Select characters" />
        </SelectTrigger>
        <SelectContent>
          {options.map((option) => (
            <SelectItem key={option.id} value={option.label}>
              <div className="flex items-center">
                <div
                  className={`mr-2 h-2 w-2 rounded-full ${
                    values.includes(option.label) ? 'bg-blue-600' : 'bg-gray-300'
                  }`}
                ></div>
                {option.label}
              </div>
            </SelectItem>
          ))}
        </SelectContent>
      </Select>
      <Button variant="ghost" size="icon" onClick={onEditClick}>
        <Edit className="h-4 w-4" />
      </Button>
    </div>
  );
}
