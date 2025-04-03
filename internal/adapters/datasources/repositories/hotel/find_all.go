package hotel

import (
	"context"

	domain "github.com/tapiaw38/globalstay-service-be/internal/domain/hotel"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type FilterOptions struct {
	Name         string
	Type         string
	Description  string
	Active       bool
	Latitude     float64
	Longitude    float64
	LocationName string
	Limit        int64
	Offset       int64
}

func (r *repository) FindAll(ctx context.Context, filter FilterOptions) ([]domain.Hotel, error) {
	filters := bson.M{}
	if filter.Name != "" {
		filters["name"] = bson.M{"$regex": filter.Name}
	}

	if filter.Type != "" {
		filters["type"] = filter.Type
	}

	if filter.Description != "" {
		filters["description"] = bson.M{"$regex": filter.Description}
	}

	if filter.Active {
		filters["active"] = filter.Active
	}

	if filter.LocationName != "" {
		filters["location_name"] = filter.LocationName
	}

	opts := options.Find()
	if filter.Limit > 0 {
		opts.SetLimit(filter.Limit)
	}

	if filter.Offset > 0 {
		opts.SetSkip(filter.Offset)
	}

	cursor, err := r.client.Find(ctx, filters, opts)
	if err != nil {
		return nil, err
	}
	defer func(cursor *mongo.Cursor, ctx context.Context) {
		err := cursor.Close(ctx)
		if err != nil {
			panic(err)
		}
	}(cursor, ctx)

	var hotels []domain.Hotel
	for cursor.Next(ctx) {
		var hotel HotelDocument
		if err := cursor.Decode(&hotel); err != nil {
			return nil, err
		}
		hotels = append(hotels, unmarshal(hotel))
	}

	return hotels, nil
}
