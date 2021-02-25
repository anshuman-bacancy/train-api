package utils

import "github.com/spf13/viper"

type Config struct {
	ConnString string `mapstructure:"CONNECTION_STRING"`
	DbName string `mapstructure:"DATABASE_NAME"`
	Collection string `mapstructure:"COLLECTION"`
}

func LoadConfig(path string) (config Config, configErr error) {
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
