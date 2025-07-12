# AI Novel Prompter MCP Server

## Overview

This MCP (Model Context Protocol) server provides comprehensive tools for novel writing, story management, and prose improvement. It exposes the core functionality of the AI Novel Prompter application through a standardized MCP interface.

## Features

### 🎭 Story Context Management (15 tools)
- **Characters**: Create, read, update, delete, and search character profiles
- **Locations**: Manage story locations with detailed descriptions
- **Codex**: World-building entries and lore management
- **Rules**: Writing guidelines and style rules

### 📚 Chapter & Text Management (12 tools)
- **Chapters**: Full chapter CRUD operations with auto-numbering
- **Story Beats**: Chapter planning and story structure
- **Future Notes**: Planned developments and plot points
- **Sample Chapters**: Reference chapters for style consistency

### ✍️ Prose Improvement Tools (8 tools)
- **Prose Prompts**: Categorized improvement prompts (tropes, style, grammar)
- **Analysis**: Apply specific improvement prompts to text
- **Sessions**: Manage prose improvement workflows
- **Change Tracking**: Review and apply suggested improvements

### 🔍 Search & Analysis (4 tools)
- **Global Search**: Search across all story elements
- **Text Analysis**: Extract style, tone, and writing traits
- **Character Mentions**: Track character appearances across chapters
- **Timeline Events**: Story chronology and event extraction

### 🤖 Prompt Generation (2 tools)
- **Chapter Prompts**: Generate AI prompts with full story context
- **Templates**: ChatGPT and Claude-formatted prompt templates

## Installation

### For Claude Code Integration (Recommended)

**Quick Installation:**
```bash
cd mcp
# Windows
.uild_claude_code.bat

# macOS/Linux  
chmod +x build_claude_code.sh
./build_claude_code.sh
```

See **[CLAUDE_CODE_INSTALLATION.md](CLAUDE_CODE_INSTALLATION.md)** for complete setup guide.

### For HTTP API Server
```bash
cd mcp
go mod download
go build -o mcp-server main.go
```

### For Custom Integration
```bash
cd mcp
go mod download
go build -o ainovelprompter-mcp mcp_stdio_server.go
```

## Usage

```go
// Initialize MCP Server
server, err := mcp.NewMCPServer()
if err != nil {
    log.Fatal(err)
}

// Get available tools
tools := server.GetTools()

// Execute a tool
result, err := server.ExecuteTool("get_characters", map[string]interface{}{
    "search": "protagonist",
})
```

## Tool Categories

### Story Context Management
- `get_characters`, `get_character_by_id`, `create_character`, `update_character`, `delete_character`
- `get_locations`, `get_location_by_id`, `create_location`, `search_locations`
- `get_codex_entries`, `get_codex_entry_by_id`, `create_codex_entry`, `search_codex`
- `get_rules`, `get_rule_by_id`, `build_writing_context`

### Chapter Management
- `get_chapters`, `get_chapter_content`, `get_previous_chapter`, `create_chapter`, `update_chapter`, `delete_chapter`
- `get_story_beats`, `save_story_beats`, `get_future_notes`, `create_future_note`
- `get_sample_chapters`, `get_sample_chapter_by_id`, `create_sample_chapter`

### Prose Improvement
- `get_prose_prompts`, `get_prose_prompt_by_id`, `create_prose_prompt`, `analyze_prose`
- `get_prose_prompt_by_category`, `create_prose_session`, `get_prose_session`, `update_prose_session`

### Search & Analysis
- `search_all_content`, `analyze_text_traits`, `get_character_mentions`, `get_timeline_events`

### Prompt Generation
- `generate_chapter_prompt`, `get_prompt_template`

## Data Storage

The MCP server uses file-based JSON storage compatible with the main AI Novel Prompter application:

- **Storage Location**: `~/.ai-novel-prompter/`
- **Format**: JSON files for each data type
- **Thread Safety**: Concurrent access protection with mutexes
- **Auto-backup**: Automatic data persistence

## Configuration

No additional configuration required. The server automatically:
- Creates data directory if it doesn't exist
- Initializes default data structures
- Provides backward compatibility with existing data

## Error Handling

All tools include comprehensive error handling:
- Input validation and type checking
- File operation error recovery
- Graceful degradation for missing data
- Detailed error messages with context

## Performance

- **Concurrent**: Thread-safe operations for multiple clients
- **Efficient**: Optimized JSON parsing and search algorithms
- **Scalable**: Handles large story datasets efficiently
- **Memory**: Low memory footprint with lazy loading

## Development

```bash
# Run tests
go test ./...

# Format code
go fmt ./...

# Vet code
go vet ./...

# Build for production
go build -ldflags="-s -w" -o mcp-server main.go
```

## Integration

This MCP server provides multiple integration options:

### 🎯 Claude Code Integration
- **Full MCP Protocol Support** via `mcp_stdio_server.go`
- **41 Novel Writing Tools** directly in your coding environment
- **Installation Guide**: [CLAUDE_CODE_INSTALLATION.md](CLAUDE_CODE_INSTALLATION.md)
- **Build Scripts**: `build_claude_code.sh` / `build_claude_code.bat`
- **Config Examples**: [claude_code_config_examples.md](claude_code_config_examples.md)

### 🌐 HTTP API Server
- **REST API Interface** via `http_server.go`
- **Web Testing Interface** at http://localhost:8080/test
- **JSON Responses** for easy integration
- **Cross-Platform Access** from any HTTP client

### 🔧 Direct Integration
- **Go Library Import** via `server.go`
- **Custom MCP Clients** 
- **Main AI Novel Prompter App** compatibility
- **Third-Party Tools** supporting MCP protocol

### 📁 Files Overview
- `mcp_stdio_server.go` - Claude Code MCP server (stdio protocol)
- `http_server.go` - HTTP API server with web interface
- `main.go` - Basic MCP server demonstration
- `server.go` - Core MCP server library

## License

Part of the AI Novel Prompter project. See main project LICENSE for details.
