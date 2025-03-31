package service

import (
	"context"
	domain "github.com/tapiaw38/reservation-service-be/internal/domain/service"
	"github.com/tapiaw38/reservation-service-be/internal/platform/appcontext"
)

type (
	UpdateUsecase interface {
		Update(context.Context, string, domain.Service) (*UpdateOutput, error)
	}

	updateUsecase struct {
		contextFactory appcontext.Factory
	}

	UpdateOutput struct {
		Data ServiceOutputData `json:"data"`
	}
)

func NewUpdateUsecase(contextFactory appcontext.Factory) UpdateUsecase {
	return &updateUsecase{
		contextFactory: contextFactory,
	}
}

func (u *updateUsecase) Update(ctx context.Context, id string, input domain.Service) (*UpdateOutput, error) {
	app := u.contextFactory()

	id, err := app.Repositories.Service.Update(ctx, id, input)
	if err != nil {
		return nil, err
	}

	service, err := app.Repositories.Service.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return &UpdateOutput{
		Data: toServiceOutputData(*service),
	}, nil
}
