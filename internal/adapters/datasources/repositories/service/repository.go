package service

import (
	"context"
	domain "github.com/tapiaw38/reservation-service-be/internal/domain/service"
	"github.com/tapiaw38/reservation-service-be/internal/platform/nosql"
)

type (
	Repository interface {
		Create(context.Context, domain.Service) (string, error)
		FindAll(context.Context, FilterOptions) ([]domain.Service, error)
		FindByID(context.Context, string) (*domain.Service, error)
		Update(context.Context, string, domain.Service) (string, error)
		Delete(context.Context, string) (string, error)
	}

	repository struct {
		client nosql.Client
	}

	FilterOptions struct {
		BusinessID string
		Name       string
		Limit      int64
		Offset     int64
	}
)

func NewRepository(client nosql.Client) Repository {
	return &repository{client: client}
}
