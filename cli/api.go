package cli

import (
	"encoding/json"
	"net/http"
)

type ExampleRes struct {
	Name   string `json:"name"`
	Commit struct {
		Sha string `json:"sha"`
		URL string `json:"url"`
	} `json:"commit"`
	Protected bool `json:"protected"`
}

type ApiResponse interface {
	ExampleRes
}

func fetchJSON[T ApiResponse](hc *http.Client, url string, data *T) error {
	resp, err := hc.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	err = json.NewDecoder(resp.Body).Decode(data)
	if err != nil {
		return err
	}
	return nil
}
