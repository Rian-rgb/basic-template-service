package infra

import (
	"basic-template-service/internal/domain/payment_type"
	"basic-template-service/lib/database"
	"log"

	"gorm.io/gorm"
)

func InitPostgresql(config *ServerConfig) *gorm.DB {
	dbCfg := database.PostgresConfig{
		Host:         *config.DBHost,
		User:         *config.DBUser,
		Password:     *config.DBPassword,
		DBName:       *config.DBName,
		Port:         *config.DBPort,
		MaxIdleConns: 5,
		MaxOpenConns: 20,
	}

	dbClient, err := database.NewPostgresClient(dbCfg)
	if err != nil {
		log.Fatalf("the App shutdown because failed connect DB: %v", err)
	}

	err = dbClient.AutoMigrate(
		&payment_type.Entity{},
	)
	if err != nil {
		log.Fatalf("failed doing auto migrate to internal tabel: %v", err)
	}

	return dbClient
}
