package handler

import (
	"net/http"

	paymentType "basic-template-service/internal/domain/payment_type"
	paymentTypeDTO "basic-template-service/internal/dto/payment_type"
	"basic-template-service/internal/errors"
	"basic-template-service/lib/logger"
	"basic-template-service/lib/response"

	appErrors "basic-template-service/lib/errors"

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

	paymentTypeEntities, err := hdl.PaymentTypeSvc.GetList(ctx.Request.Context())
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

// @Summary		Create A New Payment Type
// @Description	Creates a new payment type.
// @Tags		Payment Type
// @Accept		json
// @Produce		json
//
// @Param		request			body	paymentTypeDTO.CreateRequest	true	"Request Body"
//
// @Success		201	{object}	response.SuccessResponse					"Created"
// @Failure		409	{object}	response.ErrorResponse						"Duplicate Reference"
// @Failure		500	{object}	response.ErrorResponse						"Internal Server Error"
//
// @Security	BearerAuth
// @Router		/payment-type [post]
func (hdl *PaymentTypeHandler) Create(ctx *gin.Context) {
	var (
		req            paymentTypeDTO.CreateRequest
		codeBadRequest = appErrors.ErrCodeBadRequest
	)

	if err := ctx.ShouldBindJSON(&req); err != nil {
		logger.Error("failed to parse JSON request: %s", err)
		response.SendBadRequest(ctx, codeBadRequest, response.InvalidJSONFormatMessage, nil)
		return
	}

	errFields := req.Validate()
	if errFields != nil {
		logger.Warn("request body validation failed")
		response.SendBadRequest(ctx, codeBadRequest, response.InvalidRequestMessage, errFields)
		return
	}

	paymentType := req.ToEntity()
	err := hdl.PaymentTypeSvc.Save(ctx.Request.Context(), paymentType)
	if err != nil {
		errors.HandleServiceError(ctx, err)
		return
	}

	response.SendSuccess(ctx, http.StatusCreated, response.SuccessMessage, nil)
}
