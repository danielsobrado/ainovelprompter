# AI Novel Prompter

AI Novel Prompter can generate writing prompts for novels based on user-specified characteristics. 

## Features

- User registration and authentication
- Text creation and management
- Chapter creation and management
- Feedback submission and management
- Prompt generation based on traits
- Integration with a local ollama service
- Based on Berry template (https://codedthemes.gitbook.io/berry)
- Inspered on Jason Hamilton Youtube (https://www.youtube.com/@TheNerdyNovelist)

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

## Contributing

All comments are welcome. Open an issue or send a pull request if you find any bugs or have recommendations for improvement.

## License

This project is licensed under: Attribution-NonCommercial-NoDerivatives (BY-NC-ND) license See: https://creativecommons.org/licenses/by-nc-nd/4.0/deed.en