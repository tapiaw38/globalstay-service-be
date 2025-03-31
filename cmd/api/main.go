package main

import (
	"context"
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/tapiaw38/reservation-service-be/internal/adapters/datasources"
	"github.com/tapiaw38/reservation-service-be/internal/adapters/web"
	"github.com/tapiaw38/reservation-service-be/internal/platform/appcontext"
	"github.com/tapiaw38/reservation-service-be/internal/platform/config"
	"github.com/tapiaw38/reservation-service-be/internal/platform/nosql"
	"github.com/tapiaw38/reservation-service-be/internal/usecases"
)

func main() {
	scope := config.GetScope()

	log.Printf("scope identifier: %s", scope)

	if err := initConfig(); err != nil {
		panic(err)
	}

	if err := run(); err != nil {
		panic(err)
	}
}

func run() error {
	configService := config.GetConfigService()

	businessClient, err := nosql.NewClient(configService.NoSQLConfig.Business)
	if err != nil {
		return err
	}
	defer func(businessClient nosql.Client, ctx context.Context) {
		_ = businessClient.Disconnect(ctx)
	}(businessClient, context.TODO())

	servicesClient, err := nosql.NewClient(configService.NoSQLConfig.Services)
	if err != nil {
		return err
	}
	defer func(servicesClient nosql.Client, ctx context.Context) {
		_ = servicesClient.Disconnect(ctx)
	}(servicesClient, context.TODO())

	reservationsClient, err := nosql.NewClient(configService.NoSQLConfig.Reservations)
	if err != nil {
		return err
	}
	defer func(reservationsClient nosql.Client, ctx context.Context) {
		_ = reservationsClient.Disconnect(ctx)
	}(reservationsClient, context.TODO())

	if configService.ServerConfig.GinMode == config.DebugMode {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	app := gin.Default()
	ginConfig := cors.DefaultConfig()
	ginConfig.AllowOrigins = []string{"*"}
	ginConfig.AllowCredentials = true
	ginConfig.AllowMethods = []string{"*"}
	ginConfig.AllowHeaders = []string{"*"}
	ginConfig.ExposeHeaders = []string{"*"}
	app.Use(cors.New(ginConfig))

	bootstrap(app, businessClient, servicesClient, reservationsClient, &configService)

	return app.Run(":" + configService.ServerConfig.Port)
}

func bootstrap(
	app *gin.Engine,
	businessClient nosql.Client,
	servicesClient nosql.Client,
	reservationsClient nosql.Client,
	configService *config.ConfigurationService,
) {
	datasources := datasources.CreateDatasources(businessClient, servicesClient, reservationsClient)
	contextFactory := appcontext.NewFactory(datasources, configService)
	useCases := usecases.CreateUsecases(contextFactory)
	web.RegisterApplicationRoutes(app, useCases)
}
