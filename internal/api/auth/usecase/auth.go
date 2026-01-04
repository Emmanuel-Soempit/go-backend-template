package usecase

import (
	"context"
	"go-backend-template/internal/api/auth/dtos"
)

type AuthUsecase interface {
	// Add your auth-related methods here
	Login(ctx context.Context, email string, password string) (*dtos.LoginResponse, error)
	Register(ctx context.Context, user dtos.RegisterUserPayload) (*dtos.UserDTO, error)
}
