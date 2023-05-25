package rate

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func rateRoutes(rg *gin.RouterGroup) {
	rg.GET("/rate", func(c *gin.Context) {
		c.JSON(http.StatusOK, 1)
	})
}
