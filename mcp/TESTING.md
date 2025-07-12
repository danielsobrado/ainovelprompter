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

# 2. Start HTTP server for web-based testing
go run http_server.go
# Then visit: http://localhost:8080/test

# 3. Test basic MCP server
go run main.go
```

## Testing Options

### 1. 🚀 Comprehensive Test (`test_comprehensive.go`)
- Tests all 44 MCP tools
- Validates data operations (create, read, update)
- Tests error handling
- Runs automatically

**Usage:**
```bash
go run test_comprehensive.go
```

**Expected Output:**
```
=== AI Novel Prompter MCP Server Test ===
✅ MCP Server initialized successfully!
✅ Found 44 MCP tools:
   • Story Context: 15 tools
   • Chapter Management: 12 tools
   • Prose Improvement: 8 tools
   • Search & Analysis: 4 tools
   • Prompt Generation: 2 tools
[... detailed test results ...]
🎉 All tests completed successfully!
```

### 2. 🌐 HTTP API Server (`http_server.go`)
- Exposes MCP tools via REST API
- Web-based testing interface
- JSON responses for easy integration

**Usage:**
```bash
go run http_server.go
```

**Endpoints:**
- `GET /` - API documentation
- `GET /tools` - List all MCP tools  
- `GET /test` - Run automated tests
- `POST /execute` - Execute MCP tool

**Test URLs:**
- http://localhost:8080/
- http://localhost:8080/tools
- http://localhost:8080/test

### 3. 🔧 Basic MCP Server (`main.go`)
- Simple MCP server demonstration
- Shows tool discovery and basic execution
- Minimal output for debugging

**Usage:**
```bash
go run main.go
```

## PowerShell Testing Scripts

### `test.ps1` - Full Test Suite
- Checks Go installation
- Downloads dependencies
- Compiles all components
- Runs comprehensive tests
- Optionally starts HTTP server

### `test_api.ps1` - HTTP API Testing
- Tests all HTTP endpoints
- Validates API responses
- Tests error handling
- Requires HTTP server to be running

**Usage:**
```powershell
# Terminal 1: Start HTTP server
.\http_server.exe

# Terminal 2: Run API tests
.\test_api.ps1
```

## Expected Results

### ✅ Success Indicators
- All tools compile without errors
- MCP server initializes successfully
- 44 tools are discovered and listed
- Basic operations (create character, search, etc.) work
- Error handling functions correctly
- HTTP endpoints return valid JSON

### ❌ Common Issues & Solutions

**"Go is not installed"**
- Install Go from https://golang.org/dl/
- Add Go to your PATH environment variable

**"Failed to download dependencies"**
- Check internet connection
- Run `go mod tidy` in the mcp directory

**"Compilation failed"**
- Check that all .go files are in the correct directories
- Verify import paths match the module structure

**"Server failed to start"**
- Check if port 8080 is already in use
- Verify file permissions for data directory
- Look for error messages in console output

## Data Storage

During testing, the MCP server will create data files in:
- **Windows**: `C:\Users\[USERNAME]\.ai-novel-prompter\`
- **Files**: `characters.json`, `chapters.json`, `prose_improvement_prompts.json`, etc.

These files are compatible with the main AI Novel Prompter desktop application.

## Advanced Testing

### Manual Tool Testing
```go
// Create server
server, _ := mcp.NewMCPServer()

// Test specific tool
result, err := server.ExecuteTool("get_characters", map[string]interface{}{
    "search": "protagonist",
})
```

### Performance Testing
- Test with large datasets (100+ characters/chapters)
- Concurrent tool execution
- Memory usage monitoring

### Integration Testing
- Test with actual Claude Desktop MCP integration
- Validate MCP protocol compliance
- Test with other MCP clients

## Troubleshooting

### Enable Debug Output
```go
// Add debug logging in main.go
log.SetLevel(log.DebugLevel)
```

### Check Data Files
```bash
# View created data
cat ~/.ai-novel-prompter/characters.json
cat ~/.ai-novel-prompter/chapters.json
```

### Verify Tool Registration
```bash
# List all available tools
curl http://localhost:8080/tools
```

## Next Steps

After successful testing:
1. Integrate with Claude Desktop MCP configuration
2. Add to your writing workflow
3. Extend with custom tools as needed
4. Monitor performance with real data

## Support

If tests fail:
1. Check the console output for specific error messages
2. Verify Go version (1.19+ recommended)
3. Ensure all dependencies are downloaded
4. Check file permissions for data directory
5. Review the troubleshooting section above
