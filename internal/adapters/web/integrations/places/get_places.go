package places

import (
	"context"
	"os"

	"googlemaps.github.io/maps"
)

type (
	Place struct {
		ID               string
		Name             string
		Description      string
		Photos           []string
		Address          string
		FormattedAddress string
		Rating           float32
		UserRatingsTotal int
		Types            []string
		PriceLevel       int
		BusinessStatus   string
		Latitude         float64
		Longitude        float64
	}
)

func (p *Place) CompleteMinimumPlaceInformation() bool {
	return p.Name != "" &&
		p.Latitude != 0 &&
		p.Longitude != 0 &&
		len(p.Photos) > 0
}

func (i integration) GetPlaces(latitude, longitude float64, radius uint, types maps.PlaceType) ([]Place, error) {
	if radius == 0 {
		radius = googleMapsRadius
	}

	request := &maps.NearbySearchRequest{
		Location: &maps.LatLng{
			Lat: latitude,
			Lng: longitude,
		},
		Radius: radius,
		Type:   types,
	}

	response, err := i.client.NearbySearch(context.Background(), request)
	if err != nil {
		return nil, err
	}

	places := make([]Place, len(response.Results))

	for _, result := range response.Results {
		place := Place{
			ID:               result.PlaceID,
			Name:             result.Name,
			Photos:           extractPhoto(result.Photos),
			Address:          result.FormattedAddress,
			FormattedAddress: result.FormattedAddress,
			Rating:           result.Rating,
			UserRatingsTotal: result.UserRatingsTotal,
			Types:            result.Types,
			PriceLevel:       result.PriceLevel,
			BusinessStatus:   result.BusinessStatus,
			Latitude:         result.Geometry.Location.Lat,
			Longitude:        result.Geometry.Location.Lng,
		}

		places = append(places, place)
	}

	return places, nil
}

func extractPhoto(photos []maps.Photo) []string {
	var photoURLs []string
	apiKey := os.Getenv("GOOGLE_MAPS_KEY")

	for _, photo := range photos {
		photoURL := "https://maps.googleapis.com/maps/api/place/photo?maxwidth=800&photoreference=" + photo.PhotoReference + "&key=" + apiKey
		photoURLs = append(photoURLs, photoURL)
	}

	return photoURLs
}
