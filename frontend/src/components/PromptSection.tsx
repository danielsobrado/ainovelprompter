import React from 'react';
import { Card } from '@/components/ui/card';
import RawPrompt from './RawPrompt';
import FinalPrompt from './FinalPrompt';
import ActionButtons from './ActionButtons';

interface PromptSectionProps {
  rawPrompt: string;
  setRawPrompt: (value: string) => void;
  finalPrompt: string;
  tokenCount: number;
  onCopy: () => void;
  onGenerateChatGPT: () => void;
  onGenerateClaude: () => void;
}

export function PromptSection(props: PromptSectionProps) {
  const {
    rawPrompt,
    setRawPrompt,
    finalPrompt,
    tokenCount,
    onCopy,
    onGenerateChatGPT,
    onGenerateClaude,
  } = props;

  return (
    <div className="space-y-6">
      <RawPrompt
        value={rawPrompt}
        onChange={setRawPrompt}
      />
      
      <FinalPrompt
        value={finalPrompt}
        tokenCount={tokenCount}
        onChange={() => {}}
      />
      
      <ActionButtons
        onCopy={onCopy}
        onGenerateChatGPT={onGenerateChatGPT}
        onGenerateClaude={onGenerateClaude}
      />
    </div>
  );
}