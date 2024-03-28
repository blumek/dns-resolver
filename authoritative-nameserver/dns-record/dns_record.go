package dns_record

import "github.com/goccy/go-json"

const (
	A     = "A"
	AAAA  = "AAAA"
	CNAME = "CNAME"
)

type DNSRecord struct {
	DomainName string
	Type_      string
	Value      string
	TimeToLive int64
}

func (dnsRecord DNSRecord) MarshalBinary() ([]byte, error) {
	return json.Marshal(dnsRecord)
}
