package hotel

import (
	"context"

	domain "github.com/tapiaw38/globalstay-service-be/internal/domain/hotel"
	"github.com/tapiaw38/globalstay-service-be/internal/platform/nosql"
)

type (
	Repository interface {
		Create(context.Context, domain.Hotel) (string, error)
		FindByID(context.Context, string) (*domain.Hotel, error)
		FindAll(context.Context, FilterOptions) ([]domain.Hotel, error)
		Update(context.Context, string, domain.Hotel) (string, error)
		Delete(context.Context, string) (string, error)
		FindRoomsByHotelID(ctx context.Context, hotelID string) ([]domain.Room, error)
		FindServicesRoomByHotelID(ctx context.Context, hotelID string) ([]domain.Service, error)
		UpdateRoomByHotelID(ctx context.Context, hotelID string, room domain.Room) (*domain.Room, error)
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
