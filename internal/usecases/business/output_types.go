package business

import domain "github.com/tapiaw38/reservation-service-be/internal/domain/business"

type (
	BusinessOutputData struct {
		ID          string  `json:"id"`
		UserID      string  `json:"user_id"`
		Type        string  `json:"type"`
		Name        string  `json:"name"`
		Description string  `json:"description"`
		PhoneNumber string  `json:"phone_number"`
		Email       string  `json:"email"`
		Content     string  `json:"content"`
		Address     string  `json:"address"`
		Active      bool    `json:"active"`
		Latitude    float64 `json:"latitude"`
		Longitude   float64 `json:"longitude"`
	}
)

func toBusinessOutputData(business domain.Business) BusinessOutputData {
	return BusinessOutputData{
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
