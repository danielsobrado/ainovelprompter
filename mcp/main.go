package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/danielsobrado/ainovelprompter/mcp"
)

// Configuration constants
const (
	DEFAULT_DATA_DIR = ".ai-novel-prompter"
	SERVER_NAME     = "ai-novel-prompter-mcp"
	SERVER_VERSION  = "1.0.0"
	PROTOCOL_VERSION = "2024-11-05"
)

// MCPMessage represents a JSON-RPC 2.0 message for MCP
type MCPMessage struct {
	JSONRPC string      `json:"jsonrpc"`
	ID      interface{} `json:"id,omitempty"`
	Method  string      `json:"method,omitempty"`
	Params  interface{} `json:"params,omitempty"`
	Result  interface{} `json:"result,omitempty"`
	Error   *MCPError   `json:"error,omitempty"`
}

// MCPError represents an error in the MCP protocol
type MCPError struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

// InitializeParams contains initialization parameters for the server
type InitializeParams struct {
	ProtocolVersion string                 `json:"protocolVersion"`
	Capabilities    map[string]interface{} `json:"capabilities"`
	ClientInfo      ClientInfo             `json:"clientInfo"`
}

// ClientInfo contains information about the client
type ClientInfo struct {
	Name    string `json:"name"`
	Version string `json:"version"`
}

// ServerCapabilities defines what the server can do
type ServerCapabilities struct {
	Tools map[string]interface{} `json:"tools,omitempty"`
}

// ServerInfo contains information about this server
type ServerInfo struct {
	Name    string `json:"name"`
	Version string `json:"version"`
}

// MCPStdioServer handles MCP communication over stdio
type MCPStdioServer struct {
	mcpServer    *mcp.MCPServer
	initialized  bool
	capabilities ServerCapabilities
}

func main() {
	// Parse command line arguments
	var dataDir string
	var showHelp bool
	
	flag.StringVar(&dataDir, "data-dir", "", "Data directory path (defaults to ~/.ai-novel-prompter)")
	flag.StringVar(&dataDir, "d", "", "Data directory path (short form)")
	flag.BoolVar(&showHelp, "help", false, "Show help message")
	flag.BoolVar(&showHelp, "h", false, "Show help message (short form)")
	
	flag.Parse()
	
	if showHelp {
		showHelpMessage()
		os.Exit(0)
	}
	
	// Set up logging to stderr so it doesn't interfere with MCP communication
	log.SetOutput(os.Stderr)
	
	// Resolve data directory
	resolvedDataDir := resolveDataDirectory(dataDir)
	
	server := &MCPStdioServer{}
	
	// Initialize our MCP server with the specified data directory
	mcpServer, err := mcp.NewMCPServerWithDataDir(resolvedDataDir)
	if err != nil {
		log.Fatalf("Failed to create MCP server: %v", err)
	}
	server.mcpServer = mcpServer
	
	// Set up capabilities
	server.capabilities = ServerCapabilities{
		Tools: map[string]interface{}{
			"listChanged": false,
		},
	}
	
	log.Printf("AI Novel Prompter MCP Server starting...")
	log.Printf("Data directory: %s", resolvedDataDir)
	
	// Handle stdin/stdout communication
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}
		
		log.Printf("Received: %s", line)
		
		var message MCPMessage
		if err := json.Unmarshal([]byte(line), &message); err != nil {
			log.Printf("Error parsing JSON: %v", err)
			server.sendError(nil, -32700, "Parse error", nil)
			continue
		}
		
		response := server.handleMessage(message)
		if response != nil {
			server.sendMessage(*response)
		}
	}
	
	if err := scanner.Err(); err != nil {
		log.Printf("Error reading stdin: %v", err)
	}
}

func showHelpMessage() {
	fmt.Fprintf(os.Stderr, `AI Novel Prompter MCP Server

Usage: %s [OPTIONS]

Options:
  -d, --data-dir PATH    Data directory path (default: ~/.ai-novel-prompter)
  -h, --help            Show this help message

Examples:
  %s                                        # Use default data directory
  %s -d ./my-story                         # Use relative path
  %s --data-dir /path/to/story/data       # Use absolute path
  %s -d "C:\My Stories\Novel Data"        # Windows path with spaces

The server communicates using the MCP (Model Context Protocol) over stdin/stdout.
`, os.Args[0], os.Args[0], os.Args[0], os.Args[0], os.Args[0])
}

func resolveDataDirectory(dataDir string) string {
	if dataDir == "" {
		// Use default location in user home directory
		homeDir, err := os.UserHomeDir()
		if err != nil {
			log.Fatalf("Failed to get user home directory: %v", err)
		}
		return filepath.Join(homeDir, DEFAULT_DATA_DIR)
	}
	
	// Expand relative paths to absolute paths
	absPath, err := filepath.Abs(dataDir)
	if err != nil {
		log.Fatalf("Failed to resolve data directory path: %v", err)
	}
	
	return absPath
}

func (s *MCPStdioServer) handleMessage(message MCPMessage) *MCPMessage {
	switch message.Method {
	case "initialize":
		return s.handleInitialize(message)
	case "initialized", "notifications/initialized":
		return s.handleInitialized(message)
	case "tools/list":
		return s.handleToolsList(message)
	case "tools/call":
		return s.handleToolsCall(message)
	default:
		return &MCPMessage{
			JSONRPC: "2.0",
			ID:      message.ID,
			Error: &MCPError{
				Code:    -32601,
				Message: "Method not found",
				Data:    message.Method,
			},
		}
	}
}

func (s *MCPStdioServer) handleInitialize(message MCPMessage) *MCPMessage {
	log.Println("Handling initialize request")
	
	result := map[string]interface{}{
		"protocolVersion": PROTOCOL_VERSION,
		"capabilities":    s.capabilities,
		"serverInfo": ServerInfo{
			Name:    SERVER_NAME,
			Version: SERVER_VERSION,
		},
	}
	
	return &MCPMessage{
		JSONRPC: "2.0",
		ID:      message.ID,
		Result:  result,
	}
}

func (s *MCPStdioServer) handleInitialized(message MCPMessage) *MCPMessage {
	log.Println("Handling initialized notification")
	s.initialized = true
	return nil // No response for notifications
}

func (s *MCPStdioServer) handleToolsList(message MCPMessage) *MCPMessage {
	log.Println("Handling tools/list request")
	
	if !s.initialized {
		return &MCPMessage{
			JSONRPC: "2.0",
			ID:      message.ID,
			Error: &MCPError{
				Code:    -32002,
				Message: "Server not initialized",
			},
		}
	}
	
	tools := s.mcpServer.GetTools()
	mcpTools := make([]map[string]interface{}, len(tools))
	
	for i, tool := range tools {
		// Convert parameters to MCP format
		properties := make(map[string]interface{})
		required := make([]string, 0)
		
		for name, param := range tool.Parameters {
			properties[name] = map[string]interface{}{
				"type":        param.Type,
				"description": param.Description,
			}
			if param.Required {
				required = append(required, name)
			}
		}
		
		inputSchema := map[string]interface{}{
			"type":       "object",
			"properties": properties,
		}
		
		if len(required) > 0 {
			inputSchema["required"] = required
		}
		
		mcpTools[i] = map[string]interface{}{
			"name":        tool.Name,
			"description": tool.Description,
			"inputSchema": inputSchema,
		}
	}
	
	result := map[string]interface{}{
		"tools": mcpTools,
	}
	
	return &MCPMessage{
		JSONRPC: "2.0",
		ID:      message.ID,
		Result:  result,
	}
}

func (s *MCPStdioServer) handleToolsCall(message MCPMessage) *MCPMessage {
	log.Println("Handling tools/call request")
	
	if !s.initialized {
		return &MCPMessage{
			JSONRPC: "2.0",
			ID:      message.ID,
			Error: &MCPError{
				Code:    -32002,
				Message: "Server not initialized",
			},
		}
	}
	
	params, ok := message.Params.(map[string]interface{})
	if !ok {
		return &MCPMessage{
			JSONRPC: "2.0",
			ID:      message.ID,
			Error: &MCPError{
				Code:    -32602,
				Message: "Invalid params",
			},
		}
	}
	
	toolName, ok := params["name"].(string)
	if !ok {
		return &MCPMessage{
			JSONRPC: "2.0",
			ID:      message.ID,
			Error: &MCPError{
				Code:    -32602,
				Message: "Missing tool name",
			},
		}
	}
	
	arguments, ok := params["arguments"].(map[string]interface{})
	if !ok {
		arguments = make(map[string]interface{})
	}
	
	log.Printf("Executing tool: %s with arguments: %v", toolName, arguments)
	
	// Execute the tool using our MCP server
	result, err := s.mcpServer.ExecuteTool(toolName, arguments)
	if err != nil {
		log.Printf("Tool execution error: %v", err)
		return &MCPMessage{
			JSONRPC: "2.0",
			ID:      message.ID,
			Error: &MCPError{
				Code:    -32603,
				Message: "Tool execution failed",
				Data:    err.Error(),
			},
		}
	}
	
	// Format result as MCP tool response
	toolResult := map[string]interface{}{
		"content": []map[string]interface{}{
			{
				"type": "text",
				"text": formatResult(result),
			},
		},
	}
	
	return &MCPMessage{
		JSONRPC: "2.0",
		ID:      message.ID,
		Result:  toolResult,
	}
}

func (s *MCPStdioServer) sendMessage(message MCPMessage) {
	data, err := json.Marshal(message)
	if err != nil {
		log.Printf("Error marshaling message: %v", err)
		return
	}
	
	fmt.Println(string(data))
	log.Printf("Sent: %s", string(data))
}

func (s *MCPStdioServer) sendError(id interface{}, code int, message string, data interface{}) {
	response := MCPMessage{
		JSONRPC: "2.0",
		ID:      id,
		Error: &MCPError{
			Code:    code,
			Message: message,
			Data:    data,
		},
	}
	s.sendMessage(response)
}

func formatResult(result interface{}) string {
	data, err := json.MarshalIndent(result, "", "  ")
	if err != nil {
		return fmt.Sprintf("Error formatting result: %v", err)
	}
	return string(data)
}
