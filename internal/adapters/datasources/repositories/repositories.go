package repositories

import (
	"github.com/tapiaw38/globalstay-service-be/internal/adapters/datasources"
	"github.com/tapiaw38/globalstay-service-be/internal/adapters/datasources/repositories/hotel"
	"github.com/tapiaw38/globalstay-service-be/internal/adapters/datasources/repositories/service"
)

type Repositories struct {
	Hotel   hotel.Repository
	Service service.Repository
}

func CreateRepositories(datasources *datasources.Datasources) *Repositories {
	return &Repositories{
		Hotel:   hotel.NewRepository(datasources.NoSQLHotelClient),
		Service: service.NewRepository(datasources.NoSQLServicesClient),
	}
}
