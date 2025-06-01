// frontend/src/utils/constants.ts

export const DEFAULT_PROSE_IMPROVEMENT_PROMPTS: ProseImprovementPrompt[] = [
  {
    id: '1',
    label: 'Fix AI Tropes',
    prompt: `Task: Fix AI Writing Tropes in Text

Core AI Tropes to Eliminate:
1. Eye-based emotional indicators:
   - "eyes gleamed/sparkled/narrowed/softened"
   - "something in their eyes"
   - "eyes took on a [adjective] look"
   
2. Vague emotional descriptions:
   - "expression softened/hardened"
   - "felt a surge of [emotion]"
   - "couldn't help but [action]"
   
3. Mechanical thought indicators:
   - "gears turning"
   - "mind raced"
   - "thoughts swirled"
   
4. Filter words and unnecessary narration:
   - "could see/feel/hear"
   - "seemed to"
   - "appeared to be"
   
5. Overused transitions:
   - "after a long moment"
   - "fell silent"
   - "let out a breath they didn't know they were holding"

Replacement Strategies:
- Use specific physical actions instead of internal states
- Show emotion through dialogue and behavior, not description
- Cut filter words - just state what happens
- Replace clichés with fresh, specific details

Output Format:
JSON array with:
- initial: original text
- improved: revised text
- reason: specific explanation of why this change improves the text
- trope_category: which AI trope this addresses

Focus on:
- Precision over abstraction
- External actions over internal states
- Specific details over generic descriptions
- Natural dialogue over exposition`,
    category: 'tropes',
    order: 0
  },
  {
    id: '2',
    label: 'Improve Dialogue',
    prompt: `Task: Improve Dialogue Quality

Focus Areas:
1. Remove exposition from dialogue
2. Make each character's voice distinct
3. Add subtext and conflict
4. Remove on-the-nose dialogue
5. Use action beats instead of dialogue tags

Output Format:
JSON array with:
- initial: original dialogue
- improved: revised dialogue
- reason: explanation of improvement
- trope_category: "dialogue"`,
    category: 'style',
    order: 1
  },
  {
    id: '3',
    label: 'Enhance Descriptions',
    prompt: `Task: Enhance Descriptive Passages

Goals:
1. Replace generic descriptions with specific, sensory details
2. Remove purple prose and overwrought metaphors
3. Ground descriptions in character perspective
4. Use active voice
5. Show character emotions through environment

Output Format:
JSON array with:
- initial: original description
- improved: enhanced description
- reason: explanation of changes
- trope_category: "description"`,
    category: 'style',
    order: 2
  },
  {
    id: '4',
    label: 'Fix Pacing Issues',
    prompt: `Task: Identify and Fix Pacing Issues

Target Areas:
1. Overly long sentences that slow reading
2. Repetitive information
3. Unnecessary backstory or info-dumps
4. Rushed action sequences
5. Dragging introspection

Output Format:
JSON array with:
- initial: problematic passage
- improved: better-paced version
- reason: how this improves pacing
- trope_category: "pacing"`,
    category: 'style',
    order: 3
  },
  {
    id: '5',
    label: 'Grammar and Clarity',
    prompt: `Task: Fix Grammar and Improve Clarity

Focus on:
1. Correct grammatical errors
2. Eliminate awkward phrasing
3. Fix unclear pronoun references
4. Improve sentence structure
5. Remove redundancies

Output Format:
JSON array with:
- initial: text with issues
- improved: corrected text
- reason: what was fixed
- trope_category: "grammar"`,
    category: 'grammar',
    order: 4
  }
];