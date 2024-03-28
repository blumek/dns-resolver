package main

import (
	. "bluemek.com/authoritative_nameserver/api"
	. "bluemek.com/authoritative_nameserver/configuration"
	. "bluemek.com/authoritative_nameserver/repository/redis"
	. "bluemek.com/authoritative_nameserver/use-case"
)

func main() {
	clientSettings := toRedisSettings(GetConfiguration())
	redisClient := NewRedisClient(clientSettings)
	Bootstrap(redisClient)

	repository := NewRedisDNSRecordsRepository(redisClient)
	getDNSRecordsUseCase := NewGetDNSRecordsUseCase(repository)
	dnsRecordsWebApi := NewDNSRecordsWebApi(getDNSRecordsUseCase)

	dnsRecordsWebApi.Run()
}

func toRedisSettings(configuration Configuration) ClientSettings {
	return ClientSettings{
		IpAddress: configuration.Redis.Host,
		Port:      configuration.Redis.Port,
		Password:  configuration.Redis.Password,
	}
}
