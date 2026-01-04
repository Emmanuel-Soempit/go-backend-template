package http

import (
	"go-backend-template/ent"
	"go-backend-template/internal/api/auth/delivery/http/handler"
	"go-backend-template/internal/api/auth/repository"
	"go-backend-template/internal/api/auth/usecase"
	"go-backend-template/internal/middleware"

	// "go-backend-template/internal/middleware"

	"github.com/gofiber/fiber/v2"
)

func RegisterAuthRoutes(router fiber.Router, client *ent.Client) {
	authGroup := router.Group("/auth")

	// How to implement jwt for all auth routes
	authGroup.Use(middleware.CheckJwtToken)

	// Intiaizes structs
	userRepo := repository.NewEntUserRepo(client)
	authUseCase := usecase.NewAuthUsecase(userRepo)
	authHandler := handler.NewAuthHandler(authUseCase)

	// Routes
	authGroup.Post("/login", authHandler.LoginHandler)
	authGroup.Post("/register", authHandler.RegisterHandler)

}
