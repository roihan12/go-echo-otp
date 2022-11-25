package reporedis

import (
	"context"
	"go-echo-otp/businesses/users"
	"time"

	"github.com/go-redis/redis/v8"
)

type redisClient struct {
	rdb *redis.Client
}

func NewRedisRepository(rdb *redis.Client) users.Redis {
	return &redisClient{
		rdb: rdb,
	}
}

func (c *redisClient) Set(ctx context.Context, key string, value interface{}) error {
	err := c.rdb.Set(ctx, key, value, 10*time.Minute).Err()
	return err
}

func (c *redisClient) Get(ctx context.Context, key string) (string, error) {
	val, err := c.rdb.Get(ctx, key).Result()
	if err != nil {
		return "", err
	}

	return val, nil
}
