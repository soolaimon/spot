package spot

import (
	"reflect"
	"testing"
)

func TestGetArtist(t *testing.T) {

	desc := "Passing in a artist id returns the track from spotify"
	expected := reflect.TypeOf(Artist{})
	id := "4iHNK0tOyZPYnBU7nGAgpQ"

	artist, err := GetArtist(id)
	if reflect.TypeOf(artist) != expected {
		fail(t, desc, expected, err)
	}

	expectedString := id
	if artist.SpotId != id {
		fail(t, desc, expectedString, err)
	}

	desc = "Passing in a nonexistent artist id should throw err"
	id = "flimflam"
	expectedString = "404 Not Found"

	_, err = GetArtist(id)
	if err.Error() != expectedString {
		fail(t, desc, expectedString, err.Error())
	}

}
