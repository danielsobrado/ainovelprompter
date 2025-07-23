// frontend/src/components/ProseImprovement/index.tsx
import React, { useState, useCallback, useEffect } from 'react';
import { Card, CardContent, CardHeader, CardTitle } from '@/components/ui/card';
import { Tabs, TabsList, TabsTrigger, TabsContent } from '@/components/ui/tabs';
import { Button } from '@/components/ui/button';
import { Textarea } from '@/components/ui/textarea';
// import { ScrollArea } from '@/components/ui/scroll-area'; // Not directly used here anymore
// import { Badge } from '@/components/ui/badge'; // Not directly used here anymore
import { Settings } from 'lucide-react';
import type { ProseImprovementPrompt, ProseChange, ProseImprovementSession, LLMProvider } from '@/types';
// import { DEFAULT_PROSE_IMPROVEMENT_PROMPTS } from '@/utils/constants'; // Loaded by hook
import { useLLMProvider } from '@/hooks/useLLMProvider';
import { useProseImprovement } from '@/hooks/useProseImprovement';
import { useWailsReady } from '@/contexts/WailsReadyContext'; // Use the context
import { ProviderSettings } from './ProviderSettings';
import { PromptManager } from './PromptManager';
import { ChangeReviewer } from './ChangeReviewer';
import { ProcessingView } from './ProcessingView';

// Import Wails generated Go function types
import {
    GetInitialLLMSettings,
    ReadLLMSettingsFile,
    WriteLLMSettingsFile
} from '../../../wailsjs/go/main/App'; // Adjust path if your App struct is in a different Go package

export function ProseImprovementTab() {
  const [inputText, setInputText] = useState('');
  const [session, setSession] = useState<ProseImprovementSession | null>(null);
  const [isProcessing, setIsProcessing] = useState(false);
  
  const hardcodedDefaultProvider: LLMProvider = { type: 'manual', config: {} };
  const [selectedProvider, setSelectedProvider] = useState<LLMProvider>(hardcodedDefaultProvider);
  const [isLoadingProviderSettings, setIsLoadingProviderSettings] = useState(true);
  
  // Use the context instead of local state
  const { wailsReady } = useWailsReady();

  const [showProviderSettings, setShowProviderSettings] = useState(false);
  const [manualResponse, setManualResponse] = useState('');
  
  const { executePrompt, isLoading: isLLMLoading, error: llmError, clearError } = useLLMProvider(selectedProvider);
  const { 
    prompts, 
    updatePrompts,
    parseChanges,
    applyChanges,
    isLoadingPrompts 
  } = useProseImprovement();

  // Add a separate timeout for the entire component loading
  useEffect(() => {
    const globalTimeout = setTimeout(() => {
      console.warn("Global component timeout - forcing loading states to false");
      setIsLoadingProviderSettings(false);
    }, 8000); // 8 second global timeout

    return () => clearTimeout(globalTimeout);
  }, []);

  useEffect(() => {
    if (!wailsReady) {
        console.log("Provider settings load: Wails not ready yet.");
        return;
    }

    const loadProviderSettings = async () => {
      console.log("loadProviderSettings called");
      setIsLoadingProviderSettings(true);
      try {
        const persistedSettingsJson = await ReadLLMSettingsFile();
        const persistedSettings = JSON.parse(persistedSettingsJson || "{}") as Partial<LLMProvider>;

        if (persistedSettings && persistedSettings.type && Object.keys(persistedSettings).length > 0) {
          console.log("Loaded provider settings from llm_provider_settings.json:", persistedSettings);
          setSelectedProvider(persistedSettings as LLMProvider);
        } else {
          console.log("No valid persisted LLM settings found, fetching from Go backend (.env/config.yaml)...");
          const initialEnvSettings = await GetInitialLLMSettings();
          console.log("Settings from Go backend (GetInitialLLMSettings):", initialEnvSettings);

          let providerFromEnv: LLMProvider = { ...hardcodedDefaultProvider }; 

          if (initialEnvSettings && Object.keys(initialEnvSettings).length > 0) {
            // Prioritize OpenRouter if key is present
            if (initialEnvSettings.openrouter_api_key) {
              providerFromEnv = {
                type: 'openrouter',
                config: {
                  apiKey: initialEnvSettings.openrouter_api_key,
                  model: initialEnvSettings.openrouter_default_model || 'anthropic/claude-3-haiku',
                },
              };
            } else if (initialEnvSettings.lmstudio_api_url) { // Fallback to LMStudio if its URL is set
              providerFromEnv = {
                type: 'lmstudio',
                config: {
                  apiUrl: initialEnvSettings.lmstudio_api_url,
                  model: initialEnvSettings.lmstudio_default_model || 'local-model',
                },
              };
            }
            // If neither, it remains 'manual' from hardcodedDefaultProvider
            console.log("Using provider settings from Go backend (.env/config.yaml):", providerFromEnv);
            setSelectedProvider(providerFromEnv);
            await WriteLLMSettingsFile(JSON.stringify(providerFromEnv));
            console.log("Saved initial .env/config.yaml settings to llm_provider_settings.json");
          } else {
            console.log("No settings from Go backend either, using hardcoded default provider.");
            setSelectedProvider(hardcodedDefaultProvider);
            // Optionally save hardcoded default if you want the file to always exist after first run
            // await WriteLLMSettingsFile(JSON.stringify(hardcodedDefaultProvider));
          }
        }
      } catch (error) {
        console.error("Error loading provider settings:", error);
        setSelectedProvider(hardcodedDefaultProvider);
      } finally {
        setIsLoadingProviderSettings(false);
      }
    };

    loadProviderSettings();
    
    // Timeout fallback - if loading takes too long, proceed with default
    const timeoutId = setTimeout(() => {
      console.warn("Provider settings loading timeout - proceeding with default");
      setSelectedProvider(hardcodedDefaultProvider);
      setIsLoadingProviderSettings(false);
    }, 5000); // 5 second timeout

    // Clean up timeout if component unmounts
    return () => clearTimeout(timeoutId);
  }, [wailsReady]); // Re-run when wailsReady changes

  const handleProviderChange = useCallback(async (newProvider: LLMProvider) => {
    if (!wailsReady) {
        console.warn("Cannot save provider settings: Wails not ready.");
        setSelectedProvider(newProvider); // Update state locally anyway
        return;
    }
    setSelectedProvider(newProvider);
    try {
      console.log("Saving provider settings to llm_provider_settings.json:", newProvider);
      await WriteLLMSettingsFile(JSON.stringify(newProvider));
    } catch (error) {
      console.error("Error saving provider settings:", error);
    }
  }, [wailsReady]); // Add wailsReady as a dependency


  const startSession = useCallback(() => {
    if (!inputText.trim()) return;
    const newSession: ProseImprovementSession = {
      id: crypto.randomUUID(),
      originalText: inputText,
      currentText: inputText,
      prompts: prompts, // Assumes prompts are loaded
      currentPromptIndex: 0,
      changes: [],
      createdAt: new Date(),
      updatedAt: new Date()
    };
    setSession(newSession);
  }, [inputText, prompts]);
  const processNextPrompt = useCallback(async () => {
    if (!session || session.currentPromptIndex >= session.prompts.length || isProcessing) return;
    
    setIsProcessing(true);
    const currentPrompt = session.prompts[session.currentPromptIndex];
    
    try {
      const fullPrompt = await buildFullPrompt(currentPrompt, session.currentText, selectedProvider);
      console.log("Full Prompt being sent to LLM (or copied for manual):", fullPrompt);
      console.log("Selected Provider for this prompt:", selectedProvider);
      console.log("Current Prompt Object:", currentPrompt);
      
      if (selectedProvider.type === 'manual') {
        await navigator.clipboard.writeText(fullPrompt);
        // User will paste response manually - no further automatic processing here
        // The UI for manual input will trigger processManualResponse
      } else {
        const response = await executePrompt(fullPrompt);
        console.log("Raw LLM Response String:", response);
        const changes = parseChanges(response);
        
        setSession(prev => {
          if (!prev) return null;
          // Filter out empty/invalid changes that might result from parsing errors
          const validChanges = changes.filter(c => c.initial || c.improved); 
          return {
            ...prev,
            changes: [...prev.changes, ...validChanges],
            currentPromptIndex: prev.currentPromptIndex + 1,
            updatedAt: new Date()
          };
        });
      }
    } catch (error) {
      console.error('Error processing prompt:', error);
      // Optionally, allow user to advance or retry
    } finally {
      setIsProcessing(false);
    }
  }, [session, selectedProvider, executePrompt, parseChanges, isProcessing, prompts]); // Added prompts and isProcessing

  const processManualResponse = useCallback(() => {
    if (!session || !manualResponse.trim()) return;
    console.log("Raw Manual Response String:", manualResponse);
    const changes = parseChanges(manualResponse);
    setSession(prev => {
      if (!prev) return null;
      const validChanges = changes.filter(c => c.initial || c.improved);
      return {
        ...prev,
        changes: [...prev.changes, ...validChanges],
        currentPromptIndex: prev.currentPromptIndex + 1,
        updatedAt: new Date()
      };
    });
    setManualResponse('');
  }, [session, manualResponse, parseChanges]);

  const handleChangeDecision = useCallback((changeId: string, decision: 'accepted' | 'rejected') => {
    setSession(prev => {
      if (!prev) return prev;
      const updatedChanges = prev.changes.map(change =>
        change.id === changeId ? { ...change, status: decision } : change
      );
      const acceptedChanges = updatedChanges.filter(c => c.status === 'accepted');
      const newText = applyChanges(prev.originalText, acceptedChanges);
      return {
        ...prev,
        changes: updatedChanges,
        currentText: newText,
        updatedAt: new Date()
      };
    });
  }, [applyChanges]);

  if (isLoadingPrompts || isLoadingProviderSettings) {
    return (
      <div className="p-4 text-center space-y-2">
        <div>Loading settings and prompts...</div>
        <div className="text-sm text-muted-foreground">
          Wails Ready: {wailsReady ? 'Yes' : 'No'} | 
          Loading Prompts: {isLoadingPrompts ? 'Yes' : 'No'} | 
          Loading Provider: {isLoadingProviderSettings ? 'Yes' : 'No'}
        </div>
      </div>
    );
  }

  return (
    <div className="space-y-4">
      <Card>
        <CardHeader className="flex flex-row items-center justify-between">
          <CardTitle>Prose Improvement</CardTitle>
          <Button
            variant="outline"
            size="sm"
            onClick={() => setShowProviderSettings(true)}
          >
            <Settings className="mr-2 h-4 w-4" />
            Provider Settings ({selectedProvider.type})
          </Button>
        </CardHeader>
        <CardContent>
          <Tabs defaultValue="input" className="w-full">
            <TabsList className="grid w-full grid-cols-4">
              <TabsTrigger value="input">Input Text</TabsTrigger>
              <TabsTrigger value="prompts">Prompts ({prompts.length})</TabsTrigger>
              <TabsTrigger value="process" disabled={!session}>Process</TabsTrigger>
              <TabsTrigger value="review" disabled={!session || session.changes.length === 0}>Review Changes</TabsTrigger>
            </TabsList>
            
            <TabsContent value="input" className="space-y-4 pt-4">
              <div>
                <Textarea
                  placeholder="Paste your text here for improvement..."
                  value={inputText}
                  onChange={(e) => setInputText(e.target.value)}
                  className="min-h-[400px]"
                />
                <div className="mt-2 flex justify-between items-center">
                  <span className="text-sm text-muted-foreground">
                    {inputText.length} characters
                  </span>
                  <Button 
                    onClick={startSession}
                    disabled={!inputText.trim() || prompts.length === 0}
                  >
                    Start Improvement Session
                  </Button>
                </div>
              </div>
            </TabsContent>
            
            <TabsContent value="prompts" className="pt-4">
              <PromptManager
                prompts={prompts}
                onPromptsChange={updatePrompts}
              />
            </TabsContent>
            
            <TabsContent value="process" className="pt-4">
              {session ? (
                <ProcessingView
                  session={session}
                  isProcessing={isProcessing || isLLMLoading} // Combine processing states
                  llmError={llmError} // Pass LLM error
                  onProcessNext={processNextPrompt}
                  onClearError={clearError} // Add clear error function
                  selectedProvider={selectedProvider}
                  manualResponse={manualResponse}
                  onManualResponseChange={setManualResponse}
                  onProcessManualResponse={processManualResponse}
                />
              ) : (
                <div className="text-center py-12 text-muted-foreground">
                  No active session. Start by entering text in the Input tab.
                </div>
              )}
            </TabsContent>
            
            <TabsContent value="review" className="pt-4">
              {session && session.changes.length > 0 ? (
                <ChangeReviewer
                  changes={session.changes}
                  originalText={session.originalText}
                  currentText={session.currentText}
                  onChangeDecision={handleChangeDecision}
                />
              ) : (
                <div className="text-center py-12 text-muted-foreground">
                  No changes to review yet. Process some prompts first or ensure prompts generated valid changes.
                </div>
              )}
            </TabsContent>
          </Tabs>
        </CardContent>
      </Card>
      
      <ProviderSettings
        isOpen={showProviderSettings}
        onClose={() => setShowProviderSettings(false)}
        provider={selectedProvider}
        onProviderChange={handleProviderChange}
      />
    </div>
  );
}

async function buildFullPrompt(
  prompt: ProseImprovementPrompt, 
  text: string, 
  selectedProvider: LLMProvider
): Promise<string> {
  try {
    // Use the new backend function to get the resolved prompt
    const providerJSON = JSON.stringify(selectedProvider);
    const resolvedPromptText = await window.go.main.App.GetResolvedProsePrompt(prompt.id, providerJSON);
    
    // Replace the placeholder with the actual text
    const placeholder = "[TEXT_TO_ANALYZE_PLACEHOLDER]";
    if (resolvedPromptText.includes(placeholder)) {
      return resolvedPromptText.replace(placeholder, text);
    }
    
    // Fallback if no placeholder found - append text
    return `${resolvedPromptText}\n\nText to analyze:\n${text}`;
  } catch (error) {
    console.error('Error resolving prompt from backend:', error);
    
    // Fallback to defaultPromptText if backend fails
    const fallbackPrompt = prompt.defaultPromptText || "Please analyze and improve the following text:";
    const placeholder = "[TEXT_TO_ANALYZE_PLACEHOLDER]";
    if (fallbackPrompt.includes(placeholder)) {
      return fallbackPrompt.replace(placeholder, text);
    }
    return `${fallbackPrompt}\n\nText to analyze:\n${text}`;
  }
}