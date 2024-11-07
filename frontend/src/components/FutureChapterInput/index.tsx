import React from 'react';
import { Textarea } from '@/components/ui/textarea';
import { ScrollArea } from '@/components/ui/scroll-area';

interface FutureChapterInputProps {
  value: string;
  onChange: (value: string) => void;
}

export function FutureChapterInput({ value, onChange }: FutureChapterInputProps) {
  return (
    <div className="space-y-1">
      <ScrollArea className="h-[180px] w-full rounded-md border">
        <Textarea
          placeholder="Add any notes about future chapters, planned plot points, or important developments..."
          value={value}
          onChange={(e) => onChange(e.target.value)}
          className="min-h-[175px] border-0"
        />
      </ScrollArea>
      <div className="text-xs text-muted-foreground">
        Character count: {value.length}
      </div>
    </div>
  );
}