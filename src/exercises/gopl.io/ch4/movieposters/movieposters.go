package movieposters

// QueryURL is the url for the Open Movie Database
const QueryURL = "https://omdbapi.com"

// PosterSearchResult is the overall results wrapper
type PosterSearchResult struct {
  Items        []*MovieInfo `json:"Search"`
	TotalResults string `json:"totalResults"`
	Response     string
}

// MovieInfo represents OMDB information about a particular movie
type MovieInfo struct {
	Title     string
  Year      string
  IMDBId    string `json:"imdbID"`
  Type      string
  PosterURL string `json:"Poster"`
}