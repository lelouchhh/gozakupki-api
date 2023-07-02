package domain

import "context"

type Auth struct {
	ID       int64  `json:"-"`
	Login    string `json:"login"`
	Password string `json:"password" validate:"required,min=6"`
	Email    string `json:"email" validate:"required,email"`
	DevInfo  string `json:"devInfo"`
	Hash     string `json:"hash"`
}

type AuthUsecase interface {
	SignIn(ctx context.Context, auth Auth) (string, error)
	SignUp(ctx context.Context, auth Auth) error
	ConfirmUser(ctx context.Context, auth Auth) error
	CheckToken(ctx context.Context, token string) error
}
type AuthRepository interface {
	GetUser(ctx context.Context, auth Auth) (Auth, error)
	SignUp(ctx context.Context, auth Auth) error
	ConfirmUserByEmail(ctx context.Context, auth Auth) error
}
