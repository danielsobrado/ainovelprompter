import React from 'react';
import { Card } from '@/components/ui/card';
import { Tabs, TabsList, TabsTrigger, TabsContent } from '@/components/ui/tabs';
import { PreviousChapterInput } from '../PreviousChapterInput';
import { FutureChapterInput } from '../FutureChapterInput';
import BeatsInput from '../BeatsInput';

interface ChapterSectionProps {
  previousChapter: string;
  setPreviousChapter: (value: string) => void;
  futureChapterNotes: string;
  setFutureChapterNotes: (value: string) => void;
  beats: string;
  setBeats: (value: string) => void;
}

export function ChapterSection({
  previousChapter,
  setPreviousChapter,
  futureChapterNotes,
  setFutureChapterNotes,
  beats,
  setBeats,
}: ChapterSectionProps) {
  return (
    <Card className="p-6">
      <Tabs defaultValue="beats" className="w-full">
        <TabsList className="grid w-full grid-cols-3 mb-6">
          <TabsTrigger value="beats">Story Beats</TabsTrigger>
          <TabsTrigger value="previous">Previous Chapter</TabsTrigger>
          <TabsTrigger value="future">Future Notes</TabsTrigger>
        </TabsList>
        
        <div className="mt-4">
          <TabsContent value="beats" className="m-0">
            <BeatsInput
              value={beats}
              onChange={setBeats}
            />
          </TabsContent>
          <TabsContent value="previous" className="m-0">
            <PreviousChapterInput
              value={previousChapter}
              onChange={setPreviousChapter}
            />
          </TabsContent>
          <TabsContent value="future" className="m-0">
            <FutureChapterInput
              value={futureChapterNotes}
              onChange={setFutureChapterNotes}
            />
          </TabsContent>
        </div>
      </Tabs>
    </Card>
  );
}