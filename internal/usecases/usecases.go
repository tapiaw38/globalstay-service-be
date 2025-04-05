package usecases

import (
	"github.com/tapiaw38/globalstay-service-be/internal/platform/appcontext"
	"github.com/tapiaw38/globalstay-service-be/internal/usecases/hotel"
	"github.com/tapiaw38/globalstay-service-be/internal/usecases/location"
	"github.com/tapiaw38/globalstay-service-be/internal/usecases/service"
)

type Usecases struct {
	Hotel    Hotel
	Location Location
	Service  Service
}

type Hotel struct {
	CreateUsecase     hotel.CreateUsecase
	GetUsecase        hotel.GetUsecase
	UpdateUsecase     hotel.UpdateUsecase
	ListUsecase       hotel.ListUsecase
	ScanPlacesUsecase hotel.ScanPlacesUsecase
	UpdateRoomUsecase hotel.UpdateRoomUsecase
}

type Location struct {
	CreateUsecase location.CreateUsecase
	UpdateUsecase location.UpdateUsecase
	ListUsecase   location.ListUsecase
}

type Service struct {
	ListUsecase   service.ListUsecase
	GetUsecase    service.GetUsecase
	CreateUsecase service.CreateUsecase
	UpdateUsecase service.UpdateUsecase
	DeleteUsecase service.DeleteUsecase
}

func CreateUsecases(contextFactory appcontext.Factory) *Usecases {
	return &Usecases{
		Hotel: Hotel{
			CreateUsecase:     hotel.NewCreateUsecase(contextFactory),
			GetUsecase:        hotel.NewGetUsecase(contextFactory),
			UpdateUsecase:     hotel.NewUpdateUsecase(contextFactory),
			ListUsecase:       hotel.NewListUsecase(contextFactory),
			ScanPlacesUsecase: hotel.NewExecuteScanPlacesUsecase(contextFactory),
			UpdateRoomUsecase: hotel.NewUpdateRoomUsecase(contextFactory),
		},
		Location: Location{
			CreateUsecase: location.NewCreateUsecase(contextFactory),
			UpdateUsecase: location.NewUpdateUsecase(contextFactory),
			ListUsecase:   location.NewListUsecase(contextFactory),
		},
		Service: Service{
			ListUsecase:   service.NewListUsecase(contextFactory),
			GetUsecase:    service.NewGetUsecase(contextFactory),
			CreateUsecase: service.NewCreateUsecase(contextFactory),
			UpdateUsecase: service.NewUpdateUsecase(contextFactory),
			DeleteUsecase: service.NewDeleteUsecase(contextFactory),
		},
	}
}
