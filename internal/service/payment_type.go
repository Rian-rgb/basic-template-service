package service

import (
	paymentType "basic-template-service/internal/domain/payment_type"
	internalErrors "basic-template-service/internal/errors"
	"basic-template-service/lib/logger"
	"context"
	"fmt"
)

type PaymentTypeService struct {
	PaymentTypeRepo paymentType.IRepository
}

func (svc *PaymentTypeService) GetList(ctx context.Context) ([]paymentType.Entity, error) {

	paymentTypesEntity, err := svc.PaymentTypeRepo.FindAll(ctx)
	if err != nil {
		logger.Error("failed to find all payment types: %s", err)
		return nil, internalErrors.ErrInternalServerError
	}

	return paymentTypesEntity, nil
}

func (svc *PaymentTypeService) Save(
	ctx context.Context,
	paymentType *paymentType.Entity,
) error {

	nameExists, err := svc.PaymentTypeRepo.ExistsByName(ctx, paymentType.Name)
	if err != nil {
		logger.Error("failed to verify name existence: %s", err)
		return internalErrors.ErrInternalServerError
	}

	if nameExists {
		logger.Error("payment type already exists")
		return fmt.Errorf(
			"payment type %w",
			internalErrors.ErrAlreadyExists,
		)
	}

	err = svc.PaymentTypeRepo.Save(ctx, paymentType)
	if err != nil {
		logger.Error("failed to save payment type: %s", err)
		return internalErrors.ErrInternalServerError
	}

	return nil
}

func (svc *PaymentTypeService) Update(
	ctx context.Context,
	paymentType *paymentType.Entity,
) error {

	idExists, err := svc.PaymentTypeRepo.ExistsByID(ctx, paymentType.ID)
	if err != nil {
		logger.Error("failed to verify id existence: %s", err)
		return internalErrors.ErrInternalServerError
	}

	if !idExists {
		logger.Error("ID not found")
		return fmt.Errorf(
			"payment type %w",
			internalErrors.ErrNotFound,
		)
	}

	nameExists, err := svc.PaymentTypeRepo.ExistsByName(ctx, paymentType.Name)
	if err != nil {
		logger.Error("failed to verify name existence: %s", err)
		return internalErrors.ErrInternalServerError
	}

	if nameExists {
		logger.Error("payment type already exists")
		return fmt.Errorf(
			"payment type %w",
			internalErrors.ErrAlreadyExists,
		)
	}

	err = svc.PaymentTypeRepo.Update(ctx, paymentType)
	if err != nil {
		logger.Error("failed to update payment type: %s", err)
		return internalErrors.ErrInternalServerError
	}

	return nil
}
