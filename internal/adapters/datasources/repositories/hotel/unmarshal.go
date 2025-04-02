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

	return HotelDocument{
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

	return domain_hotel.Hotel{
		ID:           hotelDocument.ID,
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
		Pictures:     pictures,
		Reviews:      reviews,
		AveragePrice: hotelDocument.AveragePrice,
	}
}

func unmarshalService(serviceDocument ServiceDocument) domain_hotel.Service {
	return domain_hotel.Service{
		ID:          serviceDocument.ID,
		Name:        serviceDocument.Name,
		Description: serviceDocument.Description,
		Pictures:    serviceDocument.Pictures,
	}
}
