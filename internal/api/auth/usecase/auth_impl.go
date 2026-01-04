package usecase

import (
	"context"
	"errors"
	"fmt"
	"go-backend-template/internal/api/auth/dtos"
	"go-backend-template/internal/api/auth/repository"
	"go-backend-template/internal/utils"
	"log"
)

type authUseCase struct {
	userRepo repository.UserRepo
}

func NewAuthUsecase(userRepo repository.UserRepo) AuthUsecase {
	return &authUseCase{
		userRepo: userRepo,
	}
}

// Implement authentication logic here
func (u *authUseCase) Login(ctx context.Context, email string, password string) (*dtos.LoginResponse, error) {
	// Find user by email
	user, err := u.userRepo.FindByEmail(ctx, email)
	if err != nil {
		return nil, err
	}

	// Verify password (this should use proper password hashing)

	if err := utils.VerifyPassword(user.Password, password); err != nil {
		return nil, fmt.Errorf("invalid credentials")
	}

	token, err := utils.GenerateJwtToken(user)
	if err != nil {
		return nil, err
	}

	userDTO := dtos.UserDTO{
		ID:        user.ID,
		Firstname: user.Firstname,
		Lastname:  user.Lastname,
		Email:     user.Email,
	}

	data := dtos.LoginResponse{
		User:  userDTO,
		Token: token,
	}
	return &data, nil
}

func (u *authUseCase) Register(ctx context.Context, user dtos.RegisterUserPayload) (*dtos.UserDTO, error) {
	// Implementation would go here

	hashed, err := utils.HashPassword(user.Password)
	if err != nil {
		return nil, errors.New("could not hash password")
	}

	user.Password = string(hashed)

	newUser, err := u.userRepo.CreateNewUser(ctx, user)
	if err != nil {
		return nil, err
	}
	userDTO := dtos.UserDTO{
		ID:        newUser.ID,
		Firstname: newUser.Firstname,
		Lastname:  newUser.Lastname,
		Email:     newUser.Email,
	}
	log.Println("Register business logic - user registered successfully")
	return &userDTO, nil
}
