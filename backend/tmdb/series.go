package tmdb

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/google/go-querystring/query"
)

// https://developer.themoviedb.org/reference/tv-series-details
// https://api.themoviedb.org/3/tv/{series_id}
func (cl *Client) TVSeries(seriesId string, params DetailsParams) (*TVSeriesDetails, error) {
	queryParams, err := query.Values(params)
	if err != nil {
		return nil, err
	}
	path := "tv/" + seriesId

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

	var tvDetails TVSeriesDetails
	if err := json.Unmarshal(body, &tvDetails); err != nil {
		return nil, err
	}

	return &tvDetails, nil
}

// Convenience methods for simple calls without parameters
func (cl *Client) TVSeriesByID(seriesId string) (*TVSeriesDetails, error) {
	return cl.TVSeries(seriesId, DetailsParams{})
}

// https://developer.themoviedb.org/reference/tv-series-episode-groups
// https://api.themoviedb.org/3/tv/{series_id}/episode_groups
func (cl *Client) EpisodeGroups(seriesId string) (*EpisodeGroupList, error) {
	path := "tv/" + seriesId + "/episode_groups"

	resp, err := cl.get(path, url.Values{})
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

	var episodeGroups EpisodeGroupList
	if err := json.Unmarshal(body, &episodeGroups); err != nil {
		return nil, err
	}

	return &episodeGroups, nil
}
