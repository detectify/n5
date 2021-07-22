package validate

import "github.com/detectify/n5/ip"

// IP determines whether a string is a valid IP address
func IP(s string) error {
	return ip.Validate(s)
}

// NonReservedIP determines whether a string is a valid, non-reserved IP address
func NonReservedIP(s string) error {
	return ip.ValidateNonReserved(s)
}
