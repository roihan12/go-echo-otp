package users

import (
	"context"
	"go-echo-otp/businesses/users"

	"gorm.io/gorm"
)

type userRepository struct {
	conn *gorm.DB
}

func NewMySQLRepository(conn *gorm.DB) users.UserRepository {
	return &userRepository{
		conn: conn,
	}
}

func (u *userRepository) Create(ctx context.Context, user *users.Domain) (users.Domain, error) {

	rec := FromDomain(user)

	if err := u.conn.
		WithContext(ctx).
		Create(&rec).Error; err != nil {
		return rec.ToDomain(), err
	}
	return rec.ToDomain(), nil
}

func (u *userRepository) FindByEmail(ctx context.Context, email string) (users.Domain, error) {
	var user User

	if err := u.conn.WithContext(ctx).
		Where("email = ?", email).First(&user).Error; err != nil {
		return users.Domain{}, err
	}

	return user.ToDomain(), nil

}
