// SampleChapterSelector.tsx
import React from 'react';
import { Label } from '@/components/ui/label';
import {
  Select,
  SelectTrigger,
  SelectValue,
  SelectContent,
  SelectItem,
} from '@/components/ui/select';
import { Button } from '@/components/ui/button';
import { Edit } from 'lucide-react';
import { SampleChapterOption } from '../types';

interface SampleChapterSelectorProps {
  value: string;
  onChange: (value: string) => void;
  options: SampleChapterOption[];
  onEditClick: () => void;
}

export default function SampleChapterSelector({
  value,
  onChange,
  options,
  onEditClick,
}: SampleChapterSelectorProps) {
  return (
    <div className="flex items-center space-x-2">
      <Label htmlFor="sample-chapter" className="whitespace-nowrap">
        Sample Chapter
      </Label>
      <Select value={value} onValueChange={onChange}>
        <SelectTrigger id="sample-chapter" className="w-full">
          <SelectValue placeholder="Select sample chapter" />
        </SelectTrigger>
        <SelectContent>
          {options.map((option) => (
            <SelectItem key={option.id} value={option.label}>
              <div className="flex items-center">
                <div
                  className={`mr-2 h-2 w-2 rounded-full ${
                    value === option.label ? 'bg-blue-600' : 'bg-gray-300'
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