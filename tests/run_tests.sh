#!/bin/bash
# AI Novel Prompter - Master Test Runner

echo "=== AI Novel Prompter - Master Test Runner ==="
echo "Organized test suite for systematic testing"

# Change to project root
cd "$(dirname "$0")"

echo ""
echo "Available Test Categories:"
echo ""
echo "1. Quick Verification    - ./tests/scripts/test_quick.ps1"
echo "2. All Tests            - ./tests/scripts/test_all_fixes.ps1"  
echo "3. Specific Fixes       - ./tests/scripts/test_specific_fixes.ps1"
echo "4. Directory Debug      - ./tests/scripts/test_directory_debug.ps1"
echo "5. Manual Commands      - Show individual test commands"
echo ""

# Prompt user for choice
echo -n "Select test category (1-5) or 'q' to quit: "
read choice

case $choice in
    1)
        echo "Running Quick Verification..."
        powershell -ExecutionPolicy Bypass -File "./tests/scripts/test_quick.ps1"
        ;;
    2)
        echo "Running All Tests..."
        powershell -ExecutionPolicy Bypass -File "./tests/scripts/test_all_fixes.ps1"
        ;;
    3)
        echo "Running Specific Fix Tests..."
        powershell -ExecutionPolicy Bypass -File "./tests/scripts/test_specific_fixes.ps1"
        ;;
    4)
        echo "Running Directory Debug..."
        powershell -ExecutionPolicy Bypass -File "./tests/scripts/test_directory_debug.ps1"
        ;;
    5)
        echo "Manual Test Commands:"
        echo ""
        echo "# Root level app tests:"
        echo "go test -v ./tests/app_quick_test.go ./app.go"
        echo "go test -v ./tests/app_extended_test.go ./app.go"
        echo ""
        echo "# MCP storage tests:"
        echo "cd mcp && go test -v ./tests/..."
        echo ""
        echo "# Specific tests:"
        echo "go test -v -run TestBasicStorageOperations ./tests/app_quick_test.go ./app.go"
        echo "go test -v -run TestMigrationFunctionality ./tests/app_extended_test.go ./app.go"
        ;;
    q)
        echo "Exiting..."
        exit 0
        ;;
    *)
        echo "Invalid choice. Running Quick Verification by default..."
        powershell -ExecutionPolicy Bypass -File "./tests/scripts/test_quick.ps1"
        ;;
esac

echo ""
echo "=== Test Structure Information ==="
echo "📁 tests/                 - Root level app tests"
echo "📁 tests/scripts/         - Test automation scripts"  
echo "📁 mcp/tests/             - MCP storage subsystem tests"
echo "📄 tests/README.md        - Comprehensive test documentation"
echo ""
echo "For detailed test documentation, see: tests/README.md"
