package drivers

import (
	userDomain "go-echo-otp/businesses/users"
	userDB "go-echo-otp/drivers/mysql/users"
	otpDB "go-echo-otp/drivers/redis/repo_redis"
	mail "go-echo-otp/drivers/thirdparty/mailer"

	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
)

func NewUserRepository(conn *gorm.DB) userDomain.UserRepository {
	return userDB.NewMySQLRepository(conn)
}

func NewOtpRepository(redis *redis.Client) userDomain.Redis {
	return otpDB.NewRedisRepository(redis)
}

func NewSmtpRepository(config mail.Email) userDomain.EmailRepository {
	return mail.NewSmtpEmail(config)
}
