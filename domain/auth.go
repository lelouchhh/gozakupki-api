package domain

import "context"

type Auth struct {
	ID       int64  `json:"-"`
	Login    string `json:"login" validate:"min=6"`
	Password string `json:"password" validate:"min=6"`
	Email    string `json:"email" validate:"email"`
	Hash     string `json:"hash"`
}

type AuthUsecase interface {
	SignIn(ctx context.Context, auth Auth) (string, error)
	SignUp(ctx context.Context, auth Auth) error
	ConfirmUser(ctx context.Context, hash string) error
	CheckToken(ctx context.Context, token string) error
}
type AuthRepository interface {
	GetUser(ctx context.Context, auth Auth) (Auth, error)
	SignUp(ctx context.Context, auth Auth) error
	ConfirmUserByEmail(ctx context.Context, hash string) error
}
