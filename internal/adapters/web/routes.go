package web

import (
	"github.com/gin-gonic/gin"
	"github.com/tapiaw38/globalstay-service-be/internal/adapters/web/handlers/hotel"
	"github.com/tapiaw38/globalstay-service-be/internal/adapters/web/handlers/location"
	"github.com/tapiaw38/globalstay-service-be/internal/adapters/web/handlers/service"
	hotel_job "github.com/tapiaw38/globalstay-service-be/internal/adapters/web/jobs/hotel"
	"github.com/tapiaw38/globalstay-service-be/internal/usecases"
)

func RegisterApplicationRoutes(app *gin.Engine, useCases *usecases.Usecases) {
	routeGroup := app.Group("/api/v1")

	routeGroup.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	routeGroup.POST("/hotel", hotel.NewCreateHandler(useCases.Hotel.CreateUsecase))
	routeGroup.GET("/hotel/:id", hotel.NewGetHandler(useCases.Hotel.GetUsecase))
	routeGroup.PATCH("/hotel/:id", hotel.NewUpdateHandler(useCases.Hotel.UpdateUsecase))
	routeGroup.GET("/hotels", hotel.NewListHandler(useCases.Hotel.ListUsecase))
	routeGroup.PATCH("/hotel/room/:hotel_id", hotel.NewUpdateRoomHandler(useCases.Hotel.UpdateRoomUsecase))

	routeGroup.POST("/location", location.NewCreateHandler(useCases.Location.CreateUsecase))
	routeGroup.PATCH("/location/:id", location.NewUpdateHandler(useCases.Location.UpdateUsecase))
	routeGroup.GET("/locations", location.NewListHandler(useCases.Location.ListUsecase))

	routeGroup.GET("/services", service.NewListHandler(useCases.Service.ListUsecase))
	routeGroup.POST("/services", service.NewCreateHandler(useCases.Service.CreateUsecase))
	routeGroup.GET("/services/:id", service.NewGetHandler(useCases.Service.GetUsecase))
	routeGroup.PUT("/services/:id", service.NewUpdateHandler(useCases.Service.UpdateUsecase))
	routeGroup.DELETE("/services/:id", service.NewDeleteHandler(useCases.Service.DeleteUsecase))
}

func RegisterJobRoutes(app *gin.Engine, useCases *usecases.Usecases) {
	routeGroup := app.Group("/api/v1")

	routeGroup.POST("/jobs/hotel/scan_places", hotel_job.NewScanPlacesJob(useCases.Hotel.ScanPlacesUsecase))
}
