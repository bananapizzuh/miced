package modrinth

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"
)

type Project struct {
	ClientSide       string    `json:"client_side"`
	ServerSide       string    `json:"server_side"`
	GameVersions     []string  `json:"game_versions"`
	ID               string    `json:"id"`
	Slug             string    `json:"slug"`
	ProjectType      string    `json:"project_type"`
	Team             string    `json:"team"`
	Organization     any       `json:"organization"`
	Title            string    `json:"title"`
	Description      string    `json:"description"`
	Body             string    `json:"body"`
	BodyURL          any       `json:"body_url"`
	Published        time.Time `json:"published"`
	Updated          time.Time `json:"updated"`
	Approved         time.Time `json:"approved"`
	Queued           any       `json:"queued"`
	Status           string    `json:"status"`
	RequestedStatus  string    `json:"requested_status"`
	ModeratorMessage any       `json:"moderator_message"`
	License          struct {
		ID   string `json:"id"`
		Name string `json:"name"`
		URL  any    `json:"url"`
	} `json:"license"`
	Downloads            int      `json:"downloads"`
	Followers            int      `json:"followers"`
	Categories           []string `json:"categories"`
	AdditionalCategories []any    `json:"additional_categories"`
	Loaders              []string `json:"loaders"`
	Versions             []string `json:"versions"`
	IconURL              string   `json:"icon_url"`
	IssuesURL            string   `json:"issues_url"`
	SourceURL            string   `json:"source_url"`
	WikiURL              string   `json:"wiki_url"`
	DiscordURL           string   `json:"discord_url"`
	DonationUrls         []any    `json:"donation_urls"`
	Gallery              []struct {
		URL         string    `json:"url"`
		Featured    bool      `json:"featured"`
		Title       string    `json:"title"`
		Description any       `json:"description"`
		Created     time.Time `json:"created"`
		Ordering    int       `json:"ordering"`
	} `json:"gallery"`
	Color              int    `json:"color"`
	ThreadID           string `json:"thread_id"`
	MonetizationStatus string `json:"monetization_status"`
}

func (c *Client) GetProjects(ids []string) (string, error) {

	endpt := baseURL.ResolveReference(&url.URL{Path: "projects/"})
	req, err := http.NewRequest(http.MethodGet, endpt.String(), nil)
	if err != nil {
		return "", err
	}

	idsFormatted := fmt.Sprintf(`["%s"]`, strings.Join(ids, `", "`))
	println(idsFormatted)
	params := url.Values{}
	params.Add("ids", idsFormatted)
	req.URL.RawQuery = params.Encode()

	req.Header.Add("Accept", "application/json")
	req.Header.Add("User-Agent", c.userAgent)

	if c.authentication != "" {
		req.Header.Add("Authorization", c.authentication)
	}

	res, err := c.client.Do(req)
	if err != nil {
		return "", err
	}

	defer res.Body.Close()
	buf := new(strings.Builder)
	_, err = io.Copy(buf, res.Body)
	if err != nil {
		return "", err
	}
	return buf.String(), nil
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
