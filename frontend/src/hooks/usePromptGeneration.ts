import { useState, useCallback, useEffect, useRef } from 'react';
import { generateChatGPTPrompt, generateClaudePrompt } from '../promptGenerators';
import type { PromptData } from '../types';

type PromptType = 'ChatGPT' | 'Claude';

export function usePromptGeneration() {
  const [promptType, setPromptType] = useState<PromptType>('ChatGPT');
  const [finalPrompt, setFinalPrompt] = useState('');
  const [tokenCount, setTokenCount] = useState(0);
  const prevPromptType = useRef<PromptType>('ChatGPT');

  const generatePrompt = useCallback((data: PromptData) => {
    const prompt = promptType === 'ChatGPT' 
      ? generateChatGPTPrompt(data)
      : generateClaudePrompt(data);
    
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