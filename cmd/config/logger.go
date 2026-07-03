package config

import (
	"github.com/spf13/viper"
)

func LoadEnv() {
	logger := GetLogger()

	viper.AddConfigPath(".")
	viper.SetConfigFile(".env")
	viper.SetConfigType("env")

	if err := viper.ReadInConfig(); err != nil {
		viper.AutomaticEnv()
		logger.Debug().Msgf("Failed to load env file. Env Loaded from AutomaticEnv(). Error: %v", err)
	} else {
		logger.Info().Msg("Env Loaded from .env")
	}
}
