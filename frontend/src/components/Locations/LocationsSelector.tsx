import React from 'react';
import { Label } from '@/components/ui/label';
import { Button } from '@/components/ui/button';
import { Edit } from 'lucide-react';
import { LocationOption } from '../../types';
import { FancyMultiSelect } from '@/components/ui/fancy-multi-select';
import RefreshButton from '../RefreshButton';

interface LocationsSelectorProps {
  values: string[];
  onChange: (values: string[]) => void;
  onEditClick: () => void;
  options: LocationOption[];
}

export default function LocationsSelector({
  values,
  onChange,
  onEditClick,
  options,
}: LocationsSelectorProps) {
  const selectOptions = options.map(option => ({
    label: option.label,
    value: option.label
  }));

  return (
    <div className="flex items-center space-x-2">
      <Label htmlFor="locations" className="whitespace-nowrap">
        Locations
      </Label>
      <div className="flex-1">
        <FancyMultiSelect
          options={selectOptions}
          selected={values}
          onChange={onChange}
          placeholder="Select locations..."
        />
      </div>
      <RefreshButton 
        type="locations" 
        size="icon" 
        variant="ghost"
      />
      <Button variant="ghost" size="icon" onClick={onEditClick}>
        <Edit className="h-4 w-4" />
      </Button>
    </div>
  );
}