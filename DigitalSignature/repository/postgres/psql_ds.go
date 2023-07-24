package usecase

import (
	"context"
	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/jmoiron/sqlx"
	"gozakupki-api/domain"
)

type myDSRepository struct {
	db *sqlx.DB
}

func (m myDSRepository) SignIn(ctx context.Context, s domain.DigitalSignature) (string, error) {
	//TODO implement me
	panic("implement me")
}

func (m myDSRepository) SignUp(ctx context.Context, s domain.DigitalSignature) error {
	//TODO implement me
	panic("implement me")
}

func (m myDSRepository) FulfillData(ctx context.Context, s domain.DigitalSignature) error {
	//TODO implement me
	panic("implement me")
}

func NewDSRepository(conn *sqlx.DB) domain.DigitalSignatureRepository {
	return &myDSRepository{conn}
}
