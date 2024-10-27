import React from 'react';
import { Textarea } from '@/components/ui/textarea';
import { ScrollArea } from '@/components/ui/scroll-area';

interface FutureChapterInputProps {
  value: string;
  onChange: (value: string) => void;
}

export function FutureChapterInput({ value, onChange }: FutureChapterInputProps) {
  return (
    <div className="space-y-2">
      <ScrollArea className="h-[400px] w-full rounded-md border">
        <Textarea
          placeholder="Add any notes about future chapters, planned plot points, or important developments..."
          value={value}
          onChange={(e) => onChange(e.target.value)}
          className="min-h-[380px] border-0"
        />
      </ScrollArea>
      <div className="text-sm text-muted-foreground">
        Character count: {value.length}
      </div>
    </div>
  );
}