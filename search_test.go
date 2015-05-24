package spot

import "testing"

func TestBuildHardUrl(t *testing.T) {

	var params = map[string]string{
		"artist": "Mariah Carey",
		"track":  "always be my baby",
	}

	expected := "https://api.spotify.com/v1/search?q=artist:mariah+carey+track:always+be+my+baby+&type=track"
	url, _ := BuildHardUrl(params, []string{"track"})
	if url != expected {
		fail(t, expected, url)
	}

}

func fail(t *testing.T, expected, got string) {
	t.Errorf("Expected: %q, Got %q", expected, got)
}
