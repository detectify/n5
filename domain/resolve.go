package domain

import (
	"fmt"
	"net"

	"github.com/detectify/n5/ip"
)

// Resolve resolves the domain to one of more IP addresses
//
// Checks against reserved IP addresses, and returns an error if the domain name is invalid, or no IPs could be looked
// up, or one or more IPs are reserved.
func Resolve(domain string) ([]net.IP, error) {
	d, err := Parse(domain)
	if err != nil {
		return nil, err
	}

	ips, err := net.LookupIP(d.String())
	if err != nil {
		return nil, fmt.Errorf("failed to lookup IP for domain: %w", err)
	}

	for _, i := range ips {
		if ip.IsReservedIP(i) {
			return nil, fmt.Errorf("domain resolves to reserved IP: %s", i.String())
		}
	}
	return ips, nil
}

// Resolves checks whether the domain can resolve to one or more IP addresses
//
// Checks against reserved IP addresses, and returns an error if the domain name is invalid, or no IPs could be looked
// up, or one or more IPs are reserved.
func Resolves(domain string) bool {
	ips, err := Resolve(domain)
	if err != nil {
		return false
	}
	return len(ips) > 0
}
