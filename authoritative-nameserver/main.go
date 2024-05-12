package main

import (
	. "bluemek.com/authoritative_nameserver/api"
	. "bluemek.com/authoritative_nameserver/configuration"
	"bluemek.com/authoritative_nameserver/repository"
	. "bluemek.com/authoritative_nameserver/repository/redis"
	. "bluemek.com/authoritative_nameserver/use-case"
	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"
	"go.uber.org/zap"
	"net/http"
)

func main() {
	fx.New(
		fx.WithLogger(zapLogger()),
		fx.Provide(
			NewHTTPServer,
			UseRoutes(NewGinEngine),
			AsRoute(NewGetDNSRecordsRoute),
			NewGetDNSRecordsUseCase,
			NewRedisDNSRecordsRepository,
			NewRedisClient,
			NewConfiguration,
			zap.NewProduction,
		),
		fx.Invoke(
			startHttpServer(),
			bootstrapRedis(),
		),
	).Run()
}

func UseRoutes(component any) any {
	return fx.Annotate(
		component,
		fx.ParamTags(`group:"routes"`),
	)
}

func AsRoute(component any) any {
	return fx.Annotate(
		component,
		fx.As(new(Route)),
		fx.ResultTags(`group:"routes"`),
	)
}

func zapLogger() func(logger *zap.Logger) fxevent.Logger {
	return func(logger *zap.Logger) fxevent.Logger {
		return &fxevent.ZapLogger{Logger: logger}
	}
}

func bootstrapRedis() func(dnsRecordsRepository repository.DNSRecordsRepository) {
	return func(dnsRecordsRepository repository.DNSRecordsRepository) {
		Bootstrap(dnsRecordsRepository)
	}
}

func startHttpServer() func(*http.Server) {
	return func(*http.Server) {}
}
