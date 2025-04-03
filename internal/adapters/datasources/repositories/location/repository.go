package location

import (
	"context"

	domain "github.com/tapiaw38/globalstay-service-be/internal/domain/location"
	"github.com/tapiaw38/globalstay-service-be/internal/platform/nosql"
)

type (
	Repository interface {
		FindAll(context.Context) ([]domain.Location, error)
		FindByID(context.Context, string) (*domain.Location, error)
		Create(context.Context, domain.Location) (string, error)
		Update(context.Context, string, domain.Location) (string, error)
	}

	repository struct {
		client nosql.Client
	}
)

func NewRepository(client nosql.Client) Repository {
	return &repository{
		client: client,
	}
}
