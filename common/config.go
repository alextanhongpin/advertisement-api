package common

import (
	"encoding/json"
	"fmt"
	"os"
)

type Configuration struct {
	Port string `json:"port"`
}

var config Configuration

func init() {
	config = loadConfig()
}

func loadConfig() Configuration {
	file, _ := os.Open("common/config.json")
	decoder := json.NewDecoder(file)
	config := Configuration{}
	err := decoder.Decode(&config)
	if err != nil {
		fmt.Println("Error reading config:", err)
	}
	return config
}

func GetConfig() Configuration {
	if (Configuration{}) == config {
		config = loadConfig()
		return config
	}
	return config
}
