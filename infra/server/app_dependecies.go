package server

import (
	"basic-template-service/infra"

	"gorm.io/gorm"
)

type AppDependencies struct {
	Config     *infra.ServerConfig
	PostgresDB *gorm.DB
}
