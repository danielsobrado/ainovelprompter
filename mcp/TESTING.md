# MCP Server Testing Guide

## Quick Start (Windows)

### Option 1: Automated Testing (Recommended)
```powershell
# Run the comprehensive test script
.\test.ps1
```

### Option 2: Manual Testing
```bash
# 1. Test compilation and basic functionality
go run test_comprehensive.go

# 2. Test new folder storage system
cd ../cmd/test-storage
go run main.go

# 3. Start HTTP server for web-based testing
go run http_server.go
# Then visit: http://localhost:8080/test

# 4. Test basic MCP server
go run main.go
```

## Testing Options

### 1. 🧪 Storage System Test (`cmd/test-storage/main.go`)
- **NEW**: Tests folder-based storage with versioning
- Validates version management (create, update, restore)
- Tests migration from old JSON format
- Validates data directory management
- Tests legacy interface compatibility
- Performance and statistics validation

**Usage:**
```bash
cd cmd/test-storage
go run main.go
```

**Expected Output:**
```
🧪 AI Novel Prompter - Folder Storage Test Suite
================================================

📁 Test 1: Basic Storage Operations
✓ Created character with version ID: abc123_20250113_150000_create
✓ Updated character with version ID: abc123_20250113_150001_update
✓ Retrieved latest character: Test Character - Updated description
✓ Successfully deleted character

📚 Test 2: Version Management
✓ Found 3 versions with proper timestamps
✓ Restored to previous version successfully
✓ Restoration verified with correct data

🔄 Test 3: Legacy Interface Compatibility
✓ All legacy MCP tools work unchanged
✓ Backward compatibility maintained

🏗️ Test 4: Multiple Entity Types
✓ Characters, locations, codex, rules all supported
✓ Version tracking for all entity types

📊 Test 5: Storage Statistics
✓ Storage analytics working correctly
✓ Version counts and file sizes accurate

🔄 Test 6: Migration from JSON
✓ Old JSON format detected and migrated
✓ All data preserved with version history

📂 Test 7: Data Directory Management
✓ CLI data directory options working
✓ Dynamic directory switching supported

🎉 All tests passed! Storage system ready for production.
```

### 2. 🚀 Comprehensive MCP Test (`test_comprehensive.go`)
- Tests all MCP tools including **NEW version management tools**
- Validates data operations (create, read, update, delete, restore)
- Tests new storage statistics and cleanup functions
- Tests error handling for versioned operations

**Usage:**
```bash
go run test_comprehensive.go
```

**Expected Output:**
```
=== AI Novel Prompter MCP Server Test ===
✅ MCP Server initialized with folder storage!
✅ Found 50+ MCP tools (includes new version management):
   • Story Context: 15 tools
   • Chapter Management: 12 tools  
   • Prose Improvement: 8 tools
   • Search & Analysis: 4 tools
   • Prompt Generation: 2 tools
   • Version Management: 5 tools (NEW)
   • Storage Management: 4 tools (NEW)
   • Migration Tools: 2 tools (NEW)

📚 Testing Version Management Tools:
✅ get_entity_versions - Retrieved version history
✅ get_entity_version - Retrieved specific version
✅ restore_entity_version - Restored to previous version
✅ get_storage_stats - Retrieved storage statistics
✅ cleanup_old_versions - Cleaned up old versions

🔄 Testing Storage Management:
✅ set_data_directory - Changed data directory
✅ get_data_directory - Retrieved current directory
✅ migrate_from_json - Migrated old format data

🎉 All tests completed successfully with versioning!
```

### 3. 🌐 HTTP API Server (`http_server.go`)
- Exposes all MCP tools via REST API
- **NEW**: Includes version management endpoints
- Web-based testing interface with storage analytics
- JSON responses for easy integration

**Usage:**
```bash
go run http_server.go
```

**New Endpoints:**
- `GET /` - API documentation (updated)
- `GET /tools` - List all MCP tools (now 50+)
- `GET /test` - Run automated tests (includes storage tests)
- `POST /execute` - Execute MCP tool (includes version tools)
- `GET /storage/stats` - **NEW**: Storage statistics
- `GET /version/history/{entity}/{id}` - **NEW**: Version history
- `POST /version/restore` - **NEW**: Restore version

**Test URLs:**
- http://localhost:8080/ (updated with new tools)
- http://localhost:8080/tools (shows version management tools)
- http://localhost:8080/test (includes storage system tests)
- http://localhost:8080/storage/stats (NEW - storage analytics)

### 4. 🔧 Basic MCP Server (`main.go`)
- Simple MCP server demonstration
- Shows tool discovery including new version tools
- Minimal output for debugging storage operations

**Usage:**
```bash
go run main.go
```

## PowerShell Testing Scripts

### `test.ps1` - Full Test Suite (Updated)
- Checks Go installation
- Downloads dependencies
- Compiles all components
- **NEW**: Runs storage system tests
- **NEW**: Tests version management
- **NEW**: Validates migration functionality
- Runs comprehensive MCP tests
- Optionally starts HTTP server

### `test_api.ps1` - HTTP API Testing (Updated)
- Tests all HTTP endpoints including new ones
- **NEW**: Tests version management APIs
- **NEW**: Tests storage statistics endpoints
- Validates API responses for versioned data
- Tests error handling for version operations

**Usage:**
```powershell
# Terminal 1: Start HTTP server
.\http_server.exe

# Terminal 2: Run API tests (now includes version APIs)
.\test_api.ps1
```

## New Version Management Tools

### MCP Tools Added
1. **`get_entity_versions`** - Get version history for any entity
2. **`get_entity_version`** - Retrieve specific version by timestamp
3. **`restore_entity_version`** - Restore entity to previous version
4. **`get_storage_stats`** - Storage usage statistics and analytics
5. **`cleanup_old_versions`** - Remove old versions based on retention policy
6. **`set_data_directory`** - Change data directory at runtime
7. **`get_data_directory`** - Get current data directory
8. **`migrate_from_json`** - Migrate from old JSON format

### Testing Version Management

**Manual Tool Testing:**
```go
// Test version creation
result, err := server.ExecuteTool("create_character", map[string]interface{}{
    "name": "Test Character",
    "description": "A test character",
})

// Get version history
versions, err := server.ExecuteTool("get_entity_versions", map[string]interface{}{
    "entityType": "characters",
    "entityId": characterID,
})

// Restore to previous version
restored, err := server.ExecuteTool("restore_entity_version", map[string]interface{}{
    "entityType": "characters", 
    "entityId": characterID,
    "timestamp": "2025-01-13T15:30:00Z",
})

// Get storage statistics
stats, err := server.ExecuteTool("get_storage_stats", map[string]interface{}{})
```

## Expected Results

### ✅ Success Indicators
- All tools compile without errors including storage system
- MCP server initializes with folder storage successfully
- 50+ tools are discovered (including version management tools)
- Basic operations work with version tracking
- Version history and restore operations function correctly
- Storage statistics provide accurate data
- Migration from old JSON format works seamlessly
- Error handling functions correctly for version operations
- HTTP endpoints return valid JSON for all operations

### ❌ Common Issues & Solutions

**"Storage directory not accessible"**
- Check file permissions for data directory
- Verify parent directory exists and is writable
- Check disk space availability

**"Version timestamp format error"**
- Ensure timestamps are in RFC3339 format
- Check system clock accuracy
- Verify timezone handling

**"Migration failed"**
- Check old JSON file format and structure
- Verify backup directory permissions
- Ensure sufficient disk space for migration

**"Version restore failed"**
- Verify target version exists
- Check entity ID format
- Ensure restore permissions

**"Legacy interface broken"**
- Check backward compatibility layer
- Verify legacy tool registration
- Test with existing data

## Data Storage Evolution

### Old Format (Pre-Versioning)
```
~/.ai-novel-prompter/
├── characters.json
├── locations.json
├── chapters.json
└── settings.json
```

### New Format (With Versioning)
```
~/.ai-novel-prompter/
├── characters/
│   ├── aragorn_20250113_120000_create.json
│   ├── aragorn_20250113_130000_update.json
│   └── gandalf_20250113_140000_create.json
├── locations/
├── codex/
├── chapters/
└── .metadata/
    ├── config.json
    └── indexes.json
```

**Migration Benefits:**
- Complete version history for all entities
- Atomic file operations prevent corruption
- Easy backup and restore at entity level
- Storage analytics and cleanup capabilities
- Multiple data directory support for project isolation

## Advanced Testing

### Migration Testing
```bash
# Create old format test data
mkdir test_old_format
echo '[{"id":"1","name":"Old Character"}]' > test_old_format/characters.json

# Test migration
result, err := server.ExecuteTool("migrate_from_json", map[string]interface{}{
    "oldDataDir": "test_old_format",
})
```

### Version Management Testing
```bash
# Create entity and multiple versions
character := createCharacter("Test Character")
updateCharacter(character.ID, "Updated description")
updateCharacter(character.ID, "Final description")

# Test version history
versions := getVersions("characters", character.ID)
assert.Equal(t, 3, len(versions)) // create + 2 updates

# Test restore
restore(character.ID, versions[1].Timestamp)
current := getLatest("characters", character.ID)
assert.Equal(t, "Updated description", current.Description)
```

### Performance Testing
- Test with large datasets (1000+ entities with multiple versions)
- Concurrent version creation and restoration
- Storage cleanup with large version histories
- Migration performance with large JSON files

### Integration Testing
- Test with actual Claude Desktop MCP integration
- Validate MCP protocol compliance for version tools
- Test with other MCP clients using version management

## Troubleshooting Storage System

### Enable Debug Output
```go
// Add debug logging for storage operations
log.SetLevel(log.DebugLevel)
```

### Check Storage Structure
```bash
# View folder structure
tree ~/.ai-novel-prompter/

# Check version files
ls -la ~/.ai-novel-prompter/characters/

# Verify metadata
cat ~/.ai-novel-prompter/.metadata/config.json
```

### Verify Version Tools
```bash
# Test version management via API
curl -X POST http://localhost:8080/execute \
  -H "Content-Type: application/json" \
  -d '{"tool": "get_storage_stats", "params": {}}'

curl -X POST http://localhost:8080/execute \
  -H "Content-Type: application/json" \
  -d '{"tool": "get_entity_versions", "params": {"entityType": "characters", "entityId": "char123"}}'
```

### Storage Integrity Check
```bash
# Run storage test program
cd cmd/test-storage
go run main.go

# Check for corruption
go test ./mcp/storage/... -v
```

## Next Steps

After successful testing:
1. **Version Management**: Integrate version history UI in desktop app
2. **Data Directory Management**: Configure multiple project directories
3. **Migration**: Migrate existing installations to new storage format
4. **Backup Strategy**: Implement automated backup of version histories
5. **Performance Optimization**: Monitor and optimize storage operations
6. **Team Collaboration**: Share versioned project directories
7. **Integration**: Add to Claude Desktop MCP configuration with new tools

## Support

If tests fail:
1. Check console output for specific error messages
2. Verify Go version (1.19+ recommended)
3. Ensure all dependencies are downloaded
4. Check file permissions for data directories
5. Verify disk space for version storage
6. Test storage system independently with `cmd/test-storage/main.go`
7. Check migration from old format if applicable
8. Review storage system troubleshooting section

For version management issues:
1. Verify timestamp format and system clock
2. Check entity ID consistency
3. Test version restoration with simple cases
4. Validate storage directory structure
5. Check backup and cleanup operations
