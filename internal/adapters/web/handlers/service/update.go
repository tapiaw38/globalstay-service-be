package service

import (
	"github.com/gin-gonic/gin"
	usecase "github.com/tapiaw38/reservation-service-be/internal/usecases/service"
	"net/http"
)

func NewUpdateHandler(useCase usecase.UpdateUsecase) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		if id == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "id is required"})
			return
		}

		var input serviceInput
		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		service, err := useCase.Update(c, id, toService(input))
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, service)
	}
}
