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
export interface TaskTypeOption extends BaseOption { }

/**
 * Rule Option
 */
export interface RuleOption extends BaseOption { }

/**
 * Character Option
 */
export interface CharacterOption extends BaseOption { }

/**
 * Location Option
 */
export interface LocationOption extends BaseOption { }

/**
 * Codex Option
 */
export interface CodexOption extends BaseOption { }

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

/**
 * Prose Improvement Prompt interface
 */
export interface ProsePromptVariant {
  variantLabel?: string;
  targetModelFamilies?: string[]; // e.g., ["claude", "gpt", "ollama"]
  targetModels?: string[];      // e.g., ["anthropic/claude-3-opus", "openai/gpt-4o"]
  promptText: string;
}

export interface ProseImprovementPrompt { // This now represents a "Prompt Definition"
  id: string; // e.g., "enhance-imagery"
  label: string; // User-friendly label, e.g., "Enhance Imagery"
  category: 'tropes' | 'style' | 'grammar' | 'custom';
  order: number;
  description?: string;
  defaultPromptText: string; // Fallback prompt
  variants: ProsePromptVariant[];
  // The 'prompt' field is removed as it's now split into defaultPromptText and variants
}

export interface ProseChange {
  id: string;
  initial: string;
  improved: string;
  reason: string;
  trope_category?: string;
  status: 'pending' | 'accepted' | 'rejected';
  startIndex?: number;
  endIndex?: number;
}

export interface ProseImprovementSession {
  id: string;
  originalText: string;
  currentText: string;
  prompts: ProseImprovementPrompt[];
  currentPromptIndex: number;
  changes: ProseChange[];
  createdAt: Date;
  updatedAt: Date;
}

export interface LLMProvider {
  type: 'manual' | 'lmstudio' | 'openrouter';
  config?: {
    apiUrl?: string;
    apiKey?: string;
    model?: string;
  };
}

/**
 * New Backend-Driven Prompt System Types
 */
export interface PromptVariant {
  variantLabel?: string;
  targetModelFamilies?: string[];
  targetModels?: string[];
  promptText: string;
}

export interface PromptDefinition {
  id: string;
  label: string;
  category: string;
  order: number;
  description?: string;
  defaultPromptText: string;
  variants?: PromptVariant[];
}