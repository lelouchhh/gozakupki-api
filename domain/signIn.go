package domain

type SignIn struct {
	ID       int64  `json:"-"`
	Login    string `json:"login" validate:"required,min=6"`
	Password string `json:"password" validate:"required,min=6"`
	Email    string `json:"email" validate:"required,email"`
}
