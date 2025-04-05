package hotel

import (
	"context"

	domain "github.com/tapiaw38/globalstay-service-be/internal/domain/hotel"
	"github.com/tapiaw38/globalstay-service-be/internal/platform/appcontext"
)

type (
	UpdateRoomUsecase interface {
		Execute(context.Context, string, domain.Room) (*UpdateRoomOutput, error)
	}

	updateRoomUsecase struct {
		contextFactory appcontext.Factory
	}

	UpdateRoomOutput struct {
		Data RoomOutputData `json:"data"`
	}
)

func NewUpdateRoomUsecase(contextFactory appcontext.Factory) UpdateRoomUsecase {
	return &updateRoomUsecase{
		contextFactory: contextFactory,
	}
}

func (u *updateRoomUsecase) Execute(ctx context.Context, hotelID string, input domain.Room) (*UpdateRoomOutput, error) {
	app := u.contextFactory()

	roomUpdated, err := app.Repositories.Hotel.UpdateRoomByHotelID(ctx, hotelID, input)
	if err != nil {
		return nil, err
	}

	if roomUpdated == nil {
		return nil, nil
	}

	return &UpdateRoomOutput{
		Data: toRoomOutputData(*roomUpdated),
	}, nil
}
