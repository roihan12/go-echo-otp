package request

import "go-echo-otp/businesses/users"

type (
	CreateUserRequest struct {
		Name     string `json:"name"  validate:"required"`
		Email    string `json:"email" validate:"required"`
		Password string `json:"password" validate:"required"`
	}

	GetOtpRequest struct {
		Email string `json:"email" validate:"required"`
	}

	LoginWithOTPRequest struct {
		Email string `json:"email" validate:"required"`
		Otp   string `json:"otp" validate:"required"`
	}
)

func (req *GetOtpRequest) ToDomain() *users.Domain {
	return &users.Domain{
		Email: req.Email,
	}
}

func (req *LoginWithOTPRequest) ToDomain() *users.Domain {
	return &users.Domain{
		Email: req.Email,
		Otp:   req.Otp,
	}
}

func (req *CreateUserRequest) ToDomain() *users.Domain {
	return &users.Domain{
		Name:     req.Name,
		Email:    req.Email,
		Password: req.Password,
	}
}
