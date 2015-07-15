package spot

import (
	"errors"
	"net/url"
	"sort"
	"strings"
)

type SearchResults struct {
	Tracks  TrackResults
	Albums  AlbumResults
	Artists ArtistResults
}

func HardSearch(params map[string]string, itemTypes []string) (results SearchResults, err error) {
	endpoint, err := BuildHardQuery(params, itemTypes)

	var searchResults SearchResults
	err = fetchJson(endpoint, &searchResults)
	if err != nil {
		panic(err)
	}

	return searchResults, err

}

var acceptableParams = []string{"artist", "track", "album"}
var acceptableTypes = []string{"artist", "track", "album", "playlist"}

// Accept a map of search params (e.g track name, artist name, playlist name, etc) and
// search type slice [artist, playlist, track]. Types can be any number of types acceptable by api, must  have >= one.
func BuildHardQuery(params map[string]string, itemTypes []string) (searchUrl string, err error) {

	var orderer []string
	// ensure keys of params are all Spotify field filters
	var passes bool
	for k, _ := range params {
		passes = checkParam(k, acceptableParams)
		if passes == false {
			err := errors.New("HardSearch can only search for artist, track, or album. Found: " + k)
			return "", err
		}
		// while we're at it, we create our array of param types for ordering below
		orderer = append(orderer, k)
	}

	// ensure types are compatible
	passes = true
	for _, v := range itemTypes {
		passes = checkParam(v, acceptableTypes)
		if passes == false {
			err := errors.New("HardSearch can only search by type artist, track, album, or playlist. Found: " + v)
			return "", err
		}
	}

	queryString := ""

	//	order params for predictability
	sort.Strings(orderer)

	// format params into spotify search query (q=track:Track+Name+artist:Artist+Name)
	for _, t := range orderer {
		queryString += url.QueryEscape(strings.ToLower(t)) + ":" + url.QueryEscape(strings.ToLower(params[t])) + "+"
	}

	return "search?q=" + queryString + "&type=" + strings.Join(itemTypes, ","), err
}

// TODO: Find song - Return first HardSearch result.

// TODO: Soft search

func checkParam(param string, check []string) bool {

	for _, p := range check {
		if param == p {
			return true
		}
	}
	return false
}
