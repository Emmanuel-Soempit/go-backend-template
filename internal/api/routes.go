package api

import (
	"go-backend-template/ent"

	"github.com/gofiber/fiber/v2"
)

func InitializeApiRoutes(app *fiber.App, client *ent.Client) {

	app.Get("/health-check", func(c *fiber.Ctx) error {
		return c.SendString("Server is running!...")
	})

}
