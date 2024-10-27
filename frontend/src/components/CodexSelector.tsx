// CodexSelector.tsx

import React from 'react';
import { Label } from '@/components/ui/label';
import Select from 'react-select';
import { Button } from '@/components/ui/button';

export interface CodexOption {
  id: string;
  label: string;
  description: string;
}

interface CodexSelectorProps {
  values: string[];
  onChange: (selected: string[]) => void;
  options: CodexOption[];
  onEditClick: () => void;
}

export default function CodexSelector({
  values,
  onChange,
  options,
  onEditClick,
}: CodexSelectorProps) {
  const handleChange = (selectedOptions: any) => {
    const selectedValues = selectedOptions ? selectedOptions.map((opt: any) => opt.label) : [];
    onChange(selectedValues);
  };

  return (
    <div className="space-y-2">
      <Label>Codex</Label>
      <Select
        isMulti
        options={options}
        getOptionLabel={(option) => option.label}
        getOptionValue={(option) => option.id}
        value={options.filter((opt) => values.includes(opt.label))}
        onChange={handleChange}
      />
      <Button variant="outline" size="sm" onClick={onEditClick}>
        Edit
      </Button>
    </div>
  );
}
