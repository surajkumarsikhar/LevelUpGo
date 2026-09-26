package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Config struct {
	Server   ServerConfig   `json:"server"`
	Database DatabaseConfig `json:"database"`
}

type ServerConfig struct {
	Host string `json:"host"`
	Port int    `json:"port"`
}

type DatabaseConfig struct {
	Host     string `json:"host"`
	Port     int    `json:"port"`
	Database string `json:"database"`
}

// LoadConfig reads a JSON config file and returns a Config struct.
func LoadConfig(filename string) (*Config, error) {
	// TODO: Read file with os.ReadFile
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	config := Config{}
	// TODO: Unmarshal JSON into Config struct
	err = json.Unmarshal(data, &config)
	if err != nil {
		return nil, err
	}
	// TODO: Return pointer to config
	return &config, nil
}

// SaveConfig writes a Config struct to a JSON file with pretty printing.
func SaveConfig(config *Config, filename string) error {
	// TODO: Marshal config with json.MarshalIndent(config, "", "  ")
	jsonByte, err := json.MarshalIndent(config, "", "  ")
	if err != nil {
		return err
	}
	// TODO: Write to file with os.WriteFile
	err = os.WriteFile(filename, jsonByte, 0644)
	if err != nil {
		return err
	}
	return nil
}

func main() {
	config := &Config{
		Server: ServerConfig{
			Host: "localhost",
			Port: 8080,
		},
		Database: DatabaseConfig{
			Host:     "localhost",
			Port:     5432,
			Database: "myapp",
		},
	}

	err := SaveConfig(config, "config.json")
	if err != nil {
		fmt.Printf("Error saving: %v\n", err)
		return
	}
	fmt.Println("Config saved")

	loaded, err := LoadConfig("config.json")
	if err != nil {
		fmt.Printf("Error loading: %v\n", err)
		return
	}

	fmt.Printf("Loaded: %s:%d\n", loaded.Server.Host, loaded.Server.Port)
}
