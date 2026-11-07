package payment_type

import (
	"context"

	"github.com/google/uuid"
)

type IService interface {
	GetList(
		ctx context.Context,
	) ([]Entity, error)

	Save(
		ctx context.Context,
		paymentType *Entity,
	) error

	Update(
		ctx context.Context,
		paymentType *Entity,
	) error

	DeleteByID(
		ctx context.Context,
		id uuid.UUID,
	) error

	GetDetail(
		ctx context.Context,
		id uuid.UUID,
	) (*Entity, error)
}
