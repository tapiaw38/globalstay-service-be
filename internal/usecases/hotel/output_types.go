package hotel

import domain "github.com/tapiaw38/globalstay-service-be/internal/domain/hotel"

type (
	HotelOutputData struct {
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

func toHotelOutputData(hotel domain.Hotel) HotelOutputData {
	reviews := make([]ReviewOutputData, len(hotel.Reviews))
	for i, review := range hotel.Reviews {
		reviews[i] = toReviewOutputData(review)
	}

	pictures := make([]string, len(hotel.Pictures))
	copy(pictures, hotel.Pictures)

	return HotelOutputData{
		ID:           hotel.ID,
		UserID:       hotel.UserID,
		Type:         hotel.Type,
		Name:         hotel.Name,
		Description:  hotel.Description,
		PhoneNumber:  hotel.PhoneNumber,
		Email:        hotel.Email,
		Content:      hotel.Content,
		Address:      hotel.Address,
		Active:       hotel.Active,
		Latitude:     hotel.Latitude,
		Longitude:    hotel.Longitude,
		Pictures:     pictures,
		Reviews:      reviews,
		AveragePrice: hotel.AveragePrice,
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
