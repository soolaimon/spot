package spot

import "testing"

var baseUrl string = "https://api.spotify.com/v1/"

func fail(t *testing.T, desc string, expected, got interface{}) {
	t.Errorf("Testing %q | Expected: %q, Got %q", desc, expected, got)
}
