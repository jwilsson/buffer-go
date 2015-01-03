package buffer

import (
	"fmt"
	"net/url"
)

type LinksService struct {
	client *Client
}

type Link struct {
	Shares int `json:"shares,omitempty"`
}

func (s *LinksService) GetShares(urlStr string) (int, error) {
	u := fmt.Sprintf("/1/links/shares.json?url=%v", url.QueryEscape(urlStr))

	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return -1, err
	}

	link := new(Link)
	_, err = s.client.Do(req, link)

	return link.Shares, err
}
