package reservation

import (
	"context"

	domain "github.com/tapiaw38/globalstay-service-be/internal/domain/reservation"
	"go.mongodb.org/mongo-driver/bson"
)

func (r *repository) Get(ctx context.Context, id string) (domain.Reservation, error) {
	var reservation domain.Reservation
	err := r.client.GetCollection().FindOne(ctx, bson.M{"_id": id}).Decode(&reservation)
	if err != nil {
		return domain.Reservation{}, err
	}

	return reservation, nil
}
