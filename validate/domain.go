package validate

import (
	"github.com/detectify/n5/domain"
)

// DomainName determines whether a string is a valid domain name according to RFC 1034 and 3696
func DomainName(s string) error {
	return domain.Validate(s)
}
