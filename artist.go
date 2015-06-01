package spot

type Artist struct {
	Href   string
	SpotId string `json:"id"`
	Name   string
	Type   string
	Uri    string
}
