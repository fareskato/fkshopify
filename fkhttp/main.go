package fkhttp

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// HttpShopifyGet send get request to shopify and return the requested resource.
func HttpShopifyGet[T any](t T, url string) (*T, *http.Response, error) {
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

// HttpShopifyRequestWithHeaders general purpose function sends(get, post, put)
// to interact with shopify admin rest api.
func HttpShopifyRequestWithHeaders[T any](ctx context.Context, method string, url string, t T) (*http.Response, error) {

	payload, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest(method, url, bytes.NewBuffer(payload))
	if err != nil {
		return nil, err
	}
	// set headers
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Content-Length", fmt.Sprint(len(payload)))
	// context
	req = req.WithContext(ctx)
	// send the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
