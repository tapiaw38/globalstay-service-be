package business

import (
	"context"

	domain "github.com/tapiaw38/reservation-service-be/internal/domain/business"
	"go.mongodb.org/mongo-driver/bson"
)

func (r *repository) FindServices(ctx context.Context, id string) ([]domain.Service, error) {
	var businessDocument BusinessDocument

	err := r.client.FindOne(ctx, bson.M{"_id": id}).Decode(&businessDocument)
	if err != nil {
		return nil, err
	}

	var services []domain.Service
	for _, service := range businessDocument.Services {
		services = append(services, unmarshalService(service))
	}

	return services, nil
}
