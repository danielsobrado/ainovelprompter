import type { PromptData, TaskTypeOption } from '../types';

interface OptionWithDescription {
  label: string;
  description: string;
}

interface PromptOptions {
  rules: OptionWithDescription[];
  characters: OptionWithDescription[];
  locations: OptionWithDescription[];
  codex: OptionWithDescription[];
  taskTypes: TaskTypeOption[];
}

export function generateChatGPTPrompt(data: PromptData, options: PromptOptions): string {
  let prompt = '';

  // Add Previous Chapter
  if (data.previousChapterText.trim()) {
    prompt += `Previous Chapter:\n${data.previousChapterText}\n\n`;
  }

  // Add Sample Chapter
  if (data.sampleChapter) {
    prompt += `Sample Chapter:\n${data.sampleChapter}\n\n`;
  }

  // Add Future Chapter Notes
  if (data.futureChapterNotes.trim()) {
    prompt += `Future Chapter Notes:\n${data.futureChapterNotes}\n\n`;
  }

  // Add Next Chapter Beats
  if (data.nextChapterBeats.trim()) {
    prompt += `Next Chapter Beats:\n${data.nextChapterBeats}\n\n`;
  }

  // Add Rules with descriptions
  if (data.selectedRules.length > 0) {
    prompt += `Rules:\n${data.selectedRules.map(rule => {
      const ruleData = options.rules.find(r => r.label === rule);
      return ruleData ? ruleData.description : '';
    }).filter(Boolean).join('\n')}\n\n`;
  }

  // Add Characters with descriptions
  if (data.selectedCharacters.length > 0) {
    prompt += `Characters:\n${data.selectedCharacters.map(char => {
      const charData = options.characters.find(c => c.label === char);
      return charData ? charData.description : '';
    }).filter(Boolean).join('\n')}\n\n`;
  }

  // Add Locations with descriptions
  if (data.selectedLocations.length > 0) {
    prompt += `Locations:\n${data.selectedLocations.map(loc => {
      const locData = options.locations.find(l => l.label === loc);
      return locData ? locData.description : '';
    }).filter(Boolean).join('\n')}\n\n`;
  }

  // Add Codex Entries with descriptions
  if (data.selectedCodexEntries.length > 0) {
    prompt += `Codex Entries:\n${data.selectedCodexEntries.map(entry => {
      const codexData = options.codex.find(c => c.label === entry);
      return codexData ? codexData.description : '';
    }).filter(Boolean).join('\n')}\n\n`;
  }

  // Add Task Instructions
  if (data.taskTypeChecked && data.taskType) {
    // If checkbox is selected, use the task type description
    const taskTypeData = options.taskTypes.find(t => t.label === data.taskType);
    if (taskTypeData) {
      prompt += `Task:\n${taskTypeData.description}\n\n`;
    }
  } else if (data.rawPrompt.trim()) {
    // If checkbox is not selected, use the raw prompt
    prompt += `Task:\n${data.rawPrompt}\n\n`;
  }

  return prompt.trim();
}

export function generateClaudePrompt(data: PromptData, options: PromptOptions): string {
  let prompt = '';

  // Add Previous Chapter
  if (data.previousChapterText.trim()) {
    prompt += `<previous_chapter>\n${data.previousChapterText}\n</previous_chapter>\n\n`;
  }

  // Add Sample Chapter
  if (data.sampleChapter) {
    prompt += `<sample_chapter>\n${data.sampleChapter}\n</sample_chapter>\n\n`;
  }

  // Add Future Chapter Notes
  if (data.futureChapterNotes.trim()) {
    prompt += `<future_chapters>\n${data.futureChapterNotes}\n</future_chapters>\n\n`;
  }

  // Add Next Chapter Beats
  if (data.nextChapterBeats.trim()) {
    prompt += `<beats>\n${data.nextChapterBeats}\n</beats>\n\n`;
  }

  // Add Rules with descriptions
  if (data.selectedRules.length > 0) {
    prompt += `<rules>\n${data.selectedRules.map(rule => {
      const ruleData = options.rules.find(r => r.label === rule);
      return ruleData ? `  <rule>${ruleData.description}</rule>` : '';
    }).filter(Boolean).join('\n')}\n</rules>\n\n`;
  }

  // Add Characters with descriptions
  if (data.selectedCharacters.length > 0) {
    prompt += `<characters>\n${data.selectedCharacters.map(char => {
      const charData = options.characters.find(c => c.label === char);
      return charData ? `  <character>${charData.description}</character>` : '';
    }).filter(Boolean).join('\n')}\n</characters>\n\n`;
  }

  // Add Locations with descriptions
  if (data.selectedLocations.length > 0) {
    prompt += `<locations>\n${data.selectedLocations.map(loc => {
      const locData = options.locations.find(l => l.label === loc);
      return locData ? `  <location>${locData.description}</location>` : '';
    }).filter(Boolean).join('\n')}\n</locations>\n\n`;
  }

  // Add Codex Entries with descriptions
  if (data.selectedCodexEntries.length > 0) {
    prompt += `<codex>\n${data.selectedCodexEntries.map(entry => {
      const codexData = options.codex.find(c => c.label === entry);
      return codexData ? `  <entry>${codexData.description}</entry>` : '';
    }).filter(Boolean).join('\n')}\n</codex>\n\n`;
  }

  // Add Task Instructions
  if (data.taskTypeChecked && data.taskType) {
    // If checkbox is selected, use the task type description
    const taskTypeData = options.taskTypes.find(t => t.label === data.taskType);
    if (taskTypeData) {
      prompt += `<task>\n${taskTypeData.description}\n</task>\n`;
    }
  } else if (data.rawPrompt.trim()) {
    // If checkbox is not selected, use the raw prompt
    prompt += `<task>\n${data.rawPrompt}\n</task>\n`;
  }

  return prompt.trim();
}