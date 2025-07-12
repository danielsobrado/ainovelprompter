# Claude Code MCP Installation Guide

## Overview

This guide will help you install the AI Novel Prompter MCP server to work with Claude Code, giving you access to 41 powerful novel writing tools directly in your coding environment.

## Features You'll Get

### 🎭 Story Management (15+ tools)
- **Characters**: Create, manage, and search character profiles
- **Locations**: Detailed location management and descriptions  
- **Codex**: World-building entries and lore management
- **Rules**: Writing guidelines and style constraints

### 📚 Chapter Management (12+ tools)
- **Chapters**: Full CRUD operations with auto-numbering
- **Story Beats**: Chapter planning and structure
- **Future Notes**: Plot development tracking
- **Sample Chapters**: Style reference management

### ✍️ Prose Improvement (8+ tools)  
- **Analysis**: Apply improvement prompts to text
- **Sessions**: Manage prose improvement workflows
- **Prompts**: Categorized improvement suggestions
- **Change Tracking**: Review and apply edits

### 🔍 Search & Analysis (4+ tools)
- **Global Search**: Cross-story element searching
- **Text Analysis**: Style and tone extraction
- **Character Tracking**: Mention analysis across chapters
- **Timeline**: Story chronology management

### 🤖 AI Integration (2+ tools)
- **Prompt Generation**: Context-aware AI prompts
- **Templates**: ChatGPT and Claude formatting

## Installation Steps

### Step 1: Build the MCP Server

**Windows:**
```batch
cd mcp
.\build_claude_code.bat
```

**macOS/Linux:**
```bash
cd mcp
chmod +x build_claude_code.sh
./build_claude_code.sh
```

**Manual Build:**
```bash
cd mcp
go mod tidy
go build -o ainovelprompter-mcp mcp_stdio_server.go
```

### Step 2: Configure Claude Code

1. **Find your Claude Code config file:**
   - **Windows**: `%APPDATA%\Claude\claude_desktop_config.json`
   - **macOS**: `~/Library/Application Support/Claude/claude_desktop_config.json`
   - **Linux**: `~/.config/claude/claude_desktop_config.json`

2. **Add the MCP server configuration:**

**Windows Example:**
```json
{
  "mcpServers": {
    "ai-novel-prompter": {
      "command": "C:\\Development\\workspace\\GitHub\\ainovelprompter\\mcp\\ainovelprompter-mcp.exe",
      "args": [],
      "env": {}
    }
  }
}
```

**macOS/Linux Example:**
```json
{
  "mcpServers": {
    "ai-novel-prompter": {
      "command": "/absolute/path/to/your/ainovelprompter/mcp/ainovelprompter-mcp",
      "args": [],
      "env": {}
    }
  }
}
```

3. **Update the path** to match your actual project location

### Step 3: Restart Claude Code

1. Close Claude Code completely
2. Restart Claude Code
3. The MCP server should automatically connect

## Verification

### Test Connection
Open Claude Code and try these commands:

```bash
# List all available tools
claude code --mcp-tools

# Test a simple tool
claude code --mcp-call get_characters

# Create a test character
claude code --mcp-call create_character --params '{"name":"Test Character","description":"A test character for verification"}'
```

### Example Usage in Claude Code

```bash
# Generate a chapter prompt with full story context
"Use the MCP server to generate a ChatGPT prompt for Chapter 5, including characters Alice and Bob, the tavern location, and the 'show don't tell' writing rule"

# Analyze prose quality
"Use the prose improvement tools to analyze this text for style and suggest improvements: 'The man walked quickly down the dark street.'"

# Search across your story
"Search all story content for mentions of 'magic system' and show me what you find"
```

## Available Tools

### Character Management
- `get_characters` - List all characters (supports search)
- `get_character_by_id` - Get specific character details
- `create_character` - Add new characters
- `update_character` - Modify existing characters
- `delete_character` - Remove characters

### Location Management  
- `get_locations` - List all locations
- `create_location` - Add new locations
- `get_location_by_id` - Get location details

### Chapter Management
- `get_chapters` - List chapters (with filtering)
- `get_chapter_content` - Get full chapter text
- `create_chapter` - Add new chapters
- `update_chapter` - Modify chapters
- `get_previous_chapter` - Get preceding chapter

### Prose Tools
- `get_prose_prompts` - List improvement prompts
- `analyze_prose` - Apply prompts to text
- `create_prose_session` - Start improvement workflow
- `get_prose_session` - Check session status

### Search & Analysis
- `search_all_content` - Global search
- `analyze_text_traits` - Extract style information
- `get_character_mentions` - Track character appearances
- `get_timeline_events` - Story chronology

### Prompt Generation
- `generate_chapter_prompt` - Create AI prompts with context
- `get_prompt_template` - Get formatting templates

## Troubleshooting

### Common Issues

1. **Server Not Found**
   - Verify the executable path in config
   - Ensure the file is built and executable
   - Check that Claude Code can access the file

2. **Tools Not Loading**
   - Restart Claude Code completely
   - Check the server logs in Claude Code settings
   - Verify JSON configuration syntax

3. **Data Not Found**
   - The server creates data in `~/.ai-novel-prompter/`
   - Ensure write permissions exist
   - Start with creating test data

### Debug Mode

Enable debug logging by setting environment variable:
```json
{
  "mcpServers": {
    "ai-novel-prompter": {
      "command": "path/to/ainovelprompter-mcp.exe",
      "args": [],
      "env": {
        "MCP_DEBUG": "true"
      }
    }
  }
}
```

### Manual Testing

Test the server directly:
```bash
echo '{"jsonrpc":"2.0","id":1,"method":"tools/list","params":{}}' | ./ainovelprompter-mcp
```

## Data Storage

The MCP server stores data in:
- **Windows**: `%USERPROFILE%\.ai-novel-prompter\`
- **macOS/Linux**: `~/.ai-novel-prompter/`

Files created:
- `characters.json` - Character data
- `locations.json` - Location data  
- `chapters.json` - Chapter content
- `rules.json` - Writing rules
- `codex.json` - World-building entries
- `prose_prompts.json` - Improvement prompts

## Benefits in Claude Code

### Novel Writing Workflow
1. **Context Management**: Keep track of characters, locations, and rules
2. **Chapter Planning**: Organize story beats and future developments  
3. **Prose Improvement**: Systematic text enhancement
4. **AI Integration**: Generate contextual prompts for writing assistance

### Development Integration
- Use story data for documentation
- Generate test content for applications
- Analyze writing patterns programmatically
- Automate story consistency checks

## Advanced Configuration

### Custom Data Directory
```json
{
  "mcpServers": {
    "ai-novel-prompter": {
      "command": "path/to/ainovelprompter-mcp.exe",
      "args": [],
      "env": {
        "AINOVEL_DATA_DIR": "/custom/path/to/data"
      }
    }
  }
}
```

### Multiple Projects
```json
{
  "mcpServers": {
    "novel-project-1": {
      "command": "path/to/ainovelprompter-mcp.exe",
      "env": {"AINOVEL_DATA_DIR": "/path/to/project1"}
    },
    "novel-project-2": {
      "command": "path/to/ainovelprompter-mcp.exe", 
      "env": {"AINOVEL_DATA_DIR": "/path/to/project2"}
    }
  }
}
```

This gives you access to powerful novel writing tools directly in Claude Code for enhanced productivity!

## Support

### Getting Help
- Check the `TESTING.md` file for testing procedures
- Review server logs in Claude Code
- Test individual tools using the HTTP server (`http_server.go`)

### Contributing
- The MCP server is fully extensible
- Add new tools in `handlers/`
- Follow the existing pattern for consistency
- Test thoroughly before deployment