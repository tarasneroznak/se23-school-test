package subscription

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func subscriptionRoutes(rg *gin.RouterGroup) {
	rg.POST("/subscribe", subscribeController)
	rg.POST("/sendEmails", sendEmailsController)
}

type SubscribeForm struct {
	Email string `form:"email" validate:"email,required"`
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
	var form SubscribeForm
	if err := c.ShouldBind(&form); err != nil {
		c.Status(http.StatusBadRequest)
		return
	}
	validate := validator.New()
	if err := validate.Struct(form); err != nil {
		c.Status(http.StatusBadRequest)
		return
	}
	err := subscribe(form.Email)
	if err != nil {
		c.Status(http.StatusConflict)
		return
	}
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
	err := sendEmails()
	if err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}
	c.Status(http.StatusOK)
}
