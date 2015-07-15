package spot

import (
	"reflect"
	"testing"
)

func TestGetAlbum(t *testing.T) {

	desc := "Passing in a album id returns the track from spotify"
	expected := reflect.TypeOf(Album{})
	id := "1ibYM4abQtSVQFQWvDSo4J"

	album, err := GetAlbum(id)
	if reflect.TypeOf(album) != expected {
		fail(t, desc, expected, err)
	}

	expectedString := id
	if album.SpotId != id {
		fail(t, desc, expectedString, err)
	}

	desc = "Passing in a nonexistent album id should throw err"
	id = "flimflam"
	expectedString = "404 Not Found"

	_, err = GetAlbum(id)
	if err.Error() != expectedString {
		fail(t, desc, expectedString, err.Error())
	}

}
