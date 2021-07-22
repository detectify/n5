package ip

import (
	"fmt"
	"net"
	"strings"
)

// Validate indicates whether the specified string is a valid IP address
func Validate(s string) error {
	ip := net.ParseIP(s)
	if ip == nil {
		return fmt.Errorf("%s is not an IP address", s)
	}
	return nil
}

// ValidateNonReserved indicates whether the specified string is a valid non-reserved IP address
func ValidateNonReserved(s string) error {
	ip := net.ParseIP(s)
	if ip == nil {
		return fmt.Errorf("%s is not an IP address", s)
	}
	if IsReservedIP(ip) {
		return fmt.Errorf("%s is a reserved IP address", s)
	}
	return nil
}

// IsIP indicates whether the specified string is an IP address
func IsIP(s string) bool {
	ip := net.ParseIP(s)
	return ip != nil
}

// IsIPv4 indicates whether the specified string is an IP v4 address
func IsIPv4(s string) bool {
	ip := net.ParseIP(s)
	return ip != nil && strings.Contains(ip.String(), ".")
}

// IsIPv6 indicates whether the specified string is an IP v6 address
func IsIPv6(s string) bool {
	ip := net.ParseIP(s)
	return ip != nil && strings.Contains(ip.String(), ":")
}
