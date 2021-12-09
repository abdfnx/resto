package validation

import (
	"fmt"
	"net/url"
	"strings"
)

func CheckURL(urlStr string) (string, error) {
	if urlStr == "" {
		return "", fmt.Errorf("URL is needed\n")
	}

	prefixCheck := strings.HasPrefix(urlStr, "http://") || strings.HasPrefix(urlStr, "https://")
	if !prefixCheck {
		return "", fmt.Errorf("URL missing protocol or contains invalid protocol\n")
	}

	if _, err := url.Parse(urlStr); err != nil {
		return "", fmt.Errorf("URL is invalid\n")
	}

	return urlStr, nil
}
