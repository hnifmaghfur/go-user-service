# Go User Service

A RESTful API service for user management built with Go.

## Project Structure

```
├── .env                  # Environment variables for local development
├── .env.example          # Example environment variables template
├── .github               # GitHub workflows and configuration
├── .gitignore            # Git ignore rules
├── Makefile              # Build and development commands
├── README.md             # Project documentation
├── cmd/                  # Application entry points
│   └── service/          # Main service executable
├── docs/                 # Documentation files
├── go.mod                # Go module definition
├── internal/             # Private application code
│   ├── config/           # Configuration handling
│   ├── db/               # Database connections and migrations
│   ├── models/           # Data models
│   ├── repositories/     # Data access layer
│   ├── requests/         # Request models/DTOs
│   ├── responses/        # Response models/DTOs
│   ├── server/           # HTTP server setup
│   └── services/         # Business logic
├── migrations/           # Database migration files
└── tests/                # Integration and end-to-end tests
```

## Prerequisites

- Go 1.24.3 or higher
- PostgreSQL (or your database of choice)
- Make (optional, for using the Makefile commands)

## Getting Started

### Environment Setup

1. Copy the example environment file:

```bash
cp .env.example .env
```

2. Edit the `.env` file with your local configuration

### Running the Application

```bash
# Build and run the service
make run

# Or using Go directly
go run cmd/service/main.go
```

### Development Commands

```bash
# Run tests
make test

# Format code
make fmt

# Lint code
make lint

# Generate API documentation
make docs

# Run database migrations
make migrate
```

## API Documentation

API documentation is available at `/docs` when the server is running.

## License

[MIT](LICENSE)
