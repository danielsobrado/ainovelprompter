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
  } as const;  export const DEFAULT_PROSE_IMPROVEMENT_PROMPTS: ProseImprovementPrompt[] = [
    {
      id: 'enhance-imagery',
      label: 'Enhance Imagery',
      category: 'style',
      order: 1,
      description: 'Identifies specific phrases and suggests more vivid alternatives with sensory details.',
      defaultPromptText: `Review the following text and enhance the imagery. Focus on sensory details (sight, sound, smell, taste, touch) to make the descriptions more vivid and immersive.

REQUIRED JSON FORMAT: Your response must be a valid JSON array where each object has these exact keys:
- "original_text": The original text that needs enhancement
- "enhanced_text": The improved version with enhanced imagery  
- "reason": Explanation of what was enhanced and why

EXAMPLE RESPONSE:
[
  {
    "original_text": "The forest was quiet.",
    "enhanced_text": "The ancient forest stood in hushed reverence, the only sounds the whisper of wind through pine needles and the distant caw of a crow.",
    "reason": "Added specific auditory and visual details to create a more immersive and evocative scene."
  }
]

CRITICAL REQUIREMENTS:
- Your entire response MUST be a single, valid JSON array starting with '[' and ending with ']'
- Do NOT include any text before '[' or after ']'
- If no enhancements are needed, return: []
- Each object must have all three keys: "original_text", "enhanced_text", "reason"

Text to analyze:
[TEXT_TO_ANALYZE_PLACEHOLDER]`,
      variants: [
        {
          variantLabel: "Claude Optimized",
          targetModelFamilies: ["claude", "anthropic"],
          promptText: `<instructions>
You are a master of prose enhancement. For the following text, identify segments that can benefit from enhanced imagery and sensory details.

Return only a JSON array of objects, each with "original_text", "enhanced_text", and "reason" keys.

Example response format:
[
  {
    "original_text": "It was dark.",
    "enhanced_text": "Shadows danced across the moonlit ground, casting shifting patterns of silver and black.",
    "reason": "Replaced vague 'dark' with specific visual imagery of shadows and moonlight."
  }
]
</instructions>

<document_to_analyze>
[TEXT_TO_ANALYZE_PLACEHOLDER]
</document_to_analyze>`
        }
      ]
    },
    {
      id: 'strengthen-verbs',
      label: 'Strengthen Verbs',
      category: 'style',
      order: 2,
      description: 'Replaces weak verbs with stronger, more active ones.',
      defaultPromptText: `Identify weak verbs (e.g., is, was, have, go, get, make, do) in the provided text and replace them with stronger, more active verbs.

REQUIRED JSON FORMAT: Your response must be a valid JSON array where each object has these exact keys:
- "original_text": The original text containing the weak verb
- "improved_text": The text with the stronger verb replacement
- "reason": Explanation of why this verb is stronger

EXAMPLE RESPONSE:
[
  {
    "original_text": "The house was big.",
    "improved_text": "The house loomed.",
    "reason": "Replaced 'was big' with 'loomed' to provide a stronger visual and sense of imposing size."
  }
]

CRITICAL REQUIREMENTS:
- Your entire response MUST be a single, valid JSON array starting with '[' and ending with ']'
- Do NOT include any text before '[' or after ']'
- If no weak verbs need strengthening, return: []
- Each object must have all three keys: "original_text", "improved_text", "reason"

Text to analyze:
[TEXT_TO_ANALYZE_PLACEHOLDER]`,
      variants: []
    },
    {
      id: 'enhance-descriptions',
      label: 'Enhance Descriptions',
      category: 'style',
      order: 3,
      description: 'Makes descriptions more vivid, specific, and engaging.',
      defaultPromptText: `Review the text and enhance its descriptions. Focus on making them more vivid, specific, and engaging by elaborating on existing descriptions, adding sensory details, or using stronger imagery.

REQUIRED JSON FORMAT: Your response must be a valid JSON array where each object has these exact keys:
- "original_text": The original descriptive text
- "enhanced_text": The improved, more vivid description
- "reason": Explanation of what makes the new description better

EXAMPLE RESPONSE:
[
  {
    "original_text": "The car was red.",
    "enhanced_text": "The cherry-red convertible gleamed under the afternoon sun, its polished surface reflecting the azure sky.",
    "reason": "Added specificity (convertible, cherry-red), visual imagery (gleamed, reflecting), and environmental details (afternoon sun, azure sky)."
  }
]

CRITICAL REQUIREMENTS:
- Your entire response MUST be a single, valid JSON array starting with '[' and ending with ']'
- Do NOT include any text before '[' or after ']'
- If no descriptions need enhancement, return: []
- Each object must have all three keys: "original_text", "enhanced_text", "reason"

Text to analyze:
[TEXT_TO_ANALYZE_PLACEHOLDER]`,
      variants: []
    },
    {
      id: 'grammar-punctuation',
      label: 'Grammar and Punctuation',
      category: 'grammar',
      order: 4,
      description: 'Corrects grammar, punctuation, and spelling errors.',
      defaultPromptText: `Perform a thorough grammar and punctuation check on the text. Correct any errors found including spelling, punctuation, sentence structure, and grammatical mistakes.

REQUIRED JSON FORMAT: Your response must be a valid JSON array where each object has these exact keys:
- "original_text": The original text with errors
- "corrected_text": The corrected version
- "reason": Explanation of what grammar/punctuation issues were fixed

EXAMPLE RESPONSE:
[
  {
    "original_text": "Its a nice day isnt it.",
    "corrected_text": "It's a nice day, isn't it?",
    "reason": "Added apostrophe in 'It's', comma before tag question 'isn't it', and question mark at end."
  }
]

CRITICAL REQUIREMENTS:
- Your entire response MUST be a single, valid JSON array starting with '[' and ending with ']'
- Do NOT include any text before '[' or after ']'
- If no grammar/punctuation errors are found, return: []
- Each object must have all three keys: "original_text", "corrected_text", "reason"

Text to analyze:
[TEXT_TO_ANALYZE_PLACEHOLDER]`,      variants: []
    },
  ];
  
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