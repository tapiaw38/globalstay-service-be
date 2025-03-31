package datasources

import "github.com/tapiaw38/reservation-service-be/internal/platform/nosql"

type Datasources struct {
	NoSQLBusinessClient     nosql.Client
	NoSQLServicesClient     nosql.Client
	NoSQLReservationsClient nosql.Client
}

func CreateDatasources(
	noSQLBusinessClient nosql.Client,
	noSQLServicesClient nosql.Client,
	noSQLReservationsClient nosql.Client,
) *Datasources {
	return &Datasources{
		NoSQLBusinessClient:     noSQLBusinessClient,
		NoSQLServicesClient:     noSQLServicesClient,
		NoSQLReservationsClient: noSQLReservationsClient,
	}
}
