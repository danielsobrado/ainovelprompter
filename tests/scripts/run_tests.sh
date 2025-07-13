#!/bin/bash
echo "Running Go tests..."

echo ""
echo "=== Building project ==="
go build

echo ""
echo "=== Running quick tests ==="
go test -v ./app_quick_test.go ./app.go

echo ""
echo "=== Running extended tests ==="
go test -v ./app_extended_test.go ./app.go

echo ""
echo "=== Running storage tests ==="
cd mcp
go test -v ./storage/...
cd ..

echo ""
echo "Tests completed."
