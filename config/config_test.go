package config

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"
)


func TestConfigFileGet(t *testing.T) {
	dir, _ := filepath.Abs(filepath.Dir(os.Args[0]))
	copy("./config_test.json", dir + "/config.json")

	LoadConfig()

	if Conf.WeatherAPI != "api-weather-example" {
		t.Errorf("Weather api config expected %s got %s.", "api-weather-example", Conf.WeatherAPI)
	}

	remove(Conf.LogFile)
}

func remove(file string) {
	err := os.Remove(file)
	if err != nil {
		fmt.Println("Could not remove config file: " + file)
	}
}

func copy(sourceFile string, destinationFile string) {
	input, err := ioutil.ReadFile(sourceFile)
	if err != nil {
		fmt.Println(err)
		return
	}

	err = ioutil.WriteFile(destinationFile, input, 0644)
	if err != nil {
		fmt.Println("Error creating", destinationFile)
		fmt.Println(err)
		return
	}
}