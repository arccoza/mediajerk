package tmdb

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
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

func (cl *Client) get(path string) (res *http.Response, err error) {
	endpoint, err := url.JoinPath(cl.baseURL, path)
	if err != nil {
		return nil, err
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

func encodeCommonParams(params CommonSearchParams) url.Values {
	values := url.Values{}
	if params.Query != "" {
		values.Set("query", params.Query)
	}
	if params.Language != "" {
		values.Set("language", params.Language)
	}
	if params.Page > 0 {
		values.Set("page", strconv.Itoa(int(params.Page)))
	}
	if params.IncludeAdult {
		values.Set("include_adult", "true")
	}
	return values
}

func encodeMovieParams(params MovieSearchParams) url.Values {
	values := encodeCommonParams(params.CommonSearchParams)
	if params.PrimaryReleaseYear != "" {
		values.Set("primary_release_year", params.PrimaryReleaseYear)
	}
	if params.Region != "" {
		values.Set("region", params.Region)
	}
	if params.Year != "" {
		values.Set("year", params.Year)
	}
	return values
}

func encodeTVParams(params TVSearchParams) url.Values {
	values := encodeCommonParams(params.CommonSearchParams)
	if params.FirstAirDateYear > 0 {
		values.Set("first_air_date_year", strconv.Itoa(int(params.FirstAirDateYear)))
	}
	if params.Year > 0 {
		values.Set("year", strconv.Itoa(int(params.Year)))
	}
	return values
}

func encodeDetailsParams(params DetailsParams) url.Values {
	values := url.Values{}
	if params.Language != "" {
		values.Set("language", params.Language)
	}
	if params.AppendToResponse != "" {
		values.Set("append_to_response", params.AppendToResponse)
	}
	return values
}

func (cl *Client) SearchMovie(params MovieSearchParams) (*SearchResponse[Movie], error) {
	queryParams := encodeMovieParams(params)
	path := "search/movie?" + queryParams.Encode()

	resp, err := cl.get(path)
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
	queryParams := encodeTVParams(params)
	path := "search/tv?" + queryParams.Encode()

	resp, err := cl.get(path)
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

// https://developer.themoviedb.org/reference/movie-details
// https://api.themoviedb.org/3/movie/{movie_id}
func (cl *Client) Movies(movieId string, params DetailsParams) (*MovieDetails, error) {
	queryParams := encodeDetailsParams(params)
	path := "movie/" + movieId
	if queryParams.Encode() != "" {
		path += "?" + queryParams.Encode()
	}

	resp, err := cl.get(path)
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

// https://developer.themoviedb.org/reference/tv-series-details
// https://api.themoviedb.org/3/tv/{series_id}
func (cl *Client) TVSeries(seriesId string, params DetailsParams) (*TVSeriesDetails, error) {
	queryParams := encodeDetailsParams(params)
	path := "tv/" + seriesId
	if queryParams.Encode() != "" {
		path += "?" + queryParams.Encode()
	}

	resp, err := cl.get(path)
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

// https://developer.themoviedb.org/reference/tv-season-details
// https://api.themoviedb.org/3/tv/{series_id}/season/{season_number}
func (cl *Client) TVSeason(seriesId string, seasonNum int, params DetailsParams) (*TVSeasonDetails, error) {
	queryParams := encodeDetailsParams(params)
	path := "tv/" + seriesId + "/season/" + strconv.Itoa(seasonNum)
	if queryParams.Encode() != "" {
		path += "?" + queryParams.Encode()
	}

	resp, err := cl.get(path)
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

// https://developer.themoviedb.org/reference/tv-series-episode-groups
// https://api.themoviedb.org/3/tv/{series_id}/episode_groups
func (cl *Client) EpisodeGroups(seriesId string) (*EpisodeGroupList, error) {
	path := "tv/" + seriesId + "/episode_groups"

	resp, err := cl.get(path)
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

// https://developer.themoviedb.org/reference/tv-episode-group-details
// https://api.themoviedb.org/3/tv/episode_group/{tv_episode_group_id}
func (cl *Client) EpisodesGroupedBy(groupId string) (*EpisodeGroupDetails, error) {
	path := "tv/episode_group/" + groupId

	resp, err := cl.get(path)
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

	var groupDetails EpisodeGroupDetails
	if err := json.Unmarshal(body, &groupDetails); err != nil {
		return nil, err
	}

	return &groupDetails, nil
}

// Convenience methods for simple calls without parameters
func (cl *Client) MovieByID(movieId string) (*MovieDetails, error) {
	return cl.Movies(movieId, DetailsParams{})
}

func (cl *Client) TVSeriesByID(seriesId string) (*TVSeriesDetails, error) {
	return cl.TVSeries(seriesId, DetailsParams{})
}

func (cl *Client) TVSeasonByID(seriesId string, seasonNum int) (*TVSeasonDetails, error) {
	return cl.TVSeason(seriesId, seasonNum, DetailsParams{})
}
