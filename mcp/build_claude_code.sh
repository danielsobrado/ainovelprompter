#!/bin/bash

# Build AI Novel Prompter MCP Server for Claude Code
echo "🔨 Building AI Novel Prompter MCP Server for Claude Code..."

# Check if we're in the right directory
if [ ! -f "mcp_stdio_server.go" ]; then
    echo "❌ Error: mcp_stdio_server.go not found. Please run this script from the mcp/ directory."
    exit 1
fi

# Check if Go is installed
if ! command -v go &> /dev/null; then
    echo "❌ Error: Go is not installed. Please install Go from https://golang.org/dl/"
    exit 1
fi

echo "📦 Installing dependencies..."
go mod tidy

if [ $? -ne 0 ]; then
    echo "❌ Error: Failed to download dependencies"
    exit 1
fi

echo "🔧 Building MCP server executable..."
go build -o ainovelprompter-mcp mcp_stdio_server.go

if [ $? -ne 0 ]; then
    echo "❌ Error: Build failed"
    exit 1
fi

# Make executable (Unix/Mac)
chmod +x ainovelprompter-mcp

echo "✅ MCP Server built successfully!"
echo "📍 Executable location: $(pwd)/ainovelprompter-mcp"

# Test the server
echo "🧪 Testing MCP server..."
echo '{"jsonrpc":"2.0","id":1,"method":"initialize","params":{"protocolVersion":"2024-11-05","capabilities":{},"clientInfo":{"name":"test","version":"1.0.0"}}}' | ./ainovelprompter-mcp > /dev/null 2>&1

if [ $? -eq 0 ]; then
    echo "✅ Server test passed!"
else
    echo "⚠️  Server test failed, but executable was created"
fi

echo ""
echo "🎉 Ready for Claude Code integration!"
echo ""
echo "Next steps:"
echo "1. Add this server to your Claude Code configuration:"
echo "   Path: $(pwd)/ainovelprompter-mcp"
echo "2. Configuration file location:"
echo "   - macOS: ~/Library/Application Support/Claude/claude_desktop_config.json"
echo "   - Linux: ~/.config/claude/claude_desktop_config.json"
echo "3. Restart Claude Code to load the MCP server"
echo ""
echo "See CLAUDE_CODE_INSTALLATION.md for detailed setup instructions."