package spot

import (
	"reflect"
	"testing"
)

func TestGetTrack(t *testing.T) {

	desc := "Passing in a track id returns the track from spotify"
	expected := reflect.TypeOf(Track{})
	id := "2aBxt229cbLDOvtL7Xbb9x"

	track, err := GetTrack(id)
	if reflect.TypeOf(track) != expected {
		fail(t, desc, expected, err)
	}

	expectedString := id
	if track.SpotId != id {
		fail(t, desc, expectedString, err)
	}

	desc = "Passing in a nonexistent track id should throw err"
	id = "flimflam"
	expectedString = "404 Not Found"

	_, err = GetTrack(id)
	if err.Error() != expectedString {
		fail(t, desc, expectedString, err.Error())
	}

}
