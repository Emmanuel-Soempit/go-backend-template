package handler

import (
	"github.com/gofiber/fiber/v2"
)

type AuthHandler interface {
	LoginHandler(*fiber.Ctx) error
	RegisterHandler(*fiber.Ctx) error
}
