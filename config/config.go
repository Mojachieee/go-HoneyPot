package config

import (
	"encoding/json"
	"log"
	"os"
)

// Config is the struct for all configurable data
type Config struct {
	DBConfig DatabaseConfig `json:"db"`
}

// DatabaseConfig is the config struct for the database
type DatabaseConfig struct {
	Host     string `json:"host"`
	Name     string `json:"name"`
	Username string `json:"username"`
	Password string `json:"password"`
}

// Read reads the configuration file and returns a struct of it
func Read() Config {
	file, err := os.Open("config.json")
	if err != nil {
		log.Fatalf("Could not read config: %v", err)
	}
	defer file.Close()
	decoder := json.NewDecoder(file)
	config := Config{}
	err = decoder.Decode(&config)
	if err != nil {
		log.Fatalf("Could not read config: %v", err)
	}
	return config
}
