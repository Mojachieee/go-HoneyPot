package config

import (
	"encoding/json"
	"log"
	"os"
)

// Config is the struct for all configurable data
type Config struct {
	DBConfig Database `json:"db"`
}

// Database is the config struct for the database
type Database struct {
	Host     string `json:"host"`
	Name     string `json:"name"`
	Username string `json:"username"`
	Password string `json:"password"`
	Port     string `json:"port"`
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
