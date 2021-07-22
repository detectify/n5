package domain_test

import (
	"testing"

	"github.com/detectify/n5/domain"
	"github.com/stretchr/testify/require"
)

func TestValidateDomainName(t *testing.T) {
	require.NoError(t, domain.Validate("foo.com"))
	require.NoError(t, domain.Validate("FOO.COM"))
	require.NoError(t, domain.Validate("foo.bar.com"))
	require.NoError(t, domain.Validate("foo.longtldname"))
	require.NoError(t, domain.Validate("foo.invalid"))
	require.NoError(t, domain.Validate("foo.com."))
	require.NoError(t, domain.Validate("foo.bar.com."))
	require.NoError(t, domain.Validate("foo-bar.com."))
	require.NoError(t, domain.Validate("foo_bar.com."))
	require.NoError(t, domain.Validate("xn--tst-qla.se"))
	require.NoError(t, domain.Validate("foo"))
	require.NoError(t, domain.Validate("_foo.com"))
	require.NoError(t, domain.Validate("foo_.com"))
	require.NoError(t, domain.Validate("1foo.com"))

	require.Error(t, domain.Validate("foo-.com"))
	require.Error(t, domain.Validate("-foo.com"))
	require.Error(t, domain.Validate("-.foo.com"))
	require.Error(t, domain.Validate("exämple.com"))
	require.Error(t, domain.Validate("foo.com.."))
	require.Error(t, domain.Validate("@foo.com"))
	require.Error(t, domain.Validate("@.foo.com"))
	require.Error(t, domain.Validate("*.foo.com"))
	require.Error(t, domain.Validate("foo!.com"))
	require.Error(t, domain.Validate("aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa.com"))
	require.Error(t, domain.Validate("fo?o.com"))
}
