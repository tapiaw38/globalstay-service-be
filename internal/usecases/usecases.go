package usecases

import (
	"github.com/tapiaw38/reservation-service-be/internal/platform/appcontext"
	"github.com/tapiaw38/reservation-service-be/internal/usecases/business"
	"github.com/tapiaw38/reservation-service-be/internal/usecases/service"
)

type Usecases struct {
	Business Business
	Service  Service
}

type Business struct {
	CreateUsecase business.CreateUsecase
	ListUsecase   business.ListUsecase
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
		Business: Business{
			CreateUsecase: business.NewCreateUsecase(contextFactory),
			ListUsecase:   business.NewListUsecase(contextFactory),
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
