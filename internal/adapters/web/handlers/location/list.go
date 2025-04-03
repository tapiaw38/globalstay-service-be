package location

import (
	"net/http"

	"github.com/gin-gonic/gin"
	usecase "github.com/tapiaw38/globalstay-service-be/internal/usecases/location"
)

func NewListHandler(usecase usecase.ListUsecase) gin.HandlerFunc {
	return func(c *gin.Context) {
		output, err := usecase.Execute(c)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, output)
	}
}
