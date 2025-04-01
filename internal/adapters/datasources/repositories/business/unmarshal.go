package business

import (
	domain_business "github.com/tapiaw38/reservation-service-be/internal/domain/business"
)

func marshal(business domain_business.Business) BusinessDocument {
	pictures := make([]string, len(business.Pictures))
	copy(pictures, business.Pictures)

	services := make([]ServiceDocument, len(business.Services))
	for i, service := range business.Services {
		services[i] = ServiceDocument{
			ID:          service.ID,
			Name:        service.Name,
			Description: service.Description,
			Pictures:    service.Pictures,
		}
	}

	reviews := make([]ReviewDocument, len(business.Reviews))
	for i, review := range business.Reviews {
		reviews[i] = ReviewDocument{
			UserName: review.UserName,
			Rating:   review.Rating,
			Comment:  review.Comment,
		}
	}

	return BusinessDocument{
		ID:           business.ID,
		UserID:       business.UserID,
		Type:         business.Type,
		Name:         business.Name,
		Description:  business.Description,
		PhoneNumber:  business.PhoneNumber,
		Email:        business.Email,
		Content:      business.Content,
		Address:      business.Address,
		Active:       business.Active,
		Latitude:     business.Latitude,
		Longitude:    business.Longitude,
		Pictures:     pictures,
		Services:     services,
		Reviews:      reviews,
		AveragePrice: business.AveragePrice,
	}
}

func unmarshal(businessDocument BusinessDocument) domain_business.Business {
	pictures := make([]string, len(businessDocument.Pictures))
	copy(pictures, businessDocument.Pictures)

	services := make([]domain_business.Service, len(businessDocument.Services))
	for i, service := range businessDocument.Services {
		services[i] = domain_business.Service{
			ID:          service.ID,
			Name:        service.Name,
			Description: service.Description,
			Pictures:    service.Pictures,
		}
	}

	reviews := make([]domain_business.Review, len(businessDocument.Reviews))
	for i, review := range businessDocument.Reviews {
		reviews[i] = domain_business.Review{
			UserName: review.UserName,
			Rating:   review.Rating,
			Comment:  review.Comment,
		}
	}

	return domain_business.Business{
		ID:           businessDocument.ID,
		UserID:       businessDocument.UserID,
		Type:         businessDocument.Type,
		Name:         businessDocument.Name,
		Description:  businessDocument.Description,
		PhoneNumber:  businessDocument.PhoneNumber,
		Email:        businessDocument.Email,
		Content:      businessDocument.Content,
		Address:      businessDocument.Address,
		Active:       businessDocument.Active,
		Latitude:     businessDocument.Latitude,
		Longitude:    businessDocument.Longitude,
		Pictures:     pictures,
		Services:     services,
		Reviews:      reviews,
		AveragePrice: businessDocument.AveragePrice,
	}
}

func unmarshalService(serviceDocument ServiceDocument) domain_business.Service {
	return domain_business.Service{
		ID:          serviceDocument.ID,
		Name:        serviceDocument.Name,
		Description: serviceDocument.Description,
		Pictures:    serviceDocument.Pictures,
	}
}
