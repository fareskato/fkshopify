package fkhttp

import (
	"bytes"
	"context"
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

func HttpShopifyRequestWithHeaders(method string, ctx context.Context, url, token string, payload []byte) (*http.Response, error) {
	req, err := http.NewRequest(method, url, bytes.NewBuffer(payload))
	if err != nil {
		return nil, err
	}
	// set headers
	req.Header.Set("X-Shopify-Access-Token", token)
	req.Header.Set("Content-Type", "application/json")
	// req.Header.Set("Content-Length", fmt.Sprint(len(payload)))
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
