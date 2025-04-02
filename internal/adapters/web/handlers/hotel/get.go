package hotel

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tapiaw38/globalstay-service-be/internal/usecases/hotel"
)

func NewGetHandler(usecase hotel.GetUsecase) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		if id == "" {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "id is required",
			})
			return
		}

		hotel, err := usecase.Get(c, id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, hotel)
	}
}
