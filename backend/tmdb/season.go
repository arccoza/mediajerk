package tmdb

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/google/go-querystring/query"
)

// https://developer.themoviedb.org/reference/tv-season-details
// https://api.themoviedb.org/3/tv/{series_id}/season/{season_number}
func (cl *Client) TVSeason(seriesId string, seasonNum int, params DetailsParams) (*TVSeasonDetails, error) {
	queryParams, err := query.Values(params)
	if err != nil {
		return nil, err
	}
	path := "tv/" + seriesId + "/season/" + strconv.Itoa(seasonNum)

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

	var seasonDetails TVSeasonDetails
	if err := json.Unmarshal(body, &seasonDetails); err != nil {
		return nil, err
	}

	return &seasonDetails, nil
}

// Convenience methods for simple calls without parameters
func (cl *Client) TVSeasonByID(seriesId string, seasonNum int) (*TVSeasonDetails, error) {
	return cl.TVSeason(seriesId, seasonNum, DetailsParams{})
}
