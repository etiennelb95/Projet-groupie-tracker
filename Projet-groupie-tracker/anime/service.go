package anime

import (
	"fmt"
	"net/url"
	"strconv"
	
	"jikan-api-wrapper/common"
)

// AnimeData represents the structure of anime data from the Jikan API
type AnimeData struct {
	MalID         int         `json:"mal_id"`
	URL           string      `json:"url"`
	Title         string      `json:"title"`
	TitleEnglish  string      `json:"title_english"`
	TitleJapanese string      `json:"title_japanese"`
	TitleSynonyms []string    `json:"title_synonyms"`
	Type          string      `json:"type"`
	Source        string      `json:"source"`
	Episodes      int         `json:"episodes"`
	Status        string      `json:"status"`
	Airing        bool        `json:"airing"`
	Synopsis      string      `json:"synopsis"`
	Background    string      `json:"background"`
	Season        string      `json:"season"`
	Year          int         `json:"year"`
	Score         float64     `json:"score"`
	ScoredBy      int         `json:"scored_by"`
	Rank          int         `json:"rank"`
	Popularity    int         `json:"popularity"`
	Members       int         `json:"members"`
	Favorites     int         `json:"favorites"`
	Genres        []GenreInfo `json:"genres"`
	Studios       []StudioInfo `json:"studios"`
	Images        struct {
		JPG  ImageInfo `json:"jpg"`
		WebP ImageInfo `json:"webp"`
	} `json:"images"`
}

type GenreInfo struct {
	MalID int    `json:"mal_id"`
	Type  string `json:"type"`
	Name  string `json:"name"`
}

type StudioInfo struct {
	MalID int    `json:"mal_id"`
	Type  string `json:"type"`
	Name  string `json:"name"`
}

type ImageInfo struct {
	ImageURL       string `json:"image_url"`
	SmallImageURL  string `json:"small_image_url"`
	LargeImageURL  string `json:"large_image_url"`
}

// AnimeResponse represents the structure of Jikan API anime search response
type AnimeResponse struct {
	Data []AnimeData `json:"data"`
	Pagination struct {
		LastVisiblePage int  `json:"last_visible_page"`
		HasNextPage     bool `json:"has_next_page"`
		CurrentPage     int  `json:"current_page"`
		Items           struct {
			Count    int `json:"count"`
			PerPage  int `json:"per_page"`
			Total    int `json:"total"`
		} `json:"items"`
	} `json:"pagination"`
}

// SingleAnimeResponse represents the structure of Jikan API single anime response
type SingleAnimeResponse struct {
	Data AnimeData `json:"data"`
}

// SearchAnime searches for anime by title using the Jikan API
func SearchAnime(query string, limit, page int) ([]AnimeData, int, error) {
	encodedQuery := url.QueryEscape(query)
	endpoint := fmt.Sprintf("/anime?q=%s&limit=%d&page=%d", encodedQuery, limit, page)
	
	var response AnimeResponse
	err := common.DefaultClient.GetAndUnmarshal(endpoint, &response)
	if err != nil {
		return nil, 0, err
	}
	
	return response.Data, response.Pagination.Items.Total, nil
}

// GetAnimeDetails retrieves detailed information about a specific anime
func GetAnimeDetails(id int) (*AnimeData, error) {
	endpoint := fmt.Sprintf("/anime/%d", id)
	
	var response SingleAnimeResponse
	err := common.DefaultClient.GetAndUnmarshal(endpoint, &response)
	if err != nil {
		return nil, err
	}
	
	return &response.Data, nil
}

// GetTopAnime retrieves the top-rated anime
func GetTopAnime(filter string, limit, page int) ([]AnimeData, int, error) {
	endpoint := fmt.Sprintf("/top/anime?limit=%d&page=%d", limit, page)
	
	if filter != "" {
		endpoint += "&filter=" + url.QueryEscape(filter)
	}
	
	var response AnimeResponse
	err := common.DefaultClient.GetAndUnmarshal(endpoint, &response)
	if err != nil {
		return nil, 0, err
	}
	
	return response.Data, response.Pagination.Items.Total, nil
}