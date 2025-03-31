package service

import (
	"github.com/gin-gonic/gin"
	usecase "github.com/tapiaw38/reservation-service-be/internal/usecases/service"
	"net/http"
)

func NewCreateHandler(useCase usecase.CreateUsecase) gin.HandlerFunc {
	return func(c *gin.Context) {
		var input serviceInput
		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		service, err := useCase.Create(c, toService(input))
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusCreated, service)
	}
}
