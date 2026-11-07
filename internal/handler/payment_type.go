package handler

import (
	"net/http"

	paymentType "basic-template-service/internal/domain/payment_type"
	paymentTypeDTO "basic-template-service/internal/dto/payment_type"
	"basic-template-service/internal/errors"
	"basic-template-service/lib/response"

	"github.com/gin-gonic/gin"
)

type PaymentTypeHandler struct {
	PaymentTypeSvc paymentType.IService
}

// @Summary		Get Payment Types
// @Description	Gets all available payment types.
// @Tags		Payment Type
// @Accept		json
// @Produce		json
//
// @Success		200	{object}	response.SuccessResponse{data=[]paymentTypeDTO.GetListResponse}		"Success"
// @Failure		500	{object}	response.ErrorResponse												"Internal Server Error"
//
// @Security	BearerAuth
// @Router		/payment-type/list [get]
func (hdl *PaymentTypeHandler) GetList(ctx *gin.Context) {

	paymentTypeEntities, err := hdl.PaymentTypeSvc.GetAll(ctx.Request.Context())
	if err != nil {
		errors.HandleServiceError(ctx, err)
		return
	}

	paymentTypeListDTO := make([]paymentTypeDTO.GetListResponse, 0, len(paymentTypeEntities))

	for _, user := range paymentTypeEntities {
		paymentTypeListDTO = append(paymentTypeListDTO, paymentTypeDTO.GetListResponse{
			ID:   user.ID,
			Name: user.Name,
		})
	}

	resp := paymentTypeListDTO

	response.SendSuccess(ctx, http.StatusOK, response.SuccessMessage, resp)
}
