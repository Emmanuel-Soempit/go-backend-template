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
│   │   ├── auth/
│   │   │   ├── delivery/
│   │   │   │   └── http/
│   │   │   │       ├── handler/
│   │   │   │       │   ├── handler.go
│   │   │   │       │   └── handler_Impl.go
│   │   │   │       └── routes.go
│   │   │   ├── dtos/
│   │   │   │   └── dtos.go
│   │   │   ├── repository/
│   │   │   │   ├── ent_user.go
│   │   │   │   └── user_repo.go
│   │   │   └── usecase/
│   │   │       ├── auth.go
│   │   │       └── auth_impl.go
│   │   └── routes.go
│   ├── config/
│   │   ├── app.go
│   │   └── config.go
│   ├── middleware/
│   │   ├── jwt.go
│   │   └── rate_limiter.go
│   └── utils/
│       ├── jwt.go
│       ├── password.go
│       ├── responses.go
│       └── tracer.go
└── tmp/
    └── main
```

## **Project Structure Description**

1. **internal**
   The `internal` package contains application-specific code that is not intended for external use. This follows Go's convention for internal packages.
   
   1.1. **api**
   Located within `internal`, this package handles HTTP routing and API endpoints for the Fiber web framework.
   
   1.1.1. **auth**
   Located within `internal/api`, this package contains the authentication module following clean architecture principles.
   
   1.1.1.1. **delivery/http**
   Located within `internal/api/auth`, this package handles HTTP delivery for authentication endpoints.
   
   1.1.1.1.1. **handler**
   Located within `internal/api/auth/delivery/http`, this package contains HTTP handlers for authentication.
   
   1.1.1.1.1.1. **handler.go**
   Located within `internal/api/auth/delivery/http/handler`, this file defines the AuthHandler interface.
   
   1.1.1.1.1.2. **handler_Impl.go**
   Located within `internal/api/auth/delivery/http/handler`, this file implements the authentication HTTP handlers for login and registration.
   
   1.1.1.1.2. **routes.go**
   Located within `internal/api/auth/delivery/http`, this file defines authentication routes and middleware setup.
   
   1.1.1.2. **dtos**
   Located within `internal/api/auth`, this package contains Data Transfer Objects for authentication requests and responses.
   
   1.1.1.2.1. **dtos.go**
   Located within `internal/api/auth/dtos`, this file defines DTOs for login, registration, and user responses.
   
   1.1.1.3. **repository**
   Located within `internal/api/auth`, this package handles data access for user entities.
   
   1.1.1.3.1. **user_repo.go**
   Located within `internal/api/auth/repository`, this file defines the UserRepo interface.
   
   1.1.1.3.2. **ent_user.go**
   Located within `internal/api/auth/repository`, this file implements the UserRepo using Ent ORM.
   
   1.1.1.4. **usecase**
   Located within `internal/api/auth`, this package contains business logic for authentication.
   
   1.1.1.4.1. **auth.go**
   Located within `internal/api/auth/usecase`, this file defines the AuthUsecase interface.
   
   1.1.1.4.2. **auth_impl.go**
   Located within `internal/api/auth/usecase`, this file implements authentication business logic including login and registration.
   
   1.1.2. **routes.go**
   Located within `internal/api`, this file defines main API routes and initializes authentication routes.
   
   1.2. **config**
   Located within `internal`, this package manages application configuration, database connections, and middleware setup.
   
   1.2.1. **app.go**
   Located within `internal/config`, this file contains application configuration including database setup, CORS middleware, and logging configuration.
   
   1.2.2. **config.go**
   Located within `internal/config`, this file provides main configuration initialization and Fiber app setup, coordinating all configuration components.
   
   1.3. **middleware**
   Located within `internal`, this package contains HTTP middleware for authentication, authorization, and rate limiting.
   
   1.3.1. **jwt.go**
   Located within `internal/middleware`, this file provides JWT token validation middleware.
   
   1.3.2. **rate_limiter.go**
   Located within `internal/middleware`, this file provides rate limiting middleware with configurable limits for different endpoints.
   
   1.4. **utils**
   Located within `internal`, this package contains utility functions for JWT, password hashing, and API responses.
   
   1.4.1. **jwt.go**
   Located within `internal/utils`, this file provides JWT token generation and verification functions.
   
   1.4.2. **password.go**
   Located within `internal/utils`, this file provides password hashing and verification functions using bcrypt.
   
   1.4.3. **responses.go**
   Located within `internal/utils`, this file provides standardized API response helpers.
   
   1.4.4. **tracer.go**
   Located within `internal/utils`, this file provides tracing utilities for debugging.

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

#### Authentication
```bash
# Register
POST /api/v1/auth/register
Content-Type: application/json
{
  "firstname": "John",
  "lastname": "Doe", 
  "email": "john@example.com",
  "password": "password123"
}

# Login
POST /api/v1/auth/login
Content-Type: application/json
{
  "email": "john@example.com",
  "password": "password123"
}
```

### Development Workflow

1. **Make changes to code**
2. **Air automatically restarts** the server
3. **If you modify Ent schemas**: Run `go generate ./ent`
4. **Database migrations** happen automatically on startup
