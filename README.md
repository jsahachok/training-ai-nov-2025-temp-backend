# Workshop4 Backend - Go + Fiber

Backend API built with Go and Fiber framework.

## Features

- Fast HTTP server using Fiber v2
- CORS middleware enabled
- Request logging
- Health check endpoint
- JSON responses

## API Endpoints

- `GET /` - Returns hello world message
- `GET /health` - Health check endpoint

## Installation

1. Make sure you have Go installed (version 1.19 or later)
2. Install dependencies:
   ```bash
   go mod tidy
   ```

## Running the Server

```bash
go run main.go
```

The server will start on port 3000.

## Testing

Test the main endpoint:
```bash
curl http://localhost:3000/
```

Expected response:
```json
{
  "message": "hello world"
}
```

Test health endpoint:
```bash
curl http://localhost:3000/health
```

Expected response:
```json
{
  "status": "ok",
  "message": "Server is running"
}
```

## Build

To build the application:
```bash
go build -o server main.go
```

## Technologies Used

- [Go](https://golang.org/) - Programming language
- [Fiber](https://gofiber.io/) - Web framework