package domain

const (
	eTLDUndefined byte = iota // not part of the the public suffix list
	eTLDICANN                 // part of the public suffix list, delegated by ICANN
	eTLDPrivate               // part of the public suffix list, submitted by a domain holder
)
