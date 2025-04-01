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

	services, err := app.Repositories.Business.FindServices(ctx, id)
	if err != nil {
		return nil, err
	}

	servicesUpdated := updateBusinessServices(services, input)
	input.Services = servicesUpdated

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

func updateBusinessServices(services []domain.Service, input domain.Business) []domain.Service {
	existingServicesMap := make(map[string]domain.Service)
	for _, service := range services {
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
