#!/bin/bash

echo "=== AI Novel Prompter Test Verification ==="
echo "Testing fixes for remaining issues..."

# Change to project directory
cd "C:\Development\workspace\GitHub\ainovelprompter"

echo ""
echo "=== Build Test ==="
echo "Building project..."
if go build -v; then
    echo "✅ Build successful"
else
    echo "❌ Build failed"
    exit 1
fi

echo ""
echo "=== Quick Tests ==="
echo "Running quick tests..."
if go test -v ./app_quick_test.go ./app.go; then
    echo "✅ Quick tests passed"
else
    echo "❌ Quick tests failed"
fi

echo ""
echo "=== Extended Tests ==="
echo "Running extended tests..."
if go test -v ./app_extended_test.go ./app.go; then
    echo "✅ Extended tests passed"
else
    echo "❌ Extended tests failed"
fi

echo ""
echo "=== Storage Tests ==="
echo "Running storage tests..."
cd mcp
if go test -v ./storage/...; then
    echo "✅ Storage tests passed"
else
    echo "❌ Storage tests failed"
fi
cd ..

echo ""
echo "=== Test Summary ==="
echo "All tests completed. Check output above for specific failures."
