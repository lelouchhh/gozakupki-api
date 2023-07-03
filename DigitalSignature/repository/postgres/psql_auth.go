package usecase

import (
	"context"
	"fmt"
	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/jmoiron/sqlx"
	"gozakupki-api/domain"
	"gozakupki-api/pkg/hash"
	"strings"
	"time"
)

type myAuthRepository struct {
	db *sqlx.DB
}

func NewAuthRepository(conn *sqlx.DB) domain.AuthRepository {
	return &myAuthRepository{conn}
}

func (a *myAuthRepository) GetUser(ctx context.Context, auth domain.Auth) (domain.Auth, error) {
	var user domain.Auth
	fmt.Println(auth)
	if auth.Email == "" {
		err := a.db.GetContext(ctx, &user, "select user_id as id, email, login from auth.t_user where login = $1 and hex_pass = $2 and user_status = $3;", auth.Login, auth.Password, "1")
		if err != nil {
			return domain.Auth{}, domain.ErrUnauthorized
		}

	} else {
		err := a.db.GetContext(ctx, &user, "select user_id as id, email, login from auth.t_user where email = $1 and hex_pass = $2 and user_status = $3;", auth.Email, auth.Password, "1")
		if err != nil {

			return domain.Auth{}, domain.ErrUnauthorized
		}
	}
	return user, nil
}

func (a *myAuthRepository) SignUp(ctx context.Context, auth domain.Auth) error {
	err := a.doesUserExist(ctx, auth)
	if err != nil {
		return err
	}
	query := `insert into auth.t_user (login, hex_pass, email, email_pass, user_status, time_reg, user_role) values ($1, $2, $3, $4, $5, $6, $7);`
	stmt, err := a.db.PrepareContext(ctx, query)
	if err != nil {
		return domain.ErrInternalServerError
	}
	emailHash, err := hash.GenerateHash(auth.Login, auth.Email)
	if err != nil {
		return err
	}
	_, err = stmt.ExecContext(ctx, auth.Login, auth.Password, strings.ToLower(auth.Email), emailHash, "1", time.Now(), "client")
	if err != nil {
		return domain.ErrBadParamInput
	}

	return nil
}
func (a *myAuthRepository) doesUserExist(ctx context.Context, auth domain.Auth) error {
	var doesExist bool
	err := a.db.GetContext(ctx, &doesExist, "select exists(select 1 from auth.t_user where email = $1 or login = $2)", auth.Email, auth.Login)
	if err != nil {
		return err
	}
	if doesExist {
		return fmt.Errorf("user already registred")
	}
	return nil
}
func (a *myAuthRepository) ConfirmUserByEmail(ctx context.Context, hash string) error {
	query := `update auth.t_user set user_status = '1' where email_pass = $1`
	stmt, err := a.db.PrepareContext(ctx, query)
	if err != nil {
		return domain.ErrInternalServerError
	}
	_, err = stmt.ExecContext(ctx, hash)
	if err != nil {
		return domain.ErrBadParamInput
	}
	return err
}

func (a *myAuthRepository) ResetUserEmailPass(ctx context.Context, auth domain.Auth) error {
	//TODO implement me
	panic("implement me")
}
