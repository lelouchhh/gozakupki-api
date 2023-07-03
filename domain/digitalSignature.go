package domain

import (
	"context"
	"fmt"
	"strings"
)

type DigitalSignature struct {
	LastName   string `json:"SN"`
	SecondName string `json:"second_name"`
	FirstName  string `json:"first_name"`
	OtherName  string `json:"G"`
	Company    string `json:"CN"`
	City       string `json:"L"`
	S          string `json:"S"`
	INN        string `json:"INN" validate:"required"`
}

func (d *DigitalSignature) SeparateNames() error {
	strs := strings.Split(d.OtherName, " ")
	if len(strs) != 2 {
		return fmt.Errorf("name is not separatable")
	}
	d.FirstName = strs[0]
	d.LastName = strs[1]
	return nil
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
