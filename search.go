package spot

import (
	"errors"
	"net/url"
	"strings"
)

// Accept a map of search params (e.g track name, artist name, playlist name, etc) and
// search type slice [artist, playlist, track]. Types can be any number of types acceptable by api, must  have >= one.
func HardSearch(params map[string]string, itemTypes []string) (searchUrl string, err error) {

	// ensure keys of params are all Spotify field filters
	var passes bool
	for k, _ := range params {
		passes = checkParam(k)
		if passes == false {
			err := errors.New("HardSearch can only search for artist, track, or album. Found:" + k)
			return "", err
		}
	}

	queryString := ""

	// format params into spotify search query (q=track:Track+Name+artist:Artist+Name)
	for k, v := range params {
		queryString += url.QueryEscape(k) + ":" + url.QueryEscape(v) + "+"
	}

	return "https://api.spotify.com/v1/search?q=" + queryString + "&type=" + strings.Join(itemTypes, ","), err
}

// TODO: Soft search

func checkParam(param string) bool {

	var acceptableParams = []string{"artist", "track", "album"}

	for _, p := range acceptableParams {
		if param == p {
			return true
		}
	}
	return false
}
