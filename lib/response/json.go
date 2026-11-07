package response

import (
	"basic-template-service/lib/errors"

	"github.com/gin-gonic/gin"
)

func SendSuccess(ctx *gin.Context, httpCode int, message string, data interface{}) {
	ctx.JSON(httpCode, SuccessResponse{
		Message: message,
		Data:    data,
	})
}

func SendSuccessWithMeta(ctx *gin.Context, httpCode int, message string, data interface{}, meta PaginationMeta) {
	ctx.JSON(httpCode, SuccessResponse{
		Message: message,
		Data:    data,
		Meta:    meta,
	})
}

func SendError(ctx *gin.Context, httpCode int, errorCode errors.Code, message string) {
	ctx.JSON(httpCode, ErrorResponse{
		ErrorCode: errorCode,
		Message:   message,
	})
}

func SendBadRequest(ctx *gin.Context, errorCode errors.Code, message string, validationErrors []ValidationErrorField) {
	ctx.JSON(400, BadRequestResponse{
		ErrorCode: errorCode,
		Message:   message,
		Errors:    validationErrors,
	})
}
