package utils

import (
	"trains/models"

	"github.com/spf13/viper"
)

// LoadConfig loads config from .env file
func LoadConfig(path string) (config models.Config, configErr error) {
	viper.AddConfigPath(path)
	viper.SetConfigName(".env")
	viper.SetConfigType("env")
	viper.AutomaticEnv()

	configErr = viper.ReadInConfig()
	if configErr != nil {
		return
	}
	configErr = viper.Unmarshal(&config)
	return
}
