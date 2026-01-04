package repository

import (
	"context"
	"go-backend-template/ent"
	"go-backend-template/ent/user"
	"go-backend-template/internal/api/auth/dtos"
	"log"
)

type entUserRepo struct {
	client *ent.Client
}

func NewEntUserRepo(client *ent.Client) UserRepo {
	return &entUserRepo{client: client}
}

// User Repo Implementations
func (r *entUserRepo) FindByEmail(ctx context.Context, email string) (*ent.User, error) {
	u, err := r.client.User.Query().Where(user.Email(email)).Only(ctx)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return u, nil
}

func (r *entUserRepo) CreateNewUser(ctx context.Context, user dtos.RegisterUserPayload) (*ent.User, error) {
	u, err := r.client.User.
		Create().
		SetFirstname(user.Firstname).
		SetLastname(user.Lastname).
		SetPassword(user.Password).
		SetEmail(user.Email).
		Save(ctx)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return u, nil
}
