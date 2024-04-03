package main

import (
	. "bluemek.com/authoritative_nameserver/api"
	. "bluemek.com/authoritative_nameserver/configuration"
	. "bluemek.com/authoritative_nameserver/repository/redis"
	. "bluemek.com/authoritative_nameserver/use-case"
	"github.com/redis/go-redis/v9"
	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"
	"go.uber.org/zap"
	"net/http"
)

func main() {
	fx.New(
		fx.Provide(NewHTTPServer),
		fx.Provide(NewGinEngine),
		fx.Provide(NewGetDNSRecordsUseCase),
		fx.Provide(NewRedisDNSRecordsRepository),
		fx.Provide(NewRedisClient),
		fx.Provide(NewConfiguration),
		fx.Provide(zap.NewProduction()),
		fx.WithLogger(zapLogger()),
		fx.Invoke(startHttpServer()),
		fx.Invoke(bootstrapRedis()),
	).Run()
}

func zapLogger() func(log *zap.Logger) fxevent.Logger {
	return func(log *zap.Logger) fxevent.Logger {
		return &fxevent.ZapLogger{Logger: log}
	}
}

func bootstrapRedis() func(client *redis.Client) {
	return func(client *redis.Client) {
		Bootstrap(client)
	}
}

func startHttpServer() func(*http.Server) {
	return func(*http.Server) {}
}
