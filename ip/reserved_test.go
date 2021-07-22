package ip_test

import (
	"testing"

	"github.com/detectify/n5/ip"
	"github.com/stretchr/testify/require"
)

func TestIsReserved_WithReservedIPv4_ShouldReturnTrue(t *testing.T) {
	require.True(t, ip.IsReserved("0.0.0.0"))
	require.True(t, ip.IsReserved("0.0.0.100"))
	require.True(t, ip.IsReserved("0.255.255.255"))
	require.True(t, ip.IsReserved("192.168.0.1"))
	require.True(t, ip.IsReserved("224.0.0.0"))
	require.True(t, ip.IsReserved("239.255.255.255"))
	require.True(t, ip.IsReserved("240.0.0.0"))
	require.True(t, ip.IsReserved("255.255.255.254"))
	require.True(t, ip.IsReserved("255.255.255.255"))
}

func TestIsReserved_WithNonReservedIPv4_ShouldReturnFalse(t *testing.T) {
	require.False(t, ip.IsReserved("1.1.1.1"))
	require.False(t, ip.IsReserved("8.8.8.8"))
	require.False(t, ip.IsReserved("1.2.3.4"))
	require.False(t, ip.IsReserved("223.0.0.0"))
	require.False(t, ip.IsReserved("223.255.255.255"))
}

func TestIsReserved_WithMalformedIPv4_ShouldReturnFalse(t *testing.T) {
	require.False(t, ip.IsReserved("127.1"))
	require.False(t, ip.IsReserved("0.0.0.e"))
	require.False(t, ip.IsReserved("0.255.255.256"))
}

func TestIsReserved_WithReservedIPv6_ShouldReturnTrue(t *testing.T) {
	require.True(t, ip.IsReserved("::1"))
	require.True(t, ip.IsReserved("::ffff:0.0.0.0"))
	require.True(t, ip.IsReserved("::ffff:255.255.255.255"))
	require.True(t, ip.IsReserved("2001::"))
	require.True(t, ip.IsReserved("2001::ffff:ffff:ffff:ffff:ffff:ffff"))
	require.True(t, ip.IsReserved("2002::"))
	require.True(t, ip.IsReserved("2002::ffff:ffff:ffff:ffff:ffff:ffff"))
}

func TestIsReserved_WithNonReservedIPv6_ShouldReturnFalse(t *testing.T) {
	require.False(t, ip.IsReserved("1965:0db8:0000:0000:0000:8a2e:0370:7334"))
	require.False(t, ip.IsReserved("2003::"))
}
