package urlutil

import (
	"net/url"
	"strings"
)

const (
	HTTP             = "http"
	HTTPS            = "https"
	SchemeSeparator  = "://"
	DefaultHTTPPort  = "80"
	DefaultHTTPSPort = "443"
)

func Hostname(s string) (string, error) {
	if !strings.HasPrefix(s, HTTP) && !strings.HasPrefix(s, HTTPS) {
		s = HTTP + SchemeSeparator + s
	}
	u, err := url.Parse(s)
	if err != nil {
		return "", err
	}
	return u.Hostname(), nil
}

func Host(s string) (string, error) {
	if !strings.HasPrefix(s, HTTP) && !strings.HasPrefix(s, HTTPS) {
		s = HTTP + SchemeSeparator + s
	}
	u, err := url.Parse(s)
	if err != nil {
		return "", err
	}
	return u.Host, nil
}
