package spot

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
)

type Track struct {
	SpotId           string `json:"id"`
	Name             string
	Explicit         bool
	Artists          []Artist `json:"artists"`
	AvailableMarkets []string `json:"available_markets"`
	Album            Album
	DurationMS       int `json:"duration_ms"` // Duratino in milliseconds
	Popularity       int
	Type             string
	PreviewUrl       string `json:"preview_url"`
	Uri              string
	ExternalIds      map[string]string `json:"external_ids"`
	Href             string
}

type TrackResults struct {
	Href  string
	Items []Track
	Total int
}

func GetTrack(id string) (Track, error) {
	res, err := http.Get(baseUrl + "tracks/" + id)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	if res.Status == "404 Not Found" {
		err = errors.New("The call did not return a track. The id you passed GetTrack is probably incorrect. Got: " + id)
		return Track{}, err
	}

	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}

	var track Track
	err = json.Unmarshal(data, &track)

	return track, err
}
