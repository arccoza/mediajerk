package tmdb

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/google/go-querystring/query"
)

func (cl *Client) SearchMovie(params MovieSearchParams) (*SearchResponse[Movie], error) {
	queryParams, err := query.Values(params)
	if err != nil {
		return nil, err
	}
	path := "search/movie"

	resp, err := cl.get(path, queryParams)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API request failed with status %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var searchResp SearchResponse[Movie]
	if err := json.Unmarshal(body, &searchResp); err != nil {
		return nil, err
	}

	return &searchResp, nil
}

func (cl *Client) SearchTV(params TVSearchParams) (*SearchResponse[TVShow], error) {
	queryParams, err := query.Values(params)
	if err != nil {
		return nil, err
	}
	path := "search/tv"

	resp, err := cl.get(path, queryParams)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API request failed with status %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var searchResp SearchResponse[TVShow]
	if err := json.Unmarshal(body, &searchResp); err != nil {
		return nil, err
	}

	return &searchResp, nil
}

// https://developer.themoviedb.org/reference/search-multi
// https://api.themoviedb.org/3/search/multi
func (cl *Client) SearchMulti(params CommonSearchParams) (*SearchResponse[MultiMedia], error) {
	queryParams, err := query.Values(params)
	if err != nil {
		return nil, err
	}
	path := "search/multi"

	resp, err := cl.get(path, queryParams)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API request failed with status %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var searchResp SearchResponse[MultiMedia]
	if err := json.Unmarshal(body, &searchResp); err != nil {
		return nil, err
	}

	return &searchResp, nil
}

// Convenience methods for simple query-only searches
func (cl *Client) SearchMovieByQuery(query string) (*SearchResponse[Movie], error) {
	params := MovieSearchParams{
		CommonSearchParams: CommonSearchParams{
			Query: query,
		},
	}
	return cl.SearchMovie(params)
}

func (cl *Client) SearchTVByQuery(query string) (*SearchResponse[TVShow], error) {
	params := TVSearchParams{
		CommonSearchParams: CommonSearchParams{
			Query: query,
		},
	}
	return cl.SearchTV(params)
}

func (cl *Client) SearchMultiByQuery(query string) (*SearchResponse[MultiMedia], error) {
	params := CommonSearchParams{
		Query: query,
	}
	return cl.SearchMulti(params)
}
