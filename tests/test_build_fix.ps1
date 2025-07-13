# Quick Build Test
Write-Host "=== Testing Build Fix ===" -ForegroundColor Cyan

Set-Location "C:\Development\workspace\GitHub\ainovelprompter"

Write-Host ""
Write-Host "=== Build Test ===" -ForegroundColor Yellow
$build = go build -v
if ($LASTEXITCODE -eq 0) {
    Write-Host "[+] Build successful!" -ForegroundColor Green
    Write-Host "The build error has been fixed." -ForegroundColor Green
} else {
    Write-Host "[-] Build still failing:" -ForegroundColor Red
    Write-Host $build -ForegroundColor White
}

Write-Host ""
Write-Host "=== Quick Function Test ===" -ForegroundColor Yellow
if ($LASTEXITCODE -eq 0) {
    Write-Host "Running a quick test to verify functionality..." -ForegroundColor Yellow
    $quickTest = go test -v -run TestBasicStorageOperations .\tests\app_quick_test.go .\app.go
    if ($LASTEXITCODE -eq 0) {
        Write-Host "[+] Quick test PASSED!" -ForegroundColor Green
        Write-Host "The build fix is working correctly." -ForegroundColor Green
    } else {
        Write-Host "[-] Quick test failed" -ForegroundColor Red
    }
} else {
    Write-Host "Skipping functional test due to build failure." -ForegroundColor Yellow
}
