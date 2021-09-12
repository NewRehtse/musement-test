package main

import (
	"fmt"
	"musement/config"
	"musement/musement"
	"musement/utils"
	"net/http"
)

func main() {
	fmt.Println("Hello, World!")

	config.LoadConfig()

	client := utils.NewApiClient(http.Client{})

	musementAPIService := musement.NewMusementApiService(client)

	cities, _ := musementAPIService.GetCities()
	fmt.Println(cities)
}
