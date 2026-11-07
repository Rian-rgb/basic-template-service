package handler

import (
	"net/http"

	"github.com/google/uuid"

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

	paymentTypeEntity := req.ToEntity()
	err := hdl.PaymentTypeSvc.Save(ctx.Request.Context(), paymentTypeEntity)
	if err != nil {
		errors.HandleServiceError(ctx, err)
		return
	}

	response.SendSuccess(ctx, http.StatusCreated, response.SuccessMessage, nil)
}

// @Summary		Update A Payment Type
// @Description	Updates a payment type and validates that the name is not already used by another payment type.
// @Tags		Payment Type
// @Accept		json
// @Produce		json
//
// @Param       id              path    string                      	true    "ID"
// @Param		request			body	paymentTypeDTO.UpdateRequest	true	"Request Body"
//
// @Success		200	{object}	response.SuccessResponse					"Success"
// @Failure		400	{object}	response.BadRequestResponse					"Bad Request"
// @Failure		409	{object}	response.ErrorResponse						"Duplicate Reference"
// @Failure		500	{object}	response.ErrorResponse						"Internal Server Error"
//
// @Router		/payment-type/{id} [put]
func (hdl *PaymentTypeHandler) Update(ctx *gin.Context) {
	var (
		req                 paymentTypeDTO.UpdateRequest
		codeBadRequest      = appErrors.ErrCodeBadRequest
		errCodeUnauthorized = appErrors.ErrCodeUnauthorized
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

	idString := ctx.Param("id")
	id, err := uuid.Parse(idString)
	if err != nil {
		logger.Error("Invalid ID")
		response.SendError(ctx, errCodeUnauthorized.ToHTTPStatus(), errCodeUnauthorized, response.InvalidIDMessage)
		return
	}

	paymentTypeEntity := req.ToEntity(id)
	err = hdl.PaymentTypeSvc.Update(ctx.Request.Context(), paymentTypeEntity)
	if err != nil {
		errors.HandleServiceError(ctx, err)
		return
	}

	response.SendSuccess(ctx, http.StatusOK, response.SuccessMessage, nil)
}

// @Summary		Delete A Payment Type
// @Description	Deletes a payment type by ID.
// @Tags		Payment Type
// @Accept		json
// @Produce		json
//
// @Param       id              path    string                      true    "ID"
//
// @Success		200	{object}	response.SuccessResponse					"Success"
// @Failure		401	{object}	response.ErrorResponse						"Unauthorized"
// @Failure		500	{object}	response.ErrorResponse						"Internal Server Error"
//
// @Router		/payment-type/{id} [delete]
func (hdl *PaymentTypeHandler) Delete(ctx *gin.Context) {
	var (
		errCodeUnauthorized = appErrors.ErrCodeUnauthorized
	)

	idString := ctx.Param("id")
	id, err := uuid.Parse(idString)
	if err != nil {
		logger.Error("Invalid ID")
		response.SendError(ctx, errCodeUnauthorized.ToHTTPStatus(), errCodeUnauthorized, response.InvalidIDMessage)
		return
	}

	err = hdl.PaymentTypeSvc.DeleteByID(ctx.Request.Context(), id)
	if err != nil {
		errors.HandleServiceError(ctx, err)
		return
	}

	response.SendSuccess(ctx, http.StatusOK, response.SuccessMessage, nil)
}

// @Summary		Get Payment Type Detail
// @Description	Gets payment type detail by ID.
// @Tags		Payment Type
// @Accept		json
// @Produce		json
//
// @Param       id              path    string                      true    "ID"
//
// @Success		200	{object}	response.SuccessResponse{data=paymentTypeDTO.GetDetailResponse}		"Success"
// @Failure		401	{object}	response.ErrorResponse												"Unauthorized"
// @Failure		404	{object}	response.ErrorResponse												"Not Found"
// @Failure		500	{object}	response.ErrorResponse												"Internal Server Error"
//
// @Router		/payment-type/{id} [get]
func (hdl *PaymentTypeHandler) GetDetail(ctx *gin.Context) {
	var (
		errCodeUnauthorized = appErrors.ErrCodeUnauthorized
	)

	idString := ctx.Param("id")
	id, err := uuid.Parse(idString)
	if err != nil {
		logger.Error("Invalid ID")
		response.SendError(ctx, errCodeUnauthorized.ToHTTPStatus(), errCodeUnauthorized, response.InvalidIDMessage)
		return
	}

	paymentTypeEntity, err := hdl.PaymentTypeSvc.GetDetail(ctx.Request.Context(), id)
	if err != nil {
		errors.HandleServiceError(ctx, err)
		return
	}

	resp := paymentTypeDTO.GetDetailResponse{
		ID:       paymentTypeEntity.ID,
		Name:     paymentTypeEntity.Name,
		IsActive: paymentTypeEntity.IsActive,
	}

	response.SendSuccess(ctx, http.StatusOK, response.SuccessMessage, resp)
}
