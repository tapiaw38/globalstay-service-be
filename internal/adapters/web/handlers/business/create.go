package business

import (
	"net/http"

	"github.com/gin-gonic/gin"
	usecase "github.com/tapiaw38/reservation-service-be/internal/usecases/business"
)

func NewCreateHandler(usecase usecase.CreateUsecase) gin.HandlerFunc {
	return func(c *gin.Context) {
		var input BusinessInput
		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		business, err := usecase.Create(c, toBusiness(input))
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusCreated, business)
	}
}
