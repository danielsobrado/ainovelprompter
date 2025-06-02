// frontend/src/hooks/useProseImprovement.ts
import { useState, useCallback, useEffect } from 'react';
import type { ProseImprovementPrompt, ProseChange } from '@/types';
import { DEFAULT_PROSE_IMPROVEMENT_PROMPTS } from '@/utils/constants';
import { useWailsReady } from '@/contexts/WailsReadyContext'; // Import the context hook
// Import Wails runtime for Go function calls if not already globally available
// For newer Wails, direct calls like window.go.main.App... are typical.
// Ensure your wailsjsdev.js or wailsjsruntime.js is loaded.

export function useProseImprovement() {
  const { wailsReady } = useWailsReady(); // Consume the context
  const [prompts, setPrompts] = useState<ProseImprovementPrompt[]>([]);
  const [isLoadingPrompts, setIsLoadingPrompts] = useState(true);

  useEffect(() => {
    const loadPrompts = async () => {
      setIsLoadingPrompts(true);
      try {
        // Check if Wails bindings are ready
        if (wailsReady && window.go && window.go.main && window.go.main.App) {
          const savedPromptsJson = await window.go.main.App.ReadProsePromptsFile();
          if (savedPromptsJson) {
            const savedPrompts = JSON.parse(savedPromptsJson);
            if (Array.isArray(savedPrompts) && savedPrompts.length > 0) {
              setPrompts(savedPrompts);
            } else {
              setPrompts([...DEFAULT_PROSE_IMPROVEMENT_PROMPTS]);
            }
          } else {
            setPrompts([...DEFAULT_PROSE_IMPROVEMENT_PROMPTS]);
          }
        } else if (!wailsReady) {
          console.log("useProseImprovement: Wails not ready, waiting...");
          // No need for setTimeout here, effect will re-run when wailsReady changes
        } else {
          console.warn("Wails Go bindings not available, using defaults for prompts.");
          setPrompts([...DEFAULT_PROSE_IMPROVEMENT_PROMPTS]);
        }
      } catch (error) {
        console.error("Error loading prose prompts, using defaults:", error);
        setPrompts([...DEFAULT_PROSE_IMPROVEMENT_PROMPTS]);
      } finally {
        // Only set loading to false if wails is ready or if we decided to use defaults
        if (wailsReady || !(window.go && window.go.main && window.go.main.App)) {
          setIsLoadingPrompts(false);
        }
      }
    };

    if (wailsReady) { // Only attempt to load if Wails is ready
      loadPrompts();
    } else {
      // If Wails is not ready yet, ensure prompts are default and not loading indefinitely
      // This handles the initial state before Wails signals readiness.
      setPrompts([...DEFAULT_PROSE_IMPROVEMENT_PROMPTS]);
      setIsLoadingPrompts(true); // Keep loading true until wailsReady or error
    }
  }, [wailsReady]); // Re-run when wailsReady changes

  const parseChanges = useCallback((response: string): ProseChange[] => {
    try {
      console.log("Attempting to parse LLM response:", response);
      let jsonString: string | null = null;

      // 1. Try to extract JSON from a markdown code block (```json ... ```
      let match = response.match(/```json\s*([\s\S]*?)\s*```/);
      if (match && match[1]) {
        jsonString = match[1].trim();
        console.log("Extracted JSON from markdown block:", jsonString);
      } else {
        // 2. If no markdown block, check for root array or object
        const trimmedResponse = response.trim();
        if (trimmedResponse.startsWith('[')) {
          const firstBracket = response.indexOf('[');
          const lastBracket = response.lastIndexOf(']');
          if (lastBracket > firstBracket) {
            jsonString = response.substring(firstBracket, lastBracket + 1).trim();
            console.log("Extracted root JSON array:", jsonString);
          } else {
            console.warn("Truncated JSON array detected. Aborting parse.");
            throw new Error('Truncated JSON array in response.');
          }
        } else if (trimmedResponse.startsWith('{')) {
          const firstBrace = response.indexOf('{');
          const lastBrace = response.lastIndexOf('}');
          if (lastBrace > firstBrace) {
            jsonString = response.substring(firstBrace, lastBrace + 1).trim();
            console.log("Extracted root JSON object:", jsonString);
          } else {
            console.warn("Truncated JSON object detected. Aborting parse.");
            throw new Error('Truncated JSON object in response.');
          }
        } else {
          // Fallback: try to find the first significant JSON structure
          const firstBracket = response.indexOf('[');
          const lastBracket = response.lastIndexOf(']');
          if (firstBracket !== -1 && lastBracket > firstBracket) {
            jsonString = response.substring(firstBracket, lastBracket + 1).trim();
            console.log("Extracted potential JSON array string (fallback):", jsonString);
          } else {
            const firstBrace = response.indexOf('{');
            const lastBrace = response.lastIndexOf('}');
            if (firstBrace !== -1 && lastBrace > firstBrace) {
              jsonString = response.substring(firstBrace, lastBrace + 1).trim();
              console.log("Extracted potential JSON object string (fallback):", jsonString);
            }
          }
        }
      }

      if (!jsonString) {
        console.error('No valid JSON object or array found within the response string after attempting extraction:', response);
        throw new Error('No valid JSON object or array found within the response string.');
      }

      const parsed = JSON.parse(jsonString); // This will still fail if jsonString is truncated
      console.log("Successfully parsed JSON from response:", parsed);

      let changesArray: any[] = [];

      if (Array.isArray(parsed)) {
        changesArray = parsed;
      } else if (typeof parsed === 'object' && parsed !== null && Array.isArray(parsed.changes)) {
        console.log("Detected 'changes' array within a root object.");
        changesArray = parsed.changes;
      } else if (typeof parsed === 'object' && parsed !== null && parsed.sections && Array.isArray(parsed.sections)) {
        console.log("Detected 'sections' array format. Processing sections as individual changes.");
        return parsed.sections.map((section: any, index: number) => {
          const changeItem: ProseChange = {
            id: crypto.randomUUID(),
            initial: section.original_text || `Original for: ${section.heading || `Section ${index + 1}`}` || " (Original section text not provided in this response format)",
            improved: section.text || '',
            reason: `Enhancements for section: ${section.heading || `Section ${index + 1}`}`,
            trope_category: parsed.title || 'imagery enhancement',
            status: 'pending' as 'pending',
          };
          console.log(`Mapping object section ${index}:`, section, "to ProseChange:", changeItem);
          return changeItem;
        });
      } else if (typeof parsed === 'object' && parsed !== null && (parsed.initial || parsed.original || parsed.original_verb || parsed.weak_verb || parsed.revised || parsed.new_verb)) {
        console.log("Detected single change object format (not in an array). Wrapping in array.");
        changesArray = [parsed];
      } else {
        console.error("Parsed JSON is not an array or a known object structure. Parsed data:", parsed, "Original response:", response);
        throw new Error('Parsed JSON is not an array or a known object structure.');
      }

      // Common mapping logic for array-based changes
      return changesArray.map((item: any, index: number) => {
        const changeItem = {
          id: crypto.randomUUID(),
          initial: item.original_text || item.original_verb || item.original || item.weak_verb || item.original_sentence || item.initial || '', 
          improved: item.enhanced_text || item.improved_verb || item.corrected || item.change || item.modified || item.revised || item.new_verb || item.improved || item.replacement || '', 
          reason: item.reason || item.reasoning || item.explanation || '', 
          trope_category: item.trope_category || item.category,
          status: 'pending' as 'pending',
        };
        console.log(`Mapping item ${index} from changesArray:`, item, "to ProseChange:", changeItem);
        return changeItem;
      });

    } catch (error) {
      console.error('Error in parseChanges:', error, "Original response string:", response);
      return [];
    }
  }, []);

  const updateAndSavePrompts = useCallback(async (newPrompts: ProseImprovementPrompt[] | ((prevState: ProseImprovementPrompt[]) => ProseImprovementPrompt[])) => {
    setPrompts(prevPrompts => {
      const updated = typeof newPrompts === 'function' ? newPrompts(prevPrompts) : newPrompts;
      window.go.main.App.WriteProsePromptsFile(JSON.stringify(updated)).catch(err => {
        console.error("Error saving prose prompts:", err);
      });
      return updated;
    });
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
    updatePrompts: updateAndSavePrompts, // Use the new save-aware updater
    parseChanges,
    applyChanges,
    findChangePositions,
    isLoadingPrompts, // Expose loading state if needed by UI
  };
}