# Go Web Template

A production-ready Go web application template following best practices and clean architecture principles.

## Features

- 🏗️ **Clean Architecture** - Organized with handlers, services, and repositories
- 🔧 **RESTful API** - Well-structured API endpoints with JSON responses
- ✅ **Comprehensive Testing** - Unit tests with table-driven patterns
- 🔒 **Middleware** - Logging, recovery, CORS, and content-type middleware
- ⚙️ **Configuration** - Environment-based configuration with validation
- 🐳 **Docker Ready** - Multi-stage Dockerfile for optimized builds
- 📖 **API Documentation** - OpenAPI 3.0 specification included

## Project Structure

```
.
├── api/                    # API documentation (OpenAPI specs)
├── cmd/
│   └── api/               # Application entry point
├── internal/
│   ├── config/            # Configuration management
│   ├── handler/           # HTTP request handlers
│   ├── middleware/        # HTTP middleware
│   ├── models/            # Data models and DTOs
│   ├── repository/        # Data access layer
│   ├── server/            # HTTP server setup
│   └── service/           # Business logic layer
├── static/                # Static assets (CSS, JS)
├── templates/             # HTML templates
├── Dockerfile             # Docker build configuration
├── Makefile               # Build automation
├── go.mod                 # Go module definition
└── README.md              # Project documentation
```

## Requirements

- Go 1.21 or higher
- Docker (optional, for containerization)

## Quick Start

### Using Go

```bash
# Clone the repository
git clone https://github.com/vidsrvcom/gotemp.git
cd gotemp

# Install dependencies
go mod tidy

# Run the application
go run main.go

# Or run from cmd/api
go run ./cmd/api

# Or use make
make run
```

### Using Docker

```bash
# Build the image
make docker-build

# Run the container
make docker-run
```

The application will be available at `http://localhost:8080`.

## API Endpoints

| Method | Endpoint | Description |
|--------|----------|-------------|
| GET | `/` | Home page |
| GET | `/health` | Health check |
| GET | `/api/users` | List all users |
| POST | `/api/users` | Create a new user |
| GET | `/api/users/{id}` | Get user by ID |
| PUT | `/api/users/{id}` | Update user |
| DELETE | `/api/users/{id}` | Delete user |

## Development

### Running Tests

```bash
# Run all tests
make test

# Run tests with coverage
make test-coverage
```

### Linting

```bash
make lint
```

### Building

```bash
make build
```

## Configuration

The application can be configured using environment variables:

| Variable | Default | Description |
|----------|---------|-------------|
| `HOST` | `0.0.0.0` | Server host |
| `PORT` | `8080` | Server port |
| `ENVIRONMENT` | `development` | Environment (development/production) |
| `LOG_LEVEL` | `info` | Log level |

You can also create a `.env` file in the project root:

```env
HOST=0.0.0.0
PORT=8080
ENVIRONMENT=development
LOG_LEVEL=info
```

## Architecture

This project follows clean architecture principles:

1. **Handlers** - Handle HTTP requests and responses
2. **Services** - Contain business logic
3. **Repositories** - Handle data persistence
4. **Models** - Define data structures

Dependencies flow inward: Handlers → Services → Repositories

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
