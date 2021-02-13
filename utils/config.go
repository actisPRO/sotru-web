package utils

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

type Config struct {
	ServerAddress   string `json:"server_address"`
	SessionSecret   string `json:"session_secret"`
	DBHost          string `json:"db_host"`
	DBName          string `json:"db_name"`
	DBUser          string `json:"db_user"`
	DBPassword      string `json:"db_password"`
	DiscordOAuthURL string `json:"discord_oauth_url"`
	DiscordClient   string `json:"discord_client"`
	DiscordSecret   string `json:"discord_secret"`
}

// Loads configuration from the specified JSON file.
func ReadConfig(path string) (Config, error) {
	config := Config{}
	jsonFile, err := os.Open(path)
	if err != nil {
		return config, err
	}
	defer jsonFile.Close()

	bytes, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		return config, err
	}

	err = json.Unmarshal(bytes, &config)
	if err != nil {
		return config, err
	}

	return config, nil
}
