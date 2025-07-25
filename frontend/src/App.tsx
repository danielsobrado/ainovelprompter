import React, { useState, useCallback, useEffect } from 'react';
import { EventsOn } from '../wailsjs/runtime/runtime';
import { 
  ReadCharactersFile, 
  WriteCharactersFile, 
  ReadLocationsFile, 
  WriteLocationsFile, 
  ReadCodexFile, 
  WriteCodexFile, 
  ReadRulesFile, 
  WriteRulesFile, 
  ReadTaskTypesFile, 
  WriteTaskTypesFile, 
  ReadSampleChaptersFile, 
  WriteSampleChaptersFile 
} from '../wailsjs/go/main/App';
import WailsReadyContext from './contexts/WailsReadyContext'; // Import the context
import { DataRefreshProvider, useDataRefresh } from './contexts/DataRefreshContext';
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
import { Tabs, TabsContent, TabsList, TabsTrigger } from "@/components/ui/tabs";
import { CheckedState } from '@radix-ui/react-checkbox';
import { Trash2 } from 'lucide-react';
import { Button } from '@/components/ui/button';
import TaskTypeEditModal from './components/TaskTypeSelector/TaskTypeEditModal';
import RulesEditModal from './components/RulesSelector/RulesEditModal';
import CharactersEditModal from './components/Characters/CharactersEditModal';
import LocationsEditModal from './components/Locations/LocationsEditModal';
import CodexEditModal from './components/Codex/CodexEditModal';
import SampleChapterEditModal from './components/SampleChapter/SampleChapterEditModal';
import { ChapterSection } from './components/ChapterSection';
import { ProseImprovementTab } from './components/ProseImprovement';

// Main app content component  
function AppContent() {
  const [wailsReady, setWailsReady] = useState(false);

  useEffect(() => {
    let unlisten: (() => void) | null = null;
    let timeoutId: NodeJS.Timeout | null = null;

    const checkWailsReady = () => {
      // Check if Wails runtime and Go bindings are available
      if (window.runtime && window.go && window.go.main && window.go.main.App) {
        console.log("Wails is ready!");
        setWailsReady(true);
        return true;
      }
      return false;
    };

    // Try immediate check first
    if (checkWailsReady()) {
      return;
    }

    // Set up event listener for Wails ready event
    try {
      unlisten = EventsOn("wails:ready", () => {
        console.log("Wails ready event received!");
        setWailsReady(true);
      });
    } catch (error) {
      console.warn("Failed to set up Wails ready event listener:", error);
    }

    // Fallback: periodic check with timeout
    const startTime = Date.now();
    const maxWaitTime = 10000; // 10 seconds max wait
    
    const periodicCheck = () => {
      if (checkWailsReady()) {
        return;
      }
      
      if (Date.now() - startTime > maxWaitTime) {
        console.warn("Wails readiness timeout - proceeding anyway");
        setWailsReady(true);
        return;
      }
      
      timeoutId = setTimeout(periodicCheck, 100);
    };

    periodicCheck();

    return () => {
      if (unlisten && typeof unlisten === 'function') {
        unlisten();
      }
      if (timeoutId) {
        clearTimeout(timeoutId);
      }
    };
  }, []);

  // Option management hooks
  const taskTypes = useOptionManagement({ 
    initialOptions: [], 
    readFile: ReadTaskTypesFile,
    writeFile: WriteTaskTypesFile
  });
  
  const rules = useOptionManagement({ 
    initialOptions: [], 
    readFile: ReadRulesFile,
    writeFile: WriteRulesFile
  });
  
  const characters = useOptionManagement({ 
    initialOptions: [], 
    readFile: ReadCharactersFile,
    writeFile: WriteCharactersFile
  });
  
  const locations = useOptionManagement({ 
    initialOptions: [], 
    readFile: ReadLocationsFile,
    writeFile: WriteLocationsFile
  });
  
  const codex = useOptionManagement({ 
    initialOptions: [], 
    readFile: ReadCodexFile,
    writeFile: WriteCodexFile
  });
  
  const sampleChapters = useOptionManagement({ 
    initialOptions: [], 
    readFile: ReadSampleChaptersFile,
    writeFile: WriteSampleChaptersFile
  });

  // Set up refresh context
  const { setRefreshFunctions } = useDataRefresh();

  // Initialize refresh functions
  useEffect(() => {
    if (wailsReady) {
      const refreshFunctions = {
        refreshTaskTypes: async () => {
          console.log('Refreshing task types...');
          await taskTypes.refreshOptions();
        },
        refreshRules: async () => {
          console.log('Refreshing rules...');
          await rules.refreshOptions();
        },
        refreshCharacters: async () => {
          console.log('Refreshing characters...');
          await characters.refreshOptions();
        },
        refreshLocations: async () => {
          console.log('Refreshing locations...');
          await locations.refreshOptions();
        },
        refreshCodex: async () => {
          console.log('Refreshing codex...');
          await codex.refreshOptions();
        },
        refreshSampleChapters: async () => {
          console.log('Refreshing sample chapters...');
          await sampleChapters.refreshOptions();
        },
        refreshAll: async () => {
          console.log('Refreshing all data...');
          await Promise.all([
            taskTypes.refreshOptions(),
            rules.refreshOptions(),
            characters.refreshOptions(),
            locations.refreshOptions(),
            codex.refreshOptions(),
            sampleChapters.refreshOptions(),
          ]);
        },
      };
      setRefreshFunctions(refreshFunctions);
    }
  }, [wailsReady, taskTypes.refreshOptions, rules.refreshOptions, characters.refreshOptions, 
      locations.refreshOptions, codex.refreshOptions, sampleChapters.refreshOptions, setRefreshFunctions]);

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
  const [taskTypeChecked, setTaskTypeChecked] = useState(false);  
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
    setRawPrompt('');
    
    // Reset task type
    setSelectedTaskType('');  // This will trigger the effect above to uncheck the box
    
    // Clear all selections
    sampleChapters.setSelectedValues([]);
    rules.setSelectedValues([]);
    characters.setSelectedValues([]);
    locations.setSelectedValues([]);
    codex.setSelectedValues([]);
    
    // Reset prompt type
    setPromptType('ChatGPT');
  }, [sampleChapters, rules, characters, locations, codex, setPromptType]);
  
  useEffect(() => {
    // Only check the box if there's a selected task type
    if (selectedTaskType && !taskTypeChecked) {
      setTaskTypeChecked(true);
    }
    // Uncheck if no task type is selected
    if (!selectedTaskType && taskTypeChecked) {
      setTaskTypeChecked(false);
    }
  }, [selectedTaskType]);

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
    <WailsReadyContext.Provider value={{ wailsReady }}>
      <AppLayout>
        <Tabs defaultValue="prompt-generation" className="w-full">
          <TabsList className="grid w-full grid-cols-2 mb-4">
            <TabsTrigger value="prompt-generation">Prompt Generation</TabsTrigger>
            <TabsTrigger value="prose-improvement">Prose Improvement</TabsTrigger>
          </TabsList>

          <TabsContent value="prompt-generation" className="space-y-6">
            {/* Top Row with Task Type, Sample Chapter, and Clear All */}
            <div className="grid grid-cols-[1fr_1fr_auto] gap-3 items-center">
              <TaskTypeSelector
                value={selectedTaskType}
                onChange={(value: string) => {
                  setSelectedTaskType(value);
                }}
                checked={taskTypeChecked}
                onCheckedChange={(checked: CheckedState) => {
                  setTaskTypeChecked(checked === true);
                  // If unchecking, clear the selected task type
                  if (!checked) {
                    setSelectedTaskType('');
                  }
                }}
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
            <div>
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
            <div className="space-y-3">
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
            <div>
            <PromptSection
              rawPrompt={rawPrompt}
              setRawPrompt={handleRawPromptChange}
              finalPrompt={finalPrompt}
              tokenCount={tokenCount}
              onCopy={handleCopy}
              onGenerateChatGPT={() => setPromptType('ChatGPT')}
              onGenerateClaude={() => setPromptType('Claude')}
              taskTypeChecked={taskTypeChecked}
              currentPromptType={promptType}
            />
            </div>
          </TabsContent>

          <TabsContent value="prose-improvement">
            <ProseImprovementTab />
          </TabsContent>
        </Tabs>

        {/* Edit Modals */}
        <TaskTypeEditModal
          isOpen={isTaskTypeEditOpen}
          onClose={() => setIsTaskTypeEditOpen(false)}
          options={taskTypes.options}
          onSave={taskTypes.saveOptions}
        />

        <SampleChapterEditModal
          isOpen={isSampleChapterEditOpen}
          onClose={() => setIsSampleChapterEditOpen(false)}
          options={sampleChapters.options}
          onSave={sampleChapters.saveOptions}
        />

        <RulesEditModal
          isOpen={isRulesEditOpen}
          onClose={() => setIsRulesEditOpen(false)}
          options={rules.options}
          onSave={rules.saveOptions}
        />

        <CharactersEditModal
          isOpen={isCharactersEditOpen}
          onClose={() => setIsCharactersEditOpen(false)}
          options={characters.options}
          onSave={characters.saveOptions}
        />

        <LocationsEditModal
          isOpen={isLocationsEditOpen}
          onClose={() => setIsLocationsEditOpen(false)}
          options={locations.options}
          onSave={locations.saveOptions}
        />

        <CodexEditModal
          isOpen={isCodexEditOpen}
          onClose={() => setIsCodexEditOpen(false)}
          options={codex.options}
          onSave={codex.saveOptions}
        />
      </AppLayout>
    </WailsReadyContext.Provider>
  );
}

// Main App component with DataRefreshProvider
export function App() {
  return (
    <DataRefreshProvider>
      <AppContent />
    </DataRefreshProvider>
  );
}

export default App;