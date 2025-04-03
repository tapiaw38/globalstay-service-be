package hotel

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tapiaw38/globalstay-service-be/internal/usecases/hotel"
)

func NewScanPlacesJob(usecase hotel.ScanPlacesUsecase) gin.HandlerFunc {
	return func(c *gin.Context) {
		if err := usecase.Execute(c.Request.Context()); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, nil)
	}
}
