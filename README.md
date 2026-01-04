# **Go Backend Template**

A Go backend template that uses Air for hot reloading, Fiber for web framework, and Ent for database ORM.

## **Project Structure**

```
go-backend-template/
├── .env
├── .env.example
├── .gitignore
├── docker-compose.yml
├── go.mod
├── go.sum
├── main.go
├── README.md
├── ent/
├── internal/
│   ├── api/
│   │   └── routes.go
│   ├── config/
│   │   ├── app.go
│   │   └── config.go
│   └── logger.go
└── tmp/
    └── main
```

## **Project Structure Description**

1. **internal**
   The `internal` package contains application-specific code that is not intended for external use. This follows Go's convention for internal packages.
   
   1.1. **api**
   Located within `internal`, this package handles HTTP routing and API endpoints for the Fiber web framework.
   
   1.1.1. **routes.go**
   Located within `internal/api`, this file defines API routes and handlers for the application, including health checks and other endpoints.
   
   1.2. **config**
   Located within `internal`, this package manages application configuration, database connections, and middleware setup.
   
   1.2.1. **app.go**
   Located within `internal/config`, this file contains application configuration including database setup, CORS middleware, and logging configuration.
   
   1.2.2. **config.go**
   Located within `internal/config`, this file provides main configuration initialization and Fiber app setup, coordinating all configuration components.
   
   1.3. **logger.go**
   Located within `internal`, this file provides logging configuration and utilities for the application.

## **Getting Started**

### Prerequisites
- Docker and Docker Compose
- Go 1.24+
- Air (for hot reloading)

### Setup Instructions

1. **Clone the repository**
   ```bash
   git clone git@github.com:Emmanuel-Soempit/go-backend-template.git
   cd go-backend-template
   ```

2. **Install dependencies**
   ```bash
   go mod download
   ```

3. **Set up environment variables**
   ```bash
   cp .env.example .env
   ```
   Update `.env` with your configuration:
   ```env
   DATABASE_URL="postgres://postgres:postgres@localhost:5432/go_backend_db?sslmode=disable"
   PORT=3000
   JWT_SECRET="your-super-secret-jwt-key"
   ```

4. **Start the database**
   ```bash
   docker compose up -d
   ```

5. **Install Air for hot reloading** (if not already installed)
   ```bash
   go install github.com/cosmtrek/air@latest
   ```

6. **Generate Ent code** (if you modify schemas)
   ```bash
   go generate ./ent
   ```

7. **Run the application**
   ```bash
   air
   ```

### API Endpoints

The application will be available at `http://localhost:3000`

#### Health Check
```bash
GET /health-check
```

### Development Workflow

1. **Make changes to code**
2. **Air automatically restarts** the server
3. **If you modify Ent schemas**: Run `go generate ./ent`
4. **Database migrations** happen automatically on startup
