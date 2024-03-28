package main

import (
	dnsrecord "bluemek.com/authoritative_nameserver/dns-record"
	redis_ "bluemek.com/authoritative_nameserver/repository/redis"
	"context"
	"github.com/redis/go-redis/v9"
)

func Bootstrap(client *redis.Client) {
	dnsRecords := []dnsrecord.DNSRecord{
		{
			DomainName: "google.com",
			Type_:      dnsrecord.A,
			Value:      "192.0.2.1",
			TimeToLive: 14400,
		},
	}

	err := client.Set(context.Background(), "google.com", redis_.DNSRecords{DNSRecords: dnsRecords}, 0).Err()
	if err != nil {
		panic(err)
	}
}
