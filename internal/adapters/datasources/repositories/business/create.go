package business

import (
	"context"

	domain "github.com/tapiaw38/reservation-service-be/internal/domain/business"
)

func (r *repository) Create(ctx context.Context, business domain.Business) (string, error) {
	businessDocument := marshal(business)
	result, err := r.client.InsertOne(ctx, businessDocument)
	if err != nil {
		return "", err
	}

	return result.InsertedID.(string), nil
}
