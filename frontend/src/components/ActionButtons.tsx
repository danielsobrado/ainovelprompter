// ActionButtons.tsx
import React from 'react';
import { Button } from '@/components/ui/button';
import { Copy, RefreshCw } from 'lucide-react';

interface ActionButtonsProps {
  onCopy: () => void;
  onGenerateChatGPT: () => void;
  onGenerateClaude: () => void;
}

export default function ActionButtons({
  onCopy,
  onGenerateChatGPT,
  onGenerateClaude,
}: ActionButtonsProps) {
  return (
    <div className="flex space-x-2 mt-4">
      <Button onClick={onCopy} variant="default">
        <Copy className="mr-2 h-4 w-4" /> Copy
      </Button>
      <Button onClick={onGenerateChatGPT} variant="default">
        <RefreshCw className="mr-2 h-4 w-4" /> Generate ChatGPT
      </Button>
      <Button onClick={onGenerateClaude} variant="default">
        <RefreshCw className="mr-2 h-4 w-4" /> Generate Claude
      </Button>
    </div>
  );
}
