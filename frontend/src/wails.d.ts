// c:\Development\workspace\GitHub\ainovelprompter\frontend\src\wails.d.ts
interface Window {
  runtime?: any;
  go: {
    main: {
      App: {
        ReadProsePromptsFile: () => Promise<string>;
        WriteProsePromptsFile: (content: string) => Promise<void>;
        GetInitialLLMSettings: () => Promise<any>;
        ReadLLMSettingsFile: () => Promise<string>;
        WriteLLMSettingsFile: (content: string) => Promise<void>;
        ReadPromptDefinitionsFile: () => Promise<string>;
        WritePromptDefinitionsFile: (jsonData: string) => Promise<void>;
        GetResolvedProsePrompt: (taskID: string, providerJSON: string) => Promise<string>;
        // Add other Go methods exposed to the frontend here
        // For example, if you have ReadSettings, WriteSettings, etc.
        // ReadSettings: () => Promise<string>; 
        // WriteSettings: (settings: string) => Promise<void>;
      };
    };
  };
}
