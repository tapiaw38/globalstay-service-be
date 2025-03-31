package service

import (
	"github.com/gin-gonic/gin"
	webutils "github.com/tapiaw38/reservation-service-be/internal/platform/web"
	usecase "github.com/tapiaw38/reservation-service-be/internal/usecases/service"
	"net/http"
	"net/url"
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
		BusinessID: queries.Get("business_id"),
		Name:       queries.Get("name"),
		Limit:      webutils.ParseInt64QueryValue(queries.Get("limit")),
		Offset:     webutils.ParseInt64QueryValue(queries.Get("offset")),
	}
}
