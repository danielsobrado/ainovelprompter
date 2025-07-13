// Temporary type declarations for Wails bindings
declare global {
  interface Window {
    GetDataDirectory(): Promise<string>;
    GetRecentDataDirectories(): Promise<string[]>;
    GetStorageStats(): Promise<{
      totalFiles: number;
      totalSize: number;
      entitiesByType: Record<string, number>;
      versionsByType: Record<string, number>;
      oldestTimestamp: string;
      newestTimestamp: string;
    }>;
    ValidateDataDirectory(dir: string): Promise<boolean>;
    SetDataDirectory(dir: string): Promise<void>;
    AddRecentDataDirectory(dir: string): Promise<void>;
    SelectDirectory(): Promise<string>;
  }
}

export {};
