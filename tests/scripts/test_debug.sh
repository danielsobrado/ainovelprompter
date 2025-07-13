#!/bin/bash

# Set working directory
cd "C:\Development\workspace\GitHub\ainovelprompter"

echo "=== Current directory: $(pwd) ==="
echo "=== Go files present: ==="
ls -la *.go

echo ""
echo "=== Testing Go build ==="
go build -v

echo ""
echo "=== Running quick test individually ==="
go test -v -run TestBasicStorageOperations ./app_quick_test.go ./app.go

echo ""
echo "=== Running all quick tests ==="
go test -v ./app_quick_test.go ./app.go
