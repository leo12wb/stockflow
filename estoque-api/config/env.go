package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string
	DBSSLMode  string
	JWTSecret  string
	Port       string
	Env        string
}

var AppConfig Config

func LoadConfig() {
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		log.Println("Arquivo .env não encontrado, usando variáveis de ambiente")
	}

	AppConfig = Config{
		DBHost:     viper.GetString("DB_HOST"),
		DBPort:     viper.GetString("DB_PORT"),
		DBUser:     viper.GetString("DB_USER"),
		DBPassword: viper.GetString("DB_PASSWORD"),
		DBName:     viper.GetString("DB_NAME"),
		DBSSLMode:  viper.GetString("DB_SSLMODE"),
		JWTSecret:  viper.GetString("JWT_SECRET"),
		Port:       viper.GetString("PORT"),
		Env:        viper.GetString("ENV"),
	}

	if AppConfig.Port == "" {
		AppConfig.Port = "8080"
	}
	if AppConfig.JWTSecret == "" {
		AppConfig.JWTSecret = "default-secret-change-in-production"
	}
}
