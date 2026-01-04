package utils

import (
	"github.com/gofiber/fiber/v2"
)

// Response is the standard API response structure.
type Response struct {
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
	Error   interface{} `json:"error,omitempty"`
}

// Success sends a 200 OK response with optional data.
func Success(ctx *fiber.Ctx, message string, data interface{}) error {
	return ctx.Status(fiber.StatusOK).JSON(Response{
		Status:  "success",
		Message: message,
		Data:    data,
	})
}

// Created sends a 201 Created response with optional data.
func Created(ctx *fiber.Ctx, message string, data interface{}) error {
	return ctx.Status(fiber.StatusCreated).JSON(Response{
		Status:  "created",
		Message: message,
		Data:    data,
	})
}

// Failed sends a 400 Bad Request response with an error message.
func Failed(ctx *fiber.Ctx, message string, err interface{}) error {
	return ctx.Status(fiber.StatusBadRequest).JSON(Response{
		Status:  "failed",
		Message: message,
		Error:   err,
	})
}

// Unauthorized sends a 401 Unauthorized response.
func Unauthorized(ctx *fiber.Ctx, message string) error {
	return ctx.Status(fiber.StatusUnauthorized).JSON(Response{
		Status:  "unauthorized",
		Message: message,
	})
}

// NotFound sends a 404 Not Found response.
func NotFound(ctx *fiber.Ctx, message string) error {
	return ctx.Status(fiber.StatusNotFound).JSON(Response{
		Status:  "not_found",
		Message: message,
	})
}

// InternalError sends a 500 Internal Server Error response.
func InternalError(ctx *fiber.Ctx, message string, err interface{}) error {
	return ctx.Status(fiber.StatusInternalServerError).JSON(Response{
		Status:  "error",
		Message: message,
		Error:   err,
	})
}
