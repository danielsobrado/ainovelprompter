@echo off
echo Stopping MCP Server...
taskkill /F /IM ainovelprompter-mcp.exe 2>NUL
timeout /T 2 /NOBREAK >NUL

echo Building MCP Server...
cd "C:\Development\workspace\GitHub\ainovelprompter\mcp"
call build_claude_code.bat

echo MCP Server rebuilt successfully!
pause
