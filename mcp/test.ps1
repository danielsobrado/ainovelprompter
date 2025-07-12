# AI Novel Prompter MCP Server Test Script
Write-Host "=== AI Novel Prompter MCP Server Test ===" -ForegroundColor Cyan
Write-Host ""

# Change to script directory
Set-Location $PSScriptRoot
Write-Host "Current directory: $(Get-Location)" -ForegroundColor Gray

# Test Go installation
Write-Host ""
Write-Host "1. Checking Go installation..." -ForegroundColor Yellow
try {
    $goVersion = go version
    Write-Host "✅ $goVersion" -ForegroundColor Green
} catch {
    Write-Host "❌ ERROR: Go is not installed or not in PATH" -ForegroundColor Red
    Read-Host "Press Enter to exit"
    exit 1
}

# Download dependencies
Write-Host ""
Write-Host "2. Downloading dependencies..." -ForegroundColor Yellow
try {
    go mod download
    Write-Host "✅ Dependencies downloaded successfully" -ForegroundColor Green
} catch {
    Write-Host "❌ ERROR: Failed to download dependencies" -ForegroundColor Red
    Write-Host $_.Exception.Message -ForegroundColor Red
    Read-Host "Press Enter to exit"
    exit 1
}

# Test compilation
Write-Host ""
Write-Host "3. Testing compilation..." -ForegroundColor Yellow
try {
    go build -o test_server.exe test_comprehensive.go
    Write-Host "✅ Compilation successful" -ForegroundColor Green
} catch {
    Write-Host "❌ ERROR: Compilation failed" -ForegroundColor Red
    Write-Host $_.Exception.Message -ForegroundColor Red
    Read-Host "Press Enter to exit"
    exit 1
}

# Run comprehensive tests
Write-Host ""
Write-Host "4. Running comprehensive tests..." -ForegroundColor Yellow
Write-Host "=====================================" -ForegroundColor Gray
try {
    & ".\test_server.exe"
    Write-Host "=====================================" -ForegroundColor Gray
    Write-Host "✅ Tests completed" -ForegroundColor Green
} catch {
    Write-Host "❌ ERROR: Tests failed" -ForegroundColor Red
    Write-Host $_.Exception.Message -ForegroundColor Red
}

# Build HTTP server
Write-Host ""
Write-Host "5. Building HTTP server..." -ForegroundColor Yellow
try {
    go build -o http_server.exe http_server.go
    Write-Host "✅ HTTP server compiled successfully" -ForegroundColor Green
} catch {
    Write-Host "❌ ERROR: HTTP server compilation failed" -ForegroundColor Red
    Write-Host $_.Exception.Message -ForegroundColor Red
}

# Build main MCP server
Write-Host ""
Write-Host "6. Building MCP server..." -ForegroundColor Yellow
try {
    go build -o mcp_server.exe main.go
    Write-Host "✅ MCP server compiled successfully" -ForegroundColor Green
} catch {
    Write-Host "❌ ERROR: MCP server compilation failed" -ForegroundColor Red
    Write-Host $_.Exception.Message -ForegroundColor Red
}

# Summary
Write-Host ""
Write-Host "=== All Tests Completed ===" -ForegroundColor Cyan
Write-Host ""
Write-Host "You can now run:" -ForegroundColor White
Write-Host "  .\mcp_server.exe          - Basic MCP server test" -ForegroundColor Gray
Write-Host "  .\test_server.exe         - Comprehensive functionality test" -ForegroundColor Gray  
Write-Host "  .\http_server.exe         - HTTP API server" -ForegroundColor Gray
Write-Host ""
Write-Host "To test via HTTP API:" -ForegroundColor White
Write-Host "  1. Run: .\http_server.exe" -ForegroundColor Gray
Write-Host "  2. Open: http://localhost:8080/test" -ForegroundColor Gray
Write-Host "  3. Or:   http://localhost:8080/tools" -ForegroundColor Gray
Write-Host ""

# Offer to start HTTP server
$response = Read-Host "Would you like to start the HTTP server now? (y/n)"
if ($response -eq 'y' -or $response -eq 'Y') {
    Write-Host ""
    Write-Host "Starting HTTP server..." -ForegroundColor Yellow
    Write-Host "Visit http://localhost:8080 for API documentation" -ForegroundColor Green
    Write-Host "Visit http://localhost:8080/test to run tests" -ForegroundColor Green
    Write-Host "Press Ctrl+C to stop the server" -ForegroundColor Yellow
    Write-Host ""
    
    try {
        & ".\http_server.exe"
    } catch {
        Write-Host "❌ ERROR: Failed to start HTTP server" -ForegroundColor Red
        Write-Host $_.Exception.Message -ForegroundColor Red
    }
} else {
    Write-Host "HTTP server not started. You can run .\http_server.exe manually." -ForegroundColor Gray
}

Read-Host "Press Enter to exit"
