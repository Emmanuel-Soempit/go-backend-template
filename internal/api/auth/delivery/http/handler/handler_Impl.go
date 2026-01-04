package handler

import (
	"go-backend-template/internal/api/auth/dtos"
	"go-backend-template/internal/api/auth/usecase"
	"log"

	"github.com/gofiber/fiber/v2"
)

type Handler struct {
	authUc usecase.AuthUsecase
}

func NewAuthHandler(authUsecase usecase.AuthUsecase) AuthHandler {
	return &Handler{
		authUc: authUsecase,
	}
}

func (h *Handler) LoginHandler(ctx *fiber.Ctx) error {

	data := new(dtos.LoginUserPayload)

	if err := ctx.BodyParser(data); err != nil {
		log.Println("Error parsing body")
		return err
	}

	log.Println(data.Email)

	dataToSend, err := h.authUc.Login(ctx.Context(), data.Email, data.Password)
	if err != nil {
		log.Println("Error logging in")
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Login Succesful",
		"data":    dataToSend,
	})
}

func (h *Handler) RegisterHandler(ctx *fiber.Ctx) error {
	data := new(dtos.RegisterUserPayload)

	if err := ctx.BodyParser(data); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	u, err := h.authUc.Register(ctx.Context(), *data)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "User created succesfully",
		"data":    u,
	})

}
