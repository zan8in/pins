package urlutil

import (
	"fmt"
	"net/url"
	"strings"

	iputil "github.com/zan8in/pins/ip"
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

func Domain(s string) (string, error) {
	if !strings.HasPrefix(s, HTTP) && !strings.HasPrefix(s, HTTPS) {
		s = HTTP + SchemeSeparator + s
	}
	u, err := url.Parse(s)
	if err != nil {
		return "", err
	}

	host := u.Hostname()

	if iputil.IsIP(host) {
		return host, nil
	}

	parts := strings.Split(host, ".")
	if len(parts) >= 2 {
		mainDomain := parts[len(parts)-2] + "." + parts[len(parts)-1]
		return mainDomain, nil
	}

	return "", fmt.Errorf("无法获取主域名")
}

// URL Encode all characters
// URLEncodeAllChar("1' and sleep(5) and '1'='1")
// %31%27%20%61%6E%64%20%73%6C%65%65%70%28%35%29%20%61%6E%64%20%27%31%27%3D%27%31
func URLEncodeAllChar(s string) string {
	b := []byte(s)
	nb := ""
	for _, v := range b {
		nb += "%" + fmt.Sprintf("%X", v)
	}
	return nb
}
