# QR Generator

A Go-based QR code generation service with a clean architecture foundation.

## Features

- **Clean Architecture**: Well-structured codebase following clean architecture principles
- **Security**: API key authentication and security headers
- **Configuration Management**: Flexible configuration via YAML and environment variables
- **Simple & Lightweight**: No database dependencies, served on main path/domain

## Tech Stack

- **Backend**: Go 1.24.0
- **Web Framework**: Gin
- **Configuration**: Viper
- **Architecture**: Clean Architecture

## Quick Start

### Prerequisites

- Go 1.24.0 or later

### Installation

1. Clone the repository:
```bash
git clone https://github.com/rizkyduut/qr-generator.git
cd qr-generator
```

2. Copy the configuration example:
```bash
cp config.example.yaml config.yaml
```

3. Update `config.yaml` with your preferences.

4. Install dependencies:
```bash
go mod download
```

5. Run the application:
```bash
go run cmd/server/main.go
```

The server will start on port 8080 (configurable).

## Configuration

The application uses a YAML configuration file. Key settings:

### Server Configuration
```yaml
server:
  port: ":8080"
  max_body_size: 1048576  # 1MB
  read_timeout: 10        # seconds
  write_timeout: 10       # seconds
```

### Security Configuration
```yaml
security:
  api_key: "your-api-key"
  enable_security_headers: true
```

## Environment Variables

You can also configure the application using environment variables with the `QR_` prefix:

- `QR_SERVER_PORT` - Server port
- `QR_SERVER_MAX_BODY_SIZE` - Maximum request body size
- `QR_SERVER_READ_TIMEOUT` - Read timeout in seconds
- `QR_SERVER_WRITE_TIMEOUT` - Write timeout in seconds
- `QR_SECURITY_API_KEY` - API key for authentication
- `QR_SECURITY_ENABLE_SECURITY_HEADERS` - Enable security headers

## API Endpoints

### Health Check
- **GET** `/ping` - Health check endpoint
  - Returns: `{"message": "pong!"}`

## Project Structure

```
├── cmd/server/           # Application entry point
├── internal/
│   └── config/          # Configuration management
├── config.yaml          # Configuration file (not in git)
├── config.example.yaml  # Configuration template
└── README.md
```

## Development

### Building
```bash
go build -o main cmd/server/main.go
```

### Running Tests
```bash
go test ./...
```
