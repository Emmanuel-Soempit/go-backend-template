package middleware

import (
	"go-backend-template/internal/utils"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func CheckJwtToken(ctx *fiber.Ctx) error {
	authHeader := ctx.Get("Authorization")
	if authHeader == "" {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Missing Authorization header",
		})
	}

	parts := strings.SplitN(authHeader, " ", 2)
	if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Invalid Authorization header format",
		})
	}

	tokenString := parts[1]
	claims, err := utils.VerfyJwtToken(tokenString)
	if err != nil {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	// Optionally, set claims in context locals for downstream handlers
	ctx.Locals("user", claims)

	return ctx.Next()
}
