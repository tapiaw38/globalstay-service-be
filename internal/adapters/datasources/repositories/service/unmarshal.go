package service

import domain "github.com/tapiaw38/reservation-service-be/internal/domain/service"

func unmarshal(serviceDocument ServiceDocument) domain.Service {
	return domain.Service{
		ID:          serviceDocument.ID,
		BusinessID:  serviceDocument.BusinessID,
		Name:        serviceDocument.Name,
		Description: serviceDocument.Description,
		Pictures:    serviceDocument.Pictures,
	}
}

func marshal(service domain.Service) ServiceDocument {
	return ServiceDocument{
		ID:          service.ID,
		BusinessID:  service.BusinessID,
		Name:        service.Name,
		Description: service.Description,
		Pictures:    service.Pictures,
	}
}
