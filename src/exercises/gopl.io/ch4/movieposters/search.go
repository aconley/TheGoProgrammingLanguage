package movieposters

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

// getOMDBInfo queries the OMDB to get movie information
func getOMDBInfoByTitle(title string) (*PosterSearchResult, error) {
	q := url.QueryEscape(title)
	resp, err := http.Get(QueryURL + "/?s=" + q)
    if resp != nil {
        defer resp.Body.Close()
    }
	if err != nil {
		return nil, err
	}
	
	// We must close resp.Body on all execution paths.
	// (Chapter 5 presents 'defer', which makes this simpler.)
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("search query failed: %s", resp.Status)
	}

	var result PosterSearchResult
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}
	return &result, nil
}

// GetPosterURLs gets a slice of poster urls for movies
func GetPosterURLs(title string) ([]string, error) {
    info, err := getOMDBInfoByTitle(title);
    if err != nil {
        return nil, err
    }
    n := len(info.Items)
    if n == 0 {
        // Not an error, per se
        return nil, nil
    }
    
    urls := make([]string, 0, n)
    for _, item := range(info.Items) {
        if item.Type == "movie" && item.PosterURL != "N/A" {
            urls = append(urls, item.PosterURL)
        }
    }
    return urls, nil
}