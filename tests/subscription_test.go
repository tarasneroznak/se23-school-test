package rate_test

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strconv"
	"testing"

	"github.com/gin-gonic/gin"

	"github.com/stretchr/testify/assert"

	subscription "main/resources/subscription"
)

func makeSubscribeRequest(data url.Values) (*httptest.ResponseRecorder, *http.Request) {
	bodyReader := bytes.NewReader([]byte(data.Encode()))
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/subscribe", bodyReader)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Content-Length", strconv.Itoa(len(data.Encode())))
	return w, req
}

func TestSubscribeRoute(t *testing.T) {
	gin.SetMode(gin.TestMode)
	app := gin.Default()
	rg := app.Group("/")
	subscription.Init(rg)
	data := url.Values{}
	data.Set("email", "test@test.com")
	w, req := makeSubscribeRequest(data)
	app.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)
}

func TestSubscribeRouteWithInvalidEmail(t *testing.T) {
	gin.SetMode(gin.TestMode)
	app := gin.Default()
	rg := app.Group("/")

	subscription.Init(rg)
	data := url.Values{}
	data.Set("email", "test")
	w, req := makeSubscribeRequest(data)
	app.ServeHTTP(w, req)
	assert.Equal(t, 400, w.Code)
}

func TestSubscribeRouteTwiceWithTheSameEmail(t *testing.T) {
	gin.SetMode(gin.TestMode)
	app := gin.Default()
	rg := app.Group("/")
	subscription.Init(rg)
	data := url.Values{}
	data.Set("email", "test1@test.com")
	w, req := makeSubscribeRequest(data)
	app.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)
	w, req = makeSubscribeRequest(data)
	app.ServeHTTP(w, req)
	assert.Equal(t, 409, w.Code)
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
