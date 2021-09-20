// Package utils has several functions useful for general purposes
package utils

import (
	"bytes"
	"errors"
	"io/ioutil"
	log "musement/logger"
	"net/http"
)

// Client interface abstract a client to get data from different resources
type Client interface {
	GetDataFromUrl(url string) ([]byte, int, error)
}

type client struct {
	c http.Client
}

var errorResourceNotFound = errors.New("Resource not found")
var errorServiceUnavailable = errors.New("Service unavailable")
var errorEmptyUrl = errors.New("URL is empty")

// NewApiClient creates a new Client to make http requests
func NewApiClient(c http.Client) Client {
	return &client{c}
}

// GetDataFromUrl gets info from a given url
func (c *client) GetDataFromUrl(url string) ([]byte, int, error) {
	if len(url) == 0 {
		log.Errorf("url is empty\n")
		return nil, 500, errorEmptyUrl
	}
	r, err := http.NewRequest("GET", url, bytes.NewBuffer([]byte{}))

	resp, err := c.c.Do(r)
	if err != nil {
		log.Errorf("cannot make request, error: %s\n", err)
		return nil, 500, err
	}

	if resp.StatusCode == 404 {
		log.Errorf("Error resource not found")
		return nil, resp.StatusCode, errorResourceNotFound
	}

	if resp.StatusCode == 503 {
		log.Errorf("Error 503")
		return nil, resp.StatusCode, errorServiceUnavailable
	}

	if resp.StatusCode != 200 {
		log.Errorf("Error getting data with statusCode %d and error %s", resp.StatusCode, err)
		return nil, resp.StatusCode, errorServiceUnavailable
	}

	body, err := ioutil.ReadAll(resp.Body)

	defer resp.Body.Close()

	if err != nil {
		log.Errorf("cannot read response, error: %v\n", err)
		return nil, 500, err
	}

	return body, resp.StatusCode, nil
}
