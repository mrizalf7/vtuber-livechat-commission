// -- Active: 1666923799638@@127.0.0.1@5432
package routes

import (
	"main/controller"
	"github.com/gin-gonic/gin"
)

func PriceListRoute(router *gin.Engine) {
	router.GET("/",controller.Home)
	router.GET("/PriceList", controller.GetPriceList)
	router.POST("/PriceList", controller.CreatePriceList)
	router.DELETE("/PriceList/:id", controller.DeletePriceList)
	router.PUT("/PriceList/:id", controller.UpdatePriceList)
}

func CommissionRoute(router *gin.Engine) {
	router.GET("/Commission", controller.GetCommission)
	router.POST("/Commission", controller.CreateCommission)
	router.DELETE("/Commission/:id", controller.DeleteCommission)
	router.PUT("/Commission/:id", controller.UpdateCommission)
}

func ProjectRoute(router *gin.Engine) {
	router.GET("/Project", controller.GetProject)
	router.POST("/Project", controller.CreateProject)
	router.DELETE("/Project/:id", controller.DeleteProject)
	router.PUT("/Project/:id", controller.UpdateProject)
}

func UserRoute(router *gin.Engine) {

	// router.GET("/User", controller.GetUser)
	// router.POST("/User", controller.CreateUser)
	// router.DELETE("/User/:id", controller.DeleteUser)
	// router.PUT("/User/:id", controller.UpdateUser)
	router.POST("/User",controller.CreateUser)
	router.GET("/User", controller.GetUser)
	router.POST("/User/Token",controller.GenerateToken)
}
