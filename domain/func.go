package domain

// Apex returns the formatted apex domain part of the domain name
//
// Returns empty string if the domain name is empty or invalid or an effective top level domain.
func Apex(domain string) string {
	result, err := Parse(domain)
	if err != nil {
		return ""
	}

	return result.Apex().String()
}

// EffectiveTLD returns the effective top level domain part of the domain name
//
// Returns empty string if the domain name is empty or invalid.
func EffectiveTLD(domain string) string {
	result, err := Parse(domain)
	if err != nil {
		return ""
	}

	return result.EffectiveTLD()
}

// Subdomain returns the subdomain part of the domain name
//
// Returns empty string if the domain name is empty or invalid.
func Subdomain(domain string) string {
	result, err := Parse(domain)
	if err != nil {
		return ""
	}

	return result.Subdomain()
}

// Parent returns the parent domain name
//
// Returns empty string if the domain name is empty or invalid.
func Parent(domain string) string {
	result, err := Parse(domain)
	if err != nil {
		return ""
	}

	return result.Parent().String()
}

// Sanitize performs sanitization of a domain name by converting to IDN format and removing "*." and "@." prefixes
//
// Returns empty string if the domain name is empty or invalid.
func Sanitize(domain string) string {
	result, err := Parse(domain)
	if err != nil {
		return ""
	}

	return result.String()
}

// IsApex returns whether the domain is an apex domain
//
// Returns false if the domain name is empty or invalid.
func IsApex(domain string) bool {
	result, err := Parse(domain)
	if err != nil {
		return false
	}

	return result.IsApex()
}

// IsApexOrSubdomain returns whether the domain is an apex domain or a subdomain
//
// Returns false if the domain name is empty, eTLD or invalid.
func IsApexOrSubdomain(domain string) bool {
	result, err := Parse(domain)
	if err != nil {
		return false
	}

	return result.IsApexOrSubdomain()
}

// IsICANN returns whether the eTLD (public suffix) is managed by the Internet Corporation for Assigned Names and
// Numbers
//
// Returns false if the domain name is empty, eTLD or invalid.
// The complete list of public eTLDs can be found at https://publicsuffix.org/
func IsICANN(domain string) bool {
	result, err := Parse(domain)
	if err != nil {
		return false
	}

	return result.IsICANN()
}

// IsEffectiveTLD returns whether the domain is an effective top level (public suffix) domain
//
// Returns false if the domain name is empty or invalid.
func IsEffectiveTLD(domain string) bool {
	result, err := Parse(domain)
	if err != nil {
		return false
	}

	return result.IsEffectiveTLD()
}

// HasPublicSuffix returns whether the domain name is under the public suffix, under which Internet users can directly
// register names
//
// Returns false if the domain name is not an apex or subdomain, or is invalid.
// The complete list of public eTLDs can be found at https://publicsuffix.org/
func HasPublicSuffix(domain string) bool {
	result, err := Parse(domain)
	if err != nil {
		return false
	}

	return result.HasPublicSuffix()
}

// FQDN returns the fully-qualified domain name
//
// Returns empty string if the domain name is empty or invalid.
func FQDN(domain string) string {
	result, err := Parse(domain)
	if err != nil {
		return ""
	}

	return result.FQDN()
}

// Contains indicates whether this domain matches another domains or is a parent of another domain
//
// Returns false if either domain names are empty or invalid.
func Contains(first, second string) bool {
	firstD, err := Parse(first)
	if err != nil {
		return false
	}
	secondD, err := Parse(second)
	if err != nil {
		return false
	}

	return firstD.Contains(secondD)
}

// Contains indicates whether this domain matches another domain or is a parent of another domain
//
// Returns false if either domain names are empty or invalid.
func IsParentOf(first, second string) bool {
	firstD, err := Parse(first)
	if err != nil {
		return false
	}
	secondD, err := Parse(second)
	if err != nil {
		return false
	}

	return firstD.IsParentOf(secondD)
}

// IsSubdomainOf indicates whether the first domain is a subdomain of the second domain
//
// Returns false if either domain names are empty or invalid.
func IsSubdomainOf(first, second string) bool {
	firstD, err := Parse(first)
	if err != nil {
		return false
	}
	secondD, err := Parse(second)
	if err != nil {
		return false
	}

	return firstD.IsSubdomainOf(secondD)
}

// IsSubdomainOf indicates whether the first domain is an immediate subdomain of the second domain
//
// Returns false if either domain names are empty or invalid.
func IsImmediateSubdomainOf(first, second string) bool {
	firstD, err := Parse(first)
	if err != nil {
		return false
	}
	secondD, err := Parse(second)
	if err != nil {
		return false
	}

	return firstD.IsImmediateSubdomainOf(secondD)
}
