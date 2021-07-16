package domain

import (
	"strings"

	"golang.org/x/net/idna"
)

// RootDomain is the internet root, i.e. "."
//
// Not to be confused with root servers (root-servers.net) or apex domains, which are sometimes referred as root.
var RootDomain = Name{}

// Name holds a structured domain name
type Name struct {
	labels   []string
	category byte
}

// Apex returns the apex domain part of the domain name
func (n Name) Apex() Name {
	if len(n.labels) < 2 {
		// as eTLD or root does not have apex, return root
		// should not have practical impact
		return RootDomain
	}

	return Name{
		labels: n.labels[len(n.labels)-2:],
	}
}

// EffectiveTLD returns the effective top level domain (public suffix) part of the domain name
func (n Name) EffectiveTLD() string {
	if len(n.labels) == 0 {
		return ""
	}

	return n.labels[len(n.labels)-1]
}

// Subdomain returns the subdomain part of the domain name
func (n Name) Subdomain() string {
	if len(n.labels) < 2 {
		return ""
	}

	return strings.Join(n.labels[:len(n.labels)-2], ".")
}

// Parent returns the parent domain name
func (n Name) Parent() Name {
	if len(n.labels) == 0 {
		// for sanity, assume root is the root's parent
		// should not have practical impact
		return n
	}

	return Name{
		labels: n.labels[1:],
	}
}

// IsApex returns whether the domain is an apex domain
func (n Name) IsApex() bool {
	return len(n.labels) == 2
}

// IsApexOrSubdomain returns whether the domain is an apex domain or a subdomain
func (n Name) IsApexOrSubdomain() bool {
	return len(n.labels) > 1
}

// IsICANN returns whether the eTLD (public suffix) is managed by the Internet Corporation for Assigned Names and
// Numbers
//
// The complete list of ICANN eTLDs can be found at https://publicsuffix.org/
func (n Name) IsICANN() bool {
	return n.category == eTLDICANN
}

// IsEffectiveTLD returns whether the domain is an effective top level (public suffix) domain
func (n Name) IsEffectiveTLD() bool {
	return len(n.labels) == 1
}

// HasPublicSuffix returns whether the domain name is under the public suffix, under which Internet users can directly
// register names
//
// The public suffix eTLDs themselves are not considered public, hence this only applies to apex and subdomains.
// The complete list of public eTLDs can be found at https://publicsuffix.org/
func (n Name) HasPublicSuffix() bool {
	return len(n.labels) > 1 && n.category != eTLDUndefined
}

// FQDN returns the fully-qualified domain name
func (n Name) FQDN() string {
	return n.String() + "."
}

// String returns the domain as a string
func (n Name) String() string {
	return strings.Join(n.labels, ".")
}

// Unicode returns the domain in Unicode format
func (n Name) Unicode() string {
	u, _ := idna.ToUnicode(n.String())
	return u
}

// Contains indicates whether this domain matches another domain or is a parent of another domain
func (n Name) Contains(other Name) bool {
	if len(other.labels) == len(n.labels) {
		for i, l := range n.labels {
			if other.labels[i] != l {
				return false
			}
		}
		return true
	}

	return n.IsParentOf(other)
}

// IsParentOf indicates whether this domain is a parent of another domain
func (n Name) IsParentOf(other Name) bool {
	return other.IsSubdomainOf(n)
}

// IsSubdomainOf indicates whether this domain is a subdomain of another domain
func (n Name) IsSubdomainOf(other Name) bool {
	if len(other.labels) == 0 {
		// root domain contains everything
		return true
	}
	if len(other.labels) >= len(n.labels) {
		return false
	}
	for i := len(other.labels) - 1; i >= 0; i-- {
		j := len(n.labels) - len(other.labels) + i
		if other.labels[i] != n.labels[j] {
			return false
		}
	}
	return true
}

// IsImmediateSubdomainOf indicates whether this domain is an immediate subdomain of another domain
func (n Name) IsImmediateSubdomainOf(other Name) bool {
	if len(other.labels) != len(n.labels)-1 {
		return false
	}

	return n.IsSubdomainOf(other)
}
