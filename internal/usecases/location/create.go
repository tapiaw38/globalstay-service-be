package location

import (
	"context"

	"github.com/google/uuid"
	domain "github.com/tapiaw38/globalstay-service-be/internal/domain/location"
	"github.com/tapiaw38/globalstay-service-be/internal/platform/appcontext"
)

type (
	CreateUsecase interface {
		Execute(ctx context.Context, location domain.Location) (*CreateOutput, error)
	}

	createUsecase struct {
		contextFactory appcontext.Factory
	}

	CreateOutput struct {
		Data LocationOutputData `json:"data"`
	}
)

func NewCreateUsecase(contextFactory appcontext.Factory) CreateUsecase {
	return &createUsecase{
		contextFactory: contextFactory,
	}
}

func (u *createUsecase) Execute(ctx context.Context, input domain.Location) (*CreateOutput, error) {
	app := u.contextFactory()

	locationID, err := uuid.NewUUID()
	if err != nil {
		return nil, err
	}

	input.ID = locationID.String()
	id, err := app.Repositories.Location.Create(ctx, input)
	if err != nil {
		return nil, err
	}

	location, err := app.Repositories.Location.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return &CreateOutput{
		Data: toLocationOutputData(*location),
	}, nil
}
