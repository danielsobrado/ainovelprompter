import type { PromptData } from './types';

export function generateChatGPTPrompt(data: PromptData): string {
  const {
    taskType,
    taskTypeChecked,
    sampleChapter,
    previousChapterText,
    nextChapterBeats,
    futureChapterNotes,
    selectedRules,
    selectedCharacters,
    selectedLocations,
    selectedCodexEntries,
    rawPrompt,
  } = data;

  let prompt = '';

  // Add Previous Chapter
  if (previousChapterText.trim()) {
    prompt += `Previous Chapter:\n${previousChapterText}\n\n`;
  }

  // Add Sample Chapter
  if (sampleChapter) {
    prompt += `Sample Chapter:\n${sampleChapter}\n\n`;
  }

  // Add Future Chapter Notes
  if (futureChapterNotes.trim()) {
    prompt += `Future Chapter Notes:\n${futureChapterNotes}\n\n`;
  }

  // Add Task Type
  if (taskTypeChecked && taskType) {
    prompt += `Task Type:\n${taskType}\n\n`;
  }

  // Add Next Chapter Beats
  if (nextChapterBeats.trim()) {
    prompt += `Next Chapter Beats:\n${nextChapterBeats}\n\n`;
  }

  // Add Rules
  if (selectedRules.length > 0) {
    prompt += `Rules:\n${selectedRules.join('\n')}\n\n`;
  }

  // Add Characters
  if (selectedCharacters.length > 0) {
    prompt += `Characters:\n${selectedCharacters.join('\n')}\n\n`;
  }

  // Add Locations
  if (selectedLocations.length > 0) {
    prompt += `Locations:\n${selectedLocations.join('\n')}\n\n`;
  }

  // Add Codex Entries
  if (selectedCodexEntries.length > 0) {
    prompt += `Codex Entries:\n${selectedCodexEntries.join('\n')}\n\n`;
  }

  // Add Raw Prompt
  if (rawPrompt.trim()) {
    prompt += rawPrompt;
  }

  return prompt.trim();
}

export function generateClaudePrompt(data: PromptData): string {
  const {
    taskType,
    taskTypeChecked,
    sampleChapter,
    previousChapterText,
    nextChapterBeats,
    futureChapterNotes,
    selectedRules,
    selectedCharacters,
    selectedLocations,
    selectedCodexEntries,
    rawPrompt,
  } = data;

  let prompt = '<prompt>\n';

  // Add Previous Chapter
  if (previousChapterText.trim()) {
    prompt += `<previous_chapter>\n${previousChapterText}\n</previous_chapter>\n\n`;
  }

  // Add Sample Chapter
  if (sampleChapter) {
    prompt += `<sample_chapter>\n${sampleChapter}\n</sample_chapter>\n\n`;
  }

  // Add Future Chapter Notes
  if (futureChapterNotes.trim()) {
    prompt += `<future_chapters>\n${futureChapterNotes}\n</future_chapters>\n\n`;
  }

  // Add Task Type
  if (taskTypeChecked && taskType) {
    prompt += `<task_type>${taskType}</task_type>\n\n`;
  }

  // Add Next Chapter Beats
  if (nextChapterBeats.trim()) {
    prompt += `<beats>\n${nextChapterBeats}\n</beats>\n\n`;
  }

  // Add Rules
  if (selectedRules.length > 0) {
    prompt += `<rules>\n${selectedRules.map(rule => `  <rule>${rule}</rule>`).join('\n')}\n</rules>\n\n`;
  }

  // Add Characters
  if (selectedCharacters.length > 0) {
    prompt += `<characters>\n${selectedCharacters.map(char => `  <character>${char}</character>`).join('\n')}\n</characters>\n\n`;
  }

  // Add Locations
  if (selectedLocations.length > 0) {
    prompt += `<locations>\n${selectedLocations.map(loc => `  <location>${loc}</location>`).join('\n')}\n</locations>\n\n`;
  }

  // Add Codex Entries
  if (selectedCodexEntries.length > 0) {
    prompt += `<codex>\n${selectedCodexEntries.map(entry => `  <entry>${entry}</entry>`).join('\n')}\n</codex>\n\n`;
  }

  // Add Raw Prompt
  if (rawPrompt.trim()) {
    prompt += `<instruction>\n${rawPrompt}\n</instruction>\n`;
  }

  prompt += '</prompt>';

  return prompt.trim();
}