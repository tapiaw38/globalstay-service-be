package hotel

import (
	"context"

	domain "github.com/tapiaw38/globalstay-service-be/internal/domain/hotel"
	"github.com/tapiaw38/globalstay-service-be/internal/platform/appcontext"
)

type (
	UpdateUsecase interface {
		Execute(context.Context, string, domain.Hotel) (*UpdateOutput, error)
	}

	updateUsecase struct {
		contextFactory appcontext.Factory
	}

	UpdateOutput struct {
		Data HotelOutputData `json:"data"`
	}
)

func NewUpdateUsecase(contextFactory appcontext.Factory) UpdateUsecase {
	return &updateUsecase{
		contextFactory: contextFactory,
	}
}

func (u *updateUsecase) Execute(ctx context.Context, id string, input domain.Hotel) (*UpdateOutput, error) {
	app := u.contextFactory()

	_, err := app.Repositories.Hotel.Update(ctx, id, input)
	if err != nil {
		return nil, err
	}

	hotelUpdated, err := app.Repositories.Hotel.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return &UpdateOutput{
		Data: toHotelOutputData(*hotelUpdated),
	}, nil
}

// func updateHotelServices(services []domain.Service, input domain.Hotel) []domain.Service {
// 	existingServicesMap := make(map[string]domain.Service)
// 	for _, service := range services {
// 		existingServicesMap[service.ID] = service
// 	}

// 	for _, inputService := range input.Services {
// 		existingServicesMap[inputService.ID] = inputService
// 	}

// 	var updatedServices []domain.Service
// 	for _, service := range existingServicesMap {
// 		updatedServices = append(updatedServices, service)
// 	}

// 	return updatedServices
// }
