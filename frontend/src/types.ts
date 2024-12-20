// types.ts

/**
 * Base interface for all option types
 */
export interface BaseOption {
    id: string;
    label: string;
    description: string;
  }
  
  /**
   * Sample Chapter Option
   */
  export interface SampleChapterOption extends Omit<BaseOption, 'description'> {
    content: string;
  }
  
  /**
   * Task Type Option
   */
  export interface TaskTypeOption extends BaseOption {}
  
  /**
   * Rule Option
   */
  export interface RuleOption extends BaseOption {}
  
  /**
   * Character Option
   */
  export interface CharacterOption extends BaseOption {}
  
  /**
   * Location Option
   */
  export interface LocationOption extends BaseOption {}
  
  /**
   * Codex Option
   */
  export interface CodexOption extends BaseOption {}
  
  /**
   * Prompt Data interface for generating prompts
   */
  export interface PromptData {
    taskType: string;
    taskTypeChecked: boolean;
    sampleChapter: string;
    previousChapterText: string;
    nextChapterBeats: string;
    futureChapterNotes: string;
    selectedRules: string[];
    selectedCharacters: string[];
    selectedLocations: string[];
    selectedCodexEntries: string[];
    rawPrompt: string;
  }
  
  /**
   * Prompt type enum
   */
  export type PromptType = 'ChatGPT' | 'Claude';
  
  /**
   * Common modal props for edit modals
   */
  export interface EditModalProps<T extends BaseOption> {
    isOpen: boolean;
    onClose: () => void;
    options: T[];
    onSave: (options: T[]) => void;
  }