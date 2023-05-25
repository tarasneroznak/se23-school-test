package rate_test

import (
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/gin-gonic/gin"

	"github.com/stretchr/testify/assert"

	rate "main/resources/rate"
)

func TestRateRoute(t *testing.T) {
	gin.SetMode(gin.TestMode)
	app := gin.Default()
	rg := app.Group("/")

	rate.Init(rg)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/rate", nil)
	app.ServeHTTP(w, req)

	num, err := strconv.ParseInt(w.Body.String(), 10, 0)
	if err != nil {
		t.Fatalf("Response is not integer")
		return
	}

	assert.Equal(t, 200, w.Code)
	assert.GreaterOrEqual(t, num, int64(900000))
	assert.LessOrEqual(t, num, int64(1100000))
}
