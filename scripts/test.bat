@echo off
setlocal enabledelayedexpansion

REM Test runner script for AI Novel Prompter (Windows)

echo 🧪 Running AI Novel Prompter Test Suite
echo =======================================

REM Parse command line arguments
set FRONTEND_ONLY=false
set BACKEND_ONLY=false
set COVERAGE=false
set WATCH=false
set UI=false

:parse_args
if "%~1"=="" goto :main
if "%~1"=="--frontend-only" (
    set FRONTEND_ONLY=true
    shift
    goto :parse_args
)
if "%~1"=="--backend-only" (
    set BACKEND_ONLY=true
    shift
    goto :parse_args
)
if "%~1"=="--coverage" (
    set COVERAGE=true
    shift
    goto :parse_args
)
if "%~1"=="--watch" (
    set WATCH=true
    shift
    goto :parse_args
)
if "%~1"=="--ui" (
    set UI=true
    shift
    goto :parse_args
)
if "%~1"=="--help" (
    echo Usage: %0 [OPTIONS]
    echo Options:
    echo   --frontend-only    Run only frontend tests
    echo   --backend-only     Run only backend tests
    echo   --coverage         Generate coverage reports
    echo   --watch           Run tests in watch mode
    echo   --ui              Run frontend tests with UI
    echo   --help            Show this help message
    exit /b 0
)
echo Unknown option: %~1
echo Use --help for usage information
exit /b 1

:main
REM Check if we're in the right directory
if not exist "go.mod" (
    echo [ERROR] This script must be run from the project root directory
    exit /b 1
)
if not exist "frontend" (
    echo [ERROR] Frontend directory not found
    exit /b 1
)

echo [INFO] Checking dependencies...

REM Install Go dependencies
if "%FRONTEND_ONLY%"=="false" (
    echo [INFO] Installing Go dependencies...
    go mod download
    go mod tidy
    if errorlevel 1 (
        echo [ERROR] Failed to install Go dependencies
        exit /b 1
    )
)

REM Install Node dependencies
if "%BACKEND_ONLY%"=="false" (
    if not exist "frontend\node_modules" (
        echo [INFO] Installing Node.js dependencies...
        cd frontend
        npm install
        if errorlevel 1 (
            echo [ERROR] Failed to install Node.js dependencies
            cd ..
            exit /b 1
        )
        cd ..
    )
)

REM Run tests based on options
if "%FRONTEND_ONLY%"=="true" (
    call :run_frontend_tests
    if errorlevel 1 exit /b 1
) else if "%BACKEND_ONLY%"=="true" (
    call :run_go_tests
    if errorlevel 1 exit /b 1
) else (
    call :run_go_tests
    if errorlevel 1 (
        echo [ERROR] Backend tests failed. Skipping frontend tests.
        exit /b 1
    )
    call :run_frontend_tests
    if errorlevel 1 exit /b 1
)

echo.
echo [SUCCESS] All tests passed! 🎉
echo.

if "%COVERAGE%"=="true" (
    echo [INFO] Coverage reports generated:
    if "%FRONTEND_ONLY%"=="false" echo   - Go: coverage.html
    if "%BACKEND_ONLY%"=="false" echo   - Frontend: frontend\coverage\
)

exit /b 0

:run_go_tests
echo [INFO] Running Go tests...

if "%COVERAGE%"=="true" (
    echo [INFO] Running Go tests with coverage...
    go test -v -race -coverprofile=coverage.out -covermode=atomic ./...
    if errorlevel 1 (
        echo [ERROR] Go tests failed!
        exit /b 1
    )
    go tool cover -html=coverage.out -o coverage.html
    echo [SUCCESS] Go coverage report generated: coverage.html
) else (
    go test -v -race ./...
    if errorlevel 1 (
        echo [ERROR] Go tests failed!
        exit /b 1
    )
)

echo [SUCCESS] Go tests passed!
exit /b 0

:run_frontend_tests
echo [INFO] Running frontend tests...

cd frontend

if "%UI%"=="true" (
    echo [INFO] Starting test UI...
    npm run test:ui
) else if "%WATCH%"=="true" (
    echo [INFO] Starting test watcher...
    npm run test:watch
) else if "%COVERAGE%"=="true" (
    echo [INFO] Running frontend tests with coverage...
    npm run test:coverage
    if errorlevel 1 (
        echo [ERROR] Frontend tests failed!
        cd ..
        exit /b 1
    )
    echo [SUCCESS] Frontend coverage report generated in coverage\ directory
) else (
    npm run test:run
    if errorlevel 1 (
        echo [ERROR] Frontend tests failed!
        cd ..
        exit /b 1
    )
)

cd ..
echo [SUCCESS] Frontend tests passed!
exit /b 0
