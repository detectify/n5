package url_test

import (
	"testing"

	"github.com/detectify/n5/url"
	"github.com/stretchr/testify/require"
)

func TestHost_WithURL_ShouldReturnHost(t *testing.T) {
	require.Equal(t, "example.com", url.Host("http://example.com"))
	require.Equal(t, "www.example.com", url.Host("http://www.example.com/index?query=value"))
	require.Equal(t, "www.example.co.uk", url.Host("http://www.example.co.uk:8080/index/inner#location"))
	require.Equal(t, "www.example.com", url.Host("https://user:pass@www.example.com/content.html"))
	require.Equal(t, "www.example.com", url.Host("http://user:pass@www.example.com"))
	require.Equal(t, "10.0.1.1", url.Host("http://10.0.1.1:8080"))
	require.Equal(t, "10.0.1.1", url.Host("http://10.0.1.1:8080/index?query=value"))
	require.Equal(t, "пример.мкд", url.Host("https://пример.мкд/следно.html"))
}

func TestHost_WithMissingScheme_ShouldReturnHost(t *testing.T) {
	require.Equal(t, "example.com", url.Host("example.com"))
	require.Equal(t, "www.example.com", url.Host("user:pass@www.example.com"))
	require.Equal(t, "www.example.com", url.Host("user:pass@www.example.com/content.html"))
	require.Equal(t, "10.0.1.1", url.Host("10.0.1.1"))
	require.Equal(t, "10.0.1.1", url.Host("10.0.1.1:8080"))
}

func TestHost_WithInvalidScheme_ShouldReturnHost(t *testing.T) {
	require.Equal(t, "example.com", url.Host("https:pwn://example.com/"))
	require.Equal(t, "www.example.com", url.Host("https:https://www.example.com/api/v2/admin/content"))
}

func TestHost_WithInvalidPathOrQuery_ShouldReturnHost(t *testing.T) {
	require.Equal(t, "www.example.com", url.Host("https://www.example.com/api/v2/admin/content/?Þ"))
	require.Equal(t, "www.example.com", url.Host("https://www.example.com/api/v2/admin/content/?Þ~Ç\\\\"))
	require.Equal(t, "www.example.com", url.Host("https://www.example.com/api/v2/\\x1fY\\\\x00\\"))
}
