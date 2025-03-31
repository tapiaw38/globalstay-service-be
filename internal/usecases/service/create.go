package service

import (
	"context"
	"github.com/google/uuid"
	domain "github.com/tapiaw38/reservation-service-be/internal/domain/service"
	"github.com/tapiaw38/reservation-service-be/internal/platform/appcontext"
)

type (
	CreateUsecase interface {
		Create(context.Context, domain.Service) (*CreateOutput, error)
	}

	createUsecase struct {
		contextFactory appcontext.Factory
	}

	CreateOutput struct {
		Data ServiceOutputData `json:"data"`
	}
)

func NewCreateUsecase(contextFactory appcontext.Factory) CreateUsecase {
	return &createUsecase{
		contextFactory: contextFactory,
	}
}

func (u *createUsecase) Create(ctx context.Context, input domain.Service) (*CreateOutput, error) {
	app := u.contextFactory()

	serviceID, err := uuid.NewUUID()
	if err != nil {
		return nil, err
	}

	input.ID = serviceID.String()
	id, err := app.Repositories.Service.Create(ctx, input)
	if err != nil {
		return nil, err
	}

	service, err := app.Repositories.Service.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return &CreateOutput{
		Data: toServiceOutputData(*service),
	}, nil
}
