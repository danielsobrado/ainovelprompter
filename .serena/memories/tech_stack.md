# AI Novel Prompter - Tech Stack

## Frontend
- **React 18.2** - Main UI framework
- **TypeScript 5.0+** - Type safety and development experience
- **Tailwind CSS 3.1+** - Utility-first styling
- **shadcn/ui** - Modern component library built on Radix UI
- **Vite 4.3+** - Build tool and development server

### Key Frontend Dependencies
- **@radix-ui/react-*** - Headless UI components (tabs, dialog, select, etc.)
- **lucide-react** - Icon library
- **class-variance-authority** - Component variant management
- **clsx** - Conditional class names
- **react-diff-viewer** - Text diff visualization for prose improvement

## Backend
- **Go 1.24.3** - Primary backend language
- **Wails v2.10.1** - Go + frontend desktop framework
- **Viper** - Configuration management
- **godotenv** - Environment variable loading

### Backend Dependencies
- **github.com/spf13/viper** - Configuration management
- **github.com/joho/godotenv** - .env file support

## Desktop Framework
- **Wails** - Cross-platform desktop app framework
  - Provides Go backend with web frontend
  - Native OS integration
  - File system access
  - Window management

## Additional Features (Web Version)
- **Gin** - HTTP web framework
- **GORM** - ORM for database operations  
- **PostgreSQL** - Database for web version
- **JWT** - Authentication

## Development Tools
- **ESLint** - JavaScript/TypeScript linting
- **PostCSS** - CSS processing
- **Autoprefixer** - CSS vendor prefixes
- **TypeScript Compiler** - Type checking