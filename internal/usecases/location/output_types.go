package location

import domain "github.com/tapiaw38/globalstay-service-be/internal/domain/location"

type (
	LocationOutputData struct {
		ID        string  `json:"id"`
		Name      string  `json:"name"`
		City      string  `json:"city"`
		State     string  `json:"state"`
		Country   string  `json:"country"`
		Latitude  float64 `json:"latitude"`
		Longitude float64 `json:"longitude"`
		Radius    uint    `json:"radius"`
	}
)

func toLocationOutputData(location domain.Location) LocationOutputData {
	return LocationOutputData{
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
