// frontend/src/components/ProseImprovement/index.tsx
import React, { useState, useCallback } from 'react';
import { Card, CardContent, CardHeader, CardTitle } from '@/components/ui/card';
import { Tabs, TabsList, TabsTrigger, TabsContent } from '@/components/ui/tabs';
import { Button } from '@/components/ui/button';
import { Textarea } from '@/components/ui/textarea';
import { ScrollArea } from '@/components/ui/scroll-area';
import { Badge } from '@/components/ui/badge';
import { 
  Play, 
  Pause, 
  SkipForward, 
  Copy, 
  Check, 
  X, 
  Settings,
  ChevronRight,
  RefreshCw
} from 'lucide-react';
import type { ProseImprovementPrompt, ProseChange, ProseImprovementSession, LLMProvider } from '@/types';
import { DEFAULT_PROSE_IMPROVEMENT_PROMPTS } from '@/utils/constants';
import { useLLMProvider } from '@/hooks/useLLMProvider';
import { useProseImprovement } from '@/hooks/useProseImprovement';
import { ProviderSettings } from './ProviderSettings';
import { PromptManager } from './PromptManager';
import { ChangeReviewer } from './ChangeReviewer';
import { ProcessingView } from './ProcessingView';

export function ProseImprovementTab() {
  const [inputText, setInputText] = useState('');
  const [session, setSession] = useState<ProseImprovementSession | null>(null);
  const [isProcessing, setIsProcessing] = useState(false);
  const [selectedProvider, setSelectedProvider] = useState<LLMProvider>({ type: 'manual' });
  const [showProviderSettings, setShowProviderSettings] = useState(false);
  const [manualResponse, setManualResponse] = useState('');
  
  const { executePrompt } = useLLMProvider(selectedProvider);
  const { 
    prompts, 
    updatePrompts,
    parseChanges,
    applyChanges 
  } = useProseImprovement();

  const startSession = useCallback(() => {
    if (!inputText.trim()) return;
    
    const newSession: ProseImprovementSession = {
      id: crypto.randomUUID(),
      originalText: inputText,
      currentText: inputText,
      prompts: prompts,
      currentPromptIndex: 0,
      changes: [],
      createdAt: new Date(),
      updatedAt: new Date()
    };
    
    setSession(newSession);
  }, [inputText, prompts]);

  const processNextPrompt = useCallback(async () => {
    if (!session || session.currentPromptIndex >= session.prompts.length) return;
    
    setIsProcessing(true);
    const currentPrompt = session.prompts[session.currentPromptIndex];
    
    try {
      // Build the full prompt with context
      const fullPrompt = buildFullPrompt(currentPrompt, session.currentText);
      
      if (selectedProvider.type === 'manual') {
        // Copy to clipboard
        await navigator.clipboard.writeText(fullPrompt);
        // User will paste response manually
      } else {
        // Execute via API
        const response = await executePrompt(fullPrompt);
        console.log("Raw LLM Response String:", response); // <-- ADD THIS LOG
        const changes = parseChanges(response);
        
        setSession(prev => ({
          ...prev!,
          changes: [...prev!.changes, ...changes],
          currentPromptIndex: prev!.currentPromptIndex + 1,
          updatedAt: new Date()
        }));
      }
    } catch (error) {
      console.error('Error processing prompt:', error);
    } finally {
      setIsProcessing(false);
    }
  }, [session, selectedProvider, executePrompt, parseChanges]);

  const processManualResponse = useCallback(() => {
    if (!session || !manualResponse.trim()) return;
    console.log("Raw Manual Response String:", manualResponse); // <-- ADD THIS LOG
    const changes = parseChanges(manualResponse);
    
    setSession(prev => ({
      ...prev!,
      changes: [...prev!.changes, ...changes],
      currentPromptIndex: prev!.currentPromptIndex + 1,
      updatedAt: new Date()
    }));
    
    setManualResponse('');
  }, [session, manualResponse, parseChanges]);

  const handleChangeDecision = useCallback((changeId: string, decision: 'accepted' | 'rejected') => {
    setSession(prev => {
      if (!prev) return prev;
      
      const updatedChanges = prev.changes.map(change =>
        change.id === changeId ? { ...change, status: decision } : change
      );
      
      // Apply accepted changes to current text
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
            Provider Settings
          </Button>
        </CardHeader>
        <CardContent>
          <Tabs defaultValue="input" className="w-full">
            <TabsList className="grid w-full grid-cols-4">
              <TabsTrigger value="input">Input Text</TabsTrigger>
              <TabsTrigger value="prompts">Prompts</TabsTrigger>
              <TabsTrigger value="process">Process</TabsTrigger>
              <TabsTrigger value="review">Review Changes</TabsTrigger>
            </TabsList>
            
            <TabsContent value="input" className="space-y-4">
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
                    disabled={!inputText.trim()}
                  >
                    Start Improvement Session
                  </Button>
                </div>
              </div>
            </TabsContent>
            
            <TabsContent value="prompts">
              <PromptManager
                prompts={prompts}
                onPromptsChange={updatePrompts}
              />
            </TabsContent>
            
            <TabsContent value="process">
              {session ? (
                <ProcessingView
                  session={session}
                  isProcessing={isProcessing}
                  onProcessNext={processNextPrompt}
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
            
            <TabsContent value="review">
              {session && session.changes.length > 0 ? (
                <ChangeReviewer
                  changes={session.changes}
                  originalText={session.originalText}
                  currentText={session.currentText}
                  onChangeDecision={handleChangeDecision}
                />
              ) : (
                <div className="text-center py-12 text-muted-foreground">
                  No changes to review yet. Process some prompts first.
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
        onProviderChange={setSelectedProvider}
      />
    </div>
  );
}

function buildFullPrompt(prompt: ProseImprovementPrompt, text: string): string {
  return `${prompt.prompt}\n\nText to analyze:\n${text}`;
}