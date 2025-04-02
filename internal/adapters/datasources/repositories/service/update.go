package service

import (
	"context"
	"fmt"

	domain "github.com/tapiaw38/globalstay-service-be/internal/domain/service"
	"go.mongodb.org/mongo-driver/bson"
)

func (r *repository) Update(ctx context.Context, id string, service domain.Service) (string, error) {
	serviceDocument := marshal(service)
	result, err := r.client.GetCollection().UpdateOne(
		ctx,
		bson.M{"_id": id},
		bson.M{"$set": serviceDocument},
	)
	if err != nil {
		return "", err
	}

	if result.MatchedCount == 0 {
		return "", fmt.Errorf("no se encontró el documento con ID: %s", id)
	}

	return id, nil
}
