package usecase

import (
	"context"
	"gozakupki-api/domain"
	"reflect"
	"testing"
	"time"
)

func TestNewAuthUsecase(t *testing.T) {
	type args struct {
		a       domain.AuthRepository
		timeout time.Duration
	}
	tests := []struct {
		name string
		args args
		want domain.AuthUsecase
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewAuthUsecase(tt.args.a, tt.args.timeout); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewAuthUsecase() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_authUsecase_CheckToken(t *testing.T) {
	type fields struct {
		AuthRepo       domain.AuthRepository
		MailRepo       domain.Mail
		contextTimeout time.Duration
	}
	type args struct {
		ctx   context.Context
		token string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := authUsecase{
				AuthRepo:       tt.fields.AuthRepo,
				MailRepo:       tt.fields.MailRepo,
				contextTimeout: tt.fields.contextTimeout,
			}
			if err := a.CheckToken(tt.args.ctx, tt.args.token); (err != nil) != tt.wantErr {
				t.Errorf("CheckToken() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_authUsecase_ConfirmUser(t *testing.T) {
	type fields struct {
		AuthRepo       domain.AuthRepository
		MailRepo       domain.Mail
		contextTimeout time.Duration
	}
	type args struct {
		ctx  context.Context
		auth domain.Auth
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := authUsecase{
				AuthRepo:       tt.fields.AuthRepo,
				MailRepo:       tt.fields.MailRepo,
				contextTimeout: tt.fields.contextTimeout,
			}
			if err := a.ConfirmUser(tt.args.ctx, tt.args.auth); (err != nil) != tt.wantErr {
				t.Errorf("ConfirmUser() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_authUsecase_ResetUserEmailPass(t *testing.T) {
	type fields struct {
		AuthRepo       domain.AuthRepository
		MailRepo       domain.Mail
		contextTimeout time.Duration
	}
	type args struct {
		ctx  context.Context
		auth domain.Auth
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := authUsecase{
				AuthRepo:       tt.fields.AuthRepo,
				MailRepo:       tt.fields.MailRepo,
				contextTimeout: tt.fields.contextTimeout,
			}
			a.ResetUserEmailPass(tt.args.ctx, tt.args.auth)
		})
	}
}

func Test_authUsecase_SignIn(t *testing.T) {
	type fields struct {
		AuthRepo       domain.AuthRepository
		MailRepo       domain.Mail
		contextTimeout time.Duration
	}
	type args struct {
		ctx  context.Context
		auth domain.Auth
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := authUsecase{
				AuthRepo:       tt.fields.AuthRepo,
				MailRepo:       tt.fields.MailRepo,
				contextTimeout: tt.fields.contextTimeout,
			}
			got, err := a.SignIn(tt.args.ctx, tt.args.auth)
			if (err != nil) != tt.wantErr {
				t.Errorf("SignIn() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("SignIn() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_authUsecase_SignUp(t *testing.T) {
	type fields struct {
		AuthRepo       domain.AuthRepository
		MailRepo       domain.Mail
		contextTimeout time.Duration
	}
	type args struct {
		ctx  context.Context
		auth domain.Auth
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := authUsecase{
				AuthRepo:       tt.fields.AuthRepo,
				MailRepo:       tt.fields.MailRepo,
				contextTimeout: tt.fields.contextTimeout,
			}
			if err := a.SignUp(tt.args.ctx, tt.args.auth); (err != nil) != tt.wantErr {
				t.Errorf("SignUp() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
