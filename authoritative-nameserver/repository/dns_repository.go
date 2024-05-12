package repository

import . "bluemek.com/authoritative_nameserver/dns-record"

type DNSRecordsRepository interface {
	FindDNSRecordsForDomain(domainName string) []DNSRecord
	Save(dnsRecord DNSRecord) error
}
