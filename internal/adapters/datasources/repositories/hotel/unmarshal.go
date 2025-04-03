package hotel

import (
	domain_hotel "github.com/tapiaw38/globalstay-service-be/internal/domain/hotel"
)

func marshal(hotel domain_hotel.Hotel) HotelDocument {
	pictures := make([]string, len(hotel.Pictures))
	copy(pictures, hotel.Pictures)

	reviews := make([]ReviewDocument, len(hotel.Reviews))
	for i, review := range hotel.Reviews {
		reviews[i] = ReviewDocument{
			UserName: review.UserName,
			Rating:   review.Rating,
			Comment:  review.Comment,
		}
	}

	rooms := make([]RoomDocument, len(hotel.Rooms))
	for i, room := range hotel.Rooms {
		services := make([]ServiceDocument, len(room.Services))
		for j, service := range room.Services {
			services[j] = ServiceDocument{
				ID:          service.ID,
				Name:        service.Name,
				Description: service.Description,
				Pictures:    service.Pictures,
			}
		}

		rooms[i] = RoomDocument{
			ID:            room.ID,
			Number:        room.Number,
			Type:          room.Type,
			IsOccupied:    room.IsOccupied,
			PersonCount:   room.PersonCount,
			GuestName:     room.GuestName,
			PricePerNight: room.PricePerNight,
			Pictures:      room.Pictures,
			Services:      services,
		}
	}

	return HotelDocument{
		ID:           hotel.ID,
		PlaceID:      hotel.PlaceID,
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

func unmarshal(hotelDocument HotelDocument) domain_hotel.Hotel {
	pictures := make([]string, len(hotelDocument.Pictures))
	copy(pictures, hotelDocument.Pictures)

	reviews := make([]domain_hotel.Review, len(hotelDocument.Reviews))
	for i, review := range hotelDocument.Reviews {
		reviews[i] = domain_hotel.Review{
			UserName: review.UserName,
			Rating:   review.Rating,
			Comment:  review.Comment,
		}
	}

	rooms := make([]domain_hotel.Room, len(hotelDocument.Rooms))
	for i, room := range hotelDocument.Rooms {
		services := make([]domain_hotel.Service, len(room.Services))
		for j, service := range room.Services {
			services[j] = domain_hotel.Service{
				ID:          service.ID,
				Name:        service.Name,
				Description: service.Description,
				Pictures:    service.Pictures,
			}
		}

		rooms[i] = domain_hotel.Room{
			ID:            room.ID,
			Number:        room.Number,
			Type:          room.Type,
			IsOccupied:    room.IsOccupied,
			PersonCount:   room.PersonCount,
			GuestName:     room.GuestName,
			PricePerNight: room.PricePerNight,
			Pictures:      room.Pictures,
			Services:      services,
		}
	}

	return domain_hotel.Hotel{
		ID:           hotelDocument.ID,
		PlaceID:      hotelDocument.PlaceID,
		UserID:       hotelDocument.UserID,
		Type:         hotelDocument.Type,
		Name:         hotelDocument.Name,
		Description:  hotelDocument.Description,
		PhoneNumber:  hotelDocument.PhoneNumber,
		Email:        hotelDocument.Email,
		Content:      hotelDocument.Content,
		Address:      hotelDocument.Address,
		Active:       hotelDocument.Active,
		Latitude:     hotelDocument.Latitude,
		Longitude:    hotelDocument.Longitude,
		Rooms:        rooms,
		Pictures:     pictures,
		Reviews:      reviews,
		AveragePrice: hotelDocument.AveragePrice,
		LocationName: hotelDocument.LocationName,
	}
}
