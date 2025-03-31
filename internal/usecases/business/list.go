package business

import (
	"context"

	"github.com/tapiaw38/reservation-service-be/internal/adapters/datasources/repositories/business"
	"github.com/tapiaw38/reservation-service-be/internal/platform/appcontext"
)

type (
	ListUsecase interface {
		List(context.Context, FilterOptions) (*ListOutput, error)
	}

	listUsecase struct {
		contextFactory appcontext.Factory
	}

	FilterOptions business.FilterOptions

	ListOutput struct {
		Data []BusinessOutputData `json:"data"`
	}
)

func NewListUsecase(contextFactory appcontext.Factory) ListUsecase {
	return &listUsecase{
		contextFactory: contextFactory,
	}
}

func (u *listUsecase) List(ctx context.Context, filter FilterOptions) (*ListOutput, error) {
	app := u.contextFactory()

	businesses, err := app.Repositories.Business.FindAll(ctx, business.FilterOptions(filter))
	if err != nil {
		return nil, err
	}

	outputData := make([]BusinessOutputData, 0)
	for _, business := range businesses {
		outputData = append(outputData, toBusinessOutputData(business))
	}

	return &ListOutput{
		Data: outputData,
	}, nil
}
