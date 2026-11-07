package payment_type

import (
	paymentType "basic-template-service/internal/domain/payment_type"
	"basic-template-service/lib/response"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

type UpdateRequest struct {
	Name string `json:"name" validate:"required"`
}

func (req UpdateRequest) Validate() []response.ValidationErrorField {
	v := validator.New()
	err := v.Struct(req)

	return response.MapValidationErrors(err)
}

func (req *UpdateRequest) ToEntity(
	id uuid.UUID,
) *paymentType.Entity {
	return &paymentType.Entity{
		ID:   id,
		Name: req.Name,
	}
}
