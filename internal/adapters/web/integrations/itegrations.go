package integrations

import (
	"github.com/tapiaw38/globalstay-service-be/internal/adapters/web/integrations/places"
	"github.com/tapiaw38/globalstay-service-be/internal/platform/config"
)

type Integrations struct {
	Places places.Integration
}

func CreateIntegrations(cfg *config.ConfigurationService) *Integrations {
	return &Integrations{
		Places: places.NewGoogleMapsClient(cfg),
	}
}
