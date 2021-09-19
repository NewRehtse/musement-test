package musement

import (
	"fmt"
	"io/ioutil"
	"musement/utils"
	"os"
	"testing"
)

func TestGetCitiesFromMusement(t *testing.T) {
	jsonFile, err := os.Open("./mocks/cities.json")
	if err != nil {
		fmt.Println("Error opening mock file ./mocks/cities.json")
	}

	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	client := utils.NewApiMockClient(byteValue, 200, nil)

	s := NewMusementApiService(client)

	cities, err := s.GetCities()

	if err != nil {
		t.Error("Error was not expected")
	}

	if len(cities.Cities) != 4 {
		t.Errorf("Expected 4 cities but got %d", len(cities.Cities))
	}
}

func TestGetCitiesFromMusementAndGiveError404(t *testing.T) {
	client := utils.NewApiMockClient(nil, 404, nil)

	s := NewMusementApiService(client)

	_, err := s.GetCities()

	if err == nil {
		t.Error("Error should be not found resource")
	}
}

func TestGetCitiesFromMusementAndGiveError503(t *testing.T) {
	client := utils.NewApiMockClient(nil, 503, nil)

	s := NewMusementApiService(client)

	_, err := s.GetCities()

	if err == nil {
		t.Error("Error should be not found resource")
	}
}

func TestGetCitiesFromMusementAndGiveError400(t *testing.T) {
	client := utils.NewApiMockClient(nil, 400, nil)

	s := NewMusementApiService(client)

	_, err := s.GetCities()

	if err == nil {
		t.Error("Error should be not found resource")
	}
}
