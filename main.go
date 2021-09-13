package main

import (
	"fmt"
	"musement/config"
	"musement/musement"
	"musement/utils"
	"musement/weather"
	"net/http"
	"strings"
)

const defaultConditionsDay = 2

func main() {
	config.LoadConfig()
	client := utils.NewApiClient(http.Client{})
	musementAPIService := musement.NewMusementApiService(client)
	weatherAPIService := weather.NewWeatherApiService(client)

	cities, _ := musementAPIService.GetCities()
	days := defaultConditionsDay
	for _, city := range cities.Cities {
		w, _ := weatherAPIService.GetWeather(city.Latitude, city.Longitude, days)
		conditions := getForecastFromWeather(w)
		conditionsStr := strings.Join(conditions, " - ")
		str := fmt.Sprintf("Processed city %s | %s", city.Name, conditionsStr)
		fmt.Println(str)
	}
}

func getForecastFromWeather(w *weather.Weather) []string {
	fDay := w.Forecasts.ForecastsDay
	var conditions []string
	for _, f := range fDay {
		conditions = append(conditions, f.Day.Condition.Text)
	}

	return conditions
}
