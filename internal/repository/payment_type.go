package repository

import (
	paymentType "basic-template-service/internal/domain/payment_type"
	"context"

	"gorm.io/gorm"
)

type PaymentTypeRepository struct {
	DB *gorm.DB
}

func (repo *PaymentTypeRepository) FindAll(ctx context.Context) (paymentTypeEntities []paymentType.Entity, err error) {

	err = repo.DB.
		WithContext(ctx).
		Select("id", "name").
		Find(&paymentTypeEntities).
		Error

	if err != nil {
		return nil, err
	}

	return paymentTypeEntities, nil
}
