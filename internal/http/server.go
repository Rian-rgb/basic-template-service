package http

import (
	"basic-template-service/infra/server"
	"basic-template-service/internal/http/routes"
	"basic-template-service/lib/logger"

	"github.com/gin-gonic/gin"
)

func NewServer(
	dependency *server.Dependency,
	appDeps *server.AppDependencies,
) *gin.Engine {
	router := gin.Default()

	routes.RegisterRoutes(
		router,
		dependency,
		appDeps,
	)

	return router
}

func Serve(
	dependency *server.Dependency,
	appDeps *server.AppDependencies,
) {
	router := NewServer(
		dependency,
		appDeps,
	)

	port := *appDeps.Config.Port

	logger.Info("HTTP server listening on port %s", port)

	if err := router.Run(":" + port); err != nil {
		logger.Error("failed to start HTTP server: %s", err)
	}
}
