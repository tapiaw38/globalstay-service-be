package location

import (
	"context"

	domain "github.com/tapiaw38/globalstay-service-be/internal/domain/location"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func (r *repository) FindAll(ctx context.Context) ([]domain.Location, error) {
	cursor, err := r.client.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer func(cursor *mongo.Cursor, ctx context.Context) {
		err := cursor.Close(ctx)
		if err != nil {
			panic(err)
		}
	}(cursor, ctx)

	var locations []domain.Location
	for cursor.Next(ctx) {
		var location LocationDocument
		if err := cursor.Decode(&location); err != nil {
			return nil, err
		}
		locations = append(locations, unmarshal(location))
	}

	return locations, nil
}
