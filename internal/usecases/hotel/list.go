package hotel

import (
	"context"

	"github.com/tapiaw38/globalstay-service-be/internal/adapters/datasources/repositories/hotel"
	domain "github.com/tapiaw38/globalstay-service-be/internal/domain/location"
	"github.com/tapiaw38/globalstay-service-be/internal/platform/appcontext"
)

type (
	ListUsecase interface {
		Execute(context.Context, FilterOptions) (*ListOutput, error)
	}

	listUsecase struct {
		contextFactory appcontext.Factory
	}

	FilterOptions hotel.FilterOptions

	ListOutput struct {
		Data []HotelOutputData `json:"data"`
	}
)

func NewListUsecase(contextFactory appcontext.Factory) ListUsecase {
	return &listUsecase{
		contextFactory: contextFactory,
	}
}

func (u *listUsecase) Execute(ctx context.Context, filter FilterOptions) (*ListOutput, error) {
	app := u.contextFactory()

	locations, err := app.Repositories.Location.FindAll(ctx)
	if err != nil {
		return nil, err
	}

	if filter.Latitude != 0 && filter.Longitude != 0 {
		locationNearest := domain.FindNearestLocation(filter.Latitude, filter.Longitude, locations)
		filter.LocationName = locationNearest.Name
	}

	hoteles, err := app.Repositories.Hotel.FindAll(ctx, hotel.FilterOptions(filter))
	if err != nil {
		return nil, err
	}

	outputData := make([]HotelOutputData, 0)
	for _, hotel := range hoteles {
		outputData = append(outputData, toHotelOutputData(hotel))
	}

	return &ListOutput{
		Data: outputData,
	}, nil
}
