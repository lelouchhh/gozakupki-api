package usecase

import (
	"context"
	"gozakupki-api/domain"
	"gozakupki-api/pkg/JWT"
	"gozakupki-api/pkg/hash"
	"gozakupki-api/pkg/random"
	"time"
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

	auth.Password = hash.GetMD5Hash(auth.Password + "xd")

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
	auth.Password = hash.GetMD5Hash(auth.Password + "xd")
	err := a.AuthRepo.SignUp(ctx, auth)
	if err != nil {
		return err
	}
	err = a.MailRepo.SendSingleData(ctx, domain.Mail{Data: auth.Hash, To: auth.Email, EndPoint: "basic/confirm_account"})
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

func (a authUsecase) ResetPassword(ctx context.Context, auth domain.Auth) error {
	ctx, cancel := context.WithTimeout(ctx, a.contextTimeout)
	defer cancel()
	password := random.GeneratePassword(random.PasswordLength, random.MinSpecialChar, random.MinNum, random.MinUpperCase)

	auth, err := a.AuthRepo.GetUserByEmail(ctx, auth)
	if err != nil {
		return err
	}

	err = a.MailRepo.SendSingleData(ctx, domain.Mail{To: auth.Email, Data: password, EndPoint: "basic/reset_password"})
	if err != nil {
		return err
	}
	token, err := JWT.GenerateToken(auth.ID, auth.Login, auth.Email)
	if err != nil {
		return err
	}
	auth.Password = token

	err = a.AuthRepo.ResetPassword(ctx, auth)
	if err != nil {
		return err
	}
	return nil
}

//func (d *DeliverUsecase) Delete(c context.Context, id int) error {
//	return nil
//}
//
//func (d *DeliverUsecase) GetById(c context.Context, id string) (domain.Deliver, error) {
//	return domain.Deliver{}, nil
//}
