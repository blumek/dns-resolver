package api

import (
	usecase "bluemek.com/authoritative_nameserver/use-case"
	"context"
	"errors"
	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
	"go.uber.org/zap"
	_ "go.uber.org/zap"
	"net/http"
	"time"
)

func NewHTTPServer(lifecycle fx.Lifecycle, handler *gin.Engine, logger *zap.Logger) *http.Server {
	server := &http.Server{
		Addr:    ":8080",
		Handler: handler,
	}

	lifecycle.Append(fx.Hook{
		OnStart: handleOnStart(server, logger),
		OnStop:  handleOnStop(server, logger),
	})
	return server
}

func handleOnStart(server *http.Server, logger *zap.Logger) func(ctx context.Context) error {
	return func(ctx context.Context) error {
		logger.Info("An HTTP Server is about to start")
		go func() {
			if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
				logger.Error("An error occurred while listening and serving requests", zap.Error(err))
			}
		}()
		return nil
	}
}

func handleOnStop(server *http.Server, logger *zap.Logger) func(ctx context.Context) error {
	return func(ctx context.Context) error {
		logger.Info("An HTTP Server is about to shutdown")
		return server.Shutdown(ctx)
	}
}

func NewGinEngine(getDNSRecordsUseCase *usecase.GetDNSRecordsUseCase, logger *zap.Logger) *gin.Engine {
	router := gin.New()

	router.Use(ginzap.Ginzap(logger, time.RFC3339, true))
	router.Use(ginzap.RecoveryWithZap(logger, true))

	router.GET("/records/:domain-name", handleGetDNSRecordsForDomain(getDNSRecordsUseCase))
	return router
}

func handleGetDNSRecordsForDomain(getDNSRecordsUseCase *usecase.GetDNSRecordsUseCase) func(context *gin.Context) {
	return func(context *gin.Context) {
		dnsName := context.Param("domain-name")
		context.JSON(200, getDNSRecordsUseCase.GetDNSRecordsForDomain(dnsName))
	}
}
