package hotel

import (
	"context"

	domain "github.com/tapiaw38/globalstay-service-be/internal/domain/hotel"
	"go.mongodb.org/mongo-driver/bson"
)

func (r *repository) FindByID(ctx context.Context, id string) (*domain.Hotel, error) {
	var hotelDocument HotelDocument
	err := r.client.FindOne(ctx, bson.M{"_id": id}).Decode(&hotelDocument)
	if err != nil {
		return nil, err
	}

	hotel := unmarshal(hotelDocument)

	return &hotel, nil
}
