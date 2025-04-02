package hotel

import (
	"context"

	"github.com/tapiaw38/globalstay-service-be/internal/platform/appcontext"
)

type (
	GetUsecase interface {
		Get(context.Context, string) (*GetOutput, error)
	}

	getUsecase struct {
		contextFactory appcontext.Factory
	}

	GetOutput struct {
		Data HotelOutputData `json:"data"`
	}
)

func NewGetUsecase(contextFactory appcontext.Factory) GetUsecase {
	return &getUsecase{
		contextFactory: contextFactory,
	}
}

func (u *getUsecase) Get(ctx context.Context, id string) (*GetOutput, error) {
	app := u.contextFactory()

	hotel, err := app.Repositories.Hotel.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return &GetOutput{
		Data: toHotelOutputData(*hotel),
	}, nil
}
