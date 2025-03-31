package business

import (
	"net/http"
	"net/url"

	webutils "github.com/tapiaw38/reservation-service-be/internal/platform/web"
	usecase "github.com/tapiaw38/reservation-service-be/internal/usecases/business"

	"github.com/gin-gonic/gin"
)

func NewListHandler(usecase usecase.ListUsecase) gin.HandlerFunc {
	return func(c *gin.Context) {
		filter := parseListBusinessFilter(c.Request.URL.Query())
		businesses, err := usecase.List(c, filter)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, businesses)
	}
}

func parseListBusinessFilter(queries url.Values) usecase.FilterOptions {
	return usecase.FilterOptions{
		Name:        queries.Get("name"),
		Type:        queries.Get("type"),
		Description: queries.Get("description"),
		Active:      webutils.ParseBoolQueryValue(queries.Get("active")),
		Latitude:    webutils.ParseFloat64QueryValue(queries.Get("latitude")),
		Longitude:   webutils.ParseFloat64QueryValue(queries.Get("longitude")),
		Limit:       webutils.ParseInt64QueryValue(queries.Get("limit")),
		Offset:      webutils.ParseInt64QueryValue(queries.Get("offset")),
	}
}
