package usecase

import (
	"context"
	"fmt"
	"gozakupki-api/pkg/JWT"
	"gozakupki-api/pkg/hash"
	"time"

	"gozakupki-api/domain"
)

type authUsecase struct {
	AuthRepo       domain.AuthRepository
	MailRepo       domain.MailRepository
	contextTimeout time.Duration
}

func NewAuthUsecase(a domain.AuthRepository, m domain.MailRepository, timeout time.Duration) domain.AuthUsecase {
	return &authUsecase{
		AuthRepo:       a,
		MailRepo:       m,
		contextTimeout: timeout,
	}
}
func (a authUsecase) CheckToken(ctx context.Context, token string) error {
	err := JWT.IsValid(token)
	if err != nil {
		return err
	}
	return nil
}

func (a authUsecase) SignIn(ctx context.Context, auth domain.Auth) (string, error) {
	ctx, cancel := context.WithTimeout(ctx, a.contextTimeout)
	defer cancel()

	auth.Password = hash.GetMD5Hash(auth.Login + "+" + auth.Password)

	user, err := a.AuthRepo.GetUser(ctx, auth)
	if err != nil {
		return "", err
	}
	token, err := JWT.GenerateToken(user.ID, user.Login, user.Email)
	if err != nil {
		return "", domain.ErrInternalServerError
	}
	return token, nil

}

func (a authUsecase) SignUp(ctx context.Context, auth domain.Auth) error {
	ctx, cancel := context.WithTimeout(ctx, a.contextTimeout)
	defer cancel()
	auth.Password = hash.GetMD5Hash(auth.Login + "+" + auth.Password)
	err := a.AuthRepo.SignUp(ctx, auth)
	if err != nil {
		return err
	}
	err = a.MailRepo.SendSingleMessage(ctx, domain.Mail{Message: auth.Hash, To: auth.Email})
	fmt.Println(err)
	if err != nil {
		return err
	}
	return nil

}

func (a authUsecase) ConfirmUser(ctx context.Context, h string) error {
	ctx, cancel := context.WithTimeout(ctx, a.contextTimeout)
	defer cancel()
	err := hash.IsValid(h)
	if err != nil {
		return err
	}
	err = a.AuthRepo.ConfirmUserByEmail(ctx, h)
	if err != nil {
		return err
	}
	return nil
}

func (a authUsecase) ResetUserEmailPass(ctx context.Context, auth domain.Auth) {
	//TODO implement me
	panic("implement me")
}

//func (d *DeliverUsecase) Delete(c context.Context, id int) error {
//	return nil
//}
//
//func (d *DeliverUsecase) GetById(c context.Context, id string) (domain.Deliver, error) {
//	return domain.Deliver{}, nil
//}