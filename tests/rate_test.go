package rate_test

import (
	"net/http"
	"net/http/httptest"
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

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "1", w.Body.String())
}
