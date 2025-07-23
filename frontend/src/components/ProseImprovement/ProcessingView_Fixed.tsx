// frontend/src/components/ProseImprovement/ProcessingView.tsx
import React from 'react';
import { Card } from '@/components/ui/card';
import { Button } from '@/components/ui/button';
import { Textarea } from '@/components/ui/textarea';
import { Copy, Play, RefreshCw, AlertCircle } from 'lucide-react';
import { Alert, AlertDescription, AlertTitle } from '@/components/ui/alert';
import type { ProseImprovementSession, LLMProvider } from '@/types';

interface ProcessingViewProps {
  session: ProseImprovementSession;
  isProcessing: boolean;
  llmError: string | null;
  onProcessNext: () => void;
  onClearError: () => void;
  selectedProvider: LLMProvider;
  manualResponse: string;
  onManualResponseChange: (value: string) => void;
  onProcessManualResponse: () => void;
}

export function ProcessingView({
  session,
  isProcessing,
  llmError,
  onProcessNext,
  onClearError,
  selectedProvider,
  manualResponse,
  onManualResponseChange,
  onProcessManualResponse
}: ProcessingViewProps) {
  const currentPrompt = session.prompts[session.currentPromptIndex];
  const progress = (session.currentPromptIndex / session.prompts.length) * 100;

  const copyFullPrompt = async () => {
    if (!currentPrompt) return;
    // Use defaultPromptText as fallback for manual copying
    const promptText = currentPrompt.defaultPromptText || "Analyze the following text:";
    const placeholder = "[TEXT_TO_ANALYZE_PLACEHOLDER]";
    let fullPrompt: string;
    
    if (promptText.includes(placeholder)) {
      fullPrompt = promptText.replace(placeholder, session.currentText);
    } else {
      fullPrompt = `${promptText}\n\nText to analyze:\n${session.currentText}`;
    }
    
    await navigator.clipboard.writeText(fullPrompt);
  };

  const handleRetry = () => {
    onClearError();
    onProcessNext();
  };

  return (
    <div className="space-y-4">
      {/* Progress bar */}
      <div className="space-y-2">
        <div className="flex justify-between text-sm">
          <span>Progress: {session.currentPromptIndex} / {session.prompts.length}</span>
          <span>{Math.round(progress)}%</span>
        </div>
        <div className="h-2 bg-muted rounded-full overflow-hidden">
          <div 
            className="h-full bg-primary transition-all duration-300"
            style={{ width: `${progress}%` }}
          />
        </div>
      </div>

      {/* Enhanced LLM Error Message with Recovery Options */}
      {llmError && (
        <Alert variant="destructive">
          <AlertCircle className="h-4 w-4" />
          <AlertTitle>LLM Provider Error</AlertTitle>
          <AlertDescription className="space-y-3">
            <div className="font-mono text-sm bg-destructive/10 p-2 rounded">
              {llmError}
            </div>
            <div className="text-sm text-muted-foreground">
              This error occurred while trying to process your text with the {selectedProvider.type} provider.
            </div>
            <div className="flex flex-wrap gap-2">
              <Button 
                variant="outline" 
                size="sm" 
                onClick={handleRetry}
                disabled={isProcessing}
              >
                <RefreshCw className="mr-2 h-4 w-4" />
                Retry
              </Button>
              <Button 
                variant="ghost" 
                size="sm" 
                onClick={onClearError}
              >
                Clear Error
              </Button>
              <Button 
                variant="ghost" 
                size="sm" 
                onClick={() => window.location.reload()}
              >
                Reset App
              </Button>
            </div>
            <div className="text-xs text-muted-foreground">
              💡 <strong>Troubleshooting:</strong> Check your API keys in Provider Settings, verify network connection, or try switching to manual mode.
            </div>
          </AlertDescription>
        </Alert>
      )}

      {/* Current prompt */}
      {currentPrompt ? (
        <Card className="p-4">
          <h3 className="font-semibold mb-2">Current Prompt: {currentPrompt.label}</h3>
          <p className="text-sm text-muted-foreground mb-4">
            {currentPrompt.description || "Processing text with this prompt..."}
          </p>
          
          {selectedProvider.type === 'manual' ? (
            <div className="space-y-3">
              <Button onClick={copyFullPrompt} variant="outline">
                <Copy className="mr-2 h-4 w-4" />
                Copy Full Prompt
              </Button>
              <Textarea
                placeholder="Paste the AI response here..."
                value={manualResponse}
                onChange={(e) => onManualResponseChange(e.target.value)}
                className="min-h-[200px]"
              />
              <Button 
                onClick={onProcessManualResponse}
                disabled={!manualResponse.trim()}
              >
                Process Response
              </Button>
            </div>
          ) : (
            <Button 
              onClick={onProcessNext}
              disabled={isProcessing || !!llmError}
            >
              {isProcessing ? (
                <>
                  <RefreshCw className="mr-2 h-4 w-4 animate-spin" />
                  Processing...
                </>
              ) : (
                <>
                  <Play className="mr-2 h-4 w-4" />
                  Execute Prompt
                </>
              )}
            </Button>
          )}
        </Card>
      ) : (
        <Card className="p-4">
          <div className="text-center text-muted-foreground">
            All prompts have been processed. Review your changes in the Review tab.
          </div>
        </Card>
      )}

      {/* Stats */}
      <div className="grid grid-cols-3 gap-4">
        <Card className="p-4">
          <div className="text-2xl font-bold">{session.changes.length}</div>
          <div className="text-sm text-muted-foreground">Total Changes</div>
        </Card>
        <Card className="p-4">
          <div className="text-2xl font-bold">
            {session.changes.filter(c => c.status === 'accepted').length}
          </div>
          <div className="text-sm text-muted-foreground">Accepted</div>
        </Card>
        <Card className="p-4">
          <div className="text-2xl font-bold">
            {session.changes.filter(c => c.status === 'pending').length}
          </div>
          <div className="text-sm text-muted-foreground">Pending Review</div>
        </Card>
      </div>
    </div>
  );
}
