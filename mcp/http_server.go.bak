package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/danielsobrado/ainovelprompter/mcp"
)

type HTTPServer struct {
	mcpServer *mcp.MCPServer
}

type APIResponse struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data,omitempty"`
	Error   string      `json:"error,omitempty"`
	Tools   []string    `json:"tools,omitempty"`
}

func main() {
	fmt.Println("=== AI Novel Prompter HTTP API Server ===")

	// Initialize MCP Server
	mcpServer, err := mcp.NewMCPServer()
	if err != nil {
		log.Fatalf("Failed to create MCP server: %v", err)
	}

	httpServer := &HTTPServer{mcpServer: mcpServer}

	// Setup routes
	http.HandleFunc("/", httpServer.handleRoot)
	http.HandleFunc("/tools", httpServer.handleTools)
	http.HandleFunc("/execute", httpServer.handleExecute)
	http.HandleFunc("/test", httpServer.handleTest)

	// Start server
	port := "8080"
	fmt.Printf("🚀 HTTP API Server starting on http://localhost:%s\n", port)
	fmt.Println("\nAvailable endpoints:")
	fmt.Println("  GET  /          - API documentation")
	fmt.Println("  GET  /tools     - List all available MCP tools")
	fmt.Println("  POST /execute   - Execute MCP tool")
	fmt.Println("  GET  /test      - Run basic functionality tests")
	fmt.Println("\nTesting URLs:")
	fmt.Printf("  http://localhost:%s/tools\n", port)
	fmt.Printf("  http://localhost:%s/test\n", port)

	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func (s *HTTPServer) handleRoot(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	
	docs := map[string]interface{}{
		"service": "AI Novel Prompter MCP Server",
		"version": "1.0.0",
		"endpoints": map[string]interface{}{
			"GET /tools": "List all available MCP tools",
			"POST /execute": map[string]interface{}{
				"description": "Execute an MCP tool",
				"body": map[string]interface{}{
					"tool":   "string - tool name",
					"params": "object - tool parameters",
				},
				"example": map[string]interface{}{
					"tool": "get_characters",
					"params": map[string]interface{}{
						"search": "protagonist",
					},
				},
			},
			"GET /test": "Run basic functionality tests",
		},
		"example_tools": []string{
			"get_characters", "create_character", "get_chapters", "create_chapter",
			"get_prose_prompts", "analyze_prose", "search_all_content", "generate_chapter_prompt",
		},
	}

	json.NewEncoder(w).Encode(docs)
}

func (s *HTTPServer) handleTools(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	tools := s.mcpServer.GetTools()
	var toolNames []string
	toolDetails := make(map[string]interface{})

	for _, tool := range tools {
		toolNames = append(toolNames, tool.Name)
		toolDetails[tool.Name] = map[string]interface{}{
			"description": tool.Description,
			"parameters":  tool.Parameters,
		}
	}

	response := APIResponse{
		Success: true,
		Data: map[string]interface{}{
			"count":   len(tools),
			"tools":   toolNames,
			"details": toolDetails,
		},
	}

	json.NewEncoder(w).Encode(response)
}

func (s *HTTPServer) handleExecute(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var request struct {
		Tool   string                 `json:"tool"`
		Params map[string]interface{} `json:"params"`
	}

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		response := APIResponse{
			Success: false,
			Error:   fmt.Sprintf("Invalid JSON: %v", err),
		}
		json.NewEncoder(w).Encode(response)
		return
	}

	if request.Tool == "" {
		response := APIResponse{
			Success: false,
			Error:   "Tool name is required",
		}
		json.NewEncoder(w).Encode(response)
		return
	}

	if request.Params == nil {
		request.Params = make(map[string]interface{})
	}

	// Execute the MCP tool
	result, err := s.mcpServer.ExecuteTool(request.Tool, request.Params)
	if err != nil {
		response := APIResponse{
			Success: false,
			Error:   err.Error(),
		}
		json.NewEncoder(w).Encode(response)
		return
	}

	response := APIResponse{
		Success: true,
		Data:    result,
	}

	json.NewEncoder(w).Encode(response)
}

func (s *HTTPServer) handleTest(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	tests := []struct {
		name   string
		tool   string
		params map[string]interface{}
	}{
		{
			name: "Get Characters (empty)",
			tool: "get_characters",
			params: map[string]interface{}{},
		},
		{
			name: "Create Test Character",
			tool: "create_character",
			params: map[string]interface{}{
				"name":        "HTTP Test Character",
				"description": "A character created via HTTP API",
				"notes":       "Created during HTTP testing",
			},
		},
		{
			name: "Get Characters (after create)",
			tool: "get_characters",
			params: map[string]interface{}{},
		},
		{
			name: "Get Prose Prompts",
			tool: "get_prose_prompts",
			params: map[string]interface{}{},
		},
		{
			name: "Search All Content",
			tool: "search_all_content",
			params: map[string]interface{}{
				"query": "test",
				"limit": 5,
			},
		},
		{
			name: "Analyze Text Traits",
			tool: "analyze_text_traits",
			params: map[string]interface{}{
				"text": "This is a test sentence for analysis. It contains multiple sentences! How does it analyze?",
			},
		},
		{
			name: "Get Prompt Template",
			tool: "get_prompt_template",
			params: map[string]interface{}{
				"format": "ChatGPT",
			},
		},
	}

	var results []map[string]interface{}
	successCount := 0

	for _, test := range tests {
		result, err := s.mcpServer.ExecuteTool(test.tool, test.params)
		
		testResult := map[string]interface{}{
			"name":    test.name,
			"tool":    test.tool,
			"params":  test.params,
			"success": err == nil,
		}

		if err != nil {
			testResult["error"] = err.Error()
		} else {
			testResult["result"] = formatForDisplay(result)
			successCount++
		}

		results = append(results, testResult)
	}

	// Also test error handling
	_, err := s.mcpServer.ExecuteTool("invalid_tool", map[string]interface{}{})
	errorHandlingWorks := err != nil

	summary := map[string]interface{}{
		"total_tests":         len(tests),
		"successful_tests":    successCount,
		"failed_tests":        len(tests) - successCount,
		"error_handling_works": errorHandlingWorks,
		"success_rate":        fmt.Sprintf("%.1f%%", float64(successCount)/float64(len(tests))*100),
	}

	response := APIResponse{
		Success: true,
		Data: map[string]interface{}{
			"summary": summary,
			"tests":   results,
		},
	}

	json.NewEncoder(w).Encode(response)
}

func formatForDisplay(result interface{}) interface{} {
	// Convert result to JSON and back to limit depth and size
	jsonData, err := json.Marshal(result)
	if err != nil {
		return fmt.Sprintf("Error formatting: %v", err)
	}

	// Limit size for display
	str := string(jsonData)
	if len(str) > 500 {
		str = str[:500] + "...(truncated)"
	}

	// Try to parse back to object for better display
	var obj interface{}
	if err := json.Unmarshal([]byte(str), &obj); err == nil {
		return obj
	}

	return str
}
