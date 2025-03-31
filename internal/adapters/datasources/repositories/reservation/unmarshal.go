package reservation

import (
	domain_business "github.com/tapiaw38/reservation-service-be/internal/domain/business"
	domain_reservation "github.com/tapiaw38/reservation-service-be/internal/domain/reservation"
	domain_service "github.com/tapiaw38/reservation-service-be/internal/domain/service"
)

func unmarshal(reservationDocument ReservationDocument) domain_reservation.Reservation {
	services := make([]domain_service.Service, len(reservationDocument.Services))
	for i, service := range reservationDocument.Services {
		services[i] = unmarshalService(service)
	}

	return domain_reservation.Reservation{
		ID:       reservationDocument.ID,
		Business: unmarshalBusiness(reservationDocument.Business),
		Services: services,
		Schedule: unmarshalSchedule(reservationDocument.Schedule),
	}
}

func unmarshalService(service ServiceDocument) domain_service.Service {
	return domain_service.Service{
		ID:          service.ID,
		BusinessID:  service.BusinessID,
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

func unmarshalBusiness(business BusinessDocument) domain_business.Business {
	return domain_business.Business{
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

func marshal(reservation domain_reservation.Reservation) ReservationDocument {
	services := make([]ServiceDocument, len(reservation.Services))
	for i, service := range reservation.Services {
		services[i] = marshalService(service)
	}

	return ReservationDocument{
		ID:       reservation.ID,
		Business: marshalBusiness(reservation.Business),
		Services: services,
		Schedule: marshalSchedule(reservation.Schedule),
	}
}

func marshalService(service domain_service.Service) ServiceDocument {
	return ServiceDocument{
		ID:          service.ID,
		BusinessID:  service.BusinessID,
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

func marshalBusiness(business domain_business.Business) BusinessDocument {
	return BusinessDocument{
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
