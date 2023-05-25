package subscription

import (
	"github.com/gin-gonic/gin"
)

func Init(router *gin.RouterGroup) {
	subscriptionRoutes(router)
}
