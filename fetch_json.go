package spot

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
)

func fetchJson(endpoint string, obj interface{}) error {

	res, err := http.Get(baseUrl + endpoint)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	if res.Status == "404 Not Found" {
		err = errors.New("404 Not Found")
		return err
	}

	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(data, obj)

	return err

}
