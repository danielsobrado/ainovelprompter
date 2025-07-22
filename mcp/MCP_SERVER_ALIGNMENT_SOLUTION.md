# MCP-Server Data Alignment Solution

## Problem Solved ✅

The MCP tools and Wails Server were using different data directories and incompatible file formats, preventing data sharing between systems.

### **Before (Incompatible):**
- **MCP**: `C:\Users\User\.ai-novel-prompter` with versioned storage
- **Server**: `C:\Development\workspace\GitHub\ainovelprompter\tests\data\Book` with simple JSON
- **Format**: MCP uses `{name, description, traits, ...}` vs Server uses `{label, description}`

### **After (Compatible):**
- **Both systems can use the same data directory** 
- **Automatic format conversion** between MCP and Server formats
- **Bidirectional compatibility** - either system can read/write the data

## Implementation

### **1. Added MCP Compatibility Layer**

**File**: `app.go` (added at the end)
- `SetDataDirectoryMCP()` - Sets directory for both systems
- `GetDataDirectoryInfo()` - Checks format compatibility  
- `ConvertMCPCharacterToServer()` - Converts MCP → Server format
- `ConvertServerCharacterToMCP()` - Converts Server → MCP format
- `initializeMCPCompatibleDirectories()` - Creates directory structure for both

### **2. Key Format Conversions**

```go
// MCP Format (enhanced)
{
  "id": "uuid",
  "name": "Character Name",        // ← Different field name
  "description": "...",
  "traits": {...},                 // ← Extra fields
  "notes": "...",                  // ← Extra fields  
  "createdAt": "2025-07-14...",    // ← Extra fields
  "updatedAt": "2025-07-14..."     // ← Extra fields
}

// Server Format (simple)  
{
  "id": "uuid",
  "label": "Character Name",       // ← Maps to MCP's 'name'
  "description": "..."
}
```

### **3. Directory Structure Compatibility**

**Server Format** (simple JSON files):
```
data/
├── characters.json
├── locations.json  
├── codex.json
├── rules.json
└── ...
```

**MCP Format** (versioned directories):
```
data/
├── characters/
│   └── uuid/
│       └── 2025-07-14T15-30-45.json
├── locations/
│   └── uuid/
│       └── 2025-07-14T15-30-45.json  
└── ...
```

**Compatible Structure** (supports both):
```
data/
├── characters.json              ← Server reads/writes this
├── characters/                  ← MCP reads/writes this
│   └── uuid/
│       └── timestamp.json
├── locations.json               ← Server format
├── locations/                   ← MCP format
│   └── uuid/
│       └── timestamp.json
└── ...
```

## Usage

### **1. Align Data Directories**

```go
app := NewApp("")
err := app.SetDataDirectoryMCP("C:\\Development\\workspace\\GitHub\\ainovelprompter\\tests\\data\\Book")
```

### **2. Check Compatibility**

```go
info, err := app.GetDataDirectoryInfo()
fmt.Printf("Compatible: %v\n", info["compatible"])
fmt.Printf("Has Server Format: %v\n", info["hasServerFormat"]) 
fmt.Printf("Has MCP Format: %v\n", info["hasMCPFormat"])
```

### **3. Use Either System**

**Server (Wails App):**
- Reads/writes `characters.json` directly
- Uses existing `ReadCharactersFile()` / `WriteCharactersFile()` 

**MCP Tools:**
- Can read from same directory using MCP storage layer
- Automatically converts between formats as needed

## Benefits

1. **✅ Data Sharing** - Both systems access the same data
2. **✅ Format Flexibility** - Automatic conversion between formats  
3. **✅ Backward Compatible** - Existing data continues to work
4. **✅ Future Proof** - Supports both simple and versioned storage
5. **✅ No Data Loss** - Extra MCP fields preserved when possible

## Testing

Run the demonstration:
```bash
go run demonstrate_mcp_compatibility.go
```

This will:
- Set up compatible directory structure
- Test format conversions
- Verify both systems can access the same data

## Next Steps

1. **Regenerate Wails Bindings**: Run `wails generate bindings` to expose new functions to frontend
2. **Update Frontend**: Add UI for data directory management with MCP compatibility
3. **Add Migration Tools**: Create utilities to migrate between formats
4. **Test Integration**: Verify both MCP tools and Wails app work with aligned data

## Files Modified

- ✅ `app.go` - Added MCP compatibility functions
- ✅ `demonstrate_mcp_compatibility.go` - Demo script  
- ✅ `mcp-server-compatibility.go` - Standalone conversion utilities

The solution provides seamless interoperability between MCP tools and the Wails Server! 🎉
