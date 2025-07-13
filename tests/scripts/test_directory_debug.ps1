# Test Directory Switching Debug
Write-Host "=== Testing Directory Switching Debug ===" -ForegroundColor Cyan

Set-Location "C:\Development\workspace\GitHub\ainovelprompter"

Write-Host ""
Write-Host "=== Running Directory Switching Debug Test ===" -ForegroundColor Yellow
$switchTest = go test -v -run TestSetDataDirectoryDebug .\tests\test_directory_switching.go .\app.go
if ($LASTEXITCODE -eq 0) {
    Write-Host "[+] Directory switching test PASSED" -ForegroundColor Green
} else {
    Write-Host "[-] Directory switching test FAILED" -ForegroundColor Red
    Write-Host "Debug output:" -ForegroundColor Yellow
    Write-Host $switchTest -ForegroundColor White
}

Write-Host ""
Write-Host "This test will show exactly what happens during directory switching." -ForegroundColor Yellow
