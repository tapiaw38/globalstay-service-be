package hotel

import (
	"context"

	domain "github.com/tapiaw38/globalstay-service-be/internal/domain/hotel"
	"go.mongodb.org/mongo-driver/bson"
)

func (r *repository) FindRoomsByHotelID(ctx context.Context, hotelID string) ([]domain.Room, error) {
	var hotelDocument HotelDocument
	err := r.client.FindOne(ctx, bson.M{"_id": hotelID}).Decode(&hotelDocument)
	if err != nil {
		return nil, err
	}

	hotels := unmarshal(hotelDocument)

	rooms := make([]domain.Room, len(hotels.Rooms))
	for i, room := range hotels.Rooms {
		rooms[i] = domain.Room{
			Number:        room.Number,
			Type:          room.Type,
			PricePerNight: room.PricePerNight,
			Pictures:      room.Pictures,
			Services:      room.Services,
		}
	}

	return rooms, nil
}
