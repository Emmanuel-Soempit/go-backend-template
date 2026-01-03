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

### internal
The `internal` package contains application-specific code that is not intended for external use. This follows Go's convention for internal packages.

#### api
Located within `internal`, this package handles HTTP routing and API endpoints for the Fiber web framework.

##### routes.go
Located within `internal/api`, this file defines API routes and handlers for the application, including health checks and other endpoints.

#### config
Located within `internal`, this package manages application configuration, database connections, and middleware setup.

##### app.go
Located within `internal/config`, this file contains application configuration including database setup, CORS middleware, and logging configuration.

##### config.go
Located within `internal/config`, this file provides main configuration initialization and Fiber app setup, coordinating all configuration components.

#### logger.go
Located within `internal`, this file provides logging configuration and utilities for the application.
