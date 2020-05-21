package config

import (
	"encoding/json"
	"os"
)

type Config struct {
	DataSourceName string `env:"DSD"`
	DriverName     string `env:"DRIVER_NAME"`
}

func GetConfig() (Config, error) {
	file, _ := os.Open("pkg/config/tsconfig.json")
	decoder := json.NewDecoder(file)
	config := new(Config)
	err := decoder.Decode(&config)

	return *config, err
}
