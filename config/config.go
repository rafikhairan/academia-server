package config

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/spf13/viper"
	"time"
)

type Config struct {
	AppName  string
	Database Database
	JWT
}

type Database struct {
	Username string
	Password string
	Host     string
	Port     string
	Name     string
}

type JWT struct {
	SigningMethod *jwt.SigningMethodHMAC
	Expiration    time.Duration
	SignatureKey  string
}

var AppConfig Config

func init() {
	viper.SetConfigFile(".env")

	if err := viper.ReadInConfig(); err != nil {
		panic("Error reading config file")
	}

	AppConfig = Config{
		JWT: JWT{
			SigningMethod: jwt.SigningMethodHS256,
			Expiration:    viper.GetDuration("JWT_EXPIRATION"),
			SignatureKey:  viper.GetString("JWT_SIGNATURE_KEY"),
		},
		Database: Database{
			Username: viper.GetString("DATABASE_USERNAME"),
			Password: viper.GetString("DATABASE_PASSWORD"),
			Host:     viper.GetString("DATABASE_HOST"),
			Port:     viper.GetString("DATABASE_PORT"),
			Name:     viper.GetString("DATABASE_NAME"),
		},
	}
}
