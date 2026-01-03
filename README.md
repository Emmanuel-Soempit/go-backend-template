# Go Backend Template

A Go backend template that uses Air for hot reloading, Fiber for web framework, and Ent for database ORM.

## Project Structure

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

## Project Structure Description

### internal
The `internal` package contains application-specific code that is not intended for external use. This follows Go's convention for internal packages.

#### api
Handles HTTP routing and API endpoints for the Fiber web framework.

##### routes.go
Defines API routes and handlers for the application, including health checks and other endpoints.

#### config
Manages application configuration, database connections, and middleware setup.

##### app.go
Contains application configuration including database setup, CORS middleware, and logging configuration.

##### config.go
Main configuration initialization and Fiber app setup, coordinating all configuration components.

#### logger.go
Logging configuration and utilities for the application.
