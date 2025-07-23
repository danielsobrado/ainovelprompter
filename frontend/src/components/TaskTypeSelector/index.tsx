import React from 'react';
import {
  Select,
  SelectContent,
  SelectItem,
  SelectTrigger,
  SelectValue,
} from '@/components/ui/select';
import { Label } from '@/components/ui/label';
import { Checkbox } from '@/components/ui/checkbox';
import { Button } from '@/components/ui/button';
import { Edit } from 'lucide-react';
import { TaskTypeOption } from '../../types';
import { CheckedState } from '@radix-ui/react-checkbox';
import RefreshButton from '../RefreshButton';

interface TaskTypeSelectorProps {
  value: string;
  onChange: (value: string) => void;
  checked: boolean;
  onCheckedChange: (checked: CheckedState) => void;
  onEditClick: () => void;
  options: TaskTypeOption[];
}

export default function TaskTypeSelector({
  value,
  onChange,
  checked,
  onCheckedChange,
  onEditClick,
  options,
}: TaskTypeSelectorProps) {
  return (
    <div className="flex items-center space-x-2">
      <Checkbox
        id="task-type-checkbox"
        checked={checked}
        onCheckedChange={onCheckedChange}
        disabled={!value} // Disable checkbox if no task type is selected
      />
      <Label htmlFor="task-type" className="whitespace-nowrap">
        Task Type
      </Label>
      <Select 
        value={value} 
        onValueChange={onChange}
      >
        <SelectTrigger id="task-type" className="w-full">
          <SelectValue placeholder="Select task type" />
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
      <RefreshButton 
        type="taskTypes" 
        size="icon" 
        variant="ghost"
      />
      <Button variant="ghost" size="icon" onClick={onEditClick}>
        <Edit className="h-4 w-4" />
      </Button>
    </div>
  );
}