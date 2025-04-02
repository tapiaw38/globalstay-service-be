package reservation

import (
	"context"

	domain "github.com/tapiaw38/globalstay-service-be/internal/domain/reservation"
	"github.com/tapiaw38/globalstay-service-be/internal/platform/nosql"
)

type (
	Repository interface {
		Create(context.Context, domain.Reservation) (string, error)
		Get(context.Context, string) (domain.Reservation, error)
		List(context.Context) ([]domain.Reservation, error)
		Update(context.Context, domain.Reservation) (string, error)
		Delete(context.Context, string) error
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
