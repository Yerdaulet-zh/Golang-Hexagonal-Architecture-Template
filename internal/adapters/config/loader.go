package config

import (
	"errors"
	"fmt"
	"log/slog"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

func InitConfig() error {
	viper.SetConfigName("configs")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.AddConfigPath("./configs")

	if err := viper.ReadInConfig(); err != nil {
		var fileLookupError viper.ConfigFileNotFoundError
		if errors.As(err, &fileLookupError) {
			return fmt.Errorf("configuration file not found: %w", err)
		}
		return err
	}

	if err := godotenv.Load(); err != nil {
		slog.Warn("No .env file found, proceeding with environment variables")
	}
	return nil
}
