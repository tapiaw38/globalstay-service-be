package reservation

import (
	"context"

	domain "github.com/tapiaw38/globalstay-service-be/internal/domain/reservation"
	"go.mongodb.org/mongo-driver/bson"
)

func (r *repository) List(ctx context.Context) ([]domain.Reservation, error) {
	cursor, err := r.client.GetCollection().Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}

	var reservations []domain.Reservation
	if err = cursor.All(ctx, &reservations); err != nil {
		return nil, err
	}

	return reservations, nil
}
