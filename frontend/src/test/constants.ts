# Test configuration constants

# File size limits for test files (in bytes)
MAX_TEST_FILE_SIZE = 1024 * 1024  # 1MB

# Test timeouts (in milliseconds)
TEST_TIMEOUT = 30000  # 30 seconds
ASYNC_TIMEOUT = 5000  # 5 seconds

# Coverage thresholds
COVERAGE_THRESHOLDS = {
    "statements": 80,
    "branches": 75,
    "functions": 85,
    "lines": 80
}

# Mock data sizes
MAX_MOCK_CHARACTERS = 10
MAX_MOCK_LOCATIONS = 8
MAX_MOCK_RULES = 12
MAX_MOCK_CODEX_ENTRIES = 5
MAX_MOCK_SAMPLE_CHAPTERS = 6
MAX_MOCK_TASK_TYPES = 8

# Test data directories
TEST_DATA_DIR = "test_data"
TEMP_DIR_PREFIX = "ai_novel_prompter_test_"

# Default test values
DEFAULT_TIMEOUT = 10000
DEFAULT_RETRY_ATTEMPTS = 3
DEFAULT_DEBOUNCE_DELAY = 100

export {
    MAX_TEST_FILE_SIZE,
    TEST_TIMEOUT,
    ASYNC_TIMEOUT,
    COVERAGE_THRESHOLDS,
    MAX_MOCK_CHARACTERS,
    MAX_MOCK_LOCATIONS,
    MAX_MOCK_RULES,
    MAX_MOCK_CODEX_ENTRIES,
    MAX_MOCK_SAMPLE_CHAPTERS,
    MAX_MOCK_TASK_TYPES,
    TEST_DATA_DIR,
    TEMP_DIR_PREFIX,
    DEFAULT_TIMEOUT,
    DEFAULT_RETRY_ATTEMPTS,
    DEFAULT_DEBOUNCE_DELAY
}
