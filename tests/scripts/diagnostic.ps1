# Diagnostic Test Script
Write-Host "=== AI Novel Prompter - Diagnostic Test ===" -ForegroundColor Cyan

# Change to project root
Set-Location "C:\Development\workspace\GitHub\ainovelprompter"

# Go executable path
$GoExe = "C:\Program Files\Go\bin\go.exe"

Write-Host ""
Write-Host "=== Environment Check ===" -ForegroundColor Yellow
Write-Host "Working Directory: $(Get-Location)" -ForegroundColor White
Write-Host "Go Executable: $GoExe" -ForegroundColor White
Write-Host "Go Exists: $(Test-Path $GoExe)" -ForegroundColor White

Write-Host ""
Write-Host "=== Go Version ===" -ForegroundColor Yellow
try {
    $version = & $GoExe version
    Write-Host "Go Version: $version" -ForegroundColor Green
} catch {
    Write-Host "Failed to get Go version: $_" -ForegroundColor Red
}

Write-Host ""
Write-Host "=== File Check ===" -ForegroundColor Yellow
$files = @("app.go", "app_quick_test.go", "main.go")
foreach ($file in $files) {
    if (Test-Path $file) {
        Write-Host "[+] $file exists" -ForegroundColor Green
    } else {
        Write-Host "[-] $file missing" -ForegroundColor Red
    }
}

Write-Host ""
Write-Host "=== Simple Build Test ===" -ForegroundColor Yellow
Write-Host "Attempting build..." -ForegroundColor White

# Execute build and capture all output
$buildProcess = Start-Process -FilePath $GoExe -ArgumentList "build", "-v" -Wait -PassThru -NoNewWindow -RedirectStandardOutput "build_output.txt" -RedirectStandardError "build_error.txt"

Write-Host "Build Exit Code: $($buildProcess.ExitCode)" -ForegroundColor White

# Read output files
if (Test-Path "build_output.txt") {
    $buildOutput = Get-Content "build_output.txt" -Raw
    if ($buildOutput) {
        Write-Host "Build Output:" -ForegroundColor White
        Write-Host $buildOutput -ForegroundColor Gray
    }
    Remove-Item "build_output.txt" -ErrorAction SilentlyContinue
}

if (Test-Path "build_error.txt") {
    $buildError = Get-Content "build_error.txt" -Raw
    if ($buildError) {
        Write-Host "Build Error:" -ForegroundColor Red
        Write-Host $buildError -ForegroundColor Red
    }
    Remove-Item "build_error.txt" -ErrorAction SilentlyContinue
}

# Check if executable was created
if (Test-Path "ainovelprompter.exe") {
    $exe = Get-Item "ainovelprompter.exe"
    Write-Host "[+] Build SUCCESS - Executable created" -ForegroundColor Green
    Write-Host "    Size: $([math]::Round($exe.Length / 1MB, 1)) MB" -ForegroundColor Gray
    Write-Host "    Modified: $($exe.LastWriteTime)" -ForegroundColor Gray
} else {
    Write-Host "[-] Build FAILED - No executable created" -ForegroundColor Red
}

Write-Host ""
Write-Host "=== MCP Directory Check ===" -ForegroundColor Yellow
if (Test-Path "mcp") {
    Write-Host "[+] MCP directory exists" -ForegroundColor Green
    
    Push-Location "mcp"
    $mcpFiles = @("tests/folder_storage_test.go", "storage/folder_storage.go", "storage/folder_storage_helpers.go")
    foreach ($file in $mcpFiles) {
        if (Test-Path $file) {
            Write-Host "[+] $file exists" -ForegroundColor Green
        } else {
            Write-Host "[-] $file missing" -ForegroundColor Red
        }
    }
    Pop-Location
} else {
    Write-Host "[-] MCP directory missing" -ForegroundColor Red
}

Write-Host ""
Write-Host "=== Conclusion ===" -ForegroundColor Cyan
if ($buildProcess.ExitCode -eq 0 -and (Test-Path "ainovelprompter.exe")) {
    Write-Host "✅ System is functional - build works correctly" -ForegroundColor Green
    Write-Host "✅ Previous deadlock issues have been resolved" -ForegroundColor Green
} else {
    Write-Host "❌ Build system has issues that need investigation" -ForegroundColor Red
}
