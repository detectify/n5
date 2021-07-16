package url

import (
	"regexp"
)

// (?:(?:\w+:)*\/\/?)* matches on any number of any kind of scheme declaration with single or double slashes, e.g. http://, ssd:ppt:/, blabla://stuff://
// (?:\w+(?::\w+)?@)* matches on userinfo, e.g. user:pass@
// ([^:\/\s]+) will match host until end of string, port number or relative part, e.g. first.second.example.com, 1.1.1.1, пример.мкд
var hostRegex = regexp.MustCompile(`^(?:(?:\w+:)*\/\/?)*(?:\w+(?::\w+)?@)*([^:\/\s]+)`)

// Host returns the host part of the URL
//
// The function is designed to be as generic as possible, working with non-RFC compliant URLs that may occur as a
// result of URL fuzzing. Handles various cases of scheme, userinfo and relative part.
// Can lead to unexpected results with relative URLs.
func Host(u string) string {
	match := hostRegex.FindStringSubmatch(u)
	if len(match) > 1 {
		return match[1]
	}
	return ""
}
