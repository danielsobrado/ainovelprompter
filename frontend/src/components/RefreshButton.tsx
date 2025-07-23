import React from 'react';
import { Button } from './ui/button';
import { RefreshCw } from 'lucide-react';
import { useDataRefresh } from '../contexts/DataRefreshContext';

interface RefreshButtonProps {
  type?: 'taskTypes' | 'rules' | 'characters' | 'locations' | 'codex' | 'sampleChapters' | 'all';
  size?: 'sm' | 'default' | 'lg' | 'icon';
  variant?: 'default' | 'destructive' | 'outline' | 'secondary' | 'ghost' | 'link';
  className?: string;
  onRefresh?: () => void;
  disabled?: boolean;
}

export function RefreshButton({ 
  type = 'all', 
  size = 'icon', 
  variant = 'ghost', 
  className = '',
  onRefresh,
  disabled = false
}: RefreshButtonProps) {
  const { refreshFunctions, isRefreshing } = useDataRefresh();

  const handleRefresh = async () => {
    if (!refreshFunctions || disabled) return;

    try {
      switch (type) {
        case 'taskTypes':
          await refreshFunctions.refreshTaskTypes();
          break;
        case 'rules':
          await refreshFunctions.refreshRules();
          break;
        case 'characters':
          await refreshFunctions.refreshCharacters();
          break;
        case 'locations':
          await refreshFunctions.refreshLocations();
          break;
        case 'codex':
          await refreshFunctions.refreshCodex();
          break;
        case 'sampleChapters':
          await refreshFunctions.refreshSampleChapters();
          break;
        case 'all':
        default:
          await refreshFunctions.refreshAll();
          break;
      }
      onRefresh?.();
    } catch (error) {
      console.error('Failed to refresh data:', error);
    }
  };

  const getTooltipText = () => {
    switch (type) {
      case 'taskTypes':
        return 'Refresh task types';
      case 'rules':
        return 'Refresh rules';
      case 'characters':
        return 'Refresh characters';
      case 'locations':
        return 'Refresh locations';
      case 'codex':
        return 'Refresh codex entries';
      case 'sampleChapters':
        return 'Refresh sample chapters';
      case 'all':
      default:
        return 'Refresh all data';
    }
  };

  const isButtonLoading = isRefreshing;

  return (
    <Button
      variant={variant}
      size={size}
      onClick={handleRefresh}
      disabled={disabled || isButtonLoading || !refreshFunctions}
      className={className}
      title={getTooltipText()}
    >
      <RefreshCw 
        className={`h-4 w-4 ${isButtonLoading ? 'animate-spin' : ''}`} 
      />
      {size !== 'icon' && (
        <span className="ml-2">
          {isButtonLoading ? 'Refreshing...' : 'Refresh'}
        </span>
      )}
    </Button>
  );
}

export default RefreshButton;
