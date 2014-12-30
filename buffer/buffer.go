package buffer

import (
	"encoding/json"
	"net/http"
	"net/url"
)

const (
	ApiURL = "https://api.bufferapp.com"
)

type Client struct {
	AccessToken string
	BaseURL     *url.URL
	client      *http.Client

	Profiles *ProfilesService
	User     *UserService
}

func NewClient(accessToken string, httpClient *http.Client) *Client {
	if httpClient == nil {
		httpClient = http.DefaultClient
	}

	baseURL, _ := url.Parse(ApiURL)

	client := &Client{
		AccessToken: accessToken,
		BaseURL:     baseURL,
		client:      httpClient,
	}

	client.Profiles = &ProfilesService{client: client}
	client.User = &UserService{client: client}

	return client
}

func (c *Client) Do(req *http.Request, v interface{}) (*http.Response, error) {
	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	if v != nil {
		err = json.NewDecoder(resp.Body).Decode(v)
	}

	return resp, err
}

func (c *Client) NewRequest(method, urlStr string, params interface{}) (*http.Request, error) {
	rel, err := url.Parse(urlStr)
	if err != nil {
		return nil, err
	}

	u := c.BaseURL.ResolveReference(rel)
	q := u.Query()
	u.RawQuery = q.Encode()

	req, err := http.NewRequest(method, u.String(), nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Authorization", "Bearer "+c.AccessToken)

	return req, nil
}
