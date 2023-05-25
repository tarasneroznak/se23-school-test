package rate_test

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"

	"github.com/stretchr/testify/assert"

	subscription "main/resources/subscription"
)

func TestSubscribeRoute(t *testing.T) {
	gin.SetMode(gin.TestMode)
	app := gin.Default()
	rg := app.Group("/")

	subscription.Init(rg)

	jsonBody := []byte(`{"email": "test@test.com"}`)
	bodyReader := bytes.NewReader(jsonBody)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/subscribe", bodyReader)
	app.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
}

func TestSendEmailsRoute(t *testing.T) {
	gin.SetMode(gin.TestMode)
	app := gin.Default()
	rg := app.Group("/")

	subscription.Init(rg)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/sendEmails", nil)
	app.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
}
