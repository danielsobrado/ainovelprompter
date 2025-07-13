// Test configuration constants

// File size limits for test files (in bytes)
export const MAX_TEST_FILE_SIZE = 1024 * 1024; // 1MB

// Test timeouts (in milliseconds)
export const TEST_TIMEOUT = 30000; // 30 seconds
export const ASYNC_TIMEOUT = 5000; // 5 seconds

// Coverage thresholds
export const COVERAGE_THRESHOLDS = {
    statements: 80,
    branches: 75,
    functions: 85,
    lines: 80
};

// Mock data sizes
export const MAX_MOCK_CHARACTERS = 10;
export const MAX_MOCK_LOCATIONS = 8;
export const MAX_MOCK_RULES = 12;
export const MAX_MOCK_CODEX_ENTRIES = 5;
export const MAX_MOCK_SAMPLE_CHAPTERS = 6;
export const MAX_MOCK_TASK_TYPES = 8;

// Test data directories
export const TEST_DATA_DIR = "test_data";
export const TEMP_DIR_PREFIX = "ai_novel_prompter_test_";

// Default test values
export const DEFAULT_TIMEOUT = 10000;
export const DEFAULT_RETRY_ATTEMPTS = 3;
export const DEFAULT_DEBOUNCE_DELAY = 100;
