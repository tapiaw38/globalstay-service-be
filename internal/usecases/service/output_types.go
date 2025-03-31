package service

import domain "github.com/tapiaw38/reservation-service-be/internal/domain/service"

type (
	ServiceOutputData struct {
		ID          string   `json:"id"`
		Name        string   `json:"name"`
		Description string   `json:"description"`
		Pictures    []string `json:"pictures"`
	}

	ServiceDeleteOutputData struct {
		ID string `json:"id"`
	}
)

func toServiceOutputData(service domain.Service) ServiceOutputData {
	return ServiceOutputData{
		ID:          service.ID,
		Name:        service.Name,
		Description: service.Description,
		Pictures:    service.Pictures,
	}
}
