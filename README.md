# AI Novel Prompter

AI Novel Prompter can generate writing prompts for novels based on user-specified characteristics.

# Wails based Tool

AI Novel Prompter is a desktop application designed to help writers create consistent and well-structured prompts for AI writing assistants like ChatGPT and Claude, and also to refine existing prose using LLMs. The tool helps manage story elements, character details, generate prompts for continuing your novel, and iteratively improve text through AI-powered suggestions.

The Executable is on build/bin [Executable](https://github.com/danielsobrado/ainovelprompter/blob/main/build/bin/AINovelPrompter_0.0.1.exe)

## Features

### 1. Task & Chapter Management (Prompt Generation Tab)
- **Task Type Selection**: Define and customize different types of writing tasks (e.g., "Write Next Chapter", "Revise Chapter").
- **Sample Chapter Management**: Store and reference sample chapters for style consistency.
- **Chapter Content Tabs**:
  - Story Beats: Plan the main points for your next chapter.
  - Previous Chapter: Reference the last written chapter.
  - Future Notes: Keep track of planned future developments.

![AI Novel 1](https://github.com/danielsobrado/ainovelprompter/blob/main/images/AINovel1.jpg)

### 2. Story Element Management (Prompt Generation Tab)
Each category can be edited, saved, and reused across different prompts:
- **Rules**: Define writing rules and style guidelines.
- **Characters**: Manage character profiles and details.
- **Locations**: Keep track of story locations and their descriptions.
- **Codex**: Store world-building elements and lore.

### 3. Prompt Generation (Prompt Generation Tab)
- **Dual AI Support**:
  - ChatGPT-optimized formatting.
  - Claude-optimized XML formatting.
- **Real-time Preview**: See your formatted prompt as you build it.
- **Token Counting**: Track token usage for AI model limits.
- **Custom Instructions**: Add specific requirements or guidelines to the main prompt.

### 4. Prose Improvement (Prose Improvement Tab)
- **Iterative Text Refinement**: Paste your text and apply a series of AI-powered improvement prompts.
- **Customizable Improvement Prompts**:
    - Manage a list of prompts (e.g., "Enhance Imagery," "Strengthen Verbs," "Check for Clichés," "Grammar and Punctuation").
    - Add, edit, delete, and reorder these improvement prompts.
    - Prompts include examples of the expected JSON output format to guide the LLM.
- **LLM Provider Integration**:
    - Supports Manual (copy/paste), LM Studio, and OpenRouter.
    - Configure API URLs, API keys, and model identifiers for each provider.
    - **Flexible Configuration**: Provider settings (API keys, models) can be loaded from a `.env` file (recommended for sensitive data), `server/config.yaml`, or overridden and saved via the UI to a local `llm_provider_settings.json` file.
- **Step-by-Step Processing**: Execute improvement prompts one by one against your text.
- **Change Review System**:
    - AI suggestions are parsed from JSON responses.
    - Review each suggested change (`initial` vs. `improved` text, and the `reason` for the change).
    - Accept or reject suggestions.
    - View context for each change within the original text.
- **Live Preview of Accepted Changes**:
    - A diff viewer (`react-diff-viewer`) shows the `originalText` vs. the `currentText` (with all accepted changes applied).
- **Robust JSON Parsing**: The system attempts to extract and parse JSON even if the LLM includes conversational text around the JSON payload.

### 5. Data Persistence & Versioning

**📁 Folder-Based Storage with Complete Version History**
- Revolutionary **folder-based system** replacing single JSON files
- **Complete version history** with timestamped versioning for every entity
- **Never lose data** - every create, update, and delete operation is preserved
- **One-click restore** to any previous version of any entity
- **Multiple data directories** support for different projects/stories

**📂 Storage Structure:**
```
.ai-novel-prompter/
├── characters/
│   ├── aragorn_20250113_120000_create.json
│   ├── aragorn_20250113_130000_update.json
│   └── gandalf_20250113_140000_create.json
├── locations/
│   ├── rivendell_20250113_120000_create.json
│   └── mordor_20250113_130000_create.json
├── codex/
├── rules/
├── chapters/
├── story_beats/
├── future_notes/
├── sample_chapters/
├── task_types/
├── prose_prompts/
└── .metadata/
    ├── config.json
    └── indexes.json
```

**🔄 Version Management Features:**
- **Complete Audit Trail**: Every change tracked with timestamp and operation type
- **Intelligent File Naming**: `{entity_name}_{YYYYMMDD_HHMMSS}_{operation}.json` format
- **Instant Restore**: Restore any entity to any previous version with one click
- **Storage Analytics**: Detailed statistics on files, versions, and storage usage
- **Retention Policies**: Configurable cleanup of old versions to manage disk space
- **Atomic Operations**: All file operations are atomic and safe from corruption

**💾 Automatic Migration Support:**
- **Seamless Upgrade**: Automatic detection and migration from old JSON format
- **Safe Migration**: Backup creation before migration ensures data safety
- **Validation**: Complete verification that all data migrated successfully
- **Zero Downtime**: All existing functionality preserved during transition

### 6. Command Line Interface

**🚀 Enhanced CLI with Data Directory Management:**
```bash
# Use default data directory (~/.ai-novel-prompter)
./ainovelprompter

# Specify custom data directory for project isolation
./ainovelprompter --data-dir /path/to/project
./ainovelprompter -d "./Fantasy Novel"

# Multiple projects with different directories
./ainovelprompter -d "./Project A"
./ainovelprompter -d "./Project B" 

# Show help and available options
./ainovelprompter --help
./ainovelprompter -h
```

**📁 Data Directory Benefits:**
- **Project Isolation**: Keep different stories completely separate
- **Easy Backup**: Copy entire project folder for instant backup
- **Team Collaboration**: Share project folders via cloud storage
- **Version Control**: Track entire project history with git
- **Portable Projects**: Move projects between computers easily

### 7. User Interface
- **Clean, Modern Design**: Built with shadcn/ui components.
- **Tabbed Interface**: Organized access to "Prompt Generation" and "Prose Improvement" features.
- **Modal Editors**: Easy editing of all manageable data types.
- **Version History UI**: Timeline view of all entity versions with visual operation indicators.
- **Data Directory Manager**: Frontend interface for switching between project directories.

## Technical Stack

- **Frontend**:
  - React
  - TypeScript
  - Tailwind CSS
  - shadcn/ui components
  - `react-diff-viewer` for visualizing text changes.
- **Backend**:
  - Go
  - Wails framework
  - Atomic file operations for data integrity
  - JSON-based storage with versioning
  - `github.com/joho/godotenv` for `.env` file support.
  - `github.com/spf13/viper` for configuration management.

## File Management & Configuration

- **User Data**: Saved in configurable data directories with complete version history
  - **Default**: `~/.ai-novel-prompter` or `C:\Users\YourUser\.ai-novel-prompter`
  - **Custom**: Any directory specified via CLI or frontend settings
- **Version Management**: Each entity maintains complete history with atomic file operations
- **LLM Provider Configuration Priority**:
    1. Settings saved via the UI (`llm_provider_settings.json`).
    2. Values from a `.env` file located in the project root (for `wails dev`) or next to the executable (for built app). Environment variables should be prefixed (e.g., `APP_OPENROUTER_API_KEY`).
    3. Values from `server/config.yaml`.
    4. Hardcoded defaults in the application.
- **`.env` File**: Create a `.env` file in your project root for development to store API keys and default models (e.g., `APP_OPENROUTER_API_KEY="your_key"`). **Add `.env` to your `.gitignore` file.**

## Installation

```bash
# Clone the repository
git clone [repository-url]
cd ainovelprompter

# Install Go dependencies for the backend (if not already handled by Wails)
# cd server
# go mod tidy 
# cd .. 

# Install frontend dependencies
cd frontend
npm install
cd ..

# Generate Wails bindings (important after Go backend changes)
wails generate bindings

# Build and run the application in development mode
wails dev

# Or run with custom data directory
wails dev --data-dir ./my-project-data
```

### Command Line Options

The application now supports CLI options for enhanced project management:

```bash
# Use default data directory (~/.ai-novel-prompter)
./ainovelprompter

# Specify custom data directory
./ainovelprompter --data-dir /path/to/project
./ainovelprompter -d "./Fantasy Novel"

# Multiple projects with different directories
./ainovelprompter -d "./Project A"
./ainovelprompter -d "./Project B"

# Show help
./ainovelprompter --help
./ainovelprompter -h
```

**Data Directory Benefits:**
- **Project Isolation**: Keep different stories completely separate
- **Easy Backup**: Copy entire project folder for backup
- **Team Collaboration**: Share project folders via cloud storage
- **Version Control**: Track entire project history with git

## Building for Production

To build a redistributable, production mode package, use `wails build`.

```bash
wails build
```

The Executable is on build/bin [Executable](https://github.com/danielsobrado/ainovelprompter/blob/main/build/bin/AINovelPrompter_0.0.1.exe)

Or generate it with:

```console
wails build -nsis
```
This can be done for Mac as well see the latest part of [this guide](https://wails.io/docs/guides/windows-installer/)

The built application will be available in the `build` directory.

## Usage Guide

### A. Prompt Generation 
1.  **Initial Setup**:
    *   In the "Prompt Generation" tab, define your task types (e.g., "Write Next Chapter," "Revise Chapter") by clicking the edit icon next to the "Task Type" selector.
    *   Add sample chapters for style reference via the "Sample Chapters" selector and its edit icon.
    *   Set up your "Rules," "Characters," "Locations," and "Codex" entries using their respective selectors and edit icons. All these are saved locally with complete version history.
2.  **Creating a Prompt**:
    *   Select your desired "Task Type."
    *   Input content into the "Story Beats," "Previous Chapter," and "Future Notes" tabs.
    *   Select the relevant "Rules," "Characters," "Locations," and "Codex" entries that apply to the chapter you're planning.
    *   Add any specific "Custom Instructions" in the text area provided.
3.  **Generating Output**:
    *   Choose between "ChatGPT" or "Claude" optimized prompt formats using the buttons in the "Generated Prompt" section.
    *   Review the dynamically generated prompt in the preview area.
    *   Use the "Copy to Clipboard" button and paste the prompt into your preferred AI writing assistant.

### B. Prose Improvement 
1.  **Configure LLM Provider**:
    *   Navigate to the "Prose Improvement" tab.
    *   Click the "Provider Settings" button (cog icon).
    *   In the modal:
        *   Select your LLM provider: Manual (for copy/pasting), LM Studio, or OpenRouter.
        *   **LM Studio**: Enter the API URL (e.g., `http://localhost:1234/v1/chat/completions`) and the model identifier loaded in your LM Studio instance.
        *   **OpenRouter**: Enter your OpenRouter API Key and the desired model identifier (e.g., `anthropic/claude-3-haiku`, `openai/gpt-4o`). This is a free-text field.
    *   These settings are loaded with a priority:
        1.  Previously saved settings from the UI (stored in `llm_provider_settings.json`).
        2.  Settings from your project's `.env` file (e.g., `APP_OPENROUTER_API_KEY`).
        3.  Settings from `server/config.yaml`.
        4.  Application defaults.
    *   Changes made here are saved to `llm_provider_settings.json` for future sessions.
2.  **Manage Improvement Prompts**:
    *   Within the "Prose Improvement" tab, go to the "Prompts" sub-tab.
    *   Review the default prompts (e.g., "Enhance Imagery," "Strengthen Verbs," "Check for Clichés," "Grammar and Punctuation").
    *   You can add your own custom improvement prompts, edit existing ones (including their label, prompt text, category, and order), or delete them.
    *   **Important**: For prompts expected to return structured data (like lists of changes), ensure the prompt text clearly instructs the LLM on the desired JSON format and **includes an example of the JSON structure with specific keys** (e.g., `"initial_text"`, `"improved_text"`, `"reason"`). This greatly improves parsing reliability.
    *   These prompts are saved locally with full version history.
3.  **Process Text**:
    *   Go to the "Input Text" sub-tab, paste the text you want to improve, and click "Start Improvement Session."
    *   Navigate to the "Process" sub-tab.
    *   The first prompt from your managed list will be displayed.
    *   Click "Execute Prompt."
        *   If using LM Studio or OpenRouter, the application will send the prompt and text to the LLM.
        *   If using Manual mode, the full prompt (including your text) will be copied to your clipboard. Paste this into your chosen LLM, get the response, then paste the LLM's JSON response back into the "Paste AI response here..." text area in the app and click "Process Response."
    *   The application will attempt to parse the LLM's JSON response.
4.  **Review Changes**:
    *   After processing, navigate to the "Review Changes" sub-tab.
    *   Suggested changes from the LLM will be listed. Each card will show:
        *   The original text segment (`initial`).
        *   The LLM's suggested improvement (`improved`).
        *   The LLM's `reason` for the change.
        *   Context snippets from the original text (if available and enabled).
    *   You can "Accept" or "Reject" each suggestion.
    *   The "Current Text Preview" at the bottom uses `react-diff-viewer` to show a live diff between your original input text for the session and the current state of the text with all accepted changes applied.
    *   Once you've reviewed changes from one prompt, you can go back to the "Process" tab to execute the next prompt in your list against the updated text.

## Data Directory Management & Version Control

### Managing Data Directories

**🏠 Default Location**: `~/.ai-novel-prompter` (Windows: `C:\Users\{username}\.ai-novel-prompter`)

**📁 Custom Directories**: Use CLI options or frontend settings to work with multiple projects:

```bash
# Different story projects
./ainovelprompter -d "./Fantasy Novel"
./ainovelprompter -d "./Sci-Fi Story" 
./ainovelprompter -d "./Historical Fiction"
```

**🔧 Frontend Management**:
1. Open **Settings** → **Data Directory Manager**
2. **Browse** for new directory or **select from recent**
3. **Validate** directory before switching
4. **View storage statistics** and entity counts
5. **Migrate** from old JSON format if needed

### Version History & Restore

**📖 Viewing History**:
- Every entity maintains complete version history with timestamps
- Timeline shows all create/update/delete operations with visual indicators
- Click **History** icon next to any entity to view versions
- Preview any version before restoring

**⏪ Restoring Versions**:
1. Navigate to entity's version history
2. Select desired version from timeline
3. **Preview** version content to verify
4. Click **Restore** to create new version from selected point
5. Current version becomes new historical entry
6. **Atomic operation** - safe from corruption or data loss

**🧹 Cleanup & Maintenance**:
- **Automatic cleanup** removes old versions based on configurable retention policy
- **Manual cleanup** via storage statistics interface  
- **Backup entire directory** before major changes
- **Storage analytics** show disk usage by entity type and version count

### Migration from Legacy Format

**🔄 Automatic Detection**: App detects old JSON files and offers migration wizard

**📋 Migration Process**:
1. **Backup Creation**: Automatic backup of old format with timestamp
2. **Entity Conversion**: Each JSON entry becomes properly versioned entity
3. **Timestamp Assignment**: Proper creation/update timestamps assigned
4. **Validation**: Complete verification that all data migrated successfully
5. **Cleanup**: Option to archive old JSON files post-migration

**⚠️ Migration Notes**:
- **One-time process** per data directory
- **Irreversible** (but backup is created automatically)
- **Maintains all data** including relationships and metadata
- **Preserves settings** and configurations
- **Zero data loss** with complete validation

### Project Organization Best Practices

**📂 Directory Structure**:
```
Stories/
├── Fantasy-Series/
│   ├── Book-1/
│   └── Book-2/
├── Standalone-Novel/
└── Short-Stories/
```

**💾 Backup Strategy**:
- **Cloud sync** entire project directories for automatic backup
- **Git repositories** for version control across devices and team collaboration
- **Scheduled backups** of active projects to external storage
- **Export functionality** for sharing specific entities or projects

**👥 Team Collaboration**:
- **Share project directories** via cloud storage (Dropbox, Google Drive, etc.)
- **Version conflicts** resolved automatically through timestamp comparison
- **Entity-level granularity** minimizes merge conflicts between team members
- **Complete audit trail** shows who changed what and when

## Development

### Adding New Features
- **Prompt Generation**: The codebase supports easy addition of new selectors (for managing different types of story elements) and options. Modal components for editing these elements follow a consistent pattern. Data persistence is handled via versioned folder storage with atomic operations.
- **Prose Improvement**:
    - New default improvement prompts can be added to the `DEFAULT_PROSE_IMPROVEMENT_PROMPTS` array in `src/utils/constants.ts`. Remember to include clear JSON output examples in the prompt text.
    - The JSON parser in `src/hooks/useProseImprovement.ts` is designed to be somewhat flexible with key names from LLM responses by checking for common variations (e.g., `original`, `initial`, `original_text`). However, it benefits greatly from well-defined prompts that specify exact key names.
    - Support for new LLM providers can be added by extending the `LLMProvider` type in `src/types.ts` and updating the logic in `src/hooks/useLLMProvider.ts` and `src/components/ProseImprovement/ProviderSettings.tsx`.
- **Data Persistence**: All data operations use the versioned folder storage system. New entity types can be added by extending the storage interfaces and implementing the corresponding handlers. All operations are atomic and maintain complete version history.

### Testing

**📋 Comprehensive Test Suite**:
```bash
# Run all tests
./scripts/test.sh

# Frontend tests only
./scripts/test.sh --frontend-only

# Backend tests only  
./scripts/test.sh --backend-only

# With coverage reports
./scripts/test.sh --coverage

# Test storage system specifically
cd cmd/test-storage
go run main.go
```

**🧪 Storage System Tests**:
- **Version Management**: Create, update, restore operations
- **Migration Testing**: Old JSON to new folder format
- **Atomic Operations**: Data integrity under concurrent access
- **Performance**: Large dataset handling and cleanup
- **Legacy Compatibility**: All existing MCP tools work unchanged

### Customization
- **Styling**: All components use Tailwind CSS. Customizations can be made by modifying utility classes or adding custom CSS to `src/index.css`.
- **UI Components**: Built with shadcn/ui. These components can be customized as per shadcn/ui documentation.
- **Prompt Formatting (Main Tab)**: The logic for constructing prompts for ChatGPT and Claude in the "Prompt Generation" tab can be modified in `src/utils/promptGenerators.ts` and `src/utils/promptInstructions.ts`.
- **Prose Improvement Prompts**: These are primarily managed via the UI with complete version history. Default prompts are in `src/utils/constants.ts`.
- **Storage System**: The folder-based storage can be extended for new entity types by implementing the versioned storage interfaces.

---

# Web based tool

## Features

- User registration and authentication
- Text creation and management
- Chapter creation and management
- Feedback submission and management
- Prompt generation based on traits
- Integration with a local ollama service
- Based on Berry template (https://codedthemes.gitbook.io/berry)
- Inspired on Jason Hamilton Youtube (https://www.youtube.com/@TheNerdyNovelist)

## Technologies Used

- Frontend:
  - React
  - TypeScript
  - Axios
  - React Router
  - React Toastify
- Backend:
  - Go
  - Gin Web Framework
  - GORM (Go ORM)
  - PostgreSQL

## Prerequisites

Before running the application, make sure you have the following installed:

- Node.js (v18 or higher)
- Go (v1.18 or higher)
- PostgreSQL
- Docker
- Docker Compose

## Getting Started

1. Clone the repository:
   ```
   git clone https://github.com/danielsobrado/ainovelprompter.git
   ```
2. Navigate to the project directory:
   ```
   cd ainovelprompter
   ```
3. Set up the backend:

- Navigate to the `server` directory:

  ```
  cd server
  ```

- Install the Go dependencies:

  ```
  go mod download
  ```

- Update the `config.yaml` file with your database configuration.

- Run the database migrations:

  ```
  go run cmd/main.go migrate
  ```

- Start the backend server:

  ```
  go run cmd/main.go
  ```

4. Set up the frontend:

- Navigate to the `client` directory:

  ```
  cd ../client
  ```

- Install the frontend dependencies:

  ```
  npm install
  ```

- Start the frontend development server:
  ```
  npm start
  ```
5. Open your web browser and visit `http://localhost:3000` to access the application.

## Getting Started (Docker)

1. Clone the repository:
```
git clone https://github.com/danielsobrado/ainovelprompter.git
```

2. Navigate to the project directory:
```
cd ainovelprompter
```

3. Update the `docker-compose.yml` file with your database configuration.

4. Start the application using Docker Compose:
```
docker-compose up -d
```

5. Open your web browser and visit `http://localhost:3000` to access the application.

## Configuration

- Backend configuration can be modified in the `server/config.yaml` file.
- Frontend configuration can be modified in the `client/src/config.ts` file.

## Build

To build the frontend for production, run the following command in the `client` directory:
   ```
   npm run build
   ```
The production-ready files will be generated in the `client/build` directory.

## Installation and Management Guide for PostgreSQL on WSL

This small guide provides instructions on how to install PostgreSQL on the Windows Subsystem for Linux (WSL), along with steps to manage user permissions and troubleshoot common issues.

---

### Prerequisites

- Windows 10 or higher with WSL enabled. (Or just Ubuntu)
- Basic familiarity with Linux command line and SQL.

---

### Installation

1. **Open WSL Terminal**: Launch your WSL distribution (Ubuntu recommended).

2. **Update Packages**:
   ```bash
   sudo apt update
   ```

3. **Install PostgreSQL**:
   ```bash
   sudo apt install postgresql postgresql-contrib
   ```

4. **Check Installation**:
   ```bash
   psql --version
   ```

5. **Set PostgreSQL User Password**:
   ```bash
   sudo passwd postgres
   ```

---

### Database Operations

1. **Create Database**:
   ```bash
   createdb mydb
   ```

2. **Access Database**:
   ```bash
   psql mydb
   ```

3. **Import Tables from SQL File**:
   ```bash
   psql -U postgres -q mydb < /path/to/file.sql
   ```

4. **List Databases and Tables**:
   ```sql
   \l  # List databases
   \dt # List tables in the current database
   ```

5. **Switch Database**:
   ```sql
   \c dbname
   ```

---

### User Management

1. **Create New User**:
   ```sql
   CREATE USER your_db_user WITH PASSWORD 'your_db_password';
   ```

2. **Grant Privileges**:
   ```sql
   ALTER USER your_db_user CREATEDB;
   ```

---

### Troubleshooting

1. **Role Does Not Exist Error**:
   Switch to the 'postgres' user:
   ```bash
   sudo -i -u postgres
   createdb your_db_name
   ```

2. **Permission Denied to Create Extension**:
   Login as 'postgres' and execute:
   ```sql
   CREATE EXTENSION IF NOT EXISTS pg_trgm;
   ```

3. **Unknown User Error**:
   Ensure you are using a recognized system user or correctly refer to a PostgreSQL user within the SQL environment, not via `sudo`.

---

## Generating Custom Training Data to Fine-Tune a Language Model (Manual Steps)

To [generate custom training data](https://github.com/danielsobrado/ainovelprompter/tree/main/finetune_data_example1) for fine-tuning a language model to emulate the writing style of George MacDonald, the process begins by obtaining the full text of one of his novels, "The Princess and the Goblin," from Project Gutenberg. The text is then broken down into individual story beats or key moments using a prompt that instructs the AI to generate a JSON object for each beat, capturing the author, emotional tone, type of writing, and the actual text excerpt.

Next, GPT-4 is used to rewrite each of these story beats in its own words, generating a parallel set of JSON data with unique identifiers linking each rewritten beat to its original counterpart. To simplify the data and make it more useful for training, the wide variety of emotional tones is mapped to a smaller set of core tones using a Python function. The two JSON files (original and rewritten beats) are then used to generate training prompts, where the model is asked to rephrase the GPT-4 generated text in the style of the original author. Finally, these prompts and their target outputs are formatted into JSONL and JSON files, ready to be used for fine-tuning the language model to capture MacDonald's distinctive writing style.

---

## Generating Custom Training Data to Fine-Tune a Language Model (Automated)

In the previous example, the process of generating paraphrased text using a language model involved some manual tasks. The user had to manually provide the input text, run the script, and then review the generated output to ensure its quality. If the output did not meet the desired criteria, the user would need to manually retry the generation process with different parameters or make adjustments to the input text.

However, with the [updated version](https://github.com/danielsobrado/ainovelprompter/tree/main/finetune_data_example2) of the `process_text_file` function, the entire process has been fully automated. The function takes care of reading the input text file, splitting it into paragraphs, and automatically sending each paragraph to the language model for paraphrasing. It incorporates various checks and retry mechanisms to handle cases where the generated output does not meet the specified criteria, such as containing unwanted phrases, being too short or too long, or consisting of multiple paragraphs.

The automation process includes several key features:

1. **Resuming from the last processed paragraph:**
   If the script is interrupted or needs to be run multiple times, it automatically checks the output file and resumes processing from the last successfully paraphrased paragraph. This ensures that progress is not lost and the script can pick up where it left off.

2. **Retry mechanism with random seed and temperature:**
   If a generated paraphrase fails to meet the specified criteria, the script automatically retries the generation process up to a specified number of times. With each retry, it randomly changes the seed and temperature values to introduce variation in the generated responses, increasing the chances of obtaining a satisfactory output.

3. **Progress saving:**
   The script saves the progress to the output file every specified number of paragraphs (e.g., every 500 paragraphs). This safeguards against data loss in case of any interruptions or errors during the processing of a large text file.

4. **Detailed logging and summary:**
   The script provides detailed logging information, including the input paragraph, generated output, retry attempts, and reasons for failure. It also generates a summary at the end, displaying the total number of paragraphs, successfully paraphrased paragraphs, skipped paragraphs, and the total number of retries.

---
## Generating Custom Training Data to Fine-Tune a Language Model with Local LLM and LM Studio using ORPO

To [generate ORPO custom training data](https://github.com/danielsobrado/ainovelprompter/tree/main/finetune_data_example3_orpo) for fine-tuning a language model to emulate the writing style of George MacDonald. 

The input data should be in JSONL format, with each line containing a JSON object that includes the prompt and chosen response. (From the previous fine tuning)
To use the script, you need to set up the OpenAI client with your API key and specify the input and output file paths. Running the script will process the JSONL file and generate a CSV file with columns for the prompt, chosen response, and a generated rejected response. The script saves progress every 100 lines and can resume from where it left off if interrupted. Upon completion, it provides a summary of the total lines processed, written lines, skipped lines, and retry details.

---
## Fine-Tuning lessons

* Dataset Quality Matters: 95% of outcomes depend on dataset quality. A clean dataset is essential since even a little bad data can hurt the model.

* Manual Data Review: Cleaning and evaluating the dataset can greatly improve the model. This is a time-consuming but necessary step because no amount of parameter adjusting can fix a defective dataset.

* Training parameters should not improve but prevent model degradation. In robust datasets, the goal should be to avoid negative repercussions while directing the model. There is no optimal learning rate.

* Model Scale and Hardware Limitations: Larger models (33b parameters) may enable better fine-tuning but require at least 48GB VRAM, making them impractical for majority of home setups.

* Gradient Accumulation and Batch Size: Gradient accumulation helps reduce overfitting by enhancing generalisation across different datasets, but it may lower quality after a few batches.

* The size of the dataset is more important for fine-tuning a base model than a well-tuned model. Overloading a well-tuned model with excessive data might degrade its previous fine-tuning.

* An ideal learning rate schedule starts with a warmup phase, holds steady for an epoch, and then gradually decreases using a cosine schedule.

* Model Rank and Generalisation: The amount of trainable parameters affects the model's detail and generalisation. Lower-rank models generalise better but lose detail.

* LoRA's Applicability: Parameter-Efficient Fine-Tuning (PEFT) is applicable to large language models (LLMs) and systems like Stable Diffusion (SD), demonstrating its versatility.

---

## Finetuning Llama 3 issues as of May 2024

The Unsloth community has helped resolve several issues with finetuning Llama3. Here are some key points to keep in mind:

1. **Double BOS tokens**: Double BOS tokens during finetuning can break things. Unsloth automatically fixes this issue. 

2. **GGUF conversion**: GGUF conversion is broken. Be careful of double BOS and use CPU instead of GPU for conversion. Unsloth has built-in automatic GGUF conversions.

3. **Buggy base weights**: Some of Llama 3's base (not instruct) weights are "buggy" (untrained): `<|reserved_special_token_{0->250}|> <|eot_id|> <|start_header_id|> <|end_header_id|>`. This can cause NaNs and buggy results. Unsloth automatically fixes this.

4. **System prompt**: According to the Unsloth community, adding a system prompt makes finetuning of the Instruct version (and possibly the base version) much better.

5. **Quantization issues**: Quantization issues are common. See [this comparison](https://github.com/matt-c1/llama-3-quant-comparison) which shows that you can get good performance with Llama3, but using the wrong quantization can hurt performance. For finetuning, use bitsandbytes nf4 to boost accuracy. For GGUF, use the I versions as much as possible.

6. **Long context models**: Long context models are poorly trained. They simply extend the RoPE theta, sometimes without any training, and then train on a weird concatenated dataset to make it a long dataset. This approach does not work well. A smooth, continuous long context scaling would have been much better if scaling from 8K to 1M context length.

To resolve some of these issues, use [Unsloth](https://github.com/unslothai/unsloth) for finetuning Llama3.

---

## Evaluation Metrics

When fine-tuning a language model for paraphrasing in an author's style, it's important to evaluate the quality and effectiveness of the generated paraphrases. 

The following evaluation metrics can be used to assess the model's performance:

1. **BLEU (Bilingual Evaluation Understudy):**
   - BLEU measures the n-gram overlap between the generated paraphrase and the reference text, providing a score between 0 and 1.
   - To calculate BLEU scores, you can use the `sacrebleu` library in Python.
   - Example usage: `from sacrebleu import corpus_bleu; bleu_score = corpus_bleu(generated_paraphrases, [original_paragraphs])`

2. **ROUGE (Recall-Oriented Understudy for Gisting Evaluation):**
   - ROUGE measures the overlap of n-grams between the generated paraphrase and the reference text, focusing on recall.
   - To calculate ROUGE scores, you can use the `rouge` library in Python.
   - Example usage: `from rouge import Rouge; rouge = Rouge(); scores = rouge.get_scores(generated_paraphrases, original_paragraphs)`

3. **Perplexity:**
   - Perplexity quantifies the uncertainty or confusion of the model when generating text.
   - To calculate perplexity, you can use the fine-tuned language model itself.
   - Example usage: `perplexity = model.perplexity(generated_paraphrases)`

4. **Stylometric Measures:**
   - Stylometric measures capture the writing style characteristics of the target author.
   - To extract stylometric features, you can use the `stylometry` library in Python.
   - Example usage: `from stylometry import extract_features; features = extract_features(generated_paraphrases)`

### Integration with Axolotl

To integrate these evaluation metrics into your Axolotl pipeline, follow these steps:

1. Prepare your training data by creating a dataset of paragraphs from the target author's works and splitting it into training and validation sets.

2. Fine-tune your language model using the training set, following the approach discussed earlier.

3. Generate paraphrases for the paragraphs in the validation set using the fine-tuned model.

4. Implement the evaluation metrics using the respective libraries (`sacrebleu`, `rouge`, `stylometry`) and calculate the scores for each generated paraphrase.

5. Perform human evaluation by collecting ratings and feedback from human evaluators.

6. Analyze the evaluation results to assess the quality and style of the generated paraphrases and make informed decisions to improve your fine-tuning process.

Here's an example of how you can integrate these metrics into your pipeline:

```python
from sacrebleu import corpus_bleu
from rouge import Rouge
from stylometry import extract_features

# Fine-tune the model using the training set
fine_tuned_model = train_model(training_data)

# Generate paraphrases for the validation set
generated_paraphrases = generate_paraphrases(fine_tuned_model, validation_data)

# Calculate evaluation metrics
bleu_score = corpus_bleu(generated_paraphrases, [original_paragraphs])
rouge = Rouge()
rouge_scores = rouge.get_scores(generated_paraphrases, original_paragraphs)
perplexity = fine_tuned_model.perplexity(generated_paraphrases)
stylometric_features = extract_features(generated_paraphrases)

# Perform human evaluation
human_scores = collect_human_evaluations(generated_paraphrases)

# Analyze and interpret the results
analyze_results(bleu_score, rouge_scores, perplexity, stylometric_features, human_scores)
```

Remember to install the necessary libraries (sacrebleu, rouge, stylometry) and adapt the code to fit your implementation in Axolotl or similar.

---

## AI Writing Model Comparison

In this [experiment](https://github.com/danielsobrado/ainovelprompter/tree/main/compare), I explored the capabilities and differences between various AI models in generating a 1500-word text based on a detailed prompt. I tested models from https://chat.lmsys.org/, ChatGPT4, Claude 3 Opus, and some local models in LM Studio. Each model generated the text three times to observe variability in their outputs. I also created a separate prompt for evaluating the writing of the first iteration from each model and asked ChatGPT 4 and Claude Opus 3 to provide feedback.

Through this process, I observed that some models exhibit higher variability between executions, while others tend to use similar wording. There were also significant differences in the number of words generated and the amount of dialogue, descriptions, and paragraphs produced by each model. The evaluation feedback revealed that ChatGPT suggests a more "refined" prose, while Claude recommends less purple prose. Based on these findings, I compiled a list of takeaways to incorporate into the next prompt, focusing on precision, varied sentence structures, strong verbs, unique twists on fantasy motifs, consistent tone, distinct narrator voice, and engaging pacing. Another technique to consider is asking for feedback and then rewriting the text based on that feedback. 

I'm open to collaborating with others to further fine-tune prompts for each model and explore their capabilities in creative writing tasks.

## Prompting Small LLMs

- **Direct Instructions:**
  - Use clean, specific, and direct commands.
  - Avoid verbosity and unnecessary phrases.
- **Adjective Management:**
  - Be cautious with adjectives; they may influence the model's response inappropriately.
- **Delimiters and Markdown:**
  - Use backticks, brackets, or markdown to separate distinct parts of the text.
  - Markdown helps structure and segregate sections effectively.
- **Structured Formats:**
  - Utilize JSON, markdown, HTML, etc., for input and output.
  - Constrain output using JSON schema when necessary.
- **Few-shot Examples:**
  - Provide few-shot examples from various niches to avoid overfitting.
  - Use these examples to "teach" the model steps in a process.
- **Chain-of-Thought:**
  - Implement chain-of-thought prompts to improve reasoning and procedural understanding.
  - Break down tasks into steps and guide the model through them.
- **Description Before Completion:**
  - Prompt the model to describe entities before answering.
  - Ensure that description doesn't bleed into completion unintentionally.
- **Context Management:**
  - Provide essential context only, avoid unstructured paragraph dumps.
  - Direct the model towards the desired answer with sufficient but concise context.
- **Testing and Verification:**
  - Test prompts multiple times to catch unexpected outputs.
  - Use completion ranking for relevance, clarity, and coherence.
- **Use Stories:**
  - Control output with storytelling techniques.
  - For example, write a narrative that includes the desired output format.
- **GBNF Grammars:**
  - Explore GBNF grammars to constrain and control model output.
- **Read and Refine:**
  - Review and refine generated prompts to remove unnecessary phrases and ensure clarity.

## Prompting Llama 3 8b

Models have inherent formatting biases. Some models prefer hyphens for lists, others asterisks. When using these models, it's helpful to mirror their preferences for consistent outputs.

### Key Points for Llama 3 Prompting:

- **Formatting Tendencies:**
  - Llama 3 prefers lists with bolded headings and asterisks.
  - Example:
    **Bolded Title Case Heading**

    * List items with asterisks after two newlines

    * List items separated by one newline

    **Next List**

    * More list items

    * Etc...

- **Few-shot Examples:**
  - Llama 3 follows both system prompts and few-shot examples.
  - It is flexible with prompting methods but may quote few-shot examples verbatim.

- **System Prompt Adherence:**
  - Llama 3 responds well to system prompts with detailed instructions.
  - Combining system prompts and few-shot examples yields better results.

- **Context Window:**
  - The current context window is small, limiting the use of extensive few-shot examples.
  - This may be addressed in future updates.

- **Censorship:**
  - The instruct version has some censorship but is less restricted than previous versions.

- **Intelligence:**
  - Performs well in zero-shot chain-of-thought reasoning.
  - Capable of understanding and adapting to varied inputs.

- **Consistency:**
  - Generally consistent but may directly quote examples.
  - Performance can degrade with higher temperatures.

### Usage Recommendations:

- **Lists and Formatting:**
  - Use the preferred list format for better accuracy.
  - Explicitly instruct Llama 3 on desired output formats if different from its default.

- **Chat Settings:**
  - Suitable for tasks requiring intelligence and instruction following.
  - Limited by context window for large tasks.

- **Pipeline Settings:**
  - Effective for GPT-4 style pipelines using system prompts.
  - Context window limitations restrict some tasks.

Llama 3 is flexible and intelligent but has context and quoting limitations. Adjust prompting methods accordingly.


## Acknowledgments

- Built with [Wails](https://wails.io/)
- UI components from [shadcn/ui](https://ui.shadcn.com/)
- Icons from [Lucide](https://lucide.dev/)

## Contributing

All comments are welcome. Open an issue or send a pull request if you find any bugs or have recommendations for improvement.

## License

This project is licensed under: Attribution-NonCommercial-NoDerivatives (BY-NC-ND) license See: https://creativecommons.org/licenses/by-nc-nd/4.0/deed.en
