package url

import (
	"net/url"
)

// IsAbsolute indicates whether the specified string is an absolute URL
func IsAbsolute(u string) bool {
	parsed, err := url.Parse(u)
	return err == nil && IsAbsoluteURL(*parsed)
}

// IsAbsoluteURL indicates whether the specified URL is an absolute URL
func IsAbsoluteURL(u url.URL) bool {
	return len(u.Host) > 0 && len(u.Scheme) > 0
}
