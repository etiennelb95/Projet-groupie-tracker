package person

import (
	"github.com/nokusukun/jikan2go/mal_types"
	"github.com/nokusukun/jikan2go/utils"
)

func GetTop(page int) (Top, error) {
	request, err := utils.CachedReqGet(utils.Config.AppendAPIf("/top/people/%v", page))
	if err != nil {
		return Top{}, err
	}

	var a Top
	err = request.ToJSON(&a)

	return a, err
}

type Top struct {
	RequestHash        string       `json:"request_hash"`
	RequestCached      bool         `json:"request_cached"`
	RequestCacheExpiry int64        `json:"request_cache_expiry"`
	Top                []TopElement `json:"top"`
}

type TopElement struct {
	MalID     int64   `json:"mal_id"`
	Rank      int64   `json:"rank"`
	Title     string  `json:"title"`
	URL       string  `json:"url"`
	NameKanji *string `json:"name_kanji"`
	Favorites int64   `json:"favorites"`
	ImageURL  string  `json:"image_url"`
	Birthday  string  `json:"birthday"`
}

func (r TopElement) GetID() interface{} {
	return r.MalID
}

func (r TopElement) GetType() string {
	return mal_types.Person
}
