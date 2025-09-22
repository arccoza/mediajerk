package tmdb

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/google/go-querystring/query"
)

// https://developer.themoviedb.org/reference/movie-details
// https://api.themoviedb.org/3/movie/{movie_id}
func (cl *Client) Movies(movieId string, params DetailsParams) (*MovieDetails, error) {
	queryParams, err := query.Values(params)
	if err != nil {
		return nil, err
	}
	path := "movie/" + movieId

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

	var movieDetails MovieDetails
	if err := json.Unmarshal(body, &movieDetails); err != nil {
		return nil, err
	}

	return &movieDetails, nil
}

// Convenience methods for simple calls without parameters
func (cl *Client) MovieByID(movieId string) (*MovieDetails, error) {
	return cl.Movies(movieId, DetailsParams{})
}
