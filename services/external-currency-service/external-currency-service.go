package external_currency_service

import (
	"math/rand"
)

func GetCurrencyRate(from string, to string) int {
	min := 900000
	max := 1100000
	return rand.Intn(max-min) + min
}
