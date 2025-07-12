# AI Novel Prompter - Task Completion Workflow

## After Code Changes

### 1. Code Quality Checks
```bash
# Frontend linting and type checking
cd frontend
npm run lint
npx tsc --noEmit

# Go code formatting and vetting
go fmt ./...
go vet ./...
```

### 2. Generate Wails Bindings (If Backend Changed)
```bash
# Generate new bindings after Go backend changes
wails generate bindings
```

### 3. Testing
```bash
# Test in development mode
wails dev

# Manual testing of affected features
# - Test prompt generation
# - Test prose improvement
# - Test file operations
# - Test LLM provider integrations
```

### 4. Build Verification
```bash
# Clean build for production
wails build --clean

# Test the built executable
./build/bin/AINovelPrompter.exe  # Windows
# or
./build/bin/AINovelPrompter      # Linux/macOS
```

### 5. Documentation Updates
- Update README.md if new features added
- Update code comments if complex logic changed
- Update configuration examples if settings changed

## Release Workflow

### 1. Version Management
- Update version in `wails.json`
- Update version in `frontend/package.json` if needed
- Tag release in git

### 2. Cross-Platform Builds
```bash
# Build for all platforms
./scripts/build-windows.sh
./scripts/build-macos.sh
./scripts/build-macos-arm.sh
./scripts/build-macos-intel.sh
```

### 3. Distribution
- Create release binaries
- Package with NSIS installer (Windows)
- Create DMG installer (macOS)
- Upload to GitHub releases

## Development Best Practices

### Before Committing
1. Run linting and type checks
2. Test affected functionality
3. Ensure no console errors
4. Check for proper error handling
5. Verify configuration changes work

### Code Review Checklist
- [ ] TypeScript types are properly defined
- [ ] Error handling is implemented
- [ ] No hardcoded values (use constants)
- [ ] Components are properly organized
- [ ] Go functions follow naming conventions
- [ ] Memory management is proper

### Configuration Changes
- Test with different provider settings
- Verify .env file precedence
- Test configuration migration
- Document new configuration options

## Debugging Process

### Frontend Issues
1. Check browser console for errors
2. Use React DevTools
3. Verify Wails bindings are working
4. Check network requests for LLM providers

### Backend Issues
1. Check Go logs
2. Verify file permissions
3. Test Wails context functions
4. Validate configuration loading

### Integration Issues
1. Test Wails frontend-backend communication
2. Verify file operations work across platforms
3. Test LLM provider integrations
4. Check data persistence