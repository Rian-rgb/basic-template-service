package payment_type

import (
	"github.com/google/uuid"
)

type GetAllResponse struct {
	ID   uuid.UUID `json:"id"`
	Name string    `json:"name"`
}
