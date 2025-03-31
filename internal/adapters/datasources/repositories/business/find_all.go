package business

import (
	"context"

	domain "github.com/tapiaw38/reservation-service-be/internal/domain/business"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type FilterOptions struct {
	Name        string
	Type        string
	Description string
	Active      bool
	Latitude    float64
	Longitude   float64
	Limit       int64
	Offset      int64
}

func (r *repository) FindAll(ctx context.Context, filter FilterOptions) ([]domain.Business, error) {
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

	if filter.Latitude != 0 && filter.Longitude != 0 {
		filters["location"] = bson.M{
			"$geoWithin": bson.M{
				"$centerSphere": []interface{}{
					[]float64{filter.Longitude, filter.Latitude},
					1000 / 3963.2,
				},
			},
		}
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

	var businesses []domain.Business
	for cursor.Next(ctx) {
		var business BusinessDocument
		if err := cursor.Decode(&business); err != nil {
			return nil, err
		}
		businesses = append(businesses, unmarshal(business))
	}

	return businesses, nil
}
