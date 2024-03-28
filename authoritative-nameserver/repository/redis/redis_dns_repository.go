package redis

import (
	dnsrecord "bluemek.com/authoritative_nameserver/dns-record"
	"bluemek.com/authoritative_nameserver/repository"
	"context"
	"github.com/goccy/go-json"
	"github.com/redis/go-redis/v9"
)

type DNSRecords struct {
	DNSRecords []dnsrecord.DNSRecord
}

func (dnsRecords DNSRecords) MarshalBinary() ([]byte, error) {
	return json.Marshal(dnsRecords)
}

type DNSRecordsRepository struct {
	redisClient *redis.Client
}

func NewRedisDNSRecordsRepository(redisClient *redis.Client) repository.DNSRecordsRepository {
	return &DNSRecordsRepository{redisClient: redisClient}
}

func (redisDNSRecordsRepository *DNSRecordsRepository) FindDNSRecordsForDomain(domainName string) []dnsrecord.DNSRecord {
	retrievedDNSRecords, _ := redisDNSRecordsRepository.redisClient.Get(context.Background(), domainName).Result()
	return parseDNSRecordsJson(retrievedDNSRecords).DNSRecords
}

func parseDNSRecordsJson(dnsRecordsJson string) DNSRecords {
	var dnsRecords DNSRecords
	err := json.Unmarshal([]byte(dnsRecordsJson), &dnsRecords)
	if err != nil {
		return DNSRecords{}
	}
	return dnsRecords
}
