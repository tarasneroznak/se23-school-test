package subscription

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func subscriptionRoutes(rg *gin.RouterGroup) {
	rg.POST("/subscribe", subscribeController)
	rg.POST("/sendEmails", sendEmailsController)
}

// @ID subscribe
// @Summary Підписати емейл на отримання поточного курсу
// @Description Запит має перевірити, чи немає данної електронної адреси в поточній базі даних (файловій) і, в разі її відсутності, записувати її. Пізніше, за допомогою іншого запиту ми будемо відправляти лист на ті електронні адреси, які будуть в цій базі.
// @Tags subscription
// @Accept x-www-form-urlencoded
// @Produce json
// @Param email	formData string	true "Електронна адреса, яку потрібно підписати"
// @Success 200 "E-mail додано"
// @Failure	400	"BadRequest"
// @Failure	409	"Повертати, якщо e-mail вже є в базі даних (файловій)"
// @Router /subscribe [post]
func subscribeController(c *gin.Context) {
	c.Status(http.StatusOK)
}

// @ID sendEmails
// @Summary Відправити e-mail з поточним курсом на всі підписані електронні пошти.
// @Description Запит має отримувати актуальний курс BTC до UAH за допомогою third-party сервісу та відправляти його на всі електронні адреси, які були підписані раніше.
// @Tags subscription
// @Produce json
// @Success 200 "E-mailʼи відправлено"
// @Router /sendEmails [post]
func sendEmailsController(c *gin.Context) {
	c.Status(http.StatusOK)
}
