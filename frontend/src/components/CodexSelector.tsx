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
import { CodexOption } from '../types';

interface CodexSelectorProps {
  values: string[];
  onChange: (values: string[]) => void;
  onEditClick: () => void;
  options: CodexOption[];
}

export default function CodexSelector({
  values,
  onChange,
  onEditClick,
  options,
}: CodexSelectorProps) {
  return (
    <div className="flex items-center space-x-2">
      <Label htmlFor="codex" className="whitespace-nowrap">
        Codex
      </Label>
      <Select
        value="" // Empty string since multiple selections are allowed
        onValueChange={(value) => {
          const isSelected = values.includes(value);
          if (isSelected) {
            onChange(values.filter((v) => v !== value));
          } else {
            onChange([...values, value]);
          }
        }}
      >
        <SelectTrigger id="codex" className="w-full">
          <SelectValue placeholder="Select codex entries" />
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