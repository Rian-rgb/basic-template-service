package routes

import (
	_ "basic-template-service/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title           Basic Template API
// @version         0.0
// @description     API for managing basic CRUD API.
// @description     <br/><b>Developer:</b> Muhammad Brilian Satria Utama
// @description     <b>Environment:</b> Development
// @host            localhost:8080
// @BasePath        /api/v1
func registerSwaggerRoutes(router *gin.Engine) {
	router.GET(
		"/swagger/*any",
		ginSwagger.WrapHandler(swaggerFiles.Handler),
	)
}
