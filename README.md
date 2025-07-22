# Legacy File Cleanup Tool

## Problem Fixed

This tool fixes the character count mismatch between:
- **Wails UI**: Shows 3 characters (reads only directory format)
- **MCP Server**: Shows more characters (reads both directory + legacy flat files)

## Root Cause

The storage system has two file formats:
1. **New Format (Directory-per-entity)**: `characters/{entity-id}/timestamp.json`
2. **Legacy Format (Flat files)**: `characters/entityname_YYYYMMDD_HHMMSS_operation.json`

During testing and development, both formats accumulated in the storage directory. The Wails UI only reads the new directory format, while the MCP server reads both formats, causing the count mismatch.

## Solution

Remove all legacy flat files, keeping only the directory-format files.

## Usage

### Windows
```bash
# Run the batch script
cleanup.bat
```

### Unix/Linux/Mac
```bash
# Make the script executable
chmod +x cleanup.sh

# Run the shell script
./cleanup.sh
```

### Manual Execution
```bash
# Run the Go program directly
go run cleanup_legacy_files.go
```

## What It Does

1. **Scans** `~/.ai-novel-prompter/` for legacy files
2. **Identifies** legacy files by pattern:
   - Flat JSON files: `characters.json`, `locations.json`, etc.
   - Timestamped files: `name_YYYYMMDD_HHMMSS_operation.json`
3. **Removes** only legacy files, preserving directory structure
4. **Reports** what was cleaned

## Files Cleaned

### Root Directory Legacy Files
- `characters.json`
- `locations.json`
- `codex.json`
- `rules.json`
- `chapters.json`
- `story_beats.json`
- `future_notes.json`
- `sample_chapters.json`
- `task_types.json`
- `prose_prompts.json`
- `prose_prompt_definitions.json`
- `settings.json`
- `llm_provider_settings.json`

### Entity Directory Legacy Files
Pattern: `entityname_YYYYMMDD_HHMMSS_operation.json`

Examples:
- `test_character_20250716_070553_create.json`
- `aragorn_20250113_120000_update.json`
- `gandalf_20250113_140000_create.json`

## Safety

✅ **Safe to run** - Only removes legacy flat files
✅ **Preserves** all directory-format data
✅ **Non-destructive** - Keeps the correct storage format
✅ **Reversible** - Can restore from backups if needed

## Expected Result

After cleanup:
- **Wails UI**: Still shows 3 characters
- **MCP Server**: Now shows 3 characters (same count)
- **Storage**: Clean, consistent directory format only

## Verification

After running cleanup, test both interfaces:
1. Open Wails UI → Check character count
2. Query MCP server → Check character count
3. Counts should now match!

## Troubleshooting

If counts still don't match:
1. Check for other legacy files: `find ~/.ai-novel-prompter -name "*.json" -type f`
2. Restart MCP server to clear cache
3. Refresh Wails UI

## Technical Details

- **Storage Location**: `~/.ai-novel-prompter/` (Unix/Mac) or `%USERPROFILE%\.ai-novel-prompter` (Windows)
- **Preserved Format**: Directory-per-entity with RFC3339 timestamps
- **Removed Format**: Flat files with legacy naming patterns
- **Language**: Go (requires Go installation)
