package modrinth

import (
	"encoding/json"
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

type Project struct {
	ID    string `json:"id"`
	Title string `json:"title"`
}

func (c *Client) GetProject(id string) (*Project, error) {
	endpt := baseURL.ResolveReference(&url.URL{Path: "project/" + id})
	req, err := http.NewRequest(http.MethodGet, endpt.String(), nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Accept", "application/json")
	req.Header.Add("User-Agent", c.userAgent)
	if c.authentication != "" {
		req.Header.Add("Authorization", c.authentication)
	}

	res, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()
	var project Project
	if err := json.NewDecoder(res.Body).Decode(&project); err != nil {
		return nil, err
	}
	return &project, nil
}
