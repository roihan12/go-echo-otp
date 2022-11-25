package main

import (
	_middleware "go-echo-otp/app/middlewares"
	_userUseCase "go-echo-otp/businesses/users"
	_dbMySQL "go-echo-otp/drivers/mysql"
	_dbRedis "go-echo-otp/drivers/redis"
	mail "go-echo-otp/drivers/thirdparty/mailer"
	"go-echo-otp/utils"
	"net/http"

	_routes "go-echo-otp/app/routes"
	_userController "go-echo-otp/controllers/users"

	_driverFactory "go-echo-otp/drivers"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	validatorEngine "github.com/go-playground/validator"
)

func main() {
	configMySQL := _dbMySQL.ConfigDB{
		DB_URL: utils.GetConfig("DATABASE_URL"),
	}

	mysqlDB := configMySQL.InitDB()

	_dbMySQL.DBMigrate(mysqlDB)

	configRedis := _dbRedis.ConfigRedis{
		REDIS_URL: utils.GetConfig("REDIS_URL"),
	}

	configSmtp := mail.Email{
		Host:       utils.GetConfig("SMTP_HOST"),
		Port:       utils.GetInt("SMTP_PORT"),
		SenderName: utils.GetConfig("SMTP_SENDER_NAME"),
		AuthEmail:  utils.GetConfig("SMTP_EMAIL"),
		Password:   utils.GetConfig("SMTP_PASSWORD"),
	}

	// TODO USE REDIS CLIENT
	redisDb := configRedis.InitRedis()

	configLogger := _middleware.ConfigLogger{
		Format: "[${time_rfc3339}] ${status} ${method} ${host} ${path} ${latency_human}" + "\n",
	}

	e := echo.New()

	e.Validator = &utils.GoPlaygroundValidator{
		Validator: validatorEngine.New(),
	}

	e.Use(
		middleware.Recover(),   // Recover from all panics to always have your server up
		middleware.Logger(),    // Log everything to stdout
		middleware.RequestID(), // Generate a request id on the HTTP response headers for identification
	)
	e.Debug = false
	e.HideBanner = true
	// e.HTTPErrorHandler = func(err error, c echo.Context) {
	// 	// Take required information from error and context and send it to a service like New Relic
	// 	fmt.Println(c.Path(), c.QueryParams(), err.Error())

	// 	// Call the default handler to return the HTTP response
	// 	e.DefaultHTTPErrorHandler(err, c)
	// }

	e.HTTPErrorHandler = func(err error, c echo.Context) {
		report, ok := err.(*echo.HTTPError)
		if !ok {
			report = echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}

		c.Logger().Error(report)
		c.JSON(report.Code, report)
	}

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	userRepo := _driverFactory.NewUserRepository(mysqlDB)
	redisRepo := _driverFactory.NewOtpRepository(redisDb)

	mailRepo := _driverFactory.NewSmtpRepository(configSmtp)

	userUsecase := _userUseCase.NewUserUsecase(redisRepo, userRepo, mailRepo)

	userCtrl := _userController.NewUserDelivery(userUsecase)

	routesInit := _routes.ControllerList{
		LoggerMiddleware: configLogger.Init(),
		AuthController:   *userCtrl,
	}

	routesInit.RouteRegister(e)

	e.Logger.Fatal(e.Start(":1323"))
}
