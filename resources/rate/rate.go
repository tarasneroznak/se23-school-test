package rate

import (
	"github.com/gin-gonic/gin"
)

func Init(router *gin.RouterGroup) {
	rateRoutes(router)
}
