// frontend/src/hooks/useProseImprovement.ts
import { useState, useCallback, useEffect } from 'react';
import type { ProseImprovementPrompt, ProseChange } from '@/types';
import { DEFAULT_PROSE_IMPROVEMENT_PROMPTS } from '@/utils/constants';

export function useProseImprovement() {
  const [prompts, setPrompts] = useState<ProseImprovementPrompt[]>(() => {
    const saved = localStorage.getItem('prose-improvement-prompts');
    return saved ? JSON.parse(saved) : DEFAULT_PROSE_IMPROVEMENT_PROMPTS;
  });

  useEffect(() => {
    localStorage.setItem('prose-improvement-prompts', JSON.stringify(prompts));
  }, [prompts]);

  const parseChanges = useCallback((response: string): ProseChange[] => {
    try {
      console.log("Attempting to parse LLM response:", response); // Log input to parseChanges
      // Look for JSON array in the response
      const jsonMatch = response.match(/\[[\s\S]*\]/);
      if (!jsonMatch) {
        console.error('No JSON array found in response string:', response);
        throw new Error('No JSON array found in response');
      }

      const parsed = JSON.parse(jsonMatch[0]);
          console.log("Successfully parsed JSON from response:", parsed); // Log the parsed JSON
          
      return parsed.map((item: any, index: number) => { // Added index for logging
        const changeItem = {
        id: crypto.randomUUID(),
        initial: item.weak_verb || item.original_verb || item.original_sentence || item.initial || item.original || '', // Prioritize weak_verb, then original_sentence
        improved: item.new_verb || item.improved || item.replacement || '', // Prioritize new_verb
        reason: item.reasoning || item.reason || item.explanation || '', // Prioritize reasoning
        trope_category: item.trope_category || item.category,
        status: 'pending'
        };
        console.log(`Mapping item ${index}:`, item, "to ProseChange:", changeItem); // Log each item and its mapping
        return changeItem;
      });
    } catch (error) {
      console.error('Error parsing changes:', error, "Original response string:", response); // Ensure original response is logged on error
      return [];
    }
  }, []);

  const applyChanges = useCallback((text: string, changes: ProseChange[]): string => {
    let result = text;
    
    // Sort changes by position (if available) to apply from end to start
    const sortedChanges = [...changes].sort((a, b) => {
      if (a.startIndex === undefined || b.startIndex === undefined) return 0;
      return b.startIndex - a.startIndex;
    });

    for (const change of sortedChanges) {
      if (change.status !== 'accepted') continue;
      
      // Simple replacement if no position info
      if (change.startIndex === undefined) {
        result = result.replace(change.initial, change.improved);
      } else {
        // Position-based replacement
        result = 
          result.slice(0, change.startIndex) +
          change.improved +
          result.slice(change.endIndex || change.startIndex + change.initial.length);
      }
    }

    return result;
  }, []);

  const findChangePositions = useCallback((text: string, changes: ProseChange[]): ProseChange[] => {
    return changes.map(change => {
      const startIndex = text.indexOf(change.initial);
      if (startIndex === -1) {
        return change;
      }
      
      return {
        ...change,
        startIndex,
        endIndex: startIndex + change.initial.length
      };
    });
  }, []);

  return {
    prompts,
    updatePrompts: setPrompts,
    parseChanges,
    applyChanges,
    findChangePositions
  };
}