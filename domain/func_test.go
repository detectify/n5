package domain_test

import (
	"testing"

	"github.com/detectify/n5/domain"
	"github.com/stretchr/testify/require"
)

func TestContains(t *testing.T) {
	require.True(t, domain.Contains("bar.com", "bar.com"))
	require.True(t, domain.Contains("bar.com", "foo.bar.com"))
	require.True(t, domain.Contains("com", "foo.bar.com"))
	require.True(t, domain.Contains("bar.co.uk", "foo.bar.co.uk"))
	require.True(t, domain.Contains("foo.bar.com", "foo.bar.com"))

	require.False(t, domain.Contains("o.bar.com", "foo.bar.com"))
	require.False(t, domain.Contains("bar.net", "foo.bar.com"))
	require.False(t, domain.Contains("baz.foo.bar.com", "foo.bar.com"))
	require.False(t, domain.Contains("bar.com.net", "foo.bar.com"))
	require.False(t, domain.Contains("foo.bar.com", "bar.com"))
	require.False(t, domain.Contains("foo.bar.com", "com"))
}

func TestIsSubdomainOf(t *testing.T) {
	require.True(t, domain.IsSubdomainOf("foo.bar.com", "bar.com"))
	require.True(t, domain.IsSubdomainOf("foo.bar.com", "com"))
	require.True(t, domain.IsSubdomainOf("foo.bar.co.uk", "bar.co.uk"))

	require.False(t, domain.IsSubdomainOf("foo.bar.com", "foo.bar.com"))
	require.False(t, domain.IsSubdomainOf("foo.bar.com", "o.bar.com"))
	require.False(t, domain.IsSubdomainOf("foo.bar.com", "bar.net"))
	require.False(t, domain.IsSubdomainOf("foo.bar.com", "baz.foo.bar.com"))
	require.False(t, domain.IsSubdomainOf("foo.bar.com", "bar.com.net"))
	require.False(t, domain.IsSubdomainOf("bar.com", "foo.bar.com"))
	require.False(t, domain.IsSubdomainOf("com", "foo.bar.com"))
}

func TestIsImmediateSubdomainOf(t *testing.T) {
	require.True(t, domain.IsImmediateSubdomainOf("foo.bar.com", "bar.com"))
	require.True(t, domain.IsImmediateSubdomainOf("foo.bar.co.uk", "bar.co.uk"))

	require.False(t, domain.IsImmediateSubdomainOf("foo.bar.com", "com"))
	require.False(t, domain.IsImmediateSubdomainOf("foo.bar.com", "foo.bar.com"))
	require.False(t, domain.IsImmediateSubdomainOf("foo.bar.com", "o.bar.com"))
	require.False(t, domain.IsImmediateSubdomainOf("foo.bar.com", "bar.net"))
	require.False(t, domain.IsImmediateSubdomainOf("foo.bar.com", "baz.foo.bar.com"))
	require.False(t, domain.IsImmediateSubdomainOf("foo.bar.com", "bar.com.net"))
	require.False(t, domain.IsImmediateSubdomainOf("bar.com", "foo.bar.com"))
	require.False(t, domain.IsImmediateSubdomainOf("com", "foo.bar.com"))
}

func TestHasPublicSuffix(t *testing.T) {
	// ICANN
	require.True(t, domain.HasPublicSuffix("bar.com"))
	require.True(t, domain.HasPublicSuffix("foo.bar.co.uk"))
	require.True(t, domain.HasPublicSuffix("bar.exchange.aero"))
	require.True(t, domain.HasPublicSuffix("foo.bar.exchange.aero"))
	require.True(t, domain.HasPublicSuffix("foo.bar.exchange.aero"))

	// non-ICANN
	require.True(t, domain.HasPublicSuffix("foo.bar.servehttp.com"))
	require.True(t, domain.HasPublicSuffix("bar.nom.uy"))
	require.True(t, domain.HasPublicSuffix("foo.bar.yombo.me"))

	require.False(t, domain.HasPublicSuffix("com"))
	require.False(t, domain.HasPublicSuffix("co.uk"))
	require.False(t, domain.HasPublicSuffix("uk"))
	require.False(t, domain.HasPublicSuffix("yombo.me"))
}
