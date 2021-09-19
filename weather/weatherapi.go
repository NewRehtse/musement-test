package weather

import (
	"encoding/json"
	"fmt"
	"musement/config"
	log "musement/logger"
	"musement/utils"
)

type Weather struct {
	Forecasts Forecast `json:"forecast"`
}

type Forecast struct {
	ForecastsDay []ForecastDay `json:"forecastday"`
}

type ForecastDay struct {
	ForecastDate string  `json:"date"`
	Day          DayData `json:"day"`
}

type DayData struct {
	Condition ConditionData `json:"condition"`
}

type ConditionData struct {
	Text string `json:"text"`
}

type ApiService interface {
	GetWeather(latitude float32, longitude float32, days int) (*Weather, error)
}

type service struct {
	c utils.Client
}

func NewWeatherApiService(c utils.Client) ApiService {
	return &service{c}
}

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
