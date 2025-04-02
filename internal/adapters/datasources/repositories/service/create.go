package service

import (
	"context"

	domain "github.com/tapiaw38/globalstay-service-be/internal/domain/service"
)

func (r *repository) Create(ctx context.Context, service domain.Service) (string, error) {
	serviceDocument := marshal(service)
	result, err := r.client.InsertOne(ctx, serviceDocument)
	if err != nil {
		return "", err
	}

	return result.InsertedID.(string), nil
}
