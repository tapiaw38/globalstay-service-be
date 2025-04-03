package location

import (
	"context"

	domain "github.com/tapiaw38/globalstay-service-be/internal/domain/location"
)

func (r *repository) Create(ctx context.Context, location domain.Location) (string, error) {
	locationDocument := marshal(location)
	result, err := r.client.InsertOne(ctx, locationDocument)
	if err != nil {
		return "", err
	}

	return result.InsertedID.(string), nil
}
