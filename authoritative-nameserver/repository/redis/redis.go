package redis

import (
	. "bluemek.com/authoritative_nameserver/configuration"
	"fmt"
	"github.com/redis/go-redis/v9"
)

func NewRedisClient(configuration Configuration) *redis.Client {
	return redis.NewClient(toRedisOptions(configuration))
}

func toRedisOptions(configuration Configuration) *redis.Options {
	return &redis.Options{
		Addr:     fmt.Sprintf("%s:%d", configuration.Redis.Host, configuration.Redis.Port),
		Password: configuration.Redis.Password,
		DB:       0,
	}
}
