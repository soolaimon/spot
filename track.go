package spot

type Track struct {
	SpotId           string `json:"id"`
	Name             string
	Explicit         bool
	Artists          []Artist `json:"artists"`
	AvailableMarkets []string `json:"available_markets"`
	Album            Album
	DurationMS       int `json:"duration_ms"` // Duratino in milliseconds
	Popularity       int
	Type             string
	PreviewUrl       string `json:"preview_url"`
	Uri              string
	ExternalIds      map[string]string `json:"external_ids"`
	Href             string
}

type TrackResults struct {
	Href  string
	Items []Track
	Total int
}
