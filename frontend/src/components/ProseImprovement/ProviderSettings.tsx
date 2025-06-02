// frontend/src/components/ProseImprovement/ProviderSettings.tsx
import React from 'react';
import {
  Dialog,
  DialogContent,
  DialogHeader,
  DialogTitle,
  DialogDescription, // Import DialogDescription
} from '@/components/ui/dialog';
import { Input } from '@/components/ui/input';
import { Label } from '@/components/ui/label';
import { RadioGroup, RadioGroupItem } from '@/components/ui/radio-group';
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
          {/* Add a description */}
          <DialogDescription>
            Configure your connection to Manual, LM Studio, or OpenRouter providers.
          </DialogDescription>
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
                <Input
                  value={provider.config?.model || 'anthropic/claude-3-haiku'}
                  onChange={e => onProviderChange({
                    ...provider,
                    config: { ...provider.config, model: e.target.value }
                  })}
                  placeholder="e.g., anthropic/claude-3-haiku, openai/gpt-4-turbo"
                />
                <p className="text-xs text-muted-foreground mt-1">
                  Enter the full model identifier from OpenRouter.
                </p>
              </div>
            </div>
          )}
        </div>
      </DialogContent>
    </Dialog>
  );
}