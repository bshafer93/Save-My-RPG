package smrpg

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"reflect"
	"savemyrpg/dal"
	"time"
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
	JWT_SECRET_KEY         string `json:"jwt_secret_key"`
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

func Init() bool {

	_, err := LoadConfiguration("/go/src/savemyrpgserver/config.json")
	if err != nil {
		return false
	}

	cert, err := tls.LoadX509KeyPair(config.SERVER_CERT, config.SERVER_KEY)
	if err != nil {
		return false
	}

	tls_config := &tls.Config{Certificates: []tls.Certificate{cert}}
	mux := http.NewServeMux()
	mux.HandleFunc("/", handler)
	mux.HandleFunc("/serverinfo", ServerInfoHandler)
	mux.HandleFunc("/getfullsave", SendFullFile)
	mux.HandleFunc("/login", Login)
	mux.HandleFunc("/rc", RetrieveAllJoinedCampaigns)
	mux.HandleFunc("/cs", RetrieveAllCampaignSaves)
	mux.HandleFunc("/jc", UserJoinCampaign)
	mux.HandleFunc("/rci", RetrieveCampaign)

	server = &http.Server{
		Addr:              ":" + config.SERVER_PORT,
		ReadTimeout:       30 * time.Second,
		WriteTimeout:      30 * time.Second,
		IdleTimeout:       30 * time.Second,
		ReadHeaderTimeout: 30 * time.Second,
		TLSConfig:         tls_config,
		Handler:           mux,
	}

	if !dal.Init() {
		return false
	}

	return true
}
