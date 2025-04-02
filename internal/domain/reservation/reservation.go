package domain

import (
	hotel_domain "github.com/tapiaw38/globalstay-service-be/internal/domain/hotel"
	service_domain "github.com/tapiaw38/globalstay-service-be/internal/domain/service"
)

type (
	Reservation struct {
		ID       string
		Hotel    hotel_domain.Hotel
		Services []service_domain.Service
		Schedule []Schedule
	}

	Schedule struct {
		StartDate      string
		EndDate        string
		ReservationDay string
	}
)
