package usecase

import (
	"context"
	"time"

	"gozakupki-api/domain"
)

type authUsecase struct {
	DsRepo         domain.DigitalSignatureRepository
	contextTimeout time.Duration
}

func NewAuthUsecase(a domain.DigitalSignatureRepository, timeout time.Duration) domain.AuthUsecase {
	return &authUsecase{
		DsRepo:         a,
		contextTimeout: timeout,
	}
}
func (a authUsecase) SignIn(ctx context.Context, auth domain.Auth) (string, error) {
	//TODO implement me
	panic("implement me")
}

func (a authUsecase) SignUp(ctx context.Context, auth domain.Auth) error {
	//TODO implement me
	panic("implement me")
}

func (a authUsecase) ConfirmUser(ctx context.Context, hash string) error {
	//TODO implement me
	panic("implement me")
}

func (a authUsecase) CheckToken(ctx context.Context, token string) error {
	//TODO implement me
	panic("implement me")
}
