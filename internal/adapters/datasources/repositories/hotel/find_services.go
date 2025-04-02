package hotel

import (
	"context"

	domain "github.com/tapiaw38/globalstay-service-be/internal/domain/hotel"
	"go.mongodb.org/mongo-driver/bson"
)

func (r *repository) FindServices(ctx context.Context, id string) ([]domain.Service, error) {
	var hotelDocument HotelDocument

	err := r.client.FindOne(ctx, bson.M{"_id": id}).Decode(&hotelDocument)
	if err != nil {
		return nil, err
	}

	var services []domain.Service
	for _, service := range hotelDocument.Services {
		services = append(services, unmarshalService(service))
	}

	return services, nil
}
