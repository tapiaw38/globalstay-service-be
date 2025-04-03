package hotel

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tapiaw38/globalstay-service-be/internal/usecases/hotel"
)

func NewScanPlacesJob(usecase hotel.ScanPlacesUsecase) gin.HandlerFunc {
	return func(c *gin.Context) {
		var input ScanPlacesInput
		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if err := usecase.Execute(c, input.Location, input.Radius, input.Types); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, nil)
	}
}
