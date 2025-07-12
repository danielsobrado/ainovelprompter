@echo off
echo === AI Novel Prompter MCP Server Test ===
echo.

cd /d "%~dp0"
echo Current directory: %CD%

echo.
echo 1. Checking Go installation...
go version
if %ERRORLEVEL% neq 0 (
    echo ERROR: Go is not installed or not in PATH
    pause
    exit /b 1
)

echo.
echo 2. Downloading dependencies...
go mod download
if %ERRORLEVEL% neq 0 (
    echo ERROR: Failed to download dependencies
    pause
    exit /b 1
)

echo.
echo 3. Testing compilation...
go build -o test_server.exe test_comprehensive.go
if %ERRORLEVEL% neq 0 (
    echo ERROR: Compilation failed
    pause
    exit /b 1
)
echo SUCCESS: Compilation successful

echo.
echo 4. Running comprehensive tests...
echo =====================================
test_server.exe
echo =====================================

echo.
echo 5. Building HTTP server...
go build -o http_server.exe http_server.go
if %ERRORLEVEL% neq 0 (
    echo ERROR: HTTP server compilation failed
    pause
    exit /b 1
)
echo SUCCESS: HTTP server compiled

echo.
echo 6. Testing basic MCP server...
go build -o mcp_server.exe main.go
if %ERRORLEVEL% neq 0 (
    echo ERROR: MCP server compilation failed
    pause
    exit /b 1
)
echo SUCCESS: MCP server compiled

echo.
echo === All Tests Completed ===
echo.
echo You can now run:
echo   mcp_server.exe           - Basic MCP server test
echo   test_server.exe          - Comprehensive functionality test  
echo   http_server.exe          - HTTP API server (then visit http://localhost:8080)
echo.
echo To test via HTTP API:
echo   1. Run: http_server.exe
echo   2. Open: http://localhost:8080/test
echo   3. Or:   http://localhost:8080/tools
echo.

pause
