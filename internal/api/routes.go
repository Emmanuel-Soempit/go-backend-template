package api

import (
	"go-backend-template/ent"
	"go-backend-template/internal/middleware"

	"github.com/gofiber/fiber/v2"
)

func InitializeApiRoutes(app *fiber.App, client *ent.Client) {

	app.Use(middleware.RateLimiter(middleware.DefaultConfig))
	app.Get("/health-check", func(c *fiber.Ctx) error {
		return c.SendString("Server is running!...")
	})

}
