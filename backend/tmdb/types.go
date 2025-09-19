package tmdb

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
