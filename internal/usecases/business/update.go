package business

import (
	"context"

	domain "github.com/tapiaw38/reservation-service-be/internal/domain/business"
	"github.com/tapiaw38/reservation-service-be/internal/platform/appcontext"
)

type (
	UpdateUsecase interface {
		Update(context.Context, string, domain.Business) (*UpdateOutput, error)
	}

	updateUsecase struct {
		contextFactory appcontext.Factory
	}

	UpdateOutput struct {
		Data BusinessOutputData `json:"data"`
	}
)

func NewUpdateUsecase(contextFactory appcontext.Factory) UpdateUsecase {
	return &updateUsecase{
		contextFactory: contextFactory,
	}
}

func (u *updateUsecase) Update(ctx context.Context, id string, input domain.Business) (*UpdateOutput, error) {
	app := u.contextFactory()

	// TODO: crear un endpoint que solo devuelva los servicios de un negocio
	business, err := app.Repositories.Business.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}

	services := updateBusinessServices(*business, input)
	input.Services = services

	_, err = app.Repositories.Business.Update(ctx, id, input)
	if err != nil {
		return nil, err
	}

	businessUpdated, err := app.Repositories.Business.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return &UpdateOutput{
		Data: toBusinessOutputData(*businessUpdated),
	}, nil
}

func updateBusinessServices(business domain.Business, input domain.Business) []domain.Service {
	existingServicesMap := make(map[string]domain.Service)
	for _, service := range business.Services {
		existingServicesMap[service.ID] = service
	}

	for _, inputService := range input.Services {
		existingServicesMap[inputService.ID] = inputService
	}

	var updatedServices []domain.Service
	for _, service := range existingServicesMap {
		updatedServices = append(updatedServices, service)
	}

	return updatedServices
}
