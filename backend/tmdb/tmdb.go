package tmdb

import (
	"fmt"
	"net/http"
	"net/url"
)

type Client struct {
	baseURL string
	key     string
}

func NewClient(key string) *Client {
	return &Client{
		key:     key,
		baseURL: "https://api.themoviedb.org/3",
	}
}

func (cl *Client) get(path string, params url.Values) (res *http.Response, err error) {
	endpoint, err := url.JoinPath(cl.baseURL, path)
	if err != nil {
		return nil, err
	}

	// Add query parameters if present
	if len(params) > 0 {
		endpoint += "?" + params.Encode()
	}

	req, err := http.NewRequest("GET", endpoint, nil)
	if err != nil {
		return nil, err
	}

	// Add headers
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", cl.key))
	req.Header.Set("Content-Type", "application/json")

	// Make request
	client := &http.Client{}
	res, err = client.Do(req)
	if err != nil {
		return nil, err
	}

	return
}

const (
	DefaultGrouping = 0
	AirDateGrouping = iota + 1
	AbsoluteGrouping
	DVDGrouping
	DigitalGrouping
	StoryArcGrouping
	ProductionGrouping
	TVGrouping
)
