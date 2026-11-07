package repository

import (
	paymentType "basic-template-service/internal/domain/payment_type"
	"basic-template-service/lib/utils"
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

func (repo *PaymentTypeRepository) Save(
	ctx context.Context,
	financialTransaction *paymentType.Entity,
) error {

	financialTransaction.ID = utils.GenerateUUID()
	return repo.DB.WithContext(ctx).Create(financialTransaction).Error
}

func (repo *PaymentTypeRepository) ExistsByName(
	ctx context.Context,
	name string,
) (bool, error) {
	var count int64

	err := repo.DB.
		Model(&paymentType.Entity{}).
		WithContext(ctx).
		Where("name = ?", name).
		Count(&count).
		Error

	if err != nil {
		return false, err
	}

	return count > 0, nil
}
