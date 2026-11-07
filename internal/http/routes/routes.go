package routes

import (
	"basic-template-service/infra/server"
	"basic-template-service/internal/http/constants"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(
	router *gin.Engine,
	dependency *server.Dependency,
	appDeps *server.AppDependencies,
) {
	api := router.Group(constants.APIBasePath)

	registerSwaggerRoutes(router)

	registerPaymentTypeRoutes(
		api,
		dependency,
		appDeps,
	)
}
