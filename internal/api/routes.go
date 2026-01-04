package api

import (
	"go-backend-template/ent"
	authRoutes "go-backend-template/internal/api/auth/delivery/http"

	"github.com/gofiber/fiber/v2"
)

func InitializeApiRoutes(app *fiber.App, client *ent.Client) {
	app.Static("/public", "./public")

	app.Get("/health-check", func(c *fiber.Ctx) error {
		return c.SendString("Server is running!...")
	})

	routeGroup := app.Group("/api/v1")

	// Register auth routes
	authRoutes.RegisterAuthRoutes(routeGroup, client)
}
