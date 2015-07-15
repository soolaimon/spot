package spot

type Artist struct {
	Href       string
	SpotId     string `json:"id"`
	Name       string
	Type       string
	Uri        string
	Genres     []string
	Popularity int
	Images     []Image
}

type ArtistResults struct {
	Href  string
	Items []Artist
	total int
}

func GetArtist(id string) (Artist, error) {
	var artist Artist
	err := fetchJson("artists/"+id, &artist)
	if err != nil {
		return artist, err
	}

	return artist, err
}
