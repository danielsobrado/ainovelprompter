// BeatsInput.tsx

import React from 'react';
import { Label } from '@/components/ui/label';
import { Textarea } from '@/components/ui/textarea';

interface BeatsInputProps {
  value: string;
  onChange: (value: string) => void;
}

export default function BeatsInput({ value, onChange }: BeatsInputProps) {
  return (
    <div className="mt-4">
      <Label htmlFor="beats">Beats</Label>
      <Textarea
        id="beats"
        value={value}
        onChange={(e) => onChange(e.target.value)}
        placeholder="Enter the beats of your story..."
      />
    </div>
  );
}
