package musement

import (
	"encoding/json"
	"fmt"
	"musement/config"
	log "musement/logger"
	"musement/utils"
)

type Cities struct {
	Cities []City
}

type City struct {
	Id int `json:"id"`
	Uuid string `json:"uuid"`
	Name string `json:"name"`
	Latitude float32 `json:"latitude"`
	Longitude float32 `json:"longitude"`
}

type ApiService interface {
	GetCities() (*Cities, error)
}

type service struct {
	c utils.Client
}

func NewMusementApiService(c utils.Client) ApiService {
	return &service{c}
}

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
