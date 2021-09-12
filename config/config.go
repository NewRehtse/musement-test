// Package config gets configuration from config.json
package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	log "musement/logger"
)

// Config represents the configuration information.
type Config struct {
	Environment   string `json:"environment"`
	MusementAPI   string `json:"musement-api"`
	WeatherAPI    string `json:"weather-api"`
	WeatherAPIKey string `json:"weather-api-key"`
	Debug         bool   `json:"debug"`
	LogFile       string `json:"log_file"`
}

// Conf contains the initialized configuration struct
var (
	Conf = Config{
		MusementAPI: "",
		WeatherAPI: "",
		WeatherAPIKey: "",
		Debug:     false,
		LogFile:   "",
	}
)

const DEFAULT_CONFIG_FILE = "config.json"

// LoadFileConfig loads the configuration from the specified filepath
func LoadConfig() {
	log.DebugEnabled = Conf.Debug
	log.LogFile = Conf.LogFile

	// Read config file
	dir, _ := filepath.Abs(filepath.Dir(os.Args[0]))
	configFilePath := dir + "/" + DEFAULT_CONFIG_FILE
	fmt.Println(configFilePath)
	if !fileExists(configFilePath) {
		configFilePath = "./config.json"
	}

	fmt.Println(configFilePath)
	log.Error(configFilePath)
	configFile, err := readFile(configFilePath)
	if nil == err {
		err = json.Unmarshal(configFile, &Conf)

		if nil != err {
			log.Errorf("Error: %v", err)
		}
	}

	if Conf.Debug {
		jsonConfig, _ := json.MarshalIndent(Conf, ">>> ", "  ")
		fmt.Println(">>> " + string(jsonConfig))
	}
	log.DebugEnabled = Conf.Debug
	log.LogFile = Conf.LogFile
	log.ReloadConfiguration()
}

func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

func readFile(filepath string) (fileContent []byte, err error) {
	log.Infof("LoadConfig file %s", filepath)

	fileContent, err = ioutil.ReadFile(filepath)

	if err != nil {
		log.Errorf("LoadConfig file error: %v", err)
	}

	return
}
