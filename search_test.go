package spot

import (
	"reflect"
	"testing"
)

func TestBuildHardQuery(t *testing.T) {

	desc := "Passing in incorrect query param name returns error"
	expected := "HardSearch can only search for artist, track, or album. Found: song"
	params := map[string]string{
		"artist": "Billy Joel",
		"song":   "Big Shot",
	}
	url, err := BuildHardQuery(params, []string{"track"})
	if err.Error() != expected {
		fail(t, desc, expected, err.Error())
	}

	desc = "Normal search returns correct url"
	expected = "search?q=artist:mariah+carey+track:always+be+my+baby+&type=track"
	params = map[string]string{
		"artist": "Mariah Carey",
		"track":  "always be my baby",
	}
	url, _ = BuildHardQuery(params, []string{"track"})
	if url != expected {
		fail(t, desc, expected, url)
	}

	desc = "Normal search with mutltiple types returns correct url"
	expected = "search?q=artist:mariah+carey+track:always+be+my+baby+&type=track,album"
	params = map[string]string{
		"artist": "Mariah Carey",
		"track":  "always be my baby",
	}
	url, _ = BuildHardQuery(params, []string{"track", "album"})
	if url != expected {
		fail(t, desc, expected, url)
	}

	desc = "Passing in incorrect type returns error"
	expected = "HardSearch can only search by type artist, track, album, or playlist. Found: film"
	url, err = BuildHardQuery(params, []string{"film", "playlist"})
	if err.Error() != expected {
		fail(t, desc, expected, err.Error())
	}
}

func TestHardSearch(t *testing.T) {

	params := map[string]string{
		"artist": "Billy Joel",
		"track":  "Big Shot",
	}
	types := []string{"track"}

	desc := "Search for track returns SearchResult struct."
	expected := reflect.TypeOf(SearchResults{})
	results, err := HardSearch(params, types)
	if reflect.TypeOf(results) != expected {
		fail(t, desc, expected, err)
	}

	desc = "Search for track- SearchResults.TrackResults is TrackResults{}"
	expected = reflect.TypeOf(TrackResults{})
	if reflect.TypeOf(results.Tracks) != expected {
		fail(t, desc, expected, err)
	}

	//	err := godotenv.Load()
	//	if err != nil {
	//		log.Fatal("Error loading .env file")
	//	}
	//
	//	spotifyKey := os.Getenv("SPOTIFY_CLIENT_ID")
	//	spotifySecret := os.Getenv("SPOTIFY_CLIENT_SECRET")
}
