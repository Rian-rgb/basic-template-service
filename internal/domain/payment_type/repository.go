package payment_type

import "context"

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
}
