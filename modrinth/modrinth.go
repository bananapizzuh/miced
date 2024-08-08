package modrinth

import (
	"net/http"
	"net/url"
	"time"
)

var baseURL = url.URL{
	Scheme: "https",
	Host:   "api.modrinth.com",
	Path:   "/v2/",
}

type Client struct {
	client         *http.Client
	authentication string
	userAgent      string
}

func NewClient(Authentication string, UserAgent string) *Client {
	client := &http.Client{Timeout: time.Minute}

	return &Client{
		client:         client,
		authentication: Authentication,
		userAgent:      UserAgent,
	}
}
