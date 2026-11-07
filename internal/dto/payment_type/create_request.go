package payment_type

import (
	"basic-template-service/lib/response"
	"strings"

	paymentType "basic-template-service/internal/domain/payment_type"

	"github.com/go-playground/validator/v10"
)

type CreateRequest struct {
	Name string `json:"name" validate:"required"`
}

func (req CreateRequest) Validate() []response.ValidationErrorField {
	v := validator.New()
	err := v.Struct(req)

	return response.MapValidationErrors(err)
}

func (req *CreateRequest) ToEntity() *paymentType.Entity {
	return &paymentType.Entity{
		Name: strings.ToLower(req.Name),
	}
}
