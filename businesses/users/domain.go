package users

import (
	"context"
	"go-echo-otp/app/middlewares"
	"time"

	"gorm.io/gorm"
)

type Domain struct {
	ID        int
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
	Name      string
	Email     string
	Password  string
	Otp       string
}

type UserRepository interface {
	Create(ctx context.Context, user *Domain) (Domain, error)
	FindByEmail(ctx context.Context, email string) (Domain, error)
}

type Redis interface {
	Set(ctx context.Context, key string, value interface{}) error
	Get(tx context.Context, key string) (string, error)
}

type EmailRepository interface {
	SendEmail(ctx context.Context, to, subject, message string) error
}

type UserUsecase interface {
	RequestOtp(ctx context.Context, requestOtp *Domain) error
	CreateUser(ctx context.Context, requestUser *Domain) (Domain, error)
	LoginWithOTP(ctx context.Context, requestLogin *Domain) (*middlewares.NewTokenResponse, error)
}
