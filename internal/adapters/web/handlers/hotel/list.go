package hotel

import (
	"net/http"
	"net/url"

	webutils "github.com/tapiaw38/globalstay-service-be/internal/platform/web"
	usecase "github.com/tapiaw38/globalstay-service-be/internal/usecases/hotel"

	"github.com/gin-gonic/gin"
)

func NewListHandler(usecase usecase.ListUsecase) gin.HandlerFunc {
	return func(c *gin.Context) {
		filter := parseListHotelFilter(c.Request.URL.Query())
		hoteles, err := usecase.Execute(c, filter)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, hoteles)
	}
}

func parseListHotelFilter(queries url.Values) usecase.FilterOptions {
	return usecase.FilterOptions{
		Name:         queries.Get("name"),
		Type:         queries.Get("type"),
		Description:  queries.Get("description"),
		Active:       webutils.ParseBoolQueryValue(queries.Get("active")),
		Latitude:     webutils.ParseFloat64QueryValue(queries.Get("latitude")),
		Longitude:    webutils.ParseFloat64QueryValue(queries.Get("longitude")),
		LocationName: queries.Get("location_name"),
		Limit:        webutils.ParseInt64QueryValue(queries.Get("limit")),
		Offset:       webutils.ParseInt64QueryValue(queries.Get("offset")),
	}
}
