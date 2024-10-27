// App.tsx

import React, { useState, useEffect, useCallback, useRef } from 'react';
import Header from './components/Header';
import SampleChaptersSelector from './components/SampleChaptersSelector';
import TaskTypeSelector from './components/TaskTypeSelector';
import RulesSelector from './components/RulesSelector';
import BeatsInput from './components/BeatsInput';
import CharactersSelector from './components/CharactersSelector';
import LocationsSelector from './components/LocationsSelector';
import CodexSelector from './components/CodexSelector';
import RawPrompt from './components/RawPrompt';
import FinalPrompt from './components/FinalPrompt';
import ActionButtons from './components/ActionButtons';
import { SettingsModal } from './components/SettingsModal';
import SampleChapterEditModal from './components/SampleChapterEditModal';
import TaskTypeEditModal from './components/TaskTypeEditModal';
import RulesEditModal from './components/RulesEditModal';
import CharactersEditModal from './components/CharactersEditModal';
import LocationsEditModal from './components/LocationsEditModal';
import CodexEditModal from './components/CodexEditModal';
import { v4 as uuidv4 } from 'uuid';
import { CharacterOption, CodexOption, LocationOption, RuleOption, SampleChapterOption, TaskTypeOption } from './types';

function App() {

  const [isSettingsOpen, setIsSettingsOpen] = useState<boolean>(false);

  // Sample Chapters
  const [sampleChapterOptions, setSampleChapterOptions] = useState<SampleChapterOption[]>([]);
  const [previousChapter, setPreviousChapter] = useState<string>('');
  const [sampleChapter, setSampleChapter] = useState<string>('');
  const [nextChapterBeats, setNextChapterBeats] = useState<string>('');
  const [isSampleChapterEditOpen, setIsSampleChapterEditOpen] = useState<boolean>(false);

  // Task Type
  const [taskType, setTaskType] = useState<string>('');
  const [taskTypeChecked, setTaskTypeChecked] = useState<boolean>(true);
  const [taskTypeOptions, setTaskTypeOptions] = useState<TaskTypeOption[]>([]);
  const [isTaskTypeEditOpen, setIsTaskTypeEditOpen] = useState<boolean>(false);

  // Rules
  const [selectedRules, setSelectedRules] = useState<string[]>([]);
  const [rulesOptions, setRulesOptions] = useState<RuleOption[]>([]);
  const [isRulesEditOpen, setIsRulesEditOpen] = useState<boolean>(false);

  // Beats
  const [beats, setBeats] = useState<string>('');

  // Characters
  const [selectedCharacters, setSelectedCharacters] = useState<string[]>([]);
  const [characterOptions, setCharacterOptions] = useState<CharacterOption[]>([]);
  const [isCharactersEditOpen, setIsCharactersEditOpen] = useState<boolean>(false);

  // Locations
  const [selectedLocations, setSelectedLocations] = useState<string[]>([]);
  const [locationOptions, setLocationOptions] = useState<LocationOption[]>([]);
  const [isLocationsEditOpen, setIsLocationsEditOpen] = useState<boolean>(false);

  // Codex
  const [selectedCodexEntries, setSelectedCodexEntries] = useState<string[]>([]);
  const [codexOptions, setCodexOptions] = useState<CodexOption[]>([]);
  const [isCodexEditOpen, setIsCodexEditOpen] = useState<boolean>(false);

  // Raw Prompt
  const [rawPrompt, setRawPrompt] = useState<string>('');

  // Final Prompt
  const [finalPrompt, setFinalPrompt] = useState<string>('');
  const [tokenCount, setTokenCount] = useState<number>(0);

  // Current Prompt Type
  const [currentPromptType, setCurrentPromptType] = useState<'ChatGPT' | 'Claude'>('ChatGPT');

  // Previous Prompt Type
  const prevPromptType = useRef<'ChatGPT' | 'Claude'>('ChatGPT');

  // Default Task Instructions
  const DEFAULT_CHATGPT_INSTRUCTION =
    'You are a creative writer tasked with composing the next chapter based on the provided beats and context. Follow the rules strictly.';

  const DEFAULT_CLAUDE_INSTRUCTION =
    'You are a creative writer tasked with composing the next chapter based on the above <BEATS> and context. Follow the <RULES> strictly.';

  // Load options on component mount
  useEffect(() => {
    // Load Sample Chapters
    const loadSampleChapters = async () => {
      try {
        // Replace with your data loading logic
        const options: SampleChapterOption[] = [
          {
            id: uuidv4(),
            label: 'Chapter 1',
            content: 'Content of Chapter 1...',
          },
          // Add more sample chapters
        ];
        setSampleChapterOptions(options);
        setSampleChapter(options[0]?.label || '');
      } catch (error) {
        console.error('Error loading sample chapters:', error);
      }
    };

    // Load Task Types
    const loadTaskTypes = async () => {
      try {
        const options: TaskTypeOption[] = [
          {
            id: uuidv4(),
            label: 'Write Next Chapter',
            description: 'Compose the next chapter of the story.',
          },
          // Add more task types
        ];
        setTaskTypeOptions(options);
        setTaskType(options[0]?.label || '');
      } catch (error) {
        console.error('Error loading task types:', error);
      }
    };

    // Load Rules
    const loadRules = async () => {
      try {
        const options: RuleOption[] = [
          {
            id: uuidv4(),
            label: 'Maintain Character Consistency',
            description: 'Ensure characters behave consistently.',
          },
          // Add more rules
        ];
        setRulesOptions(options);
      } catch (error) {
        console.error('Error loading rules:', error);
      }
    };

    // Load Characters
    const loadCharacters = async () => {
      try {
        const options: CharacterOption[] = [
          {
            id: uuidv4(),
            label: 'John Doe',
            description: 'A brave hero with a mysterious past.',
          },
          // Add more characters
        ];
        setCharacterOptions(options);
      } catch (error) {
        console.error('Error loading characters:', error);
      }
    };

    // Load Locations
    const loadLocations = async () => {
      try {
        const options: LocationOption[] = [
          {
            id: uuidv4(),
            label: 'Enchanted Forest',
            description: 'A mystical forest filled with magical creatures.',
          },
          // Add more locations
        ];
        setLocationOptions(options);
      } catch (error) {
        console.error('Error loading locations:', error);
      }
    };

    // Load Codex Entries
    const loadCodexEntries = async () => {
      try {
        const options: CodexOption[] = [
          {
            id: uuidv4(),
            label: 'Magic System',
            description: 'Explanation of how magic works in the world.',
          },
          // Add more codex entries
        ];
        setCodexOptions(options);
      } catch (error) {
        console.error('Error loading codex entries:', error);
      }
    };

    // Call all load functions
    loadSampleChapters();
    loadTaskTypes();
    loadRules();
    loadCharacters();
    loadLocations();
    loadCodexEntries();
  }, []);

  // Helper functions to get descriptions
  const getSampleChapterContent = (label: string): string => {
    const option = sampleChapterOptions.find((opt) => opt.label === label);
    return option ? option.content : '';
  };

  const getTaskTypeDescription = (label: string): string => {
    const option = taskTypeOptions.find((opt) => opt.label === label);
    return option ? option.description : '';
  };

  const getRuleDescription = (label: string): string => {
    const option = rulesOptions.find((opt) => opt.label === label);
    return option ? option.description : '';
  };

  const getCharacterDescription = (label: string): string => {
    const option = characterOptions.find((opt) => opt.label === label);
    return option ? option.description : '';
  };

  const getLocationDescription = (label: string): string => {
    const option = locationOptions.find((opt) => opt.label === label);
    return option ? option.description : '';
  };

  const getCodexDescription = (label: string): string => {
    const option = codexOptions.find((opt) => opt.label === label);
    return option ? option.description : '';
  };

  const calculateTokenCount = (text: string): number => {
    // Approximate token count; for more accurate results, integrate a tokenizer
    return text.trim().split(/\s+/).length;
  };

  // Generate prompt for ChatGPT
  const generateChatGPTPrompt = useCallback(() => {
    let prompt = '';

    // Add Sample Chapter
    if (sampleChapter) {
      const chapterContent = getSampleChapterContent(sampleChapter);
      prompt += `Sample Chapter:\n${chapterContent}\n\n`;
    }

    // Add Task Description
    if (taskTypeChecked && taskType) {
      const taskTypeDescription = getTaskTypeDescription(taskType);
      prompt += `Task:\n${taskTypeDescription}\n\n`;
    }

    // Add Rules
    if (selectedRules.length > 0) {
      const rulesDescriptions = selectedRules
        .map((label) => getRuleDescription(label))
        .join('\n');
      prompt += `Rules:\n${rulesDescriptions}\n\n`;
    }

    // Add Beats
    if (beats.trim() !== '') {
      prompt += `Beats:\n${beats}\n\n`;
    }

    // Add Characters
    if (selectedCharacters.length > 0) {
      const charactersDescriptions = selectedCharacters
        .map((label) => getCharacterDescription(label))
        .join('\n');
      prompt += `Characters:\n${charactersDescriptions}\n\n`;
    }

    // Add Locations
    if (selectedLocations.length > 0) {
      const locationsDescriptions = selectedLocations
        .map((label) => getLocationDescription(label))
        .join('\n');
      prompt += `Locations:\n${locationsDescriptions}\n\n`;
    }

    // Add Codex
    if (selectedCodexEntries.length > 0) {
      const codexDescriptions = selectedCodexEntries
        .map((label) => getCodexDescription(label))
        .join('\n');
      prompt += `Codex:\n${codexDescriptions}\n\n`;
    }

    // Append Raw Prompt (Task Instruction)
    prompt += rawPrompt;

    setFinalPrompt(prompt);
    setTokenCount(calculateTokenCount(prompt));
  }, [
    sampleChapter,
    sampleChapterOptions,
    taskType,
    taskTypeChecked,
    selectedRules,
    rulesOptions,
    beats,
    selectedCharacters,
    characterOptions,
    selectedLocations,
    locationOptions,
    selectedCodexEntries,
    codexOptions,
    rawPrompt,
  ]);

  // Generate prompt for Claude
  const generateClaudePrompt = useCallback(() => {
    let prompt = '';

    // Add Sample Chapter
    if (sampleChapter) {
      const chapterContent = getSampleChapterContent(sampleChapter);
      prompt += `<SAMPLE_CHAPTER>\n${chapterContent}\n</SAMPLE_CHAPTER>\n\n`;
    }

    // Add Task Description
    if (taskTypeChecked && taskType) {
      const taskTypeDescription = getTaskTypeDescription(taskType);
      prompt += `<TASK>\n${taskTypeDescription}\n</TASK>\n\n`;
    }

    // Add Rules
    if (selectedRules.length > 0) {
      const rulesDescriptions = selectedRules
        .map((label) => getRuleDescription(label))
        .join('\n');
      prompt += `<RULES>\n${rulesDescriptions}\n</RULES>\n\n`;
    }

    // Add Beats
    if (beats.trim() !== '') {
      prompt += `<BEATS>\n${beats}\n</BEATS>\n\n`;
    }

    // Add Characters
    if (selectedCharacters.length > 0) {
      const charactersDescriptions = selectedCharacters
        .map((label) => getCharacterDescription(label))
        .join('\n');
      prompt += `<CHARACTERS>\n${charactersDescriptions}\n</CHARACTERS>\n\n`;
    }

    // Add Locations
    if (selectedLocations.length > 0) {
      const locationsDescriptions = selectedLocations
        .map((label) => getLocationDescription(label))
        .join('\n');
      prompt += `<LOCATIONS>\n${locationsDescriptions}\n</LOCATIONS>\n\n`;
    }

    // Add Codex
    if (selectedCodexEntries.length > 0) {
      const codexDescriptions = selectedCodexEntries
        .map((label) => getCodexDescription(label))
        .join('\n');
      prompt += `<CODEX>\n${codexDescriptions}\n</CODEX>\n\n`;
    }

    // Append Raw Prompt (Task Instruction)
    prompt += rawPrompt;

    setFinalPrompt(prompt);
    setTokenCount(calculateTokenCount(prompt));
  }, [
    sampleChapter,
    sampleChapterOptions,
    taskType,
    taskTypeChecked,
    selectedRules,
    rulesOptions,
    beats,
    selectedCharacters,
    characterOptions,
    selectedLocations,
    locationOptions,
    selectedCodexEntries,
    codexOptions,
    rawPrompt,
  ]);

  // useEffect to handle prompt generation and default task instruction
  useEffect(() => {
    // Update task instruction if it matches the default of the previous prompt type
    const prevDefaultInstruction =
      prevPromptType.current === 'ChatGPT'
        ? DEFAULT_CHATGPT_INSTRUCTION
        : DEFAULT_CLAUDE_INSTRUCTION;

    const currentDefaultInstruction =
      currentPromptType === 'ChatGPT'
        ? DEFAULT_CHATGPT_INSTRUCTION
        : DEFAULT_CLAUDE_INSTRUCTION;

    if (rawPrompt.trim() === '' || rawPrompt === prevDefaultInstruction) {
      setRawPrompt(currentDefaultInstruction);
    }

    // Generate the prompt based on the current prompt type
    if (currentPromptType === 'ChatGPT') {
      generateChatGPTPrompt();
    } else if (currentPromptType === 'Claude') {
      generateClaudePrompt();
    }

    // Update the previous prompt type
    prevPromptType.current = currentPromptType;
  }, [
    currentPromptType,
    sampleChapter,
    sampleChapterOptions,
    taskType,
    taskTypeChecked,
    selectedRules,
    rulesOptions,
    beats,
    selectedCharacters,
    characterOptions,
    selectedLocations,
    locationOptions,
    selectedCodexEntries,
    codexOptions,
    rawPrompt,
    generateChatGPTPrompt,
    generateClaudePrompt,
  ]);

  // Handle Copy
  const handleCopy = () => {
    navigator.clipboard
      .writeText(finalPrompt)
      .then(() => console.log('Prompt copied to clipboard'))
      .catch((err) => console.error('Failed to copy prompt: ', err));
  };

  // Handle Generate Buttons
  const handleGenerateChatGPT = () => {
    setCurrentPromptType('ChatGPT');
  };

  const handleGenerateClaude = () => {
    setCurrentPromptType('Claude');
  };

  // Handle Settings Click
  const handleSettingsClick = () => {
    // Handle settings modal
  };

  return (
    <div className="container mx-auto p-4 space-y-4">
      <Header onSettingsClick={handleSettingsClick} />

      {/* Sample Chapters */}
      <SampleChaptersSelector
        previousChapter={previousChapter}
        onPreviousChapterChange={setPreviousChapter}
        sampleChapter={sampleChapter}
        onSampleChapterChange={setSampleChapter}
        nextChapterBeats={nextChapterBeats}
        onNextChapterBeatsChange={setNextChapterBeats}
        options={sampleChapterOptions}
        onEditClick={() => setIsSampleChapterEditOpen(true)}
      />
      {/* Task Type */}
      <TaskTypeSelector
        value={taskType}
        onChange={setTaskType}
        checked={taskTypeChecked}
        onCheckedChange={(checked: boolean) => setTaskTypeChecked(!!checked)}
        onEditClick={() => setIsTaskTypeEditOpen(true)}
        options={taskTypeOptions}
      />

      {/* Rules */}
      <RulesSelector
        values={selectedRules}
        onChange={setSelectedRules}
        options={rulesOptions}
        onEditClick={() => setIsRulesEditOpen(true)}
      />

      {/* Beats */}
      <BeatsInput value={beats} onChange={setBeats} />

      {/* Characters */}
      <CharactersSelector
        values={selectedCharacters}
        onChange={setSelectedCharacters}
        options={characterOptions}
        onEditClick={() => setIsCharactersEditOpen(true)}
      />

      {/* Locations */}
      <LocationsSelector
        values={selectedLocations}
        onChange={setSelectedLocations}
        options={locationOptions}
        onEditClick={() => setIsLocationsEditOpen(true)}
      />

      {/* Codex */}
      <CodexSelector
        values={selectedCodexEntries}
        onChange={setSelectedCodexEntries}
        options={codexOptions}
        onEditClick={() => setIsCodexEditOpen(true)}
      />

      {/* Raw Prompt */}
      <RawPrompt value={rawPrompt} onChange={setRawPrompt} />

      {/* Final Prompt */}
      <FinalPrompt
        value={finalPrompt}
        tokenCount={tokenCount}
        onChange={(value) => {
          setFinalPrompt(value);
          setTokenCount(calculateTokenCount(value));
        }}
      />

      {/* Action Buttons */}
      <ActionButtons
        onCopy={handleCopy}
        onGenerateChatGPT={handleGenerateChatGPT}
        onGenerateClaude={handleGenerateClaude}
      />

      {/* Modals for editing options */}
      <SampleChapterEditModal
        isOpen={isSampleChapterEditOpen}
        onClose={() => setIsSampleChapterEditOpen(false)}
        options={sampleChapterOptions}
        onSave={setSampleChapterOptions}
      />

      <TaskTypeEditModal
        isOpen={isTaskTypeEditOpen}
        onClose={() => setIsTaskTypeEditOpen(false)}
        options={taskTypeOptions}
        onSave={setTaskTypeOptions}
      />

      <RulesEditModal
        isOpen={isRulesEditOpen}
        onClose={() => setIsRulesEditOpen(false)}
        options={rulesOptions}
        onSave={setRulesOptions}
      />

      <CharactersEditModal
        isOpen={isCharactersEditOpen}
        onClose={() => setIsCharactersEditOpen(false)}
        options={characterOptions}
        onSave={setCharacterOptions}
      />

      <LocationsEditModal
        isOpen={isLocationsEditOpen}
        onClose={() => setIsLocationsEditOpen(false)}
        options={locationOptions}
        onSave={setLocationOptions}
      />

      <CodexEditModal
        isOpen={isCodexEditOpen}
        onClose={() => setIsCodexEditOpen(false)}
        options={codexOptions}
        onSave={setCodexOptions}
      />

      <SettingsModal isOpen={isSettingsOpen} onClose={() => setIsSettingsOpen(false)} />
    </div>
  );
}

export default App;
