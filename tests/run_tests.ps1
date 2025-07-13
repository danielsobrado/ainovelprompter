# AI Novel Prompter - Master Test Runner
Write-Host "=== AI Novel Prompter - Master Test Runner ===" -ForegroundColor Cyan
Write-Host "Organized test suite for systematic testing" -ForegroundColor Yellow

# Change to project root
Set-Location "C:\Development\workspace\GitHub\ainovelprompter"

Write-Host ""
Write-Host "Available Test Categories:" -ForegroundColor White
Write-Host ""
Write-Host "1. Quick Verification    - .\tests\scripts\test_quick.ps1" -ForegroundColor Green
Write-Host "2. All Tests            - .\tests\scripts\test_all_fixes.ps1" -ForegroundColor Yellow  
Write-Host "3. Specific Fixes       - .\tests\scripts\test_specific_fixes.ps1" -ForegroundColor Yellow
Write-Host "4. Directory Debug      - .\tests\scripts\test_directory_debug.ps1" -ForegroundColor Red
Write-Host "5. Filename Parsing     - .\tests\scripts\test_filename_parsing.ps1" -ForegroundColor Cyan
Write-Host ""

# Prompt user for choice
Write-Host "Select test category (0-5) or 'q' to quit: " -NoNewline -ForegroundColor White
$choice = Read-Host

switch ($choice) {
    "1" { 
        Write-Host "Running Quick Verification..." -ForegroundColor Green
        & ".\tests\scripts\test_quick.ps1"
    }
    "2" { 
        Write-Host "Running All Tests..." -ForegroundColor Yellow
        & ".\tests\scripts\test_all_fixes.ps1"
    }
    "3" { 
        Write-Host "Running Specific Fix Tests..." -ForegroundColor Yellow
        & ".\tests\scripts\test_specific_fixes.ps1"
    }
    "4" { 
        Write-Host "Running Directory Debug..." -ForegroundColor Red
        & ".\tests\scripts\test_directory_debug.ps1"
    }
    "5" { 
        Write-Host "Running Filename Parsing Tests..." -ForegroundColor Cyan
        & ".\tests\scripts\test_filename_parsing.ps1"
    }
    "q" { 
        Write-Host "Exiting..." -ForegroundColor Gray
        exit 0
    }
    default { 
        Write-Host "Invalid choice. Running Quick Verification by default..." -ForegroundColor Yellow
        & ".\tests\scripts\test_quick.ps1"
    }
}

Write-Host ""
Write-Host "=== Test Structure Information ===" -ForegroundColor Cyan
Write-Host "[DIR] tests/                 - Root level app tests" -ForegroundColor White
Write-Host "[DIR] tests/scripts/         - Test automation scripts" -ForegroundColor White  
Write-Host "[DIR] mcp/tests/             - MCP storage subsystem tests" -ForegroundColor White
Write-Host "[FILE] tests/README.md        - Comprehensive test documentation" -ForegroundColor White
Write-Host ""
Write-Host "For detailed test documentation, see: tests/README.md" -ForegroundColor Yellow
