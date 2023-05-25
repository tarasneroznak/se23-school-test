package rate

import (
	external_currency_service "main/services/external-currency-service"
)

func getBTCtoUAH() int {
	return external_currency_service.GetCurrencyRate("BTC", "UAH")
}
