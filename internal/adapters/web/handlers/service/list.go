package service

import (
	"net/http"
	"net/url"

	"github.com/gin-gonic/gin"
	webutils "github.com/tapiaw38/globalstay-service-be/internal/platform/web"
	usecase "github.com/tapiaw38/globalstay-service-be/internal/usecases/service"
)

func NewListHandler(useCase usecase.ListUsecase) gin.HandlerFunc {
	return func(c *gin.Context) {
		filter := parseListFilter(c.Request.URL.Query())
		services, err := useCase.List(c, filter)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, services)
	}
}

func parseListFilter(queries url.Values) usecase.FilterOptions {
	return usecase.FilterOptions{
		HotelID: queries.Get("hotel_id"),
		Name:    queries.Get("name"),
		Limit:   webutils.ParseInt64QueryValue(queries.Get("limit")),
		Offset:  webutils.ParseInt64QueryValue(queries.Get("offset")),
	}
}
