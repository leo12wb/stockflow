package config

import (
	"log"

	"go.uber.org/zap"
)

var Logger *zap.Logger

func InitLogger() {
	var err error
	if AppConfig.Env == "production" {
		Logger, err = zap.NewProduction()
	} else {
		Logger, err = zap.NewDevelopment()
	}
	if err != nil {
		log.Fatalf("Erro ao inicializar logger: %v", err)
	}
}
