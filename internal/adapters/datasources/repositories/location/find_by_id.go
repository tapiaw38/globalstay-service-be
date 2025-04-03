package location

import (
	"context"

	domain "github.com/tapiaw38/globalstay-service-be/internal/domain/location"
	"go.mongodb.org/mongo-driver/bson"
)

func (r *repository) FindByID(ctx context.Context, id string) (*domain.Location, error) {
	var locationDocument LocationDocument
	err := r.client.FindOne(ctx, bson.M{"id": id}).Decode(&locationDocument)
	if err != nil {
		return nil, err
	}

	location := unmarshal(locationDocument)

	return &location, nil
}
