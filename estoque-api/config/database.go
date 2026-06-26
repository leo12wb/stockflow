package config

import (
	"fmt"
	"log"

	"github.com/seu-usuario/estoque-api/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func ConnectDatabase() {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=America/Sao_Paulo",
		AppConfig.DBHost,
		AppConfig.DBUser,
		AppConfig.DBPassword,
		AppConfig.DBName,
		AppConfig.DBPort,
		AppConfig.DBSSLMode,
	)

	var gormLogger logger.Interface
	if AppConfig.Env == "development" {
		gormLogger = logger.Default.LogMode(logger.Info)
	} else {
		gormLogger = logger.Default.LogMode(logger.Error)
	}

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: gormLogger,
	})
	if err != nil {
		log.Fatalf("Erro ao conectar ao banco de dados: %v", err)
	}

	err = db.AutoMigrate(
		&models.Usuario{},
		&models.Fornecedor{},
		&models.Produto{},
		&models.Movimentacao{},
	)
	if err != nil {
		log.Fatalf("Erro ao executar migrations: %v", err)
	}

	DB = db
	log.Println("Banco de dados conectado e migrations executadas com sucesso!")
}
