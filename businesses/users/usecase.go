package users

import (
	"context"
	"errors"
	"go-echo-otp/app/middlewares"
	"go-echo-otp/utils"
)

type userUsecase struct {
	mailRepo  EmailRepository
	redisRepo Redis
	userRepo  UserRepository
}

func NewUserUsecase(redisRepo Redis, user UserRepository, mailRepo EmailRepository) UserUsecase {
	return &userUsecase{
		mailRepo:  mailRepo,
		redisRepo: redisRepo,
		userRepo:  user,
	}
}

func (u *userUsecase) CreateUser(ctx context.Context, request *Domain) (Domain, error) {
	hashPassword, _ := utils.NewPassword(request.Password)

	user, err := u.userRepo.Create(ctx, &Domain{
		Name:     request.Name,
		Email:    request.Email,
		Password: hashPassword,
	})
	if err != nil {
		return *request, err
	}

	to := request.Email
	subject := " Akun Profcouse"
	message := "" +
		"<img src=\"https://firebasestorage.googleapis.com/v0/b/crudfirebase-91413.appspot.com/o/logo.png?alt=media&token=4fa0b90f-6b13-41f3-96a3-53e277ff4d5c\" alt=\"Logo Prof Course\" width=\"75\">" +
		"<p>Dear " + request.Name + "</p><br><p> password anda sekarang adalah : " + request.Password + " " + "" +
		"<p>Anda harus menjaga informasi anda</p>" +
		"<br>" +
		"<p>Terima kasih</p>" +
		"<br>" +
		"Prof Course"

	//send email
	go u.mailRepo.SendEmail(ctx, to, subject, message)

	return user, nil
}

func (u *userUsecase) RequestOtp(ctx context.Context, request *Domain) error {
	// get email
	user, err := u.userRepo.FindByEmail(ctx, request.Email)
	if err != nil {
		return err
	}

	otp := utils.RandNumeric()
	cacheKey := "otp:" + user.Email
	err = u.redisRepo.Set(ctx, cacheKey, otp)
	if err != nil {
		return err
	}

	to := user.Email
	subject := "Login Akun Profcouse"
	message := "" +
		"<img src=\"\" alt=\"Logo Eduworld Course\" width=\"75\">" +
		"<p>Dear " + user.Name + "</p><br><p> Berikut ini adalah OTP Anda : " + otp + " " + "" +
		"<p>Anda harus menjaga informasi anda</p>" +
		"<br>" +
		"<p>Terima kasih</p>" +
		"<br>" +
		"Eduworld Course"

	//send email
	go u.mailRepo.SendEmail(ctx, to, subject, message)

	return nil

}

func (u *userUsecase) LoginWithOTP(ctx context.Context, request *Domain) (*middlewares.NewTokenResponse, error) {
	cacheKey := "otp:" + request.Email

	otpValue, _ := u.redisRepo.Get(ctx, cacheKey)

	if request.Otp != otpValue {
		return nil, errors.New("otp is not valid")
	}

	user, err := u.userRepo.FindByEmail(ctx, request.Email)
	if err != nil {
		return nil, err
	}

	token, err := middlewares.NewCustomToken(middlewares.NewTokenRequest{
		UserID:    user.ID,
		UserEmail: user.Email,
	}, middlewares.DurationShort)
	if err != nil {
		return nil, err
	}

	return token, nil
}
