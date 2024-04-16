# AI Novel Prompter

AI Novel Prompter can generate writing prompts for novels based on user-specified characteristics. 

## Features

- User registration and authentication
- Text creation and management
- Chapter creation and management
- Feedback submission and management
- Prompt generation based on traits

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

- Node.js (v14 or higher)
- Go (v1.16 or higher)
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

## Contributing

All comments are welcome. Open an issue or send a pull request if you find any bugs or have recommendations for improvement.

## License

This project is licensed under: Attribution-NonCommercial-NoDerivatives (BY-NC-ND) license See: https://creativecommons.org/licenses/by-nc-nd/4.0/deed.en