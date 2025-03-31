package domain

import (
	business_domain "github.com/tapiaw38/reservation-service-be/internal/domain/business"
	service_domain "github.com/tapiaw38/reservation-service-be/internal/domain/service"
)

type (
	Reservation struct {
		ID       string
		Business business_domain.Business
		Services []service_domain.Service
		Schedule []Schedule
	}

	Schedule struct {
		StartDate      string
		EndDate        string
		ReservationDay string
	}
)
