package reservation

import (
	domain_hotel "github.com/tapiaw38/globalstay-service-be/internal/domain/hotel"
	domain_reservation "github.com/tapiaw38/globalstay-service-be/internal/domain/reservation"
	domain_service "github.com/tapiaw38/globalstay-service-be/internal/domain/service"
)

func unmarshal(reservationDocument ReservationDocument) domain_reservation.Reservation {
	services := make([]domain_service.Service, len(reservationDocument.Services))
	for i, service := range reservationDocument.Services {
		services[i] = unmarshalService(service)
	}

	return domain_reservation.Reservation{
		ID:       reservationDocument.ID,
		Hotel:    unmarshalHotel(reservationDocument.Hotel),
		Services: services,
		Schedule: unmarshalSchedule(reservationDocument.Schedule),
	}
}

func unmarshalService(service ServiceDocument) domain_service.Service {
	return domain_service.Service{
		ID:          service.ID,
		Name:        service.Name,
		Description: service.Description,
		Pictures:    service.Pictures,
	}
}

func unmarshalSchedule(schedule []ScheduleDocument) []domain_reservation.Schedule {
	var schedules []domain_reservation.Schedule

	for _, s := range schedule {
		schedules = append(schedules, domain_reservation.Schedule{
			StartDate:      s.StartDate,
			EndDate:        s.EndDate,
			ReservationDay: s.ReservationDay,
		})
	}

	return schedules
}

func unmarshalHotel(hotel HotelDocument) domain_hotel.Hotel {
	return domain_hotel.Hotel{
		ID:          hotel.ID,
		UserID:      hotel.UserID,
		Type:        hotel.Type,
		Name:        hotel.Name,
		Description: hotel.Description,
		PhoneNumber: hotel.PhoneNumber,
		Email:       hotel.Email,
		Content:     hotel.Content,
		Address:     hotel.Address,
		Active:      hotel.Active,
		Latitude:    hotel.Latitude,
		Longitude:   hotel.Longitude,
	}
}

func marshal(reservation domain_reservation.Reservation) ReservationDocument {
	services := make([]ServiceDocument, len(reservation.Services))
	for i, service := range reservation.Services {
		services[i] = marshalService(service)
	}

	return ReservationDocument{
		ID:       reservation.ID,
		Hotel:    marshalHotel(reservation.Hotel),
		Services: services,
		Schedule: marshalSchedule(reservation.Schedule),
	}
}

func marshalService(service domain_service.Service) ServiceDocument {
	return ServiceDocument{
		ID:          service.ID,
		Name:        service.Name,
		Description: service.Description,
		Pictures:    service.Pictures,
	}
}

func marshalSchedule(schedule []domain_reservation.Schedule) []ScheduleDocument {
	var schedules []ScheduleDocument

	for _, s := range schedule {
		schedules = append(schedules, ScheduleDocument{
			StartDate:      s.StartDate,
			EndDate:        s.EndDate,
			ReservationDay: s.ReservationDay,
		})
	}

	return schedules
}

func marshalHotel(hotel domain_hotel.Hotel) HotelDocument {
	return HotelDocument{
		ID:          hotel.ID,
		UserID:      hotel.UserID,
		Type:        hotel.Type,
		Name:        hotel.Name,
		Description: hotel.Description,
		PhoneNumber: hotel.PhoneNumber,
		Email:       hotel.Email,
		Content:     hotel.Content,
		Address:     hotel.Address,
		Active:      hotel.Active,
		Latitude:    hotel.Latitude,
		Longitude:   hotel.Longitude,
	}
}
