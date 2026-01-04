package api

import (
	"go-backend-template/ent"
	authRoutes "go-backend-template/internal/api/auth/delivery/http"

	"github.com/gofiber/fiber/v2"
)

func InitializeApiRoutes(app *fiber.App, client *ent.Client) {

	// Example implementation of global rate limiting using default configs
	// app.Use(middleware.RateLimiter(middleware.DefaultConfig))

	app.Static("/public", "./public")
	app.Get("/health-check", func(c *fiber.Ctx) error {
		return c.SendString("Server is running!...")
	})

	routeGroup := app.Group("/api/v1")

	// Register auth routes
	authRoutes.RegisterAuthRoutes(routeGroup, client)
}
