package location

import (
	"context"
	"fmt"

	domain "github.com/tapiaw38/globalstay-service-be/internal/domain/location"
	"go.mongodb.org/mongo-driver/bson"
)

func (r *repository) Update(ctx context.Context, id string, location domain.Location) (string, error) {
	locationDocument := marshal(location)
	result, err := r.client.UpdateOne(ctx, bson.M{"id": id}, bson.M{"$set": locationDocument})
	if err != nil {
		return "", err
	}

	if result.MatchedCount == 0 {
		return "", fmt.Errorf("no document found with id %s", id)
	}

	return id, nil
}
