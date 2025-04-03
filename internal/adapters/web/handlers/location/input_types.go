package location

import domain "github.com/tapiaw38/globalstay-service-be/internal/domain/location"

type (
	LocationInput struct {
		Name      string  `json:"name"`
		City      string  `json:"city"`
		State     string  `json:"state"`
		Country   string  `json:"country"`
		Latitude  float64 `json:"latitude"`
		Longitude float64 `json:"longitude"`
		Radius    uint    `json:"radius"`
	}
)

func toLocationInput(input LocationInput) domain.Location {
	return domain.Location{
		Name:      input.Name,
		City:      input.City,
		State:     input.State,
		Country:   input.Country,
		Latitude:  input.Latitude,
		Longitude: input.Longitude,
		Radius:    input.Radius,
	}
}
