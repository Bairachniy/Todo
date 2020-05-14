package config

import (
	"encoding/json"
	"os"
)

type Config struct {
	DataSourceName string `env:dataSourceName`
	DriverName     string `env:driverName`
}

func GetConfig() (Config, error) {
	file, _ := os.Open("pkg/config/tsconfig.json")
	decoder := json.NewDecoder(file)
	config := new(Config)
	err := decoder.Decode(&config)
	//if err != nil {
	//	panic(err)
	//}
	return *config, err
}
