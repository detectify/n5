# n5
Growing collection of network utilities for Go (Golang), such as domain names, IPs and URLs.

[![Go Reference](https://pkg.go.dev/badge/github.com/detectify/n5.svg)](https://pkg.go.dev/github.com/detectify/n5)
![build](https://github.com/detectify/n5/actions/workflows/go.yml/badge.svg)
[![Go Report Card](https://goreportcard.com/badge/github.com/detectify/n5)](https://goreportcard.com/report/github.com/detectify/n5)

## Installation
To install n5, use `go get`:

```bash
go get github.com/detectify/n5
```

## `domain` package
The `domain` package provides functions for managing domain names, which is represented as a sequence of labels 
in the type `domain.Name`. Supports various levels of names (up to TLD), internationalized names and names not on the 
[public suffix list](https://publicsuffix.org).
To create a domain name, parse the string representation, or extract it from a string containing the domain name, such 
as a URL.

```go
import (
    "fmt"
    "github.com/detectify/n5/domain"
)

func main() {
    name, err := domain.Parse("www.example.com")
    fmt.Println("apex = " + name.Apex()) // prints "apex = example.com"
    fmt.Println("tld = " + name.EffectiveTLD()) // prints "tld = com"
    fmt.Println("FQDN = " + name.FQDN()) // prints "tld = www.example.com."
    
    name, err = domain.Extract("https://пример.мкд/index.html")
    fmt.Println("domain (punycoded) = " + name.String()) 
       // prints "domain (punycoded) = xn--e1afmkfd.xn--d1alf"
    fmt.Println("domain (original) = " + name.Unicode()) 
       // prints "domain (original) = пример.мкд"
}
```

### Notes
- Domain name validation is based on the domain name definition specified in [RFC 1034](https://www.ietf.org/rfc/rfc1034.txt), 
  following the [recommended domain name syntax](https://datatracker.ietf.org/doc/html/rfc1034#section-3.5) (reaffirmed 
  in [RFC 2181](https://datatracker.ietf.org/doc/html/rfc2181#section-11)), which is matching the host name 
  definition in [RFC 952](https://datatracker.ietf.org/doc/html/rfc952), extended in 
  [RFC 1123](https://datatracker.ietf.org/doc/html/rfc1123#section-2). Exception is allowing usage of '_' in labels as 
  there are various such names in existence. The name validation is taken from a 
  [Gist](https://gist.github.com/chmike/d4126a3247a6d9a70922fc0e8b4f4013) by [chmike](https://gist.github.com/chmike)
  with added support for `_` character.
- For checking against the public suffix list the [github.com/weppos/publicsuffix-go](https://github.com/weppos/publicsuffix-go) 
  package is used with the [default list](https://pkg.go.dev/github.com/weppos/publicsuffix-go/publicsuffix#pkg-variables).

## `http` package
The `http` package provides a function for validating the HTTP method (`http.ValidateMethod`).

## `ip` package
The `ip` package provides functions for validating IP v4/v6 addresses, including cross-checking against [reserved IPs](https://en.wikipedia.org/wiki/Reserved_IP_addresses).
Separate functions are available for all, v4 and v6 IPs, for example `ip.IsIP`, `ip.IsIPv4`, `ip.IsIPv6`.

## `url` package
The `url` package provides functions for validating absolute URLs (`url.IsAbsolute`) and extracting hostname from URL (`url.Host`).

### Notes
- Host extraction aims to support non-standard URL formats for which the [`url.URL`](https://pkg.go.dev/net/url#URL) type returns error.

## Contributing
Please feel free to submit issues, fork the repository and send pull requests. In addition to fixes, new features are also welcome if you feel they are within the scope of the package. Feel free to reach out and discuss if you have any questions.

## License
This project is licensed under the terms of the MIT license.
