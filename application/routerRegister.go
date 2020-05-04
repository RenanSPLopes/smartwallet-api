package main

import (
	"smartwallet-api/application/controllers"

	"github.com/gin-gonic/gin"
)

func registerRoutes() *gin.Engine {
	r := gin.Default()
	r.GET("/ping", controllers.Ping)

	marketDataController := ProvideMarketDataController()
	r.GET("/marketdata", marketDataController.GetAll)
	r.GET("/marketdata/:id", marketDataController.GetById)
	
	return r
}
