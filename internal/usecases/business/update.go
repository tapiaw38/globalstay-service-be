package business

import (
	"context"

	domain "github.com/tapiaw38/reservation-service-be/internal/domain/business"
	"github.com/tapiaw38/reservation-service-be/internal/platform/appcontext"
)

type (
	UpdateUsecase interface {
		Update(context.Context, string, domain.Business) (*UpdateOutput, error)
	}

	updateUsecase struct {
		contextFactory appcontext.Factory
	}

	UpdateOutput struct {
		Data BusinessOutputData `json:"data"`
	}
)

func NewUpdateUsecase(contextFactory appcontext.Factory) UpdateUsecase {
	return &updateUsecase{
		contextFactory: contextFactory,
	}
}

func (u *updateUsecase) Update(ctx context.Context, id string, input domain.Business) (*UpdateOutput, error) {
	app := u.contextFactory()

	_, err := app.Repositories.Business.Update(ctx, id, input)
	if err != nil {
		return nil, err
	}

	business, err := app.Repositories.Business.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return &UpdateOutput{
		Data: toBusinessOutputData(*business),
	}, nil
}
