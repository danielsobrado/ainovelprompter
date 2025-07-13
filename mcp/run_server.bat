@echo off
REM AI Novel Prompter MCP Server - Quick Start Script
REM This script helps you run the MCP server with different configurations

setlocal enabledelayedexpansion

echo 🎯 AI Novel Prompter MCP Server Launcher
echo.

REM Check if executable exists
if not exist "ainovelprompter-mcp.exe" (
    echo ❌ MCP server executable not found!
    echo    Run build_claude_code.bat first to build the server.
    echo.
    pause
    exit /b 1
)

echo Choose an option:
echo.
echo 1. Use default data directory (~/.ai-novel-prompter)
echo 2. Use project directory (current directory)
echo 3. Use custom directory
echo 4. Show help
echo.

set /p choice="Enter your choice (1-4): "

if "%choice%"=="1" (
    echo.
    echo 🚀 Starting MCP server with default data directory...
    ainovelprompter-mcp.exe
) else if "%choice%"=="2" (
    echo.
    echo 🚀 Starting MCP server with project directory as data directory...
    ainovelprompter-mcp.exe --data-dir "%CD%"
) else if "%choice%"=="3" (
    echo.
    set /p custom_dir="Enter custom data directory path: "
    echo.
    echo 🚀 Starting MCP server with custom directory: !custom_dir!
    ainovelprompter-mcp.exe --data-dir "!custom_dir!"
) else if "%choice%"=="4" (
    echo.
    ainovelprompter-mcp.exe --help
    echo.
    pause
) else (
    echo.
    echo ❌ Invalid choice. Please run the script again.
    echo.
    pause
)

endlocal
