# AI Novel Prompter - Test Suite

## 📁 Test Organization

The test suite is organized into logical sections:

### 🏗️ Root Level Tests (`/tests/`)
Tests for main application functionality:
- **`app_test.go`** - Main application tests
- **`app_basic_test.go`** - Basic application functionality  
- **`app_quick_test.go`** - Quick validation tests (core functionality)
- **`app_extended_test.go`** - Extended functionality tests (migration, etc.)
- **`file_drop_handler_test.go`** - File drop functionality tests
- **`file_operations_test.go`** - File operations tests
- **`test_filename_fix.go`** - Custom test for filename parsing fix
- **`test_directory_switching.go`** - Custom test for directory switching

### 🔧 MCP Storage Tests (`/mcp/tests/`)
Tests for MCP storage subsystem:
- **`folder_storage_test.go`** - Folder storage implementation tests
- **`storage_test.go`** - General storage interface tests  
- **`migration_test.go`** - Data migration tests

### 📜 Test Scripts (`/tests/scripts/`)
Automation scripts for running tests:
- **`test_quick_verification.ps1`** - Quick verification of core fixes
- **`test_all_fixes.ps1`** - Comprehensive test suite
- **`test_comprehensive_fixes.ps1`** - Enhanced comprehensive testing
- **`test_specific_fixes.ps1`** - Individual fix testing
- **`test_filename_parsing.ps1`** - Filename parsing fix verification
- **`test_directory_debug.ps1`** - Directory switching debug
- **`run_tests.bat/.sh`** - Basic test runners

## 🚀 Quick Start

### Run Core Verification
```powershell
cd tests/scripts
.\test_quick_verification.ps1
```

### Run All Tests  
```powershell
cd tests/scripts
.\test_all_fixes.ps1
```

### Run Individual Test Categories
```bash
# Root level app tests
cd tests
go test -v ./app_quick_test.go ../app.go

# MCP storage tests  
cd mcp/tests
go test -v ./folder_storage_test.go ../storage/*.go

# Specific functionality tests
cd tests
go test -v -run TestBasicStorageOperations ./app_quick_test.go ../app.go
```

## 🔍 Test Categories

### ✅ **Core Functionality Tests** 
- Storage operations (create, read, update, delete)
- Character management 
- Legacy interface compatibility
- File naming conventions

### ✅ **Integration Tests**
- App integration with storage
- Data directory management
- Migration from old JSON format
- Version management

### ✅ **Performance Tests**
- Storage operation benchmarks
- Version history performance
- Cache efficiency

### ✅ **Edge Case Tests**
- Directory switching
- Filename parsing edge cases
- Error handling
- Concurrent access

## 🛠️ Test Utilities

### Key Fixes Verified
- ✅ **Filename parsing fix** - Resolves cache indexing issues
- ✅ **Migration functionality** - Ensures proper data migration  
- ✅ **Directory switching** - Validates multi-project support
- ✅ **Path separator normalization** - Cross-platform compatibility
- ✅ **Deadlock resolution** - Performance improvements maintained

### Test Data
- Tests use temporary directories for isolation
- No external dependencies required
- Automatic cleanup after test runs

## 📊 Expected Results

After applying the core fixes:
- ✅ **Quick tests should PASS** - Core functionality working
- ✅ **Migration tests should retrieve characters correctly** 
- ✅ **Storage statistics should count files properly**
- ✅ **Directory switching should maintain data integrity**

## 🔧 Troubleshooting

### Common Issues
1. **"0 characters after migration"** - Fixed by filename parsing improvement
2. **Directory switching failures** - Debug with `test_directory_debug.ps1`
3. **Path separator issues** - Resolved by normalization logic
4. **Build failures** - Check Go module dependencies

### Debug Scripts
- Use `test_directory_debug.ps1` for directory switching issues
- Use `test_filename_parsing.ps1` for cache indexing problems
- Check individual test outputs for specific failure details

## 📈 Performance Notes

- **Critical Performance Fix**: Deadlock resolution reduces operation time from 10+ minutes to milliseconds
- **Cache Optimization**: Filename parsing fix ensures proper cache rebuilding
- **Storage Efficiency**: Version management optimized for minimal overhead

## 🔄 Continuous Testing

For development workflow:
```bash
# Watch for changes and run tests
cd tests/scripts
.\test_quick_verification.ps1  # After each change
.\test_all_fixes.ps1          # Before commits
```

The test suite ensures the application maintains high quality and performance while supporting new features and fixes.
