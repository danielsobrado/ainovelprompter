package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/danielsobrado/ainovelprompter/mcp"
)

type MCPMessage struct {
	JSONRPC string      `json:"jsonrpc"`
	ID      interface{} `json:"id,omitempty"`
	Method  string      `json:"method,omitempty"`
	Params  interface{} `json:"params,omitempty"`
	Result  interface{} `json:"result,omitempty"`
	Error   *MCPError   `json:"error,omitempty"`
}

type MCPError struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

type InitializeParams struct {
	ProtocolVersion string                 `json:"protocolVersion"`
	Capabilities    map[string]interface{} `json:"capabilities"`
	ClientInfo      ClientInfo             `json:"clientInfo"`
}

type ClientInfo struct {
	Name    string `json:"name"`
	Version string `json:"version"`
}

type ServerCapabilities struct {
	Tools map[string]interface{} `json:"tools,omitempty"`
}

type ServerInfo struct {
	Name    string `json:"name"`
	Version string `json:"version"`
}

type MCPStdioServer struct {
	mcpServer    *mcp.MCPServer
	initialized  bool
	capabilities ServerCapabilities
}

func main() {
	// Parse command line arguments
	var dataDir string
	var logLevel string
	var showHelp bool

	flag.StringVar(&dataDir, "data-dir", "", "Data directory path (defaults to ~/.ai-novel-prompter)")
	flag.StringVar(&dataDir, "d", "", "Data directory path (short form)")
	flag.StringVar(&logLevel, "log-level", "", "Log level (DEBUG, INFO, WARN, ERROR)")
	flag.StringVar(&logLevel, "l", "", "Log level (short form)")
	flag.BoolVar(&showHelp, "help", false, "Show help message")
	flag.BoolVar(&showHelp, "h", false, "Show help message (short form)")

	flag.Parse()

	if showHelp {
		showHelpMessage()
		os.Exit(0)
	}

	// Set log level if provided
	if logLevel != "" {
		os.Setenv("AINOVEL_LOG_LEVEL", logLevel)
		log.Printf("Log level set to: %s", logLevel)
	}

	// Set up logging to stderr so it doesn't interfere with MCP communication
	log.SetOutput(os.Stderr)

	server := &MCPStdioServer{}

	// Initialize our MCP server with optional data directory
	var mcpServer *mcp.MCPServer
	var err error

	if dataDir != "" {
		log.Printf("Using data directory: %s", dataDir)
		mcpServer, err = mcp.NewMCPServerWithDataDir(dataDir)
	} else {
		mcpServer, err = mcp.NewMCPServer()
	}

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

	log.Println("AI Novel Prompter MCP Server starting...")

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
		"protocolVersion": "2024-11-05",
		"capabilities":    s.capabilities,
		"serverInfo": ServerInfo{
			Name:    "ai-novel-prompter-mcp",
			Version: "1.0.0",
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

// showHelpMessage displays the command line help
func showHelpMessage() {
	fmt.Fprintf(os.Stderr, `AI Novel Prompter MCP Server

Usage: %s [OPTIONS]

Options:
  -d, --data-dir PATH      Data directory path (default: ~/.ai-novel-prompter)
  -l, --log-level LEVEL    Log level (DEBUG, INFO, WARN, ERROR) (default: INFO)
  -h, --help              Show this help message

Examples:
  %s                                          # Use defaults
  %s -d ./my-story                           # Use relative path
  %s --data-dir /path/to/story/data         # Use absolute path
  %s -d "C:\My Stories\Novel Data"         # Windows path with spaces
  %s --log-level DEBUG                      # Enable debug logging
  %s -d ./data -l DEBUG                     # Custom data dir with debug logging

Log Levels:
  DEBUG    - Detailed operation logging (verbose)
  INFO     - Standard operational logging (default)
  WARN     - Warnings and unexpected conditions only
  ERROR    - Errors only

The MCP server will create the data directory if it doesn't exist.
This allows sharing data between the desktop app and MCP server.
`, os.Args[0], os.Args[0], os.Args[0], os.Args[0], os.Args[0], os.Args[0], os.Args[0])
}
