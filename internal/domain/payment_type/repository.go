package payment_type

import (
	"context"

	"github.com/google/uuid"
)

type IRepository interface {
	FindAll(
		ctx context.Context,
	) (paymentTypeEntities []Entity, err error)

	Save(
		ctx context.Context,
		financialTransaction *Entity,
	) error

	ExistsByName(
		ctx context.Context,
		name string,
	) (bool, error)

	Update(
		ctx context.Context,
		paymentType *Entity,
	) error

	ExistsByID(
		ctx context.Context,
		id uuid.UUID,
	) (bool, error)
}
