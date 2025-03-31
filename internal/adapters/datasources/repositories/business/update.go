package business

import (
	"context"
	"fmt"

	domain "github.com/tapiaw38/reservation-service-be/internal/domain/business"
	"go.mongodb.org/mongo-driver/bson"
)

func (r *repository) Update(ctx context.Context, id string, business domain.Business) (string, error) {
	businessDocument := marshal(business)
	result, err := r.client.GetCollection().UpdateOne(
		ctx,
		bson.M{"_id": id},
		bson.M{"$set": businessDocument},
	)
	if err != nil {
		return "", err
	}

	if result.MatchedCount == 0 {
		return "", fmt.Errorf("no se encontró el documento con ID: %s", id)
	}

	return id, nil
}
