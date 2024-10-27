// components/PromptSection/index.tsx
import React from 'react';
import { Card, CardContent, CardHeader, CardTitle } from '@/components/ui/card';
import { Button } from '@/components/ui/button';
import { Textarea } from '@/components/ui/textarea';
import { ScrollArea } from '@/components/ui/scroll-area';
import { Copy, RefreshCw } from 'lucide-react';

interface PromptSectionProps {
  rawPrompt: string;
  setRawPrompt: (value: string) => void;
  finalPrompt: string;
  tokenCount: number;
  onCopy: () => void;
  onGenerateChatGPT: () => void;
  onGenerateClaude: () => void;
}

export default function PromptSection({
  rawPrompt,
  setRawPrompt,
  finalPrompt,
  tokenCount,
  onCopy,
  onGenerateChatGPT,
  onGenerateClaude,
}: PromptSectionProps) {
  return (
    <div className="space-y-6">
      {/* Raw Prompt Input */}
      <Card>
        <CardHeader>
          <CardTitle className="text-lg font-semibold">Custom Instructions</CardTitle>
        </CardHeader>
        <CardContent>
          <Textarea
            placeholder="Enter any additional instructions or requirements..."
            value={rawPrompt}
            onChange={(e) => setRawPrompt(e.target.value)}
            className="min-h-[100px]"
          />
        </CardContent>
      </Card>

      {/* Generated Prompt */}
      <Card>
        <CardHeader className="flex flex-row items-center justify-between">
          <CardTitle className="text-lg font-semibold">Generated Prompt</CardTitle>
          <div className="flex items-center space-x-2">
            <Button
              variant="outline"
              size="sm"
              onClick={onGenerateChatGPT}
            >
              ChatGPT
            </Button>
            <Button
              variant="outline"
              size="sm"
              onClick={onGenerateClaude}
            >
              Claude
            </Button>
          </div>
        </CardHeader>
        <CardContent>
          <div className="space-y-4">
            <ScrollArea className="h-[300px] w-full rounded-md border">
              <div className="p-4">
                <pre className="whitespace-pre-wrap font-mono text-sm">{finalPrompt}</pre>
              </div>
            </ScrollArea>
            <div className="flex items-center justify-between">
              <span className="text-sm text-muted-foreground">
                Token count: {tokenCount}
              </span>
              <Button
                variant="outline"
                size="sm"
                onClick={onCopy}
              >
                <Copy className="mr-2 h-4 w-4" />
                Copy to Clipboard
              </Button>
            </div>
          </div>
        </CardContent>
      </Card>
    </div>
  );
}