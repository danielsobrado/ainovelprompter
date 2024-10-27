// LocationsSelector.tsx

import React from 'react';
import { Label } from '@/components/ui/label';
import Select from 'react-select';
import { Button } from '@/components/ui/button';

export interface LocationOption {
  id: string;
  label: string;
  description: string;
}

interface LocationsSelectorProps {
  values: string[];
  onChange: (values: string[]) => void;
  onEditClick: () => void;
  options: LocationOption[];
}

export default function LocationsSelector({
  values,
  onChange,
  options,
  onEditClick,
}: LocationsSelectorProps) {
  const handleChange = (selectedOptions: any) => {
    const selectedValues = selectedOptions ? selectedOptions.map((opt: any) => opt.label) : [];
    onChange(selectedValues);
  };

  return (
    <div className="space-y-2">
      <Label>Locations</Label>
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
