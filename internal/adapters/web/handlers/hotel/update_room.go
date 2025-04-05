package hotel

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tapiaw38/globalstay-service-be/internal/usecases/hotel"
)

func NewUpdateRoomHandler(usecase hotel.UpdateRoomUsecase) gin.HandlerFunc {
	return func(c *gin.Context) {
		var input RoomInput
		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
			return
		}

		hotelID := c.Param("hotel_id")
		if hotelID == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Hotel ID is required"})
			return
		}

		if input.Number == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Room number is required"})
			return
		}

		output, err := usecase.Execute(c.Request.Context(), hotelID, toRoom(input))
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		if output == nil {
			c.JSON(404, gin.H{"error": "Room not found"})
			return
		}

		c.JSON(http.StatusOK, output)
	}
}
