package spot

type Album struct {
	AvailableMarkets []string
	Artists          []Artist
	spotId           string `json:"id"`
	Uri              string
	Genres           []string
	Name             string
	ReleaseDate      string `json:"release_date"`
	ReleasePrecision string `json:"release_date_precision"`
	Tracks           TrackResults
}

type AlbumResults struct {
	Href  string
	Items []Album
	Total int
}
