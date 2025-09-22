package tmdb

import (
	"encoding/json"
	"regexp"
	"sort"
	"strconv"
)

type CommonSearchParams struct {
	Query        string `url:"query,omitempty"`
	IncludeAdult bool   `url:"include_adult,omitempty"`
	Language     string `url:"language,omitempty"`
	Page         int32  `url:"page,omitempty"`
}

type MovieSearchParams struct {
	CommonSearchParams
	PrimaryReleaseYear string `url:"primary_release_year,omitempty"`
	Region             string `url:"region,omitempty"`
	Year               string `url:"year,omitempty"`
}

type TVSearchParams struct {
	CommonSearchParams
	FirstAirDateYear int32 `url:"first_air_date_year,omitempty"`
	Year             int32 `url:"year,omitempty"`
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

type Person struct {
	ID                 int          `json:"id"`
	Name               string       `json:"name"`
	ProfilePath        *string      `json:"profile_path"`
	Adult              bool         `json:"adult"`
	Popularity         float64      `json:"popularity"`
	KnownForDepartment string       `json:"known_for_department"`
	Gender             int          `json:"gender"`
	KnownFor           []MultiMedia `json:"known_for"`
}

type MultiMedia struct {
	MediaType string `json:"media_type"`

	// Embedded types - only one will be populated based on MediaType
	// Use json:"-" to exclude from automatic JSON processing since we handle it manually
	Movie  `json:"-"`
	TVShow `json:"-"`
	Person `json:"-"`
}

func (m *MultiMedia) UnmarshalJSON(data []byte) error {
	// First extract the media_type
	var mediaTypeStruct struct {
		MediaType string `json:"media_type"`
	}
	if err := json.Unmarshal(data, &mediaTypeStruct); err != nil {
		return err
	}
	m.MediaType = mediaTypeStruct.MediaType

	// Based on media_type, unmarshal into the appropriate embedded struct
	switch m.MediaType {
	case "movie":
		return json.Unmarshal(data, &m.Movie)
	case "tv":
		return json.Unmarshal(data, &m.TVShow)
	case "person":
		return json.Unmarshal(data, &m.Person)
	default:
		// For unknown types, try to unmarshal as much as possible
		// This provides forward compatibility
		json.Unmarshal(data, &m.Movie)
		json.Unmarshal(data, &m.TVShow)
		json.Unmarshal(data, &m.Person)
	}

	return nil
}

type SearchResponse[T any] struct {
	Page         int `json:"page"`
	Results      []T `json:"results"`
	TotalPages   int `json:"total_pages"`
	TotalResults int `json:"total_results"`
}

type DetailsParams struct {
	Language         string `url:"language,omitempty"`
	AppendToResponse string `url:"append_to_response,omitempty"`
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

	// Append response fields (populated by custom UnmarshalJSON)
	Videos  *VideosResponse  `json:"videos,omitempty"`
	Images  *ImagesResponse  `json:"images,omitempty"`
	Credits *CreditsResponse `json:"credits,omitempty"`
}

func (m *MovieDetails) UnmarshalJSON(data []byte) error {
	// Create alias to avoid recursion
	type MovieDetailsAlias MovieDetails

	// Unmarshal normally - the append fields are handled by normal JSON tags
	var alias MovieDetailsAlias
	if err := json.Unmarshal(data, &alias); err != nil {
		return err
	}

	// Copy all fields
	*m = MovieDetails(alias)

	return nil
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

	// Append response fields (populated by custom UnmarshalJSON)
	FullSeasons   []TVSeasonDetails `json:"-"` // From season/N keys
	Videos        *VideosResponse   `json:"videos,omitempty"`
	Images        *ImagesResponse   `json:"images,omitempty"`
	Credits       *CreditsResponse  `json:"credits,omitempty"`
	EpisodeGroups *EpisodeGroupList `json:"episode_groups,omitempty"`
}

func (t *TVSeriesDetails) UnmarshalJSON(data []byte) error {
	// Create alias to avoid recursion
	type TVSeriesDetailsAlias TVSeriesDetails

	// First unmarshal normally into alias
	var alias TVSeriesDetailsAlias
	if err := json.Unmarshal(data, &alias); err != nil {
		return err
	}

	// Copy all normal fields
	*t = TVSeriesDetails(alias)

	// Parse raw JSON to extract dynamic season keys
	var raw map[string]json.RawMessage
	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}

	// Extract season/N keys
	seasonRegex := regexp.MustCompile(`^season/(\d+)$`)
	var seasons []TVSeasonDetails
	var seasonNumbers []int

	for key, value := range raw {
		if matches := seasonRegex.FindStringSubmatch(key); matches != nil {
			seasonNum, err := strconv.Atoi(matches[1])
			if err != nil {
				continue
			}

			var season TVSeasonDetails
			if err := json.Unmarshal(value, &season); err != nil {
				continue
			}

			seasons = append(seasons, season)
			seasonNumbers = append(seasonNumbers, seasonNum)
		}
	}

	// Sort seasons by season number
	if len(seasons) > 0 {
		// Create pairs of (seasonNumber, season) for sorting
		type seasonPair struct {
			num    int
			season TVSeasonDetails
		}

		pairs := make([]seasonPair, len(seasons))
		for i := range seasons {
			pairs[i] = seasonPair{seasonNumbers[i], seasons[i]}
		}

		sort.Slice(pairs, func(i, j int) bool {
			return pairs[i].num < pairs[j].num
		})

		// Extract sorted seasons
		sortedSeasons := make([]TVSeasonDetails, len(pairs))
		for i, pair := range pairs {
			sortedSeasons[i] = pair.season
		}

		t.FullSeasons = sortedSeasons
	}

	return nil
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
	ID      int            `json:"id"`
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

// Append response types for common TMDB append_to_response endpoints
type Video struct {
	ID          string `json:"id"`
	ISO639_1    string `json:"iso_639_1"`
	ISO3166_1   string `json:"iso_3166_1"`
	Key         string `json:"key"`
	Name        string `json:"name"`
	Official    bool   `json:"official"`
	PublishedAt string `json:"published_at"`
	Site        string `json:"site"`
	Size        int    `json:"size"`
	Type        string `json:"type"`
}

type VideosResponse struct {
	Results []Video `json:"results"`
}

type Image struct {
	AspectRatio float64 `json:"aspect_ratio"`
	Height      int     `json:"height"`
	ISO639_1    *string `json:"iso_639_1"`
	FilePath    string  `json:"file_path"`
	VoteAverage float64 `json:"vote_average"`
	VoteCount   int     `json:"vote_count"`
	Width       int     `json:"width"`
}

type ImagesResponse struct {
	Backdrops []Image `json:"backdrops"`
	Logos     []Image `json:"logos"`
	Posters   []Image `json:"posters"`
}

type CastMember struct {
	Adult              bool    `json:"adult"`
	Gender             int     `json:"gender"`
	ID                 int     `json:"id"`
	KnownForDepartment string  `json:"known_for_department"`
	Name               string  `json:"name"`
	OriginalName       string  `json:"original_name"`
	Popularity         float64 `json:"popularity"`
	ProfilePath        *string `json:"profile_path"`
	CastID             int     `json:"cast_id"`
	Character          string  `json:"character"`
	CreditID           string  `json:"credit_id"`
	Order              int     `json:"order"`
}

type CrewMember struct {
	Adult              bool    `json:"adult"`
	Gender             int     `json:"gender"`
	ID                 int     `json:"id"`
	KnownForDepartment string  `json:"known_for_department"`
	Name               string  `json:"name"`
	OriginalName       string  `json:"original_name"`
	Popularity         float64 `json:"popularity"`
	ProfilePath        *string `json:"profile_path"`
	CreditID           string  `json:"credit_id"`
	Department         string  `json:"department"`
	Job                string  `json:"job"`
}

type CreditsResponse struct {
	Cast []CastMember `json:"cast"`
	Crew []CrewMember `json:"crew"`
}
