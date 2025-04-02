package service

import (
	"context"

	"github.com/tapiaw38/globalstay-service-be/internal/adapters/datasources/repositories/service"
	"github.com/tapiaw38/globalstay-service-be/internal/platform/appcontext"
)

type (
	ListUsecase interface {
		List(context.Context, FilterOptions) (*ListOutput, error)
	}

	listUsecase struct {
		contextFactory appcontext.Factory
	}

	FilterOptions service.FilterOptions

	ListOutput struct {
		Data []ServiceOutputData `json:"data"`
	}
)

func NewListUsecase(contextFactory appcontext.Factory) ListUsecase {
	return &listUsecase{
		contextFactory: contextFactory,
	}
}

func (u *listUsecase) List(ctx context.Context, filter FilterOptions) (*ListOutput, error) {
	app := u.contextFactory()

	services, err := app.Repositories.Service.FindAll(ctx, service.FilterOptions(filter))
	if err != nil {
		return nil, err
	}

	outputData := make([]ServiceOutputData, 0)
	for _, service := range services {
		outputData = append(outputData, toServiceOutputData(service))
	}

	return &ListOutput{
		Data: outputData,
	}, nil
}
