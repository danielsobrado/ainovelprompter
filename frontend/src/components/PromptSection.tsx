import React from 'react';
import { Card, CardContent, CardHeader, CardTitle } from '@/components/ui/card';
import { Button } from '@/components/ui/button';
import { Textarea } from '@/components/ui/textarea';
import { ScrollArea } from '@/components/ui/scroll-area';
import { Copy } from 'lucide-react';
import { cn } from '@/lib/utils';

interface PromptSectionProps {
  rawPrompt: string;
  setRawPrompt: (value: string) => void;
  finalPrompt: string;
  tokenCount: number;
  onCopy: () => void;
  onGenerateChatGPT: () => void;
  onGenerateClaude: () => void;
  taskTypeChecked: boolean;
  currentPromptType: 'ChatGPT' | 'Claude';  
}

export function PromptSection({
  rawPrompt,
  setRawPrompt,
  finalPrompt,
  tokenCount,
  onCopy,
  onGenerateChatGPT,
  onGenerateClaude,
  taskTypeChecked,
  currentPromptType,
}: PromptSectionProps) {
  return (
    <div className="space-y-6">
      {/* Raw Prompt Input */}
      <Card>
        <CardHeader>
          <CardTitle className="text-lg font-semibold">Task Instructions</CardTitle>
        </CardHeader>
        <CardContent>
          <Textarea
            placeholder="Task instructions will be generated based on your selections..."
            value={rawPrompt}
            onChange={(e) => setRawPrompt(e.target.value)}
            className="min-h-[100px]"
            disabled={taskTypeChecked}
          />
          {taskTypeChecked && (
            <p className="mt-2 text-sm text-muted-foreground">
              Task instructions are currently determined by the selected task type.
              Uncheck the task type checkbox to use custom instructions.
            </p>
          )}
        </CardContent>
      </Card>

      {/* Generated Prompt */}
      <Card>
        <CardHeader className="flex flex-row items-center justify-between">
          <CardTitle className="text-lg font-semibold">Generated Prompt</CardTitle>
          <div className="flex items-center space-x-2">
            <Button
              variant={currentPromptType === 'ChatGPT' ? "default" : "outline"}
              size="sm"
              onClick={onGenerateChatGPT}
              className={cn(
                "transition-colors",
                currentPromptType === 'ChatGPT' && "bg-primary text-primary-foreground hover:bg-primary/90"
              )}
            >
              ChatGPT
            </Button>
            <Button
              variant={currentPromptType === 'Claude' ? "default" : "outline"}
              size="sm"
              onClick={onGenerateClaude}
              className={cn(
                "transition-colors",
                currentPromptType === 'Claude' && "bg-primary text-primary-foreground hover:bg-primary/90"
              )}
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