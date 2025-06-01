// frontend/src/components/ProseImprovement/ProviderSettings.tsx
import React from 'react';
import {
  Dialog,
  DialogContent,
  DialogHeader,
  DialogTitle,
} from '@/components/ui/dialog';
import { Input } from '@/components/ui/input';
import { Label } from '@/components/ui/label';
import { RadioGroup, RadioGroupItem } from '@/components/ui/radio-group';
import {
  Select,
  SelectContent,
  SelectItem,
  SelectTrigger,
  SelectValue,
} from '@/components/ui/select';
import type { LLMProvider } from '@/types';

interface ProviderSettingsProps {
  isOpen: boolean;
  onClose: () => void;
  provider: LLMProvider;
  onProviderChange: (provider: LLMProvider) => void;
}

export function ProviderSettings({
  isOpen,
  onClose,
  provider,
  onProviderChange
}: ProviderSettingsProps) {
  return (
    <Dialog open={isOpen} onOpenChange={onClose}>
      <DialogContent>
        <DialogHeader>
          <DialogTitle>LLM Provider Settings</DialogTitle>
        </DialogHeader>
        
        <div className="space-y-4">
          <RadioGroup
            value={provider.type}
            onValueChange={(type: any) => onProviderChange({ ...provider, type })}
          >
            <div className="flex items-center space-x-2">
              <RadioGroupItem value="manual" id="manual" />
              <Label htmlFor="manual">Manual (Copy/Paste)</Label>
            </div>
            <div className="flex items-center space-x-2">
              <RadioGroupItem value="lmstudio" id="lmstudio" />
              <Label htmlFor="lmstudio">LM Studio</Label>
            </div>
            <div className="flex items-center space-x-2">
              <RadioGroupItem value="openrouter" id="openrouter" />
              <Label htmlFor="openrouter">OpenRouter</Label>
            </div>
          </RadioGroup>

          {provider.type === 'lmstudio' && (
            <div className="space-y-3">
              <div>
                <Label>API URL</Label>
                <Input
                  value={provider.config?.apiUrl || 'http://localhost:1234/v1/chat/completions'}
                  onChange={e => onProviderChange({
                    ...provider,
                    config: { ...provider.config, apiUrl: e.target.value }
                  })}
                  placeholder="http://localhost:1234/v1/chat/completions"
                />
              </div>
              <div>
                <Label>Model</Label>
                <Input
                  value={provider.config?.model || 'local-model'}
                  onChange={e => onProviderChange({
                    ...provider,
                    config: { ...provider.config, model: e.target.value }
                  })}
                  placeholder="local-model"
                />
              </div>
            </div>
          )}

          {provider.type === 'openrouter' && (
            <div className="space-y-3">
              <div>
                <Label>API Key</Label>
                <Input
                  type="password"
                  value={provider.config?.apiKey || ''}
                  onChange={e => onProviderChange({
                    ...provider,
                    config: { ...provider.config, apiKey: e.target.value }
                  })}
                  placeholder="sk-or-..."
                />
              </div>
              <div>
                <Label>Model</Label>
                <Select
                  value={provider.config?.model || 'anthropic/claude-3-haiku'}
                  onValueChange={model => onProviderChange({
                    ...provider,
                    config: { ...provider.config, model }
                  })}
                >
                  <SelectTrigger>
                    <SelectValue />
                  </SelectTrigger>
                  <SelectContent>
                    <SelectItem value="anthropic/claude-3-haiku">Claude 3 Haiku</SelectItem>
                    <SelectItem value="anthropic/claude-3-sonnet">Claude 3 Sonnet</SelectItem>
                    <SelectItem value="openai/gpt-4-turbo">GPT-4 Turbo</SelectItem>
                    <SelectItem value="openai/gpt-3.5-turbo">GPT-3.5 Turbo</SelectItem>
                    <SelectItem value="google/gemini-pro">Gemini Pro</SelectItem>
                    <SelectItem value="meta-llama/llama-2-70b-chat">Llama 2 70B</SelectItem>
                    <SelectItem value="mistralai/mistral-medium">Mistral Medium</SelectItem>
                  </SelectContent>
                </Select>
              </div>
            </div>
          )}
        </div>
      </DialogContent>
    </Dialog>
  );
}