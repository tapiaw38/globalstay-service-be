package appcontext

import (
	"github.com/tapiaw38/reservation-service-be/internal/adapters/datasources"
	"github.com/tapiaw38/reservation-service-be/internal/adapters/datasources/repositories"
	"github.com/tapiaw38/reservation-service-be/internal/platform/config"
)

type Context struct {
	Repositories  *repositories.Repositories
	ConfigService *config.ConfigurationService
}

type Option func(*Context)

type Factory func(opts ...Option) *Context

func NewFactory(
	datasources *datasources.Datasources,
	configService *config.ConfigurationService,
) func(opts ...Option) *Context {
	return func(opts ...Option) *Context {
		return &Context{
			Repositories:  repositories.CreateRepositories(datasources),
			ConfigService: configService,
		}
	}
}
