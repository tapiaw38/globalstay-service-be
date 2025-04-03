package hotel

import (
	"context"

	"github.com/google/uuid"
	domain "github.com/tapiaw38/globalstay-service-be/internal/domain/hotel"
	"github.com/tapiaw38/globalstay-service-be/internal/platform/appcontext"
)

type (
	CreateUsecase interface {
		Execute(context.Context, domain.Hotel) (*CreateOutput, error)
	}

	createUsecase struct {
		contextFactory appcontext.Factory
	}

	CreateOutput struct {
		Data HotelOutputData `json:"data"`
	}
)

func NewCreateUsecase(contextFactory appcontext.Factory) CreateUsecase {
	return &createUsecase{
		contextFactory: contextFactory,
	}
}

func (u *createUsecase) Execute(ctx context.Context, input domain.Hotel) (*CreateOutput, error) {
	app := u.contextFactory()

	hotelID, err := uuid.NewUUID()
	if err != nil {
		return nil, err
	}

	input.ID = hotelID.String()
	id, err := app.Repositories.Hotel.Create(ctx, input)
	if err != nil {
		return nil, err
	}

	hotel, err := app.Repositories.Hotel.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return &CreateOutput{
		Data: toHotelOutputData(*hotel),
	}, nil
}
