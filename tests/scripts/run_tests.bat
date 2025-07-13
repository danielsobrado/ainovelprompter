@echo off
echo Running Go tests...

echo.
echo === Testing current directory ==="
echo Current directory: %CD%
dir *.go

echo.
echo === Building project ===
go build
if %ERRORLEVEL% neq 0 (
    echo Build failed with error code %ERRORLEVEL%
    exit /b %ERRORLEVEL%
)

echo.
echo === Running quick tests ===
go test -v ./app_quick_test.go ./app.go

echo.
echo === Running extended tests ===
go test -v ./app_extended_test.go ./app.go

echo.
echo === Running storage tests ===
cd mcp
go test -v ./storage/...
cd ..

echo.
echo Tests completed.
