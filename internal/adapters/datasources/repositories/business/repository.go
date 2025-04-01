package business

import (
	"context"

	domain "github.com/tapiaw38/reservation-service-be/internal/domain/business"
	"github.com/tapiaw38/reservation-service-be/internal/platform/nosql"
)

type (
	Repository interface {
		Create(context.Context, domain.Business) (string, error)
		FindByID(context.Context, string) (*domain.Business, error)
		FindAll(context.Context, FilterOptions) ([]domain.Business, error)
		Update(context.Context, string, domain.Business) (string, error)
		Delete(context.Context, string) (string, error)
		FindServices(context.Context, string) ([]domain.Service, error)
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
