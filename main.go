package main

import (
	"main/resources/rate"
	"main/resources/subscription"

	"github.com/gin-gonic/gin"
)

var app = gin.Default()

func main() {
	router := app.Group("/api")

	rate.Init(router)
	subscription.Init(router)

	app.Run(":8080")
}
