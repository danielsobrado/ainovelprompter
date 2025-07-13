import React, { useState, useEffect } from 'react';
import { Button } from './ui/button';
import { Badge } from './ui/badge';
import { Card, CardContent, CardDescription, CardHeader, CardTitle } from './ui/card';
import { Dialog, DialogContent, DialogHeader, DialogTitle, DialogTrigger } from './ui/dialog';
import { Alert, AlertDescription } from './ui/alert';
import { Tabs, TabsContent, TabsList, TabsTrigger } from './ui/tabs';
import { ScrollArea } from './ui/scroll-area';
import { 
  History, 
  RotateCcw, 
  Eye, 
  Calendar, 
  User, 
  FileText, 
  AlertTriangle, 
  CheckCircle,
  Clock,
  Plus,
  Edit,
  Trash2
} from 'lucide-react';

interface Version {
  id: string;
  entityId: string;
  timestamp: string;
  operation: 'create' | 'update' | 'delete';
  filePath: string;
  active: boolean;
}

interface EntityVersionHistoryProps {
  entityType: string;
  entityId: string;
  entityName?: string;
  onVersionRestore?: (version: Version) => void;
}

interface VersionDiffProps {
  currentVersion: any;
  previousVersion: any;
  entityType: string;
}

const EntityVersionHistory: React.FC<EntityVersionHistoryProps> = ({
  entityType,
  entityId,
  entityName,
  onVersionRestore
}) => {
  const [versions, setVersions] = useState<Version[]>([]);
  const [selectedVersion, setSelectedVersion] = useState<Version | null>(null);
  const [selectedEntity, setSelectedEntity] = useState<any>(null);
  const [isLoading, setIsLoading] = useState(false);
  const [error, setError] = useState<string>('');
  const [success, setSuccess] = useState<string>('');

  useEffect(() => {
    loadVersionHistory();
  }, [entityType, entityId]);

  const loadVersionHistory = async () => {
    try {
      setIsLoading(true);
      setError('');

      // Call MCP to get version history
      const response = await fetch('/api/mcp/call', {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify({
          tool: 'get_entity_versions',
          params: {
            entityType,
            entityId
          }
        })
      });

      if (!response.ok) {
        throw new Error('Failed to load version history');
      }

      const data = await response.json();
      setVersions(data.result || []);
    } catch (err) {
      setError(err instanceof Error ? err.message : 'Failed to load version history');
    } finally {
      setIsLoading(false);
    }
  };

  const loadVersionData = async (version: Version) => {
    try {
      setIsLoading(true);

      const response = await fetch('/api/mcp/call', {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify({
          tool: 'get_entity_version',
          params: {
            entityType,
            entityId,
            timestamp: version.timestamp
          }
        })
      });

      if (!response.ok) {
        throw new Error('Failed to load version data');
      }

      const data = await response.json();
      setSelectedEntity(data.result);
      setSelectedVersion(version);
    } catch (err) {
      setError(err instanceof Error ? err.message : 'Failed to load version data');
    } finally {
      setIsLoading(false);
    }
  };

  const handleRestoreVersion = async (version: Version) => {
    try {
      setIsLoading(true);
      setError('');

      const response = await fetch('/api/mcp/call', {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify({
          tool: 'restore_entity_version',
          params: {
            entityType,
            entityId,
            timestamp: version.timestamp
          }
        })
      });

      if (!response.ok) {
        throw new Error('Failed to restore version');
      }

      setSuccess(`Version from ${formatTimestamp(version.timestamp)} restored successfully`);
      
      // Reload version history to show new restore point
      await loadVersionHistory();
      
      // Notify parent component
      if (onVersionRestore) {
        onVersionRestore(version);
      }
    } catch (err) {
      setError(err instanceof Error ? err.message : 'Failed to restore version');
    } finally {
      setIsLoading(false);
    }
  };

  const formatTimestamp = (timestamp: string): string => {
    return new Date(timestamp).toLocaleString();
  };

  const getOperationIcon = (operation: string) => {
    switch (operation) {
      case 'create':
        return <Plus className="h-4 w-4 text-green-600" />;
      case 'update':
        return <Edit className="h-4 w-4 text-blue-600" />;
      case 'delete':
        return <Trash2 className="h-4 w-4 text-red-600" />;
      default:
        return <FileText className="h-4 w-4 text-gray-600" />;
    }
  };

  const getOperationColor = (operation: string) => {
    switch (operation) {
      case 'create':
        return 'bg-green-100 text-green-800';
      case 'update':
        return 'bg-blue-100 text-blue-800';
      case 'delete':
        return 'bg-red-100 text-red-800';
      default:
        return 'bg-gray-100 text-gray-800';
    }
  };

  return (
    <div className="space-y-4">
      <div className="flex items-center justify-between">
        <div>
          <h3 className="text-lg font-semibold flex items-center gap-2">
            <History className="h-5 w-5" />
            Version History
          </h3>
          {entityName && (
            <p className="text-sm text-gray-600">
              {entityName} ({entityType})
            </p>
          )}
        </div>
        <Button
          variant="outline"
          size="sm"
          onClick={loadVersionHistory}
          disabled={isLoading}
        >
          {isLoading ? 'Loading...' : 'Refresh'}
        </Button>
      </div>

      {/* Version Timeline */}
      <Card>
        <CardHeader>
          <CardTitle className="text-base">Timeline</CardTitle>
          <CardDescription>
            {versions.length} version{versions.length !== 1 ? 's' : ''} found
          </CardDescription>
        </CardHeader>
        <CardContent>
          <ScrollArea className="h-64">
            <div className="space-y-3">
              {versions.map((version, index) => (
                <div
                  key={version.id}
                  className={`flex items-center justify-between p-3 rounded-lg border ${
                    version.active ? 'border-blue-200 bg-blue-50' : 'border-gray-200'
                  }`}
                >
                  <div className="flex items-center gap-3">
                    {getOperationIcon(version.operation)}
                    <div>
                      <div className="flex items-center gap-2">
                        <Badge className={getOperationColor(version.operation)}>
                          {version.operation}
                        </Badge>
                        {version.active && (
                          <Badge variant="outline">Current</Badge>
                        )}
                      </div>
                      <div className="text-sm text-gray-600 mt-1">
                        <Calendar className="h-3 w-3 inline mr-1" />
                        {formatTimestamp(version.timestamp)}
                      </div>
                    </div>
                  </div>
                  
                  <div className="flex gap-2">
                    <Dialog>
                      <DialogTrigger asChild>
                        <Button
                          variant="outline"
                          size="sm"
                          onClick={() => loadVersionData(version)}
                        >
                          <Eye className="h-4 w-4" />
                        </Button>
                      </DialogTrigger>
                      <DialogContent className="max-w-2xl max-h-[80vh]">
                        <DialogHeader>
                          <DialogTitle>
                            Version Details - {formatTimestamp(version.timestamp)}
                          </DialogTitle>
                        </DialogHeader>
                        <VersionViewer
                          version={version}
                          entity={selectedEntity}
                          entityType={entityType}
                        />
                      </DialogContent>
                    </Dialog>
                    
                    {!version.active && version.operation !== 'delete' && (
                      <Button
                        variant="outline"
                        size="sm"
                        onClick={() => handleRestoreVersion(version)}
                        disabled={isLoading}
                      >
                        <RotateCcw className="h-4 w-4" />
                      </Button>
                    )}
                  </div>
                </div>
              ))}
            </div>
          </ScrollArea>
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

interface VersionViewerProps {
  version: Version;
  entity: any;
  entityType: string;
}

const VersionViewer: React.FC<VersionViewerProps> = ({ version, entity, entityType }) => {
  if (!entity) {
    return (
      <div className="flex items-center justify-center h-32">
        <Clock className="h-6 w-6 animate-spin mr-2" />
        Loading version data...
      </div>
    );
  }

  const renderEntityData = (data: any) => {
    const excludeFields = ['id', 'createdAt', 'updatedAt'];
    
    return (
      <div className="space-y-3">
        {Object.entries(data)
          .filter(([key]) => !excludeFields.includes(key))
          .map(([key, value]) => (
            <div key={key}>
              <label className="text-sm font-medium text-gray-700 capitalize">
                {key.replace(/([A-Z])/g, ' $1').toLowerCase()}
              </label>
              <div className="mt-1">
                {typeof value === 'string' && value.length > 100 ? (
                  <ScrollArea className="h-32">
                    <pre className="text-sm bg-gray-50 p-2 rounded whitespace-pre-wrap">
                      {value}
                    </pre>
                  </ScrollArea>
                ) : (
                  <div className="text-sm bg-gray-50 p-2 rounded">
                    {typeof value === 'object' ? JSON.stringify(value, null, 2) : String(value)}
                  </div>
                )}
              </div>
            </div>
          ))}
      </div>
    );
  };

  return (
    <Tabs defaultValue="content" className="w-full">
      <TabsList className="grid w-full grid-cols-2">
        <TabsTrigger value="content">Content</TabsTrigger>
        <TabsTrigger value="metadata">Metadata</TabsTrigger>
      </TabsList>
      
      <TabsContent value="content" className="mt-4">
        <ScrollArea className="h-96">
          {renderEntityData(entity)}
        </ScrollArea>
      </TabsContent>
      
      <TabsContent value="metadata" className="mt-4">
        <div className="space-y-3">
          <div>
            <label className="text-sm font-medium text-gray-700">Version ID</label>
            <div className="text-sm bg-gray-50 p-2 rounded font-mono">{version.id}</div>
          </div>
          <div>
            <label className="text-sm font-medium text-gray-700">Operation</label>
            <div className="text-sm bg-gray-50 p-2 rounded">{version.operation}</div>
          </div>
          <div>
            <label className="text-sm font-medium text-gray-700">Timestamp</label>
            <div className="text-sm bg-gray-50 p-2 rounded">
              {new Date(version.timestamp).toISOString()}
            </div>
          </div>
          <div>
            <label className="text-sm font-medium text-gray-700">File Path</label>
            <div className="text-sm bg-gray-50 p-2 rounded font-mono">{version.filePath}</div>
          </div>
          <div>
            <label className="text-sm font-medium text-gray-700">Status</label>
            <div className="text-sm bg-gray-50 p-2 rounded">
              {version.active ? 'Active (Current)' : 'Inactive (Historical)'}
            </div>
          </div>
          {entity.createdAt && (
            <div>
              <label className="text-sm font-medium text-gray-700">Created At</label>
              <div className="text-sm bg-gray-50 p-2 rounded">
                {new Date(entity.createdAt).toLocaleString()}
              </div>
            </div>
          )}
          {entity.updatedAt && (
            <div>
              <label className="text-sm font-medium text-gray-700">Updated At</label>
              <div className="text-sm bg-gray-50 p-2 rounded">
                {new Date(entity.updatedAt).toLocaleString()}
              </div>
            </div>
          )}
        </div>
      </TabsContent>
    </Tabs>
  );
};

export default EntityVersionHistory;
