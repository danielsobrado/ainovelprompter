// SampleChaptersSelector.tsx

import React from 'react';
import { Label } from '@/components/ui/label';
import { Select, SelectTrigger, SelectValue, SelectContent, SelectItem } from '@/components/ui/select';
import { Button } from '@/components/ui/button';
import { Edit } from 'lucide-react';

interface SampleChapterOption {
  id: string;
  label: string;
  content: string;
}

interface SampleChaptersSelectorProps {
  value: string;
  onChange: (value: string) => void;
  options: SampleChapterOption[];
  onEditClick: () => void;
}

export default function SampleChaptersSelector({
  value,
  onChange,
  options,
  onEditClick,
}: SampleChaptersSelectorProps) {
  return (
    <div className="flex items-center space-x-2">
      <Label htmlFor="task-type" className="whitespace-nowrap">
        Sample Chapters
      </Label>
      <div className="flex items-center space-x-2">
        <Select value={value} onValueChange={onChange}>
          <SelectTrigger className="w-[200px]">
            <SelectValue placeholder="Select a sample chapter" />
          </SelectTrigger>
          <SelectContent>
            {options.map((option) => (
              <SelectItem key={option.id} value={option.label}>
                {option.label}
              </SelectItem>
            ))}
          </SelectContent>
        </Select>
        <Button variant="ghost" size="icon" onClick={onEditClick}>
            <Edit className="h-4 w-4" />
        </Button>
      </div>
    </div>
  );
}
