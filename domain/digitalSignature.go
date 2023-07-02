package domain

import "context"

type DigitalSignature struct {
}

type DigitalSignatureUsecase interface {
	SignIn(ctx context.Context, s DigitalSignature) (string, error)
	SignUp(ctx context.Context, s DigitalSignature) error
}

type DigitalSignatureRepository interface {
	SignIn(ctx context.Context, s DigitalSignature) (string, error)
	SignUp(ctx context.Context, s DigitalSignature) error
	FulfillData(ctx context.Context, s DigitalSignature) error
}
