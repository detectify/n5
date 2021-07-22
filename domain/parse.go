package domain

import (
	"fmt"
	"strings"

	"github.com/weppos/publicsuffix-go/publicsuffix"
	"golang.org/x/net/idna"
)

// Parse parses the specified domain name and returns it in structured format
//
// Removes wildcard ("*." and "@.") prefixes, and converts the name to IDN format.
// Returns an error if the domain name is empty or invalid.
func Parse(s string) (Name, error) {
	formattedName := strings.Trim(strings.ToLower(s), ".")

	if strings.HasPrefix(formattedName, "*.") {
		formattedName = strings.Replace(formattedName, "*.", "", 1)
	}
	if strings.HasPrefix(formattedName, "@.") {
		formattedName = strings.Replace(formattedName, "@.", "", 1)
	}

	if len(formattedName) == 0 {
		return Name{}, fmt.Errorf("domain name is empty")
	}

	var err error
	formattedName, err = idna.ToASCII(formattedName)
	if err != nil {
		return Name{}, fmt.Errorf("domain name %s is invalid: %w", s, err)
	}

	if err = Validate(formattedName); err != nil {
		return Name{}, fmt.Errorf("domain name %s is invalid: %w", s, err)
	}

	rule := publicsuffix.DefaultList.Find(formattedName, publicsuffix.DefaultFindOptions)
	if rule == nil {
		return Name{}, fmt.Errorf("domain name %s is invalid: no rule found", s)
	}

	category := eTLDUndefined
	if rule.Private {
		category = eTLDPrivate
	} else if len(rule.Value) > 0 {
		// empty value indicates the default rule
		category = eTLDICANN
	}

	decomposedName := rule.Decompose(formattedName)
	if decomposedName[1] == "" {
		// no TLD found, which means it's already a TLD
		return Name{
			labels:   []string{formattedName},
			category: category,
		}, nil
	}

	labelsNoTDL := strings.TrimSuffix(formattedName, decomposedName[1])
	labelsNoTDL = strings.TrimSuffix(labelsNoTDL, ".")

	if len(labelsNoTDL) == 0 {
		return Name{
			labels:   []string{decomposedName[1]},
			category: category,
		}, nil
	}

	return Name{
		labels:   append(strings.Split(labelsNoTDL, "."), decomposedName[1]),
		category: category,
	}, nil
}

// MustParse parses the specified domain name and returns it in structured format
//
// Removes wildcard ("*." and "@.") prefixes, and converts the name to IDN format.
// Panics if the domain name is empty or invalid.
func MustParse(s string) Name {
	result, err := Parse(s)
	if err != nil {
		panic(err)
	}
	return result
}
