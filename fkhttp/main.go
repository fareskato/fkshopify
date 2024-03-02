package fkhttp

import (
	"encoding/json"
	"io"
	"net/http"
)

func HttpGet[T any](t T, url string) (*T, *http.Response, error) {
	var response *T
	res, err := http.Get(url)
	if err != nil {
		return nil, nil, err
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, nil, err
	}

	if err := json.Unmarshal(body, &response); err != nil {
		return nil, nil, err
	}
	return response, res, nil
}
