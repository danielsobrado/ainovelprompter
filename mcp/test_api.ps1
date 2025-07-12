# PowerShell script to test MCP HTTP API endpoints
# Run this after starting the HTTP server with: .\http_server.exe

Write-Host "=== Testing AI Novel Prompter MCP HTTP API ===" -ForegroundColor Cyan
Write-Host ""

$baseUrl = "http://localhost:8080"

# Function to make HTTP requests
function Test-Endpoint {
    param(
        [string]$Method,
        [string]$Url,
        [string]$Body = $null,
        [string]$Description
    )
    
    Write-Host "Testing: $Description" -ForegroundColor Yellow
    Write-Host "  $Method $Url" -ForegroundColor Gray
    
    try {
        if ($Body) {
            $response = Invoke-RestMethod -Uri $Url -Method $Method -Body $Body -ContentType "application/json"
        } else {
            $response = Invoke-RestMethod -Uri $Url -Method $Method
        }
        
        Write-Host "  ✅ Status: Success" -ForegroundColor Green
        if ($response.success -eq $false) {
            Write-Host "  ❌ API Error: $($response.error)" -ForegroundColor Red
        } else {
            Write-Host "  📊 Response: $($response | ConvertTo-Json -Depth 2 -Compress)" -ForegroundColor White
        }
    } catch {
        Write-Host "  ❌ HTTP Error: $($_.Exception.Message)" -ForegroundColor Red
    }
    Write-Host ""
}

# Check if server is running
Write-Host "Checking if HTTP server is running..." -ForegroundColor Yellow
try {
    $ping = Invoke-RestMethod -Uri "$baseUrl/" -Method GET -TimeoutSec 5
    Write-Host "✅ Server is running!" -ForegroundColor Green
} catch {
    Write-Host "❌ Server is not running. Please start it with: .\http_server.exe" -ForegroundColor Red
    Write-Host "Then run this script again." -ForegroundColor Yellow
    Read-Host "Press Enter to exit"
    exit 1
}
Write-Host ""

# Test 1: Get API documentation
Test-Endpoint -Method "GET" -Url "$baseUrl/" -Description "API Documentation"

# Test 2: Get all tools
Test-Endpoint -Method "GET" -Url "$baseUrl/tools" -Description "List all MCP tools"

# Test 3: Run automated tests
Test-Endpoint -Method "GET" -Url "$baseUrl/test" -Description "Run comprehensive tests"

# Test 4: Execute specific tools
Write-Host "=== Testing Individual Tools ===" -ForegroundColor Cyan
Write-Host ""

# Test get_characters
$body = @{
    tool = "get_characters"
    params = @{}
} | ConvertTo-Json

Test-Endpoint -Method "POST" -Url "$baseUrl/execute" -Body $body -Description "Get Characters (empty)"

# Test create_character
$body = @{
    tool = "create_character"
    params = @{
        name = "API Test Character"
        description = "A character created via HTTP API"
        notes = "Created during API testing"
    }
} | ConvertTo-Json

Test-Endpoint -Method "POST" -Url "$baseUrl/execute" -Body $body -Description "Create Character"

# Test get_characters again
$body = @{
    tool = "get_characters"
    params = @{}
} | ConvertTo-Json

Test-Endpoint -Method "POST" -Url "$baseUrl/execute" -Body $body -Description "Get Characters (after create)"

# Test search
$body = @{
    tool = "search_all_content"
    params = @{
        query = "test"
        limit = 5
    }
} | ConvertTo-Json

Test-Endpoint -Method "POST" -Url "$baseUrl/execute" -Body $body -Description "Search All Content"

# Test prose prompts
$body = @{
    tool = "get_prose_prompts"
    params = @{}
} | ConvertTo-Json

Test-Endpoint -Method "POST" -Url "$baseUrl/execute" -Body $body -Description "Get Prose Prompts"

# Test text analysis
$body = @{
    tool = "analyze_text_traits"
    params = @{
        text = "This is a sample text for analysis. It has multiple sentences! Does it work? Yes, it should work well."
    }
} | ConvertTo-Json

Test-Endpoint -Method "POST" -Url "$baseUrl/execute" -Body $body -Description "Analyze Text Traits"

# Test prompt generation
$body = @{
    tool = "generate_chapter_prompt"
    params = @{
        promptType = "ChatGPT"
        taskType = "Write the next chapter continuing the story"
        nextChapterBeats = "The protagonist makes an important discovery"
    }
} | ConvertTo-Json

Test-Endpoint -Method "POST" -Url "$baseUrl/execute" -Body $body -Description "Generate Chapter Prompt"

# Test error handling
Write-Host "=== Testing Error Handling ===" -ForegroundColor Cyan
Write-Host ""

# Test invalid tool
$body = @{
    tool = "invalid_tool_name"
    params = @{}
} | ConvertTo-Json

Test-Endpoint -Method "POST" -Url "$baseUrl/execute" -Body $body -Description "Invalid Tool (should fail)"

# Test missing parameters
$body = @{
    tool = "create_character"
    params = @{}
} | ConvertTo-Json

Test-Endpoint -Method "POST" -Url "$baseUrl/execute" -Body $body -Description "Missing Required Parameters (should fail)"

# Summary
Write-Host "=== Test Summary ===" -ForegroundColor Cyan
Write-Host ""
Write-Host "✅ All HTTP endpoint tests completed!" -ForegroundColor Green
Write-Host ""
Write-Host "Key endpoints tested:" -ForegroundColor White
Write-Host "  • GET  /          - API documentation" -ForegroundColor Gray
Write-Host "  • GET  /tools     - Tool discovery" -ForegroundColor Gray
Write-Host "  • GET  /test      - Automated testing" -ForegroundColor Gray
Write-Host "  • POST /execute   - Tool execution" -ForegroundColor Gray
Write-Host ""
Write-Host "Tool categories tested:" -ForegroundColor White
Write-Host "  • Story Context Management" -ForegroundColor Gray
Write-Host "  • Search & Analysis" -ForegroundColor Gray
Write-Host "  • Prose Improvement" -ForegroundColor Gray
Write-Host "  • Prompt Generation" -ForegroundColor Gray
Write-Host "  • Error Handling" -ForegroundColor Gray
Write-Host ""

Read-Host "Press Enter to exit"
