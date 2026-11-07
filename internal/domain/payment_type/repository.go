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

	DeleteByID(
		ctx context.Context,
		id uuid.UUID,
	) error

	FindByID(
		ctx context.Context,
		id uuid.UUID,
	) (paymentTypeEntity *Entity, err error)

	FindPagination(
		ctx context.Context,
		search string,
		offset int,
		limit int,
		sort string,
		sortBy string,
	) (paymentTypeEntities []Entity, totalData int64, err error)
}
