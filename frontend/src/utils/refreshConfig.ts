// Refresh configuration constants
export const REFRESH_CONFIG = {
  // Auto-refresh intervals (in milliseconds)
  AUTO_REFRESH_INTERVAL: 30000, // 30 seconds
  STORAGE_STATS_REFRESH_INTERVAL: 30000, // 30 seconds
  
  // Manual refresh debounce (prevent rapid clicking)
  MANUAL_REFRESH_DEBOUNCE: 500, // 0.5 seconds
  
  // Refresh on focus settings
  REFRESH_ON_WINDOW_FOCUS: true,
  REFRESH_ON_DROPDOWN_OPEN: true,
  
  // Visual feedback
  SHOW_REFRESH_ANIMATIONS: true,
  SHOW_LAST_REFRESH_TIME: true,
} as const;

export type RefreshConfig = typeof REFRESH_CONFIG;
