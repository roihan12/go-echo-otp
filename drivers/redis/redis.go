package redis

import (
	"go-echo-otp/utils"
	"log"

	"github.com/go-redis/redis/v8"
)

type ConfigRedis struct {
	REDIS_URL string
}

func (config *ConfigRedis) InitRedis() *redis.Client {

	opt, err := redis.ParseURL(utils.GetConfig("REDIS_URL"))
	if err != nil {
		panic(err)
	}
	rdb := redis.NewClient(opt)

	log.Println("connected to the database redis")

	return rdb
}
