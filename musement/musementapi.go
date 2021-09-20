// Package musement abstract musement api functions
package musement

import (
	"encoding/json"
	"fmt"
	"musement/config"
	log "musement/logger"
	"musement/utils"
)

// Cities Struct to save cities information
type Cities struct {
	Cities []City
}

// City Struct to save city information
type City struct {
	Id int `json:"id"`
	Uuid string `json:"uuid"`
	Name string `json:"name"`
	Latitude float32 `json:"latitude"`
	Longitude float32 `json:"longitude"`
}

// ApiService is an interface to abstract musement api functions
type ApiService interface {
	GetCities() (*Cities, error)
}

type service struct {
	c utils.Client
}

// NewMusementApiService creates new service to use musement api function
func NewMusementApiService(c utils.Client) ApiService {
	return &service{c}
}

// GetCities function gets all cities from musmenet api
func (s *service) GetCities() (*Cities, error) {
	url := fmt.Sprintf("%s/cities", config.Conf.MusementAPI)
	response, _, err := s.c.GetDataFromUrl(url)
	if err != nil {
		log.Errorf("Could not get cities from api, error: %s", err)
		return nil, err
	}

	var result []City
	err = json.Unmarshal(response, &result)
	if err != nil {
		log. Errorf("Could not parse json from api. Error: %s", err)
		return nil, err
	}

	return &Cities{result}, nil
}
