package business

import (
	"context"

	"github.com/google/uuid"
	domain "github.com/tapiaw38/reservation-service-be/internal/domain/business"
	"github.com/tapiaw38/reservation-service-be/internal/platform/appcontext"
)

type (
	CreateUsecase interface {
		Create(context.Context, domain.Business) (*CreateOutput, error)
	}

	createUsecase struct {
		contextFactory appcontext.Factory
	}

	CreateOutput struct {
		Data BusinessOutputData `json:"data"`
	}
)

func NewCreateUsecase(contextFactory appcontext.Factory) CreateUsecase {
	return &createUsecase{
		contextFactory: contextFactory,
	}
}

func (u *createUsecase) Create(ctx context.Context, input domain.Business) (*CreateOutput, error) {
	app := u.contextFactory()

	businessID, err := uuid.NewUUID()
	if err != nil {
		return nil, err
	}

	input.ID = businessID.String()
	id, err := app.Repositories.Business.Create(ctx, input)
	if err != nil {
		return nil, err
	}

	business, err := app.Repositories.Business.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return &CreateOutput{
		Data: toBusinessOutputData(*business),
	}, nil
}
