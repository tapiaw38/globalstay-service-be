package hotel

import (
	"net/http"

	"github.com/gin-gonic/gin"
	usecase "github.com/tapiaw38/globalstay-service-be/internal/usecases/hotel"
)

func NewUpdateHandler(usecase usecase.UpdateUsecase) gin.HandlerFunc {
	return func(c *gin.Context) {
		var input HotelInput
		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		id := c.Param("id")
		if id == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "missing id parameter"})
			return
		}

		updatedHotel, err := usecase.Execute(c, id, toHotel(input))
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, updatedHotel)
	}
}
