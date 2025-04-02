package reservation

import (
	"context"

	domain "github.com/tapiaw38/globalstay-service-be/internal/domain/reservation"
)

func (r *repository) Create(ctx context.Context, reservation domain.Reservation) (string, error) {
	reservationDocument := marshal(reservation)
	result, err := r.client.GetCollection().InsertOne(ctx, reservationDocument)
	if err != nil {
		return "", err
	}

	return result.InsertedID.(string), nil
}
