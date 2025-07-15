# Claude Code Configuration Examples

## Basic Configuration

### Windows
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

### macOS
```json
{
  "mcpServers": {
    "ai-novel-prompter": {
      "command": "/Users/username/Projects/ainovelprompter/mcp/ainovelprompter-mcp",
      "args": [],
      "env": {}
    }
  }
}
```

### Linux
```json
{
  "mcpServers": {
    "ai-novel-prompter": {
      "command": "/home/username/projects/ainovelprompter/mcp/ainovelprompter-mcp",
      "args": [],
      "env": {}
    }
  }
}
```

## Sharing Data Between Desktop App and MCP Server

Both the desktop application and MCP server can share the same data directory:

**Desktop App:**
```bash
# Start desktop app with specific data directory
ainovelprompter.exe --data-dir "C:\My Stories\Current Novel"
```

**MCP Server Configuration:**
```json
{
  "mcpServers": {
    "ai-novel-prompter": {
      "command": "path/to/ainovelprompter-mcp.exe",
      "args": ["--data-dir", "C:\\My Stories\\Current Novel"]
    }
  }
}
```

This allows you to:
- Edit your story in the desktop app
- Use Claude Code for AI-assisted writing
- Both tools access the same characters, chapters, and story data
- Changes are immediately available in both applications

## Advanced Configurations

### With Custom Data Directory
```json
{
  "mcpServers": {
    "ai-novel-prompter": {
      "command": "path/to/ainovelprompter-mcp.exe",
      "args": [],
      "env": {
        "AINOVEL_DATA_DIR": "C:\\MyNovels\\ProjectData"
      }
    }
  }
}
```

### With Debug Logging
```json
{
  "mcpServers": {
    "ai-novel-prompter": {
      "command": "path/to/ainovelprompter-mcp.exe",
      "args": [],
      "env": {
        "MCP_DEBUG": "true",
        "LOG_LEVEL": "debug"
      }
    }
  }
}
```

### Multiple Novel Projects

**Option 1: Using Environment Variables**
```json
{
  "mcpServers": {
    "fantasy-novel": {
      "command": "path/to/ainovelprompter-mcp.exe",
      "args": [],
      "env": {
        "AINOVEL_DATA_DIR": "C:\Novels\FantasyProject"
      }
    },
    "scifi-novel": {
      "command": "path/to/ainovelprompter-mcp.exe",
      "args": [],
      "env": {
        "AINOVEL_DATA_DIR": "C:\Novels\SciFiProject"
      }
    }
  }
}
```

**Option 2: Using Command Line Arguments**
```json
{
  "mcpServers": {
    "fantasy-novel": {
      "command": "path/to/ainovelprompter-mcp.exe",
      "args": ["--data-dir", "C:\Novels\FantasyProject"]
    },
    "scifi-novel": {
      "command": "path/to/ainovelprompter-mcp.exe",
      "args": ["-d", "C:\Novels\SciFiProject"]
    }
  }
}
```
```

### With Timeout Configuration
```json
{
  "mcpServers": {
    "ai-novel-prompter": {
      "command": "path/to/ainovelprompter-mcp.exe",
      "args": [],
      "env": {
        "MCP_TIMEOUT": "30000",
        "MAX_RESULTS": "100"
      }
    }
  }
}
```

## Configuration File Locations

- **Windows**: `%APPDATA%\Claude\claude_desktop_config.json`
- **macOS**: `~/Library/Application Support/Claude/claude_desktop_config.json`
- **Linux**: `~/.config/claude/claude_desktop_config.json`

## Environment Variables

| Variable | Description | Default |
|----------|-------------|---------|
| `AINOVEL_DATA_DIR` | Custom data directory | `~/.ai-novel-prompter/` |
| `MCP_DEBUG` | Enable debug logging | `false` |
| `LOG_LEVEL` | Logging level | `info` |
| `MCP_TIMEOUT` | Timeout in milliseconds | `10000` |
| `MAX_RESULTS` | Maximum search results | `50` |

## Troubleshooting Configuration

### Test Your Configuration
1. Save your configuration file
2. Restart Claude Code
3. Open Claude Code terminal
4. Run: `claude code --mcp-servers` (if available)

### Common Issues

**Invalid JSON syntax:**
- Use a JSON validator to check your configuration
- Ensure all strings are quoted
- Check for trailing commas

**Path not found:**
- Use absolute paths instead of relative paths
- On Windows, use `\\` or `/` for path separators
- Verify the executable exists and is accessible

**Permission denied:**
- Ensure the executable has proper permissions
- On Unix systems: `chmod +x ainovelprompter-mcp`
- Check firewall settings if applicable

### Validation Commands

**Check if executable works:**
```bash
# Test direct execution
echo '{"jsonrpc":"2.0","id":1,"method":"initialize","params":{"protocolVersion":"2024-11-05","capabilities":{},"clientInfo":{"name":"test","version":"1.0.0"}}}' | ./ainovelprompter-mcp
```

**Validate JSON configuration:**
```bash
# On Unix systems
cat ~/.config/claude/claude_desktop_config.json | jq .

# On Windows (if jq is installed)
type %APPDATA%\Claude\claude_desktop_config.json | jq .
```