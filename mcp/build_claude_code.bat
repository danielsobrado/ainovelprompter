@echo off
REM Build AI Novel Prompter MCP Server for Claude Code (Windows)
echo 🔨 Building AI Novel Prompter MCP Server for Claude Code...

REM Check if we're in the right directory
if not exist "mcp_stdio_server.go" (
    echo ❌ Error: mcp_stdio_server.go not found. Please run this script from the mcp\ directory.
    pause
    exit /b 1
)

REM Check if Go is installed
go version >nul 2>&1
if %ERRORLEVEL% NEQ 0 (
    echo ❌ Error: Go is not installed. Please install Go from https://golang.org/dl/
    pause
    exit /b 1
)

echo 📦 Installing dependencies...
go mod tidy

if %ERRORLEVEL% NEQ 0 (
    echo ❌ Error: Failed to download dependencies
    pause
    exit /b 1
)

echo 🔧 Building MCP server executable...
go build -o ainovelprompter-mcp.exe main.go

if %ERRORLEVEL% NEQ 0 (
    echo ❌ Error: Build failed
    pause
    exit /b 1
)

echo ✅ MCP Server built successfully!
echo 📍 Executable location: %CD%\ainovelprompter-mcp.exe

REM Test the server (basic check)
echo 🧪 Testing MCP server...
echo {"jsonrpc":"2.0","id":1,"method":"initialize","params":{"protocolVersion":"2024-11-05","capabilities":{},"clientInfo":{"name":"test","version":"1.0.0"}}} | ainovelprompter-mcp.exe >nul 2>&1

if %ERRORLEVEL% EQU 0 (
    echo ✅ Server test passed!
) else (
    echo ⚠️  Server test failed, but executable was created
)

echo.
echo 🎉 Ready for Claude Code integration!
echo.
echo Next steps:
echo 1. Add this server to your Claude Code configuration:
echo    Path: %CD%\ainovelprompter-mcp.exe
echo 2. Configuration file location:
echo    %%APPDATA%%\Claude\claude_desktop_config.json
echo 3. Restart Claude Code to load the MCP server
echo.
echo See CLAUDE_CODE_INSTALLATION.md for detailed setup instructions.
echo.
pause