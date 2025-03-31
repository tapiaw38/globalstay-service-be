package service

import domain "github.com/tapiaw38/reservation-service-be/internal/domain/service"

type serviceInput struct {
	BusinessID  string   `json:"business_id,omitempty"`
	Name        string   `json:"name,omitempty"`
	Description string   `json:"description,omitempty"`
	Pictures    []string `json:"pictures,omitempty"`
}

func toService(input serviceInput) domain.Service {
	return domain.Service{
		BusinessID:  input.BusinessID,
		Name:        input.Name,
		Description: input.Description,
		Pictures:    input.Pictures,
	}
}
