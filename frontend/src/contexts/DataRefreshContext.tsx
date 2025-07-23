import React, { createContext, useContext, useCallback } from 'react';

interface RefreshFunctions {
  refreshTaskTypes: () => Promise<void>;
  refreshRules: () => Promise<void>;
  refreshCharacters: () => Promise<void>;
  refreshLocations: () => Promise<void>;
  refreshCodex: () => Promise<void>;
  refreshSampleChapters: () => Promise<void>;
  refreshAll: () => Promise<void>;
}

interface DataRefreshContextType {
  refreshFunctions: RefreshFunctions | null;
  setRefreshFunctions: (functions: RefreshFunctions) => void;
  isRefreshing: boolean;
  lastRefresh: Date | null;
}

const DataRefreshContext = createContext<DataRefreshContextType>({
  refreshFunctions: null,
  setRefreshFunctions: () => {},
  isRefreshing: false,
  lastRefresh: null,
});

interface DataRefreshProviderProps {
  children: React.ReactNode;
}

export function DataRefreshProvider({ children }: DataRefreshProviderProps) {
  const [refreshFunctions, setRefreshFunctions] = React.useState<RefreshFunctions | null>(null);
  const [isRefreshing, setIsRefreshing] = React.useState(false);
  const [lastRefresh, setLastRefresh] = React.useState<Date | null>(null);

  const wrappedSetRefreshFunctions = useCallback((functions: RefreshFunctions) => {
    // Wrap refresh functions to track global refresh state
    const wrappedFunctions: RefreshFunctions = {
      refreshTaskTypes: async () => {
        setIsRefreshing(true);
        try {
          await functions.refreshTaskTypes();
          setLastRefresh(new Date());
        } finally {
          setIsRefreshing(false);
        }
      },
      refreshRules: async () => {
        setIsRefreshing(true);
        try {
          await functions.refreshRules();
          setLastRefresh(new Date());
        } finally {
          setIsRefreshing(false);
        }
      },
      refreshCharacters: async () => {
        setIsRefreshing(true);
        try {
          await functions.refreshCharacters();
          setLastRefresh(new Date());
        } finally {
          setIsRefreshing(false);
        }
      },
      refreshLocations: async () => {
        setIsRefreshing(true);
        try {
          await functions.refreshLocations();
          setLastRefresh(new Date());
        } finally {
          setIsRefreshing(false);
        }
      },
      refreshCodex: async () => {
        setIsRefreshing(true);
        try {
          await functions.refreshCodex();
          setLastRefresh(new Date());
        } finally {
          setIsRefreshing(false);
        }
      },
      refreshSampleChapters: async () => {
        setIsRefreshing(true);
        try {
          await functions.refreshSampleChapters();
          setLastRefresh(new Date());
        } finally {
          setIsRefreshing(false);
        }
      },
      refreshAll: async () => {
        setIsRefreshing(true);
        try {
          await functions.refreshAll();
          setLastRefresh(new Date());
        } finally {
          setIsRefreshing(false);
        }
      },
    };

    setRefreshFunctions(wrappedFunctions);
  }, []);

  const value: DataRefreshContextType = {
    refreshFunctions,
    setRefreshFunctions: wrappedSetRefreshFunctions,
    isRefreshing,
    lastRefresh,
  };

  return (
    <DataRefreshContext.Provider value={value}>
      {children}
    </DataRefreshContext.Provider>
  );
}

// Hook to use refresh functionality
export function useDataRefresh() {
  return useContext(DataRefreshContext);
}

export default DataRefreshContext;
