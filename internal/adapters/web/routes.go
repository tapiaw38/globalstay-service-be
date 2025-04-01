package web

import (
	"github.com/gin-gonic/gin"
	"github.com/tapiaw38/reservation-service-be/internal/adapters/web/handlers/business"
	"github.com/tapiaw38/reservation-service-be/internal/adapters/web/handlers/service"
	"github.com/tapiaw38/reservation-service-be/internal/usecases"
)

func RegisterApplicationRoutes(app *gin.Engine, useCases *usecases.Usecases) {
	routeGroup := app.Group("/api/v1")

	routeGroup.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	routeGroup.POST("/business", business.NewCreateHandler(useCases.Business.CreateUsecase))
	routeGroup.GET("/business/:id", business.NewGetHandler(useCases.Business.GetUsecase))
	routeGroup.PATCH("/business/:id", business.NewUpdateHandler(useCases.Business.UpdateUsecase))
	routeGroup.GET("/businesses", business.NewListHandler(useCases.Business.ListUsecase))

	routeGroup.GET("/services", service.NewListHandler(useCases.Service.ListUsecase))
	routeGroup.POST("/services", service.NewCreateHandler(useCases.Service.CreateUsecase))
	routeGroup.GET("/services/:id", service.NewGetHandler(useCases.Service.GetUsecase))
	routeGroup.PUT("/services/:id", service.NewUpdateHandler(useCases.Service.UpdateUsecase))
	routeGroup.DELETE("/services/:id", service.NewDeleteHandler(useCases.Service.DeleteUsecase))
}
