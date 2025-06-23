package config

import (
	"log"

	"github.com/go-playground/validator/v10"
	"github.com/spf13/viper"
)

type Config struct {
	ServerAdr  string `mapstructure:"SERVER_ADR" validate:"required"`
	DBHost     string `mapstructure:"DB_HOST" validate:"required"`
	DBName     string `mapstructure:"DB_NAME" validate:"required"`
	DBUser     string `mapstructure:"DB_USER" validate:"required"`
	DbPort     string `mapstructure:"DB_PORT" validate:"required"`
	DBPassword string `mapstructure:"DB_PASSWORD" validate:"required"`
	Key        string `mapstructure:"KEY" validate:"required"`
}

func LoadConfig() (Config, error) {
	var config Config

	viper.SetConfigFile("../.env")
	viper.SetConfigType("env")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error loading the config file: %v", err)
		return config, err
	}

	if err := viper.Unmarshal(&config); err != nil {
		return config, err
	}

	if err := validator.New().Struct(&config); err != nil {
		return config, err
	}

	return config, nil
}

// SERVER_ADRESS= localhost:8080
// DB_HOST= 127.0.0.1
// DB_NAME= ECOMMERCE
// DB_USER= root
// DB_PORT= 3306
// DB_PASSWORD= new_password
// KEY= secret_user_key
