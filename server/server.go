package server

import (
	"github.com/gin-gonic/gin"

	"github.com/lucassauro/go-api-boilerplate/utils"
	"github.com/lucassauro/go-api-boilerplate/controllers"
)

func SetupRouter() (router *gin.Engine) {
	router = gin.Default()

	router.GET("/echo", controllers.Echo)
	router.POST("/echo", controllers.Echo)
	router.OPTIONS("/echo", controllers.Echo)

	// HEALTH CHECK
	router.GET("/testdb", controllers.TestDB)
	router.GET("/testsdb", controllers.TestsDB)

	return router
}


func Server() error {
	router := SetupRouter()

	router.HandleMethodNotAllowed = true

	utils.GetNetworkAddress()

	router.SetTrustedProxies(nil)

	if err := router.Run(); err != nil {
		return err
	}

	return nil
}
