package config

import (
	"encoding/json"
	"log"
	"os"
)

type Config struct {
	DataSourceName string `env:"DSN"`
	DriverName     string `env:"DRIVER_NAME"`
}

func GetConfig() (Config, error) {

	file, _ := os.Open("pkg/config/tsconfig.json")
	decoder := json.NewDecoder(file)
	config := new(Config)
	err := decoder.Decode(&config)
	if err != nil {
		log.Fatalln(err, "config - 20")
	}
	return Config{
		DataSourceName: config.DataSourceName,
		DriverName:     config.DriverName,
	}, nil
}
