package service

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
		Data ServiceOutputData `json:"data"`
	}
)

func NewGetUsecase(contextFactory appcontext.Factory) GetUsecase {
	return &getUsecase{
		contextFactory: contextFactory,
	}
}

func (u *getUsecase) Get(ctx context.Context, id string) (*GetOutput, error) {
	app := u.contextFactory()

	service, err := app.Repositories.Service.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return &GetOutput{
		Data: toServiceOutputData(*service),
	}, nil
}
