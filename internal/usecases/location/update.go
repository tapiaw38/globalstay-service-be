package location

import (
	"context"

	domain "github.com/tapiaw38/globalstay-service-be/internal/domain/location"
	"github.com/tapiaw38/globalstay-service-be/internal/platform/appcontext"
)

type (
	UpdateUsecase interface {
		Execute(ctx context.Context, id string, location domain.Location) (*UpdateOutput, error)
	}

	updateUsecase struct {
		contextFactory appcontext.Factory
	}

	UpdateOutput struct {
		Data LocationOutputData `json:"data"`
	}
)

func NewUpdateUsecase(contextFactory appcontext.Factory) UpdateUsecase {
	return &updateUsecase{
		contextFactory: contextFactory,
	}
}

func (u *updateUsecase) Execute(ctx context.Context, id string, input domain.Location) (*UpdateOutput, error) {
	app := u.contextFactory()

	if _, err := app.Repositories.Location.Update(ctx, id, input); err != nil {
		return nil, err
	}

	location, err := app.Repositories.Location.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return &UpdateOutput{
		Data: toLocationOutputData(*location),
	}, nil
}
