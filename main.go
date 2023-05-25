package main

import (
	_ "main/docs"
	"main/resources/rate"
	"main/resources/subscription"

	"github.com/gin-gonic/gin"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

var app = gin.Default()

// @title GSES3 BTC application
// @version 1.0.0
// @host localhost:8080
// @BasePath /api
// @schemes http
func main() {
	router := app.Group("/api")

	rate.Init(router)
	subscription.Init(router)

	url := ginSwagger.URL("http://localhost:8080/api/swagger/doc.json")
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	app.Run(":8080")
}
