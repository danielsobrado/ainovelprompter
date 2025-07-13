# Quick Test Runner - Fixed for Windows Go Path
Write-Host "=== AI Novel Prompter - Quick Test Runner ===" -ForegroundColor Cyan

# Change to project root
Set-Location "C:\Development\workspace\GitHub\ainovelprompter"

# Go executable path
$GoExe = "C:\Program Files\Go\bin\go.exe"

Write-Host ""
Write-Host "=== Build Verification ===" -ForegroundColor Yellow
$build = & $GoExe build -v
if ($LASTEXITCODE -eq 0) {
    Write-Host "[+] Build successful" -ForegroundColor Green
} else {
    Write-Host "[-] Build failed" -ForegroundColor Red
    exit 1
}

Write-Host ""
Write-Host "=== Core Application Tests ===" -ForegroundColor Yellow
$quickTests = & $GoExe test -v .\app_quick_test.go .\app.go
if ($LASTEXITCODE -eq 0) {
    Write-Host "[+] Quick tests PASSED" -ForegroundColor Green
} else {
    Write-Host "[-] Quick tests FAILED" -ForegroundColor Red
}

Write-Host ""
Write-Host "=== Basic Storage Operations ===" -ForegroundColor Yellow  
$basicTest = & $GoExe test -v -run TestBasicStorageOperations .\app_quick_test.go .\app.go
if ($LASTEXITCODE -eq 0) {
    Write-Host "[+] Basic storage PASSED" -ForegroundColor Green
} else {
    Write-Host "[-] Basic storage FAILED" -ForegroundColor Red
}

Write-Host ""
Write-Host "=== MCP Storage Tests ===" -ForegroundColor Yellow
Set-Location "mcp"
$mcpTests = & $GoExe test -v .\tests\folder_storage_test.go .\storage\*.go
if ($LASTEXITCODE -eq 0) {
    Write-Host "[+] MCP storage tests PASSED" -ForegroundColor Green
} else {
    Write-Host "[-] MCP storage tests FAILED" -ForegroundColor Red
}
Set-Location ".."

Write-Host ""
Write-Host "=== Summary ===" -ForegroundColor Cyan
if ($LASTEXITCODE -eq 0) {
    Write-Host "[+] Core functionality verified!" -ForegroundColor Green
    Write-Host "Your fixes are working correctly." -ForegroundColor Green
} else {
    Write-Host "[!] Some issues remain - check individual test outputs above." -ForegroundColor Yellow
}

Write-Host ""
Write-Host "Available test scripts:" -ForegroundColor White
Write-Host "- .\tests\scripts\test_all_fixes.ps1 - Full test suite" -ForegroundColor Gray
Write-Host "- .\tests\scripts\test_specific_fixes.ps1 - Individual fix tests" -ForegroundColor Gray
Write-Host "- .\tests\scripts\test_directory_debug.ps1 - Directory switching debug" -ForegroundColor Gray

Write-Host ""
Write-Host "=== Test Structure Information ===" -ForegroundColor Cyan
Write-Host "[DIR] tests/                 - Root level app tests" -ForegroundColor Gray
Write-Host "[DIR] tests/scripts/         - Test automation scripts" -ForegroundColor Gray  
Write-Host "[DIR] mcp/tests/             - MCP storage subsystem tests" -ForegroundColor Gray
Write-Host "[FILE] tests/README.md        - Comprehensive test documentation" -ForegroundColor Gray
Write-Host "For detailed test documentation, see: tests/README.md" -ForegroundColor Gray
