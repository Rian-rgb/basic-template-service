package service

import (
	paymentType "basic-template-service/internal/domain/payment_type"
	internalErrors "basic-template-service/internal/errors"
	"basic-template-service/lib/logger"
	"context"
)

type PaymentTypeService struct {
	PaymentTypeRepo paymentType.IRepository
}

func (svc *PaymentTypeService) GetAll(ctx context.Context) ([]paymentType.Entity, error) {

	paymentTypesEntity, err := svc.PaymentTypeRepo.FindAll(ctx)
	if err != nil {
		logger.Error("failed to find all payment types: %s", err)
		return nil, internalErrors.ErrInternalServerError
	}

	return paymentTypesEntity, nil
}
