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
  SampleChaptersSelector, 
  PreviousChapterInput,
  FutureChapterInput,
  PromptSection,
} from './components';
import { DEFAULT_INSTRUCTIONS, PROMPT_TYPES } from './utils/constants';
import { generateDynamicInstructions } from './utils/promptInstructions';
import type { PromptData } from './types';
import { CheckedState } from '@radix-ui/react-checkbox';
import { Trash2 } from 'lucide-react';
import { Button } from '@/components/ui/button';
import TaskTypeEditModal from './components/TaskTypeEditModal';
import RulesEditModal from './components/RulesEditModal';
import CharactersEditModal from './components/CharactersEditModal';
import LocationsEditModal from './components/LocationsEditModal';
import CodexEditModal from './components/CodexEditModal';
import SampleChapterEditModal from './components/SampleChapterEditModal';
import { ChapterSection } from './components/ChapterSection';

export function App() {
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
  
  // Handle clear all
  const handleClearAll = useCallback(() => {
    // Clear all text inputs
    setPreviousChapter('');
    setFutureChapterNotes('');
    setBeats('');
    
    // Reset task type and instructions
    setSelectedTaskType('');
    setTaskTypeChecked(true);  // This should come before setRawPrompt
    setRawPrompt('');  // Clear the raw prompt first
    
    // Clear all selections
    sampleChapters.setSelectedValues([]);
    rules.setSelectedValues([]);
    characters.setSelectedValues([]);
    locations.setSelectedValues([]);
    codex.setSelectedValues([]);
    
    // Reset prompt type to ChatGPT
    setPromptType('ChatGPT');
  }, [sampleChapters, rules, characters, locations, codex, setPromptType]);

  // Update raw prompt when selections change
  useEffect(() => {
    if (!taskTypeChecked) {
      // Only generate dynamic instructions if there's any content to reference
      const hasContent = previousChapter || 
                        beats || 
                        characters.selectedValues.length > 0 || 
                        locations.selectedValues.length > 0 || 
                        codex.selectedValues.length > 0 || 
                        rules.selectedValues.length > 0;

      if (hasContent) {
        const newInstructions = generateDynamicInstructions({
          previousChapter,
          beats,
          selectedCharacters: characters.selectedValues,
          selectedLocations: locations.selectedValues,
          selectedCodexEntries: codex.selectedValues,
          selectedRules: rules.selectedValues,
          isClaudeFormat: promptType === 'Claude'
        });
        setRawPrompt(newInstructions);
      } else {
        // If there's no content, set to empty
        setRawPrompt('');
      }
    }
  }, [
    taskTypeChecked,
    previousChapter,
    beats,
    characters.selectedValues,
    locations.selectedValues,
    codex.selectedValues,
    rules.selectedValues,
    promptType
  ]);

  // Handle default instructions
  useEffect(() => {
    if (taskTypeChecked && !selectedTaskType) {
      setRawPrompt(DEFAULT_INSTRUCTIONS[promptType]);
    }
  }, [taskTypeChecked, selectedTaskType, promptType]);

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

    generatePrompt(
      promptData,
      {
        rules: rules.options,
        characters: characters.options,
        locations: locations.options,
        codex: codex.options,
        taskTypes: taskTypes.options,
      }
    );
  }, [
    selectedTaskType,
    taskTypeChecked,
    previousChapter,
    sampleChapters.selectedValues,
    futureChapterNotes,
    beats,
    rules.selectedValues,
    rules.options,
    characters.selectedValues,
    characters.options,
    locations.selectedValues,
    locations.options,
    codex.selectedValues,
    codex.options,
    taskTypes.options,
    rawPrompt,
    generatePrompt,
  ]);

  return (
    <AppLayout>
      {/* Top Row with Task Type, Sample Chapter, and Clear All */}
      <div className="grid grid-cols-[1fr_1fr_auto] gap-4 mb-6 items-center">
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

        <Button 
          variant="destructive"
          onClick={handleClearAll}
          size="sm"
        >
          <Trash2 className="mr-2 h-4 w-4" />
          Clear All
        </Button>
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
        taskTypeChecked={taskTypeChecked}
        currentPromptType={promptType}  // Add this prop
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

export default App;