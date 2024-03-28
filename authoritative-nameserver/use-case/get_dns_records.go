package use_case

import (
	. "bluemek.com/authoritative_nameserver/dns-record"
	. "bluemek.com/authoritative_nameserver/repository"
)

type GetDNSRecordsUseCase struct {
	repository DNSRecordsRepository
}

func NewGetDNSRecordsUseCase(recordsRepository DNSRecordsRepository) *GetDNSRecordsUseCase {
	return &GetDNSRecordsUseCase{repository: recordsRepository}
}

func (getDNSRecordsUseCase *GetDNSRecordsUseCase) GetDNSRecordsForDomain(domainName string) []DNSRecord {
	return getDNSRecordsUseCase.repository.FindDNSRecordsForDomain(domainName)
}
