package domain

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

// Validate determines whether a string is a valid domain name
//
// Validation is based on the domain name definition specified in RFC 1034, following the recommended domain name
// syntax, which is matching the host name definition in RFC 952, extended in RFC 1123. Exception is allowing usage of
// '_' in labels as there are various such names in existence.
func Validate(s string) error {
	// source: https://gist.github.com/chmike/d4126a3247a6d9a70922fc0e8b4f4013
	name := strings.TrimSuffix(s, ".")
	switch {
	case len(name) == 0:
		return nil
	case len(name) > 255:
		return fmt.Errorf("name length is %d, can't exceed 255", len(name))
	}
	var l int
	for i := 0; i < len(name); i++ {
		b := name[i]
		if b == '.' {
			// check domain labels validity
			switch {
			case i == l:
				return fmt.Errorf("invalid character '%c' at offset %d: label can't begin with a period", b, i)
			case i-l > 63:
				return fmt.Errorf("byte length of label '%s' is %d, can't exceed 63", name[l:i], i-l)
			case name[l] == '-':
				return fmt.Errorf("label '%s' at offset %d begins with a hyphen", name[l:i], l)
			case name[i-1] == '-':
				return fmt.Errorf("label '%s' at offset %d ends with a hyphen", name[l:i], l)
			}
			l = i + 1
			continue
		}
		if !(b >= 'a' && b <= 'z' || b >= '0' && b <= '9' || b == '-' || b >= 'A' && b <= 'Z' || b == '_') {
			c, _ := utf8.DecodeRuneInString(name[i:])
			if c == utf8.RuneError {
				return fmt.Errorf("invalid rune at offset %d", i)
			}
			return fmt.Errorf("invalid character '%c' at offset %d", c, i)
		}
	}
	switch {
	case l == len(name):
		return fmt.Errorf("missing top level domain, domain can't end with a period")
	case len(name)-l > 63:
		return fmt.Errorf("byte length of top level domain '%s' is %d, can't exceed 63", name[l:], len(name)-l)
	case name[l] == '-':
		return fmt.Errorf("top level domain '%s' at offset %d begins with a hyphen", name[l:], l)
	case name[len(name)-1] == '-':
		return fmt.Errorf("top level domain '%s' at offset %d ends with a hyphen", name[l:], l)
	case name[l] >= '0' && name[l] <= '9':
		return fmt.Errorf("top level domain '%s' at offset %d begins with a digit", name[l:], l)
	}
	return nil
}

// IsDomainName determines whether a string is a valid domain name
func IsDomainName(s string) bool {
	return Validate(s) == nil
}
