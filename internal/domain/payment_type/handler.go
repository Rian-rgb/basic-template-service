package payment_type

import (
	"github.com/gin-gonic/gin"
)

type IHandler interface {
	GetAll(ctx *gin.Context)
}
