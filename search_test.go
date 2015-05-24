package spot

import (
	"testing"
)

func TestBuildHardUrl(t *testing.T) {

	desc := "Passing in incorrect query param name returns error"
	expected := "HardSearch can only search for artist, track, or album. Found: song"
	params := map[string]string{
		"artist": "Billy Joel",
		"song":   "Big Shot",
	}
	url, err := BuildHardUrl(params, []string{"track"})
	if err.Error() != expected {
		fail(t, desc, expected, err.Error())
	}

	desc = "Normal search returns correct url"
	expected = "https://api.spotify.com/v1/search?q=artist:mariah+carey+track:always+be+my+baby+&type=track"
	params = map[string]string{
		"artist": "Mariah Carey",
		"track":  "always be my baby",
	}
	url, _ = BuildHardUrl(params, []string{"track"})
	if url != expected {
		fail(t, desc, expected, url)
	}

	desc = "Passing in incorrect type returns error"
	expected = "HardSearch can only search by type artist, track, album, or playlist. Found: film"
	url, err = BuildHardUrl(params, []string{"film", "playlist"})
	if err.Error() != expected {
		fail(t, desc, expected, err.Error())
	}

}

func fail(t *testing.T, desc, expected, got string) {
	t.Errorf("Testing %q | Expected: %q, Got %q", desc, expected, got)
}