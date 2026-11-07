package errors

import (
	"errors"

	appErrors "basic-template-service/lib/errors"
	"basic-template-service/lib/response"

	"github.com/gin-gonic/gin"
)

func HandleServiceError(ctx *gin.Context, err error) {
	var (
		errCodeDuplicateReference  = appErrors.ErrCodeDuplicateReference
		errCodeInternalServerError = appErrors.ErrCodeInternalServerError
		errCodeUnknown             = appErrors.ErrCodeUnknown
	)
	if err == nil {
		return
	}

	switch {

	case errors.Is(err, ErrAlreadyExists):
		response.SendError(ctx, errCodeDuplicateReference.ToHTTPStatus(), errCodeDuplicateReference, err.Error())

	case errors.Is(err, ErrInternalServerError):
		response.SendError(ctx, errCodeInternalServerError.ToHTTPStatus(), errCodeInternalServerError, err.Error())

	default:
		response.SendError(ctx, errCodeUnknown.ToHTTPStatus(), errCodeUnknown, response.InternalServerErrorMessage)
	}
}
