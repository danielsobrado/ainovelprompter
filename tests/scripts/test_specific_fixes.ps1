# Test Specific Fixes Script
Write-Host "=== Testing Specific Fixes ===" -ForegroundColor Cyan

# Change to project directory
Set-Location "C:\Development\workspace\GitHub\ainovelprompter"

Write-Host ""
Write-Host "=== Test 1: Path Separator Fix ===" -ForegroundColor Yellow
$pathTest = go test -v -run TestDataDirectoryResolution .\tests\app_extended_test.go .\app.go
if ($LASTEXITCODE -eq 0) {
    Write-Host "[+] Path separator fix PASSED" -ForegroundColor Green
} else {
    Write-Host "[-] Path separator fix FAILED" -ForegroundColor Red
    Write-Host $pathTest -ForegroundColor Yellow
}

Write-Host ""
Write-Host "=== Test 2: Basic Storage Operations ===" -ForegroundColor Yellow
$basicTest = go test -v -run TestBasicStorageOperations .\tests\app_quick_test.go .\app.go
if ($LASTEXITCODE -eq 0) {
    Write-Host "[+] Basic storage operations PASSED" -ForegroundColor Green
} else {
    Write-Host "[-] Basic storage operations FAILED" -ForegroundColor Red
    Write-Host $basicTest -ForegroundColor Yellow
}

Write-Host ""
Write-Host "=== Test 3: Legacy Interface ===" -ForegroundColor Yellow
$legacyTest = go test -v -run TestLegacyInterfaceCompatibility .\tests\app_quick_test.go .\app.go
if ($LASTEXITCODE -eq 0) {
    Write-Host "[+] Legacy interface PASSED" -ForegroundColor Green
} else {
    Write-Host "[-] Legacy interface FAILED" -ForegroundColor Red
    Write-Host $legacyTest -ForegroundColor Yellow
}

Write-Host ""
Write-Host "=== Test 4: Migration Fix (with enhanced debugging) ===" -ForegroundColor Yellow
$migrationTest = go test -v -run TestMigrationFunctionality .\tests\app_extended_test.go .\app.go
if ($LASTEXITCODE -eq 0) {
    Write-Host "[+] Migration fix PASSED" -ForegroundColor Green
} else {
    Write-Host "[-] Migration fix FAILED (check debug output)" -ForegroundColor Red
    Write-Host $migrationTest -ForegroundColor Yellow
}

Write-Host ""
Write-Host "=== Test 5: Simple Storage Test ===" -ForegroundColor Yellow
Set-Location "mcp"
$storageTest = go test -v -run TestFolderStorage_CharacterOperations .\tests\folder_storage_test.go .\storage\*.go
if ($LASTEXITCODE -eq 0) {
    Write-Host "[+] Simple storage test PASSED" -ForegroundColor Green
} else {
    Write-Host "[-] Simple storage test FAILED" -ForegroundColor Red
    Write-Host $storageTest -ForegroundColor Yellow
}
Set-Location ".."

Write-Host ""
Write-Host "=== Summary ===" -ForegroundColor Cyan
Write-Host "Check individual test results above." -ForegroundColor Yellow
Write-Host "Focus on migration test debug output to understand remaining issues." -ForegroundColor Yellow
