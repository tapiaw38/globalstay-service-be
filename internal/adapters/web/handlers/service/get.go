package service

import (
	"net/http"

	"github.com/gin-gonic/gin"
	usecase "github.com/tapiaw38/globalstay-service-be/internal/usecases/service"
)

func NewGetHandler(useCase usecase.GetUsecase) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		service, err := useCase.Get(c, id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, service)
	}
}
