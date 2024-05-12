package redis

import (
	dnsrecord "bluemek.com/authoritative_nameserver/dns-record"
	"bluemek.com/authoritative_nameserver/repository"
	"context"
	"errors"
	"fmt"
	"github.com/goccy/go-json"
	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
)

type DNSRecordsRepository struct {
	redisClient *redis.Client
	logger      *zap.Logger
}

func NewRedisDNSRecordsRepository(redisClient *redis.Client) repository.DNSRecordsRepository {
	return &DNSRecordsRepository{redisClient: redisClient}
}

func (redisDNSRecordsRepository *DNSRecordsRepository) FindDNSRecordsForDomain(domainName string) []dnsrecord.DNSRecord {
	keys := redisDNSRecordsRepository.keysFor(domainName)
	return redisDNSRecordsRepository.retrieveDNSRecordsForKeys(keys)
}

func (redisDNSRecordsRepository *DNSRecordsRepository) retrieveDNSRecordsForKeys(keys []string) []dnsrecord.DNSRecord {
	dnsRecords := make([]dnsrecord.DNSRecord, 0)
	if len(keys) == 0 {
		return dnsRecords
	}

	dnsRecordJsons, retrieveDnsRecordsError := redisDNSRecordsRepository.redisClient.MGet(context.Background(), keys...).Result()
	if retrieveDnsRecordsError != nil {
		redisDNSRecordsRepository.logger.Error(
			"An error occurred while retrieving DNS records",
			zap.Error(retrieveDnsRecordsError),
		)
		return dnsRecords
	}

	for _, dnsRecordJson := range dnsRecordJsons {
		dnsRecord, _ := redisDNSRecordsRepository.parseDNSRecordJson(dnsRecordJson.(string))
		dnsRecords = append(dnsRecords, dnsRecord)
	}

	return dnsRecords
}

func (redisDNSRecordsRepository *DNSRecordsRepository) keysFor(domainName string) []string {
	keys := make([]string, 0)

	keysPrefix := fmt.Sprintf("%s|*", domainName)
	keysIterator := redisDNSRecordsRepository.redisClient.Scan(context.Background(), 0, keysPrefix, 0).Iterator()
	for keysIterator.Next(context.Background()) {
		keys = append(keys, keysIterator.Val())
	}

	if keysIteratorError := keysIterator.Err(); keysIteratorError != nil {
		redisDNSRecordsRepository.logger.Error(
			"An error occurred while retrieving keys for DNS record",
			zap.String("domainName", domainName),
			zap.Error(keysIteratorError),
		)
	}

	return keys
}

func (redisDNSRecordsRepository *DNSRecordsRepository) parseDNSRecordJson(dnsRecordJson string) (dnsrecord.DNSRecord, error) {
	var dnsRecord dnsrecord.DNSRecord
	unmarshallingError := json.Unmarshal([]byte(dnsRecordJson), &dnsRecord)
	if unmarshallingError != nil {
		return dnsRecord, errors.New("")
	}
	return dnsRecord, nil
}

func (redisDNSRecordsRepository *DNSRecordsRepository) Save(dnsRecord dnsrecord.DNSRecord) error {
	dnsRecordJson, marshalingError := json.Marshal(dnsRecord)
	if marshalingError != nil {
		redisDNSRecordsRepository.logger.Error(
			"An error occurred while marshaling DNS record to json",
			zap.String("domainName", dnsRecord.DomainName),
			zap.Error(marshalingError),
		)
		return marshalingError
	}

	dnsRecordIdentifier := fmt.Sprintf("%s|%s", dnsRecord.DomainName, uuid.NewString())
	settingError := redisDNSRecordsRepository.redisClient.Set(context.Background(), dnsRecordIdentifier, dnsRecordJson, dnsRecord.TimeToLive).Err()
	if settingError != nil {
		redisDNSRecordsRepository.logger.Error(
			"An error occurred while saving DNS record",
			zap.String("dnsRecordIdentifier", dnsRecordIdentifier),
			zap.String("dnsRecord", string(dnsRecordJson)),
			zap.Error(settingError),
		)
		return settingError
	}

	return nil
}

func Bootstrap(dnsRecordsRepository repository.DNSRecordsRepository) {
	dnsRecord := dnsrecord.DNSRecord{
		DomainName: "google.com",
		Type_:      dnsrecord.A,
		Value:      "192.0.2.1",
		TimeToLive: 14400,
	}

	err := dnsRecordsRepository.Save(dnsRecord)
	if err != nil {
		panic(err)
	}
}
