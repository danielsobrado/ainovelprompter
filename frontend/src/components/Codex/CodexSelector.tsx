import React from 'react';
import { Label } from '@/components/ui/label';
import { Button } from '@/components/ui/button';
import { Edit } from 'lucide-react';
import { CodexOption } from '../../types';
import { FancyMultiSelect } from '@/components/ui/fancy-multi-select';

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
  const selectOptions = options.map(option => ({
    label: option.label,
    value: option.label
  }));

  return (
    <div className="flex items-center space-x-2">
      <Label htmlFor="codex" className="whitespace-nowrap">
        Codex
      </Label>
      <div className="flex-1">
        <FancyMultiSelect
          options={selectOptions}
          selected={values}
          onChange={onChange}
          placeholder="Select codex entries..."
        />
      </div>
      <Button variant="ghost" size="icon" onClick={onEditClick}>
        <Edit className="h-4 w-4" />
      </Button>
    </div>
  );
}