package main

import (
	"smartwallet-api/application/controllers"
	"github.com/gin-gonic/gin"
)

func registerRoutes() *gin.Engine {
	r := gin.Default()
	r.GET("/ping", controllers.Ping)	

	return r
}