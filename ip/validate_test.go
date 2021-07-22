package ip_test

import (
	"testing"

	"github.com/detectify/n5/ip"
	"github.com/stretchr/testify/require"
)

func TestIsIPv4_WithIPv4_ShouldReturnTrue(t *testing.T) {
	require.True(t, ip.IsIPv4("1.1.1.1"))
	require.True(t, ip.IsIPv4("192.168.0.1"))
	require.True(t, ip.IsIPv4("255.255.255.255"))
	require.True(t, ip.IsIPv4("0.0.0.0"))
	require.True(t, ip.IsIPv4("127.0.0.1"))

}

func TestIsIPv4_WithIPv6Translated_ShouldReturnTrue(t *testing.T) {
	require.True(t, ip.IsIPv4("::ffff:0.0.0.0"))
	require.True(t, ip.IsIPv4("::ffff:192.168.0.1"))

}

func TestIsIPv4_WithMalformedIPv4_ShouldReturnFalse(t *testing.T) {
	require.False(t, ip.IsIPv4("127.1"))
	require.False(t, ip.IsIPv4("127.0.0.-1"))
	require.False(t, ip.IsIPv4("256.0.0.0"))
	require.False(t, ip.IsIPv4("0.0.1337.0"))
	require.False(t, ip.IsIPv4("0.0.0.e"))
	require.False(t, ip.IsIPv4("localhost"))
	require.False(t, ip.IsIPv4("2001:db8:85a3::8a2e:370:7334"))
}

func TestIsIPv6_WithIPv4_ShouldReturnTrue(t *testing.T) {
	require.True(t, ip.IsIPv6("2001:0db8:0000:0000:0000:8a2e:0370:7334"))
	require.True(t, ip.IsIPv6("2001:db8::8a2e:370:7334"))
	require.True(t, ip.IsIPv6("2002:ffff:ffff:ffff:ffff:ffff:ffff:ffff"))
	require.True(t, ip.IsIPv6("::"))
	require.True(t, ip.IsIPv6("::0"))
	require.True(t, ip.IsIPv6("::1"))
}

func TestIsIPv6_WithIPv4Translated_ShouldReturnFalse(t *testing.T) {
	require.False(t, ip.IsIPv6("::ffff:0.0.0.0"))
	require.False(t, ip.IsIPv6("::ffff:192.168.0.1"))
}

func TestIsIPv6_WithMalformedIPv6_ShouldReturnFalse(t *testing.T) {
	require.False(t, ip.IsIPv6("2001:0db8:0000:fffg:0000:8a2e:0370:7334"))
	require.False(t, ip.IsIPv6("2001:0db8:0000:-1:0000:8a2e:0370:7334"))
	require.False(t, ip.IsIPv6("::ffffff"))
	require.False(t, ip.IsIPv6("::g"))
}
