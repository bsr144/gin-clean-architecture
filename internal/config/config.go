package config

import (
	"log"

	"github.com/spf13/viper"
)

func NewDriverConfig() *DriverConfig {
	var driverConfig = new(DriverConfig)

	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.AddConfigPath("./config_files")

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			log.Fatalf("Error while loading the config file: %s", err)
		}
	}

	if err := viper.Unmarshal(driverConfig); err != nil {
		log.Fatalf("Error while unmarshal config: %s", err)
	}
	return driverConfig
}
