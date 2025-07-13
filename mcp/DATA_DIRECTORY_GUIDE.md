# MCP Server Data Directory Configuration

## Overview

The AI Novel Prompter MCP server can be configured to use different data directories for your story files. This is useful when you want to:

- Use your project directory instead of the default location
- Work with multiple stories in different directories
- Store data on a specific drive or network location

## Command Line Options

### Basic Usage

```bash
# Use default data directory (~/.ai-novel-prompter)
ainovelprompter-mcp.exe

# Use project directory as data directory
ainovelprompter-mcp.exe --data-dir "C:\Development\workspace\GitHub\ainovelprompter"

# Use relative path
ainovelprompter-mcp.exe -d "./my-story-data"

# Use absolute path (Unix/Linux)
ainovelprompter-mcp.exe --data-dir "/home/user/stories/my-novel"

# Show help
ainovelprompter-mcp.exe --help
```

### Command Line Arguments

| Argument | Short Form | Description | Example |
|----------|------------|-------------|---------|
| `--data-dir` | `-d` | Specify data directory path | `--data-dir "./story"` |
| `--help` | `-h` | Show help message | `--help` |

## Claude Desktop Configuration

Update your Claude Desktop MCP configuration to include the data directory:

### Windows Example
```json
{
  "mcpServers": {
    "ainovelprompter": {
      "command": "C:\\Development\\workspace\\GitHub\\ainovelprompter\\mcp\\ainovelprompter-mcp.exe",
      "args": ["--data-dir", "C:\\Development\\workspace\\GitHub\\ainovelprompter"],
      "env": {}
    }
  }
}
```

### Unix/Linux Example  
```json
{
  "mcpServers": {
    "ainovelprompter": {
      "command": "/path/to/ainovelprompter-mcp",
      "args": ["--data-dir", "/home/user/projects/my-novel"],
      "env": {}
    }
  }
}
```

### Multiple Projects
Configure different MCP servers for different stories:

```json
{
  "mcpServers": {
    "fantasy-novel": {
      "command": "/path/to/ainovelprompter-mcp",
      "args": ["--data-dir", "/projects/fantasy-story"],
      "env": {}
    },
    "scifi-novel": {
      "command": "/path/to/ainovelprompter-mcp",
      "args": ["--data-dir", "/projects/scifi-story"],
      "env": {}
    }
  }
}
```

## Interactive Launcher

Use the provided batch script for easy setup:

```bash
# Windows
.\run_server.bat
```

The script provides these options:
1. **Default directory** - Uses `~/.ai-novel-prompter`
2. **Project directory** - Uses current directory  
3. **Custom directory** - Specify any path
4. **Show help** - Display command options

## Data Directory Structure

The MCP server creates these files in your specified data directory:

```
your-data-directory/
├── characters.json          # Character profiles and traits
├── locations.json           # Story locations and descriptions  
├── chapters.json            # Chapter content and metadata
├── codex.json              # World-building and lore entries
├── rules.json              # Writing guidelines and constraints
├── prose_prompts.json      # Improvement prompts and categories
├── story_beats.json        # Chapter story beats
├── future_notes.json       # Future chapter notes
├── sample_chapters.json    # Reference chapters for style
└── task_types.json         # Writing task templates
```

## Troubleshooting

### Common Issues

**Server returns empty data:**
- Check that your data directory contains the JSON files
- Verify the path is correct and accessible
- Ensure the server has read/write permissions

**Path with spaces:**
- Always wrap paths with spaces in quotes
- Example: `--data-dir "C:\My Stories\Novel Data"`

**Relative vs Absolute paths:**
- Relative paths are resolved from the current working directory
- Absolute paths are safer for Claude Desktop configuration
- Use forward slashes or escaped backslashes in JSON configuration

### Verification

Test your configuration:

```bash
# Test server startup with data directory
ainovelprompter-mcp.exe --data-dir "your-path" --help

# Check if data files exist
dir "your-data-directory\*.json"  # Windows
ls "your-data-directory"/*.json   # Unix/Linux
```

## Migration

To move from default to custom data directory:

1. **Stop the MCP server**
2. **Copy data files:**
   ```bash
   # Windows
   copy "%USERPROFILE%\.ai-novel-prompter\*.json" "your-new-directory\"
   
   # Unix/Linux  
   cp ~/.ai-novel-prompter/*.json your-new-directory/
   ```
3. **Update Claude Desktop configuration** with new path
4. **Restart Claude Desktop**

## Examples

### Development Setup
```bash
# Use project directory for development
cd C:\Development\workspace\GitHub\ainovelprompter
.\mcp\ainovelprompter-mcp.exe --data-dir .
```

### Production Setup
```bash
# Use dedicated story directory
ainovelprompter-mcp.exe --data-dir "D:\Stories\MyNovel"
```

### Network Storage
```bash
# Use network drive (Windows)
ainovelprompter-mcp.exe --data-dir "\\server\stories\my-novel"

# Use mounted directory (Unix/Linux)
ainovelprompter-mcp.exe --data-dir "/mnt/stories/my-novel"
```

---

**Next Steps:**
- Build the server: `.\build_claude_code.bat`
- Test configuration: `.\run_server.bat`  
- Update Claude Desktop config with your data directory
- Restart Claude Desktop to apply changes
