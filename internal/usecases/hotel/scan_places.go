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
		Execute(context.Context) error
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

func (u *scanPlacesUsecase) Execute(ctx context.Context) error {
	app := u.contextFactory()

	locations, err := app.Repositories.Location.FindAll(ctx)
	if err != nil {
		return err
	}

	for _, location := range locations {
		places, err := app.Integrations.Places.GetPlaces(
			location.Latitude,
			location.Longitude,
			location.Radius,
			maps.PlaceTypeLodging,
		)
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

				updatedHotel := existingHotelsMap[place.ID]

				if !updatedHotel.HasCompletePlaceInformation() {
					updatedHotel.Name = place.Name
					updatedHotel.Longitude = place.Longitude
					updatedHotel.Latitude = place.Latitude
					updatedHotel.Pictures = place.Photos
					updatedHotel.Address = place.FormattedAddress
					updatedHotel.LocationName = location.Name
				}

				if _, err := app.Repositories.Hotel.Update(ctx, updatedHotel.ID, updatedHotel); err != nil {
					log.Printf("Error updating hotel: %v", err)
					continue
				}

				log.Printf("Hotel updated: %s", place.ID)

				log.Printf("Hotel already exists: %s", place.ID)
				continue
			}

			id, err := uuid.NewUUID()
			if err != nil {
				log.Printf("Error generating uuid: %v", err)
				return err
			}

			hotel := domain.Hotel{
				ID:           id.String(),
				PlaceID:      place.ID,
				Name:         place.Name,
				Description:  place.Description,
				Address:      place.FormattedAddress,
				Pictures:     place.Photos,
				Latitude:     place.Latitude,
				Longitude:    place.Longitude,
				LocationName: location.Name,
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
	}

	return nil
}
