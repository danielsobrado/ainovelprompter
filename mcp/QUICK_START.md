# AI Novel Prompter MCP - Quick Start Guide

## 🚀 Choose Your Integration Method

### 1. **Claude Code Integration** (Recommended for Developers)
Get 41 novel writing tools directly in Claude Code for enhanced AI-assisted writing.

**Quick Setup:**
```bash
cd mcp
# Windows
.\build_claude_code.bat
# macOS/Linux
./build_claude_code.sh
```

**Complete Guide:** [CLAUDE_CODE_INSTALLATION.md](CLAUDE_CODE_INSTALLATION.md)  
**Configuration Examples:** [claude_code_config_examples.md](claude_code_config_examples.md)

---

### 2. **HTTP API Server** (For Web Integration)
Access MCP tools via REST API with web testing interface.

**Quick Setup:**
```bash
cd mcp
go run http_server.go
# Visit: http://localhost:8080/test
```

**Endpoints:**
- `GET /tools` - List all available tools
- `POST /execute` - Execute any MCP tool
- `GET /test` - Automated testing interface

---

### 3. **Go Library** (For Custom Applications)
Import MCP server directly in your Go applications.

**Quick Setup:**
```go
import "github.com/danielsobrado/ainovelprompter/mcp"

server, err := mcp.NewMCPServer()
result, err := server.ExecuteTool("get_characters", params)
```

---

## 🎯 What You Get

### **Story Management** (15 tools)
- Characters, locations, codex entries, writing rules
- Search, create, update, delete operations
- Build comprehensive writing contexts

### **Chapter Management** (12 tools)  
- Full chapter CRUD with auto-numbering
- Story beats, future notes, sample chapters
- Timeline and chronology tracking

### **Prose Improvement** (8 tools)
- Categorized improvement prompts
- Prose analysis sessions with change tracking
- Style and grammar enhancement workflows

### **Search & Analysis** (4 tools)
- Global content search across all story elements
- Text trait analysis and style extraction
- Character mention tracking and timeline events

### **AI Integration** (2 tools)
- Context-aware prompt generation for AI writing
- ChatGPT and Claude formatted templates

---

## 🧪 Testing Your Setup

### Method 1: Comprehensive Test
```bash
cd mcp
go run test_comprehensive.go
```

### Method 2: HTTP Testing
```bash
# Terminal 1: Start server
go run http_server.go

# Terminal 2: Test endpoints
curl http://localhost:8080/test
```

### Method 3: PowerShell Testing (Windows)
```powershell
cd mcp
.\test.ps1
```

---

## 📊 Tool Categories

| Category | Tools | Description |
|----------|-------|-------------|
| **Story Context** | 15 | Characters, locations, codex, rules |
| **Chapters** | 12 | Chapter management, beats, notes |
| **Prose** | 8 | Analysis, improvement, sessions |
| **Search** | 4 | Global search, text analysis |
| **Prompts** | 2 | AI prompt generation, templates |
| **Total** | **41** | Complete novel writing toolkit |

---

## 🗂️ Data Storage

MCP server creates data in:
- **Windows**: `%USERPROFILE%\.ai-novel-prompter\`
- **macOS/Linux**: `~/.ai-novel-prompter/`

**JSON Files Created:**
- `characters.json` - Character profiles and traits
- `locations.json` - Story locations and descriptions
- `chapters.json` - Chapter content and metadata
- `codex.json` - World-building and lore entries
- `rules.json` - Writing guidelines and constraints
- `prose_prompts.json` - Improvement prompts and categories

---

## 🔧 Configuration Options

### Environment Variables
| Variable | Purpose | Default |
|----------|---------|---------|
| `AINOVEL_DATA_DIR` | Custom data directory | `~/.ai-novel-prompter/` |
| `MCP_DEBUG` | Enable debug logging | `false` |
| `LOG_LEVEL` | Set logging verbosity | `info` |

### Multiple Projects
Configure separate data directories for different novels:
```json
{
  "mcpServers": {
    "fantasy-novel": {
      "command": "path/to/ainovelprompter-mcp",
      "env": {"AINOVEL_DATA_DIR": "/path/to/fantasy-project"}
    },
    "scifi-novel": {
      "command": "path/to/ainovelprompter-mcp", 
      "env": {"AINOVEL_DATA_DIR": "/path/to/scifi-project"}
    }
  }
}
```

---

## 📚 Documentation Files

| File | Purpose |
|------|---------|
| `README.md` | Main documentation and overview |
| `CLAUDE_CODE_INSTALLATION.md` | Complete Claude Code setup guide |
| `claude_code_config_examples.md` | Configuration examples and troubleshooting |
| `TESTING.md` | Comprehensive testing procedures |
| `QUICK_START.md` | This file - quick setup guide |

---

## 🆘 Need Help?

### Common Issues
1. **Build Failures**: Check Go installation and run `go mod tidy`
2. **Claude Code Not Connecting**: Verify executable path and JSON syntax
3. **Tools Not Working**: Check data directory permissions
4. **Performance Issues**: Monitor data file sizes and enable debug logging

### Testing Commands
```bash
# Test MCP server directly
echo '{"jsonrpc":"2.0","id":1,"method":"tools/list","params":{}}' | ./ainovelprompter-mcp

# Test HTTP API
curl http://localhost:8080/tools

# Run comprehensive tests
go run test_comprehensive.go
```

### Getting Support
- Review `TESTING.md` for detailed testing procedures
- Check Claude Code logs for MCP connection issues  
- Use HTTP server for debugging individual tools
- Verify data directory structure and permissions

---

**🎉 Ready to enhance your novel writing with AI-powered tools!**