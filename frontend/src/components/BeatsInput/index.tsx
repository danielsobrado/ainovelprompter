import React from 'react';
import { Textarea } from '@/components/ui/textarea';
import { ScrollArea } from '@/components/ui/scroll-area';

interface BeatsInputProps {
  value: string;
  onChange: (value: string) => void;
}

export default function BeatsInput({ value, onChange }: BeatsInputProps) {
  return (
    <div className="space-y-2">
      <ScrollArea className="h-[400px] w-full rounded-md border">
        <Textarea
          placeholder="Enter the main story beats for the next chapter..."
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