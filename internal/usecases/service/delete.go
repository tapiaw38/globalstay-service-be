package service

import (
	"context"
	"github.com/tapiaw38/reservation-service-be/internal/platform/appcontext"
)

type (
	DeleteUsecase interface {
		Delete(context.Context, string) (*DeleteOutput, error)
	}

	deleteUsecase struct {
		contextFactory appcontext.Factory
	}

	DeleteOutput struct {
		Data ServiceDeleteOutputData `json:"data"`
	}
)

func NewDeleteUsecase(contextFactory appcontext.Factory) DeleteUsecase {
	return &deleteUsecase{
		contextFactory: contextFactory,
	}
}

func (u *deleteUsecase) Delete(ctx context.Context, id string) (*DeleteOutput, error) {
	app := u.contextFactory()

	id, err := app.Repositories.Service.Delete(ctx, id)
	if err != nil {
		return nil, err
	}

	data := ServiceDeleteOutputData{
		ID: id,
	}

	return &DeleteOutput{
		Data: data,
	}, nil
}
