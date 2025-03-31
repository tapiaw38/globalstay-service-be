package service

import (
	"github.com/gin-gonic/gin"
	usecase "github.com/tapiaw38/reservation-service-be/internal/usecases/service"
	"net/http"
)

func NewDeleteHandler(useCase usecase.DeleteUsecase) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		serviceID, err := useCase.Delete(c, id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, serviceID)
	}
}
