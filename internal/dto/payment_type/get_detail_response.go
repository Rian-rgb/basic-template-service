package payment_type

import (
	"github.com/google/uuid"
)

type GetDetailResponse struct {
	ID       uuid.UUID `json:"id"`
	Name     string    `json:"name"`
	IsActive bool      `json:"is_active"`
}
