import React, { useState, useEffect } from 'react';
import { Button } from './ui/button';
import { Input } from './ui/input';
import { Label } from './ui/label';
import { Alert, AlertDescription } from './ui/alert';
import { Dialog, DialogContent, DialogHeader, DialogTitle, DialogTrigger } from './ui/dialog';
import { Select, SelectContent, SelectItem, SelectTrigger, SelectValue } from './ui/select';
import { Card, CardContent, CardDescription, CardHeader, CardTitle } from './ui/card';
import { Badge } from './ui/badge';
import { FolderOpen, AlertTriangle, CheckCircle, Clock, Database } from 'lucide-react';
import { 
  GetDataDirectory, 
  GetRecentDataDirectories, 
  GetStorageStats, 
  ValidateDataDirectoryPath, 
  SetDataDirectory, 
  AddRecentDataDirectory, 
  SelectDirectory 
} from '../../wailsjs/go/main/App';

interface DataDirectoryManagerProps {
  onDirectoryChange?: (newPath: string) => void;
}

interface StorageStats {
  totalFiles: number;
  totalSize: number;
  entitiesByType: Record<string, number>;
  versionsByType: Record<string, number>;
  oldestTimestamp: string;
  newestTimestamp: string;
}

interface MigrationProgress {
  isRunning: boolean;
  step: string;
  progress: number;
  error?: string;
}

const DataDirectoryManager: React.FC<DataDirectoryManagerProps> = ({ onDirectoryChange }) => {
  const [currentDirectory, setCurrentDirectory] = useState<string>('');
  const [newDirectory, setNewDirectory] = useState<string>('');
  const [recentDirectories, setRecentDirectories] = useState<string[]>([]);
  const [stats, setStats] = useState<StorageStats | null>(null);
  const [migrationProgress, setMigrationProgress] = useState<MigrationProgress>({ isRunning: false, step: '', progress: 0 });
  const [isLoading, setIsLoading] = useState(false);
  const [error, setError] = useState<string>('');
  const [success, setSuccess] = useState<string>('');

  // Load current directory and stats on mount
  useEffect(() => {
    loadCurrentDirectory();
    loadRecentDirectories();
    loadStorageStats();
  }, []);

  const loadCurrentDirectory = async () => {
    try {
      const directory = await GetDataDirectory();
      setCurrentDirectory(directory);
      setNewDirectory(directory);
    } catch (err) {
      setError('Failed to load current directory');
    }
  };

  const loadRecentDirectories = async () => {
    try {
      const recent = await GetRecentDataDirectories();
      setRecentDirectories(recent);
    } catch (err) {
      // Recent directories are optional
    }
  };

  const loadStorageStats = async () => {
    try {
      // Call the backend method through Wails
      const statsData = await GetStorageStats();
      setStats(statsData);
    } catch (err) {
      // Stats are optional
    }
  };

  const handleDirectorySelect = async () => {
    try {
      setIsLoading(true);
      setError('');
      
      // Validate directory first
      await ValidateDataDirectoryPath(newDirectory);
      
      // Change directory
      await SetDataDirectory(newDirectory);
      
      // Update current directory
      setCurrentDirectory(newDirectory);
      
      // Add to recent directories
      await AddRecentDataDirectory(newDirectory);
      
      // Reload recent directories and stats
      await loadRecentDirectories();
      await loadStorageStats();
      
      setSuccess(`Data directory changed to: ${newDirectory}`);
      
      // Notify parent component
      if (onDirectoryChange) {
        onDirectoryChange(newDirectory);
      }
      
    } catch (err) {
      setError(err instanceof Error ? err.message : 'Failed to change directory');
    } finally {
      setIsLoading(false);
    }
  };

  const handleBrowseDirectory = async () => {
    try {
      // This would use Wails file dialog
      const selected = await SelectDirectory();
      if (selected) {
        setNewDirectory(selected);
      }
    } catch (err) {
      setError('Failed to browse directory');
    }
  };

  const handleMigration = async (oldPath: string) => {
    try {
      setMigrationProgress({ isRunning: true, step: 'Starting migration...', progress: 0 });
      
      // Create backup
      setMigrationProgress({ isRunning: true, step: 'Creating backup...', progress: 20 });
      
      // Start migration via MCP
      const response = await fetch('/api/mcp/migrate', {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify({
          tool: 'migrate_from_json',
          params: {
            oldPath,
            createBackup: true
          }
        })
      });
      
      if (!response.ok) {
        throw new Error('Migration failed');
      }
      
      setMigrationProgress({ isRunning: true, step: 'Migration completed', progress: 100 });
      
      // Reload stats
      await loadStorageStats();
      
      setSuccess('Migration completed successfully');
      
      setTimeout(() => {
        setMigrationProgress({ isRunning: false, step: '', progress: 0 });
      }, 2000);
      
    } catch (err) {
      setMigrationProgress({ 
        isRunning: false, 
        step: '', 
        progress: 0, 
        error: err instanceof Error ? err.message : 'Migration failed' 
      });
    }
  };

  const formatFileSize = (bytes: number): string => {
    const units = ['B', 'KB', 'MB', 'GB'];
    let size = bytes;
    let unitIndex = 0;
    
    while (size >= 1024 && unitIndex < units.length - 1) {
      size /= 1024;
      unitIndex++;
    }
    
    return `${size.toFixed(1)} ${units[unitIndex]}`;
  };

  return (
    <div className="space-y-6">
      {/* Current Directory Display */}
      <Card>
        <CardHeader>
          <CardTitle className="flex items-center gap-2">
            <Database className="h-5 w-5" />
            Current Data Directory
          </CardTitle>
          <CardDescription>
            Location where all story data is stored
          </CardDescription>
        </CardHeader>
        <CardContent>
          <div className="flex items-center gap-2 p-3 bg-gray-50 rounded-md">
            <FolderOpen className="h-4 w-4 text-gray-500" />
            <code className="text-sm font-mono">{currentDirectory}</code>
          </div>
          
          {/* Storage Statistics */}
          {stats && (
            <div className="mt-4 grid grid-cols-2 gap-4">
              <div className="text-center p-3 bg-blue-50 rounded-md">
                <div className="text-2xl font-bold text-blue-600">{stats.totalFiles}</div>
                <div className="text-sm text-blue-600">Total Files</div>
              </div>
              <div className="text-center p-3 bg-green-50 rounded-md">
                <div className="text-2xl font-bold text-green-600">{formatFileSize(stats.totalSize)}</div>
                <div className="text-sm text-green-600">Total Size</div>
              </div>
            </div>
          )}
          
          {/* Entity Breakdown */}
          {stats && (
            <div className="mt-4">
              <h4 className="text-sm font-medium mb-2">Story Elements</h4>
              <div className="flex flex-wrap gap-2">
                {Object.entries(stats.entitiesByType).map(([type, count]) => (
                  <Badge key={type} variant="secondary">
                    {type}: {count}
                  </Badge>
                ))}
              </div>
            </div>
          )}
        </CardContent>
      </Card>

      {/* Directory Selection */}
      <Card>
        <CardHeader>
          <CardTitle>Change Data Directory</CardTitle>
          <CardDescription>
            Select a new location for storing story data
          </CardDescription>
        </CardHeader>
        <CardContent className="space-y-4">
          {/* Recent Directories */}
          {recentDirectories.length > 0 && (
            <div>
              <Label htmlFor="recent-directories">Recent Directories</Label>
              <Select onValueChange={setNewDirectory}>
                <SelectTrigger>
                  <SelectValue placeholder="Select a recent directory..." />
                </SelectTrigger>
                <SelectContent>
                  {recentDirectories.map((dir) => (
                    <SelectItem key={dir} value={dir}>
                      {dir}
                    </SelectItem>
                  ))}
                </SelectContent>
              </Select>
            </div>
          )}

          {/* Manual Path Entry */}
          <div>
            <Label htmlFor="new-directory">Directory Path</Label>
            <div className="flex gap-2">
              <Input
                id="new-directory"
                value={newDirectory}
                onChange={(e) => setNewDirectory(e.target.value)}
                placeholder="Enter directory path..."
              />
              <Button
                type="button"
                variant="outline"
                onClick={handleBrowseDirectory}
              >
                Browse
              </Button>
            </div>
          </div>

          {/* Action Buttons */}
          <div className="flex gap-2">
            <Button
              onClick={handleDirectorySelect}
              disabled={isLoading || !newDirectory || newDirectory === currentDirectory}
            >
              {isLoading ? 'Changing...' : 'Change Directory'}
            </Button>
            
            <Dialog>
              <DialogTrigger asChild>
                <Button variant="outline">
                  Migrate Data
                </Button>
              </DialogTrigger>
              <DialogContent>
                <DialogHeader>
                  <DialogTitle>Migrate from Old Format</DialogTitle>
                </DialogHeader>
                <MigrationDialog
                  onMigrate={handleMigration}
                  progress={migrationProgress}
                  currentDirectory={currentDirectory}
                />
              </DialogContent>
            </Dialog>
          </div>
        </CardContent>
      </Card>

      {/* Status Messages */}
      {error && (
        <Alert variant="destructive">
          <AlertTriangle className="h-4 w-4" />
          <AlertDescription>{error}</AlertDescription>
        </Alert>
      )}
      
      {success && (
        <Alert>
          <CheckCircle className="h-4 w-4" />
          <AlertDescription>{success}</AlertDescription>
        </Alert>
      )}
    </div>
  );
};

interface MigrationDialogProps {
  onMigrate: (oldPath: string) => void;
  progress: MigrationProgress;
  currentDirectory: string;
}

const MigrationDialog: React.FC<MigrationDialogProps> = ({ onMigrate, progress, currentDirectory }) => {
  const [oldPath, setOldPath] = useState<string>('');

  const handleBrowseOldDirectory = async () => {
    try {
      const selected = await SelectDirectory();
      if (selected) {
        setOldPath(selected);
      }
    } catch (err) {
      // Handle error
    }
  };

  return (
    <div className="space-y-4">
      <div>
        <Label htmlFor="old-path">Old Data Directory</Label>
        <div className="flex gap-2">
          <Input
            id="old-path"
            value={oldPath}
            onChange={(e) => setOldPath(e.target.value)}
            placeholder="Select directory with JSON files..."
          />
          <Button
            type="button"
            variant="outline"
            onClick={handleBrowseOldDirectory}
            disabled={progress.isRunning}
          >
            Browse
          </Button>
        </div>
      </div>

      {progress.isRunning && (
        <div className="space-y-2">
          <div className="flex items-center gap-2">
            <Clock className="h-4 w-4 animate-spin" />
            <span className="text-sm">{progress.step}</span>
          </div>
          <div className="w-full bg-gray-200 rounded-full h-2">
            <div
              className="bg-blue-600 h-2 rounded-full transition-all duration-300"
              style={{ width: `${progress.progress}%` }}
            />
          </div>
        </div>
      )}

      {progress.error && (
        <Alert variant="destructive">
          <AlertTriangle className="h-4 w-4" />
          <AlertDescription>{progress.error}</AlertDescription>
        </Alert>
      )}

      <div className="text-sm text-gray-600">
        <p>This will migrate data from the old JSON format to the new folder-based format with versioning.</p>
        <p className="mt-1">Target directory: <code>{currentDirectory}</code></p>
      </div>

      <Button
        onClick={() => onMigrate(oldPath)}
        disabled={!oldPath || progress.isRunning}
        className="w-full"
      >
        {progress.isRunning ? 'Migrating...' : 'Start Migration'}
      </Button>
    </div>
  );
};

export default DataDirectoryManager;
