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
