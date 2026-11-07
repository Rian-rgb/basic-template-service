package payment_type

import "context"

type IService interface {
	GetList(
		ctx context.Context,
	) ([]Entity, error)

	Save(
		ctx context.Context,
		paymentType *Entity,
	) error
}
