// utils/constants.ts
import type { ProseImprovementPrompt } from '../types';

export const PROMPT_TYPES = {
    CHATGPT: 'ChatGPT',
    CLAUDE: 'Claude'
  } as const;
  
  export type PromptType = keyof typeof PROMPT_TYPES;
  
  // Default instructions for each AI model
  export const DEFAULT_INSTRUCTIONS = {
    [PROMPT_TYPES.CHATGPT]: `You are a creative writer tasked with composing the next chapter based on the provided context and requirements. Please follow these guidelines:
  
  1. Maintain consistency with previous chapters
  2. Follow the provided story beats
  3. Stay true to character voices and personalities
  4. Keep the established tone and style
  5. Incorporate relevant worldbuilding elements from the codex
  6. Respect all provided rules and constraints
  
  Your task is to write the next chapter that seamlessly continues the story while incorporating all the elements specified above.`,
  
    [PROMPT_TYPES.CLAUDE]: `<instructions>
  You are a creative writer tasked with composing the next chapter based on the provided context and requirements.
  
  <guidelines>
  1. Maintain consistency with the <previous_chapter>
  2. Follow the provided <beats> precisely
  3. Stay true to character voices defined in <characters>
  4. Keep the established tone and style
  5. Incorporate relevant worldbuilding from <codex>
  6. Respect all <rules> strictly
  </guidelines>
  
  Your task is to write the next chapter that seamlessly continues the story while incorporating all the elements above.
  </instructions>`,
  } as const;

  export const DEFAULT_TASK_INSTRUCTION = DEFAULT_INSTRUCTIONS[PROMPT_TYPES.CHATGPT];

  export const DEFAULT_CLAUDE_INSTRUCTION = DEFAULT_INSTRUCTIONS[PROMPT_TYPES.CLAUDE];
  
  // Token limits for different models
  export const TOKEN_LIMITS = {
    [PROMPT_TYPES.CHATGPT]: 4096,
    [PROMPT_TYPES.CLAUDE]: 100000,
  } as const;
  
  // Default placeholder texts
  export const PLACEHOLDERS = {
    PREVIOUS_CHAPTER: "Paste the content of the previous chapter here...",
    FUTURE_NOTES: "Add any notes about planned future chapters, plot points, or developments...",
    BEATS: "Enter the main story beats for the next chapter...",
    RAW_PROMPT: "Enter any additional instructions or requirements for the AI...",
  } as const;
  
  // File size limits
  export const FILE_SIZE_LIMITS = {
    MAX_CHAPTER_SIZE: 500 * 1024, // 500KB
    MAX_TOTAL_SIZE: 2 * 1024 * 1024, // 2MB
  } as const;
  
  // Local storage keys
  export const STORAGE_KEYS = {
    SETTINGS: 'story-generator-settings',
    TASK_TYPES: 'story-generator-task-types',
    RULES: 'story-generator-rules',
    CHARACTERS: 'story-generator-characters',
    LOCATIONS: 'story-generator-locations',
    CODEX: 'story-generator-codex',
    SAMPLE_CHAPTERS: 'story-generator-sample-chapters',
    LAST_PROMPT_TYPE: 'story-generator-last-prompt-type',
  } as const;
  
  // XML tags for Claude
  export const CLAUDE_TAGS = {
    PROMPT: 'prompt',
    TASK: 'task',
    PREVIOUS_CHAPTER: 'previous_chapter',
    SAMPLE_CHAPTER: 'sample_chapter',
    FUTURE_CHAPTERS: 'future_chapters',
    BEATS: 'beats',
    RULES: 'rules',
    CHARACTERS: 'characters',
    LOCATIONS: 'locations',
    CODEX: 'codex',
    INSTRUCTION: 'instruction',
  } as const;

  export const DEFAULT_PROSE_IMPROVEMENT_PROMPTS: readonly ProseImprovementPrompt[] = [
    {
      id: '1',
      label: 'Enhance Imagery',
      prompt: 'Review the following text and enhance the imagery. Focus on sensory details (sight, sound, smell, taste, touch) to make the descriptions more vivid and immersive. Provide changes in the specified JSON format.',
      order: 1,
      category: 'style',
    },
    {
      id: '2',
      label: 'Strengthen Verbs',
      prompt: 'Identify weak verbs (e.g., is, was, have, go) in the text and replace them with stronger, more active verbs. Explain the reasoning for each change. Provide changes in the specified JSON format.',
      order: 2,
      category: 'style',
    },
    {
      id: '3',
      label: 'Check for Clichés',
      prompt: 'Scan the text for common clichés or overused phrases. Suggest fresh alternatives for each one found. Provide changes in the specified JSON format.',
      order: 3,
      category: 'tropes',
    },
    {
      id: '4',
      label: 'Grammar and Punctuation',
      prompt: 'Perform a thorough grammar and punctuation check on the text. Correct any errors found. Provide changes in the specified JSON format.',
      order: 4,
      category: 'grammar',
    },
  ] as const;
  
  // Default task types
  export const DEFAULT_TASK_TYPES = [
    {
      id: 'write-chapter',
      label: 'Write Next Chapter',
      description: 'Compose the next chapter of the story following the provided context and beats.',
    },
    {
      id: 'revise-chapter',
      label: 'Revise Chapter',
      description: 'Revise and improve an existing chapter while maintaining consistency.',
    },
    {
      id: 'expand-scene',
      label: 'Expand Scene',
      description: 'Expand and elaborate on a specific scene with more detail and depth.',
    },
  ] as const;
  
  // Default rules
  export const DEFAULT_RULES = [
    {
      id: 'consistency',
      label: 'Maintain Consistency',
      description: 'Ensure all events, character behaviors, and world elements remain consistent with previous chapters.',
    },
    {
      id: 'show-dont-tell',
      label: 'Show, Don\'t Tell',
      description: 'Present story elements through action and dialogue rather than exposition.',
    },
    {
      id: 'pacing',
      label: 'Maintain Pacing',
      description: 'Keep the story moving at an engaging pace while allowing important moments to breathe.',
    },
  ] as const;
  
  // Error messages
  export const ERROR_MESSAGES = {
    INVALID_FILE_TYPE: 'Invalid file type. Please upload a text file.',
    FILE_TOO_LARGE: 'File is too large. Please upload a smaller file.',
    FAILED_TO_COPY: 'Failed to copy to clipboard. Please try again.',
    FAILED_TO_SAVE: 'Failed to save changes. Please try again.',
    REQUIRED_FIELD: 'This field is required.',
    INVALID_INPUT: 'Invalid input. Please check your entries.',
  } as const;
  
  // Success messages
  export const SUCCESS_MESSAGES = {
    COPIED_TO_CLIPBOARD: 'Successfully copied to clipboard!',
    CHANGES_SAVED: 'Changes saved successfully!',
    FILE_UPLOADED: 'File uploaded successfully!',
  } as const;
  
  // Default export for convenience
  export default {
    PROMPT_TYPES,
    DEFAULT_INSTRUCTIONS,
    DEFAULT_TASK_INSTRUCTION,  
    DEFAULT_CLAUDE_INSTRUCTION,  
    TOKEN_LIMITS,
    PLACEHOLDERS,
    FILE_SIZE_LIMITS,
    STORAGE_KEYS,
    CLAUDE_TAGS,
    DEFAULT_TASK_TYPES,
    DEFAULT_RULES,
    ERROR_MESSAGES,
    SUCCESS_MESSAGES,
  };