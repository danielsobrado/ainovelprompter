#!/bin/bash

# Test runner script for AI Novel Prompter

set -e

echo "🧪 Running AI Novel Prompter Test Suite"
echo "======================================="

# Colors for output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
NC='\033[0m' # No Color

# Function to print colored output
print_status() {
    echo -e "${BLUE}[INFO]${NC} $1"
}

print_success() {
    echo -e "${GREEN}[SUCCESS]${NC} $1"
}

print_warning() {
    echo -e "${YELLOW}[WARNING]${NC} $1"
}

print_error() {
    echo -e "${RED}[ERROR]${NC} $1"
}

# Check if we're in the right directory
if [ ! -f "go.mod" ] || [ ! -d "frontend" ]; then
    print_error "This script must be run from the project root directory"
    exit 1
fi

# Parse command line arguments
FRONTEND_ONLY=false
BACKEND_ONLY=false
COVERAGE=false
WATCH=false
UI=false

while [[ $# -gt 0 ]]; do
    case $1 in
        --frontend-only)
            FRONTEND_ONLY=true
            shift
            ;;
        --backend-only)
            BACKEND_ONLY=true
            shift
            ;;
        --coverage)
            COVERAGE=true
            shift
            ;;
        --watch)
            WATCH=true
            shift
            ;;
        --ui)
            UI=true
            shift
            ;;
        --help)
            echo "Usage: $0 [OPTIONS]"
            echo "Options:"
            echo "  --frontend-only    Run only frontend tests"
            echo "  --backend-only     Run only backend tests"
            echo "  --coverage         Generate coverage reports"
            echo "  --watch           Run tests in watch mode"
            echo "  --ui              Run frontend tests with UI"
            echo "  --help            Show this help message"
            exit 0
            ;;
        *)
            print_error "Unknown option: $1"
            echo "Use --help for usage information"
            exit 1
            ;;
    esac
done

# Install dependencies if needed
install_dependencies() {
    print_status "Checking dependencies..."
    
    # Check Go dependencies
    if [ "$FRONTEND_ONLY" = false ]; then
        print_status "Installing Go dependencies..."
        go mod download
        go mod tidy
    fi
    
    # Check Node dependencies
    if [ "$BACKEND_ONLY" = false ]; then
        if [ ! -d "frontend/node_modules" ]; then
            print_status "Installing Node.js dependencies..."
            cd frontend
            npm install
            cd ..
        fi
    fi
}

# Run Go tests
run_go_tests() {
    print_status "Running Go tests..."
    
    if [ "$COVERAGE" = true ]; then
        print_status "Running Go tests with coverage..."
        go test -v -race -coverprofile=coverage.out -covermode=atomic ./...
        go tool cover -html=coverage.out -o coverage.html
        print_success "Go coverage report generated: coverage.html"
    else
        go test -v -race ./...
    fi
    
    if [ $? -eq 0 ]; then
        print_success "Go tests passed!"
    else
        print_error "Go tests failed!"
        return 1
    fi
}

# Run frontend tests
run_frontend_tests() {
    print_status "Running frontend tests..."
    
    cd frontend
    
    if [ "$UI" = true ]; then
        print_status "Starting test UI..."
        npm run test:ui
    elif [ "$WATCH" = true ]; then
        print_status "Starting test watcher..."
        npm run test:watch
    elif [ "$COVERAGE" = true ]; then
        print_status "Running frontend tests with coverage..."
        npm run test:coverage
        print_success "Frontend coverage report generated in coverage/ directory"
    else
        npm run test:run
    fi
    
    local exit_code=$?
    cd ..
    
    if [ $exit_code -eq 0 ]; then
        print_success "Frontend tests passed!"
    else
        print_error "Frontend tests failed!"
        return 1
    fi
}

# Run linting
run_linting() {
    print_status "Running linting..."
    
    # Go linting
    if [ "$FRONTEND_ONLY" = false ]; then
        print_status "Running Go linting..."
        
        # Format check
        if ! gofmt -l . | grep -q .; then
            print_success "Go code is properly formatted"
        else
            print_warning "Go code formatting issues found:"
            gofmt -l .
            print_status "Running gofmt..."
            gofmt -w .
        fi
        
        # Vet check
        go vet ./...
        if [ $? -eq 0 ]; then
            print_success "Go vet passed!"
        else
            print_error "Go vet failed!"
            return 1
        fi
    fi
    
    # Frontend linting
    if [ "$BACKEND_ONLY" = false ]; then
        print_status "Running frontend linting..."
        cd frontend
        npm run lint
        local exit_code=$?
        cd ..
        
        if [ $exit_code -eq 0 ]; then
            print_success "Frontend linting passed!"
        else
            print_error "Frontend linting failed!"
            return 1
        fi
    fi
}

# Main execution
main() {
    print_status "Starting test suite..."
    
    # Install dependencies
    install_dependencies
    
    # Run linting
    if [ "$WATCH" = false ] && [ "$UI" = false ]; then
        run_linting
        if [ $? -ne 0 ]; then
            print_error "Linting failed. Stopping test execution."
            exit 1
        fi
    fi
    
    # Run tests based on options
    if [ "$FRONTEND_ONLY" = true ]; then
        run_frontend_tests
    elif [ "$BACKEND_ONLY" = true ]; then
        run_go_tests
    else
        # Run both
        run_go_tests
        if [ $? -eq 0 ]; then
            run_frontend_tests
        else
            print_error "Backend tests failed. Skipping frontend tests."
            exit 1
        fi
    fi
    
    if [ $? -eq 0 ]; then
        echo ""
        print_success "All tests passed! 🎉"
        echo ""
        
        if [ "$COVERAGE" = true ]; then
            print_status "Coverage reports generated:"
            if [ "$FRONTEND_ONLY" = false ]; then
                echo "  - Go: coverage.html"
            fi
            if [ "$BACKEND_ONLY" = false ]; then
                echo "  - Frontend: frontend/coverage/"
            fi
        fi
    else
        print_error "Some tests failed! ❌"
        exit 1
    fi
}

# Trap to clean up on exit
cleanup() {
    print_status "Cleaning up..."
}

trap cleanup EXIT

# Run main function
main
