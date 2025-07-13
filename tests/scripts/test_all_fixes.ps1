# AI Novel Prompter Test Verification Script
Write-Host "=== AI Novel Prompter Test Verification ===" -ForegroundColor Cyan
Write-Host "Testing fixes for remaining issues..." -ForegroundColor Yellow

# Change to project directory
Set-Location "C:\Development\workspace\GitHub\ainovelprompter"

Write-Host ""
Write-Host "=== Build Test ===" -ForegroundColor Cyan
Write-Host "Building project..."
$buildResult = go build -v
if ($LASTEXITCODE -eq 0) {
    Write-Host "[+] Build successful" -ForegroundColor Green
} else {
    Write-Host "[-] Build failed" -ForegroundColor Red
    Write-Host "Build output: $buildResult" -ForegroundColor Red
    exit 1
}

Write-Host ""
Write-Host "=== Quick Tests ===" -ForegroundColor Cyan
Write-Host "Running quick tests..."
$quickTestResult = go test -v .\tests\app_quick_test.go .\app.go
if ($LASTEXITCODE -eq 0) {
    Write-Host "[+] Quick tests passed" -ForegroundColor Green
} else {
    Write-Host "[-] Quick tests failed" -ForegroundColor Red
    Write-Host "Quick test output:" -ForegroundColor Yellow
    Write-Host $quickTestResult -ForegroundColor Yellow
}

Write-Host ""
Write-Host "=== Extended Tests ===" -ForegroundColor Cyan
Write-Host "Running extended tests..."
$extendedTestResult = go test -v .\tests\app_extended_test.go .\app.go
if ($LASTEXITCODE -eq 0) {
    Write-Host "[+] Extended tests passed" -ForegroundColor Green
} else {
    Write-Host "[-] Extended tests failed" -ForegroundColor Red
    Write-Host "Extended test output:" -ForegroundColor Yellow
    Write-Host $extendedTestResult -ForegroundColor Yellow
}

Write-Host ""
Write-Host "=== Storage Tests ===" -ForegroundColor Cyan
Write-Host "Running storage tests..."
Set-Location "mcp"
$storageTestResult = go test -v .\tests\...
if ($LASTEXITCODE -eq 0) {
    Write-Host "[+] Storage tests passed" -ForegroundColor Green
} else {
    Write-Host "[-] Storage tests failed" -ForegroundColor Red
    Write-Host "Storage test output:" -ForegroundColor Yellow
    Write-Host $storageTestResult -ForegroundColor Yellow
}
Set-Location ".."

Write-Host ""
Write-Host "=== Individual Test Commands ===" -ForegroundColor Cyan
Write-Host "To run specific tests manually:" -ForegroundColor Yellow
Write-Host ""
Write-Host "# Test path separator fix:" -ForegroundColor White
Write-Host "go test -v -run TestDataDirectoryResolution .\tests\app_extended_test.go .\app.go" -ForegroundColor Gray
Write-Host ""
Write-Host "# Test migration fix:" -ForegroundColor White  
Write-Host "go test -v -run TestMigrationFunctionality .\tests\app_extended_test.go .\app.go" -ForegroundColor Gray
Write-Host ""
Write-Host "# Test basic storage operations:" -ForegroundColor White
Write-Host "go test -v -run TestBasicStorageOperations .\tests\app_quick_test.go .\app.go" -ForegroundColor Gray
Write-Host ""
Write-Host "# Test all storage:" -ForegroundColor White
Write-Host "cd mcp && go test -v .\tests\..." -ForegroundColor Gray

Write-Host ""
Write-Host "=== Test Summary ===" -ForegroundColor Cyan
Write-Host "All tests completed. Check output above for specific failures." -ForegroundColor Yellow
Write-Host ""
Write-Host "Key fixes applied:" -ForegroundColor White
Write-Host "[+] Path separator normalization (Windows/Unix compatibility)" -ForegroundColor Green
Write-Host "[+] Migration cache rebuild for character retrieval" -ForegroundColor Green  
Write-Host "[+] SetDataDirectory error handling (already working)" -ForegroundColor Green
Write-Host "[+] Storage type consistency (already working)" -ForegroundColor Green
Write-Host "[+] Critical deadlock fix maintained" -ForegroundColor Green
