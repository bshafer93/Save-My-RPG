package smrpg

import (
	"encoding/json"
	"fmt"
	"os"
	"reflect"
)

type Config struct {
	SERVER_HOST            string `json:"server_host"`
	SERVER_PORT            string `json:"server_port"`
	SERVER_CERT            string `json:"server_cert"`
	SERVER_KEY             string `json:"server_key"`
	DB_HOST                string `json:"db_host"`
	DB_PORT                string `json:"db_port"`
	DB_USERNAME            string `json:"db_username"`
	DB_PASSWORD            string `json:"db_password"`
	SAVES_PATH             string `json:"saves_path"`
	BUNNYNET_READ_API_KEY  string `json:"bunnynet_read_api_key"`
	BUNNYNET_WRITE_API_KEY string `json:"bunnynet_write_api_key"`
}

var config Config = Config{}

func LoadConfiguration(file_path string) (*Config, error) {

	configFile, err := os.ReadFile(file_path)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	json.Unmarshal(configFile, &config)
	fmt.Println("Config File Loaded!")
	return &config, nil
}

func PrintConfig() {
	values := reflect.ValueOf(config)
	types := values.Type()
	for i := 0; i < values.NumField(); i++ {
		fmt.Println(types.Field(i).Name, ": ", values.Field(i))
	}
}
