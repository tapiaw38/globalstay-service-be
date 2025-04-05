package hotel

import domain "github.com/tapiaw38/globalstay-service-be/internal/domain/hotel"

type (
	HotelInput struct {
		UserID       string          `json:"user_id,omitempty"`
		Type         string          `json:"type,omitempty"`
		Name         string          `json:"name,omitempty"`
		Description  string          `json:"description,omitempty"`
		PhoneNumber  string          `json:"phone_number,omitempty"`
		Email        string          `json:"email,omitempty"`
		Content      string          `json:"content,omitempty"`
		Address      string          `json:"address,omitempty"`
		Active       bool            `json:"active,omitempty"`
		Latitude     float64         `json:"latitude,omitempty"`
		Longitude    float64         `json:"longitude,omitempty"`
		Rooms        []RoomInput     `json:"rooms,omitempty"`
		Pictures     []string        `json:"pictures,omitempty"`
		Reviews      []domain.Review `json:"reviews,omitempty"`
		AveragePrice float64         `json:"average_price,omitempty"`
	}

	RoomInput struct {
		Number        string         `json:"number,omitempty"`
		Type          string         `json:"type,omitempty"`
		PersonCount   int            `json:"person_count,omitempty"`
		PricePerNight float64        `json:"price_per_night,omitempty"`
		Pictures      []string       `json:"pictures,omitempty"`
		Services      []ServiceInput `json:"services,omitempty"`
	}

	ServiceInput struct {
		Name        string   `json:"name,omitempty"`
		Description string   `json:"description,omitempty"`
		Icons       []string `json:"pictures,omitempty"`
	}
)

func toHotel(input HotelInput) domain.Hotel {
	var rooms []domain.Room
	for _, room := range input.Rooms {
		rooms = append(rooms, toRoom(room))
	}

	return domain.Hotel{
		UserID:       input.UserID,
		Type:         input.Type,
		Name:         input.Name,
		Description:  input.Description,
		PhoneNumber:  input.PhoneNumber,
		Email:        input.Email,
		Content:      input.Content,
		Address:      input.Address,
		Active:       input.Active,
		Latitude:     input.Latitude,
		Longitude:    input.Longitude,
		Rooms:        rooms,
		Pictures:     input.Pictures,
		Reviews:      input.Reviews,
		AveragePrice: input.AveragePrice,
	}
}

func toRoom(input RoomInput) domain.Room {
	var services []domain.Service
	for _, service := range input.Services {
		services = append(services, toService(service))
	}

	return domain.Room{
		Number:        input.Number,
		Type:          input.Type,
		PersonCount:   input.PersonCount,
		PricePerNight: input.PricePerNight,
		Pictures:      input.Pictures,
		Services:      services,
	}
}

func toService(input ServiceInput) domain.Service {
	return domain.Service{
		Name:        input.Name,
		Description: input.Description,
		Icons:       input.Icons,
	}
}
