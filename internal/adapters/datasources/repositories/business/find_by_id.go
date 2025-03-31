package business

import (
	"context"

	domain "github.com/tapiaw38/reservation-service-be/internal/domain/business"
	"go.mongodb.org/mongo-driver/bson"
)

func (r *repository) FindByID(ctx context.Context, id string) (*domain.Business, error) {
	var businessDocument BusinessDocument

	err := r.client.FindOne(ctx, bson.M{"_id": id}).Decode(&businessDocument)
	if err != nil {
		return nil, err
	}

	business := unmarshal(businessDocument)

	return &business, nil
}
