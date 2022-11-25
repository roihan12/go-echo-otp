package routes

import (
	"go-echo-otp/controllers/users"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type ControllerList struct {
	LoggerMiddleware echo.MiddlewareFunc
	JWTMiddleware    middleware.JWTConfig
	AuthController   users.UserDelivery
}

func (cl *ControllerList) RouteRegister(e *echo.Echo) {
	e.Use(cl.LoggerMiddleware)

	users := e.Group("/api/v1/users")

	users.POST("/register", cl.AuthController.RegisterHandler)
	users.POST("/requestotp", cl.AuthController.RequestOtpHandler)
	users.POST("/login", cl.AuthController.LoginHandler)

}
