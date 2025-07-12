## Test Suite for AI Novel Prompter

This document outlines the comprehensive test suite for the AI Novel Prompter application.

### Overview

The test suite covers:
- **Frontend Tests**: React components, hooks, utilities, and integration tests
- **Backend Tests**: Go functions, file operations, and API endpoints
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
├── *.test.go                    # Go unit tests
├── scripts/test.sh              # Test runner script
└── .github/workflows/test.yml   # CI/CD pipeline
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

# With coverage
go test -v -race -coverprofile=coverage.out ./...
go tool cover -html=coverage.out -o coverage.html

# Specific package
go test ./app_test.go
```

### Test Categories

#### 1. Unit Tests

**Frontend Components:**
- `CharactersSelector` - Character selection and management
- `TaskTypeSelector` - Task type selection functionality
- `FinalPrompt` - Prompt display and copy functionality
- `SettingsModal` - Application settings management
- `ProseImprovementTab` - Prose improvement interface

**Frontend Hooks:**
- `useOptionManagement` - Data loading and persistence
- `usePromptGeneration` - Prompt generation logic
- `useLLMProvider` - LLM integration

**Frontend Utilities:**
- Helper functions (token counting, text formatting, validation)
- Prompt generators (ChatGPT and Claude formats)
- Storage utilities

**Backend Functions:**
- File operations (read/write JSON files)
- Settings management
- Character/location/rule management
- File drop handling
- Directory processing

#### 2. Integration Tests

**Frontend Integration:**
- App component rendering
- Tab switching functionality
- Prompt generation workflow
- Data persistence across components

**Backend Integration:**
- File system operations
- Data directory management
- Configuration loading

#### 3. End-to-End Tests

**Build Verification:**
- Application builds successfully
- All assets are generated
- Dependencies are correctly resolved

### Test Configuration

#### Frontend (Vitest)
- **Environment**: jsdom for DOM simulation
- **Coverage**: v8 provider with HTML and JSON reports
- **Mocking**: Wails runtime and browser APIs
- **Setup**: Custom test utilities and mock data

#### Backend (Go testing)
- **Coverage**: Atomic mode with race detection
- **Assertions**: testify/assert for readable tests
- **Isolation**: Temporary directories for file operations
- **Cleanup**: Proper resource cleanup in all tests

### Continuous Integration

**GitHub Actions Pipeline:**
1. **Backend Tests**: Go tests with coverage
2. **Frontend Tests**: Vitest tests with coverage
3. **Integration Tests**: Build verification
4. **Security Scan**: Gosec and npm audit
5. **Code Quality**: Formatting and linting

**Quality Gates:**
- Minimum 80% code coverage
- All linting rules pass
- No high-severity security vulnerabilities
- All tests pass in CI environment

### Mock Data and Utilities

**Mock Data:**
- Sample characters, locations, rules, codex entries
- Task types and sample chapters
- Settings configurations
- File content examples

**Test Utilities:**
- Wails API mocking
- Event simulation helpers
- Async operation utilities
- Component rendering helpers

### Coverage Reports

**Frontend Coverage:**
- Line coverage: Target 85%+
- Branch coverage: Target 80%+
- Function coverage: Target 90%+

**Backend Coverage:**
- Package coverage: Target 80%+
- Critical paths: 100% coverage required
- Error handling: All error paths tested

### Best Practices

1. **Test Isolation**: Each test runs independently
2. **Descriptive Names**: Clear test descriptions
3. **Arrange-Act-Assert**: Consistent test structure
4. **Mock External Dependencies**: No real file I/O in tests
5. **Error Path Testing**: Test both success and failure scenarios
6. **Performance**: Tests complete in under 30 seconds
7. **Maintenance**: Regular test review and cleanup

### Adding New Tests

1. **Component Tests**: Create `__tests__` directory next to component
2. **Hook Tests**: Add to `hooks/__tests__/`
3. **Utility Tests**: Add to `utils/__tests__/`
4. **Backend Tests**: Create `*_test.go` files next to source
5. **Integration**: Add to main `__tests__` directories

### Debugging Tests

```bash
# Frontend debug mode
npm run test:ui

# Backend verbose output
go test -v ./...

# Single test
go test -run TestSpecificFunction

# Frontend single test
npm test -- --run "specific test name"
```

### Known Issues and Limitations

1. **Wails Mocking**: Complex Wails runtime interactions require comprehensive mocking
2. **File System**: Tests use temporary directories to avoid real file system interference
3. **Async Operations**: Proper handling of promises and async state updates
4. **Browser APIs**: Clipboard and other browser APIs are mocked for testing

### Future Improvements

1. **E2E Testing**: Add Playwright or Cypress for full user journey testing
2. **Performance Testing**: Add performance benchmarks for critical paths
3. **Visual Regression**: Screenshot testing for UI components
4. **Load Testing**: Stress testing for file operations and large prompts
5. **Accessibility Testing**: Automated a11y testing integration

### Troubleshooting

**Common Issues:**

1. **Tests timeout**: Increase timeout in vitest.config.ts
2. **Wails mocks not working**: Check mock setup in test/setup.ts
3. **File permission errors**: Ensure test cleanup is working properly
4. **Coverage not updating**: Clear coverage cache and rebuild

**Solutions:**

```bash
# Clear test cache
rm -rf frontend/node_modules/.vitest

# Rebuild and retest
npm run build && npm run test

# Check Go module cache
go clean -modcache
```

### Conclusion

This comprehensive test suite ensures the reliability and maintainability of the AI Novel Prompter application. The combination of unit tests, integration tests, and continuous integration provides confidence in code changes and helps prevent regressions.

Regular maintenance and updates to the test suite are essential as the application evolves. All contributors should follow the testing best practices and ensure new features include appropriate test coverage.
