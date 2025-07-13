# Comprehensive Fix Testing
Write-Host "=== Comprehensive Fix Testing ===" -ForegroundColor Cyan

Set-Location "C:\Development\workspace\GitHub\ainovelprompter"

Write-Host ""
Write-Host "=== Test 1: Directory Switching Debug ===" -ForegroundColor Yellow
$switchTest = go test -v -run TestSetDataDirectoryDebug .\tests\test_directory_switching.go .\app.go
if ($LASTEXITCODE -eq 0) {
    Write-Host "[+] Directory switching PASSED" -ForegroundColor Green
} else {
    Write-Host "[-] Directory switching FAILED" -ForegroundColor Red
}

Write-Host ""
Write-Host "=== Test 2: Basic Storage Operations ===" -ForegroundColor Yellow  
$basicTest = go test -v -run TestBasicStorageOperations .\tests\app_quick_test.go .\app.go
if ($LASTEXITCODE -eq 0) {
    Write-Host "[+] Basic storage PASSED" -ForegroundColor Green
} else {
    Write-Host "[-] Basic storage FAILED" -ForegroundColor Red
}

Write-Host ""
Write-Host "=== Test 3: Migration Functionality ===" -ForegroundColor Yellow
$migrationTest = go test -v -run TestMigrationFunctionality .\tests\app_extended_test.go .\app.go
if ($LASTEXITCODE -eq 0) {
    Write-Host "[+] Migration PASSED" -ForegroundColor Green
} else {
    Write-Host "[-] Migration FAILED" -ForegroundColor Red
}

Write-Host ""
Write-Host "=== Test 4: Storage Character Operations ===" -ForegroundColor Yellow
Set-Location "mcp"
$charTest = go test -v -run TestFolderStorage_CharacterOperations .\tests\folder_storage_test.go .\storage\*.go
if ($LASTEXITCODE -eq 0) {
    Write-Host "[+] Storage character operations PASSED" -ForegroundColor Green
} else {
    Write-Host "[-] Storage character operations FAILED" -ForegroundColor Red
}
Set-Location ".."

Write-Host ""
Write-Host "=== Test 5: Legacy Interface ===" -ForegroundColor Yellow
Set-Location "mcp" 
$legacyTest = go test -v -run TestFolderStorage_LegacyInterface .\tests\folder_storage_test.go .\storage\*.go
if ($LASTEXITCODE -eq 0) {
    Write-Host "[+] Legacy interface PASSED" -ForegroundColor Green
} else {
    Write-Host "[-] Legacy interface FAILED" -ForegroundColor Red
}
Set-Location ".."

Write-Host ""
Write-Host "=== Summary ===" -ForegroundColor Cyan
Write-Host "Key fixes applied:" -ForegroundColor White
Write-Host "[+] Filename parsing fix for cache rebuild" -ForegroundColor Green
Write-Host "[+] fixActiveStatus method corrected" -ForegroundColor Green  
Write-Host "[+] Path separator normalization" -ForegroundColor Green
Write-Host "[+] Migration debugging enhanced" -ForegroundColor Green
Write-Host "[+] Critical deadlock fix maintained" -ForegroundColor Green
Write-Host ""
Write-Host "If Directory Switching test fails, run it individually for detailed debug output:" -ForegroundColor Yellow
Write-Host ".\tests\scripts\test_directory_debug.ps1" -ForegroundColor Gray
