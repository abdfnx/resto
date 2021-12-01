package client

import (
	"net/http"
	"time"
)

func httpClient() *http.Client {
	return &http.Client{
		Timeout: 10 * time.Second,
	}
}
