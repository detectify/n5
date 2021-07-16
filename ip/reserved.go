package ip

import (
	"net"
)

// List of reserved IP address ranges based on https://en.wikipedia.org/wiki/Reserved_IP_addresses
var (
	// ReservedIPv4Ranges holds the reserved IPv4 ranges
	ReservedIPv4Ranges []net.IPNet

	// ReservedIPv4RangeStrings holds the reserved IPv4 ranges in string format
	ReservedIPv4RangeStrings = []string{
		"0.0.0.0/8",
		"10.0.0.0/8",
		"100.64.0.0/10",
		"127.0.0.0/8",
		"169.254.0.0/16",
		"172.16.0.0/12",
		"192.0.0.0/24",
		"192.0.2.0/24",
		"192.88.99.0/24",
		"192.168.0.0/16",
		"198.18.0.0/15",
		"198.51.100.0/24",
		"203.0.113.0/24",
		"224.0.0.0/4",
		"240.0.0.0/4",
		"255.255.255.255/32",
	}

	// ReservedIPv6Ranges holds the reserved IPv6 ranges
	ReservedIPv6Ranges []net.IPNet

	// ReservedIPv6RangeStrings holds the reserved IPv6 ranges in string format
	ReservedIPv6RangeStrings = []string{
		"::1/128",
		"::ffff:0:0:0/96",
		"64:ff9b::/96",
		"100::/64",
		"2001::/32",
		"2001:20::/28",
		"2001:db8::/32",
		"2002::/16",
		"fc00::/7",
		"fe80::/10",
		"ff00::/8",
	}
)

func init() {
	for _, r := range ReservedIPv4RangeStrings {
		_, ipNet, _ := net.ParseCIDR(r)
		ReservedIPv4Ranges = append(ReservedIPv4Ranges, *ipNet)
	}
	for _, r := range ReservedIPv6RangeStrings {
		_, ipNet, _ := net.ParseCIDR(r)
		ReservedIPv6Ranges = append(ReservedIPv6Ranges, *ipNet)
	}
}

// IsReserved checks if the specified IP address is reserved
func IsReserved(s string) bool {
	return IsReservedIP(net.ParseIP(s))
}

// IsReservedIP checks if the specified IP address is reserved
func IsReservedIP(ip net.IP) bool {
	if IsReservedIPv4(ip) {
		return true
	}
	if IsReservedIPv6(ip) {
		return true
	}
	return false
}

// IsReservedIPv4 checks if the specified IP address is reserved
func IsReservedIPv4(ip net.IP) bool {
	if ip == nil {
		return false
	}

	for _, r := range ReservedIPv4Ranges {
		if r.Contains(ip) {
			return true
		}
	}
	return false
}

// IsReservedIPv6 checks if the specified IP address is reserved
func IsReservedIPv6(ip net.IP) bool {
	if ip == nil {
		return false
	}

	for _, r := range ReservedIPv6Ranges {
		if r.Contains(ip) {
			return true
		}
	}
	return false
}
