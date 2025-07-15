package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/danielsobrado/ainovelprompter/mcp"
)

func main() {
	// Initialize MCP Server
	server, err := mcp.NewMCPServer()
	if err != nil {
		log.Fatalf("Failed to create MCP server: %v", err)
	}

	// Example: Get all tools
	tools := server.GetTools()
	fmt.Printf("MCP Server initialized with %d tools:\n", len(tools))

	for _, tool := range tools {
		fmt.Printf("- %s: %s\n", tool.Name, tool.Description)
	}

	// Example: Test a tool
	fmt.Println("\nTesting get_characters tool:")
	result, err := server.ExecuteTool("get_characters", map[string]interface{}{})
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		jsonData, _ := json.MarshalIndent(result, "", "  ")
		fmt.Printf("Result: %s\n", jsonData)
	}

	fmt.Println("\nMCP Server ready!")
}
