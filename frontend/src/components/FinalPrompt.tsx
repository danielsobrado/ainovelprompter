// FinalPrompt.tsx
import React from 'react';
import { Card, CardHeader, CardContent, CardTitle } from '@/components/ui/card';

interface FinalPromptProps {
  value: string;
  tokenCount: number;
  onChange: (value: string) => void;
}

export default function FinalPrompt({ value, tokenCount, onChange }: FinalPromptProps) {
  return (
    <div className="mt-4">
      <Card>
        <CardHeader className="flex justify-between items-center">
          <CardTitle className="text-lg">Final Prompt</CardTitle>
          <span className="text-sm text-gray-500">Tokens: {tokenCount}</span>
        </CardHeader>
        <CardContent>
          <textarea
            value={value}
            onChange={(e) => onChange(e.target.value)}
            className="h-[400px] w-full rounded-md border p-4"
          />
        </CardContent>
      </Card>
    </div>
  );
}
