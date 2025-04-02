package service

import (
	"context"

	domain "github.com/tapiaw38/globalstay-service-be/internal/domain/service"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (r *repository) FindAll(ctx context.Context, filter FilterOptions) ([]domain.Service, error) {
	filters := bson.M{}
	if filter.HotelID != "" {
		filters["hotel_id"] = filter.HotelID
	}

	if filter.Name != "" {
		filters["name"] = bson.M{"$regex": filter.Name, "$options": "i"}
	}

	opts := options.Find()
	if filter.Limit > 0 {
		opts.SetLimit(filter.Limit)
	}
	if filter.Offset > 0 {
		opts.SetSkip(filter.Offset)
	}

	cursor, err := r.client.GetCollection().Find(ctx, filters, opts)
	if err != nil {
		return nil, err
	}
	defer func(cursor *mongo.Cursor, ctx context.Context) {
		err := cursor.Close(ctx)
		if err != nil {
			panic(err)
		}
	}(cursor, ctx)

	var services []domain.Service
	for cursor.Next(ctx) {
		var service ServiceDocument
		if err := cursor.Decode(&service); err != nil {
			return nil, err
		}
		services = append(services, unmarshal(service))
	}

	return services, nil
}
