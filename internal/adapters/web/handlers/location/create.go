package location

import (
	"net/http"

	"github.com/gin-gonic/gin"
	usecase "github.com/tapiaw38/globalstay-service-be/internal/usecases/location"
)

func NewCreateHandler(usecase usecase.CreateUsecase) gin.HandlerFunc {
	return func(c *gin.Context) {
		var input LocationInput
		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		output, err := usecase.Execute(c, toLocationInput(input))
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusCreated, output)
	}
}
