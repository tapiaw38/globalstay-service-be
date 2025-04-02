package reservation

import (
	"context"

	domain "github.com/tapiaw38/globalstay-service-be/internal/domain/reservation"
	"go.mongodb.org/mongo-driver/bson"
)

func (r *repository) Update(ctx context.Context, reservation domain.Reservation) (string, error) {
	reservationDocument := marshal(reservation)
	result, err := r.client.GetCollection().UpdateOne(ctx, bson.M{"_id": reservation.ID}, bson.M{"$set": reservationDocument})
	if err != nil {
		return "", err
	}

	return result.UpsertedID.(string), nil
}
