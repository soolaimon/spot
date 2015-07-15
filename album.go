package spot

type Album struct {
	AvailableMarkets []string
	Artists          []Artist
	SpotId           string `json:"id"`
	Uri              string
	Genres           []string
	Name             string
	ReleaseDate      string `json:"release_date"`
	ReleasePrecision string `json:"release_date_precision"`
	Tracks           TrackResults
	Images           []Image
}

type AlbumResults struct {
	Href  string
	Items []Album
	Total int
}

func GetAlbum(id string) (Album, error) {
	var album Album
	err := fetchJson("albums/"+id, &album)
	if err != nil {
		return album, err
	}

	return album, err
}
