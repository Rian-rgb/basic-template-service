package payment_type

import "context"

type IService interface {
	GetAll(ctx context.Context) ([]Entity, error)
}
