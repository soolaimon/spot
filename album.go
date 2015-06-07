package spot

type Album struct {
	SpotId string
}

type AlbumResults struct {
	Href  string
	Items []Album
	Total int
}
