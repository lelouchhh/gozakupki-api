package domain

import "context"

type SignUp struct {
	Login    string `json:"login" validate:"required,min=6"`
	Password string `json:"password" validate:"required,min=6"`
	Email    string `json:"email" validate:"required,email"`
}

type SignUpUsecase interface {
	SignUp(ctx context.Context, auth Auth) error
}
type SignUpRepository interface {
	SignUp(ctx context.Context, auth Auth) error
}
