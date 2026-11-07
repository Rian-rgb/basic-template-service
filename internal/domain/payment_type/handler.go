package payment_type

import (
	"github.com/gin-gonic/gin"
)

type IHandler interface {
	GetList(ctx *gin.Context)
	Create(ctx *gin.Context)
}
