# AI Novel Prompter - Project Structure

## Root Level Files
- **`main.go`** - Main Wails application entry point
- **`app.go`** - Core application logic and Wails bindings
- **`wails.json`** - Wails configuration file
- **`go.mod/go.sum`** - Go module dependencies
- **`package.json`** - Root package configuration
- **`docker-compose.yml`** - Docker setup for web version

## Key Directories

### `/frontend/` - React Frontend
- **`src/App.tsx`** - Main React application
- **`src/components/`** - UI components organized by feature
  - **`ui/`** - shadcn/ui base components
  - **`ProseImprovement/`** - Prose improvement feature
  - **`TaskTypeSelector/`** - Task management components
- **`src/hooks/`** - Custom React hooks for state management
- **`src/utils/`** - Utility functions and constants
- **`src/types.ts`** - TypeScript type definitions
- **`wailsjs/`** - Generated Wails bindings

### `/server/` - Go Backend (Web Version)
- **`cmd/main.go`** - Web server entry point
- **`pkg/`** - Reusable packages
- **`config.yaml`** - Server configuration
- **`go.mod`** - Server-specific dependencies

### `/client/` - Legacy Web Frontend
- React-based web interface (separate from desktop app)
- Uses Material-UI (Berry template)
- Database-backed features

### `/scripts/` - Build Scripts
- **`build.sh`** - General build script
- **`build-windows.sh`** - Windows-specific build
- **`build-macos*.sh`** - macOS build variants
- **`install-wails-cli.sh`** - Wails CLI installation

### Other Directories
- **`/build/`** - Build output directory
- **`/images/`** - Documentation images
- **`/doc/`** - Documentation files
- **`/prompts/`** - Example prompts and templates
- **`/finetune_data_example*/`** - Fine-tuning datasets and examples
- **`/compare/`** - AI model comparison experiments

## Data Storage
- **Desktop App**: Local files in user home directory `~/.ai-novel-prompter/`
- **Web Version**: PostgreSQL database

## Configuration Files
- **`frontend/package.json`** - Frontend dependencies
- **`frontend/tailwind.config.js`** - Tailwind CSS configuration
- **`frontend/vite.config.ts`** - Vite build configuration
- **`frontend/tsconfig.json`** - TypeScript configuration
- **`server/config.yaml`** - Backend server configuration

## Generated Files
- **`frontend/wailsjs/`** - Auto-generated Wails bindings
- **`build/`** - Compiled application binaries