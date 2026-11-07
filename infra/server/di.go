package server

import (
	paymentType "basic-template-service/internal/domain/payment_type"
	"basic-template-service/internal/handler"
	"basic-template-service/internal/repository"
	"basic-template-service/internal/service"
)

type Dependency struct {
	PaymentTypeHdl paymentType.IHandler
}

func DependencyInject(appDeps *AppDependencies) *Dependency {

	// Repo
	paymentTypeRepo := &repository.PaymentTypeRepository{
		DB: appDeps.PostgresDB,
	}

	// Svc
	paymentTypeSvc := &service.PaymentTypeService{
		PaymentTypeRepo: paymentTypeRepo,
	}

	// Hdl
	paymentTypeHdl := &handler.PaymentTypeHandler{
		PaymentTypeSvc: paymentTypeSvc,
	}

	return &Dependency{
		PaymentTypeHdl: paymentTypeHdl,
	}
}
