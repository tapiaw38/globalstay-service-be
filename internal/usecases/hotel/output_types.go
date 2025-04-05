package hotel

import domain "github.com/tapiaw38/globalstay-service-be/internal/domain/hotel"

type (
	HotelOutputData struct {
		ID           string             `json:"id"`
		UserID       string             `json:"user_id"`
		Type         string             `json:"type"`
		Name         string             `json:"name"`
		Description  string             `json:"description"`
		PhoneNumber  string             `json:"phone_number"`
		Email        string             `json:"email"`
		Content      string             `json:"content"`
		Address      string             `json:"address"`
		Active       bool               `json:"active"`
		Latitude     float64            `json:"latitude"`
		Longitude    float64            `json:"longitude"`
		Pictures     []string           `json:"pictures"`
		Rooms        []RoomOutputData   `json:"rooms"`
		Reviews      []ReviewOutputData `json:"reviews"`
		AveragePrice float64            `json:"average_price"`
		LocationName string             `json:"location_name"`
	}

	RoomOutputData struct {
		Number        string              `json:"number"`
		Type          string              `json:"type"`
		PersonCount   int                 `json:"person_count"`
		PricePerNight float64             `json:"price_per_night"`
		Pictures      []string            `json:"pictures"`
		Services      []ServiceOutputData `json:"services"`
	}

	ServiceOutputData struct {
		Name        string   `json:"name"`
		Description string   `json:"description"`
		Icons       []string `json:"icons"`
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

	rooms := make([]RoomOutputData, len(hotel.Rooms))
	for i, room := range hotel.Rooms {
		rooms[i] = toRoomOutputData(room)
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
		Rooms:        rooms,
		Pictures:     pictures,
		Reviews:      reviews,
		AveragePrice: hotel.AveragePrice,
		LocationName: hotel.LocationName,
	}
}

func toRoomOutputData(room domain.Room) RoomOutputData {
	services := make([]ServiceOutputData, len(room.Services))
	for i, service := range room.Services {
		services[i] = toServiceOutputData(service)
		copy(services[i].Icons, service.Icons)
	}

	pictures := make([]string, len(room.Pictures))
	copy(pictures, room.Pictures)

	return RoomOutputData{
		Number:        room.Number,
		Type:          room.Type,
		PersonCount:   room.PersonCount,
		PricePerNight: room.PricePerNight,
		Pictures:      pictures,
		Services:      services,
	}
}

func toServiceOutputData(service domain.Service) ServiceOutputData {
	return ServiceOutputData{
		Name:        service.Name,
		Description: service.Description,
		Icons:       service.Icons,
	}
}

func toReviewOutputData(review domain.Review) ReviewOutputData {
	return ReviewOutputData{
		UserName: review.UserName,
		Rating:   review.Rating,
		Comment:  review.Comment,
	}
}
