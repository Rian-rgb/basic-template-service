package payment_type

import (
	"github.com/gin-gonic/gin"
)

type IHandler interface {
	GetList(ctx *gin.Context)
	Create(ctx *gin.Context)
	Update(ctx *gin.Context)
	Delete(ctx *gin.Context)
	GetDetail(ctx *gin.Context)
	GetPagination(ctx *gin.Context)
}
