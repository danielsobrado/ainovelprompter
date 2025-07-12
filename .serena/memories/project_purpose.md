# AI Novel Prompter - Project Purpose

AI Novel Prompter is a desktop application designed to help writers create consistent and well-structured prompts for AI writing assistants like ChatGPT and Claude. 

## Core Features

### Prompt Generation
- Generate customized writing prompts for continuing novels
- Support for both ChatGPT and Claude prompt formats
- Manage story elements (characters, locations, rules, codex)
- Task and chapter management with story beats planning

### Prose Improvement
- Iterative text refinement using AI-powered improvement prompts
- Support for multiple LLM providers (Manual, LM Studio, OpenRouter)
- Change review system with accept/reject workflow
- Real-time diff viewer showing text improvements

### Data Management
- Local data persistence in user's home directory
- JSON-based storage for all user data
- Configuration management with priority system (.env → config.yaml → UI settings)
- File operations and project management

## Architecture
- **Frontend**: React + TypeScript + Tailwind CSS + shadcn/ui
- **Backend**: Go + Wails framework for desktop app
- **Additional**: Web-based version with PostgreSQL backend (separate feature)

## Target Users
Writers and authors who want to:
- Create consistent AI prompts for novel writing
- Manage complex story elements efficiently  
- Improve their prose using AI assistance
- Maintain writing consistency across chapters