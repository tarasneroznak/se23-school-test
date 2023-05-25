package rate

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func rateRoutes(rg *gin.RouterGroup) {
	rg.GET("/rate", rateController)
}

// @ID rate
// @Summary Отримати поточний курс BTC до UAH
// @Description Запит має повертати поточний курс BTC до UAH використовуючи будь-який third party сервіс з публічним АРІ
// @Tags rate
// @Produce json
// @Success 200 {integer} string "Повертається актуальний курс BTC до UAH"
// @Failure	400	"Invalid status value"
// @Router /rate [get]
func rateController(c *gin.Context) {
	rate := getBTCtoUAH()
	c.JSON(http.StatusOK, rate)
}
