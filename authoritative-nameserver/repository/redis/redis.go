package redis

import (
	"fmt"
	"github.com/redis/go-redis/v9"
)

type ClientSettings struct {
	IpAddress string
	Port      uint16
	Password  string
}

func redisDefaultClientSettings() ClientSettings {
	return ClientSettings{
		IpAddress: "localhost",
		Port:      6379,
		Password:  "",
	}
}

func NewRedisClient(settings ClientSettings) *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", settings.IpAddress, settings.Port),
		Password: settings.Password,
		DB:       0,
	})
}

func NewDefaultRedisClient() *redis.Client {
	return NewRedisClient(redisDefaultClientSettings())
}
