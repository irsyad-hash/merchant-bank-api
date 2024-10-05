package config

import (
	"encoding/json"
	"log"
	"os"
)

type Config struct {
	Port      string `json:"port"`
	JWTSecret string `json:"jwt_secret"`
}

func LoadConfig() *Config {
	file, err := os.Open("config/config.json")
	if err != nil {
		log.Fatal("Cannot open config file:", err)
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	config := &Config{}
	err = decoder.Decode(config)
	if err != nil {
		log.Fatal("Cannot decode config JSON:", err)
	}

	return config
}