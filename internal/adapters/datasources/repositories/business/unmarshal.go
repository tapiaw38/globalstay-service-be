package business

import (
	domain_business "github.com/tapiaw38/reservation-service-be/internal/domain/business"
)

func marshal(business domain_business.Business) BusinessDocument {
	return BusinessDocument{
		ID:          business.ID,
		UserID:      business.UserID,
		Type:        business.Type,
		Name:        business.Name,
		Description: business.Description,
		PhoneNumber: business.PhoneNumber,
		Email:       business.Email,
		Content:     business.Content,
		Address:     business.Address,
		Active:      business.Active,
		Latitude:    business.Latitude,
		Longitude:   business.Longitude,
	}
}

func unmarshal(businessDocument BusinessDocument) domain_business.Business {
	return domain_business.Business{
		ID:          businessDocument.ID,
		UserID:      businessDocument.UserID,
		Type:        businessDocument.Type,
		Name:        businessDocument.Name,
		Description: businessDocument.Description,
		PhoneNumber: businessDocument.PhoneNumber,
		Email:       businessDocument.Email,
		Content:     businessDocument.Content,
		Address:     businessDocument.Address,
		Active:      businessDocument.Active,
		Latitude:    businessDocument.Latitude,
		Longitude:   businessDocument.Longitude,
	}
}
