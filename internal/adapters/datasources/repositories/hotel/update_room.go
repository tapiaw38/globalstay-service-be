package hotel

import (
	"context"

	domain "github.com/tapiaw38/globalstay-service-be/internal/domain/hotel"
	"go.mongodb.org/mongo-driver/bson"
)

func (r *repository) UpdateRoomByHotelID(ctx context.Context, hotelID string, room domain.Room) (*domain.Room, error) {
	filter := bson.M{
		"_id":          hotelID,
		"rooms.number": room.Number,
	}

	update := bson.M{
		"$set": bson.M{
			"rooms.$.type":          room.Type,
			"rooms.$.personCount":   room.PersonCount,
			"rooms.$.pricePerNight": room.PricePerNight,
			"rooms.$.pictures":      room.Pictures,
			"rooms.$.services":      room.Services,
		},
	}

	result, err := r.client.UpdateOne(ctx, filter, update)
	if err != nil {
		return nil, err
	}

	if result.MatchedCount == 0 {
		return nil, nil
	}

	updatedRoom := &domain.Room{
		Number:        room.Number,
		Type:          room.Type,
		PersonCount:   room.PersonCount,
		PricePerNight: room.PricePerNight,
		Pictures:      room.Pictures,
		Services:      room.Services,
	}

	return updatedRoom, nil
}
