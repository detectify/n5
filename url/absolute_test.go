package url_test

import (
	"testing"

	"github.com/detectify/n5/url"
	"github.com/stretchr/testify/require"
)

func TestIsAbsolute_WithValidURLs_ShouldReturnTrue(t *testing.T) {
	require.True(t, url.IsAbsolute("https://example.com"))
	require.True(t, url.IsAbsolute("http://blog.example.com/relative/path?query=value#fragment"))
	require.True(t, url.IsAbsolute("somescheme://example.nonexistent.tld"))
}

func TestIsAbsolute_WithInvalidURLs_ShouldReturnFalse(t *testing.T) {
	require.False(t, url.IsAbsolute(""))
	require.False(t, url.IsAbsolute("https//example.com"))
	require.False(t, url.IsAbsolute("/relative/path"))
	require.False(t, url.IsAbsolute("example.com"))
}
