package main

import (
	"fmt"
	"main/config"
	"main/routes"
	"github.com/gin-gonic/gin"
)


func main() {

	router := gin.New()

	// router.Use(gin.Logger())

	// Recovery middleware recovers from any panics and writes a 500 if there was one.
	// router.Use(gin.Recovery())
	// router.Use(AuthRequired())
	config.Connect()
	routes.PriceListRoute(router)
	routes.ProjectRoute(router)
	routes.CommissionRoute(router)
	routes.UserRoute(router)
	fmt.Println("Running on 3000")
	router.Run(":3000")
}