# Test Filename Parsing Fix
Write-Host "=== Testing Filename Parsing Fix ===" -ForegroundColor Cyan

Set-Location "C:\Development\workspace\GitHub\ainovelprompter"

Write-Host ""
Write-Host "=== Running Filename Parsing Test ===" -ForegroundColor Yellow
$filenameTest = go test -v -run TestFilenameParsingFix .\tests\test_filename_fix.go .\app.go
if ($LASTEXITCODE -eq 0) {
    Write-Host "[+] Filename parsing fix PASSED" -ForegroundColor Green
} else {
    Write-Host "[-] Filename parsing fix FAILED" -ForegroundColor Red
    Write-Host $filenameTest -ForegroundColor Yellow
}

Write-Host ""
Write-Host "=== Running Migration Debug Test ===" -ForegroundColor Yellow
$migrationDebugTest = go test -v -run TestMigrationDebug .\tests\test_filename_fix.go .\app.go
if ($LASTEXITCODE -eq 0) {
    Write-Host "[+] Migration debug test PASSED" -ForegroundColor Green
} else {
    Write-Host "[-] Migration debug test FAILED" -ForegroundColor Red
    Write-Host $migrationDebugTest -ForegroundColor Yellow
}

Write-Host ""
Write-Host "=== Test Summary ===" -ForegroundColor Cyan
Write-Host "These tests verify that the core filename parsing fix is working." -ForegroundColor Yellow
Write-Host "If these pass, the main issue should be resolved." -ForegroundColor Yellow
