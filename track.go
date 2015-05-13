package spot

type Track struct {
	Id       string `json:"id"`
	Name     string
	Explicit bool
	//Artists          []Artist `json:"artists"`
	AvailableMarkets []string `json:"available_markets"`
}
