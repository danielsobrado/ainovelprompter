# AI Novel Prompter - Suggested Commands

## Development Commands

### Frontend Development
```bash
# Navigate to frontend directory
cd frontend

# Install dependencies
npm install

# Start development server
npm run dev

# Build for production
npm run build

# Run linting
npm run lint

# Preview production build
npm run preview
```

### Wails Desktop App Commands
```bash
# Development mode (hot reload)
wails dev

# Build for current platform
wails build

# Build with cleanup
wails build --clean

# Build for specific platform
wails build --platform windows/amd64
wails build --platform darwin/arm64
wails build --platform darwin/amd64
wails build --platform linux/amd64

# Generate Wails bindings (after Go backend changes)
wails generate bindings

# Initialize new Wails project (for reference)
wails init -n myapp -t vanilla

# Check Wails doctor
wails doctor
```

### Build Scripts (Windows)
```bash
# Use provided build scripts
./scripts/build.sh              # Current platform
./scripts/build-windows.sh      # Windows AMD64
./scripts/build-macos.sh         # macOS universal
./scripts/build-macos-arm.sh     # macOS ARM64
./scripts/build-macos-intel.sh   # macOS Intel

# Install Wails CLI
./scripts/install-wails-cli.sh
```

### Go Backend Commands
```bash
# Run Go backend directly (for testing)
cd server
go run cmd/main.go

# Install Go dependencies
go mod download
go mod tidy

# Run database migrations (web version)
go run cmd/main.go migrate

# Test Go code
go test ./...

# Format Go code
go fmt ./...

# Vet Go code
go vet ./...
```

### Docker Commands (Web Version)
```bash
# Start with Docker Compose
docker-compose up -d

# Build containers
docker-compose build

# View logs
docker-compose logs

# Stop containers
docker-compose down
```

### Git Commands
```bash
# Standard workflow
git add .
git commit -m "feat: description"
git push origin main

# Create feature branch
git checkout -b feature/feature-name
git push -u origin feature/feature-name
```

### Windows System Commands
```cmd
# List files
dir

# Change directory
cd path\to\directory

# Find files
where filename

# Environment variables
set VARIABLE=value
echo %VARIABLE%

# Process management
tasklist | findstr process_name
taskkill /PID process_id
```

### Installation Commands
```bash
# Install Node.js dependencies
npm install

# Install Go (if needed)
# Download from https://golang.org/dl/

# Install Wails CLI
go install github.com/wailsapp/wails/v2/cmd/wails@latest

# Verify installations
node --version
npm --version
go version
wails version
```