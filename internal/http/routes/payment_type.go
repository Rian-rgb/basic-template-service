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
		dependency.PaymentTypeHdl.GetList,
	)

	paymentType.POST(
		"",
		dependency.PaymentTypeHdl.Create,
	)

	paymentType.PUT(
		"/:id",
		dependency.PaymentTypeHdl.Update,
	)

	paymentType.DELETE(
		"/:id",
		dependency.PaymentTypeHdl.Delete,
	)

	paymentType.GET(
		"/:id",
		dependency.PaymentTypeHdl.GetDetail,
	)

	paymentType.GET(
		"/pagination",
		dependency.PaymentTypeHdl.GetPagination,
	)
}
