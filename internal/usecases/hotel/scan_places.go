package hotel

import (
	"context"
	"log"

	"github.com/google/uuid"
	"github.com/tapiaw38/globalstay-service-be/internal/adapters/datasources/repositories/hotel"
	domain "github.com/tapiaw38/globalstay-service-be/internal/domain/hotel"
	"github.com/tapiaw38/globalstay-service-be/internal/platform/appcontext"
	"googlemaps.github.io/maps"
)

type (
	ScanPlacesUsecase interface {
		Execute(context.Context, string, uint, string) error
	}

	scanPlacesUsecase struct {
		contextFactory appcontext.Factory
	}
)

func NewExecuteScanPlacesUsecase(contextFactory appcontext.Factory) ScanPlacesUsecase {
	return &scanPlacesUsecase{
		contextFactory: contextFactory,
	}
}

func (u *scanPlacesUsecase) Execute(ctx context.Context, location string, radius uint, types string) error {
	app := u.contextFactory()

	places, err := app.Integrations.Places.GetPlaces(location, radius, maps.PlaceType(types))
	if err != nil {
		return err
	}

	hotels, err := app.Repositories.Hotel.FindAll(ctx, hotel.FilterOptions{})
	if err != nil {
		return err
	}

	existingHotelsMap := make(map[string]domain.Hotel)
	for _, hotel := range hotels {
		existingHotelsMap[hotel.PlaceID] = hotel
	}

	for _, place := range places {
		if _, ok := existingHotelsMap[place.ID]; ok {
			log.Printf("Hotel already exists: %s", place.ID)
			continue
		}

		id, err := uuid.NewUUID()
		if err != nil {
			log.Printf("Error generating uuid: %v", err)
			return err
		}

		hotel := domain.Hotel{
			ID:          id.String(),
			PlaceID:     place.ID,
			Name:        place.Name,
			Description: place.Description,
			Address:     place.Address,
			Pictures:    place.Photos,
			Latitude:    place.Latitude,
			Longitude:   place.Longitude,
		}

		if !hotel.HasCompletePlaceInformation() {
			log.Printf("Hotel does not have complete place information: %s", place.ID)
			continue
		}

		if _, err := app.Repositories.Hotel.Create(ctx, hotel); err != nil {
			log.Printf("Error creating hotel: %v", err)
			continue
		}
	}

	return nil
}
