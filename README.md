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

## Configuration

1. Copy the example configuration file and update it with your settings:

```sh
cp .env.example .env
```

2. Edit `config.json` to match your database and other configuration settings.

## Running the Application

1. Start the application:

```sh
go run main.go
```

2. The API will be available at `http://localhost:8080`.

## Testing

Run the tests using:

```sh
go test ./...
```
