package business

import (
	"context"

	"github.com/tapiaw38/reservation-service-be/internal/platform/appcontext"
)

type (
	GetUsecase interface {
		Get(context.Context, string) (*GetOutput, error)
	}

	getUsecase struct {
		contextFactory appcontext.Factory
	}

	GetOutput struct {
		Data BusinessOutputData `json:"data"`
	}
)

func NewGetUsecase(contextFactory appcontext.Factory) GetUsecase {
	return &getUsecase{
		contextFactory: contextFactory,
	}
}

func (u *getUsecase) Get(ctx context.Context, id string) (*GetOutput, error) {
	app := u.contextFactory()

	business, err := app.Repositories.Business.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return &GetOutput{
		Data: toBusinessOutputData(*business),
	}, nil
}
