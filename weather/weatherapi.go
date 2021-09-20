// Package weather abstract weather api functions
package weather

import (
	"encoding/json"
	"fmt"
	"musement/config"
	log "musement/logger"
	"musement/utils"
)

// Weather struct saves weather information
type Weather struct {
	Forecasts Forecast `json:"forecast"`
}

// Forecast struct saves forecast information
type Forecast struct {
	ForecastsDay []ForecastDay `json:"forecastday"`
}

// ForecastDay struct saves the forecast for a day
type ForecastDay struct {
	ForecastDate string  `json:"date"`
	Day          DayData `json:"day"`
}

// DayData struct saves concrete information of a day
type DayData struct {
	Condition ConditionData `json:"condition"`
}

// ConditionData struct saves the condition information
type ConditionData struct {
	Text string `json:"text"`
}

// ApiService is an interface to abstract weather api functions
type ApiService interface {
	GetWeather(latitude float32, longitude float32, days int) (*Weather, error)
}

type service struct {
	c utils.Client
}

// NewWeatherApiService function creates a service to use weather api
func NewWeatherApiService(c utils.Client) ApiService {
	return &service{c}
}

// GetWeather functions gets the weather for a specific latitude-longitude for x days
func (s *service) GetWeather(latitude float32, longitude float32, days int) (*Weather, error) {
	url := fmt.Sprintf("%s/forecast.json?key=%s&q=%f,%f&days=%d", config.Conf.WeatherAPI, config.Conf.WeatherAPIKey, latitude, longitude, days)
	response, _, err := s.c.GetDataFromUrl(url)
	if err != nil {
		log.Errorf("Could not get cities from api, error: %s", err)
		return nil, err
	}

	var result Weather
	err = json.Unmarshal(response, &result)
	if err != nil {
		log. Errorf("Could not parse json from api. Error: %s", err)
		return nil, err
	}

	return &result, nil
}
