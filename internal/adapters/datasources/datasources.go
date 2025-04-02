package datasources

import "github.com/tapiaw38/globalstay-service-be/internal/platform/nosql"

type Datasources struct {
	NoSQLHotelClient        nosql.Client
	NoSQLServicesClient     nosql.Client
	NoSQLReservationsClient nosql.Client
}

func CreateDatasources(
	noSQLHotelClient nosql.Client,
	noSQLServicesClient nosql.Client,
	noSQLReservationsClient nosql.Client,
) *Datasources {
	return &Datasources{
		NoSQLHotelClient:        noSQLHotelClient,
		NoSQLServicesClient:     noSQLServicesClient,
		NoSQLReservationsClient: noSQLReservationsClient,
	}
}
