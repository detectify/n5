package domain

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestExtract_WithValidURL_ShouldReturnDomain(t *testing.T) {
	n, err := Extract("https://example.com/index.html")
	require.NoError(t, err)
	require.Equal(t, "example.com", n.String())

	n, err = Extract("https://blog.example.com/index.html")
	require.NoError(t, err)
	require.Equal(t, "blog.example.com", n.String())

	n, err = Extract("https://blog.example.com:8080/index.html")
	require.NoError(t, err)
	require.Equal(t, "blog.example.com", n.String())

	n, err = Extract("https://user:password@blog.example.com:8080/index.html")
	require.NoError(t, err)
	require.Equal(t, "blog.example.com", n.String())
}

func TestExtract_WithValidHostname_ShouldReturnDomain(t *testing.T) {
	n, err := Extract("example.com")
	require.NoError(t, err)
	require.Equal(t, "example.com", n.String())

	n, err = Extract("blog.example.com")
	require.NoError(t, err)
	require.Equal(t, "blog.example.com", n.String())

	n, err = Extract("blog.example.com:8080")
	require.NoError(t, err)
	require.Equal(t, "blog.example.com", n.String())

	n, err = Extract("user:password@blog.example.com:8080")
	require.NoError(t, err)
	require.Equal(t, "blog.example.com", n.String())
}

func TestExtract_WithInvalidString_ShouldReturnError(t *testing.T) {
	_, err := Extract("example@")
	require.Error(t, err)

	_, err = Extract("/index.html")
	require.Error(t, err)

	_, err = Extract("?domain=example.com")
	require.Error(t, err)
}
