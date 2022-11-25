package users

import (
	"go-echo-otp/businesses/users"
	"go-echo-otp/controllers/users/request"
	"net/http"

	"github.com/labstack/echo/v4"
)

type UserDelivery struct {
	userUsecase users.UserUsecase
}

func NewUserDelivery(UserUsecase users.UserUsecase) *UserDelivery {
	return &UserDelivery{userUsecase: UserUsecase}
}

func (p *UserDelivery) RequestOtpHandler(c echo.Context) error {
	ctx := c.Request().Context()

	var req request.GetOtpRequest
	if err := c.Bind(&req); err != nil {
		return err
	}

	err := c.Validate(req)
	if err != nil {
		return err
	}

	err = p.userUsecase.RequestOtp(ctx, req.ToDomain())
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"code":   http.StatusOK,
	})

}

func (p *UserDelivery) LoginHandler(c echo.Context) error {
	ctx := c.Request().Context()
	var req request.LoginWithOTPRequest
	if err := c.Bind(&req); err != nil {
		return err
	}

	otp, err := p.userUsecase.LoginWithOTP(ctx, req.ToDomain())
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"code":   http.StatusOK,
		"data":   otp,
	})

}

func (p *UserDelivery) RegisterHandler(c echo.Context) error {
	ctx := c.Request().Context()
	var req request.CreateUserRequest
	if err := c.Bind(&req); err != nil {
		return err
	}

	UserList, err := p.userUsecase.CreateUser(ctx, req.ToDomain())
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, UserList)
}
