import React, { useState, useEffect } from 'react';
import {
  Dialog,
  DialogContent,
  DialogHeader,
  DialogTitle,
  DialogFooter,
} from '@/components/ui/dialog';
import { Button } from '@/components/ui/button';
import { Input } from '@/components/ui/input';
import { Textarea } from '@/components/ui/textarea';
import { Select, SelectTrigger, SelectContent, SelectItem, SelectValue } from '@/components/ui/select';

interface RulesOption {
  id: string;
  label: string;
  description: string;
}

interface RulessEditModalProps {
  isOpen: boolean;
  onClose: () => void;
  options: RulesOption[];
  onSave: (options: RulesOption[]) => void;
}

export default function RulessEditModal({
  isOpen,
  onClose,
  options,
  onSave,
}: RulessEditModalProps) {
  // State management
  const [localOptions, setLocalOptions] = useState<RulesOption[]>([]);
  const [selectedItemId, setSelectedItemId] = useState<string>('new');
  const [label, setLabel] = useState<string>('');
  const [description, setDescription] = useState<string>('');
  const [error, setError] = useState<string>('');

  // Initialize local options when modal opens
  useEffect(() => {
    setLocalOptions(options);
  }, [options]);

  // Handle selection changes
  useEffect(() => {
    setError('');
    if (selectedItemId === 'new') {
      setLabel('');
      setDescription('');
    } else {
      const selectedItem = localOptions.find((option) => option.id === selectedItemId);
      if (selectedItem) {
        setLabel(selectedItem.label);
        setDescription(selectedItem.description);
      }
    }
  }, [selectedItemId, localOptions]);

  const handleSave = () => {
    if (label.trim() === '') {
      setError('Label is required');
      return;
    }

    if (description.trim() === '') {
      setError('Description is required');
      return;
    }

    setError('');

    if (selectedItemId === 'new') {
      // Add new item
      const newOption: RulesOption = {
        id: crypto.randomUUID(),
        label: label.trim(),
        description: description.trim(),
      };
      setLocalOptions((prev) => [...prev, newOption]);
    } else {
      // Update existing item
      setLocalOptions((prev) =>
        prev.map((option) =>
          option.id === selectedItemId
            ? { ...option, label: label.trim(), description: description.trim() }
            : option
        )
      );
    }

    // Reset form
    setSelectedItemId('new');
    setLabel('');
    setDescription('');
  };

  const handleDelete = () => {
    if (selectedItemId !== 'new') {
      setLocalOptions((prev) => prev.filter((option) => option.id !== selectedItemId));
      setSelectedItemId('new');
      setLabel('');
      setDescription('');
    }
  };

  const handleModalSave = () => {
    onSave(localOptions);
    onClose();
  };

  return (
    <Dialog open={isOpen} onOpenChange={onClose}>
      <DialogContent className="sm:max-w-[600px]">
        <DialogHeader>
          <DialogTitle className="text-xl font-semibold">Edit Custom Instructions</DialogTitle>
        </DialogHeader>

        <div className="space-y-6 py-4">
          {error && (
            <div className="rounded-lg border border-red-500 p-4 text-red-500 text-sm">
              {error}
            </div>
          )}

          <div className="space-y-2">
            <div className="text-sm font-medium">Select Instruction</div>
            <Select 
              value={selectedItemId} 
              onValueChange={setSelectedItemId}
            >
              <SelectTrigger className="w-full">
                <SelectValue placeholder="Select or create new instruction" />
              </SelectTrigger>
              <SelectContent>
                <SelectItem value="new">Create New Custom Instruction</SelectItem>
                {localOptions.map((option) => (
                  <SelectItem key={option.id} value={option.id}>
                    {option.label}
                  </SelectItem>
                ))}
              </SelectContent>
            </Select>
          </div>

          <div className="space-y-2">
            <div className="text-sm font-medium">Label</div>
            <Input
              placeholder="Enter instruction label"
              value={label}
              onChange={(e) => setLabel(e.target.value)}
            />
          </div>

          <div className="space-y-2">
            <div className="text-sm font-medium">Description</div>
            <Textarea
              placeholder="Enter instruction description"
              value={description}
              onChange={(e) => setDescription(e.target.value)}
              className="min-h-[100px]"
            />
          </div>

          <div className="flex space-x-2">
            <Button 
              onClick={handleSave}
              className="flex-1"
            >
              {selectedItemId === 'new' ? 'Add Custom Instruction' : 'Update Instruction'}
            </Button>
            {selectedItemId !== 'new' && (
              <Button 
                variant="destructive" 
                onClick={handleDelete}
                className="flex-1"
              >
                Delete Instruction
              </Button>
            )}
          </div>
        </div>

        <DialogFooter className="gap-2 sm:gap-0">
          <Button variant="outline" onClick={onClose}>
            Cancel
          </Button>
          <Button onClick={handleModalSave}>
            Save All Changes
          </Button>
        </DialogFooter>
      </DialogContent>
    </Dialog>
  );
}