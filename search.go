package spot

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
	"sort"
	"strings"
)

type SearchResults struct {
	Tracks TrackResults `json:tracks`
}

func HardSearch(params map[string]string, itemTypes []string) (results SearchResults, err error) {
	url, err := BuildHardUrl(params, itemTypes)
	if err != nil {
		panic(err)
	}

	res, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}

	var searchResults SearchResults
	err = json.Unmarshal(data, &searchResults)
	if err != nil {
		panic(err)
	}

	return searchResults, err

}

var acceptableParams = []string{"artist", "track", "album"}
var acceptableTypes = []string{"artist", "track", "album", "playlist"}

// Accept a map of search params (e.g track name, artist name, playlist name, etc) and
// search type slice [artist, playlist, track]. Types can be any number of types acceptable by api, must  have >= one.
func BuildHardUrl(params map[string]string, itemTypes []string) (searchUrl string, err error) {

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

	return "https://api.spotify.com/v1/search?q=" + queryString + "&type=" + strings.Join(itemTypes, ","), err
}

// TODO: Soft search

func checkParam(param string, check []string) bool {

	for _, p := range check {
		if param == p {
			return true
		}
	}
	return false
}
