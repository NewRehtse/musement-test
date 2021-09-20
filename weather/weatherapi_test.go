package weather

import (
	"errors"
	"fmt"
	"io/ioutil"
	"musement/utils"
	"os"
	"testing"
)

func TestGetWeatherFromCity(t *testing.T) {
	jsonFile, err := os.Open("./mocks/weather.json")
	if err != nil {
		fmt.Println("Error opening mock file ./mocks/weather.json")
	}
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	client := utils.NewApiMockClient(byteValue, 200, nil)

	s := NewWeatherApiService(client)

	latitude := float32(51.52)
	longitude := float32(-0.11)
	days := 2

	weather, err := s.GetWeather(latitude, longitude, days)

	if err != nil {
		t.Error("Error was not expected")
	}

	if weather == nil {
		t.Error("Weather should not be nil")
	}
}

func TestGetWeatherFromCityAndGetError(t *testing.T) {
	client := utils.NewApiMockClient(nil, 400, errors.New("Error returned"))

	s := NewWeatherApiService(client)

	latitude := float32(51.52)
	longitude := float32(-0.11)
	days := 2

	_, err := s.GetWeather(latitude, longitude, days)

	if err == nil {
		t.Error("Error was expected")
	}
}

func TestGetWeatherFromCityWrongJson(t *testing.T) {
	jsonFile, err := os.Open("./mocks/weather.txt")
	if err != nil {
		fmt.Println("Error opening mock file ./mocks/weather.json")
	}
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	client := utils.NewApiMockClient(byteValue, 200, nil)

	s := NewWeatherApiService(client)

	latitude := float32(51.52)
	longitude := float32(-0.11)
	days := 2

	weather, err := s.GetWeather(latitude, longitude, days)

	if weather != nil {
		t.Error("Weather should be nil")
	}

	if err == nil {
		t.Error("Error should not be nil")
	}
}
