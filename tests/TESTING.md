## Test Suite for AI Novel Prompter

This document outlines the comprehensive test suite for the AI Novel Prompter application with folder-based storage and versioning.

### Overview

The test suite covers:
- **Frontend Tests**: React components, hooks, utilities, and integration tests
- **Backend Tests**: Go functions, folder storage with versioning, and API endpoints
- **Storage System Tests**: Version management, migration, and data integrity
- **Integration Tests**: End-to-end functionality and build verification
- **Code Quality**: Linting, formatting, and security scanning

### Test Structure

```
├── frontend/
│   ├── src/
│   │   ├── __tests__/           # Integration tests
│   │   ├── components/
│   │   │   └── **/__tests__/    # Component tests
│   │   ├── hooks/
│   │   │   └── __tests__/       # Hook tests
│   │   ├── utils/
│   │   │   └── __tests__/       # Utility tests
│   │   └── test/
│   │       ├── setup.ts         # Test setup and configuration
│   │       └── mocks.ts         # Mock data and utilities
│   ├── vitest.config.ts         # Vitest configuration
│   └── package.json             # Test scripts and dependencies
├── mcp/
│   ├── storage/
│   │   ├── folder_storage_test.go      # Folder storage system tests
│   │   ├── migration_test.go           # Migration tests
│   │   └── storage_test.go             # Legacy storage tests
│   ├── server_test.go                  # MCP server tests
│   └── TESTING.md                      # MCP-specific testing guide
├── cmd/
│   └── test-storage/
│       └── main.go                     # Comprehensive storage test program
├── *_test.go                           # Go unit tests
├── scripts/test.sh                     # Test runner script
└── .github/workflows/test.yml          # CI/CD pipeline
```

### Running Tests

#### Quick Start
```bash
# Run all tests
./scripts/test.sh

# Frontend tests only
./scripts/test.sh --frontend-only

# Backend tests only
./scripts/test.sh --backend-only

# Storage system tests
./scripts/test.sh --storage-only

# With coverage
./scripts/test.sh --coverage
```

#### Frontend Tests
```bash
cd frontend

# Run tests once
npm run test:run

# Run with coverage
npm run test:coverage

# Watch mode
npm run test:watch

# UI mode
npm run test:ui
```

#### Backend Tests
```bash
# Run all Go tests
go test ./...

# Storage system tests specifically
go test ./mcp/storage/...

# With coverage
go test -v -race -coverprofile=coverage.out ./...
go tool cover -html=coverage.out -o coverage.html

# Specific package
go test ./app_test.go
```

#### Storage System Test Program
```bash
# Run comprehensive storage test
cd cmd/test-storage
go run main.go
```

### Test Categories

#### 1. Unit Tests

**Frontend Components:**
- `CharactersSelector` - Character selection and management
- `TaskTypeSelector` - Task type selection functionality
- `FinalPrompt` - Prompt display and copy functionality
- `SettingsModal` - Application settings management
- `ProseImprovementTab` - Prose improvement interface
- `DataDirectoryManager` - Data directory management interface
- `EntityVersionHistory` - Version history and restore functionality

**Frontend Hooks:**
- `useOptionManagement` - Data loading and persistence with versioning
- `usePromptGeneration` - Prompt generation logic
- `useLLMProvider` - LLM integration
- `useVersionHistory` - Entity version management

**Frontend Utilities:**
- Helper functions (token counting, text formatting, validation)
- Prompt generators (ChatGPT and Claude formats)
- Storage utilities with version support

**Backend Functions:**
- Folder storage operations (create, read, update, delete)
- Version management (create, restore, cleanup)
- Migration from legacy JSON format
- Settings management with data directory support
- Character/location/rule management with versioning
- File drop handling
- Directory processing and validation

#### 2. Storage System Tests

**Folder Storage (`folder_storage_test.go`):**
- Basic CRUD operations with versioning
- Version history management
- Entity restoration to previous versions
- Legacy interface compatibility
- File naming conventions and slugification
- Storage statistics and analytics
- Cleanup of old versions
- Data directory management
- Atomic operations and data integrity
- Concurrent access handling
- Performance benchmarks

**Migration System (`migration_test.go`):**
- Migration from old JSON format
- Backup creation during migration
- Data validation post-migration
- Handling of missing or empty files
- Migration rollback capabilities
- Comprehensive entity type support

**Legacy Storage (`storage_test.go`):**
- Original file storage system tests
- Backward compatibility verification
- Data persistence across instances
- Error handling and edge cases
- Performance benchmarks

#### 3. Integration Tests

**Frontend Integration:**
- App component rendering with new storage
- Tab switching functionality
- Prompt generation workflow with versioned data
- Data persistence across components
- Version history UI interactions
- Data directory switching

**Backend Integration:**
- Folder storage operations
- Data directory management
- Configuration loading with custom directories
- MCP server integration with versioned storage

**MCP Server Integration (`server_test.go`):**
- Tool discovery and registration
- Character operations with versioning
- Chapter management
- Prose improvement workflows
- Search and analysis functions
- Prompt generation tools
- Version management tools
- Storage statistics tools
- Migration tools

#### 4. End-to-End Tests

**Build Verification:**
- Application builds successfully with new storage
- All assets are generated
- Dependencies are correctly resolved
- CLI options work correctly

**Storage Test Program (`cmd/test-storage/main.go`):**
- Comprehensive test of all storage operations
- Version management workflows
- Migration testing
- Data directory operations
- Statistical analysis
- Performance validation

### Test Configuration

#### Frontend (Vitest)
- **Environment**: jsdom for DOM simulation
- **Coverage**: v8 provider with HTML and JSON reports
- **Mocking**: Wails runtime, browser APIs, and storage operations
- **Setup**: Custom test utilities and mock data
- **Version History**: Mock version management functions

#### Backend (Go testing)
- **Coverage**: Atomic mode with race detection
- **Assertions**: testify/assert for readable tests
- **Isolation**: Temporary directories for storage operations
- **Cleanup**: Proper resource cleanup in all tests
- **Concurrent Testing**: Race condition detection
- **Performance**: Benchmark tests for storage operations

### Continuous Integration

**GitHub Actions Pipeline:**
1. **Backend Tests**: Go tests with coverage including storage system
2. **Storage Tests**: Dedicated storage system validation
3. **Migration Tests**: Legacy format migration validation
4. **Frontend Tests**: Vitest tests with coverage
5. **Integration Tests**: Build verification with new CLI options
6. **Security Scan**: Gosec and npm audit
7. **Code Quality**: Formatting and linting

**Quality Gates:**
- Minimum 80% code coverage (85% for storage system)
- All linting rules pass
- No high-severity security vulnerabilities
- All tests pass in CI environment
- Storage integrity validation
- Migration success validation

### Mock Data and Utilities

**Mock Data:**
- Sample characters, locations, rules, codex entries with versions
- Task types and sample chapters
- Settings configurations with data directories
- File content examples
- Version history mock data
- Storage statistics mock data

**Test Utilities:**
- Wails API mocking including new storage functions
- Event simulation helpers
- Async operation utilities
- Component rendering helpers
- Storage operation mocks
- Version management utilities

### Coverage Reports

**Frontend Coverage:**
- Line coverage: Target 85%+
- Branch coverage: Target 80%+
- Function coverage: Target 90%+
- Version management: Target 95%+

**Backend Coverage:**
- Package coverage: Target 80%+
- Storage system: Target 95%+
- Critical paths: 100% coverage required
- Error handling: All error paths tested
- Migration logic: 100% coverage required

### Storage System Validation

**Data Integrity Tests:**
- Atomic operations under concurrent access
- Corruption recovery mechanisms
- Version consistency validation
- File system permission handling
- Large dataset performance

**Version Management Tests:**
- Create/update/delete version tracking
- Version restoration accuracy
- Timeline consistency
- Cleanup operation safety
- Storage space optimization

**Migration Tests:**
- JSON to folder format conversion
- Data preservation validation
- Backup creation and verification
- Rollback capabilities
- Edge case handling (empty files, corrupt data)

### Best Practices

1. **Test Isolation**: Each test runs independently with clean storage
2. **Descriptive Names**: Clear test descriptions including version scenarios
3. **Arrange-Act-Assert**: Consistent test structure
4. **Mock External Dependencies**: No real file I/O in unit tests
5. **Error Path Testing**: Test both success and failure scenarios
6. **Performance**: Storage tests complete efficiently
7. **Maintenance**: Regular test review and cleanup
8. **Version Scenarios**: Test all version management workflows
9. **Migration Safety**: Validate all migration paths

### Adding New Tests

1. **Component Tests**: Create `__tests__` directory next to component
2. **Hook Tests**: Add to `hooks/__tests__/`
3. **Utility Tests**: Add to `utils/__tests__/`
4. **Storage Tests**: Add to `mcp/storage/*_test.go`
5. **Backend Tests**: Create `*_test.go` files next to source
6. **Integration**: Add to main `__tests__` directories
7. **Version Tests**: Include version management scenarios

### Debugging Tests

```bash
# Frontend debug mode
npm run test:ui

# Backend verbose output
go test -v ./...

# Storage tests specifically
go test -v ./mcp/storage/...

# Single test
go test -run TestSpecificFunction

# Frontend single test
npm test -- --run "specific test name"

# Run storage test program
cd cmd/test-storage && go run main.go
```

### Known Issues and Limitations

1. **Wails Mocking**: Complex Wails runtime interactions require comprehensive mocking
2. **File System**: Tests use temporary directories to avoid real file system interference
3. **Async Operations**: Proper handling of promises and async state updates
4. **Browser APIs**: Clipboard and other browser APIs are mocked for testing
5. **Version History**: Large version histories may impact test performance
6. **Migration Testing**: Requires careful setup of legacy data formats

### Storage System Specific Testing

**Test Program Output:**
```
🧪 AI Novel Prompter - Folder Storage Test Suite
================================================

📁 Test 1: Basic Storage Operations
✓ Created character with version ID: abc123
✓ Updated character with version ID: def456
✓ Retrieved latest character: Test Character - Updated description
✓ Successfully deleted character

📚 Test 2: Version Management
✓ Found 3 versions:
  1. 2025-01-13T15:30:00Z: create (active: false)
  2. 2025-01-13T15:31:00Z: update (active: false)
  3. 2025-01-13T15:32:00Z: update (active: true)
✓ Restored to version: ghi789
✓ Restoration verified: description = 'Updated description'

🔄 Test 3: Legacy Interface Compatibility
✓ Retrieved 2 characters via legacy interface
✓ Search found 2 characters
✓ Retrieved character by ID: Legacy Test 1

🏗️ Test 4: Multiple Entity Types
✓ Created location: Test Location
✓ Created codex entry: Magic System
✓ Created rule: Character Development
✓ Retrieved 1 locations
✓ Retrieved 1 codex entries
✓ Retrieved 1 rules

📊 Test 5: Storage Statistics
✓ Storage Statistics:
  - Total files: 6
  - Total size: 1024 bytes
  - Entities by type:
    * characters: 2 entities
    * locations: 1 entities
    * codex: 1 entities
    * rules: 1 entities
  - Versions by type:
    * characters: 4 versions
    * locations: 1 versions
    * codex: 1 versions
    * rules: 1 versions

🔄 Test 6: Migration from JSON
✓ Created old format data with 2 characters and 1 locations
✓ Migration completed successfully
✓ Verified migration: 2 characters, 1 locations

📂 Test 7: Data Directory Management
✓ Current data directory: /tmp/test_dir
✓ Successfully changed data directory to: /tmp/test_dir_new

🎉 All tests passed! Folder-based storage with versioning is working correctly.
✅ Storage system is ready for production use.
```

### Future Improvements

1. **E2E Testing**: Add Playwright or Cypress for full user journey testing
2. **Performance Testing**: Add performance benchmarks for critical paths
3. **Visual Regression**: Screenshot testing for UI components
4. **Load Testing**: Stress testing for large version histories
5. **Accessibility Testing**: Automated a11y testing integration
6. **Version Conflict Resolution**: Test concurrent version creation scenarios
7. **Storage Optimization**: Test compression and archival strategies

### Troubleshooting

**Common Issues:**

1. **Tests timeout**: Increase timeout in vitest.config.ts
2. **Wails mocks not working**: Check mock setup in test/setup.ts
3. **File permission errors**: Ensure test cleanup is working properly
4. **Version conflicts**: Check timestamp generation in tests
5. **Migration failures**: Verify test data format matches expected input

**Solutions:**

```bash
# Clear test cache
rm -rf frontend/node_modules/.vitest

# Clear storage test data
rm -rf /tmp/ainovelprompter_test_*

# Rebuild and retest
npm run build && npm run test

# Check Go module cache
go clean -modcache

# Verbose storage testing
cd cmd/test-storage && go run main.go
```

### Conclusion

This comprehensive test suite ensures the reliability and maintainability of the AI Novel Prompter application with its new folder-based storage and versioning system. The combination of unit tests, integration tests, storage system validation, and continuous integration provides confidence in the data integrity and version management capabilities.

The storage system tests are particularly critical as they validate:
- **Data integrity** under all conditions
- **Version management** accuracy and consistency
- **Migration safety** from legacy formats
- **Performance** with large datasets
- **Concurrent access** handling

Regular maintenance and updates to the test suite are essential as the application evolves. All contributors should follow the testing best practices and ensure new features include appropriate test coverage, especially for version management functionality.
