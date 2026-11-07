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

	GetPagination(
		ctx context.Context,
		search string,
		offset int,
		limit int,
		sort string,
		sortBy string,
	) (paymentTypeEntities []Entity, totalPage int, totalData int64, err error)
}
