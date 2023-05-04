package client

import (
	"errors"
	"fmt"
	"net/http"
)

type Client struct {
	BaseURL    string
	HTTPClient *http.Client
}

func (c *Client) SendRequest(req *http.Request) (*http.Response, error) {
	if c.HTTPClient == nil {
		return &http.Response{}, errors.New("no http client")
	}

	req.Header.Set("Accept", "application/json; charset=utf-8")
	req.Header.Set("Content-Type", "application/json; charset=utf-8")

	response, err := c.HTTPClient.Do(req)
	if err != nil {
		return &http.Response{}, fmt.Errorf("failed to send request: error: %w", err)
	}

	return response, nil
}
