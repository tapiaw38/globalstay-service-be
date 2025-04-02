package service

import (
	"net/http"

	"github.com/gin-gonic/gin"
	usecase "github.com/tapiaw38/globalstay-service-be/internal/usecases/service"
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
