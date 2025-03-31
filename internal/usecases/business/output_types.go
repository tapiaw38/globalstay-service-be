package business

import domain "github.com/tapiaw38/reservation-service-be/internal/domain/business"

type (
	BusinessOutputData struct {
		ID           string              `json:"id"`
		UserID       string              `json:"user_id"`
		Type         string              `json:"type"`
		Name         string              `json:"name"`
		Description  string              `json:"description"`
		PhoneNumber  string              `json:"phone_number"`
		Email        string              `json:"email"`
		Content      string              `json:"content"`
		Address      string              `json:"address"`
		Active       bool                `json:"active"`
		Latitude     float64             `json:"latitude"`
		Longitude    float64             `json:"longitude"`
		Pictures     []string            `json:"pictures"`
		Services     []ServiceOutputData `json:"services"`
		Reviews      []ReviewOutputData  `json:"reviews"`
		AveragePrice float64             `json:"average_price"`
	}

	ServiceOutputData struct {
		ID          string   `json:"id"`
		Name        string   `json:"name"`
		Description string   `json:"description"`
		Pictures    []string `json:"pictures"`
	}

	ReviewOutputData struct {
		UserName string  `json:"user_name"`
		Rating   float64 `json:"rating"`
		Comment  string  `json:"comment"`
	}
)

func toBusinessOutputData(business domain.Business) BusinessOutputData {
	services := make([]ServiceOutputData, len(business.Services))
	for i, service := range business.Services {
		services[i] = toServiceOutputData(service)
	}

	reviews := make([]ReviewOutputData, len(business.Reviews))
	for i, review := range business.Reviews {
		reviews[i] = toReviewOutputData(review)
	}

	pictures := make([]string, len(business.Pictures))
	copy(pictures, business.Pictures)

	return BusinessOutputData{
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

func toServiceOutputData(service domain.Service) ServiceOutputData {
	return ServiceOutputData{
		ID:          service.ID,
		Name:        service.Name,
		Description: service.Description,
		Pictures:    service.Pictures,
	}
}

func toReviewOutputData(review domain.Review) ReviewOutputData {
	return ReviewOutputData{
		UserName: review.UserName,
		Rating:   review.Rating,
		Comment:  review.Comment,
	}
}
