package hotel

import (
	"context"

	domain "github.com/tapiaw38/globalstay-service-be/internal/domain/hotel"
)

func (r *repository) Create(ctx context.Context, hotel domain.Hotel) (string, error) {
	hotelDocument := marshal(hotel)
	result, err := r.client.InsertOne(ctx, hotelDocument)
	if err != nil {
		return "", err
	}

	return result.InsertedID.(string), nil
}
