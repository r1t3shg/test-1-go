# Go Backend API

A minimal production-style Go backend API built with `net/http`.

## Features

- Health check endpoint
- Simple root endpoint
- Environment-based port configuration
- Container-ready with Docker

## Endpoints

- `GET /` - Returns "Hello from Go API"
- `GET /health` - Returns `{"status": "ok"}`

## Prerequisites

- Go 1.22 or later
- Docker (optional, for containerized deployment)

## Local Development

### Run with Go

```bash
# Run directly
go run main.go

# Or build and run
go build -o api main.go
./api
```

The server will start on port 8080 by default. Access it at `http://localhost:8080`

### Custom Port

```bash
# Set PORT environment variable
export PORT=3000
go run main.go

# On Windows PowerShell
$env:PORT=3000
go run main.go
```

## Testing

### Using curl

```bash
# Test root endpoint
curl http://localhost:8080/

# Test health endpoint
curl http://localhost:8080/health

# Test with custom port
curl http://localhost:3000/health
```

### Expected Responses

**GET /**
```
Hello from Go API
```

**GET /health**
```json
{"status":"ok"}
```

### Using PowerShell (Windows)

```powershell
# Test root endpoint
Invoke-WebRequest -Uri http://localhost:8080/ -Method GET

# Test health endpoint
Invoke-WebRequest -Uri http://localhost:8080/health -Method GET
```

## Docker

### Build Image

```bash
docker build -t go-api .
```

### Run Container

```bash
# Default port (8080)
docker run -p 8080:8080 go-api

# Custom port
docker run -p 3000:3000 -e PORT=3000 go-api
```

### Test Docker Container

```bash
# Test endpoints
curl http://localhost:8080/
curl http://localhost:8080/health
```

## Environment Variables

- `PORT` - Server port (default: 8080)

## Project Structure

```
.
├── main.go      # Main application code
├── go.mod       # Go module file
├── Dockerfile   # Multi-stage Docker build
└── README.md    # This file
```

