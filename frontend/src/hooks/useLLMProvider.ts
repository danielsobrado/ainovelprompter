// frontend/src/hooks/useLLMProvider.ts
import { useState, useCallback } from 'react';
import type { LLMProvider } from '@/types';

export function useLLMProvider(provider: LLMProvider) {
  const [isLoading, setIsLoading] = useState(false);
  const [error, setError] = useState<string | null>(null);

  // Add function to clear error state
  const clearError = useCallback(() => {
    setError(null);
  }, []);

  const executePrompt = useCallback(async (prompt: string): Promise<string> => {
    setIsLoading(true);
    setError(null);

    try {
      switch (provider.type) {
        case 'lmstudio':
          return await executeLMStudioPrompt(prompt, provider.config);
        
        case 'openrouter':
          return await executeOpenRouterPrompt(prompt, provider.config);
        
        case 'manual':
        default:
          throw new Error('Manual mode requires user input');
      }
    } catch (err) {
      setError(err instanceof Error ? err.message : 'Unknown error');
      throw err;
    } finally {
      setIsLoading(false);
    }
  }, [provider]);

  return { executePrompt, isLoading, error, clearError };
}

async function executeLMStudioPrompt(
  prompt: string, 
  config?: LLMProvider['config']
): Promise<string> {
  const apiUrl = config?.apiUrl || 'http://localhost:1234/v1/chat/completions';
  const requestBody = {
    model: config?.model || 'local-model',
    messages: [{ role: 'user', content: prompt }],
    temperature: 0.7,
    max_tokens: 2000,
    // stream: false, // LM Studio often defaults this; explicitly setting might help if there's ambiguity
  };

  console.log('[LMStudio Request] URL:', apiUrl);
  console.log('[LMStudio Request] Body:', JSON.stringify(requestBody, null, 2));

  const response = await fetch(apiUrl, {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    body: JSON.stringify(requestBody),
  });

  if (!response.ok) {
    throw new Error(`LM Studio error: ${response.statusText}`);
  }

  const data = await response.json();
  console.log('[LMStudio Response] Data:', data);
  return data.choices[0].message.content;
}

async function executeOpenRouterPrompt(
  prompt: string,
  config?: LLMProvider['config']
): Promise<string> {
  const response = await fetch('https://openrouter.ai/api/v1/chat/completions', {
    method: 'POST',
    headers: {
      'Authorization': `Bearer ${config?.apiKey}`,
      'Content-Type': 'application/json',
      'HTTP-Referer': window.location.origin,
      'X-Title': 'AI Novel Prompter',
    },
    body: JSON.stringify({
      model: config?.model || 'anthropic/claude-3-haiku',
      messages: [{ role: 'user', content: prompt }],
      max_tokens: 4096, // Added max_tokens
    }),
  });

  if (!response.ok) {
    throw new Error(`OpenRouter error: ${response.statusText}`);
  }

  const data = await response.json();
  return data.choices[0].message.content;
}