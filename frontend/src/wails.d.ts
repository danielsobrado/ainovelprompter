// c:\Development\workspace\GitHub\ainovelprompter\frontend\src\wails.d.ts
interface Window {
  go: {
    main: {
      App: {
        ReadProsePromptsFile: () => Promise<string>;
        WriteProsePromptsFile: (content: string) => Promise<void>;
        // Add other Go methods exposed to the frontend here
        // For example, if you have ReadSettings, WriteSettings, etc.
        // ReadSettings: () => Promise<string>; 
        // WriteSettings: (settings: string) => Promise<void>;
      };
    };
  };
}
