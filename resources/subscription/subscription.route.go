package subscription

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func subscriptionRoutes(rg *gin.RouterGroup) {
	rg.POST("/subscribe", func(c *gin.Context) {
		c.Status(http.StatusOK)
	})
	rg.POST("/sendEmails", func(c *gin.Context) {
		c.Status(http.StatusOK)
	})
}
