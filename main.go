package main

import (
	"fmt"
	"musement/config"
	"musement/musement"
	"musement/utils"
	"musement/weather"
	"net/http"
)

func main() {
	config.LoadConfig()
	client := utils.NewApiClient(http.Client{})
	musementAPIService := musement.NewMusementApiService(client)
	weatherAPIService := weather.NewWeatherApiService(client)

	cities, _ := musementAPIService.GetCities()
	days := 2
	for _, city := range cities.Cities {
		w, _ := weatherAPIService.GetWeather(city.Latitude, city.Longitude, days)
		str := fmt.Sprintf("Processed city %s | %s - %s", city.Name, w.Forecasts.ForecastsDay[0].Day.Condition.Text, w.Forecasts.ForecastsDay[1].Day.Condition.Text)
		fmt.Println(str)
		//break
	}
}
