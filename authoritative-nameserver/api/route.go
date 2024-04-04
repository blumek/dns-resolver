package api

import (
	usecase "bluemek.com/authoritative_nameserver/use-case"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
)

type Route interface {
	HttpMethod() string
	Pattern() string
	Handler() gin.HandlerFunc
}

type GetDNSRecordsRoute struct {
	getDNSRecordsUseCase *usecase.GetDNSRecordsUseCase
	logger               *zap.Logger
}

func NewGetDNSRecordsRoute(getDNSRecordsUseCase *usecase.GetDNSRecordsUseCase, logger *zap.Logger) *GetDNSRecordsRoute {
	return &GetDNSRecordsRoute{
		getDNSRecordsUseCase: getDNSRecordsUseCase,
		logger:               logger,
	}
}

func (*GetDNSRecordsRoute) HttpMethod() string {
	return http.MethodGet
}

func (*GetDNSRecordsRoute) Pattern() string {
	return "/records/:domain-name"
}

func (getDNSRecordsRoute *GetDNSRecordsRoute) Handler() gin.HandlerFunc {
	return func(context *gin.Context) {
		dnsName := context.Param("domain-name")
		context.JSON(200, getDNSRecordsRoute.getDNSRecordsUseCase.GetDNSRecordsForDomain(dnsName))
	}
}
