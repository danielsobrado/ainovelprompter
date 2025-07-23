# AI Novel Prompter

**AI Novel Prompter** is a comprehensive writing tool designed to help novelists create consistent, well-structured prompts for AI writing assistants while managing story elements and improving prose through AI-powered assistance.

## 🚀 Key Features

### 🖥️ Desktop Application (Wails)
- **Story Element Management**: Characters, locations, codex entries, rules, and sample chapters
- **Intelligent Prompt Generation**: Optimized for ChatGPT and Claude with real-time preview
- **Prose Improvement Engine**: Iterative text refinement with customizable AI prompts
- **Local Data Storage**: All data saved locally with versioned storage system

![AI Novel 1](https://github.com/danielsobrado/ainovelprompter/blob/main/images/wails1.png)

### 🔗 MCP Server Integration
- **Claude Desktop Integration**: Direct access to your story data through Claude Desktop
- **Real-time Character/Location Access**: Query and manage entities directly in Claude conversations
- **Configurable Logging**: Debug character loading issues with detailed logging
- **Robust Data Validation**: Prevents creation of incomplete or empty entities

![AI Novel MCP 1](https://github.com/danielsobrado/ainovelprompter/blob/main/images/mcp1.png)

## 📦 Installation & Setup

### Desktop Application

```bash
# Clone the repository
git clone https://github.com/danielsobrado/ainovelprompter.git
cd ainovelprompter

# Install frontend dependencies
cd frontend
npm install
cd ..

# Development mode
wails dev

# Build for production
wails build
```

### MCP Server for Claude Desktop

1. **Build the MCP Server**:
```bash
cd mcp
go build -o ainovelprompter-mcp.exe ..\cmd\mcp-server\main.go
```

2. **Configure Claude Desktop** (edit `claude_desktop_config.json`):
```json
{
  "mcpServers": {
    "ai-novel-prompter": {
      "command": "C:\\path\\to\\ainovelprompter-mcp.exe",
      "args": [
        "--data-dir", "C:\\path\\to\\your\\story\\data",
        "--log-level", "INFO"
      ],
      "env": {}
    }
  }
}
```

3. **Restart Claude Desktop** to load the MCP server

## ⚙️ Configuration

### MCP Server Options

| Argument | Short | Description | Default |
|----------|-------|-------------|---------|
| `--data-dir` | `-d` | Data directory path | `~/.ai-novel-prompter` |
| `--log-level` | `-l` | Logging verbosity | `INFO` |
| `--help` | `-h` | Show help message | - |

### Log Levels
- **DEBUG**: Detailed operation logging (for troubleshooting)
- **INFO**: Standard operational logging (recommended)
- **WARN**: Warnings and unexpected conditions only
- **ERROR**: Errors only

### Example Configurations

**Development/Troubleshooting**:
```json
"args": ["--data-dir", "./test-data", "--log-level", "DEBUG"]
```

**Production**:
```json
"args": ["--data-dir", "/path/to/story/data", "--log-level", "INFO"]
```

## 🛡️ Data Validation & Quality

### Enhanced Validation System
- **Required Fields**: All entities must have non-empty names and descriptions
- **Length Constraints**: Names (1-200 chars), Descriptions (1-2000 chars)
- **Automatic Sanitization**: Whitespace trimming and cleanup
- **Error Prevention**: Cannot create incomplete entities

### Validation Rules
- ✅ Character names and descriptions required
- ✅ Location names and descriptions required  
- ✅ Codex entry titles and content required
- ✅ Automatic whitespace cleanup
- ❌ Empty or whitespace-only entities rejected

## 🎯 Usage

### Desktop Application

1. **Prompt Generation Tab**:
   - Define task types (e.g., "Write Next Chapter", "Revise Chapter")
   - Manage story elements (characters, locations, rules, codex)
   - Generate AI-optimized prompts for ChatGPT or Claude
   - Real-time token counting and preview

2. **Prose Improvement Tab**:
   - Configure LLM providers (Manual, LM Studio, OpenRouter)
   - Apply customizable improvement prompts
   - Review and accept/reject AI suggestions
   - Live diff preview of changes

### Claude Desktop Integration

Once configured, you can directly interact with your story data in Claude conversations:

```
# Get all characters
Show me my characters

# Search for specific characters
Find characters related to "magic"

# Get locations
What locations do I have in my story?

# Search locations
Show me locations in the forest
```

## 🔧 Technical Stack

### Frontend
- **React** with TypeScript
- **Tailwind CSS** for styling
- **shadcn/ui** components
- **Wails** framework for desktop

### Backend
- **Go** with Wails bindings
- **MCP (Model Context Protocol)** server
- **Versioned JSON storage** system
- **Comprehensive validation** layer

### MCP Server
- **JSON-RPC 2.0** protocol
- **Stdio communication** with Claude Desktop
- **Configurable logging** system
- **Real-time data access**

## 📁 Data Storage

### Storage Format
```
~/.ai-novel-prompter/
├── characters/
│   └── [character-id]/
│       ├── 2025-07-22T10-30-15.000+04-00.json
│       └── 2025-07-22T11-45-20.000+04-00.json
├── locations/
├── codex/
├── rules/
├── sample-chapters/
├── task-types/
└── prose-prompts/
```

### Features
- **Version History**: Complete audit trail for all changes
- **Data Integrity**: Individual files reduce corruption risk
- **Scalability**: Better performance with large datasets
- **MCP Compatibility**: Direct access through Claude Desktop

## 🐛 Troubleshooting

### Character Loading Issues

1. **Enable Debug Logging**:
```json
"args": ["--log-level", "DEBUG"]
```

2. **Check Claude Desktop Logs** for detailed error information

3. **Common Issues**:
   - Invalid timestamps in filenames
   - Corrupted JSON files
   - Permission problems
   - Missing data directories

### MCP Server Issues

1. **Verify Server Binary**: Ensure `ainovelprompter-mcp.exe` is built with latest code
2. **Check Configuration**: Verify paths and arguments in `claude_desktop_config.json`
3. **Restart Claude Desktop**: Required after configuration changes
4. **Review Logs**: Look for startup errors and connection issues

## 📊 Recent Improvements

### Version 0.1.0 Features
- ✅ **MCP Server Integration**: Direct Claude Desktop access
- ✅ **Enhanced Logging**: Configurable debug levels with performance tracking
- ✅ **Data Validation**: Prevents incomplete entity creation
- ✅ **Improved Character Loading**: Robust timestamp parsing and error handling
- ✅ **Command Line Interface**: Flexible configuration options

### Validation Enhancements
- ✅ **Frontend Validation**: Immediate feedback with specific error messages
- ✅ **Backend Validation**: Server-side validation with detailed errors
- ✅ **Automatic Sanitization**: Whitespace trimming and cleanup
- ✅ **Length Constraints**: Proper limits for names and descriptions

## 🤝 Contributing

Contributions are welcome! Please feel free to:
- Open issues for bugs or feature requests
- Submit pull requests for improvements
- Share feedback and suggestions

## 📄 License

This project is licensed under the **Attribution-NonCommercial-NoDerivatives (BY-NC-ND)** license.  
See: https://creativecommons.org/licenses/by-nc-nd/4.0/deed.en

---

## 🔗 Links

- **Repository**: https://github.com/danielsobrado/ainovelprompter
- **Wails Framework**: https://wails.io/
- **shadcn/ui Components**: https://ui.shadcn.com/
- **Model Context Protocol**: https://modelcontextprotocol.io/

---

*Last Updated: July 2025 - MCP Integration & Validation Release*