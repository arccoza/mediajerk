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

type DetailsParams struct {
	Language         string
	AppendToResponse string
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

type Genre struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type ProductionCompany struct {
	ID            int     `json:"id"`
	LogoPath      *string `json:"logo_path"`
	Name          string  `json:"name"`
	OriginCountry string  `json:"origin_country"`
}

type ProductionCountry struct {
	ISO3166_1 string `json:"iso_3166_1"`
	Name      string `json:"name"`
}

type SpokenLanguage struct {
	EnglishName string `json:"english_name"`
	ISO639_1    string `json:"iso_639_1"`
	Name        string `json:"name"`
}

type Collection struct {
	ID           int     `json:"id"`
	Name         string  `json:"name"`
	PosterPath   *string `json:"poster_path"`
	BackdropPath *string `json:"backdrop_path"`
}

type MovieDetails struct {
	Adult               bool                `json:"adult"`
	BackdropPath        *string             `json:"backdrop_path"`
	BelongsToCollection *Collection         `json:"belongs_to_collection"`
	Budget              int                 `json:"budget"`
	Genres              []Genre             `json:"genres"`
	Homepage            *string             `json:"homepage"`
	ID                  int                 `json:"id"`
	IMDbID              *string             `json:"imdb_id"`
	OriginCountry       []string            `json:"origin_country"`
	OriginalLanguage    string              `json:"original_language"`
	OriginalTitle       string              `json:"original_title"`
	Overview            *string             `json:"overview"`
	Popularity          float64             `json:"popularity"`
	PosterPath          *string             `json:"poster_path"`
	ProductionCompanies []ProductionCompany `json:"production_companies"`
	ProductionCountries []ProductionCountry `json:"production_countries"`
	ReleaseDate         string              `json:"release_date"`
	Revenue             int                 `json:"revenue"`
	Runtime             *int                `json:"runtime"`
	SpokenLanguages     []SpokenLanguage    `json:"spoken_languages"`
	Status              string              `json:"status"`
	Tagline             *string             `json:"tagline"`
	Title               string              `json:"title"`
	Video               bool                `json:"video"`
	VoteAverage         float64             `json:"vote_average"`
	VoteCount           int                 `json:"vote_count"`
}

type CreatedBy struct {
	ID          int     `json:"id"`
	CreditID    string  `json:"credit_id"`
	Name        string  `json:"name"`
	Gender      int     `json:"gender"`
	ProfilePath *string `json:"profile_path"`
}

type Network struct {
	ID            int     `json:"id"`
	LogoPath      *string `json:"logo_path"`
	Name          string  `json:"name"`
	OriginCountry string  `json:"origin_country"`
}

type Season struct {
	AirDate      *string `json:"air_date"`
	EpisodeCount int     `json:"episode_count"`
	ID           int     `json:"id"`
	Name         string  `json:"name"`
	Overview     string  `json:"overview"`
	PosterPath   *string `json:"poster_path"`
	SeasonNumber int     `json:"season_number"`
	VoteAverage  float64 `json:"vote_average"`
}

type Episode struct {
	ID             int     `json:"id"`
	Name           string  `json:"name"`
	Overview       string  `json:"overview"`
	VoteAverage    float64 `json:"vote_average"`
	VoteCount      int     `json:"vote_count"`
	AirDate        *string `json:"air_date"`
	EpisodeNumber  int     `json:"episode_number"`
	EpisodeType    string  `json:"episode_type"`
	ProductionCode *string `json:"production_code"`
	Runtime        *int    `json:"runtime"`
	SeasonNumber   int     `json:"season_number"`
	ShowID         int     `json:"show_id"`
	StillPath      *string `json:"still_path"`
}

type TVSeriesDetails struct {
	Adult               bool                `json:"adult"`
	BackdropPath        *string             `json:"backdrop_path"`
	CreatedBy           []CreatedBy         `json:"created_by"`
	EpisodeRunTime      []int               `json:"episode_run_time"`
	FirstAirDate        *string             `json:"first_air_date"`
	Genres              []Genre             `json:"genres"`
	Homepage            string              `json:"homepage"`
	ID                  int                 `json:"id"`
	InProduction        bool                `json:"in_production"`
	Languages           []string            `json:"languages"`
	LastAirDate         *string             `json:"last_air_date"`
	LastEpisodeToAir    *Episode            `json:"last_episode_to_air"`
	Name                string              `json:"name"`
	NextEpisodeToAir    *Episode            `json:"next_episode_to_air"`
	Networks            []Network           `json:"networks"`
	NumberOfEpisodes    int                 `json:"number_of_episodes"`
	NumberOfSeasons     int                 `json:"number_of_seasons"`
	OriginCountry       []string            `json:"origin_country"`
	OriginalLanguage    string              `json:"original_language"`
	OriginalName        string              `json:"original_name"`
	Overview            string              `json:"overview"`
	Popularity          float64             `json:"popularity"`
	PosterPath          *string             `json:"poster_path"`
	ProductionCompanies []ProductionCompany `json:"production_companies"`
	ProductionCountries []ProductionCountry `json:"production_countries"`
	Seasons             []Season            `json:"seasons"`
	SpokenLanguages     []SpokenLanguage    `json:"spoken_languages"`
	Status              string              `json:"status"`
	Tagline             string              `json:"tagline"`
	Type                string              `json:"type"`
	VoteAverage         float64             `json:"vote_average"`
	VoteCount           int                 `json:"vote_count"`
}

type TVSeasonDetails struct {
	ID           int       `json:"id"`
	AirDate      *string   `json:"air_date"`
	Episodes     []Episode `json:"episodes"`
	Name         string    `json:"name"`
	Overview     string    `json:"overview"`
	PosterPath   *string   `json:"poster_path"`
	SeasonNumber int       `json:"season_number"`
	VoteAverage  float64   `json:"vote_average"`
}

type EpisodeGroup struct {
	Description  string `json:"description"`
	EpisodeCount int    `json:"episode_count"`
	GroupCount   int    `json:"group_count"`
	ID           string `json:"id"`
	Name         string `json:"name"`
	Network      *struct {
		ID            int     `json:"id"`
		LogoPath      *string `json:"logo_path"`
		Name          string  `json:"name"`
		OriginCountry string  `json:"origin_country"`
	} `json:"network"`
	Type int `json:"type"`
}

type EpisodeGroupList struct {
	Results []EpisodeGroup `json:"results"`
}

type EpisodeGroupDetails struct {
	Description string `json:"description"`
	ID          string `json:"id"`
	Name        string `json:"name"`
	Network     *struct {
		ID            int     `json:"id"`
		LogoPath      *string `json:"logo_path"`
		Name          string  `json:"name"`
		OriginCountry string  `json:"origin_country"`
	} `json:"network"`
	Type   int `json:"type"`
	Groups []struct {
		ID       string    `json:"id"`
		Name     string    `json:"name"`
		Order    int       `json:"order"`
		Episodes []Episode `json:"episodes"`
	} `json:"groups"`
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
