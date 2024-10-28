interface GenerateInstructionsOptions {
    previousChapter: string;
    beats: string;
    selectedCharacters: string[];
    selectedLocations: string[];
    selectedCodexEntries: string[];
    selectedRules: string[];
    isClaudeFormat: boolean;
  }
  
  export function generateDynamicInstructions({
    previousChapter,
    beats,
    selectedCharacters,
    selectedLocations,
    selectedCodexEntries,
    selectedRules,
    isClaudeFormat,
  }: GenerateInstructionsOptions): string {
    let instructions = `You are a creative writer tasked with composing the next chapter based on the provided context.`;
  
    // Add previous chapter reference if content exists
    if (previousChapter.trim()) {
      if (isClaudeFormat) {
        instructions += ` Maintain consistency with the content in <previous_chapter>.`;
      } else {
        instructions += ` Maintain consistency with the previous chapter.`;
      }
    }
  
    // Add beats reference if content exists
    if (beats.trim()) {
      if (isClaudeFormat) {
        instructions += ` Follow the story beats provided in <beats>.`;
      } else {
        instructions += ` Follow the provided story beats.`;
      }
    }
  
    // Add character references if selected
    if (selectedCharacters.length > 0) {
      if (isClaudeFormat) {
        instructions += `\n\nIncorporate the characters defined in <characters> with their established traits and backgrounds.`;
      } else {
        instructions += `\n\nIncorporate the provided characters with their established traits and backgrounds.`;
      }
    }
  
    // Add location references if selected
    if (selectedLocations.length > 0) {
      if (isClaudeFormat) {
        instructions += `\n\nUse the locations described in <locations> as your story settings.`;
      } else {
        instructions += `\n\nUse the provided locations as your story settings.`;
      }
    }
  
    // Add codex references if selected
    if (selectedCodexEntries.length > 0) {
      if (isClaudeFormat) {
        instructions += `\n\nIncorporate the world-building elements from <codex> to maintain setting consistency.`;
      } else {
        instructions += `\n\nIncorporate the provided codex entries to maintain setting consistency.`;
      }
    }
  
    // Add rules references if selected
    if (selectedRules.length > 0) {
      if (isClaudeFormat) {
        instructions += `\n\nFollow the writing guidelines specified in <rules>.`;
      } else {
        instructions += `\n\nFollow the provided writing guidelines.`;
      }
    }
  
    return instructions;
  }