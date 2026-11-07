package payment_type

import "context"

type IRepository interface {
	FindAll(
		ctx context.Context,
	) (paymentTypeEntities []Entity, err error)
}
