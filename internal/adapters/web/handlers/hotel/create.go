package hotel

import (
	"net/http"

	"github.com/gin-gonic/gin"
	usecase "github.com/tapiaw38/globalstay-service-be/internal/usecases/hotel"
)

func NewCreateHandler(usecase usecase.CreateUsecase) gin.HandlerFunc {
	return func(c *gin.Context) {
		var input HotelInput
		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		hotel, err := usecase.Create(c, toHotel(input))
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusCreated, hotel)
	}
}
