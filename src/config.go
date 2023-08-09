package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Config struct {
	server_host string `json:"server_host",envconfig:"SERVER_HOST"`
	server_port string `json:"server_port",envconfig:"SERVER_PORT"`
	db_host     string `json:"db_host",envconfig:"DB_HOST"`
	db_port     string `json:"db_port",envconfig:"DB_PORT"`
	db_username string `json:"db_username",envconfig:"DB_USERNAME"`
	db_password string `json:"db_password",envconfig:"DB_PASSWORD"`
}

var config Config = Config{}

func LoadConfiguration(file_path string) Config {
	config := Config{}

	configFile, err := os.Open(file_path)
	defer configFile.Close()
	if err != nil {
		fmt.Println(err.Error())
		fmt.Println("Failed to load config.json")
	}

	jsonParser := json.NewDecoder(configFile)
	jsonParser.Decode(&config)
	return config
}
