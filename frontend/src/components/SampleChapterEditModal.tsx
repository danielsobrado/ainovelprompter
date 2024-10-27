// SampleChapterEditModal.tsx

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
import {
  Select,
  SelectTrigger,
  SelectContent,
  SelectItem,
  SelectValue,
} from '@/components/ui/select';
import { SampleChapterOption } from '@/types';

interface SampleChapterEditModalProps {
  isOpen: boolean;
  onClose: () => void;
  options: SampleChapterOption[];
  onSave: (options: SampleChapterOption[]) => void;
}

export default function SampleChapterEditModal({
  isOpen,
  onClose,
  options,
  onSave,
}: SampleChapterEditModalProps) {
  // State management
  const [localOptions, setLocalOptions] = useState<SampleChapterOption[]>([]);
  const [selectedItemId, setSelectedItemId] = useState<string>('new');
  const [label, setLabel] = useState<string>('');
  const [content, setContent] = useState<string>('');
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
      setContent('');
    } else {
      const selectedItem = localOptions.find((option) => option.id === selectedItemId);
      if (selectedItem) {
        setLabel(selectedItem.label);
        setContent(selectedItem.content);
      }
    }
  }, [selectedItemId, localOptions]);

  const handleSave = () => {
    if (label.trim() === '') {
      setError('Label is required');
      return;
    }

    if (content.trim() === '') {
      setError('Content is required');
      return;
    }

    setError('');

    if (selectedItemId === 'new') {
      // Add new item
      const newOption: SampleChapterOption = {
        id: crypto.randomUUID(),
        label: label.trim(),
        content: content.trim(),
      };
      setLocalOptions((prev) => [...prev, newOption]);
    } else {
      // Update existing item
      setLocalOptions((prev) =>
        prev.map((option) =>
          option.id === selectedItemId
            ? { ...option, label: label.trim(), content: content.trim() }
            : option
        )
      );
    }

    // Reset form
    setSelectedItemId('new');
    setLabel('');
    setContent('');
  };

  const handleDelete = () => {
    if (selectedItemId !== 'new') {
      setLocalOptions((prev) => prev.filter((option) => option.id !== selectedItemId));
      setSelectedItemId('new');
      setLabel('');
      setContent('');
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
          <DialogTitle className="text-xl font-semibold">Edit Sample Chapters</DialogTitle>
        </DialogHeader>

        <div className="space-y-6 py-4">
          {error && (
            <div className="rounded-lg border border-red-500 p-4 text-red-500 text-sm">
              {error}
            </div>
          )}

          <div className="space-y-2">
            <div className="text-sm font-medium">Select Sample Chapter</div>
            <Select value={selectedItemId} onValueChange={setSelectedItemId}>
              <SelectTrigger className="w-full">
                <SelectValue placeholder="Select or create new sample chapter" />
              </SelectTrigger>
              <SelectContent>
                <SelectItem value="new">Create New Sample Chapter</SelectItem>
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
              placeholder="Enter chapter label"
              value={label}
              onChange={(e) => setLabel(e.target.value)}
            />
          </div>

          <div className="space-y-2">
            <div className="text-sm font-medium">Content</div>
            <Textarea
              placeholder="Enter chapter content"
              value={content}
              onChange={(e) => setContent(e.target.value)}
              className="min-h-[150px]"
            />
          </div>

          <div className="flex space-x-2">
            <Button onClick={handleSave} className="flex-1">
              {selectedItemId === 'new' ? 'Add Chapter' : 'Update Chapter'}
            </Button>
            {selectedItemId !== 'new' && (
              <Button variant="destructive" onClick={handleDelete} className="flex-1">
                Delete Chapter
              </Button>
            )}
          </div>
        </div>

        <DialogFooter className="gap-2 sm:gap-0">
          <Button variant="outline" onClick={onClose}>
            Cancel
          </Button>
          <Button onClick={handleModalSave}>Save All Changes</Button>
        </DialogFooter>
      </DialogContent>
    </Dialog>
  );
}
