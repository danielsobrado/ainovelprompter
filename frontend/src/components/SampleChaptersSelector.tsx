// SampleChaptersSelector.tsx

import React from 'react';
import { Label } from '@/components/ui/label';
import {
  Select,
  SelectTrigger,
  SelectValue,
  SelectContent,
  SelectItem,
} from '@/components/ui/select';
import { Button } from '@/components/ui/button';
import { Textarea } from '@/components/ui/textarea';
import { Edit } from 'lucide-react';
import { SampleChapterOption } from '../types';

interface SampleChaptersSelectorProps {
  previousChapter: string;
  onPreviousChapterChange: (value: string) => void;
  sampleChapter: string;
  onSampleChapterChange: (value: string) => void;
  nextChapterBeats: string;
  onNextChapterBeatsChange: (value: string) => void;
  options: SampleChapterOption[];
  onEditClick: () => void;
}

export default function SampleChaptersSelector({
  previousChapter,
  onPreviousChapterChange,
  sampleChapter,
  onSampleChapterChange,
  nextChapterBeats,
  onNextChapterBeatsChange,
  options,
  onEditClick,
}: SampleChaptersSelectorProps) {
  return (
    <div className="flex flex-col space-y-4">
      {/* Row for Previous Chapter and Sample Chapter */}
      <div className="flex items-center space-x-4">
        {/* Previous Chapter */}
        <div className="flex items-center space-x-2 w-1/2">
          <Label htmlFor="previous-chapter" className="whitespace-nowrap">
            Previous Chapter
          </Label>
          <Select value={previousChapter} onValueChange={onPreviousChapterChange}>
            <SelectTrigger id="previous-chapter" className="w-full">
              <SelectValue placeholder="Select previous chapter" />
            </SelectTrigger>
            <SelectContent>
              {options.map((option) => (
                <SelectItem key={option.id} value={option.label}>
                  {option.label}
                </SelectItem>
              ))}
            </SelectContent>
          </Select>
        </div>

        {/* Sample Chapter */}
        <div className="flex items-center space-x-2 w-1/2">
          <Label htmlFor="sample-chapter" className="whitespace-nowrap">
            Sample Chapter
          </Label>
          <Select value={sampleChapter} onValueChange={onSampleChapterChange}>
            <SelectTrigger id="sample-chapter" className="w-full">
              <SelectValue placeholder="Select sample chapter" />
            </SelectTrigger>
            <SelectContent>
              {options.map((option) => (
                <SelectItem key={option.id} value={option.label}>
                  {option.label}
                </SelectItem>
              ))}
            </SelectContent>
          </Select>
          <Button variant="ghost" size="icon" onClick={onEditClick}>
            <Edit className="h-4 w-4" />
          </Button>
        </div>
      </div>

      {/* Next Chapter Beats */}
      <div className="flex flex-col space-y-2">
        <Label htmlFor="next-chapter-beats" className="whitespace-nowrap">
          Next Chapter Beats
        </Label>
        <Textarea
          id="next-chapter-beats"
          value={nextChapterBeats}
          onChange={(e) => onNextChapterBeatsChange(e.target.value)}
          placeholder="Enter next chapter beats..."
          className="w-full min-h-[100px]"
        />
      </div>
    </div>
  );
}
