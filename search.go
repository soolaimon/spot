package spot

import (
	"net/url"
	"strings"
)

func BuildUrl(params map[string]string, itemTypes []string) (searchUrl string, err error) {

	queryString := ""

	for k, v := range params {
		queryString += url.QueryEscape(k) + ":" + url.QueryEscape(v) + "+"
	}

	return "https://api.spotify.com/v1/search?q=" + queryString + "&type=" + strings.Join(itemTypes, ","), err
}
