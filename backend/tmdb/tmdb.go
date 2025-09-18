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

type CommonSearchParams struct {
	Query        string
	IncludeAdult bool
	Language     string
	Page         int32
}

type MovieSearchParams struct {
	CommonSearchParams
	PrimaryReleaseYear string
	Region             string
	Year               string
}

type TVSearchParams struct {
	CommonSearchParams
	FirstAirDateYear int32
	Year             int32
}

type Movie struct {
	ID               int     `json:"id"`
	Title            string  `json:"title"`
	OriginalTitle    string  `json:"original_title"`
	OriginalLanguage string  `json:"original_language"`
	Overview         string  `json:"overview"`
	PosterPath       *string `json:"poster_path"`
	BackdropPath     *string `json:"backdrop_path"`
	ReleaseDate      string  `json:"release_date"`
	Adult            bool    `json:"adult"`
	Popularity       float64 `json:"popularity"`
	VoteAverage      float64 `json:"vote_average"`
	VoteCount        int     `json:"vote_count"`
	GenreIDs         []int   `json:"genre_ids"`
}

type TVShow struct {
	ID               int      `json:"id"`
	Name             string   `json:"name"`
	OriginalName     string   `json:"original_name"`
	OriginalLanguage string   `json:"original_language"`
	Overview         string   `json:"overview"`
	PosterPath       *string  `json:"poster_path"`
	BackdropPath     *string  `json:"backdrop_path"`
	FirstAirDate     string   `json:"first_air_date"`
	Adult            bool     `json:"adult"`
	Popularity       float64  `json:"popularity"`
	VoteAverage      float64  `json:"vote_average"`
	VoteCount        int      `json:"vote_count"`
	GenreIDs         []int    `json:"genre_ids"`
	OriginCountry    []string `json:"origin_country"`
}

type SearchResponse[T any] struct {
	Page         int `json:"page"`
	Results      []T `json:"results"`
	TotalPages   int `json:"total_pages"`
	TotalResults int `json:"total_results"`
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
