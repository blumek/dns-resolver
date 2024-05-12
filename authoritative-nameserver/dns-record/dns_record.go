package dns_record

import (
	"github.com/goccy/go-json"
	"time"
)

const (
	A     = "A"
	AAAA  = "AAAA"
	CNAME = "CNAME"
)

type DNSRecord struct {
	DomainName string
	Type_      string
	Value      string
	TimeToLive time.Duration
}

func (dnsRecord DNSRecord) MarshalBinary() ([]byte, error) {
	return json.Marshal(dnsRecord)
}
