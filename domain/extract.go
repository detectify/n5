package domain

import "strings"

// Extract extracts a domain name from the specified string
//
// Supports hostname, hostname with port number, origins and URLs
func Extract(s string) (Name, error) {
	// strip schema
	i := strings.Index(s, "://")
	if i >= 0 {
		s = s[i+3:]
	}
	// strip authentication
	i = strings.Index(s, "@")
	if i >= 0 {
		s = s[i+1:]
	}
	// strip relative path
	i = strings.Index(s, "/")
	if i >= 0 {
		s = s[:i]
	}
	// strip port
	i = strings.Index(s, ":")
	if i >= 0 {
		s = s[:i]
	}

	return Parse(s)
}
