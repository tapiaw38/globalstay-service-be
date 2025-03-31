package business

import domain "github.com/tapiaw38/reservation-service-be/internal/domain/business"

type (
	BusinessInput struct {
		UserID      string  `json:"user_id,omitempty"`
		Type        string  `json:"type,omitempty"`
		Name        string  `json:"name,omitempty"`
		Description string  `json:"description,omitempty"`
		PhoneNumber string  `json:"phone_number,omitempty"`
		Email       string  `json:"email,omitempty"`
		Content     string  `json:"content,omitempty"`
		Address     string  `json:"address,omitempty"`
		Active      bool    `json:"active,omitempty"`
		Latitude    float64 `json:"latitude,omitempty"`
		Longitude   float64 `json:"longitude,omitempty"`
	}
)

func toBusiness(input BusinessInput) domain.Business {
	return domain.Business{
		UserID:      input.UserID,
		Type:        input.Type,
		Name:        input.Name,
		Description: input.Description,
		PhoneNumber: input.PhoneNumber,
		Email:       input.Email,
		Content:     input.Content,
		Address:     input.Address,
		Active:      input.Active,
		Latitude:    input.Latitude,
		Longitude:   input.Longitude,
	}
}
