package hotel

import (
	"context"

	domain "github.com/tapiaw38/globalstay-service-be/internal/domain/hotel"
	"github.com/tapiaw38/globalstay-service-be/internal/platform/appcontext"
)

type (
	UpdateUsecase interface {
		Execute(context.Context, string, domain.Hotel) (*UpdateOutput, error)
	}

	updateUsecase struct {
		contextFactory appcontext.Factory
	}

	UpdateOutput struct {
		Data HotelOutputData `json:"data"`
	}
)

func NewUpdateUsecase(contextFactory appcontext.Factory) UpdateUsecase {
	return &updateUsecase{
		contextFactory: contextFactory,
	}
}

func (u *updateUsecase) Execute(ctx context.Context, id string, input domain.Hotel) (*UpdateOutput, error) {
	app := u.contextFactory()

	roomsCurrent, err := app.Repositories.Hotel.FindRoomsByHotelID(ctx, id)
	if err != nil {
		return nil, err
	}

	updatedRooms := updateHotelRooms(roomsCurrent, input.Rooms)
	input.Rooms = updatedRooms

	if _, err := app.Repositories.Hotel.Update(ctx, id, input); err != nil {
		return nil, err
	}

	hotelUpdated, err := app.Repositories.Hotel.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return &UpdateOutput{
		Data: toHotelOutputData(*hotelUpdated),
	}, nil
}

func updateHotelRooms(existingRooms []domain.Room, inputRooms []domain.Room) []domain.Room {
	existingRoomsMap := make(map[string]domain.Room)

	for _, room := range existingRooms {
		existingRoomsMap[room.Number] = room
	}

	for _, inputRoom := range inputRooms {
		existingRoomsMap[inputRoom.Number] = inputRoom
	}

	var updatedRooms []domain.Room
	for _, room := range existingRoomsMap {
		updatedRooms = append(updatedRooms, room)
	}

	return updatedRooms
}
