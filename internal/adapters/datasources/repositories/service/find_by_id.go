package service

import (
	"context"

	domain "github.com/tapiaw38/globalstay-service-be/internal/domain/service"
	"go.mongodb.org/mongo-driver/bson"
)

func (r *repository) FindByID(ctx context.Context, id string) (*domain.Service, error) {
	var serviceDocument ServiceDocument
	err := r.client.GetCollection().FindOne(ctx, bson.M{"_id": id}).Decode(&serviceDocument)
	if err != nil {
		return nil, err
	}

	service := unmarshal(serviceDocument)

	return &service, nil
}
