import type { ProseImprovementPrompt } from '../../../frontend/src/types';

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

Example:
\`\`\`json
[
  {
    "initial": "Her eyes sparkled with unshed tears as she felt a surge of sadness.",
    "improved": "She bit her lip, turning away to hide the tremor in her voice. \\"I... I can't.\\"",
    "reason": "Replaced eye-based emotion and vague 'surge of sadness' with a specific action (bit lip, turning away) and dialogue showing distress, making the emotion more tangible and less reliant on clichés.",
    "trope_category": "Eye-based emotional indicators"
  }
]
\`\`\`

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
- trope_category: "dialogue"

Example:
\`\`\`json
[
  {
    "initial": "\\"As you know, the artifact is hidden in the Dragon's Peak, guarded by a magical seal that only the royal bloodline can break,\\" she explained.",
    "improved": "\\"The Dragon's Peak then?\\" he asked, his hand instinctively going to the hilt of his sword. She nodded, her gaze distant. \\"Only our blood can break the seal.\\"",
    "reason": "Removed direct exposition. The improved version implies the information through character actions and more natural, less 'as-you-know' dialogue, adding a touch of subtext with his hand on the sword.",
    "trope_category": "dialogue"
  }
]
\`\`\``,
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
- trope_category: "description"

Example:
\`\`\`json
[
  {
    "initial": "The forest was dark and scary. The trees were tall.",
    "improved": "Ancient oaks, their branches like skeletal fingers, clawed at a bruised twilight sky. A cold wind whispered through the gnarled roots, carrying the scent of damp earth and something else... something predatory.",
    "reason": "Replaced generic 'dark and scary' and 'tall trees' with specific sensory details (skeletal fingers, bruised sky, cold wind, scent of damp earth, predatory scent) and active voice to create a more vivid and immersive atmosphere.",
    "trope_category": "description"
  }
]
\`\`\``,
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
- trope_category: "pacing"

Example:
\`\`\`json
[
  {
    "initial": "She ran very quickly for a long time down the endless, winding, dark corridor, her heart pounding like a drum in her chest because she was so scared of the monster that she knew was chasing her from behind.",
    "improved": "She sprinted down the winding corridor, shadows clawing at her heels. Each gasp echoed the monster's nearing growl. Heart hammering, she risked a glance back—nothing but deeper darkness.",
    "reason": "Shortened the overly long sentence, used stronger verbs ('sprinted', 'clawing'), and broke the information into more digestible, active phrases. This increases tension and quickens the reading pace to match the action.",
    "trope_category": "pacing"
  }
]
\`\`\``,
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
- trope_category: "grammar"

Example:
\`\`\`json
[
  {
    "initial": "The cat, it was black and fluffy, and it sat on the mat, it seemed to be very happy.",
    "improved": "The black, fluffy cat sat contentedly on the mat.",
    "reason": "Corrected pronoun redundancy ('it was'), removed filter word ('seemed to be'), and combined descriptive elements for a more concise and clear sentence. 'Contentedly' directly conveys the cat's state.",
    "trope_category": "grammar"
  }
]
\`\`\``,
    category: 'grammar',
    order: 4
  }
];