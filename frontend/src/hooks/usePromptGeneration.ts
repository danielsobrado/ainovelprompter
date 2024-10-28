import { useState, useCallback, useEffect, useRef } from 'react';
import { generateChatGPTPrompt, generateClaudePrompt } from '../utils/promptGenerators';
import type { PromptData, RuleOption, CharacterOption, LocationOption, CodexOption, TaskTypeOption } from '../types';

type PromptType = 'ChatGPT' | 'Claude';

interface GeneratePromptOptions {
    rules: RuleOption[];
    characters: CharacterOption[];
    locations: LocationOption[];
    codex: CodexOption[];
    taskTypes: TaskTypeOption[];
  }

export function usePromptGeneration() {
  const [promptType, setPromptType] = useState<PromptType>('ChatGPT');
  const [finalPrompt, setFinalPrompt] = useState('');
  const [tokenCount, setTokenCount] = useState(0);
  const prevPromptType = useRef<PromptType>('ChatGPT');

  const generatePrompt = useCallback((
    data: PromptData, 
    options: GeneratePromptOptions
  ) => {
    const prompt = promptType === 'ChatGPT' 
      ? generateChatGPTPrompt(data, options)
      : generateClaudePrompt(data, options);
    
    setFinalPrompt(prompt);
    setTokenCount(prompt.trim().split(/\s+/).length);
  }, [promptType]);

  useEffect(() => {
    prevPromptType.current = promptType;
  }, [promptType]);

  return {
    promptType,
    setPromptType,
    finalPrompt,
    setFinalPrompt,
    tokenCount,
    generatePrompt,
    prevPromptType: prevPromptType.current
  };
}