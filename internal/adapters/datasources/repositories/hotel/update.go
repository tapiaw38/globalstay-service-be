package hotel

import (
	"context"
	"fmt"

	domain "github.com/tapiaw38/globalstay-service-be/internal/domain/hotel"
	"go.mongodb.org/mongo-driver/bson"
)

func (r *repository) Update(ctx context.Context, id string, hotel domain.Hotel) (string, error) {
	hotelDocument := marshal(hotel)
	result, err := r.client.UpdateOne(
		ctx,
		bson.M{"_id": id},
		bson.M{"$set": hotelDocument},
	)
	if err != nil {
		return "", err
	}

	if result.MatchedCount == 0 {
		return "", fmt.Errorf("no se encontró el documento con ID: %s", id)
	}

	return id, nil
}
