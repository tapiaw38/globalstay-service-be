package location

import domain "github.com/tapiaw38/globalstay-service-be/internal/domain/location"

func unmarshal(hotel LocationDocument) domain.Location {
	return domain.Location{
		ID:        hotel.ID,
		Name:      hotel.Name,
		City:      hotel.City,
		State:     hotel.State,
		Country:   hotel.Country,
		Latitude:  hotel.Latitude,
		Longitude: hotel.Longitude,
		Radius:    hotel.Radius,
	}
}

func marshal(location domain.Location) LocationDocument {
	return LocationDocument{
		ID:        location.ID,
		Name:      location.Name,
		City:      location.City,
		State:     location.State,
		Country:   location.Country,
		Latitude:  location.Latitude,
		Longitude: location.Longitude,
		Radius:    location.Radius,
	}
}
