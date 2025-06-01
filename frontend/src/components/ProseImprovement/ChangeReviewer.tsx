// frontend/src/components/ProseImprovement/ChangeReviewer.tsx
import React, { useState, useEffect } from 'react';
import { Card } from '@/components/ui/card';
import { Button } from '@/components/ui/button';
import { Badge } from '@/components/ui/badge';
import { ScrollArea } from '@/components/ui/scroll-area';
import { Tabs, TabsList, TabsTrigger } from '@/components/ui/tabs';
import { Check, X, Eye, EyeOff, Copy } from 'lucide-react';
import type { ProseChange } from '@/types';
import { cn } from '@/lib/utils';

interface ChangeReviewerProps {
  changes: ProseChange[];
  originalText: string;
  currentText: string;
  onChangeDecision: (changeId: string, decision: 'accepted' | 'rejected') => void;
}

export function ChangeReviewer({
  changes,
  originalText,
  currentText,
  onChangeDecision
}: ChangeReviewerProps) {
  const [showContext, setShowContext] = useState(true);
  const [filterStatus, setFilterStatus] = useState<'all' | 'pending' | 'accepted' | 'rejected'>('all');

  const filteredChanges = changes.filter(change => {
    if (filterStatus === 'all') return true;
    return change.status === filterStatus;
  });

  // Log the first few changes to check their content when the component or filter updates
  useEffect(() => {
    if (filteredChanges.length > 0) {
      console.log("ChangeReviewer - Filtered Changes (first 3):", filteredChanges.slice(0, 3));
    }
  }, [filteredChanges]);

  const stats = {
    total: changes.length,
    pending: changes.filter(c => c.status === 'pending').length,
    accepted: changes.filter(c => c.status === 'accepted').length,
    rejected: changes.filter(c => c.status === 'rejected').length
  };

  const copyCurrentText = async () => {
    await navigator.clipboard.writeText(currentText);
  };

  return (
    <div className="p-4 space-y-4">
      {/* Filter and Stats */}
      <div className="flex gap-4">
        <Badge variant="outline">Total: {stats.total}</Badge>
        <Badge variant="outline" className="text-yellow-600">Pending: {stats.pending}</Badge>
        <Badge variant="outline" className="text-green-600">Accepted: {stats.accepted}</Badge>
        <Badge variant="outline" className="text-red-600">Rejected: {stats.rejected}</Badge>
      </div>

      {/* Controls */}
      <div className="flex justify-between items-center">
        <Tabs value={filterStatus} onValueChange={(v: any) => setFilterStatus(v)}>
          <TabsList>
            <TabsTrigger value="all">All</TabsTrigger>
            <TabsTrigger value="pending">Pending</TabsTrigger>
            <TabsTrigger value="accepted">Accepted</TabsTrigger>
            <TabsTrigger value="rejected">Rejected</TabsTrigger>
          </TabsList>
        </Tabs>
        
        <Button
          variant="outline"
          size="sm"
          onClick={() => setShowContext(!showContext)}
        >
          {showContext ? <EyeOff className="mr-2 h-4 w-4" /> : <Eye className="mr-2 h-4 w-4" />}
          {showContext ? 'Hide' : 'Show'} Context
        </Button>
      </div>

      {/* Changes list */}
      <ScrollArea className="h-[500px]">
        <div className="space-y-3">
          {filteredChanges.map((change) => (
            // console.log("Rendering change object:", change), // Uncomment to log every change object
            <Card key={change.id} className="p-4">
              <div className="space-y-3">
                {/* Status badge */}
                <div className="flex items-center justify-between">
                  <Badge
                    variant={
                      change.status === 'accepted' ? 'default' : // 'default' for accepted often is primary color
                      change.status === 'rejected' ? 'destructive' : // 'destructive' for rejected
                      'outline'
                    }
                    className={cn(
                        change.status === 'accepted' && "bg-green-600 text-white",
                        change.status === 'rejected' && "bg-red-600 text-white",
                        change.status === 'pending' && "border-yellow-500 text-yellow-600"
                    )}
                  >
                    {change.status}
                  </Badge>
                  {change.trope_category && (
                    <Badge variant="outline">{change.trope_category}</Badge>
                  )}
                </div>

                {/* Change content */}
                <div className="space-y-2">
                  <div className="p-3 bg-red-50 dark:bg-red-900/20 rounded-md">
                    <p className="text-sm font-mono text-red-700 dark:text-red-300">
                      <span className="font-semibold">Original:</span> {change.initial || '(No original text provided)'}
                    </p>
                  </div>
                  <div className="p-3 bg-green-50 dark:bg-green-900/20 rounded-md">
                    <p className="text-sm font-mono text-green-700 dark:text-green-300">
                      <span className="font-semibold">Improved:</span> {change.improved || '(No improved text provided)'}
                    </p>
                  </div>
                </div>

                {/* Reason */}
                <div className="text-sm text-muted-foreground">
                  <strong>Reason:</strong> {change.reason || '(No reason provided)'}
                </div>

                {/* Context (if enabled) */}
                {showContext && change.startIndex !== undefined && (
                  <div className="p-3 bg-muted rounded-md">
                    <p className="text-xs text-muted-foreground mb-1">Context:</p>
                    <p className="text-sm">
                      ...{originalText.slice(
                        Math.max(0, change.startIndex - 50),
                        change.startIndex
                      )}
                      <span className="font-bold text-red-600">{change.initial}</span>
                      {originalText.slice(
                        change.endIndex || change.startIndex + change.initial.length,
                        Math.min(
                          originalText.length,
                          (change.endIndex || change.startIndex + change.initial.length) + 50
                        )
                      )}...
                    </p>
                  </div>
                )}

                {/* Actions */}
                {change.status === 'pending' && (
                  <div className="flex gap-2">
                    <Button
                      size="sm"
                      variant="default"
                      onClick={() => onChangeDecision(change.id, 'accepted')}
                      className="flex-1"
                    >
                      <Check className="mr-2 h-4 w-4" />
                      Accept
                    </Button>
                    <Button
                      size="sm"
                      variant="destructive"
                      onClick={() => onChangeDecision(change.id, 'rejected')}
                      className="flex-1"
                    >
                      <X className="mr-2 h-4 w-4" />
                      Reject
                    </Button>
                  </div>
                )}
              </div>
            </Card>
          ))}
        </div>
      </ScrollArea>

      {/* Final text preview */}
      <Card className="p-4">
        <div className="flex items-center justify-between mb-2">
          <h3 className="font-semibold">Current Text Preview</h3>
          <Button
            size="sm"
            variant="outline"
            onClick={copyCurrentText}
          >
            <Copy className="mr-2 h-4 w-4" />
            Copy Text
          </Button>
        </div>
        <ScrollArea className="h-[200px]">
          <pre className="whitespace-pre-wrap font-mono text-sm">{currentText}</pre>
        </ScrollArea>
      </Card>
    </div>
  );
}