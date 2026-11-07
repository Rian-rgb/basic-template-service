package routes

import (
	"basic-template-service/infra/server"

	"github.com/gin-gonic/gin"
)

func registerPaymentTypeRoutes(
	api *gin.RouterGroup,
	dependency *server.Dependency,
	appDeps *server.AppDependencies,
) {

	paymentType := api.Group("/payment-type")

	paymentType.GET(
		"/list",
		dependency.PaymentTypeHdl.GetAll,
	)
}
