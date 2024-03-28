package api

import (
	usecase "bluemek.com/authoritative_nameserver/use-case"
	"fmt"
	"github.com/gin-gonic/gin"
)

type DNSRecordsWebApi struct {
	getDNSRecordsUseCase *usecase.GetDNSRecordsUseCase
}

func NewDNSRecordsWebApi(getDNSRecordsUseCase *usecase.GetDNSRecordsUseCase) *DNSRecordsWebApi {
	return &DNSRecordsWebApi{getDNSRecordsUseCase: getDNSRecordsUseCase}
}

func (dnsRecordsWebApi *DNSRecordsWebApi) Run() {
	router := gin.Default()
	router.GET("/records/:domain-name", dnsRecordsWebApi.handleGetDNSRecordsForDomain())

	err := router.Run()
	if router.Run() != nil {
		panic(fmt.Sprintf("Couldn't run DNSRecords WEB API as an error occured. %v", err.Error()))
	}
}

func (dnsRecordsWebApi *DNSRecordsWebApi) handleGetDNSRecordsForDomain() func(context *gin.Context) {
	return func(context *gin.Context) {
		dnsName := context.Param("domain-name")
		context.JSON(200, dnsRecordsWebApi.getDNSRecordsUseCase.GetDNSRecordsForDomain(dnsName))
	}
}
