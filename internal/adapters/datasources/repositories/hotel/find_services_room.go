package hotel

import (
	"context"

	domain "github.com/tapiaw38/globalstay-service-be/internal/domain/hotel"
	"go.mongodb.org/mongo-driver/bson"
)

func (r *repository) FindServicesRoomByHotelID(ctx context.Context, hotelID string) ([]domain.Service, error) {
	var hotelDocument HotelDocument
	err := r.client.FindOne(ctx, bson.M{"_id": hotelID}).Decode(&hotelDocument)
	if err != nil {
		return nil, err
	}

	hotels := unmarshal(hotelDocument)

	var services []domain.Service

	for _, room := range hotels.Rooms {
		for _, service := range room.Services {
			services = append(services, domain.Service{
				ID:          service.ID,
				Name:        service.Name,
				Description: service.Description,
				Icons:       service.Icons,
			})
		}
	}

	return services, nil
}
