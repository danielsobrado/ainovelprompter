// frontend/src/components/ProseImprovement/PromptManager.tsx
import React, { useState } from 'react';
import { Button } from '@/components/ui/button';
import { Card } from '@/components/ui/card';
import { Input } from '@/components/ui/input';
import { Textarea } from '@/components/ui/textarea';
import { ScrollArea } from '@/components/ui/scroll-area';
import { Badge } from '@/components/ui/badge';
import {
  Select,
  SelectContent,
  SelectItem,
  SelectTrigger,
  SelectValue,
} from '@/components/ui/select';
import { Plus, Edit2, Trash2, MoveUp, MoveDown, RotateCcw } from 'lucide-react';
import type { ProseImprovementPrompt } from '@/types';
import { DEFAULT_PROSE_IMPROVEMENT_PROMPTS } from '@/utils/constants';

interface PromptManagerProps {
  prompts: ProseImprovementPrompt[];
  onPromptsChange: (prompts: ProseImprovementPrompt[]) => void;
}

export function PromptManager({ prompts, onPromptsChange }: PromptManagerProps) {
  const [editingId, setEditingId] = useState<string | null>(null);
  const [newPrompt, setNewPrompt] = useState<Partial<ProseImprovementPrompt>>({
    label: '',
    defaultPromptText: '',
    category: 'custom',
    variants: []
  });

  const handleAdd = () => {
    if (!newPrompt.label || !newPrompt.defaultPromptText) return;
    
    const prompt: ProseImprovementPrompt = {
      id: crypto.randomUUID(),
      label: newPrompt.label,
      defaultPromptText: newPrompt.defaultPromptText,
      category: newPrompt.category || 'custom',
      order: prompts.length,
      description: newPrompt.description,
      variants: newPrompt.variants || []
    };
    
    onPromptsChange([...prompts, prompt]);
    setNewPrompt({ label: '', defaultPromptText: '', category: 'custom', variants: [] });
  };

  const handleUpdate = (id: string, updates: Partial<ProseImprovementPrompt>) => {
    onPromptsChange(
      prompts.map(p => p.id === id ? { ...p, ...updates } : p)
    );
    setEditingId(null);
  };

  const handleDelete = (id: string) => {
    onPromptsChange(prompts.filter(p => p.id !== id));
  };
  const handleReorder = (id: string, direction: 'up' | 'down') => {
    const index = prompts.findIndex(p => p.id === id);
    if (index === -1) return;
    
    const newIndex = direction === 'up' ? index - 1 : index + 1;
    if (newIndex < 0 || newIndex >= prompts.length) return;
    
    const newPrompts = [...prompts];
    [newPrompts[index], newPrompts[newIndex]] = [newPrompts[newIndex], newPrompts[index]];
    
    // Update order values
    newPrompts.forEach((p, i) => p.order = i);
    onPromptsChange(newPrompts);
  };

  const handleResetToDefaults = () => {
    if (confirm('Are you sure you want to reset all prompts to their default values? This will replace any custom changes you have made.')) {
      onPromptsChange([...DEFAULT_PROSE_IMPROVEMENT_PROMPTS]);
    }
  };

  return (
    <div className="space-y-4">
      {/* Add new prompt */}
      <Card className="p-4">        <h3 className="font-semibold mb-3">Add New Prompt</h3>
        <div className="space-y-3">
          <Input
            placeholder="Prompt label"
            value={newPrompt.label || ''}
            onChange={e => setNewPrompt({ ...newPrompt, label: e.target.value })}
          />
          <Input
            placeholder="Description (optional)"
            value={newPrompt.description || ''}
            onChange={e => setNewPrompt({ ...newPrompt, description: e.target.value })}
          />
          <Textarea
            placeholder="Default prompt text..."
            value={newPrompt.defaultPromptText || ''}
            onChange={e => setNewPrompt({ ...newPrompt, defaultPromptText: e.target.value })}
            className="min-h-[100px]"
          />
          <div className="flex gap-2">
            <Select
              value={newPrompt.category}
              onValueChange={v => setNewPrompt({ ...newPrompt, category: v as any })}
            >
              <SelectTrigger className="w-[180px]">
                <SelectValue />
              </SelectTrigger>
              <SelectContent>
                <SelectItem value="tropes">AI Tropes</SelectItem>
                <SelectItem value="style">Style</SelectItem>
                <SelectItem value="grammar">Grammar</SelectItem>
                <SelectItem value="custom">Custom</SelectItem>
              </SelectContent>
            </Select>
            <Button onClick={handleAdd} className="ml-auto">
              <Plus className="mr-2 h-4 w-4" />
              Add Prompt
            </Button>
          </div>
        </div>
      </Card>

      {/* Existing prompts */}
      <ScrollArea className="h-[400px]">
        <div className="space-y-2">
          {prompts.sort((a, b) => a.order - b.order).map((prompt, index) => (
            <Card key={prompt.id} className="p-3">              {editingId === prompt.id ? (
                <div className="space-y-2">
                  <Input
                    value={prompt.label}
                    onChange={e => handleUpdate(prompt.id, { label: e.target.value })}
                  />
                  <Input
                    placeholder="Description (optional)"
                    value={prompt.description || ''}
                    onChange={e => handleUpdate(prompt.id, { description: e.target.value })}
                  />
                  <Textarea
                    value={prompt.defaultPromptText}
                    onChange={e => handleUpdate(prompt.id, { defaultPromptText: e.target.value })}
                    className="min-h-[200px]"
                  />
                  <div className="flex justify-end gap-2">
                    <Button size="sm" variant="outline" onClick={() => setEditingId(null)}>
                      Cancel
                    </Button>
                    <Button size="sm" onClick={() => setEditingId(null)}>
                      Save
                    </Button>
                  </div>
                </div>
              ) : (
                <div className="flex items-start justify-between">                  <div className="flex-1">
                    <div className="flex items-center gap-2 mb-1">
                      <span className="font-medium">{prompt.label}</span>
                      <Badge variant="outline" className="text-xs">
                        {prompt.category}
                      </Badge>
                    </div>
                    {prompt.description && (
                      <p className="text-xs text-muted-foreground mb-2">{prompt.description}</p>
                    )}
                    <p className="text-sm text-muted-foreground line-clamp-3">
                      {prompt.defaultPromptText}
                    </p>
                    {prompt.variants && prompt.variants.length > 0 && (
                      <p className="text-xs text-blue-600 mt-1">
                        +{prompt.variants.length} variant{prompt.variants.length !== 1 ? 's' : ''} available
                      </p>
                    )}
                    {prompt.defaultPromptText.length > 150 && (
                      <Button
                        variant="ghost"
                        size="sm"
                        className="mt-1 h-auto p-0 text-xs"
                        onClick={() => setEditingId(prompt.id)}
                      >
                        View full prompt →
                      </Button>
                    )}
                  </div>
                  <div className="flex items-center gap-1 ml-2">
                    <Button
                      size="icon"
                      variant="ghost"
                      onClick={() => handleReorder(prompt.id, 'up')}
                      disabled={index === 0}
                    >
                      <MoveUp className="h-4 w-4" />
                    </Button>
                    <Button
                      size="icon"
                      variant="ghost"
                      onClick={() => handleReorder(prompt.id, 'down')}
                      disabled={index === prompts.length - 1}
                    >
                      <MoveDown className="h-4 w-4" />
                    </Button>
                    <Button
                      size="icon"
                      variant="ghost"
                      onClick={() => setEditingId(prompt.id)}
                    >
                      <Edit2 className="h-4 w-4" />
                    </Button>
                    <Button
                      size="icon"
                      variant="ghost"
                      onClick={() => handleDelete(prompt.id)}
                    >
                      <Trash2 className="h-4 w-4" />
                    </Button>
                  </div>
                </div>
              )}
            </Card>
          ))}
        </div>
      </ScrollArea>
    </div>
  );
}