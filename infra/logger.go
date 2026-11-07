package infra

import (
	logger "basic-template-service/lib/logger"
)

func InitLogger(config *ServerConfig) {
	env := *config.ENV
	if env == "" {
		env = "production"
	}

	logger.SetupLogger(env)
}
