# AI Novel Prompter - Code Style and Conventions

## TypeScript/JavaScript Conventions

### Naming Conventions
- **Components**: PascalCase (e.g., `TaskTypeSelector`, `ProseImprovement`)
- **Hooks**: camelCase starting with `use` (e.g., `useOptionManagement`, `useProseImprovement`)
- **Types/Interfaces**: PascalCase (e.g., `PromptType`, `ProseImprovementSession`)
- **Variables/Functions**: camelCase (e.g., `generatePrompt`, `handleSubmit`)
- **Constants**: UPPER_SNAKE_CASE (e.g., `DEFAULT_TASK_TYPES`, `TOKEN_LIMITS`)
- **Files**: kebab-case for utilities, PascalCase for components

### Component Structure
- Use functional components with hooks
- Props interfaces defined above component
- Export component as default at bottom
- Group related components in folders with index.tsx

### Type Definitions
- Centralized in `src/types.ts`
- Use interfaces over types where possible
- Properly typed React component props
- Enum types for fixed sets of values

### Import Organization
1. External libraries (React, etc.)
2. Internal components
3. Hooks and utilities
4. Types
5. Relative imports

## Go Conventions

### Project Structure
- Standard Go project layout with `cmd/` and `pkg/` directories
- Main application entry in root directory files
- Configuration in YAML files
- Environment variables with godotenv

### Naming Conventions
- **Packages**: lowercase (e.g., `main`)
- **Functions**: PascalCase for public, camelCase for private
- **Structs**: PascalCase (e.g., `App`)
- **Methods**: PascalCase for public, camelCase for private

### Error Handling
- Proper error propagation
- Logging with appropriate levels
- Graceful error responses to frontend

## File Organization

### Frontend Structure
```
src/
├── components/          # Reusable UI components
│   ├── ui/             # shadcn/ui components
│   └── [Feature]/      # Feature-specific components
├── hooks/              # Custom React hooks
├── utils/              # Utility functions
├── types.ts            # Type definitions
└── App.tsx            # Main application component
```

### Configuration Management
- Environment variables via .env files
- Configuration hierarchy: UI settings → .env → config.yaml → defaults
- Sensitive data in .env files (gitignored)

## Documentation Standards
- Clear README with setup instructions
- Inline comments for complex logic only
- JSDoc for public APIs
- Type annotations serve as documentation