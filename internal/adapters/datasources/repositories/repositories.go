package repositories

import (
	"github.com/tapiaw38/reservation-service-be/internal/adapters/datasources"
	"github.com/tapiaw38/reservation-service-be/internal/adapters/datasources/repositories/business"
	"github.com/tapiaw38/reservation-service-be/internal/adapters/datasources/repositories/service"
)

type Repositories struct {
	Business business.Repository
	Service  service.Repository
}

func CreateRepositories(datasources *datasources.Datasources) *Repositories {
	return &Repositories{
		Business: business.NewRepository(datasources.NoSQLBusinessClient),
		Service:  service.NewRepository(datasources.NoSQLServicesClient),
	}
}
