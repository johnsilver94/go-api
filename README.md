# Go API Project

## Introduction

This project is a RESTful API built with Go. It provides various endpoints to manage resources.

## Prerequisites

- Go 1.16 or higher
- Git
- A running instance of a database (e.g., PostgreSQL, MySQL)

## Installation

1. Clone the repository:

```sh
git clone https://github.com/johnsilver94/go-api.git
```

2. Navigate to the project directory:

```sh
cd go-api
```

3. Install dependencies:

```sh
go mod tidy
```

## Folder Structure

- `.editorconfig`: Configuration file for maintaining consistent coding styles.
- `.env`: Environment variables file.
- `.env.example`: Example environment variables file.
- `.gitignore`: Specifies files and directories to be ignored by Git.
- `.husky/`: Configuration for Husky, which manages Git hooks.
- `.vscode/`: Visual Studio Code workspace settings.
- `bin/`: Compiled binary files.
- `cmd/`: Entry points for the application.
  - `api/`: Main API server entry point.
  - `migrate/`: Database migration scripts.
- `configs/`: Configuration files.
- `db/`: Database connection and initialization.
- `http/`: HTTP request files for testing endpoints.
- `Makefile`: Makefile for build automation.

## Configuration

1. Copy the example configuration file and update it with your settings:

```sh
cp .env.example .env
```

2. Edit `config.json` to match your database and other configuration settings.

## Running the Application

1. Start the application:

```sh
make run
```

2. The API will be available at `http://localhost:{PORT}`.

## Database Migrations

1. Create a new migration:

```sh
make migration <migration_name>
```

2. Apply all migrations:

```sh
make migration-up
```

3. Rollback the last migration:

```sh
make migration-down
```

## Testing

Run the tests using:

```sh
go test ./...
```
