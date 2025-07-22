import React, { useState, useEffect } from 'react';
import { Badge } from './ui/badge';
import { Database, HardDrive } from 'lucide-react';
import { GetStorageStats } from '../../wailsjs/go/main/App';

interface StorageStats {
  totalFiles: number;
  totalSize: number;
  entitiesByType: Record<string, number>;
  versionsByType: Record<string, number>;
  oldestTimestamp: string;
  newestTimestamp: string;
}

const StorageIndicator: React.FC = () => {
  const [stats, setStats] = useState<StorageStats | null>(null);
  const [isLoading, setIsLoading] = useState(false);
  const [hasError, setHasError] = useState(false);

  useEffect(() => {
    loadStorageStats();
    // Refresh stats every 30 seconds
    const interval = setInterval(loadStorageStats, 30000);
    return () => clearInterval(interval);
  }, []);

  const loadStorageStats = async () => {
    try {
      setIsLoading(true);
      setHasError(false);
      // Call the backend method through Wails
      const statsData = await GetStorageStats();
      setStats(statsData);
    } catch (err) {
      // Track error state but don't fail completely
      console.warn('Failed to load storage stats:', err);
      setHasError(true);
      // Set fallback stats to show something useful
      setStats({
        totalFiles: 0,
        totalSize: 0,
        entitiesByType: {},
        versionsByType: {},
        oldestTimestamp: '',
        newestTimestamp: ''
      });
    } finally {
      setIsLoading(false);
    }
  };

  const formatFileSize = (bytes: number): string => {
    if (bytes === 0) return '0 B';
    
    const units = ['B', 'KB', 'MB', 'GB'];
    let size = bytes;
    let unitIndex = 0;
    
    while (size >= 1024 && unitIndex < units.length - 1) {
      size /= 1024;
      unitIndex++;
    }
    
    return `${size.toFixed(size < 10 ? 1 : 0)} ${units[unitIndex]}`;
  };

  if (isLoading) {
    return (
      <div className="flex items-center gap-2 text-xs text-gray-400">
        <Database className="h-3 w-3 animate-spin" />
        <span>Loading...</span>
      </div>
    );
  }

  // Always show something, even with fallback data
  const displayStats = stats || {
    totalFiles: 0,
    totalSize: 0,
    entitiesByType: {},
    versionsByType: {},
    oldestTimestamp: '',
    newestTimestamp: ''
  };

  return (
    <div className="flex items-center gap-2 text-xs text-gray-500">
      <div className="flex items-center gap-1">
        <Database className={`h-3 w-3 ${hasError ? 'text-orange-500' : ''}`} />
        <span>{displayStats.totalFiles} files{hasError ? ' (!)' : ''}</span>
      </div>
      <div className="flex items-center gap-1">
        <HardDrive className={`h-3 w-3 ${hasError ? 'text-orange-500' : ''}`} />
        <span>{formatFileSize(displayStats.totalSize)}</span>
      </div>
    </div>
  );
};

export default StorageIndicator;
