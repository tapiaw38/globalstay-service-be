package business

import (
	"net/http"

	"github.com/gin-gonic/gin"
	usecase "github.com/tapiaw38/reservation-service-be/internal/usecases/business"
)

func NewUpdateHandler(usecase usecase.UpdateUsecase) gin.HandlerFunc {
	return func(c *gin.Context) {
		var input BusinessInput
		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		id := c.Param("id")
		if id == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "missing id parameter"})
			return
		}

		updatedBusiness, err := usecase.Update(c, id, toBusiness(input))
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, updatedBusiness)
	}
}
