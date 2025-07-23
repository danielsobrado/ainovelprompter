import React from 'react';
import { Label } from '@/components/ui/label';
import { Button } from '@/components/ui/button';
import { Edit } from 'lucide-react';
import { RuleOption } from '../../types';
import { FancyMultiSelect } from '@/components/ui/fancy-multi-select';
import RefreshButton from '../RefreshButton';

interface RulesSelectorProps {
  values: string[];
  onChange: (values: string[]) => void;
  onEditClick: () => void;
  options: RuleOption[];
}

export default function RulesSelector({
  values,
  onChange,
  onEditClick,
  options,
}: RulesSelectorProps) {
  // Convert RuleOption[] to Option[] for FancyMultiSelect
  const selectOptions = options.map(option => ({
    label: option.label,
    value: option.label // Using label as value since that's what your current code uses
  }));

  return (
    <div className="flex items-center space-x-2">
      <Label htmlFor="rules" className="whitespace-nowrap">
        Rules
      </Label>
      <div className="flex-1">
        <FancyMultiSelect
          options={selectOptions}
          selected={values}
          onChange={onChange}
          placeholder="Select rules..."
        />
      </div>
      <RefreshButton 
        type="rules" 
        size="icon" 
        variant="ghost"
      />
      <Button variant="ghost" size="icon" onClick={onEditClick}>
        <Edit className="h-4 w-4" />
      </Button>
    </div>
  );
}