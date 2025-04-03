package location

import (
	"context"

	"github.com/tapiaw38/globalstay-service-be/internal/platform/appcontext"
)

type (
	ListUsecase interface {
		Execute(ctx context.Context) ([]LocationOutputData, error)
	}

	listUsecase struct {
		contextFactory appcontext.Factory
	}

	ListOutput struct {
		Data []LocationOutputData `json:"data"`
	}
)

func NewListUsecase(contextFactory appcontext.Factory) ListUsecase {
	return &listUsecase{
		contextFactory: contextFactory,
	}
}

func (u *listUsecase) Execute(ctx context.Context) ([]LocationOutputData, error) {
	app := u.contextFactory()

	locations, err := app.Repositories.Location.FindAll(ctx)
	if err != nil {
		return nil, err
	}

	outputData := make([]LocationOutputData, 0, len(locations))
	for _, location := range locations {
		outputData = append(outputData, toLocationOutputData(location))
	}

	return outputData, nil
}
