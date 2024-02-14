package config

import (
	"encoding/json"
	"os"
)

type Config struct {
	Token      string `json:"token"`
	OutputDir  string `json:"output_directory"`
	ConfigType string `json:"config_type"`
}

func GetConfig(configPath string) (Config, error) {

	var c Config

	f, err := os.Open(configPath)
	if err != nil {
		return Config{}, err
	}
	defer f.Close()

	decoder := json.NewDecoder(f)

	err = decoder.Decode(&c)

	return c, nil

}
