package places

import (
	"log"

	"github.com/tapiaw38/globalstay-service-be/internal/platform/config"
	"googlemaps.github.io/maps"
)

var (
	googleMapsRadius uint = 5000
)

type (
	Integration interface {
		GetPlaces(float64, float64, uint, maps.PlaceType) ([]Place, error)
	}

	integration struct {
		client *maps.Client
	}
)

func NewGoogleMapsClient(cfg *config.ConfigurationService) Integration {
	client, err := maps.NewClient(maps.WithAPIKey(cfg.GCP.MapsKey))
	if err != nil {
		log.Println(err)
		return nil
	}

	return integration{
		client: client,
	}
}
