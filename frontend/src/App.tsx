import React, { useState, useCallback, useEffect } from 'react';
import { AppLayout } from './components/AppLayout';
import { useOptionManagement } from './hooks/useOptionManagement';
import { usePromptGeneration } from './hooks/usePromptGeneration';
import {
  TaskTypeSelector,
  RulesSelector,
  BeatsInput,
  CharactersSelector,
  LocationsSelector,
  CodexSelector,
  SampleChaptersSelector, // Updated name
  PreviousChapterInput,
  FutureChapterInput,
  PromptSection,
} from './components';
import { DEFAULT_INSTRUCTIONS, PROMPT_TYPES } from './utils/constants';
import type { PromptData } from './types';
import { CheckedState } from '@radix-ui/react-checkbox';
import TaskTypeEditModal from './components/TaskTypeEditModal';
import RulesEditModal from './components/RulesEditModal';
import CharactersEditModal from './components/CharactersEditModal';
import LocationsEditModal from './components/LocationsEditModal';
import CodexEditModal from './components/CodexEditModal';
import SampleChapterEditModal from './components/SampleChapterEditModal';
import { ChapterSection } from './components/ChapterSection';

export default function App() {
  // Option management hooks
  const taskTypes = useOptionManagement({ initialOptions: [] });
  const rules = useOptionManagement({ initialOptions: [] });
  const characters = useOptionManagement({ initialOptions: [] });
  const locations = useOptionManagement({ initialOptions: [] });
  const codex = useOptionManagement({ initialOptions: [] });
  const sampleChapters = useOptionManagement({ initialOptions: [] });

  // Modal states
  const [isTaskTypeEditOpen, setIsTaskTypeEditOpen] = useState(false);
  const [isRulesEditOpen, setIsRulesEditOpen] = useState(false);
  const [isCharactersEditOpen, setIsCharactersEditOpen] = useState(false);
  const [isLocationsEditOpen, setIsLocationsEditOpen] = useState(false);
  const [isCodexEditOpen, setIsCodexEditOpen] = useState(false);
  const [isSampleChapterEditOpen, setIsSampleChapterEditOpen] = useState(false);

  // Chapter state
  const [previousChapter, setPreviousChapter] = useState('');
  const [futureChapterNotes, setFutureChapterNotes] = useState('');
  const [beats, setBeats] = useState('');
  
  // Task Type state
  const [taskTypeChecked, setTaskTypeChecked] = useState(true);
  const [selectedTaskType, setSelectedTaskType] = useState('');

  // Prompt generation
  const {
    promptType,
    setPromptType,
    finalPrompt,
    tokenCount,
    generatePrompt,
  } = usePromptGeneration();

  // Raw prompt with default instruction
  const [rawPrompt, setRawPrompt] = useState<string>(DEFAULT_INSTRUCTIONS[PROMPT_TYPES.CHATGPT]);
  
  // Handle copy to clipboard
  const handleCopy = useCallback(() => {
    navigator.clipboard.writeText(finalPrompt)
      .then(() => console.log('Copied to clipboard'))
      .catch(err => console.error('Failed to copy:', err));
  }, [finalPrompt]);

  // Handle raw prompt change with type safety
  const handleRawPromptChange = useCallback((value: string) => {
    setRawPrompt(value);
  }, []);

  // Generate prompts when data changes
  useEffect(() => {
    const promptData: PromptData = {
      taskType: selectedTaskType,
      taskTypeChecked,
      sampleChapter: sampleChapters.selectedValues[0] || '',
      previousChapterText: previousChapter,
      nextChapterBeats: beats,
      selectedRules: rules.selectedValues,
      selectedCharacters: characters.selectedValues,
      selectedLocations: locations.selectedValues,
      selectedCodexEntries: codex.selectedValues,
      futureChapterNotes,
      rawPrompt,
    };

    generatePrompt(promptData);
  }, [
    selectedTaskType,
    taskTypeChecked,
    previousChapter,
    sampleChapters.selectedValues,
    futureChapterNotes,
    beats,
    rules.selectedValues,
    characters.selectedValues,
    locations.selectedValues,
    codex.selectedValues,
    rawPrompt,
    generatePrompt,
  ]);

  return (
    <AppLayout>
      {/* Top Row with Task Type and Sample Chapter */}
      <div className="grid grid-cols-2 gap-4 mb-6">
        <TaskTypeSelector
          value={selectedTaskType}
          onChange={setSelectedTaskType}
          checked={taskTypeChecked}
          onCheckedChange={(checked: CheckedState) => setTaskTypeChecked(checked === true)}
          options={taskTypes.options}
          onEditClick={() => setIsTaskTypeEditOpen(true)}
        />
        
        <SampleChaptersSelector
          value={sampleChapters.selectedValues[0] || ''}
          onChange={(value: string) => {
            sampleChapters.setSelectedValues(value ? [value] : []);
          }}
          options={sampleChapters.options}
          onEditClick={() => setIsSampleChapterEditOpen(true)}
        />
      </div>

      {/* Tabbed Section for Beats, Previous and Future Chapters */}
      <div className="mb-6">
        <ChapterSection
          beats={beats}
          setBeats={setBeats}
          previousChapter={previousChapter}
          setPreviousChapter={setPreviousChapter}
          futureChapterNotes={futureChapterNotes}
          setFutureChapterNotes={setFutureChapterNotes}
        />
      </div>

      {/* Selectors Section */}
      <div className="space-y-4 mb-6">
        <RulesSelector
          values={rules.selectedValues}
          onChange={(values: string[]) => rules.setSelectedValues(values)}
          options={rules.options}
          onEditClick={() => setIsRulesEditOpen(true)}
        />

        <CharactersSelector
          values={characters.selectedValues}
          onChange={(values: string[]) => characters.setSelectedValues(values)}
          options={characters.options}
          onEditClick={() => setIsCharactersEditOpen(true)}
        />

        <LocationsSelector
          values={locations.selectedValues}
          onChange={(values: string[]) => locations.setSelectedValues(values)}
          options={locations.options}
          onEditClick={() => setIsLocationsEditOpen(true)}
        />

        <CodexSelector
          values={codex.selectedValues}
          onChange={(values: string[]) => codex.setSelectedValues(values)}
          options={codex.options}
          onEditClick={() => setIsCodexEditOpen(true)}
        />
      </div>

      {/* Prompt Section */}
      <div className="mb-6">
        <PromptSection
          rawPrompt={rawPrompt}
          setRawPrompt={handleRawPromptChange}
          finalPrompt={finalPrompt}
          tokenCount={tokenCount}
          onCopy={handleCopy}
          onGenerateChatGPT={() => setPromptType('ChatGPT')}
          onGenerateClaude={() => setPromptType('Claude')}
        />
      </div>

      {/* Edit Modals */}
      <TaskTypeEditModal
        isOpen={isTaskTypeEditOpen}
        onClose={() => setIsTaskTypeEditOpen(false)}
        options={taskTypes.options}
        onSave={taskTypes.setOptions}
      />

      <SampleChapterEditModal
        isOpen={isSampleChapterEditOpen}
        onClose={() => setIsSampleChapterEditOpen(false)}
        options={sampleChapters.options}
        onSave={sampleChapters.setOptions}
      />

      <RulesEditModal
        isOpen={isRulesEditOpen}
        onClose={() => setIsRulesEditOpen(false)}
        options={rules.options}
        onSave={rules.setOptions}
      />

      <CharactersEditModal
        isOpen={isCharactersEditOpen}
        onClose={() => setIsCharactersEditOpen(false)}
        options={characters.options}
        onSave={characters.setOptions}
      />

      <LocationsEditModal
        isOpen={isLocationsEditOpen}
        onClose={() => setIsLocationsEditOpen(false)}
        options={locations.options}
        onSave={locations.setOptions}
      />

      <CodexEditModal
        isOpen={isCodexEditOpen}
        onClose={() => setIsCodexEditOpen(false)}
        options={codex.options}
        onSave={codex.setOptions}
      />

    </AppLayout>
  );
}