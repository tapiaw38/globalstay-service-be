package appcontext

import (
	"github.com/tapiaw38/globalstay-service-be/internal/adapters/datasources"
	"github.com/tapiaw38/globalstay-service-be/internal/adapters/datasources/repositories"
	"github.com/tapiaw38/globalstay-service-be/internal/adapters/web/integrations"
	"github.com/tapiaw38/globalstay-service-be/internal/platform/config"
)

type Context struct {
	Repositories  *repositories.Repositories
	Integrations  *integrations.Integrations
	ConfigService *config.ConfigurationService
}

type Option func(*Context)

type Factory func(opts ...Option) *Context

func NewFactory(
	datasources *datasources.Datasources,
	integrations *integrations.Integrations,
	configService *config.ConfigurationService,
) func(opts ...Option) *Context {
	return func(opts ...Option) *Context {
		return &Context{
			Repositories:  repositories.CreateRepositories(datasources),
			Integrations:  integrations,
			ConfigService: configService,
		}
	}
}
